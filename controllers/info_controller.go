package controllers

import (
	"context"
	"fmt"

	"github.com/cbadawi/stacks-go-draft/models"

	"github.com/apimatic/go-core-runtime/utilities"
)

// InfoController represents a controller struct.
type InfoController struct {
	baseController
}

// NewInfoController creates a new instance of InfoController.
// It takes a baseController as a parameter and returns a pointer to the InfoController.
func NewInfoController(baseController baseController) *InfoController {
	infoController := InfoController{baseController: baseController}
	return &infoController
}

// GetCoreApiInfo takes context as parameters and
// returns an models.ApiResponse with models.CoreNodeInfoResponse data and
// an error if there was an issue with the request or response.
// Retrieves information about the Core API including the server version
func (i *InfoController) GetCoreApiInfo(ctx context.Context) (
	models.ApiResponse[models.CoreNodeInfoResponse],
	error) {
	req := i.prepareRequest(ctx, "GET", "/v2/info")
	req.Authenticate(true)
	var result models.CoreNodeInfoResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.CoreNodeInfoResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetStatus takes context as parameters and
// returns an models.ApiResponse with models.ServerStatusResponse data and
// an error if there was an issue with the request or response.
// Retrieves the running status of the Stacks Blockchain API, including the server version and current chain tip information.
func (i *InfoController) GetStatus(ctx context.Context) (
	models.ApiResponse[models.ServerStatusResponse],
	error) {
	req := i.prepareRequest(ctx, "GET", "/extended/v1/status")
	req.Authenticate(true)
	var result models.ServerStatusResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ServerStatusResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetNetworkBlockTimes takes context as parameters and
// returns an models.ApiResponse with models.NetworkBlockTimesResponse data and
// an error if there was an issue with the request or response.
// Retrieves the target block times for mainnet and testnet. The block time is hardcoded and will change throughout the implementation phases of the testnet.
func (i *InfoController) GetNetworkBlockTimes(ctx context.Context) (
	models.ApiResponse[models.NetworkBlockTimesResponse],
	error) {
	req := i.prepareRequest(ctx, "GET", "/extended/v1/info/network_block_times")
	req.Authenticate(true)
	var result models.NetworkBlockTimesResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.NetworkBlockTimesResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetNetworkBlockTimeByNetwork takes context, network as parameters and
// returns an models.ApiResponse with models.TargetBlockTime data and
// an error if there was an issue with the request or response.
// Retrieves the target block time for a given network. The network can be mainnet or testnet. The block time is hardcoded and will change throughout the implementation phases of the testnet.
func (i *InfoController) GetNetworkBlockTimeByNetwork(
	ctx context.Context,
	network models.NetworkEnum) (
	models.ApiResponse[models.TargetBlockTime],
	error) {
	req := i.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/extended/v1/info/network_block_time/%v", network),
	)
	req.Authenticate(true)

	var result models.TargetBlockTime
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.TargetBlockTime](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetStxSupply takes context, height as parameters and
// returns an models.ApiResponse with models.GetStxSupplyResponse data and
// an error if there was an issue with the request or response.
// Retrieves the total and unlocked STX supply. More information on Stacking can be found [here] (https://docs.stacks.co/understand-stacks/stacking).
// **Note:** This uses the estimated future total supply for the year 2050.
func (i *InfoController) GetStxSupply(
	ctx context.Context,
	height *float64) (
	models.ApiResponse[models.GetStxSupplyResponse],
	error) {
	req := i.prepareRequest(ctx, "GET", "/extended/v1/stx_supply")
	req.Authenticate(true)
	if height != nil {
		req.QueryParam("height", *height)
	}
	var result models.GetStxSupplyResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetStxSupplyResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetStxSupplyTotalSupplyPlain takes context as parameters and
// returns an models.ApiResponse with string data and
// an error if there was an issue with the request or response.
// Retrieves the total supply for STX tokens as plain text.
// **Note:** this uses the estimated future total supply for the year 2050.
func (i *InfoController) GetStxSupplyTotalSupplyPlain(ctx context.Context) (
	models.ApiResponse[string],
	error) {
	req := i.prepareRequest(ctx, "GET", "/extended/v1/stx_supply/total/plain")
	req.Authenticate(true)
	str, resp, err := req.CallAsText()
	var result string = str

	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	return models.NewApiResponse(result, resp), err
}

// GetStxSupplyCirculatingPlain takes context as parameters and
// returns an models.ApiResponse with string data and
// an error if there was an issue with the request or response.
// Retrieves the STX tokens currently in circulation that have been unlocked as plain text.
func (i *InfoController) GetStxSupplyCirculatingPlain(ctx context.Context) (
	models.ApiResponse[string],
	error) {
	req := i.prepareRequest(ctx, "GET", "/extended/v1/stx_supply/circulating/plain")
	req.Authenticate(true)
	str, resp, err := req.CallAsText()
	var result string = str

	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	return models.NewApiResponse(result, resp), err
}

// GetTotalStxSupplyLegacyFormat takes context, height as parameters and
// returns an models.ApiResponse with models.GetStxSupplyLegacyFormatResponse data and
// an error if there was an issue with the request or response.
// Retrieves total supply of STX tokens including those currently in circulation that have been unlocked.
// **Note:** this uses the estimated future total supply for the year 2050.
func (i *InfoController) GetTotalStxSupplyLegacyFormat(
	ctx context.Context,
	height *float64) (
	models.ApiResponse[models.GetStxSupplyLegacyFormatResponse],
	error) {
	req := i.prepareRequest(ctx, "GET", "/extended/v1/stx_supply/legacy_format")
	req.Authenticate(true)
	if height != nil {
		req.QueryParam("height", *height)
	}
	var result models.GetStxSupplyLegacyFormatResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetStxSupplyLegacyFormatResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetPoxInfo takes context as parameters and
// returns an models.ApiResponse with models.CoreNodePoxResponse data and
// an error if there was an issue with the request or response.
// Retrieves Proof-of-Transfer (PoX) information. Can be used for Stacking.
func (i *InfoController) GetPoxInfo(ctx context.Context) (
	models.ApiResponse[models.CoreNodePoxResponse],
	error) {
	req := i.prepareRequest(ctx, "GET", "/v2/pox")
	req.Authenticate(true)
	var result models.CoreNodePoxResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.CoreNodePoxResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}
