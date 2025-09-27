package client

// TODO: should move those to the `config` package
var rpcUrls = map[ChainID]string{
	Mainnet:  GetAlchemyRpcUrl(Mainnet),
	Arbitrum: GetAlchemyRpcUrl(Arbitrum),
}
