package rpc

import (
	"github.com/ethereum/go-ethereum/rpc"
)

type ChainID int

const (
	Mainnet  ChainID = 1
	Arbitrum ChainID = 42161
)

// TODO: should move those to the `config` package
var rpcUrls = map[ChainID]string{
	Mainnet:  GetAlchemyRpcUrl(Mainnet),
	Arbitrum: GetAlchemyRpcUrl(Arbitrum),
}

func GetClient(chainId ChainID) *rpc.Client {
	rpcUrl, ok := rpcUrls[chainId]

	if !ok {
		panic("Unsupported chain ID")
	}

	apiKey, err := GetAlchemyApiKey()

	if err != nil {
		panic(err)
	}

	rpcUrl += apiKey

	client, err := rpc.Dial(rpcUrl)

	if err != nil {
		panic(err)
	}

	return client
}
