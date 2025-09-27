package client_test

import (
	"block-parser-go/internal/client"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetClientMainnet(t *testing.T) {
	client, err := client.NewClient(client.Mainnet)

	assert.NotNil(t, client, "Expected client to be not nil")
	defer client.Close()

	assert.NoError(t, err, "Expected error to be nil")
}

func TestGetClientArbitrum(t *testing.T) {
	client, err := client.NewClient(client.Arbitrum)

	assert.NotNil(t, client, "Expected client to be not nil")
	defer client.Close()

	assert.NoError(t, err, "Expected error to be nil")
}

func TestGetClientWithInvalidChainID(t *testing.T) {
	const InvalidChainID client.ChainID = 0
	client, err := client.NewClient(InvalidChainID)

	assert.Error(t, err, "Expected error to be not nil")

	assert.EqualError(t, err, "unsupported chain ID")

	assert.Nil(t, client, "Expected client to be nil")

	if client != nil {
		defer client.Close()
	}
}
