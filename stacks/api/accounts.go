package api

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/cbadawi/go-stacks/stacks/internal/models"
)

// AccountsAPIService AccountsAPI service
type AccountsAPIService service

type ApiGetAccountAssetsRequest struct {
	ctx        context.Context
	ApiService *AccountsAPIService
	principal  string
	limit      *int32
	offset     *int32
	unanchored *bool
	untilBlock *string
}

// max number of account assets to fetch
func (r ApiGetAccountAssetsRequest) Limit(limit int32) ApiGetAccountAssetsRequest {
	r.limit = &limit
	return r
}

// index of first account assets to fetch
func (r ApiGetAccountAssetsRequest) Offset(offset int32) ApiGetAccountAssetsRequest {
	r.offset = &offset
	return r
}

// Include transaction data from unanchored (i.e. unconfirmed) microblocks
func (r ApiGetAccountAssetsRequest) Unanchored(unanchored bool) ApiGetAccountAssetsRequest {
	r.unanchored = &unanchored
	return r
}

// returned data representing the state at that point in time, rather than the current block. Note - Use either of the query parameters but not both at a time.
func (r ApiGetAccountAssetsRequest) UntilBlock(untilBlock string) ApiGetAccountAssetsRequest {
	r.untilBlock = &untilBlock
	return r
}

func (r ApiGetAccountAssetsRequest) Execute() (*models.AddressAssetsListResponse, *http.Response, error) {
	return r.ApiService.GetAccountAssetsExecute(r)
}

/*
GetAccountAssets Get account assets

Retrieves a list of all assets events associated with an account or a Contract Identifier. This includes Transfers, Mints.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param principal Stacks address or a Contract identifier
 @return ApiGetAccountAssetsRequest
*/
func (a *AccountsAPIService) GetAccountAssets(ctx context.Context, principal string) ApiGetAccountAssetsRequest {
	return ApiGetAccountAssetsRequest{
		ApiService: a,
		ctx:        ctx,
		principal:  principal,
	}
}

// Execute executes the request
//  @return models.AddressAssetsListResponse
func (a *AccountsAPIService) GetAccountAssetsExecute(r ApiGetAccountAssetsRequest) (*models.AddressAssetsListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.AddressAssetsListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountsAPIService.GetAccountAssets")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/extended/v1/address/{principal}/assets"
	localVarPath = strings.Replace(localVarPath, "{"+"principal"+"}", url.PathEscape(parameterValueToString(r.principal, "principal")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	}
	if r.unanchored != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "unanchored", r.unanchored, "")
	} else {
		var defaultValue bool = false
		r.unanchored = &defaultValue
	}
	if r.untilBlock != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "until_block", r.untilBlock, "")
	}
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

type ApiGetAccountBalanceRequest struct {
	ctx        context.Context
	ApiService *AccountsAPIService
	principal  string
	unanchored *bool
	untilBlock *string
}

// Include transaction data from unanchored (i.e. unconfirmed) microblocks
func (r ApiGetAccountBalanceRequest) Unanchored(unanchored bool) ApiGetAccountBalanceRequest {
	r.unanchored = &unanchored
	return r
}

// returned data representing the state up until that point in time, rather than the current block.
func (r ApiGetAccountBalanceRequest) UntilBlock(untilBlock string) ApiGetAccountBalanceRequest {
	r.untilBlock = &untilBlock
	return r
}

func (r ApiGetAccountBalanceRequest) Execute() (*models.AddressBalanceResponse, *http.Response, error) {
	return r.ApiService.GetAccountBalanceExecute(r)
}

/*
GetAccountBalance Get account balances

Retrieves total account balance information for a given Address or Contract Identifier. This includes the balances of  STX Tokens, Fungible Tokens and Non-Fungible Tokens for the account.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param principal Stacks address or a Contract identifier
 @return ApiGetAccountBalanceRequest
*/
func (a *AccountsAPIService) GetAccountBalance(ctx context.Context, principal string) ApiGetAccountBalanceRequest {
	return ApiGetAccountBalanceRequest{
		ApiService: a,
		ctx:        ctx,
		principal:  principal,
	}
}

