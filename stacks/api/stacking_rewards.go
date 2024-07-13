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

// StackingRewardsAPIService StackingRewardsAPI service
type StackingRewardsAPIService service

type ApiGetBurnchainRewardListRequest struct {
	ctx        context.Context
	ApiService *StackingRewardsAPIService
	limit      *int32
	offset     *int32
}

// max number of rewards to fetch
func (r ApiGetBurnchainRewardListRequest) Limit(limit int32) ApiGetBurnchainRewardListRequest {
	r.limit = &limit
	return r
}

// index of first rewards to fetch
func (r ApiGetBurnchainRewardListRequest) Offset(offset int32) ApiGetBurnchainRewardListRequest {
	r.offset = &offset
	return r
}

func (r ApiGetBurnchainRewardListRequest) Execute() (*models.BurnchainRewardListResponse, *http.Response, error) {
	return r.ApiService.GetBurnchainRewardListExecute(r)
}

/*
GetBurnchainRewardList Get recent burnchain reward recipients

Retrieves a list of recent burnchain (e.g. Bitcoin) reward recipients with the associated amounts and block info

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetBurnchainRewardListRequest
*/
func (a *StackingRewardsAPIService) GetBurnchainRewardList(ctx context.Context) ApiGetBurnchainRewardListRequest {
	return ApiGetBurnchainRewardListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return BurnchainRewardListResponse
func (a *StackingRewardsAPIService) GetBurnchainRewardListExecute(r ApiGetBurnchainRewardListRequest) (*models.BurnchainRewardListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.BurnchainRewardListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StackingRewardsAPIService.GetBurnchainRewardList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/extended/v1/burnchain/rewards"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	} else {
		var defaultValue int32 = 96
		r.limit = &defaultValue
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
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

type ApiGetBurnchainRewardListByAddressRequest struct {
	ctx        context.Context
	ApiService *StackingRewardsAPIService
	address    string
	limit      *int32
	offset     *int32
}

// max number of rewards to fetch
func (r ApiGetBurnchainRewardListByAddressRequest) Limit(limit int32) ApiGetBurnchainRewardListByAddressRequest {
	r.limit = &limit
	return r
}

// index of first rewards to fetch
func (r ApiGetBurnchainRewardListByAddressRequest) Offset(offset int32) ApiGetBurnchainRewardListByAddressRequest {
	r.offset = &offset
	return r
}

func (r ApiGetBurnchainRewardListByAddressRequest) Execute() (*models.BurnchainRewardListResponse, *http.Response, error) {
	return r.ApiService.GetBurnchainRewardListByAddressExecute(r)
}

/*
GetBurnchainRewardListByAddress Get recent burnchain reward for the given recipient

Retrieves a list of recent burnchain (e.g. Bitcoin) rewards for the given recipient with the associated amounts and block info

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param address Reward recipient address. Should either be in the native burnchain's format (e.g. B58 for Bitcoin), or if a STX principal address is provided it will be encoded as into the equivalent burnchain format
 @return ApiGetBurnchainRewardListByAddressRequest
*/
func (a *StackingRewardsAPIService) GetBurnchainRewardListByAddress(ctx context.Context, address string) ApiGetBurnchainRewardListByAddressRequest {
	return ApiGetBurnchainRewardListByAddressRequest{
		ApiService: a,
		ctx:        ctx,
		address:    address,
	}
}

// Execute executes the request
//  @return BurnchainRewardListResponse
func (a *StackingRewardsAPIService) GetBurnchainRewardListByAddressExecute(r ApiGetBurnchainRewardListByAddressRequest) (*models.BurnchainRewardListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.BurnchainRewardListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StackingRewardsAPIService.GetBurnchainRewardListByAddress")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/extended/v1/burnchain/rewards/{address}"
	localVarPath = strings.Replace(localVarPath, "{"+"address"+"}", url.PathEscape(parameterValueToString(r.address, "address")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
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

type ApiGetBurnchainRewardSlotHoldersRequest struct {
	ctx        context.Context
	ApiService *StackingRewardsAPIService
	limit      *int32
	offset     *int32
}

// max number of items to fetch
func (r ApiGetBurnchainRewardSlotHoldersRequest) Limit(limit int32) ApiGetBurnchainRewardSlotHoldersRequest {
	r.limit = &limit
	return r
}

// index of the first items to fetch
func (r ApiGetBurnchainRewardSlotHoldersRequest) Offset(offset int32) ApiGetBurnchainRewardSlotHoldersRequest {
	r.offset = &offset
	return r
}

func (r ApiGetBurnchainRewardSlotHoldersRequest) Execute() (*models.BurnchainRewardSlotHolderListResponse, *http.Response, error) {
	return r.ApiService.GetBurnchainRewardSlotHoldersExecute(r)
}

/*
GetBurnchainRewardSlotHolders Get recent reward slot holders

Retrieves a list of the Bitcoin addresses that would validly receive Proof-of-Transfer commitments.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetBurnchainRewardSlotHoldersRequest
*/
func (a *StackingRewardsAPIService) GetBurnchainRewardSlotHolders(ctx context.Context) ApiGetBurnchainRewardSlotHoldersRequest {
	return ApiGetBurnchainRewardSlotHoldersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return BurnchainRewardSlotHolderListResponse
func (a *StackingRewardsAPIService) GetBurnchainRewardSlotHoldersExecute(r ApiGetBurnchainRewardSlotHoldersRequest) (*models.BurnchainRewardSlotHolderListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.BurnchainRewardSlotHolderListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StackingRewardsAPIService.GetBurnchainRewardSlotHolders")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/extended/v1/burnchain/reward_slot_holders"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	} else {
		var defaultValue int32 = 96
		r.limit = &defaultValue
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
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

type ApiGetBurnchainRewardSlotHoldersByAddressRequest struct {
	ctx        context.Context
	ApiService *StackingRewardsAPIService
	address    string
	limit      *int32
	offset     *int32
}

// max number of items to fetch
func (r ApiGetBurnchainRewardSlotHoldersByAddressRequest) Limit(limit int32) ApiGetBurnchainRewardSlotHoldersByAddressRequest {
	r.limit = &limit
	return r
}

// index of the first items to fetch
func (r ApiGetBurnchainRewardSlotHoldersByAddressRequest) Offset(offset int32) ApiGetBurnchainRewardSlotHoldersByAddressRequest {
	r.offset = &offset
	return r
}

func (r ApiGetBurnchainRewardSlotHoldersByAddressRequest) Execute() (*models.BurnchainRewardSlotHolderListResponse, *http.Response, error) {
	return r.ApiService.GetBurnchainRewardSlotHoldersByAddressExecute(r)
}

/*
GetBurnchainRewardSlotHoldersByAddress Get recent reward slot holder entries for the given address

Retrieves a list of the Bitcoin addresses that would validly receive Proof-of-Transfer commitments for a given reward slot holder recipient address.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param address Reward slot holder recipient address. Should either be in the native burnchain's format (e.g. B58 for Bitcoin), or if a STX principal address is provided it will be encoded as into the equivalent burnchain format
 @return ApiGetBurnchainRewardSlotHoldersByAddressRequest
*/
func (a *StackingRewardsAPIService) GetBurnchainRewardSlotHoldersByAddress(ctx context.Context, address string) ApiGetBurnchainRewardSlotHoldersByAddressRequest {
	return ApiGetBurnchainRewardSlotHoldersByAddressRequest{
		ApiService: a,
		ctx:        ctx,
		address:    address,
	}
}

// Execute executes the request
//  @return BurnchainRewardSlotHolderListResponse
func (a *StackingRewardsAPIService) GetBurnchainRewardSlotHoldersByAddressExecute(r ApiGetBurnchainRewardSlotHoldersByAddressRequest) (*models.BurnchainRewardSlotHolderListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.BurnchainRewardSlotHolderListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StackingRewardsAPIService.GetBurnchainRewardSlotHoldersByAddress")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/extended/v1/burnchain/reward_slot_holders/{address}"
	localVarPath = strings.Replace(localVarPath, "{"+"address"+"}", url.PathEscape(parameterValueToString(r.address, "address")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
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

type ApiGetBurnchainRewardsTotalByAddressRequest struct {
	ctx        context.Context
	ApiService *StackingRewardsAPIService
	address    string
}

func (r ApiGetBurnchainRewardsTotalByAddressRequest) Execute() (*models.BurnchainRewardsTotal, *http.Response, error) {
	return r.ApiService.GetBurnchainRewardsTotalByAddressExecute(r)
}

/*
GetBurnchainRewardsTotalByAddress Get total burnchain rewards for the given recipient

Retrieves the total burnchain (e.g. Bitcoin) rewards for a given recipient `address`

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param address Reward recipient address. Should either be in the native burnchain's format (e.g. B58 for Bitcoin), or if a STX principal address is provided it will be encoded as into the equivalent burnchain format
 @return ApiGetBurnchainRewardsTotalByAddressRequest
*/
func (a *StackingRewardsAPIService) GetBurnchainRewardsTotalByAddress(ctx context.Context, address string) ApiGetBurnchainRewardsTotalByAddressRequest {
	return ApiGetBurnchainRewardsTotalByAddressRequest{
		ApiService: a,
		ctx:        ctx,
		address:    address,
	}
}

// Execute executes the request
//  @return BurnchainRewardsTotal
func (a *StackingRewardsAPIService) GetBurnchainRewardsTotalByAddressExecute(r ApiGetBurnchainRewardsTotalByAddressRequest) (*models.BurnchainRewardsTotal, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.BurnchainRewardsTotal
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StackingRewardsAPIService.GetBurnchainRewardsTotalByAddress")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/extended/v1/burnchain/rewards/{address}/total"
	localVarPath = strings.Replace(localVarPath, "{"+"address"+"}", url.PathEscape(parameterValueToString(r.address, "address")), -1)

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
