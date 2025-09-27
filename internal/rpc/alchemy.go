package rpc

import (
	"os"

	"block-parser-go/internal/config"
)

const alchemyApiKeyEnvName = "ALCHEMY_API_KEY"

func GetAlchemyRpcUrl(chainId int) string {
	return "https://" + GetAlchemyChainCode(chainId) + ".g.alchemy.com/v2/" + os.Getenv(alchemyApiKeyEnvName)
}

func GetAlchemyChainCode(chainId int) string {
	switch chainId {
	case 1:
		return "eth-mainnet"
	case 42161:
		return "arb-mainnet"
	default:
		panic("Unsupported chain ID")
	}
}

func GetAlchemyApiKey() string {
	return config.GetEnvValueSafely(alchemyApiKeyEnvName, alchemyApiKeyEnvName+" is not set")
}