// Execute executes the request
//  @return AddressBalanceResponse
func (a *AccountsAPIService) GetAccountBalanceExecute(r ApiGetAccountBalanceRequest) (*models.AddressBalanceResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.AddressBalanceResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountsAPIService.GetAccountBalance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/extended/v1/address/{principal}/balances"
	localVarPath = strings.Replace(localVarPath, "{"+"principal"+"}", url.PathEscape(parameterValueToString(r.principal, "principal")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.unanchored != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "unanchored", r.unanchored, "")
	} else {
		var defaultValue bool = false
		r.unanchored = &defaultValue
	}
	if r.untilBlock != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "until_block", r.untilBlock, "")
	}
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

type ApiGetAccountInboundRequest struct {
	ctx        context.Context
	ApiService *AccountsAPIService
	principal  string
	limit      *int32
	offset     *int32
	height     *float32
	unanchored *bool
	untilBlock *string
}

// number of items to return
func (r ApiGetAccountInboundRequest) Limit(limit int32) ApiGetAccountInboundRequest {
	r.limit = &limit
	return r
}

// number of items to skip
func (r ApiGetAccountInboundRequest) Offset(offset int32) ApiGetAccountInboundRequest {
	r.offset = &offset
	return r
}

// Filter for transfers only at this given block height
func (r ApiGetAccountInboundRequest) Height(height float32) ApiGetAccountInboundRequest {
	r.height = &height
	return r
}

// Include transaction data from unanchored (i.e. unconfirmed) microblocks
func (r ApiGetAccountInboundRequest) Unanchored(unanchored bool) ApiGetAccountInboundRequest {
	r.unanchored = &unanchored
	return r
}

// returned data representing the state up until that point in time, rather than the current block. Note - Use either of the query parameters but not both at a time.
func (r ApiGetAccountInboundRequest) UntilBlock(untilBlock string) ApiGetAccountInboundRequest {
	r.untilBlock = &untilBlock
	return r
}

func (r ApiGetAccountInboundRequest) Execute() (*models.AddressStxInboundListResponse, *http.Response, error) {
	return r.ApiService.GetAccountInboundExecute(r)
}

/*
GetAccountInbound Get inbound STX transfers

Retrieves a list of STX transfers with memos to the given principal. This includes regular transfers from a stx-transfer transaction type,
and transfers from contract-call transactions a the `send-many-memo` bulk sending contract.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param principal Stacks address or a Contract identifier
 @return ApiGetAccountInboundRequest
*/
func (a *AccountsAPIService) GetAccountInbound(ctx context.Context, principal string) ApiGetAccountInboundRequest {
	return ApiGetAccountInboundRequest{
		ApiService: a,
		ctx:        ctx,
		principal:  principal,
	}
}

// Execute executes the request
//  @return AddressStxInboundListResponse
func (a *AccountsAPIService) GetAccountInboundExecute(r ApiGetAccountInboundRequest) (*models.AddressStxInboundListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.AddressStxInboundListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountsAPIService.GetAccountInbound")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/extended/v1/address/{principal}/stx_inbound"
	localVarPath = strings.Replace(localVarPath, "{"+"principal"+"}", url.PathEscape(parameterValueToString(r.principal, "principal")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	}
	if r.height != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "height", r.height, "")
	}
	if r.unanchored != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "unanchored", r.unanchored, "")
	} else {
		var defaultValue bool = false
		r.unanchored = &defaultValue
	}
	if r.untilBlock != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "until_block", r.untilBlock, "")
	}
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

type ApiGetAccountInfoRequest struct {
	ctx        context.Context
	ApiService *AccountsAPIService
	principal  string
	proof      *int32
	tip        *string
}

// Returns object without the proof field if set to 0
func (r ApiGetAccountInfoRequest) Proof(proof int32) ApiGetAccountInfoRequest {
	r.proof = &proof
	return r
}

// The Stacks chain tip to query from
func (r ApiGetAccountInfoRequest) Tip(tip string) ApiGetAccountInfoRequest {
	r.tip = &tip
	return r
}

func (r ApiGetAccountInfoRequest) Execute() (*models.AccountDataResponse, *http.Response, error) {
	return r.ApiService.GetAccountInfoExecute(r)
}

/*
GetAccountInfo Get account info

Retrieves the account data for a given Account or a Contract Identifier

Where balance is the hex encoding of a unsigned 128-bit integer (big-endian), nonce is an unsigned 64-bit integer, and the proofs are provided as hex strings.

For non-existent accounts, this does not return a 404 error, rather it returns an object with balance and nonce of 0.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param principal Stacks address or a Contract identifier
 @return ApiGetAccountInfoRequest
*/
func (a *AccountsAPIService) GetAccountInfo(ctx context.Context, principal string) ApiGetAccountInfoRequest {
	return ApiGetAccountInfoRequest{
		ApiService: a,
		ctx:        ctx,
		principal:  principal,
	}
}

