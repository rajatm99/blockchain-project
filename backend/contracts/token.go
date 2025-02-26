package contracts

import (
	"backend/logger"
	"backend/token"
	"context"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type tokenService struct {
	tokenContract *token.Token
}

type TokenService interface {
	PurchaseToken(context.Context, PurchaseTokenRequest) error
}

func NewTokenService(
	tokenContract *token.Token,
) TokenService {
	return tokenService{
		tokenContract: tokenContract,
	}
}

// this function mints new coin and assigns it to the address passed in the request
// minting can only be done by owner of the contract
func (t tokenService) PurchaseToken(ctx context.Context, request PurchaseTokenRequest) (err error) {

	// private key of owner
	privateKeyString := "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	privateKeyHex, err := hex.DecodeString(privateKeyString)
	if err != nil {
		logger.Error(ctx, "error while converting private key string to hex", "error", err.Error())
		return
	}


	// Convert hex string to *ecdsa.PrivateKey
	privateKey, err := crypto.ToECDSA([]byte(privateKeyHex))
	if err != nil {
		logger.Error(ctx, "error converting to ECDSA", "error", err.Error())
		return
	}

	// TODO change
	privateKeyTransactor := bind.NewKeyedTransactor(privateKey)

	// TODO estimate gas??
	// set gas
	privateKeyTransactor.GasPrice = new(big.Int).SetInt64(7732778450)

	quantity := new(big.Int).SetInt64(request.Quantity)
	buyerAddress := common.HexToAddress(request.Buyer)

	_, err = t.tokenContract.Mint(privateKeyTransactor, quantity, buyerAddress)
	if err != nil {
		fmt.Println("error while minting transaction", err.Error())
		return
	}

	logger.Info(ctx, "transaction successful")
	
	return nil
}
