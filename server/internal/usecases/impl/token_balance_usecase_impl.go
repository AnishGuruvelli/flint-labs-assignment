package impl

import (
	"fmt"
	"net/http"
	"time"

	"flint-labs-assignment/internal/constants"
	"flint-labs-assignment/internal/customerror"
	"flint-labs-assignment/internal/dto"
	"flint-labs-assignment/pkg/client/infura"

	"flint-labs-assignment/internal/usecases"
	"flint-labs-assignment/logconf"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type TokenBalanceUsecaseImpl struct {
	infuraDataUsecase infura.InfuraDataService
}

var tokenBalanceUsecaseInstance *TokenBalanceUsecaseImpl

func NewTokenBalanceUsecase(infuraDataUsecase infura.InfuraDataService) usecases.TokenBalanceUsecase {
	if tokenBalanceUsecaseInstance == nil {
		tokenBalanceUsecaseInstance = &TokenBalanceUsecaseImpl{
			infuraDataUsecase: infuraDataUsecase,
		}
	}
	return tokenBalanceUsecaseInstance
}

func (uc *TokenBalanceUsecaseImpl) GetTokenBalance(c *gin.Context, walletAddress string) (*dto.BalanceData, *customerror.InternalErrorHandler) {
	ctxlog := logconf.Logger(c).WithFields(logrus.Fields{"service": "TokenBalanceUsecaseImpl", "method": "GetTokenBalance"})
	ctxlog.Info("Received request to fetch token balance")

	// Fetch current balance
	currentBalance, err := uc.infuraDataUsecase.GetWalletBalance(c, walletAddress)
	if err != nil {
		ctxlog.Errorf("Error fetching current balance: %v", err)
		return nil, &customerror.InternalErrorHandler{
			Error:     fmt.Errorf("error fetching current balance"),
			Code:      http.StatusInternalServerError,
			ErrorCode: constants.INTERNALSERVERERROR,
		}
	}

	// Fetch historical balance for the last 12 hours
	startTime := time.Now().Add(-12 * time.Hour)
	endTime := time.Now()

	// ctxlog.Info("getting historical balance")
	historicalBalance, err := uc.infuraDataUsecase.GetHistoricalBalance(c, walletAddress, startTime, endTime)
	if err != nil {
		ctxlog.Errorf("Error fetching historical balance: %v", err)
		// choose to continue without historical balance data
	}

	currentBalanceValue, _ := currentBalance.Float64()
	historicalBalanceValue, _ := historicalBalance.Float64()

	// Calculate the percentage change
	changePercentage := 0.0
	if historicalBalanceValue != 0 {
		changePercentage = ((currentBalanceValue - historicalBalanceValue) / historicalBalanceValue) * 100
	}

	// Create the response
	response := &dto.BalanceData{
		Balance:          currentBalanceValue,
		ChangePercentage: changePercentage,
		Timestamp:        time.Now(),
	}

	ctxlog.Info("Token balance fetched successfully")
	return response, nil
}
