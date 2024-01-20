package stacksblockchainapi

import (
	"context"
	"fmt"

	"github.com/cbadawi/stacks-go-draft/errors"

	"github.com/cbadawi/stacks-go-draft/models"

	"github.com/apimatic/go-core-runtime/utilities"
)

// MicroblocksController represents a controller struct.
type MicroblocksController struct {
	baseController
}

// NewMicroblocksController creates a new instance of MicroblocksController.
// It takes a baseController as a parameter and returns a pointer to the MicroblocksController.
func NewMicroblocksController(baseController baseController) *MicroblocksController {
	microblocksController := MicroblocksController{baseController: baseController}
	return &microblocksController
}

// GetMicroblockList takes context, limit, offset as parameters and
// returns an models.ApiResponse with models.MicroblockListResponse data and
// an error if there was an issue with the request or response.
// Retrieves a list of microblocks.
// If you need to actively monitor new microblocks, we highly recommend subscribing to [WebSockets or Socket.io](https://github.com/hirosystems/stacks-blockchain-api/tree/master/client) for real-time updates.
func (m *MicroblocksController) GetMicroblockList(
	ctx context.Context,
	limit *int,
	offset *int) (
	models.ApiResponse[models.MicroblockListResponse],
	error) {
	req := m.prepareRequest(ctx, "GET", "/extended/v1/microblock")
	req.Authenticate(true)
	if limit != nil {
		req.QueryParam("limit", *limit)
	}
	if offset != nil {
		req.QueryParam("offset", *offset)
	}
	var result models.MicroblockListResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.MicroblockListResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetMicroblockByHash takes context, hash as parameters and
// returns an models.ApiResponse with models.Microblock data and
// an error if there was an issue with the request or response.
// Retrieves a specific microblock by `hash`
func (m *MicroblocksController) GetMicroblockByHash(
	ctx context.Context,
	hash string) (
	models.ApiResponse[models.Microblock],
	error) {
	req := m.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/extended/v1/microblock/%v", hash),
	)
	req.Authenticate(true)

	var result models.Microblock
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Microblock](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	if resp.StatusCode == 404 {
		err = errors.NewApiError(404, "Cannot find microblock with given hash")
	}
	return models.NewApiResponse(result, resp), err
}

// GetUnanchoredTxs takes context as parameters and
// returns an models.ApiResponse with models.UnanchoredTransactionListResponse data and
// an error if there was an issue with the request or response.
// Retrieves transactions that have been streamed in microblocks but not yet accepted or rejected in an anchor block
func (m *MicroblocksController) GetUnanchoredTxs(ctx context.Context) (
	models.ApiResponse[models.UnanchoredTransactionListResponse],
	error) {
	req := m.prepareRequest(ctx, "GET", "/extended/v1/microblock/unanchored/txs")
	req.Authenticate(true)
	var result models.UnanchoredTransactionListResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.UnanchoredTransactionListResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}
