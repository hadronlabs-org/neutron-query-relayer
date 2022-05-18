package chain

import (
	"context"
	"fmt"
	"github.com/cosmos/cosmos-sdk/api/tendermint/abci"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/cosmos/cosmos-sdk/types"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/lidofinance/cosmos-query-relayer/internal/config"
	rpcclient "github.com/tendermint/tendermint/rpc/client"
)

var mode = signing.SignMode_SIGN_MODE_DIRECT

type TxSubmitter struct {
	ctx           context.Context
	baseTxf       tx.Factory
	codec         Codec
	rpcClient     rpcclient.Client
	chainID       string
	addressPrefix string
	signKeyName   string
}

func TestKeybase(chainID string, keyringRootDir string, codec Codec) (keyring.Keyring, error) {
	keybase, err := keyring.New(chainID, "test", keyringRootDir, nil, codec.Marshaller)
	if err != nil {
		return keybase, err
	}

	return keybase, nil
}

func NewTxSubmitter(ctx context.Context, rpcClient rpcclient.Client, codec Codec, keybase keyring.Keyring, cfg config.CosmosQueryRelayerConfig) (*TxSubmitter, error) {
	lidoCfg := cfg.LidoChain
	baseTxf := tx.Factory{}.
		WithKeybase(keybase).
		WithSignMode(mode).
		WithTxConfig(codec.TxConfig).
		WithChainID(lidoCfg.ChainID).
		WithGasAdjustment(lidoCfg.GasAdjustment).
		WithGasPrices(lidoCfg.GasPrices)

	return &TxSubmitter{
		ctx:           ctx,
		codec:         codec,
		baseTxf:       baseTxf,
		rpcClient:     rpcClient,
		chainID:       lidoCfg.ChainID,
		addressPrefix: lidoCfg.ChainPrefix,
		signKeyName:   lidoCfg.Keyring.SignKeyName,
	}, nil
}

// BuildAndSendTx builds transaction with calculated gas and fees params, signs it and submits to chain
func (cc *TxSubmitter) BuildAndSendTx(sender string, msgs []types.Msg) error {
	account, err := cc.QueryAccount(sender)
	if err != nil {
		return err
	}

	txf := cc.baseTxf.
		WithAccountNumber(account.AccountNumber).
		WithSequence(account.Sequence)

	//gasNeeded, err := cc.calculateGas(txf, msgs...)
	// FIXME: simulate query is broken for now. Turn on after migrating back to cosmos-sdk v45
	gasNeeded := uint64(2000000)
	if err != nil {
		return err
	}

	txf = txf.WithGas(gasNeeded)

	bz, err := cc.buildTxBz(txf, msgs, sender, gasNeeded)
	if err != nil {
		return err
	}
	res, err := cc.rpcClient.BroadcastTxSync(cc.ctx, bz)

	fmt.Printf("Broadcast result: code=%+v log=%v err=%+v hash=%b", res.Code, res.Log, err, res.Hash)

	if res.Code == 0 {
		return nil
	} else {
		return fmt.Errorf("error broadcasting transaction with log=%s", res.Log)
	}
}

// QueryAccount returns BaseAccount for given account address
func (cc *TxSubmitter) QueryAccount(address string) (*authtypes.BaseAccount, error) {
	request := authtypes.QueryAccountRequest{Address: address}
	req, err := request.Marshal()
	if err != nil {
		return nil, err
	}
	simQuery := abci.RequestQuery{
		Path: "/cosmos.auth.v1beta1.Query/Account",
		Data: req,
	}
	res, err := cc.rpcClient.ABCIQueryWithOptions(cc.ctx, simQuery.Path, simQuery.Data, rpcclient.DefaultABCIQueryOptions)
	if err != nil {
		return nil, err
	}

	if res.Response.Code != 0 {
		return nil, fmt.Errorf("error fetching account with address=%s log=%s", address, res.Response.Log)
	}

	var response authtypes.QueryAccountResponse
	if err := response.Unmarshal(res.Response.Value); err != nil {
		return nil, err
	}

	var account authtypes.BaseAccount
	err = account.Unmarshal(response.Account.Value)

	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (cc *TxSubmitter) buildTxBz(txf tx.Factory, msgs []types.Msg, feePayerAddress string, gasAmount uint64) ([]byte, error) {
	txBuilder := cc.codec.TxConfig.NewTxBuilder()
	err := txBuilder.SetMsgs(msgs...)
	if err != nil {
		fmt.Printf("set msgs failure")
		return nil, err
	}

	txBuilder.SetGasLimit(gasAmount)
	//txBuilder.SetMemo("bob to alice")

	feePayerBz, err := types.GetFromBech32(feePayerAddress, cc.addressPrefix)
	if err != nil {
		return nil, err
	}
	txBuilder.SetFeePayer(feePayerBz)
	// TODO: shouldn't set it like this. use gas limit and gas prices
	txBuilder.SetFeeAmount(types.NewCoins(types.NewInt64Coin("stake", 500000)))
	//txBuilder.SetFeeGranter()
	//txBuilder.SetTimeoutHeight(...)

	fmt.Printf("\nAbout to sign with txf: %+v\n\n", txf)
	err = tx.Sign(txf, cc.signKeyName, txBuilder, true)

	if err != nil {
		return nil, err
	}

	bz, err := cc.codec.TxConfig.TxEncoder()(txBuilder.GetTx())
	return bz, err
}

func (cc *TxSubmitter) calculateGas(txf tx.Factory, msgs ...types.Msg) (uint64, error) {
	simulation, err := cc.buildSimTx(txf, msgs...)
	if err != nil {
		return 0, err
	}
	// We then call the Simulate method on this client.
	simQuery := abci.RequestQuery{
		Path: "/cosmos.tx.v1beta1.Service/Simulate",
		Data: simulation,
	}
	res, err := cc.rpcClient.ABCIQueryWithOptions(cc.ctx, simQuery.Path, simQuery.Data, rpcclient.DefaultABCIQueryOptions)
	if err != nil {
		return 0, err
	}

	var simRes txtypes.SimulateResponse

	if err := simRes.Unmarshal(res.Response.Value); err != nil {
		return 0, err
	}
	if simRes.GasInfo == nil {
		return 0, fmt.Errorf("no result in simulation response with log=%s code=%d", res.Response.Log, res.Response.Code)
	}

	return uint64(txf.GasAdjustment() * float64(simRes.GasInfo.GasUsed)), nil
}

// buildSimTx creates an unsigned tx with an empty single signature and returns
// the encoded transaction or an error if the unsigned transaction cannot be built.
func (cc *TxSubmitter) buildSimTx(txf tx.Factory, msgs ...types.Msg) ([]byte, error) {
	txb, err := cc.baseTxf.BuildUnsignedTx(msgs...)
	if err != nil {
		return nil, err
	}

	// Create an empty signature literal as the ante handler will populate with a
	// sentinel pubkey.
	sig := signing.SignatureV2{
		PubKey: &secp256k1.PubKey{},
		Data: &signing.SingleSignatureData{
			SignMode: cc.baseTxf.SignMode(),
		},
		Sequence: txf.Sequence(),
	}
	if err := txb.SetSignatures(sig); err != nil {
		return nil, err
	}

	bz, err := cc.codec.TxConfig.TxEncoder()(txb.GetTx())
	if err != nil {
		return nil, nil
	}
	simReq := txtypes.SimulateRequest{TxBytes: bz}
	return simReq.Marshal()
}
