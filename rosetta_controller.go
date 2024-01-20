package main

import (
	"context"

	"github.com/cbadawi/stacks-go-draft/errors"

	"github.com/cbadawi/stacks-go-draft/models"

	"github.com/apimatic/go-core-runtime/utilities"
)

// RosettaController represents a controller struct.
type RosettaController struct {
	baseController
}

// NewRosettaController creates a new instance of RosettaController.
// It takes a baseController as a parameter and returns a pointer to the RosettaController.
func NewRosettaController(baseController baseController) *RosettaController {
	rosettaController := RosettaController{baseController: baseController}
	return &rosettaController
}

// RosettaNetworkList takes context as parameters and
// returns an models.ApiResponse with models.RosettaNetworkListResponse data and
// an error if there was an issue with the request or response.
// Retrieves a list of NetworkIdentifiers that the Rosetta server supports.
func (r *RosettaController) RosettaNetworkList(ctx context.Context) (
	models.ApiResponse[models.RosettaNetworkListResponse],
	error) {
	req := r.prepareRequest(ctx, "POST", "/rosetta/v1/network/list")
	req.Authenticate(true)
	var result models.RosettaNetworkListResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.RosettaNetworkListResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	if resp.StatusCode == 400 {
		err = errors.NewRosettaError(400, "Error")
	}
	return models.NewApiResponse(result, resp), err
}

// RosettaNetworkOptions takes context, body as parameters and
// returns an models.ApiResponse with models.RosettaNetworkOptionsResponse data and
// an error if there was an issue with the request or response.
// Retrieves the version information and allowed network-specific types for a NetworkIdentifier.
// Any NetworkIdentifier returned by /network/list should be accessible here.
// Because options are retrievable in the context of a NetworkIdentifier, it is possible to define unique options for each network.
func (r *RosettaController) RosettaNetworkOptions(
	ctx context.Context,
	body *models.RosettaOptionsRequest) (
	models.ApiResponse[models.RosettaNetworkOptionsResponse],
	error) {
	req := r.prepareRequest(ctx, "POST", "/rosetta/v1/network/options")
	req.Authenticate(true)
	req.Header("Content-Type", "application/json")
	req.Json(body)
	var result models.RosettaNetworkOptionsResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.RosettaNetworkOptionsResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	if resp.StatusCode == 400 {
		err = errors.NewRosettaError(400, "Error")
	}
	return models.NewApiResponse(result, resp), err
}

// RosettaNetworkStatus takes context, body as parameters and
// returns an models.ApiResponse with models.RosettaNetworkStatusResponse data and
// an error if there was an issue with the request or response.
// Retrieves the current status of the network requested.
// Any NetworkIdentifier returned by /network/list should be accessible here.
func (r *RosettaController) RosettaNetworkStatus(
	ctx context.Context,
	body *models.RosettaStatusRequest) (
	models.ApiResponse[models.RosettaNetworkStatusResponse],
	error) {
	req := r.prepareRequest(ctx, "POST", "/rosetta/v1/network/status")
	req.Authenticate(true)
	req.Header("Content-Type", "application/json")
	req.Json(body)
	var result models.RosettaNetworkStatusResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.RosettaNetworkStatusResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	if resp.StatusCode == 400 {
		err = errors.NewRosettaError(400, "Error")
	}
	return models.NewApiResponse(result, resp), err
}

// RosettaAccountBalance takes context, body as parameters and
// returns an models.ApiResponse with models.RosettaAccountBalanceResponse data and
// an error if there was an issue with the request or response.
// An AccountBalanceRequest is utilized to make a balance request on the /account/balance endpoint.
// If the block_identifier is populated, a historical balance query should be performed.
func (r *RosettaController) RosettaAccountBalance(
	ctx context.Context,
	body *models.RosettaAccountBalanceRequest) (
	models.ApiResponse[models.RosettaAccountBalanceResponse],
	error) {
	req := r.prepareRequest(ctx, "POST", "/rosetta/v1/account/balance")
	req.Authenticate(true)
	req.Header("Content-Type", "application/json")
	req.Json(body)
	var result models.RosettaAccountBalanceResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.RosettaAccountBalanceResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	if resp.StatusCode == 400 {
		err = errors.NewRosettaError(400, "Error")
	}
	return models.NewApiResponse(result, resp), err
}

// RosettaBlock takes context, body as parameters and
// returns an models.ApiResponse with models.RosettaBlockResponse data and
// an error if there was an issue with the request or response.
// Retrieves the Block information for a given block identifier including a list of all transactions in the block.
func (r *RosettaController) RosettaBlock(
	ctx context.Context,
	body *models.RosettaBlockRequest) (
	models.ApiResponse[models.RosettaBlockResponse],
	error) {
	req := r.prepareRequest(ctx, "POST", "/rosetta/v1/block")
	req.Authenticate(true)
	req.Header("Content-Type", "application/json")
	req.Json(body)
	var result models.RosettaBlockResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.RosettaBlockResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	if resp.StatusCode == 400 {
		err = errors.NewRosettaError(400, "Error")
	}
	return models.NewApiResponse(result, resp), err
}

