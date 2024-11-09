package functions

import (
	"fmt"
	"log"
	"math/big"
	"net/http"
	"rampx/backend/apps/swaps/structs"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/gin-gonic/gin"
)

// convertTokenDetails converts RequestTokenDetail slice to TokenDetail slice
func convertTokenDetails(reqDetails []structs.RequestTokenDetail) ([]TokenDetail, error) {
	details := make([]TokenDetail, len(reqDetails))
	for i, detail := range reqDetails {
		amount := new(big.Int)
		amount, success := amount.SetString(detail.Amount, 10)
		if !success {
			return nil, fmt.Errorf("invalid amount format for token %s", detail.TokenAddress)
		}

		permitValue := new(big.Int)
		permitValue, success = permitValue.SetString(detail.PermitValue, 10)
		if !success {
			permitValue = common.Big0
		}

		permitDeadline := new(big.Int)
		permitDeadline, success = permitDeadline.SetString(detail.PermitDeadline, 10)
		if !success {
			permitDeadline = common.Big0
		}

		var permitR [32]byte
		permitRValue, err := hexutil.Decode(detail.PermitR)
		if err != nil {
			copy(permitR[:], permitRValue)
		}

		var permitS [32]byte
		permitSValue, err := hexutil.Decode(detail.PermitS)
		if err != nil {
			copy(permitS[:], permitSValue)
		}
		log.Println("PermitV: ", detail.PermitV)

		details[i] = TokenDetail{
			TokenAddress:   detail.TokenAddress,
			Amount:         amount,
			PermitValue:    permitValue,
			PermitDeadline: permitDeadline,
			PermitV:        detail.PermitV,
			PermitR:        [32]byte(permitR),
			PermitS:        [32]byte(permitS),
		}
	}
	return details, nil
}

// convertRoutes converts RequestRoute slice to Route slice
func convertRoutes(reqRoutes []structs.RequestRoute) ([]Route, error) {
	routes := make([]Route, len(reqRoutes))
	for i, route := range reqRoutes {
		// Convert ProcessorIndex
		processorIndex := new(big.Int)
		processorIndex, success := processorIndex.SetString(route.ProcessorIndex, 10)
		if !success {
			return nil, fmt.Errorf("invalid processor index format")
		}

		// Convert SellTokenDetails
		sellTokenDetails, err := convertTokenDetails(route.SellTokenDetails)
		if err != nil {
			return nil, err
		}

		// Decode Payload
		payload, err := hexutil.Decode(route.Payload)
		if err != nil {
			return nil, fmt.Errorf("invalid payload hex: %v", err)
		}

		// Decode Metadata
		metadata, err := hexutil.Decode(route.Metadata)
		if err != nil {
			return nil, fmt.Errorf("invalid metadata hex: %v", err)
		}

		routes[i] = Route{
			ProcessorIndex:   processorIndex,
			SellTokenDetails: sellTokenDetails,
			Payload:          payload,
			Metadata:         metadata,
		}
	}
	return routes, nil
}

// HandleCrossChainSwap handles the POST request for cross-chain swap
func HandleCrossChainSwap(c *gin.Context) {
	var req structs.SwapRequestRampX
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convert sell tokens
	sellTokens, err := convertTokenDetails(req.SellTokens)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sell tokens: " + err.Error()})
		return
	}

	// Convert buy tokens
	buyTokens, err := convertTokenDetails(req.BuyTokens)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid buy tokens: " + err.Error()})
		return
	}

	// Convert routes
	routes, err := convertRoutes(req.Routes)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid routes: " + err.Error()})
		return
	}

	// Decode user signature
	userSignature, err := hexutil.Decode(req.UserSignature)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid signature hex"})
		return
	}

	// Execute the swap
	err = ExecuteCrossChainSwap(
		req.UserAddress,
		sellTokens,
		buyTokens,
		routes,
		userSignature,
		req.ChainID,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Swap executed successfully"})
}
