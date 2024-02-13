package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/apimatic/go-core-runtime/https"
	"github.com/cbadawi/go-stacks/logger"
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

func (bc baseController) logCall(req https.CallBuilder, path string, reqType string) (interface{}, *http.Response, error) {
	bc.logger.TryLog("API Call", map[string]string{"path": path}, logger.LevelInfo)
	var result any
	var resp *http.Response
	var err error
	switch reqType {
	case "json":
		result, resp, err = req.CallAsJson()
	case "text":
		result, resp, err = req.CallAsText()
	default:
		bc.logger.PrintError(errors.New("Unhandled call type"), nil)
	}
	if err != nil {
		bc.logger.TryLog("API Call Error", map[string]string{"error": fmt.Sprint(err)}, logger.LevelError)
		return result, resp, err
	}

	// if resp.StatusCode == 400 {
	// 	err := errors.New(resp.Status)
	// 	return result, resp, err
	// }

	bc.logger.TryLog("API Call Response", map[string]string{"code": fmt.Sprint(resp.StatusCode), "status": resp.Status}, logger.LevelInfo)
	return result, resp, err
}

func (bc baseController) LogCallAsJSON(req https.CallBuilder, path string) (*json.Decoder, *http.Response, error) {
	result, resp, err := bc.logCall(req, path, "json")
	if decoder, ok := result.(*json.Decoder); ok {
		return decoder, resp, err
	} else {
		err = errors.New(fmt.Sprintf("Expected json.Decoder type: %T", result))
		bc.logger.PrintError(err, map[string]string{"result": fmt.Sprint(result)})
		return nil, resp, err
	}
}

func (bc baseController) LogCallAsText(req https.CallBuilder, path string) (string, *http.Response, error) {
	result, resp, err := bc.logCall(req, path, "text")
	return result.(string), resp, err
}
