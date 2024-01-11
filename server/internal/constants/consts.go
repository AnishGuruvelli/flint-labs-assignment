package constants

import (
	"github.com/google/uuid"
)

const (
	INTERNALSERVERERROR = "APIDATA001"
	BADREQUEST          = "APIDATA002"

	APIVALIDATIONFAILED = "APIKEY001"
	APIKEYEMPTY         = "APIKEY002"
)

var CorrelationId = "correlation-id"
var HCorrelationId = "X-Correlation-Id"

var HApiKey = "X-API-Key"

var NullUUIDObj, _ = uuid.Parse("00000000-0000-0000-0000-000000000000")