// RosettaBlockTransaction takes context, body as parameters and
// returns an models.ApiResponse with models.RosettaBlockTransactionResponse data and
// an error if there was an issue with the request or response.
// Retrieves a Transaction included in a block that is not returned in a BlockResponse.
func (r *RosettaController) RosettaBlockTransaction(
	ctx context.Context,
	body *models.RosettaBlockTransactionRequest) (
	models.ApiResponse[models.RosettaBlockTransactionResponse],
	error) {
	req := r.prepareRequest(ctx, "POST", "/rosetta/v1/block/transaction")
	req.Authenticate(true)
	req.Header("Content-Type", "application/json")
	req.Json(body)
	var result models.RosettaBlockTransactionResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.RosettaBlockTransactionResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	if resp.StatusCode == 400 {
		err = errors.NewRosettaError(400, "Error")
	}
	return models.NewApiResponse(result, resp), err
}

// RosettaMempool takes context, body as parameters and
// returns an models.ApiResponse with models.RosettaMempoolResponse data and
// an error if there was an issue with the request or response.
// Retrieves a list of transactions currently in the mempool for a given network.
func (r *RosettaController) RosettaMempool(
	ctx context.Context,
	body *models.RosettaMempoolRequest) (
	models.ApiResponse[models.RosettaMempoolResponse],
	error) {
	req := r.prepareRequest(ctx, "POST", "/rosetta/v1/mempool")
	req.Authenticate(true)
	req.Header("Content-Type", "application/json")
	req.Json(body)
	var result models.RosettaMempoolResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.RosettaMempoolResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	if resp.StatusCode == 400 {
		err = errors.NewRosettaError(400, "Error")
	}
	return models.NewApiResponse(result, resp), err
}

// RosettaMempoolTransaction takes context, body as parameters and
// returns an models.ApiResponse with models.RosettaMempoolTransactionResponse data and
// an error if there was an issue with the request or response.
// Retrieves transaction details from the mempool for a given transaction id from a given network.
func (r *RosettaController) RosettaMempoolTransaction(
	ctx context.Context,
	body *models.RosettaMempoolTransactionRequest) (
	models.ApiResponse[models.RosettaMempoolTransactionResponse],
	error) {
	req := r.prepareRequest(ctx, "POST", "/rosetta/v1/mempool/transaction")
	req.Authenticate(true)
	req.Header("Content-Type", "application/json")
	req.Json(body)
	var result models.RosettaMempoolTransactionResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.RosettaMempoolTransactionResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	if resp.StatusCode == 400 {
		err = errors.NewRosettaError(400, "Error")
	}
	return models.NewApiResponse(result, resp), err
}

// RosettaConstructionDerive takes context, body as parameters and
// returns an models.ApiResponse with models.RosettaConstructionDeriveResponse data and
// an error if there was an issue with the request or response.
// Retrieves the Account Identifier information based on a Public Key for a given network
func (r *RosettaController) RosettaConstructionDerive(
	ctx context.Context,
	body *models.RosettaConstructionDeriveRequest) (
	models.ApiResponse[models.RosettaConstructionDeriveResponse],
	error) {
	req := r.prepareRequest(ctx, "POST", "/rosetta/v1/construction/derive")
	req.Authenticate(true)
	req.Header("Content-Type", "application/json")
	req.Json(body)
	var result models.RosettaConstructionDeriveResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.RosettaConstructionDeriveResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	if resp.StatusCode == 400 {
		err = errors.NewRosettaError(400, "Error")
	}
	return models.NewApiResponse(result, resp), err
}

// RosettaConstructionHash takes context, body as parameters and
// returns an models.ApiResponse with models.RosettaConstructionHashResponse data and
// an error if there was an issue with the request or response.
// Retrieves the network-specific transaction hash for a signed transaction.
func (r *RosettaController) RosettaConstructionHash(
	ctx context.Context,
	body *models.RosettaConstructionHashRequest) (
	models.ApiResponse[models.RosettaConstructionHashResponse],
	error) {
	req := r.prepareRequest(ctx, "POST", "/rosetta/v1/construction/hash")
	req.Authenticate(true)
	req.Header("Content-Type", "application/json")
	req.Json(body)
	var result models.RosettaConstructionHashResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.RosettaConstructionHashResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	if resp.StatusCode == 400 {
		err = errors.NewRosettaError(400, "Error")
	}
	return models.NewApiResponse(result, resp), err
}

// RosettaConstructionMetadata takes context, body as parameters and
// returns an models.ApiResponse with models.RosettaConstructionMetadataResponse data and
// an error if there was an issue with the request or response.
// To Do
func (r *RosettaController) RosettaConstructionMetadata(
	ctx context.Context,
	body *models.RosettaConstructionMetadataRequest) (
	models.ApiResponse[models.RosettaConstructionMetadataResponse],
	error) {
	req := r.prepareRequest(ctx, "POST", "/rosetta/v1/construction/metadata")
	req.Authenticate(true)
	req.Header("Content-Type", "application/json")
	req.Json(body)
	var result models.RosettaConstructionMetadataResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.RosettaConstructionMetadataResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	if resp.StatusCode == 400 {
		err = errors.NewRosettaError(400, "Error")
	}
	return models.NewApiResponse(result, resp), err
}

