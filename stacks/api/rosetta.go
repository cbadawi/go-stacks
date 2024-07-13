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

// RosettaAPIService RosettaAPI service
type RosettaAPIService service

type ApiRosettaAccountBalanceRequest struct {
	ctx                          context.Context
	ApiService                   *RosettaAPIService
	rosettaAccountBalanceRequest *models.RosettaAccountBalanceRequest
}

func (r ApiRosettaAccountBalanceRequest) RosettaAccountBalanceRequest(rosettaAccountBalanceRequest models.RosettaAccountBalanceRequest) ApiRosettaAccountBalanceRequest {
	r.rosettaAccountBalanceRequest = &rosettaAccountBalanceRequest
	return r
}

func (r ApiRosettaAccountBalanceRequest) Execute() (*models.RosettaAccountBalanceResponse, *http.Response, error) {
	return r.ApiService.RosettaAccountBalanceExecute(r)
}

/*
RosettaAccountBalance Get an Account Balance

An AccountBalanceRequest is utilized to make a balance request on the /account/balance endpoint.
If the block_identifier is populated, a historical balance query should be performed.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRosettaAccountBalanceRequest
*/
func (a *RosettaAPIService) RosettaAccountBalance(ctx context.Context) ApiRosettaAccountBalanceRequest {
	return ApiRosettaAccountBalanceRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return RosettaAccountBalanceResponse
func (a *RosettaAPIService) RosettaAccountBalanceExecute(r ApiRosettaAccountBalanceRequest) (*models.RosettaAccountBalanceResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.RosettaAccountBalanceResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RosettaAPIService.RosettaAccountBalance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/rosetta/v1/account/balance"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.rosettaAccountBalanceRequest == nil {
		return localVarReturnValue, nil, utils.ReportError("rosettaAccountBalanceRequest is required and must be specified")
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
	localVarPostBody = r.rosettaAccountBalanceRequest
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v models.RosettaError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
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

type ApiRosettaBlockRequest struct {
	ctx                 context.Context
	ApiService          *RosettaAPIService
	rosettaBlockRequest *models.RosettaBlockRequest
}

func (r ApiRosettaBlockRequest) RosettaBlockRequest(rosettaBlockRequest models.RosettaBlockRequest) ApiRosettaBlockRequest {
	r.rosettaBlockRequest = &rosettaBlockRequest
	return r
}

func (r ApiRosettaBlockRequest) Execute() (*models.RosettaBlockResponse, *http.Response, error) {
	return r.ApiService.RosettaBlockExecute(r)
}

/*
RosettaBlock Get a Block

Retrieves the Block information for a given block identifier including a list of all transactions in the block.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRosettaBlockRequest
*/
func (a *RosettaAPIService) RosettaBlock(ctx context.Context) ApiRosettaBlockRequest {
	return ApiRosettaBlockRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return RosettaBlockResponse
func (a *RosettaAPIService) RosettaBlockExecute(r ApiRosettaBlockRequest) (*models.RosettaBlockResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.RosettaBlockResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RosettaAPIService.RosettaBlock")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/rosetta/v1/block"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.rosettaBlockRequest == nil {
		return localVarReturnValue, nil, utils.ReportError("rosettaBlockRequest is required and must be specified")
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
	localVarPostBody = r.rosettaBlockRequest
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v models.RosettaError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
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

type ApiRosettaBlockTransactionRequest struct {
	ctx                            context.Context
	ApiService                     *RosettaAPIService
	rosettaBlockTransactionRequest *models.RosettaBlockTransactionRequest
}

func (r ApiRosettaBlockTransactionRequest) RosettaBlockTransactionRequest(rosettaBlockTransactionRequest models.RosettaBlockTransactionRequest) ApiRosettaBlockTransactionRequest {
	r.rosettaBlockTransactionRequest = &rosettaBlockTransactionRequest
	return r
}

func (r ApiRosettaBlockTransactionRequest) Execute() (*models.RosettaBlockTransactionResponse, *http.Response, error) {
	return r.ApiService.RosettaBlockTransactionExecute(r)
}

/*
RosettaBlockTransaction Get a Block Transaction

Retrieves a Transaction included in a block that is not returned in a BlockResponse.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRosettaBlockTransactionRequest
*/
func (a *RosettaAPIService) RosettaBlockTransaction(ctx context.Context) ApiRosettaBlockTransactionRequest {
	return ApiRosettaBlockTransactionRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return RosettaBlockTransactionResponse
func (a *RosettaAPIService) RosettaBlockTransactionExecute(r ApiRosettaBlockTransactionRequest) (*models.RosettaBlockTransactionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.RosettaBlockTransactionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RosettaAPIService.RosettaBlockTransaction")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/rosetta/v1/block/transaction"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.rosettaBlockTransactionRequest == nil {
		return localVarReturnValue, nil, utils.ReportError("rosettaBlockTransactionRequest is required and must be specified")
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
	localVarPostBody = r.rosettaBlockTransactionRequest
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v models.RosettaError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
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

type ApiRosettaConstructionCombineRequest struct {
	ctx                               context.Context
	ApiService                        *RosettaAPIService
	rosettaConstructionCombineRequest *models.RosettaConstructionCombineRequest
}

func (r ApiRosettaConstructionCombineRequest) RosettaConstructionCombineRequest(rosettaConstructionCombineRequest models.RosettaConstructionCombineRequest) ApiRosettaConstructionCombineRequest {
	r.rosettaConstructionCombineRequest = &rosettaConstructionCombineRequest
	return r
}

func (r ApiRosettaConstructionCombineRequest) Execute() (*models.RosettaConstructionCombineResponse, *http.Response, error) {
	return r.ApiService.RosettaConstructionCombineExecute(r)
}

/*
RosettaConstructionCombine Create Network Transaction from Signatures

Take unsigned transaction and signature, combine both and return signed transaction. The examples below are illustrative only. You'll need to use your wallet to generate actual values to use them in the request payload.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRosettaConstructionCombineRequest
*/
func (a *RosettaAPIService) RosettaConstructionCombine(ctx context.Context) ApiRosettaConstructionCombineRequest {
	return ApiRosettaConstructionCombineRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return RosettaConstructionCombineResponse
func (a *RosettaAPIService) RosettaConstructionCombineExecute(r ApiRosettaConstructionCombineRequest) (*models.RosettaConstructionCombineResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.RosettaConstructionCombineResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RosettaAPIService.RosettaConstructionCombine")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/rosetta/v1/construction/combine"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.rosettaConstructionCombineRequest == nil {
		return localVarReturnValue, nil, utils.ReportError("rosettaConstructionCombineRequest is required and must be specified")
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
	localVarPostBody = r.rosettaConstructionCombineRequest
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v models.RosettaError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
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

type ApiRosettaConstructionDeriveRequest struct {
	ctx                              context.Context
	ApiService                       *RosettaAPIService
	rosettaConstructionDeriveRequest *models.RosettaConstructionDeriveRequest
}

func (r ApiRosettaConstructionDeriveRequest) RosettaConstructionDeriveRequest(rosettaConstructionDeriveRequest models.RosettaConstructionDeriveRequest) ApiRosettaConstructionDeriveRequest {
	r.rosettaConstructionDeriveRequest = &rosettaConstructionDeriveRequest
	return r
}

func (r ApiRosettaConstructionDeriveRequest) Execute() (*models.RosettaConstructionDeriveResponse, *http.Response, error) {
	return r.ApiService.RosettaConstructionDeriveExecute(r)
}

/*
RosettaConstructionDerive Derive an AccountIdentifier from a PublicKey

Retrieves the Account Identifier information based on a Public Key for a given network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRosettaConstructionDeriveRequest
*/
func (a *RosettaAPIService) RosettaConstructionDerive(ctx context.Context) ApiRosettaConstructionDeriveRequest {
	return ApiRosettaConstructionDeriveRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return RosettaConstructionDeriveResponse
func (a *RosettaAPIService) RosettaConstructionDeriveExecute(r ApiRosettaConstructionDeriveRequest) (*models.RosettaConstructionDeriveResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.RosettaConstructionDeriveResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RosettaAPIService.RosettaConstructionDerive")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/rosetta/v1/construction/derive"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.rosettaConstructionDeriveRequest == nil {
		return localVarReturnValue, nil, utils.ReportError("rosettaConstructionDeriveRequest is required and must be specified")
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
	localVarPostBody = r.rosettaConstructionDeriveRequest
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v models.RosettaError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
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

type ApiRosettaConstructionHashRequest struct {
	ctx                            context.Context
	ApiService                     *RosettaAPIService
	rosettaConstructionHashRequest *models.RosettaConstructionHashRequest
}

func (r ApiRosettaConstructionHashRequest) RosettaConstructionHashRequest(rosettaConstructionHashRequest models.RosettaConstructionHashRequest) ApiRosettaConstructionHashRequest {
	r.rosettaConstructionHashRequest = &rosettaConstructionHashRequest
	return r
}

func (r ApiRosettaConstructionHashRequest) Execute() (*models.RosettaConstructionHashResponse, *http.Response, error) {
	return r.ApiService.RosettaConstructionHashExecute(r)
}

/*
RosettaConstructionHash Get the Hash of a Signed Transaction

Retrieves the network-specific transaction hash for a signed transaction.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRosettaConstructionHashRequest
*/
func (a *RosettaAPIService) RosettaConstructionHash(ctx context.Context) ApiRosettaConstructionHashRequest {
	return ApiRosettaConstructionHashRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return RosettaConstructionHashResponse
func (a *RosettaAPIService) RosettaConstructionHashExecute(r ApiRosettaConstructionHashRequest) (*models.RosettaConstructionHashResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.RosettaConstructionHashResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RosettaAPIService.RosettaConstructionHash")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/rosetta/v1/construction/hash"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.rosettaConstructionHashRequest == nil {
		return localVarReturnValue, nil, utils.ReportError("rosettaConstructionHashRequest is required and must be specified")
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
	localVarPostBody = r.rosettaConstructionHashRequest
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v models.RosettaError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
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

type ApiRosettaConstructionMetadataRequest struct {
	ctx                                context.Context
	ApiService                         *RosettaAPIService
	rosettaConstructionMetadataRequest *models.RosettaConstructionMetadataRequest
}

func (r ApiRosettaConstructionMetadataRequest) RosettaConstructionMetadataRequest(rosettaConstructionMetadataRequest models.RosettaConstructionMetadataRequest) ApiRosettaConstructionMetadataRequest {
	r.rosettaConstructionMetadataRequest = &rosettaConstructionMetadataRequest
	return r
}

func (r ApiRosettaConstructionMetadataRequest) Execute() (*models.RosettaConstructionMetadataResponse, *http.Response, error) {
	return r.ApiService.RosettaConstructionMetadataExecute(r)
}

/*
RosettaConstructionMetadata Get Metadata for Transaction Construction

To Do

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRosettaConstructionMetadataRequest
*/
func (a *RosettaAPIService) RosettaConstructionMetadata(ctx context.Context) ApiRosettaConstructionMetadataRequest {
	return ApiRosettaConstructionMetadataRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return RosettaConstructionMetadataResponse
func (a *RosettaAPIService) RosettaConstructionMetadataExecute(r ApiRosettaConstructionMetadataRequest) (*models.RosettaConstructionMetadataResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.RosettaConstructionMetadataResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RosettaAPIService.RosettaConstructionMetadata")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/rosetta/v1/construction/metadata"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.rosettaConstructionMetadataRequest == nil {
		return localVarReturnValue, nil, utils.ReportError("rosettaConstructionMetadataRequest is required and must be specified")
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
	localVarPostBody = r.rosettaConstructionMetadataRequest
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v models.RosettaError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
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

type ApiRosettaConstructionParseRequest struct {
	ctx                             context.Context
	ApiService                      *RosettaAPIService
	rosettaConstructionParseRequest *models.RosettaConstructionParseRequest
}

func (r ApiRosettaConstructionParseRequest) RosettaConstructionParseRequest(rosettaConstructionParseRequest models.RosettaConstructionParseRequest) ApiRosettaConstructionParseRequest {
	r.rosettaConstructionParseRequest = &rosettaConstructionParseRequest
	return r
}

func (r ApiRosettaConstructionParseRequest) Execute() (*models.RosettaConstructionParseResponse, *http.Response, error) {
	return r.ApiService.RosettaConstructionParseExecute(r)
}

/*
RosettaConstructionParse Parse a Transaction

TODO

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRosettaConstructionParseRequest
*/
func (a *RosettaAPIService) RosettaConstructionParse(ctx context.Context) ApiRosettaConstructionParseRequest {
	return ApiRosettaConstructionParseRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return RosettaConstructionParseResponse
func (a *RosettaAPIService) RosettaConstructionParseExecute(r ApiRosettaConstructionParseRequest) (*models.RosettaConstructionParseResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.RosettaConstructionParseResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RosettaAPIService.RosettaConstructionParse")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/rosetta/v1/construction/parse"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.rosettaConstructionParseRequest == nil {
		return localVarReturnValue, nil, utils.ReportError("rosettaConstructionParseRequest is required and must be specified")
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
	localVarPostBody = r.rosettaConstructionParseRequest
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v models.RosettaError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
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

type ApiRosettaConstructionPayloadsRequest struct {
	ctx                                context.Context
	ApiService                         *RosettaAPIService
	rosettaConstructionPayloadsRequest *models.RosettaConstructionPayloadsRequest
}

func (r ApiRosettaConstructionPayloadsRequest) RosettaConstructionPayloadsRequest(rosettaConstructionPayloadsRequest models.RosettaConstructionPayloadsRequest) ApiRosettaConstructionPayloadsRequest {
	r.rosettaConstructionPayloadsRequest = &rosettaConstructionPayloadsRequest
	return r
}

func (r ApiRosettaConstructionPayloadsRequest) Execute() (*models.RosettaConstructionPayloadResponse, *http.Response, error) {
	return r.ApiService.RosettaConstructionPayloadsExecute(r)
}

/*
RosettaConstructionPayloads Generate an Unsigned Transaction and Signing Payloads

Generate an unsigned transaction from operations and metadata. The examples below are illustrative only. You'll need to use your wallet to generate actual values to use them in the request payload.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRosettaConstructionPayloadsRequest
*/
func (a *RosettaAPIService) RosettaConstructionPayloads(ctx context.Context) ApiRosettaConstructionPayloadsRequest {
	return ApiRosettaConstructionPayloadsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return RosettaConstructionPayloadResponse
func (a *RosettaAPIService) RosettaConstructionPayloadsExecute(r ApiRosettaConstructionPayloadsRequest) (*models.RosettaConstructionPayloadResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.RosettaConstructionPayloadResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RosettaAPIService.RosettaConstructionPayloads")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/rosetta/v1/construction/payloads"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.rosettaConstructionPayloadsRequest == nil {
		return localVarReturnValue, nil, utils.ReportError("rosettaConstructionPayloadsRequest is required and must be specified")
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
	localVarPostBody = r.rosettaConstructionPayloadsRequest
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v models.RosettaError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
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

type ApiRosettaConstructionPreprocessRequest struct {
	ctx                                  context.Context
	ApiService                           *RosettaAPIService
	rosettaConstructionPreprocessRequest *models.RosettaConstructionPreprocessRequest
}

func (r ApiRosettaConstructionPreprocessRequest) RosettaConstructionPreprocessRequest(rosettaConstructionPreprocessRequest models.RosettaConstructionPreprocessRequest) ApiRosettaConstructionPreprocessRequest {
	r.rosettaConstructionPreprocessRequest = &rosettaConstructionPreprocessRequest
	return r
}

func (r ApiRosettaConstructionPreprocessRequest) Execute() (*models.RosettaConstructionPreprocessResponse, *http.Response, error) {
	return r.ApiService.RosettaConstructionPreprocessExecute(r)
}

/*
RosettaConstructionPreprocess Create a Request to Fetch Metadata

TODO

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRosettaConstructionPreprocessRequest
*/
func (a *RosettaAPIService) RosettaConstructionPreprocess(ctx context.Context) ApiRosettaConstructionPreprocessRequest {
	return ApiRosettaConstructionPreprocessRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return RosettaConstructionPreprocessResponse
func (a *RosettaAPIService) RosettaConstructionPreprocessExecute(r ApiRosettaConstructionPreprocessRequest) (*models.RosettaConstructionPreprocessResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.RosettaConstructionPreprocessResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RosettaAPIService.RosettaConstructionPreprocess")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/rosetta/v1/construction/preprocess"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.rosettaConstructionPreprocessRequest == nil {
		return localVarReturnValue, nil, utils.ReportError("rosettaConstructionPreprocessRequest is required and must be specified")
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
	localVarPostBody = r.rosettaConstructionPreprocessRequest
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v models.RosettaError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
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

type ApiRosettaConstructionSubmitRequest struct {
	ctx                              context.Context
	ApiService                       *RosettaAPIService
	rosettaConstructionSubmitRequest *models.RosettaConstructionSubmitRequest
}

func (r ApiRosettaConstructionSubmitRequest) RosettaConstructionSubmitRequest(rosettaConstructionSubmitRequest models.RosettaConstructionSubmitRequest) ApiRosettaConstructionSubmitRequest {
	r.rosettaConstructionSubmitRequest = &rosettaConstructionSubmitRequest
	return r
}

func (r ApiRosettaConstructionSubmitRequest) Execute() (*models.RosettaConstructionSubmitResponse, *http.Response, error) {
	return r.ApiService.RosettaConstructionSubmitExecute(r)
}

/*
RosettaConstructionSubmit Submit a Signed Transaction

Submit a pre-signed transaction to the node. The examples below are illustrative only. You'll need to use your wallet to generate actual values to use them in the request payload.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRosettaConstructionSubmitRequest
*/
func (a *RosettaAPIService) RosettaConstructionSubmit(ctx context.Context) ApiRosettaConstructionSubmitRequest {
	return ApiRosettaConstructionSubmitRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return RosettaConstructionSubmitResponse
func (a *RosettaAPIService) RosettaConstructionSubmitExecute(r ApiRosettaConstructionSubmitRequest) (*models.RosettaConstructionSubmitResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.RosettaConstructionSubmitResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RosettaAPIService.RosettaConstructionSubmit")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/rosetta/v1/construction/submit"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.rosettaConstructionSubmitRequest == nil {
		return localVarReturnValue, nil, utils.ReportError("rosettaConstructionSubmitRequest is required and must be specified")
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
	localVarPostBody = r.rosettaConstructionSubmitRequest
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v models.RosettaError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
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

type ApiRosettaMempoolRequest struct {
	ctx                   context.Context
	ApiService            *RosettaAPIService
	rosettaMempoolRequest *models.RosettaMempoolRequest
}

func (r ApiRosettaMempoolRequest) RosettaMempoolRequest(rosettaMempoolRequest models.RosettaMempoolRequest) ApiRosettaMempoolRequest {
	r.rosettaMempoolRequest = &rosettaMempoolRequest
	return r
}

func (r ApiRosettaMempoolRequest) Execute() (*models.RosettaMempoolResponse, *http.Response, error) {
	return r.ApiService.RosettaMempoolExecute(r)
}

/*
RosettaMempool Get All Mempool Transactions

Retrieves a list of transactions currently in the mempool for a given network.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRosettaMempoolRequest
*/
func (a *RosettaAPIService) RosettaMempool(ctx context.Context) ApiRosettaMempoolRequest {
	return ApiRosettaMempoolRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return RosettaMempoolResponse
func (a *RosettaAPIService) RosettaMempoolExecute(r ApiRosettaMempoolRequest) (*models.RosettaMempoolResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.RosettaMempoolResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RosettaAPIService.RosettaMempool")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/rosetta/v1/mempool"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.rosettaMempoolRequest == nil {
		return localVarReturnValue, nil, utils.ReportError("rosettaMempoolRequest is required and must be specified")
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
	localVarPostBody = r.rosettaMempoolRequest
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v models.RosettaError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
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

type ApiRosettaMempoolTransactionRequest struct {
	ctx                              context.Context
	ApiService                       *RosettaAPIService
	rosettaMempoolTransactionRequest *models.RosettaMempoolTransactionRequest
}

func (r ApiRosettaMempoolTransactionRequest) RosettaMempoolTransactionRequest(rosettaMempoolTransactionRequest models.RosettaMempoolTransactionRequest) ApiRosettaMempoolTransactionRequest {
	r.rosettaMempoolTransactionRequest = &rosettaMempoolTransactionRequest
	return r
}

func (r ApiRosettaMempoolTransactionRequest) Execute() (*models.RosettaMempoolTransactionResponse, *http.Response, error) {
	return r.ApiService.RosettaMempoolTransactionExecute(r)
}

/*
RosettaMempoolTransaction Get a Mempool Transaction

Retrieves transaction details from the mempool for a given transaction id from a given network.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRosettaMempoolTransactionRequest
*/
func (a *RosettaAPIService) RosettaMempoolTransaction(ctx context.Context) ApiRosettaMempoolTransactionRequest {
	return ApiRosettaMempoolTransactionRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return RosettaMempoolTransactionResponse
func (a *RosettaAPIService) RosettaMempoolTransactionExecute(r ApiRosettaMempoolTransactionRequest) (*models.RosettaMempoolTransactionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.RosettaMempoolTransactionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RosettaAPIService.RosettaMempoolTransaction")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/rosetta/v1/mempool/transaction"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.rosettaMempoolTransactionRequest == nil {
		return localVarReturnValue, nil, utils.ReportError("rosettaMempoolTransactionRequest is required and must be specified")
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
	localVarPostBody = r.rosettaMempoolTransactionRequest
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v models.RosettaError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
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

type ApiRosettaNetworkListRequest struct {
	ctx        context.Context
	ApiService *RosettaAPIService
}

func (r ApiRosettaNetworkListRequest) Execute() (*models.RosettaNetworkListResponse, *http.Response, error) {
	return r.ApiService.RosettaNetworkListExecute(r)
}

/*
RosettaNetworkList Get List of Available Networks

Retrieves a list of NetworkIdentifiers that the Rosetta server supports.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRosettaNetworkListRequest
*/
func (a *RosettaAPIService) RosettaNetworkList(ctx context.Context) ApiRosettaNetworkListRequest {
	return ApiRosettaNetworkListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return RosettaNetworkListResponse
func (a *RosettaAPIService) RosettaNetworkListExecute(r ApiRosettaNetworkListRequest) (*models.RosettaNetworkListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.RosettaNetworkListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RosettaAPIService.RosettaNetworkList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/rosetta/v1/network/list"

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
		if localVarHTTPResponse.StatusCode == 400 {
			var v models.RosettaError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
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

type ApiRosettaNetworkOptionsRequest struct {
	ctx                   context.Context
	ApiService            *RosettaAPIService
	rosettaOptionsRequest *models.RosettaOptionsRequest
}

func (r ApiRosettaNetworkOptionsRequest) RosettaOptionsRequest(rosettaOptionsRequest models.RosettaOptionsRequest) ApiRosettaNetworkOptionsRequest {
	r.rosettaOptionsRequest = &rosettaOptionsRequest
	return r
}

func (r ApiRosettaNetworkOptionsRequest) Execute() (*models.RosettaNetworkOptionsResponse, *http.Response, error) {
	return r.ApiService.RosettaNetworkOptionsExecute(r)
}

/*
RosettaNetworkOptions Get Network Options

Retrieves the version information and allowed network-specific types for a NetworkIdentifier.
Any NetworkIdentifier returned by /network/list should be accessible here.
Because options are retrievable in the context of a NetworkIdentifier, it is possible to define unique options for each network.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRosettaNetworkOptionsRequest
*/
func (a *RosettaAPIService) RosettaNetworkOptions(ctx context.Context) ApiRosettaNetworkOptionsRequest {
	return ApiRosettaNetworkOptionsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return RosettaNetworkOptionsResponse
func (a *RosettaAPIService) RosettaNetworkOptionsExecute(r ApiRosettaNetworkOptionsRequest) (*models.RosettaNetworkOptionsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.RosettaNetworkOptionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RosettaAPIService.RosettaNetworkOptions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/rosetta/v1/network/options"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.rosettaOptionsRequest == nil {
		return localVarReturnValue, nil, utils.ReportError("rosettaOptionsRequest is required and must be specified")
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
	localVarPostBody = r.rosettaOptionsRequest
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v models.RosettaError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
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

type ApiRosettaNetworkStatusRequest struct {
	ctx                  context.Context
	ApiService           *RosettaAPIService
	rosettaStatusRequest *models.RosettaStatusRequest
}

func (r ApiRosettaNetworkStatusRequest) RosettaStatusRequest(rosettaStatusRequest models.RosettaStatusRequest) ApiRosettaNetworkStatusRequest {
	r.rosettaStatusRequest = &rosettaStatusRequest
	return r
}

func (r ApiRosettaNetworkStatusRequest) Execute() (*models.RosettaNetworkStatusResponse, *http.Response, error) {
	return r.ApiService.RosettaNetworkStatusExecute(r)
}

/*
RosettaNetworkStatus Get Network Status

Retrieves the current status of the network requested.
Any NetworkIdentifier returned by /network/list should be accessible here.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRosettaNetworkStatusRequest
*/
func (a *RosettaAPIService) RosettaNetworkStatus(ctx context.Context) ApiRosettaNetworkStatusRequest {
	return ApiRosettaNetworkStatusRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return RosettaNetworkStatusResponse
func (a *RosettaAPIService) RosettaNetworkStatusExecute(r ApiRosettaNetworkStatusRequest) (*models.RosettaNetworkStatusResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.RosettaNetworkStatusResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RosettaAPIService.RosettaNetworkStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/rosetta/v1/network/status"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.rosettaStatusRequest == nil {
		return localVarReturnValue, nil, utils.ReportError("rosettaStatusRequest is required and must be specified")
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
	localVarPostBody = r.rosettaStatusRequest
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v models.RosettaError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
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
