package client

import (
	"errors"

	"github.com/ethereum/go-ethereum/rpc"
)

type ChainID uint16

const (
	Mainnet  ChainID = 1
	Arbitrum ChainID = 42161
)

func NewClient(chainId ChainID) (*rpc.Client, error) {
	rpcUrl, ok := rpcUrls[chainId]

	if !ok {
		return nil, errors.New("unsupported chain ID")
	}

	apiKey, err := GetAlchemyApiKey()

	if err != nil {
		return nil, err
	}

	rpcUrl += apiKey

	client, err := rpc.Dial(rpcUrl)

	if err != nil {
		return nil, err
	}

	return client, nil
}
