package stacksblockchainapi

import (
	"context"
	"stacksblockchainapi/models"

	"github.com/apimatic/go-core-runtime/utilities"
)

// FeesController represents a controller struct.
type FeesController struct {
	baseController
}

// NewFeesController creates a new instance of FeesController.
// It takes a baseController as a parameter and returns a pointer to the FeesController.
func NewFeesController(baseController baseController) *FeesController {
	feesController := FeesController{baseController: baseController}
	return &feesController
}

// GetFeeTransfer takes context as parameters and
// returns an models.ApiResponse with string data and
// an error if there was an issue with the request or response.
// Retrieves an estimated fee rate for STX transfer transactions. This a a fee rate / byte, and is returned as a JSON integer
func (f *FeesController) GetFeeTransfer(ctx context.Context) (
	models.ApiResponse[string],
	error) {
	req := f.prepareRequest(ctx, "GET", "/v2/fees/transfer")
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

// FetchFeeRate takes context, body as parameters and
// returns an models.ApiResponse with models.FeeRate data and
// an error if there was an issue with the request or response.
// Deprecated: fetch_fee_rate is deprecated
// **NOTE:** This endpoint is deprecated in favor of [Get approximate fees for a given transaction](#operation/post_fee_transaction).
// Retrieves estimated fee rate.
func (f *FeesController) FetchFeeRate(
	ctx context.Context,
	body *models.FeeRateRequest) (
	models.ApiResponse[models.FeeRate],
	error) {
	req := f.prepareRequest(ctx, "POST", "/extended/v1/fee_rate")
	req.Authenticate(true)
	req.Header("Content-Type", "application/json")
	req.Json(body)
	var result models.FeeRate
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.FeeRate](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// PostFeeTransaction takes context, body as parameters and
// returns an models.ApiResponse with models.TransactionFeeEstimateResponse data and
// an error if there was an issue with the request or response.
// Get an estimated fee for the supplied transaction.  This
// estimates the execution cost of the transaction, the current
// fee rate of the network, and returns estimates for fee
// amounts.
// * `transaction_payload` is a hex-encoded serialization of
// the TransactionPayload for the transaction.
// * `estimated_len` is an optional argument that provides the
// endpoint with an estimation of the final length (in bytes)
// of the transaction, including any post-conditions and
// signatures
// If the node cannot provide an estimate for the transaction
// (e.g., if the node has never seen a contract-call for the
// given contract and function) or if estimation is not
// configured on this node, a 400 response is returned.
// The 400 response will be a JSON error containing a `reason`
// field which can be one of the following:
// * `DatabaseError` - this Stacks node has had an internal
// database error while trying to estimate the costs of the
// supplied transaction.
// * `NoEstimateAvailable` - this Stacks node has not seen this
// kind of contract-call before, and it cannot provide an
// estimate yet.
// * `CostEstimationDisabled` - this Stacks node does not perform
// fee or cost estimation, and it cannot respond on this
// endpoint.
// The 200 response contains the following data:
// * `estimated_cost` - the estimated multi-dimensional cost of
// executing the Clarity VM on the provided transaction.
// * `estimated_cost_scalar` - a unitless integer that the Stacks
// node uses to compare how much of the block limit is consumed
// by different transactions. This value incorporates the
// estimated length of the transaction and the estimated
// execution cost of the transaction. The range of this integer
// may vary between different Stacks nodes. In order to compute
// an estimate of total fee amount for the transaction, this
// value is multiplied by the same Stacks node's estimated fee
// rate.
// * `cost_scalar_change_by_byte` - a float value that indicates how
// much the `estimated_cost_scalar` value would increase for every
// additional byte in the final transaction.
// * `estimations` - an array of estimated fee rates and total fees to
// pay in microSTX for the transaction. This array provides a range of
// estimates (default: 3) that may be used. Each element of the array
// contains the following fields:
// * `fee_rate` - the estimated value for the current fee
// rates in the network
// * `fee` - the estimated value for the total fee in
// microSTX that the given transaction should pay. These
// values are the result of computing:
// `fee_rate` x `estimated_cost_scalar`.
// If the estimated fees are less than the minimum relay
// fee `(1 ustx x estimated_len)`, then that minimum relay
// fee will be returned here instead.
// Note: If the final transaction's byte size is larger than
// supplied to `estimated_len`, then applications should increase
// this fee amount by:
// `fee_rate` x `cost_scalar_change_by_byte` x (`final_size` - `estimated_size`)
func (f *FeesController) PostFeeTransaction(
	ctx context.Context,
	body *models.TransactionFeeEstimateRequest) (
	models.ApiResponse[models.TransactionFeeEstimateResponse],
	error) {
	req := f.prepareRequest(ctx, "POST", "/v2/fees/transaction")
	req.Authenticate(true)
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(*body)
	}
	var result models.TransactionFeeEstimateResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.TransactionFeeEstimateResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}
