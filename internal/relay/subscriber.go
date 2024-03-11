package relay

import (
	"context"

	"github.com/neutron-org/neutron/x/interchainqueries/types"
	neutrontypes "github.com/neutron-org/neutron/x/interchainqueries/types"
)

// Subscriber is an interface that subscribes to Neutron and provides chain data in real time.
type Subscriber interface {
	// returns the neutron registered query by its ID
	GetNeutronRegisteredQuery(ctx context.Context, queryId string) (*neutrontypes.RegisteredQuery, error)
}

// MessageKV contains params of a KV interchain query.
type MessageKV struct {
	// QueryId is the ID of the query.
	QueryId uint64
	// KVKeys is the query parameter that describes keys list to be retrieved.
	KVKeys types.KVKeys
}

// MessageTX contains params of a TX interchain query.
type MessageTX struct {
	// QueryId is the ID of the query.
	QueryId uint64
	// TransactionsFilter is the query parameter that describes conditions for transactions search.
	TransactionsFilter string
}
