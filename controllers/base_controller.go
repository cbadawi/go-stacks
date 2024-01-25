package controllers

import (
	"net/http"

	"github.com/apimatic/go-core-runtime/https"
	"github.com/cbadawi/stacks-go-draft/logger"
)

// callBuilderFactory is an interface that defines a method to get a CallBuilderFactory.
// It allows objects to get a reference to a CallBuilderFactory for creating API call.
type callBuilderFactory interface {
	GetCallBuilder() https.CallBuilderFactory
}

type CallBuilderLogger struct {
	Cb     callBuilderFactory
	Logger *logger.Logger
}

// baseController represents a controller used as a base for other controllers.
// It encapsulates common functionality required by controllers for making API call.
type baseController struct {
	callBuilder    callBuilderFactory
	prepareRequest https.CallBuilderFactory
	logger         *logger.Logger
}

// NewBaseController creates a new instance of baseController.
// It takes a callBuilderFactory as a parameter and returns a pointer to the baseController.
func NewBaseController(cbl CallBuilderLogger) *baseController {
	baseController := baseController{callBuilder: cbl.Cb}
	baseController.prepareRequest = baseController.callBuilder.GetCallBuilder()
	baseController.logger = cbl.Logger
	return &baseController
}

// validateResponse is a function used to validate the HTTP response.
// It takes an http.Response object as a parameter and returns an error, if any.
func validateResponse(response http.Response) error {
	return nil
}
