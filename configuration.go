package main

import (
	"os"
)

// ConfigurationOptions represents a function type that can be used to apply options to the Configuration struct.
type ConfigurationOptions func(*Configuration)

// Configuration holds configuration settings.
type Configuration struct {
	environment       Environment
	httpConfiguration HttpConfiguration
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

// WithHttpConfiguration is an option that sets the HttpConfiguration in the Configuration.
func WithHttpConfiguration(httpConfiguration HttpConfiguration) ConfigurationOptions {
	return func(c *Configuration) {
		c.httpConfiguration = httpConfiguration
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
