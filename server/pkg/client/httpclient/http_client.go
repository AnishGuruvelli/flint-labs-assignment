package httpclient

import (
	"flint-labs-assignment/internal/constants"
	"flint-labs-assignment/logconf"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func PostForObject(c *gin.Context, url string, headers map[string]string, body io.Reader) (*http.Response, error) {
	ctxlog := logconf.Logger(c).WithFields(logrus.Fields{"service": "HttpClient",
		"method": "PostForObject",
		"url":    url})

	ctxlog.Info("making post request for url")

	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		ctxlog.Error("error on creating post request:", err)
		return nil, err
	}

	// Set the request headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(constants.HCorrelationId, c.GetString(constants.CorrelationId))

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Send the request
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		ctxlog.Error("error on making post request:", err.Error())
		return nil, err
	}

	return resp, nil
}

func GetForUrl(c *gin.Context, url string, headers map[string]string) (*http.Response, error) {
	ctxlog := logconf.Logger(c).WithFields(logrus.Fields{"service": "HttpClient",
		"method": "GetForUrl",
		"url":    url})

	ctxlog.Info("making get request for url")

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		ctxlog.Error("error on creating get request:", err)
		return nil, err
	}

	// Set the request headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(constants.HCorrelationId, c.GetString(constants.CorrelationId))

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Send the request
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		ctxlog.Error("error on making get request:", err.Error())
		return nil, err
	}

	return resp, nil
}
