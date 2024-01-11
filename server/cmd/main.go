package main

import (
	"flint-labs-assignment/configs"

	"flint-labs-assignment/initconfig"
	"flint-labs-assignment/pkg/httprest"
)

func init() {
	_, err := configs.NewConfig()
	if err != nil {
		panic("failed to load config: " + err.Error())
	}

	initconfig.InitialiseAppConfigs()
	httprest.InitEndpoints()
}

func main() {

}
