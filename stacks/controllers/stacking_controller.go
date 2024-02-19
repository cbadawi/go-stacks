package controllers

import (
	"context"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/models"

	"github.com/apimatic/go-core-runtime/utilities"
)

// StackingController represents a controller struct.
type StackingController struct {
	baseController
}

// NewStackingController creates a new instance of StackingController.
// It takes a baseController as a parameter and returns a pointer to the StackingController.
func NewStackingController(baseController baseController) *StackingController {
	stackingController := StackingController{baseController: baseController}
	return &stackingController
}

// GetPoolDelegations takes context, poolPrincipal, afterBlock, unanchored, limit, offset as parameters and
// returns an models.ApiResponse with models.PoolDelegationsResponse data and
// an error if there was an issue with the request or response.
// Retrieves the list of stacking pool members for a given delegator principal.
func (s *StackingController) GetPoolDelegations(
	ctx context.Context,
	poolPrincipal string,
	afterBlock *int,
	unanchored *bool,
	limit *int,
	offset *int) (
	models.ApiResponse[models.PoolDelegationsResponse],
	error) {
	path := fmt.Sprintf("/extended/beta/stacking/%v/delegations", poolPrincipal)
	req := s.prepareRequest(
		ctx,
		"GET",
		path,
	)
	req.Authenticate(true)
	if afterBlock != nil {
		req.QueryParam("after_block", *afterBlock)
	}
	if unanchored != nil {
		req.QueryParam("unanchored", *unanchored)
	}
	if limit != nil {
		req.QueryParam("limit", *limit)
	}
	if offset != nil {
		req.QueryParam("offset", *offset)
	}

	var result models.PoolDelegationsResponse
	decoder, resp, err := s.LogCallAsJSON(req, path)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.PoolDelegationsResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}
