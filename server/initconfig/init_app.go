package initconfig

import (
	"flint-labs-assignment/configs"

	"flint-labs-assignment/internal/usecases"
	usecaseimpl "flint-labs-assignment/internal/usecases/impl"
	infura "flint-labs-assignment/pkg/client/infura"
	infuraimpl "flint-labs-assignment/pkg/client/infura/impl"

	"flint-labs-assignment/logconf"

	"github.com/sirupsen/logrus"
)

type UseCaseContainer struct {
	InfuraDataUsecase   infura.InfuraDataService
	TokenBalanceUsecase usecases.TokenBalanceUsecase
}

var Usecases *UseCaseContainer

func InitialiseAppConfigs() {
	logrus.Info("initialising the app configs")
	logconf.InitializeLogger()
	_, err := configs.NewConfig()

	if err != nil {
		logrus.Error("error on initialisation of configs")
		panic(err)
	}

	infuraDataUsecase := infuraimpl.NewInfuraDataService()

	// Initialize use cases
	tokenBalanceUsecase := usecaseimpl.NewTokenBalanceUsecase(infuraDataUsecase)

	// Create the UsecaseContainer
	Usecases = &UseCaseContainer{
		InfuraDataUsecase:   infuraDataUsecase,
		TokenBalanceUsecase: tokenBalanceUsecase,
	}

	logrus.Info("initialised the app configs successfully")
}
