package controllers

import (
	"context"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/errors"

	"github.com/cbadawi/go-stacks/stacks/models"

	"github.com/apimatic/go-core-runtime/utilities"
)

// SmartContractsController represents a controller struct.
type SmartContractsController struct {
	baseController
}

// NewSmartContractsController creates a new instance of SmartContractsController.
// It takes a baseController as a parameter and returns a pointer to the SmartContractsController.
func NewSmartContractsController(baseController baseController) *SmartContractsController {
	smartContractsController := SmartContractsController{baseController: baseController}
	return &smartContractsController
}

// GetContractById takes context, contractId, unanchored as parameters and
// returns an models.ApiResponse with models.SmartContract7 data and
// an error if there was an issue with the request or response.
// Retrieves details of a contract with a given `contract_id`
func (s *SmartContractsController) GetContractById(
	ctx context.Context,
	contractId string,
	unanchored *bool) (
	models.ApiResponse[models.SmartContract7],
	error) {
	path := fmt.Sprintf("/extended/v1/contract/%v", contractId)
	req := s.prepareRequest(
		ctx,
		"GET",
		path,
	)
	req.Authenticate(true)
	if unanchored != nil {
		req.QueryParam("unanchored", *unanchored)
	}

	var result models.SmartContract7
	decoder, resp, err := s.LogCallAsJSON(req, path)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.SmartContract7](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	if resp.StatusCode == 404 {
		err = errors.NewApiError(404, "Cannot find contract of given ID")
	}
	return models.NewApiResponse(result, resp), err
}

