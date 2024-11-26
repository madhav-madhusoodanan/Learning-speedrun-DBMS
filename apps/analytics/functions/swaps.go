package functions

import (
	"context"
	"github.com/gin-gonic/gin"
	"rampx/backend/apps/analytics/structs"
	"net/http"
	"rampx/backend/db"
	"github.com/jackc/pgx/v5/pgtype"
	"math/big"
)

func ConvertStringToPgNumeric(f string) pgtype.Numeric {
	value := pgtype.Numeric{
		Valid: true,
		Int:   new(big.Int).SetInt64(0), // Using 18 decimals as common in most tokens
		Exp:   -18,
	}

	value.ScanScientific(f)
	return value
}

func AddSwap(client *db.Queries) func(c *gin.Context) {
	return func(c *gin.Context){
		ctx := context.Background()
		var req structs.AddSwapRequest
	
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		/* TODO: calculate the dollar value and the fee amount */

		request := db.CreateSwapOrderParams{
			SourceChainID: req.SourceChainID,
			SourceTokenAddress: req.SourceTokenAddress,
			SourceAmount: ConvertStringToPgNumeric(req.SourceAmount),
			DestinationChainID: req.DestChainID,
			DestinationTokenAddress: req.DestTokenAddress,
			DestinationAmount: ConvertStringToPgNumeric(req.DestAmount),
			DollarValue: ConvertStringToPgNumeric("0"),
			FeeDollarValue: ConvertStringToPgNumeric("0"),
			SourceAddress: req.SourceAddress,
			DestinationAddress: req.DestinationAddress,
			TransactionHash: req.TransactionHash,
		}

		err := client.CreateSwapOrder(ctx, request)
		if err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{})
	}
}