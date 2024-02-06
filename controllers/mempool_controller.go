package controllers

import (
	"context"

	"github.com/cbadawi/stacks-go-draft/models"

	"github.com/apimatic/go-core-runtime/utilities"
)

// MempoolController represents a controller struct.
type MempoolController struct {
	baseController
}

// NewMempoolController creates a new instance of MempoolController.
// It takes a baseController as a parameter and returns a pointer to the MempoolController.
func NewMempoolController(baseController baseController) *MempoolController {
	mempoolController := MempoolController{baseController: baseController}
	return &mempoolController
}

// GetMempoolFeePriorities takes context as parameters and
// returns an models.ApiResponse with models.MempoolFeePriorities data and
// an error if there was an issue with the request or response.
// Returns estimated fee priorities (in micro-STX) for all transactions that are currently in the mempool. Also returns priorities separated by transaction type.
func (m *MempoolController) GetMempoolFeePriorities(ctx context.Context) (
	models.ApiResponse[models.MempoolFeePriorities],
	error) {
	path := "/extended/v2/mempool/fees"
	req := m.prepareRequest(ctx, "GET", path)
	req.Authenticate(true)
	var result models.MempoolFeePriorities
	decoder, resp, err := m.LogCallAsJSON(req, path)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.MempoolFeePriorities](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}
