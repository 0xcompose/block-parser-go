package rpc

import (
	"os"

	"block-parser-go/internal/config"
)

const alchemyApiKeyEnvName = "ALCHEMY_API_KEY"

func GetAlchemyRpcUrl(chainId ChainID) string {
	return "https://" + GetAlchemyChainCode(chainId) + ".g.alchemy.com/v2/" + os.Getenv(alchemyApiKeyEnvName)
}

func GetAlchemyChainCode(chainId ChainID) string {
	switch chainId {
	case Mainnet:
		return "eth-mainnet"
	case Arbitrum:
		return "arb-mainnet"
	default:
		panic("Unsupported chain ID")
	}
}

func GetAlchemyApiKey() (string, error) {
	key, err := config.GetEnv(alchemyApiKeyEnvName, alchemyApiKeyEnvName+" is not set")

	if err != nil {
		return "", err
	}

	return key, nil
}
