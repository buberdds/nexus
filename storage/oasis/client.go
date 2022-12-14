// Package oasis implements the source storage interface
// backed by oasis-node.
package oasis

import (
	"context"

	"github.com/oasisprotocol/oasis-core/go/common/crypto/signature"
	config "github.com/oasisprotocol/oasis-sdk/client-sdk/go/config"
	connection "github.com/oasisprotocol/oasis-sdk/client-sdk/go/connection"
	runtimeSignature "github.com/oasisprotocol/oasis-sdk/client-sdk/go/crypto/signature"
)

const (
	moduleName = "storage_oasis"
)

// ClientFactory supports connections to an oasis-node instance.
type ClientFactory struct {
	connection *connection.Connection
	network    *config.Network
}

// NewClientFactory creates a new oasis-node client factory.
func NewClientFactory(ctx context.Context, network *config.Network) (*ClientFactory, error) {
	connection, err := connection.ConnectNoVerify(ctx, network)
	if err != nil {
		return nil, err
	}

	return &ClientFactory{
		connection: &connection,
		network:    network,
	}, nil
}

// Consensus creates a new ConsensusClient.
func (cf *ClientFactory) Consensus() (*ConsensusClient, error) {
	ctx := context.Background()

	connection := *cf.connection
	client := connection.Consensus()

	// Configure chain context for all signatures using chain domain separation.
	signature.SetChainContext(cf.network.ChainContext)

	c := &ConsensusClient{
		client:  client,
		network: cf.network,
	}
	doc, err := c.GenesisDocument(ctx)
	if err != nil {
		return nil, err
	}
	c.genesisHeight = doc.Height

	return c, nil
}

// Runtime creates a new RuntimeClient.
func (cf *ClientFactory) Runtime(runtimeID string) (*RuntimeClient, error) {
	ctx := context.Background()

	connection := *cf.connection
	client := connection.Runtime(&config.ParaTime{
		ID: runtimeID,
	})

	// Configure chain context for all signatures using chain domain separation.
	signature.SetChainContext(cf.network.ChainContext)

	info, err := client.GetInfo(ctx)
	if err != nil {
		return nil, err
	}

	rtCtx := runtimeSignature.DeriveChainContext(info.ID, cf.network.ChainContext)
	return &RuntimeClient{
		client:  client,
		network: cf.network,
		info:    info,
		rtCtx:   rtCtx,
	}, nil
}
