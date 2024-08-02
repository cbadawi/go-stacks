package api

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/cbadawi/go-stacks/stacks/models"
)

// MicroblocksAPIService MicroblocksAPI service
type MicroblocksAPIService service

type ApiGetMicroblockByHashRequest struct {
	ctx        context.Context
	ApiService *MicroblocksAPIService
	hash       string
}

func (r ApiGetMicroblockByHashRequest) Execute() (*models.Microblock, *http.Response, error) {
	return r.ApiService.GetMicroblockByHashExecute(r)
}

/*
GetMicroblockByHash Get microblock

Retrieves a specific microblock by `hash`

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param hash Hash of the microblock
 @return ApiGetMicroblockByHashRequest
*/
func (a *MicroblocksAPIService) GetMicroblockByHash(ctx context.Context, hash string) ApiGetMicroblockByHashRequest {
	return ApiGetMicroblockByHashRequest{
		ApiService: a,
		ctx:        ctx,
		hash:       hash,
	}
}

// Execute executes the request
//  @return Microblock
func (a *MicroblocksAPIService) GetMicroblockByHashExecute(r ApiGetMicroblockByHashRequest) (*models.Microblock, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.Microblock
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MicroblocksAPIService.GetMicroblockByHash")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/extended/v1/microblock/{hash}"
	localVarPath = strings.Replace(localVarPath, "{"+"hash"+"}", url.PathEscape(parameterValueToString(r.hash, "hash")), -1)

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

type ApiGetMicroblockListRequest struct {
	ctx        context.Context
	ApiService *MicroblocksAPIService
	limit      *int32
	offset     *int32
}

// Max number of microblocks to fetch
func (r ApiGetMicroblockListRequest) Limit(limit int32) ApiGetMicroblockListRequest {
	r.limit = &limit
	return r
}

// Index of the first microblock to fetch
func (r ApiGetMicroblockListRequest) Offset(offset int32) ApiGetMicroblockListRequest {
	r.offset = &offset
	return r
}

func (r ApiGetMicroblockListRequest) Execute() (*models.MicroblockListResponse, *http.Response, error) {
	return r.ApiService.GetMicroblockListExecute(r)
}

/*
GetMicroblockList Get recent microblocks

Retrieves a list of microblocks.

If you need to actively monitor new microblocks, we highly recommend subscribing to [WebSockets or Socket.io](https://github.com/hirosystems/stacks-blockchain-api/tree/master/client) for real-time updates.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetMicroblockListRequest
*/
func (a *MicroblocksAPIService) GetMicroblockList(ctx context.Context) ApiGetMicroblockListRequest {
	return ApiGetMicroblockListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return MicroblockListResponse
func (a *MicroblocksAPIService) GetMicroblockListExecute(r ApiGetMicroblockListRequest) (*models.MicroblockListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.MicroblockListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MicroblocksAPIService.GetMicroblockList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/extended/v1/microblock"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	} else {
		var defaultValue int32 = 20
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

type ApiGetUnanchoredTxsRequest struct {
	ctx        context.Context
	ApiService *MicroblocksAPIService
}

func (r ApiGetUnanchoredTxsRequest) Execute() (*models.UnanchoredTransactionListResponse, *http.Response, error) {
	return r.ApiService.GetUnanchoredTxsExecute(r)
}

/*
GetUnanchoredTxs Get the list of current transactions that belong to unanchored microblocks

Retrieves transactions that have been streamed in microblocks but not yet accepted or rejected in an anchor block

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetUnanchoredTxsRequest
*/
func (a *MicroblocksAPIService) GetUnanchoredTxs(ctx context.Context) ApiGetUnanchoredTxsRequest {
	return ApiGetUnanchoredTxsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return UnanchoredTransactionListResponse
func (a *MicroblocksAPIService) GetUnanchoredTxsExecute(r ApiGetUnanchoredTxsRequest) (*models.UnanchoredTransactionListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.UnanchoredTransactionListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MicroblocksAPIService.GetUnanchoredTxs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/extended/v1/microblock/unanchored/txs"

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
