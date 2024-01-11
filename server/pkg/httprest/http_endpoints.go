package httprest

import (
	"os"

	"flint-labs-assignment/pkg/api"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func InitEndpoints() {
	logrus.Info("initializing endpoints")
	r := gin.Default()

	r.Use(CORSMiddleware())

	r.GET("/v1/healthz", api.Health)

	tokenBalanceEndpointsGroup := r.Group("/v1/token-balance")
	{
		tokenBalanceEndpointsGroup.GET("/:walletAddress", api.GetTokenBalance)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081" // Default port if not defined
	}

	err := r.Run(":" + port)
	if err != nil {
		logrus.Error("Error on initializing endpoints, Err: ", err.Error())
		return
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
