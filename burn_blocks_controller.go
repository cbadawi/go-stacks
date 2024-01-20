package main

import (
	"context"
	"fmt"

	"github.com/cbadawi/stacks-go-draft/models"

	"github.com/apimatic/go-core-runtime/utilities"
)

// BurnBlocksController represents a controller struct.
type BurnBlocksController struct {
	baseController
}

// NewBurnBlocksController creates a new instance of BurnBlocksController.
// It takes a baseController as a parameter and returns a pointer to the BurnBlocksController.
func NewBurnBlocksController(baseController baseController) *BurnBlocksController {
	burnBlocksController := BurnBlocksController{baseController: baseController}
	return &burnBlocksController
}

// GetBurnBlocks takes context, limit, offset as parameters and
// returns an models.ApiResponse with models.BurnBlockListResponse data and
// an error if there was an issue with the request or response.
// Retrieves a list of recent burn blocks
func (b *BurnBlocksController) GetBurnBlocks(
	ctx context.Context,
	limit *int,
	offset *int) (
	models.ApiResponse[models.BurnBlockListResponse],
	error) {
	req := b.prepareRequest(ctx, "GET", "/extended/v2/burn-blocks")
	req.Authenticate(true)
	if limit != nil {
		req.QueryParam("limit", *limit)
	}
	if offset != nil {
		req.QueryParam("offset", *offset)
	}
	var result models.BurnBlockListResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.BurnBlockListResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetBurnBlock takes context, heightOrHash as parameters and
// returns an models.ApiResponse with models.BurnBlock data and
// an error if there was an issue with the request or response.
// Retrieves a single burn block
func (b *BurnBlocksController) GetBurnBlock(
	ctx context.Context,
	heightOrHash interface{}) (
	models.ApiResponse[models.BurnBlock],
	error) {
	req := b.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/extended/v2/burn-blocks/%v", heightOrHash),
	)
	req.Authenticate(true)

	var result models.BurnBlock
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.BurnBlock](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}
