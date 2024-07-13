package api

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"

	"github.com/cbadawi/go-stacks/stacks/internal/models"
	"github.com/cbadawi/go-stacks/stacks/utils"
)

// FeesAPIService FeesAPI service
type FeesAPIService service

type ApiFetchFeeRateRequest struct {
	ctx            context.Context
	ApiService     *FeesAPIService
	feeRateRequest *models.FeeRateRequest
}

func (r ApiFetchFeeRateRequest) FeeRateRequest(feeRateRequest models.FeeRateRequest) ApiFetchFeeRateRequest {
	r.feeRateRequest = &feeRateRequest
	return r
}

func (r ApiFetchFeeRateRequest) Execute() (*models.FeeRate, *http.Response, error) {
	return r.ApiService.FetchFeeRateExecute(r)
}

/*
FetchFeeRate Fetch fee rate

**NOTE:** This endpoint is deprecated in favor of [Get approximate fees for a given transaction](/api/get-approximate-fees-for-a-given-transaction).

Retrieves estimated fee rate.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiFetchFeeRateRequest

Deprecated
*/
func (a *FeesAPIService) FetchFeeRate(ctx context.Context) ApiFetchFeeRateRequest {
	return ApiFetchFeeRateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return FeeRate
// Deprecated
func (a *FeesAPIService) FetchFeeRateExecute(r ApiFetchFeeRateRequest) (*models.FeeRate, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.FeeRate
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FeesAPIService.FetchFeeRate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/extended/v1/fee_rate"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.feeRateRequest == nil {
		return localVarReturnValue, nil, utils.ReportError("feeRateRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.feeRateRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetFeeTransferRequest struct {
	ctx        context.Context
	ApiService *FeesAPIService
}

func (r ApiGetFeeTransferRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetFeeTransferExecute(r)
}

/*
GetFeeTransfer Get estimated fee

Retrieves an estimated fee rate for STX transfer transactions. This a a fee rate / byte, and is returned as a JSON integer

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetFeeTransferRequest
*/
func (a *FeesAPIService) GetFeeTransfer(ctx context.Context) ApiGetFeeTransferRequest {
	return ApiGetFeeTransferRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *FeesAPIService) GetFeeTransferExecute(r ApiGetFeeTransferRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FeesAPIService.GetFeeTransfer")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/fees/transfer"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPostFeeTransactionRequest struct {
	ctx                           context.Context
	ApiService                    *FeesAPIService
	transactionFeeEstimateRequest *models.TransactionFeeEstimateRequest
}

func (r ApiPostFeeTransactionRequest) TransactionFeeEstimateRequest(transactionFeeEstimateRequest models.TransactionFeeEstimateRequest) ApiPostFeeTransactionRequest {
	r.transactionFeeEstimateRequest = &transactionFeeEstimateRequest
	return r
}

func (r ApiPostFeeTransactionRequest) Execute() (*models.TransactionFeeEstimateResponse, *http.Response, error) {
	return r.ApiService.PostFeeTransactionExecute(r)
}

/*
PostFeeTransaction Get approximate fees for a given transaction

Get an estimated fee for the supplied transaction.  This
estimates the execution cost of the transaction, the current
fee rate of the network, and returns estimates for fee
amounts.
* `transaction_payload` is a hex-encoded serialization of
  the TransactionPayload for the transaction.
* `estimated_len` is an optional argument that provides the
  endpoint with an estimation of the final length (in bytes)
  of the transaction, including any post-conditions and
  signatures

If the node cannot provide an estimate for the transaction
(e.g., if the node has never seen a contract-call for the
given contract and function) or if estimation is not
configured on this node, a 400 response is returned.

The 400 response will be a JSON error containing a `reason`
field which can be one of the following:
* `DatabaseError` - this Stacks node has had an internal
  database error while trying to estimate the costs of the
  supplied transaction.
* `NoEstimateAvailable` - this Stacks node has not seen this
  kind of contract-call before, and it cannot provide an
  estimate yet.
* `CostEstimationDisabled` - this Stacks node does not perform
  fee or cost estimation, and it cannot respond on this
  endpoint.

The 200 response contains the following data:
* `estimated_cost` - the estimated multi-dimensional cost of
  executing the Clarity VM on the provided transaction.
* `estimated_cost_scalar` - a unitless integer that the Stacks
  node uses to compare how much of the block limit is consumed
  by different transactions. This value incorporates the
  estimated length of the transaction and the estimated
  execution cost of the transaction. The range of this integer
  may vary between different Stacks nodes. In order to compute
  an estimate of total fee amount for the transaction, this
  value is multiplied by the same Stacks node's estimated fee
  rate.
* `cost_scalar_change_by_byte` - a float value that indicates how
  much the `estimated_cost_scalar` value would increase for every
  additional byte in the final transaction.
* `estimations` - an array of estimated fee rates and total fees to
  pay in microSTX for the transaction. This array provides a range of
  estimates (default: 3) that may be used. Each element of the array
  contains the following fields:
    * `fee_rate` - the estimated value for the current fee
      rates in the network
    * `fee` - the estimated value for the total fee in
      microSTX that the given transaction should pay. These
      values are the result of computing:
      `fee_rate` x `estimated_cost_scalar`.
      If the estimated fees are less than the minimum relay
      fee `(1 ustx x estimated_len)`, then that minimum relay
      fee will be returned here instead.

Note: If the final transaction's byte size is larger than
supplied to `estimated_len`, then applications should increase
this fee amount by:
  `fee_rate` x `cost_scalar_change_by_byte` x (`final_size` - `estimated_size`)


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostFeeTransactionRequest
*/
func (a *FeesAPIService) PostFeeTransaction(ctx context.Context) ApiPostFeeTransactionRequest {
	return ApiPostFeeTransactionRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return TransactionFeeEstimateResponse
func (a *FeesAPIService) PostFeeTransactionExecute(r ApiPostFeeTransactionRequest) (*models.TransactionFeeEstimateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.TransactionFeeEstimateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FeesAPIService.PostFeeTransaction")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/fees/transaction"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.transactionFeeEstimateRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
