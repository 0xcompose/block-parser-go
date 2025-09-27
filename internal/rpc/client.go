package rpc

import (
	"github.com/ethereum/go-ethereum/rpc"
)

type ChainID int

const (
	Mainnet  ChainID = 1
	Arbitrum ChainID = 42161
)

var RPC_URLS = map[ChainID]string{
	Mainnet:  "https://eth-mainnet.g.alchemy.com/v2/",
	Arbitrum: "https://arb-mainnet.g.alchemy.com/v2/",
}

func GetClient(chainId ChainID) *rpc.Client {
	rpcUrl, ok := RPC_URLS[chainId]

	if !ok {
		panic("Unsupported chain ID")
	}

	rpcUrl += GetAlchemyApiKey()

	client, err := rpc.Dial(rpcUrl)

	if err != nil {
		panic(err)
	}

	return client
}
