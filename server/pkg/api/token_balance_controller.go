package api

import (
	"net/http"

	"flint-labs-assignment/initconfig"
	"flint-labs-assignment/internal/constants"
	"flint-labs-assignment/internal/customerror"
	"flint-labs-assignment/logconf"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Health(c *gin.Context) {
	c.JSON(200, gin.H{"status": "UP"})
}

func GetTokenBalance(c *gin.Context) {
	ctxlog := logconf.Logger(c).WithFields(logrus.Fields{"service": "TokenBalanceAPI", "method": "GetTokenBalance"})
	ctxlog.Info("Received request to get token balance")

	walletAddress := c.Param("walletAddress")

	balance, customErr := initconfig.Usecases.TokenBalanceUsecase.GetTokenBalance(c, walletAddress)

	if customErr != nil {
		ctxlog.Errorf("Error fetching token balance: %v", customErr)
		c.JSON(customErr.Code, InternalError(c, customErr.Error, "error while fetching token balance"))
		return
	}

	ctxlog.Info("Token balance fetched successfully")
	c.JSON(http.StatusOK, balance)
}

func InternalError(c *gin.Context, err error, msg string) *customerror.CustomError {
	var errsData []customerror.ErrorData
	errData := customerror.ErrorData{Code: err.Error(), Description: err.Error()}
	errsData = append(errsData, errData)
	custErr := customerror.NewCustomError(err.Error(), constants.INTERNALSERVERERROR, errsData)
	return custErr
}
