package usecases

import (
	"flint-labs-assignment/internal/customerror"
	"flint-labs-assignment/internal/dto"

	"github.com/gin-gonic/gin"
)

type TokenBalanceUsecase interface {
	GetTokenBalance(c *gin.Context, walletAddress string) (*dto.BalanceData, *customerror.InternalErrorHandler)
}
