package functions

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"
	"os"
	"rampx/backend/apps/swaps/contracts"
	"rampx/backend/apps/swaps/utils"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

/*
	TODO:
	1.[x] Extract user address, tokenDetails, routes and signatures from call
	2.[x] Call contract
	3.[x] Setup random choosing of rpc from list to distribute load
	4.[x] Set chain id too
	5.[x] Private key, put as ENV variable
*/

type TokenDetail struct {
	TokenAddress string
	Amount       *big.Int
}

type Route struct {
	ProcessorIndex   *big.Int
	SellTokenDetails []TokenDetail
	Payload          []byte
	Metadata         []byte
}

func ExecuteCrossChainSwap(
	userAddress string,
	sellTokens []TokenDetail,
	buyTokens []TokenDetail,
	routes []Route,
	userSignature []byte,
	chainId int64,
) error {

	privateKeyString := os.Getenv("EXECUTOR_PRIVATE_KEY")
	privateKey, err := crypto.HexToECDSA(privateKeyString)
	if err != nil {
		log.Println("[ERROR]: getting private key: ", err)
		return err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Println("[ERROR]: error casting public key to ECDSA")
	}

	signerAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	rpc := utils.GetRPC(chainId)
	client, err := ethclient.Dial(rpc)
	if err != nil {
		log.Println("[ERROR]: could not dial RPC: ", err)
		return err
	}

	nonce, err := client.PendingNonceAt(context.Background(), signerAddress)
	if err != nil {
		log.Println("[ERROR]: not able to get nonce of executor: ", err)
		return err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Println("[ERROR]: not able to find gas price: ", err)
		return err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(chainId))
	if err != nil {
		log.Println("[ERROR]: not able to generate transactor: ", err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(1000000) // in units
	auth.GasPrice = gasPrice

	// replace with actual contract deployment
	aggregatorAddress := utils.GetAggregator(chainId)
	address := common.HexToAddress(aggregatorAddress)
	instance, err := contracts.NewContracts(address, client)

	if err != nil {
		log.Println("[ERROR]: not able to connext to contract: ", err)
		return err
	}

	userAddressContract := common.HexToAddress(userAddress)

	buyTokensContract := []contracts.ExecutionVerifierTokenDetail{}

	for _, buyToken := range buyTokens {
		buyTokenContract := contracts.ExecutionVerifierTokenDetail{
			Token:  common.HexToAddress(buyToken.TokenAddress),
			Amount: buyToken.Amount,
		}

		buyTokensContract = append(buyTokensContract, buyTokenContract)
	}

	sellTokensContract := []contracts.ExecutionVerifierTokenDetail{}

	for _, sellToken := range sellTokens {
		sellTokenContract := contracts.ExecutionVerifierTokenDetail{
			Token:  common.HexToAddress(sellToken.TokenAddress),
			Amount: sellToken.Amount,
		}

		sellTokensContract = append(sellTokensContract, sellTokenContract)
	}

	routesContract := []contracts.ExecutionVerifierRoute{}

	for _, route := range routes {
		sellTokensContract = []contracts.ExecutionVerifierTokenDetail{}
		for _, sellToken := range route.SellTokenDetails {
			sellTokenContract := contracts.ExecutionVerifierTokenDetail{
				Token:  common.HexToAddress(sellToken.TokenAddress),
				Amount: sellToken.Amount,
			}

			sellTokensContract = append(sellTokensContract, sellTokenContract)
		}

		routeContract := contracts.ExecutionVerifierRoute{
			ProcessorIndex:   route.ProcessorIndex,
			SellTokenDetails: sellTokensContract,
			Payload:          route.Payload,
			Metadata:         route.Metadata,
		}

		routesContract = append(routesContract, routeContract)
	}

	_, err = instance.Execute0(auth, userAddressContract, sellTokensContract, buyTokensContract, routesContract, userSignature)

	if err != nil {
		log.Println("[ERROR]: Failed to execute transaction for user: ", err)
		return err
	}

	// log.Println("[TRANSACTION FULFILLED]: ", tx.)
	return nil
}
