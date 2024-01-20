package stacksblockchainapi

import (
	"context"
	"fmt"
	"stacksblockchainapi/models"

	"github.com/apimatic/go-core-runtime/utilities"
)

// StackingRewardsController represents a controller struct.
type StackingRewardsController struct {
	baseController
}

// NewStackingRewardsController creates a new instance of StackingRewardsController.
// It takes a baseController as a parameter and returns a pointer to the StackingRewardsController.
func NewStackingRewardsController(baseController baseController) *StackingRewardsController {
	stackingRewardsController := StackingRewardsController{baseController: baseController}
	return &stackingRewardsController
}

// GetBurnchainRewardSlotHolders takes context, limit, offset as parameters and
// returns an models.ApiResponse with models.BurnchainRewardSlotHolderListResponse data and
// an error if there was an issue with the request or response.
// Retrieves a list of the Bitcoin addresses that would validly receive Proof-of-Transfer commitments.
func (s *StackingRewardsController) GetBurnchainRewardSlotHolders(
	ctx context.Context,
	limit *int,
	offset *int) (
	models.ApiResponse[models.BurnchainRewardSlotHolderListResponse],
	error) {
	req := s.prepareRequest(ctx, "GET", "/extended/v1/burnchain/reward_slot_holders")
	req.Authenticate(true)
	if limit != nil {
		req.QueryParam("limit", *limit)
	}
	if offset != nil {
		req.QueryParam("offset", *offset)
	}
	var result models.BurnchainRewardSlotHolderListResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.BurnchainRewardSlotHolderListResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetBurnchainRewardSlotHoldersByAddress takes context, address, limit, offset as parameters and
// returns an models.ApiResponse with models.BurnchainRewardSlotHolderListResponse data and
// an error if there was an issue with the request or response.
// Retrieves a list of the Bitcoin addresses that would validly receive Proof-of-Transfer commitments for a given reward slot holder recipient address.
func (s *StackingRewardsController) GetBurnchainRewardSlotHoldersByAddress(
	ctx context.Context,
	address string,
	limit *int,
	offset *int) (
	models.ApiResponse[models.BurnchainRewardSlotHolderListResponse],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/extended/v1/burnchain/reward_slot_holders/%v", address),
	)
	req.Authenticate(true)
	if limit != nil {
		req.QueryParam("limit", *limit)
	}
	if offset != nil {
		req.QueryParam("offset", *offset)
	}

	var result models.BurnchainRewardSlotHolderListResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.BurnchainRewardSlotHolderListResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetBurnchainRewardList takes context, limit, offset as parameters and
// returns an models.ApiResponse with models.BurnchainRewardListResponse data and
// an error if there was an issue with the request or response.
// Retrieves a list of recent burnchain (e.g. Bitcoin) reward recipients with the associated amounts and block info
func (s *StackingRewardsController) GetBurnchainRewardList(
	ctx context.Context,
	limit *int,
	offset *int) (
	models.ApiResponse[models.BurnchainRewardListResponse],
	error) {
	req := s.prepareRequest(ctx, "GET", "/extended/v1/burnchain/rewards")
	req.Authenticate(true)
	if limit != nil {
		req.QueryParam("limit", *limit)
	}
	if offset != nil {
		req.QueryParam("offset", *offset)
	}
	var result models.BurnchainRewardListResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.BurnchainRewardListResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetBurnchainRewardListByAddress takes context, address, limit, offset as parameters and
// returns an models.ApiResponse with models.BurnchainRewardListResponse data and
// an error if there was an issue with the request or response.
// Retrieves a list of recent burnchain (e.g. Bitcoin) rewards for the given recipient with the associated amounts and block info
func (s *StackingRewardsController) GetBurnchainRewardListByAddress(
	ctx context.Context,
	address string,
	limit *int,
	offset *int) (
	models.ApiResponse[models.BurnchainRewardListResponse],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/extended/v1/burnchain/rewards/%v", address),
	)
	req.Authenticate(true)
	if limit != nil {
		req.QueryParam("limit", *limit)
	}
	if offset != nil {
		req.QueryParam("offset", *offset)
	}

	var result models.BurnchainRewardListResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.BurnchainRewardListResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetBurnchainRewardsTotalByAddress takes context, address as parameters and
// returns an models.ApiResponse with models.BurnchainRewardsTotal data and
// an error if there was an issue with the request or response.
// Retrieves the total burnchain (e.g. Bitcoin) rewards for a given recipient `address`
func (s *StackingRewardsController) GetBurnchainRewardsTotalByAddress(
	ctx context.Context,
	address string) (
	models.ApiResponse[models.BurnchainRewardsTotal],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/extended/v1/burnchain/rewards/%v/total", address),
	)
	req.Authenticate(true)

	var result models.BurnchainRewardsTotal
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.BurnchainRewardsTotal](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}
