package functions 

import (
	"context"
	"github.com/gin-gonic/gin"
	"rampx/backend/apps/analytics/structs"
	"net/http"
	"rampx/backend/db"
)

/* 
	1. make a query to check if any username exists
	2. check if there is an existing address on the db
	3. create user
*/

func CheckExistingUsernames(client *db.Queries) func(c *gin.Context) {
	return func(c *gin.Context){
		ctx := context.Background()
		var req structs.CheckUsernameRequest
	
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		count, err := client.GetExistingWalletsByUsername(ctx, req.Username)
		if err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"count": count})
	}
}

func CheckExistingWallets(client *db.Queries) func(c *gin.Context) {
	return func(c *gin.Context){
		ctx := context.Background()
		var req structs.CheckUserAddressRequest
	
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		request := db.GetExistingWalletsByChainAndAddressParams{
			ChainID: req.ChainId,
			UserAddress: req.Address,
		}

		count, err := client.GetExistingWalletsByChainAndAddress(ctx, request)
		if err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"count": count})
	}
}

func AddWallet(client *db.Queries) func(c *gin.Context) {
	return func(c *gin.Context){
		ctx := context.Background()
		var req structs.CreateWalletRequest
	
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		request := db.AddUserWalletParams{
			UserName: req.Username,
			ChainID: req.ChainId,
			UserAddress: req.Address,
		}

		err := client.AddUserWallet(ctx, request)
		if err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{})
	}
}