// Execute executes the request
//  @return AccountDataResponse
func (a *AccountsAPIService) GetAccountInfoExecute(r ApiGetAccountInfoRequest) (*models.AccountDataResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.AccountDataResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountsAPIService.GetAccountInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/accounts/{principal}"
	localVarPath = strings.Replace(localVarPath, "{"+"principal"+"}", url.PathEscape(parameterValueToString(r.principal, "principal")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.proof != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "proof", r.proof, "")
	}
	if r.tip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "tip", r.tip, "")
	}
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

type ApiGetAccountNoncesRequest struct {
	ctx         context.Context
	ApiService  *AccountsAPIService
	principal   string
	blockHeight *float32
	blockHash   *string
}

// Optionally get the nonce at a given block height.
func (r ApiGetAccountNoncesRequest) BlockHeight(blockHeight float32) ApiGetAccountNoncesRequest {
	r.blockHeight = &blockHeight
	return r
}

// Optionally get the nonce at a given block hash. Note - Use either of the query parameters but not both at a time.
func (r ApiGetAccountNoncesRequest) BlockHash(blockHash string) ApiGetAccountNoncesRequest {
	r.blockHash = &blockHash
	return r
}

func (r ApiGetAccountNoncesRequest) Execute() (*models.AddressNonces, *http.Response, error) {
	return r.ApiService.GetAccountNoncesExecute(r)
}

/*
GetAccountNonces Get the latest nonce used by an account

Retrieves the latest nonce values used by an account by inspecting the mempool, microblock transactions, and anchored transactions.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param principal Stacks address
 @return ApiGetAccountNoncesRequest
*/
func (a *AccountsAPIService) GetAccountNonces(ctx context.Context, principal string) ApiGetAccountNoncesRequest {
	return ApiGetAccountNoncesRequest{
		ApiService: a,
		ctx:        ctx,
		principal:  principal,
	}
}

// Execute executes the request
//  @return AddressNonces
func (a *AccountsAPIService) GetAccountNoncesExecute(r ApiGetAccountNoncesRequest) (*models.AddressNonces, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.AddressNonces
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountsAPIService.GetAccountNonces")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/extended/v1/address/{principal}/nonces"
	localVarPath = strings.Replace(localVarPath, "{"+"principal"+"}", url.PathEscape(parameterValueToString(r.principal, "principal")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.blockHeight != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "block_height", r.blockHeight, "")
	}
	if r.blockHash != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "block_hash", r.blockHash, "")
	}
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

type ApiGetAccountStxBalanceRequest struct {
	ctx        context.Context
	ApiService *AccountsAPIService
	principal  string
	unanchored *bool
	untilBlock *string
}

// Include transaction data from unanchored (i.e. unconfirmed) microblocks.
func (r ApiGetAccountStxBalanceRequest) Unanchored(unanchored bool) ApiGetAccountStxBalanceRequest {
	r.unanchored = &unanchored
	return r
}

// returned data representing the state up until that point in time, rather than the current block. Note - Use either of the query parameters but not both at a time.
func (r ApiGetAccountStxBalanceRequest) UntilBlock(untilBlock string) ApiGetAccountStxBalanceRequest {
	r.untilBlock = &untilBlock
	return r
}

func (r ApiGetAccountStxBalanceRequest) Execute() (*models.AddressStxBalanceResponse, *http.Response, error) {
	return r.ApiService.GetAccountStxBalanceExecute(r)
}

/*
GetAccountStxBalance Get account STX balance

Retrieves STX token balance for a given Address or Contract Identifier.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param principal Stacks address or a Contract identifier.
 @return ApiGetAccountStxBalanceRequest
*/
func (a *AccountsAPIService) GetAccountStxBalance(ctx context.Context, principal string) ApiGetAccountStxBalanceRequest {
	return ApiGetAccountStxBalanceRequest{
		ApiService: a,
		ctx:        ctx,
		principal:  principal,
	}
}

