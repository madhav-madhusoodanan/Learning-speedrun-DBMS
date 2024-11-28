package analytics

import (
	"rampx/backend/internal/apps/analytics/internal/functions"
	"rampx/backend/internal/db"
	"github.com/gin-gonic/gin"
)

func Setup(router *gin.RouterGroup, client *db.Queries){
	router.POST("/swap/create", functions.AddSwap(client))
	router.POST("/user/name/check", functions.CheckExistingUsernames(client))
	router.POST("/user/address/check", functions.CheckExistingWallets(client))
	router.POST("/user/create", functions.AddWallet(client))
}