// GetContractsByTrait takes context, traitAbi, limit, offset as parameters and
// returns an models.ApiResponse with models.ContractListResponse data and
// an error if there was an issue with the request or response.
// Retrieves a list of contracts based on the following traits listed in JSON format -  functions, variables, maps, fungible tokens and non-fungible tokens
func (s *SmartContractsController) GetContractsByTrait(
	ctx context.Context,
	traitAbi string,
	limit *int,
	offset *int) (
	models.ApiResponse[models.ContractListResponse],
	error) {
	path := "/extended/v1/contract/by_trait"
	req := s.prepareRequest(ctx, "GET", path)
	req.Authenticate(true)
	req.QueryParam("trait_abi", traitAbi)
	if limit != nil {
		req.QueryParam("limit", *limit)
	}
	if offset != nil {
		req.QueryParam("offset", *offset)
	}
	var result models.ContractListResponse
	decoder, resp, err := s.LogCallAsJSON(req, path)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ContractListResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetContractEventsById takes context, contractId, limit, offset, unanchored as parameters and
// returns an models.ApiResponse with models.TransactionEvent data and
// an error if there was an issue with the request or response.
// Retrieves a list of events that have been triggered by a given `contract_id`
func (s *SmartContractsController) GetContractEventsById(
	ctx context.Context,
	contractId string,
	limit *int,
	offset *int,
	unanchored *bool) (
	models.ApiResponse[models.TransactionEvent],
	error) {
	path := fmt.Sprintf("/extended/v1/contract/%v/events", contractId)
	req := s.prepareRequest(
		ctx,
		"GET",
		path,
	)
	req.Authenticate(true)
	if limit != nil {
		req.QueryParam("limit", *limit)
	}
	if offset != nil {
		req.QueryParam("offset", *offset)
	}
	if unanchored != nil {
		req.QueryParam("unanchored", *unanchored)
	}

	var result models.TransactionEvent
	decoder, resp, err := s.LogCallAsJSON(req, path)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.TransactionEvent](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetContractInterface takes context, contractAddress, contractName, tip as parameters and
// returns an models.ApiResponse with models.ContractInterfaceResponse data and
// an error if there was an issue with the request or response.
// Retrieves a contract interface with a given `contract_address` and `contract name`
func (s *SmartContractsController) GetContractInterface(
	ctx context.Context,
	contractAddress string,
	contractName string,
	tip *string) (
	models.ApiResponse[models.ContractInterfaceResponse],
	error) {
	path := fmt.Sprintf("/v2/contracts/interface/%v/%v", contractAddress, contractName)
	req := s.prepareRequest(
		ctx,
		"GET",
		path,
	)
	req.Authenticate(true)
	if tip != nil {
		req.QueryParam("tip", *tip)
	}

	var result models.ContractInterfaceResponse
	decoder, resp, err := s.LogCallAsJSON(req, path)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ContractInterfaceResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetContractDataMapEntry takes context, contractAddress, contractName, mapName, body, proof, tip as parameters and
// returns an models.ApiResponse with models.MapEntryResponse data and
// an error if there was an issue with the request or response.
// Attempt to fetch data from a contract data map. The contract is identified with Stacks Address `contract_address` and Contract Name `contract_address` in the URL path. The map is identified with [Map Name].
// The key to lookup in the map is supplied via the POST body. This should be supplied as the hex string serialization of the key (which should be a Clarity value). Note, this is a JSON string atom.
// In the response, `data` is the hex serialization of the map response. Note that map responses are Clarity option types, for non-existent values, this is a serialized none, and for all other responses, it is a serialized (some ...) object.
func (s *SmartContractsController) GetContractDataMapEntry(
	ctx context.Context,
	contractAddress string,
	contractName string,
	mapName string,
	body string,
	proof *int,
	tip *string) (
	models.ApiResponse[models.MapEntryResponse],
	error) {
	path := fmt.Sprintf("/v2/map_entry/%v/%v/%v", contractAddress, contractName, mapName)
	req := s.prepareRequest(
		ctx,
		"POST",
		path,
	)
	req.Authenticate(true)
	req.Header("Content-Type", "application/json")
	if proof != nil {
		req.QueryParam("proof", *proof)
	}
	if tip != nil {
		req.QueryParam("tip", *tip)
	}
	req.Text(body)

	var result models.MapEntryResponse
	decoder, resp, err := s.LogCallAsJSON(req, path)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.MapEntryResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	if resp.StatusCode == 400 {
		err = errors.NewApiError(400, "Failed loading data map")
	}
	return models.NewApiResponse(result, resp), err
}

// GetContractSource takes context, contractAddress, contractName, proof, tip as parameters and
// returns an models.ApiResponse with models.ContractSourceResponse data and
// an error if there was an issue with the request or response.
// Retrieves the Clarity source code of a given contract, along with the block height it was published in, and the MARF proof for the data
func (s *SmartContractsController) GetContractSource(
	ctx context.Context,
	contractAddress string,
	contractName string,
	proof *int,
	tip *string) (
	models.ApiResponse[models.ContractSourceResponse],
	error) {
	path := fmt.Sprintf("/v2/contracts/source/%v/%v", contractAddress, contractName)
	req := s.prepareRequest(
		ctx,
		"GET",
		path,
	)
	req.Authenticate(true)
	if proof != nil {
		req.QueryParam("proof", *proof)
	}
	if tip != nil {
		req.QueryParam("tip", *tip)
	}

	var result models.ContractSourceResponse
	decoder, resp, err := s.LogCallAsJSON(req, path)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ContractSourceResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// CallReadOnlyFunction takes context, contractAddress, contractName, functionName, body, tip as parameters and
// returns an models.ApiResponse with models.ReadOnlyFunctionSuccessResponse data and
// an error if there was an issue with the request or response.
// Call a read-only public function on a given smart contract.
// The smart contract and function are specified using the URL path. The arguments and the simulated tx-sender are supplied via the POST body in the following JSON format:
func (s *SmartContractsController) CallReadOnlyFunction(
	ctx context.Context,
	contractAddress string,
	contractName string,
	functionName string,
	body *models.ReadOnlyFunctionArgs,
	tip *string) (
	models.ApiResponse[models.ReadOnlyFunctionSuccessResponse],
	error) {
	path := fmt.Sprintf("/v2/contracts/call-read/%v/%v/%v", contractAddress, contractName, functionName)
	req := s.prepareRequest(
		ctx,
		"POST",
		path,
	)
	req.Authenticate(true)
	req.Header("Content-Type", "application/json")
	if tip != nil {
		req.QueryParam("tip", *tip)
	}
	req.Json(body)

	var result models.ReadOnlyFunctionSuccessResponse
	decoder, resp, err := s.LogCallAsJSON(req, path)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ReadOnlyFunctionSuccessResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}
