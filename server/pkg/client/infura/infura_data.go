package infura

import (
	"math/big"
	"time"

	"github.com/gin-gonic/gin"
)

type InfuraDataService interface {
	GetWalletBalance(c *gin.Context, address string) (*big.Float, error)
	GetHistoricalBalance(c *gin.Context, address string, startTime, endTime time.Time) (*big.Float, error)
	// "Sorry, it looks like you are trying to access an API Pro endpoint. Contact us to upgrade to API Pro.\": invalid syntax","service":"WalletCli
}
