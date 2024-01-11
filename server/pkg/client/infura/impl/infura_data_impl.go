package impl

import (
	"fmt"
	"math/big"
	"time"

	"flint-labs-assignment/logconf"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/ethereum/go-ethereum/rpc"
	"github.com/gin-gonic/gin"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type InfuraDataService struct{}

var infuraDataInstance *InfuraDataService

func NewInfuraDataService() *InfuraDataService {
	if infuraDataInstance == nil {
		infuraDataInstance = &InfuraDataService{}
	}
	return infuraDataInstance
}

func (s *InfuraDataService) GetWalletBalance(c *gin.Context, address string) (*big.Float, error) {
	ctxlog := logconf.Logger(c).WithFields(logrus.Fields{
		"service": "WalletClient",
		"method":  "GetWalletBalance",
		"address": address,
	})

	ctxlog.Info("fetching wallet balance for address: ", address)

	infuraHost := viper.GetString("infura.host")
	infuraAPIKey := viper.GetString("infura.auth.api.key")

	ethereumNodeURL := fmt.Sprintf("https://%s/v3/%s", infuraHost, infuraAPIKey)

	client, err := rpc.Dial(ethereumNodeURL)
	if err != nil {
		ctxlog.Error("error connecting to Ethereum node: ", err.Error())
		return nil, fmt.Errorf("error connecting to Ethereum node")
	}
	defer client.Close()

	var result hexutil.Big
	err = client.Call(&result, "eth_getBalance", address, "latest")
	if err != nil {
		ctxlog.Error("error calling eth_getBalance RPC: ", err.Error())
		return nil, fmt.Errorf("error calling eth_getBalance RPC")
	}

	weiBalance := new(big.Int).SetUint64(result.ToInt().Uint64())

	// Create a big.Int representing 10^18
	ethInWei := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)

	// Calculate Ether balance
	ethBalance := new(big.Float).Quo(new(big.Float).SetInt(weiBalance), new(big.Float).SetInt(ethInWei))
	return ethBalance, nil
}

func (s *InfuraDataService) GetHistoricalBalance(c *gin.Context, address string, startTime, endTime time.Time) (*big.Float, error) {
	ctxlog := logconf.Logger(c).WithFields(logrus.Fields{
		"service":   "WalletClient",
		"method":    "GetHistoricalBalance",
		"address":   address,
		"startTime": startTime,
		"endTime":   endTime,
	})

	ctxlog.Info("fetching historical balance for address: ", address)

	// etherscanAPIKey := viper.GetString("etherscan.api.key")
	// etherscanEndpoint := "https://api.etherscan.io/api"

	// // Format timestamps to Unix timestamps
	// startTimestamp := startTime.Unix()
	// endTimestamp := endTime.Unix()

	// // Etherscan API endpoint for historical balance
	// apiURL := fmt.Sprintf("%s?module=account&action=balancehistory&address=%s&startblock=0&endblock=999999999&apikey=%s&starttimestamp=%d&endtimestamp=%d",
	// 	etherscanEndpoint, address, etherscanAPIKey, startTimestamp, endTimestamp)

	// // Make the API request
	// response, err := http.Get(apiURL)
	// if err != nil {
	// 	ctxlog.Error("error calling Etherscan API: ", err.Error())
	// 	return nil, fmt.Errorf("error calling Etherscan API")
	// }
	// defer response.Body.Close()

	// // Parse the response body
	// var data map[string]interface{}
	// if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
	// 	ctxlog.Error("error decoding Etherscan API response: ", err.Error())
	// 	return nil, fmt.Errorf("error decoding Etherscan API response")
	// }

	// // Check if "result" key exists in the response
	// resultValue, ok := data["result"]
	// if !ok {
	// 	ctxlog.Error("missing 'result' key in Etherscan API response")
	// 	return nil, fmt.Errorf("missing 'result' key in Etherscan API response")
	// }

	// // Check the type of the value
	// switch resultValue := resultValue.(type) {
	// case float64:
	// 	// Value is already a float64, no need to convert
	// 	historicalBalance := resultValue
	// 	return big.NewFloat(historicalBalance), nil
	// case string:
	// 	// Attempt to convert string to float64
	// 	historicalBalance, err := strconv.ParseFloat(resultValue, 64)
	// 	if err != nil {
	// 		ctxlog.Error("error converting string to float64: ", err.Error())
	// 		return nil, fmt.Errorf("error converting string to float64")
	// 	}
	// 	return big.NewFloat(historicalBalance), nil
	// default:
	// 	ctxlog.Errorf("unexpected type for 'result': %T", resultValue)
	// 	return nil, fmt.Errorf("unexpected type for 'result'")
	// }

	return big.NewFloat(1), nil
}
