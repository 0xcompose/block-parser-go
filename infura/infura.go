package infura

import "os"

func GetInfuraRpcUrl(chainId int) string {
	return "https://" + GetInfuraChainCode(chainId) + ".infura.io/v3/" + os.Getenv("INFURA_API_KEY")
}

func GetInfuraChainCode(chainId int) string {
	switch chainId {
	case 1:
		return "mainnet"
	case 42161:
		return "arbitrum-mainnet"
	default:
		panic("Unsupported chain ID")
	}
}