// Execute executes the request
//  @return AddressStxBalanceResponse
func (a *AccountsAPIService) GetAccountStxBalanceExecute(r ApiGetAccountStxBalanceRequest) (*models.AddressStxBalanceResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.AddressStxBalanceResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountsAPIService.GetAccountStxBalance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/extended/v1/address/{principal}/stx"
	localVarPath = strings.Replace(localVarPath, "{"+"principal"+"}", url.PathEscape(parameterValueToString(r.principal, "principal")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.unanchored != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "unanchored", r.unanchored, "")
	} else {
		var defaultValue bool = false
		r.unanchored = &defaultValue
	}
	if r.untilBlock != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "until_block", r.untilBlock, "")
	}
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

type ApiGetAccountTransactionsRequest struct {
	ctx        context.Context
	ApiService *AccountsAPIService
	principal  string
	limit      *int32
	offset     *int32
	height     *float32
	unanchored *bool
	untilBlock *string
}

// max number of account transactions to fetch
func (r ApiGetAccountTransactionsRequest) Limit(limit int32) ApiGetAccountTransactionsRequest {
	r.limit = &limit
	return r
}

// index of first account transaction to fetch
func (r ApiGetAccountTransactionsRequest) Offset(offset int32) ApiGetAccountTransactionsRequest {
	r.offset = &offset
	return r
}

// Filter for transactions only at this given block height
func (r ApiGetAccountTransactionsRequest) Height(height float32) ApiGetAccountTransactionsRequest {
	r.height = &height
	return r
}

// Include transaction data from unanchored (i.e. unconfirmed) microblocks
func (r ApiGetAccountTransactionsRequest) Unanchored(unanchored bool) ApiGetAccountTransactionsRequest {
	r.unanchored = &unanchored
	return r
}

// returned data representing the state up until that point in time, rather than the current block. Note - Use either of the query parameters but not both at a time.
func (r ApiGetAccountTransactionsRequest) UntilBlock(untilBlock string) ApiGetAccountTransactionsRequest {
	r.untilBlock = &untilBlock
	return r
}

func (r ApiGetAccountTransactionsRequest) Execute() (*models.AddressTransactionsListResponse, *http.Response, error) {
	return r.ApiService.GetAccountTransactionsExecute(r)
}

/*
GetAccountTransactions Get account transactions

Retrieves a list of all Transactions for a given Address or Contract Identifier. More information on Transaction types can be found [here](https://docs.stacks.co/understand-stacks/transactions#types).

If you need to actively monitor new transactions for an address or contract id, we highly recommend subscribing to [WebSockets or Socket.io](https://github.com/hirosystems/stacks-blockchain-api/tree/master/client) for real-time updates.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param principal Stacks address or a Contract identifier
 @return ApiGetAccountTransactionsRequest

Deprecated
*/
func (a *AccountsAPIService) GetAccountTransactions(ctx context.Context, principal string) ApiGetAccountTransactionsRequest {
	return ApiGetAccountTransactionsRequest{
		ApiService: a,
		ctx:        ctx,
		principal:  principal,
	}
}

// Execute executes the request
//  @return AddressTransactionsListResponse
// Deprecated
func (a *AccountsAPIService) GetAccountTransactionsExecute(r ApiGetAccountTransactionsRequest) (*models.AddressTransactionsListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.AddressTransactionsListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountsAPIService.GetAccountTransactions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/extended/v1/address/{principal}/transactions"
	localVarPath = strings.Replace(localVarPath, "{"+"principal"+"}", url.PathEscape(parameterValueToString(r.principal, "principal")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	}
	if r.height != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "height", r.height, "")
	}
	if r.unanchored != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "unanchored", r.unanchored, "")
	} else {
		var defaultValue bool = false
		r.unanchored = &defaultValue
	}
	if r.untilBlock != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "until_block", r.untilBlock, "")
	}
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

type ApiGetAccountTransactionsWithTransfersRequest struct {
	ctx        context.Context
	ApiService *AccountsAPIService
	principal  string
	limit      *int32
	offset     *int32
	height     *float32
	unanchored *bool
	untilBlock *string
}

// max number of account transactions to fetch
func (r ApiGetAccountTransactionsWithTransfersRequest) Limit(limit int32) ApiGetAccountTransactionsWithTransfersRequest {
	r.limit = &limit
	return r
}

// index of first account transaction to fetch
func (r ApiGetAccountTransactionsWithTransfersRequest) Offset(offset int32) ApiGetAccountTransactionsWithTransfersRequest {
	r.offset = &offset
	return r
}

