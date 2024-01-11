package dto

import (
	"time"
)

// BalanceData struct represents the data sent to the frontend
type BalanceData struct {
	Balance          float64    `json:"balance"`
	ChangePercentage float64   `json:"changePercentage"`
	Timestamp        time.Time `json:"timestamp"`
}
