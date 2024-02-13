package stacksblockchainapi

import (
	"os"

	"github.com/cbadawi/stacks-go-draft/logger"
)

// ConfigurationOptions represents a function type that can be used to apply options to the Configuration struct.
type ConfigurationOptions func(*Configuration)

// Configuration holds configuration settings.
type Configuration struct {
	baseUri           string
	environment       Environment
	httpConfiguration HttpConfiguration
	logger            *logger.Logger
}

// newConfiguration creates a new Configuration with the provided options.
func newConfiguration(options ...ConfigurationOptions) Configuration {
	config := Configuration{}

	for _, option := range options {
		option(&config)
	}
	return config
}

// WithEnvironment is an option that sets the Environment in the Configuration.
func WithEnvironment(environment Environment) ConfigurationOptions {
	return func(c *Configuration) {
		c.environment = environment
	}
}

// WithBaseUri is an option that sets the Base URI in the Configuration.
func WithBaseUri(uri string) ConfigurationOptions {
	return func(c *Configuration) {
		c.baseUri = uri
	}
}

// WithVerbose is an option that sets the log verbosity in the Configuration.
func WithVerbose(v bool) ConfigurationOptions {
	return func(c *Configuration) {
		c.logger.Verbose = v
	}
}

// WithHttpConfiguration is an option that sets the HttpConfiguration in the Configuration.
func WithHttpConfiguration(httpConfiguration HttpConfiguration) ConfigurationOptions {
	return func(c *Configuration) {
		c.httpConfiguration = httpConfiguration
	}
}

// WithLogger is an option that sets the Logger in the Configuration.
func WithLogger(loggerHandler *logger.Logger) ConfigurationOptions {
	return func(c *Configuration) {
		c.logger = loggerHandler
	}
}

// Environment returns the Environment from the Configuration.
func (c *Configuration) Environment() Environment {
	return c.environment
}

// HttpConfiguration returns the HttpConfiguration from the Configuration.
func (c *Configuration) HttpConfiguration() HttpConfiguration {
	return c.httpConfiguration
}

// CreateConfigurationFromEnvironment creates a new Configuration with default settings.
// It also configures various Configuration options.
func CreateConfigurationFromEnvironment(options ...ConfigurationOptions) Configuration {
	config := DefaultConfiguration()

	environment := os.Getenv("STACKSBLOCKCHAINAPI_ENVIRONMENT")
	if environment != "" {
		config.environment = Environment(environment)
	}
	for _, option := range options {
		option(&config)
	}
	return config
}

// Server represents available servers.
type Server string

const (
	ENUMDEFAULT Server = "default"
)

// Environment represents available environments.
type Environment string

const (
	MAINNET_URI string = "https://api.mainnet.hiro.so/"
	TESTNET_URI string = "https://api.testnet.hiro.so/"
)

const (
	PRODUCTION   Environment = "production"
	ENVIRONMENT2 Environment = "environment2"
	ENVIRONMENT3 Environment = "environment3"
)

// CreateRetryConfiguration creates a new RetryConfiguration with the provided options.
func CreateRetryConfiguration(options ...RetryConfigurationOptions) RetryConfiguration {
	config := DefaultRetryConfiguration()

	for _, option := range options {
		option(&config)
	}
	return config
}

// CreateHttpConfiguration creates a new HttpConfiguration with the provided options.
func CreateHttpConfiguration(options ...HttpConfigurationOptions) HttpConfiguration {
	config := DefaultHttpConfiguration()

	for _, option := range options {
		option(&config)
	}
	return config
}

// CreateConfiguration creates a new Configuration with the provided options.
func CreateConfiguration(options ...ConfigurationOptions) Configuration {
	config := DefaultConfiguration()

	for _, option := range options {
		option(&config)
	}
	return config
}
