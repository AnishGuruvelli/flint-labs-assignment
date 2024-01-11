package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

// BalanceData struct represents the data sent to the frontend
type BalanceData struct {
	Balance          float64   `json:"balance"`
	ChangePercentage float64   `json:"changePercentage"`
	Timestamp        time.Time `json:"timestamp"`
}

var ExternalAddresses = map[string]string{
	"Mantle": "0xDCBc586cAb42a1D193CaCD165a81E5fbd9B428d7",
	"Linea":  "0xDCBc586cAb42a1D193CaCD165a81E5fbd9B428d7",
	"Kroma":  "0x7afb9de72A9A321fA535Bb36b7bF0c987b42b859",
}

var previousBalances = make(map[string]BalanceData)

func GetBalanceHandler(c *gin.Context) {
	chain := c.Param("chain")
	address := ExternalAddresses[chain]

	// Replace this with your logic to fetch the balance and calculate percentage change
	balance, err := getEthereumBalance(chain, address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching balance"})
		return
	}

	changePercentage := calculatePercentageChange(address, balance)
	timestamp := time.Now()

	// Update the previous balance
	previousBalances[address] = BalanceData{
		Balance:          balance,
		ChangePercentage: changePercentage,
		Timestamp:        timestamp,
	}

	// Create a BalanceData struct
	balanceData := BalanceData{
		Balance:          balance,
		ChangePercentage: changePercentage,
		Timestamp:        timestamp,
	}

	// Set the response headers
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, balanceData)
}

func getEthereumBalance(chain, address string) (float64, error) {
	var endpoint string

	switch chain {
	case "Mantle":
		endpoint = "https://mainnet.infura.io/v3/3e581c1579624c01861c8400629bef66"
	case "Linea":
		endpoint = "https://linea.infura.io/v3/your-infura-api-key" // Replace with your Infura API key
	case "Kroma":
		endpoint = "https://kroma.infura.io/v3/your-infura-api-key" // Replace with your Infura API key
	default:
		return 0, fmt.Errorf("invalid chain: %s", chain)
	}

	client, err := ethclient.Dial(endpoint)
	if err != nil {
		return 0, err
	}

	// Get the balance
	ethBalance, err := getBalance(client, common.HexToAddress(address))
	if err != nil {
		return 0, err
	}

	balanceFloat, _ := new(big.Float).SetInt(ethBalance).Float64()

	return balanceFloat / 1e18, nil // Convert Wei to Ether
}

func getBalance(client *ethclient.Client, address common.Address) (*big.Int, error) {
	// Call the eth_getBalance method
	ethBalance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		return nil, err
	}

	return ethBalance, nil
}

func calculatePercentageChange(address string, currentBalance float64) float64 {
	previousBalance := previousBalances[address].Balance
	if previousBalance != 0 {
		return ((currentBalance - previousBalance) / previousBalance) * 100
	}
	return 0
}

func main() {
	r := gin.Default()

	// Define routes
	r.GET("/getBalance/:chain", GetBalanceHandler)

	// Serve static files (your frontend assets)
	r.StaticFS("/", http.Dir("./client"))

	// Start the server
	port := "8080"
	err := r.Run(":" + port)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
