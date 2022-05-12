package config

import (
	"fmt"
	"github.com/vrischmann/envconfig"
	"time"
)

type CosmosQueryRelayerConfig struct {
	//	TODO: full configuration
	LidoChain struct {
		RPCAddress string `envconfig:"default=tcp://127.0.0.1:26657"`
		//RPCAddress string `envconfig:"default=tcp://public-node.terra.dev:26657"` // for tests only
		Timeout     time.Duration `envconfig:"default=5s"`
		ChainPrefix string        `envconfig:"default=terra"`
	}
	TargetChain struct {
		Timeout time.Duration `envconfig:"default=5s"`
		//RPCAddress string `envconfig:"default=tcp://rpc.cosmos.network:26657"`
		RPCAddress string `envconfig:"default=tcp://public-node.terra.dev:26657"`
		//RPCAddress string `envconfig:"default=http://167.99.25.150:26657"`
		//RPCAddress string `envconfig:"default=tcp://127.0.0.1:26657"`
		ChainID string `envconfig:"default=columbus-5"`
		//ChainID     string `envconfig:"default=testnet"`
		ChainPrefix string `envconfig:"default=terra"`

		Keyring struct {
			Backend   string `envconfig:"default=test"`
			GasPrices string `envconfig:"default=0.01uatom"` // should not be in keyring struct
		}
	}
}

func NewCosmosQueryRelayerConfig() (CosmosQueryRelayerConfig, error) {
	config := CosmosQueryRelayerConfig{}
	if err := envconfig.Init(&config); err != nil {
		return config, fmt.Errorf("failed to init config: %w", err)
	}

	return config, nil
}
