package client_test

import (
	"block-parser-go/internal/client"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetClientManager(t *testing.T) {
	clientManager := client.GetClientManager()

	assert.NotNil(t, clientManager, "Expected client manager to be not nil")
}

func TestClientManagerCreatesClients(t *testing.T) {
	clientManager := client.GetClientManager()

	client, err := clientManager.GetClient(client.Mainnet)

	assert.NotNil(t, client, "Expected client to be not nil")
	assert.NoError(t, err, "Expected error to be nil")

	if client != nil {
		defer clientManager.Close()
	}
}

// TODO: finish this test
func TestClientManagerClosesAllClients(t *testing.T) {
	t.Skip("Not implemented properly yet")
	clientManager := client.GetClientManager()

	mainnetClient, err := clientManager.GetClient(client.Mainnet)
	arbitrumClient, err := clientManager.GetClient(client.Arbitrum)

	assert.NotNil(t, mainnetClient, "Expected client to be not nil")
	assert.NotNil(t, arbitrumClient, "Expected client to be not nil")
	assert.NoError(t, err, "Expected error to be nil")

	clientManager.Close()
	assert.Nil(t, mainnetClient, "Expected client to be nil")
	assert.Nil(t, arbitrumClient, "Expected client to be nil")
}
