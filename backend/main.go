package main

import (
	contracts "backend/contracts"
	"backend/logger"
	"backend/token"
	"context"
	"log"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
)

func main() {
	ctx := context.Background()
	_, err := logger.SetupLogger()
	if err != nil {
		log.Fatal("unable to setup logger", err.Error())
	}

	// set url and contract address
	url := "http://127.0.0.1:8545"
	contractHexAddress := "0x5FbDB2315678afecb367f032d93F642f64180aa3"

	// Connect to the Ethereum client
	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal("unable to connect", url, err.Error())
	}

	// Convert the contract address from hex string to common.Address type
	contractAddress := common.HexToAddress(contractHexAddress)

	// Create a new instance of the token contract
	instance, err := token.NewToken(contractAddress, client)
	if err != nil {
		logger.Error(ctx, "error while connecting to contract", "error", err.Error())
		return
	}

	// fetch chainID
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		logger.Error(ctx, "error while fetching chain_id", "error", err.Error())
		return
	}

	// fetch token name from contract
	name, err := instance.Name(nil)
	if err != nil {
		logger.Error(ctx, "error while fetching contract name", "error", err.Error())
		return
	}
	

	logger.Info(ctx, "contracted connected", "name", name, "chain_id", chainID)

	r := mux.NewRouter()

	tokenContract := contracts.NewTokenService(instance)

	r.Handle("/purchase", contracts.HandlePurchaseToken(tokenContract)).Methods("POST")

	port := ":8080"
	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(port, r))
}