// Filter for transactions only at this given block height
func (r ApiGetAccountTransactionsWithTransfersRequest) Height(height float32) ApiGetAccountTransactionsWithTransfersRequest {
	r.height = &height
	return r
}

// Include transaction data from unanchored (i.e. unconfirmed) microblocks
func (r ApiGetAccountTransactionsWithTransfersRequest) Unanchored(unanchored bool) ApiGetAccountTransactionsWithTransfersRequest {
	r.unanchored = &unanchored
	return r
}

// returned data representing the state up until that point in time, rather than the current block.
func (r ApiGetAccountTransactionsWithTransfersRequest) UntilBlock(untilBlock string) ApiGetAccountTransactionsWithTransfersRequest {
	r.untilBlock = &untilBlock
	return r
}

func (r ApiGetAccountTransactionsWithTransfersRequest) Execute() (*models.AddressTransactionsWithTransfersListResponse, *http.Response, error) {
	return r.ApiService.GetAccountTransactionsWithTransfersExecute(r)
}

/*
GetAccountTransactionsWithTransfers Get account transactions including STX transfers for each transaction.

Retrieve all transactions for an account or contract identifier including STX transfers for each transaction.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param principal Stacks address or a Contract identifier
 @return ApiGetAccountTransactionsWithTransfersRequest

Deprecated
*/
func (a *AccountsAPIService) GetAccountTransactionsWithTransfers(ctx context.Context, principal string) ApiGetAccountTransactionsWithTransfersRequest {
	return ApiGetAccountTransactionsWithTransfersRequest{
		ApiService: a,
		ctx:        ctx,
		principal:  principal,
	}
}

// Execute executes the request
//  @return AddressTransactionsWithTransfersListResponse
// Deprecated
func (a *AccountsAPIService) GetAccountTransactionsWithTransfersExecute(r ApiGetAccountTransactionsWithTransfersRequest) (*models.AddressTransactionsWithTransfersListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.AddressTransactionsWithTransfersListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountsAPIService.GetAccountTransactionsWithTransfers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/extended/v1/address/{principal}/transactions_with_transfers"
	localVarPath = strings.Replace(localVarPath, "{"+"principal"+"}", url.PathEscape(parameterValueToString(r.principal, "principal")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	}
	if r.height != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "height", r.height, "")
	}
	if r.unanchored != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "unanchored", r.unanchored, "")
	} else {
		var defaultValue bool = false
		r.unanchored = &defaultValue
	}
	if r.untilBlock != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "until_block", r.untilBlock, "")
	}
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

type ApiGetSingleTransactionWithTransfersRequest struct {
	ctx        context.Context
	ApiService *AccountsAPIService
	principal  string
	txId       string
}

func (r ApiGetSingleTransactionWithTransfersRequest) Execute() (*models.AddressTransactionWithTransfers, *http.Response, error) {
	return r.ApiService.GetSingleTransactionWithTransfersExecute(r)
}

/*
GetSingleTransactionWithTransfers Get account transaction information for specific transaction

Retrieves transaction details for a given Transaction Id `tx_id`, for a given account or contract Identifier.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param principal Stacks address or a contract identifier
 @param txId Transaction id
 @return ApiGetSingleTransactionWithTransfersRequest

Deprecated
*/
func (a *AccountsAPIService) GetSingleTransactionWithTransfers(ctx context.Context, principal string, txId string) ApiGetSingleTransactionWithTransfersRequest {
	return ApiGetSingleTransactionWithTransfersRequest{
		ApiService: a,
		ctx:        ctx,
		principal:  principal,
		txId:       txId,
	}
}

// Execute executes the request
//  @return AddressTransactionWithTransfers
// Deprecated
func (a *AccountsAPIService) GetSingleTransactionWithTransfersExecute(r ApiGetSingleTransactionWithTransfersRequest) (*models.AddressTransactionWithTransfers, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.AddressTransactionWithTransfers
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountsAPIService.GetSingleTransactionWithTransfers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/extended/v1/address/{principal}/{tx_id}/with_transfers"
	localVarPath = strings.Replace(localVarPath, "{"+"principal"+"}", url.PathEscape(parameterValueToString(r.principal, "principal")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"tx_id"+"}", url.PathEscape(parameterValueToString(r.txId, "txId")), -1)

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
