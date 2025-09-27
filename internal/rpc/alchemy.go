package rpc

import "os"

func GetAlchemyRpcUrl(chainId int) string {
	return "https://" + GetAlchemyChainCode(chainId) + ".g.alchemy.com/v2/" + os.Getenv("ALCHEMY_API_KEY")
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