// RosettaConstructionParse takes context, body as parameters and
// returns an models.ApiResponse with models.RosettaConstructionParseResponse data and
// an error if there was an issue with the request or response.
// TODO
func (r *RosettaController) RosettaConstructionParse(
	ctx context.Context,
	body *models.RosettaConstructionParseRequest) (
	models.ApiResponse[models.RosettaConstructionParseResponse],
	error) {
	req := r.prepareRequest(ctx, "POST", "/rosetta/v1/construction/parse")
	req.Authenticate(true)
	req.Header("Content-Type", "application/json")
	req.Json(body)
	var result models.RosettaConstructionParseResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.RosettaConstructionParseResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	if resp.StatusCode == 400 {
		err = errors.NewRosettaError(400, "Error")
	}
	return models.NewApiResponse(result, resp), err
}

// RosettaConstructionPreprocess takes context, body as parameters and
// returns an models.ApiResponse with models.RosettaConstructionPreprocessResponse data and
// an error if there was an issue with the request or response.
// TODO
func (r *RosettaController) RosettaConstructionPreprocess(
	ctx context.Context,
	body *models.RosettaConstructionPreprocessRequest) (
	models.ApiResponse[models.RosettaConstructionPreprocessResponse],
	error) {
	req := r.prepareRequest(ctx, "POST", "/rosetta/v1/construction/preprocess")
	req.Authenticate(true)
	req.Header("Content-Type", "application/json")
	req.Json(body)
	var result models.RosettaConstructionPreprocessResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.RosettaConstructionPreprocessResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	if resp.StatusCode == 400 {
		err = errors.NewRosettaError(400, "Error")
	}
	return models.NewApiResponse(result, resp), err
}

// RosettaConstructionSubmit takes context, body as parameters and
// returns an models.ApiResponse with models.RosettaConstructionSubmitResponse data and
// an error if there was an issue with the request or response.
// Submit a pre-signed transaction to the node. The examples below are illustrative only. You'll need to use your wallet to generate actual values to use them in the request payload.
func (r *RosettaController) RosettaConstructionSubmit(
	ctx context.Context,
	body *models.RosettaConstructionSubmitRequest) (
	models.ApiResponse[models.RosettaConstructionSubmitResponse],
	error) {
	req := r.prepareRequest(ctx, "POST", "/rosetta/v1/construction/submit")
	req.Authenticate(true)
	req.Header("Content-Type", "application/json")
	req.Json(body)
	var result models.RosettaConstructionSubmitResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.RosettaConstructionSubmitResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	if resp.StatusCode == 400 {
		err = errors.NewRosettaError(400, "Error")
	}
	return models.NewApiResponse(result, resp), err
}

// RosettaConstructionPayloads takes context, body as parameters and
// returns an models.ApiResponse with models.RosettaConstructionPayloadResponse data and
// an error if there was an issue with the request or response.
// Generate an unsigned transaction from operations and metadata. The examples below are illustrative only. You'll need to use your wallet to generate actual values to use them in the request payload.
func (r *RosettaController) RosettaConstructionPayloads(
	ctx context.Context,
	body *models.RosettaConstructionPayloadsRequest) (
	models.ApiResponse[models.RosettaConstructionPayloadResponse],
	error) {
	req := r.prepareRequest(ctx, "POST", "/rosetta/v1/construction/payloads")
	req.Authenticate(true)
	req.Header("Content-Type", "application/json")
	req.Json(body)
	var result models.RosettaConstructionPayloadResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.RosettaConstructionPayloadResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	if resp.StatusCode == 400 {
		err = errors.NewRosettaError(400, "Error")
	}
	return models.NewApiResponse(result, resp), err
}

// RosettaConstructionCombine takes context, body as parameters and
// returns an models.ApiResponse with models.RosettaConstructionCombineResponse data and
// an error if there was an issue with the request or response.
// Take unsigned transaction and signature, combine both and return signed transaction. The examples below are illustrative only. You'll need to use your wallet to generate actual values to use them in the request payload.
func (r *RosettaController) RosettaConstructionCombine(
	ctx context.Context,
	body *models.RosettaConstructionCombineRequest) (
	models.ApiResponse[models.RosettaConstructionCombineResponse],
	error) {
	req := r.prepareRequest(ctx, "POST", "/rosetta/v1/construction/combine")
	req.Authenticate(true)
	req.Header("Content-Type", "application/json")
	req.Json(body)
	var result models.RosettaConstructionCombineResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.RosettaConstructionCombineResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	if resp.StatusCode == 400 {
		err = errors.NewRosettaError(400, "Error")
	}
	return models.NewApiResponse(result, resp), err
}
