package analytics

import (
	"rampx/backend/apps/analytics/functions"
	"github.com/gin-gonic/gin"
	"rampx/backend/db"
)

func Setup(router *gin.RouterGroup, client *db.Queries){
	router.POST("/swap/create", functions.AddSwap(client))
	router.POST("/user/name/check", functions.CheckExistingUsernames(client))
	router.POST("/user/address/check", functions.CheckExistingWallets(client))
	router.POST("/user/create", functions.AddWallet(client))
}