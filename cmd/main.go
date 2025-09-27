package main

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/0xcompose/block-parser-go/infura"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/joho/godotenv"
)

const chainId = 1

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}

	rpcUrl := infura.GetInfuraRpcUrl(chainId)
	fmt.Println("RPC URL: ", rpcUrl)
	client, err := rpc.Dial(rpcUrl)

	if err != nil {
		log.Fatalf("Failed to connect to Alchemy: %v", err)
	}

	fmt.Println("Successfully connected to Alchemy!")

	var result string

	err = client.Call(&result, "eth_blockNumber")

	if err != nil {
		log.Fatalf("Failed to get block number: %v", err)
	}

	fmt.Println("Block number: ", result)
	fmt.Printf("%x \n", result)
	blockNumber, err := strconv.ParseUint(result, 16, 64)

	if err != nil {
		log.Fatalf("Failed to parse block number: %v", err)
	}

	fmt.Printf("%d \n", blockNumber)

	defer client.Close()
}

func subscribeForNewBlocks(client *rpc.Client) {
	subscription, err := client.EthSubscribe(context.Background(), "newHeads", nil)

	if err != nil {
		log.Fatalf("Failed to subscribe to new heads: %v", err)
	}

	fmt.Println("Successfully subscribed to new heads!")
	defer subscription.Unsubscribe()
}

func GetSupportedModules(client *rpc.Client) {
	modules, err := client.SupportedModules()

	if err != nil {
		log.Fatalf("Failed to get supported modules: %v", err)
	}
	fmt.Println("Supported modules: ", modules)

}
