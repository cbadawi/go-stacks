package stacksblockchainapi

import (
	"net/http"
	"os"

	"github.com/cbadawi/stacks-go-draft/logger"
)

// DefaultRetryConfiguration returns the default RetryConfiguration for HTTP requests.
// It also configures various retry options.
func DefaultRetryConfiguration() RetryConfiguration {
	return NewRetryConfiguration(
		WithMaxRetryAttempts(0),
		WithRetryOnTimeout(true),
		WithRetryInterval(1),
		WithMaximumRetryWaitTime(0),
		WithBackoffFactor(2),
		WithHttpStatusCodesToRetry([]int64{408, 413, 429, 500, 502, 503, 504, 521, 522, 524}),
		WithHttpMethodsToRetry([]string{"GET", "PUT"}),
	)
}

// DefaultHttpConfiguration returns the default HttpConfiguration for HTTP requests.
// It also configures various HttpConfiguration options.
func DefaultHttpConfiguration() HttpConfiguration {
	return NewHttpConfiguration(
		WithTimeout(0),
		WithTransport(http.DefaultTransport),
		WithRetryConfiguration(DefaultRetryConfiguration()),
	)
}

// DefaultConfiguration returns the default Configuration.
func DefaultConfiguration() Configuration {
	// Initialize a new jsonlog.Logger which writes any messages *at or above* the INFO
	// severity level to the standard out stream.
	loggerHandler := logger.NewLogger(os.Stdout, logger.LevelInfo)

	return newConfiguration(
		WithEnvironment(PRODUCTION),
		WithHttpConfiguration(DefaultHttpConfiguration()),
		WithLogger(loggerHandler),
	)
}
