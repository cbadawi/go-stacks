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

// BlocksAPIService BlocksAPI service
type BlocksAPIService service

type ApiGetAverageBlockTimesRequest struct {
	ctx        context.Context
	ApiService *BlocksAPIService
}

func (r ApiGetAverageBlockTimesRequest) Execute() (*models.AverageBlockTimesResponse, *http.Response, error) {
	return r.ApiService.GetAverageBlockTimesExecute(r)
}

/*
GetAverageBlockTimes Get average block times

Retrieves average block times (in seconds)


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetAverageBlockTimesRequest
*/
func (a *BlocksAPIService) GetAverageBlockTimes(ctx context.Context) ApiGetAverageBlockTimesRequest {
	return ApiGetAverageBlockTimesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return AverageBlockTimesResponse
func (a *BlocksAPIService) GetAverageBlockTimesExecute(r ApiGetAverageBlockTimesRequest) (*models.AverageBlockTimesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.AverageBlockTimesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BlocksAPIService.GetAverageBlockTimes")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/extended/v2/blocks/average-times"

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

type ApiGetBlockRequest struct {
	ctx          context.Context
	ApiService   *BlocksAPIService
	heightOrHash models.GetBurnBlockHeightOrHashParameter
}

func (r ApiGetBlockRequest) Execute() (*models.NakamotoBlock, *http.Response, error) {
	return r.ApiService.GetBlockExecute(r)
}

/*
GetBlock Get block

Retrieves a single block


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param heightOrHash filter by block height, hash, index block hash or the constant `latest` to filter for the most recent block
 @return ApiGetBlockRequest
*/
func (a *BlocksAPIService) GetBlock(ctx context.Context, heightOrHash models.GetBurnBlockHeightOrHashParameter) ApiGetBlockRequest {
	return ApiGetBlockRequest{
		ApiService:   a,
		ctx:          ctx,
		heightOrHash: heightOrHash,
	}
}

func getHeightOrHash(param models.GetBurnBlockHeightOrHashParameter) interface{} {
	if param.Int32 != nil {
		return *param.Int32
	}
	if param.String != nil {
		return *param.String
	}
	return nil
}

// Execute executes the request
//  @return NakamotoBlock
func (a *BlocksAPIService) GetBlockExecute(r ApiGetBlockRequest) (*models.NakamotoBlock, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.NakamotoBlock
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BlocksAPIService.GetBlock")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/extended/v2/blocks/{height_or_hash}"
	heightOrHash := getHeightOrHash(r.heightOrHash)
	localVarPath = strings.Replace(localVarPath, "{"+"height_or_hash"+"}", url.PathEscape(parameterValueToString(heightOrHash, "heightOrHash")), -1)

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

type ApiGetBlockByBurnBlockHashRequest struct {
	ctx           context.Context
	ApiService    *BlocksAPIService
	burnBlockHash string
}

func (r ApiGetBlockByBurnBlockHashRequest) Execute() (*models.Block, *http.Response, error) {
	return r.ApiService.GetBlockByBurnBlockHashExecute(r)
}

/*
GetBlockByBurnBlockHash Get block by burnchain block hash

**NOTE:** This endpoint is deprecated in favor of [Get blocks](/api/get-blocks).

Retrieves block details of a specific block for a given burnchain block hash


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param burnBlockHash Hash of the burnchain block
 @return ApiGetBlockByBurnBlockHashRequest

Deprecated
*/
func (a *BlocksAPIService) GetBlockByBurnBlockHash(ctx context.Context, burnBlockHash string) ApiGetBlockByBurnBlockHashRequest {
	return ApiGetBlockByBurnBlockHashRequest{
		ApiService:    a,
		ctx:           ctx,
		burnBlockHash: burnBlockHash,
	}
}

// Execute executes the request
//  @return Block
// Deprecated
func (a *BlocksAPIService) GetBlockByBurnBlockHashExecute(r ApiGetBlockByBurnBlockHashRequest) (*models.Block, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.Block
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BlocksAPIService.GetBlockByBurnBlockHash")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/extended/v1/block/by_burn_block_hash/{burn_block_hash}"
	localVarPath = strings.Replace(localVarPath, "{"+"burn_block_hash"+"}", url.PathEscape(parameterValueToString(r.burnBlockHash, "burnBlockHash")), -1)

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

type ApiGetBlockByHashRequest struct {
	ctx        context.Context
	ApiService *BlocksAPIService
	hash       string
}

func (r ApiGetBlockByHashRequest) Execute() (*models.Block, *http.Response, error) {
	return r.ApiService.GetBlockByHashExecute(r)
}

/*
GetBlockByHash Get block by hash

**NOTE:** This endpoint is deprecated in favor of [Get block](/api/get-block).

Retrieves block details of a specific block for a given chain height. You can use the hash from your latest block ('get_block_list' API) to get your block details.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param hash Hash of the block
 @return ApiGetBlockByHashRequest

Deprecated
*/
func (a *BlocksAPIService) GetBlockByHash(ctx context.Context, hash string) ApiGetBlockByHashRequest {
	return ApiGetBlockByHashRequest{
		ApiService: a,
		ctx:        ctx,
		hash:       hash,
	}
}

// Execute executes the request
//  @return Block
// Deprecated
func (a *BlocksAPIService) GetBlockByHashExecute(r ApiGetBlockByHashRequest) (*models.Block, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.Block
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BlocksAPIService.GetBlockByHash")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/extended/v1/block/{hash}"
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

type ApiGetBlockByHeightRequest struct {
	ctx        context.Context
	ApiService *BlocksAPIService
	height     float32
}

func (r ApiGetBlockByHeightRequest) Execute() (*models.Block, *http.Response, error) {
	return r.ApiService.GetBlockByHeightExecute(r)
}

/*
GetBlockByHeight Get block by height

**NOTE:** This endpoint is deprecated in favor of [Get block](/api/get-block).

Retrieves block details of a specific block at a given block height


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param height Height of the block
 @return ApiGetBlockByHeightRequest

Deprecated
*/
func (a *BlocksAPIService) GetBlockByHeight(ctx context.Context, height float32) ApiGetBlockByHeightRequest {
	return ApiGetBlockByHeightRequest{
		ApiService: a,
		ctx:        ctx,
		height:     height,
	}
}

// Execute executes the request
//  @return Block
// Deprecated
func (a *BlocksAPIService) GetBlockByHeightExecute(r ApiGetBlockByHeightRequest) (*models.Block, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.Block
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BlocksAPIService.GetBlockByHeight")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/extended/v1/block/by_height/{height}"
	localVarPath = strings.Replace(localVarPath, "{"+"height"+"}", url.PathEscape(parameterValueToString(r.height, "height")), -1)

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

type ApiGetBlockListRequest struct {
	ctx        context.Context
	ApiService *BlocksAPIService
	limit      *int32
	offset     *int32
}

// max number of blocks to fetch
func (r ApiGetBlockListRequest) Limit(limit int32) ApiGetBlockListRequest {
	r.limit = &limit
	return r
}

// index of first block to fetch
func (r ApiGetBlockListRequest) Offset(offset int32) ApiGetBlockListRequest {
	r.offset = &offset
	return r
}

func (r ApiGetBlockListRequest) Execute() (*models.BlockListResponse, *http.Response, error) {
	return r.ApiService.GetBlockListExecute(r)
}

/*
GetBlockList Get recent blocks

**NOTE:** This endpoint is deprecated in favor of [Get blocks](/api/get-blocks).

Retrieves a list of recently mined blocks

If you need to actively monitor new blocks, we highly recommend subscribing to [WebSockets or Socket.io](https://github.com/hirosystems/stacks-blockchain-api/tree/master/client) for real-time updates.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetBlockListRequest

Deprecated
*/
func (a *BlocksAPIService) GetBlockList(ctx context.Context) ApiGetBlockListRequest {
	return ApiGetBlockListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return BlockListResponse
// Deprecated
func (a *BlocksAPIService) GetBlockListExecute(r ApiGetBlockListRequest) (*models.BlockListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.BlockListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BlocksAPIService.GetBlockList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/extended/v1/block"

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

type ApiGetBlocksRequest struct {
	ctx        context.Context
	ApiService *BlocksAPIService
	limit      *int32
	offset     *int32
}

// max number of blocks to fetch
func (r ApiGetBlocksRequest) Limit(limit int32) ApiGetBlocksRequest {
	r.limit = &limit
	return r
}

// index of first burn block to fetch
func (r ApiGetBlocksRequest) Offset(offset int32) ApiGetBlocksRequest {
	r.offset = &offset
	return r
}

func (r ApiGetBlocksRequest) Execute() (*models.NakamotoBlockListResponse, *http.Response, error) {
	return r.ApiService.GetBlocksExecute(r)
}

/*
GetBlocks Get blocks

Retrieves a list of recently mined blocks


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetBlocksRequest
*/
func (a *BlocksAPIService) GetBlocks(ctx context.Context) ApiGetBlocksRequest {
	return ApiGetBlocksRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return NakamotoBlockListResponse
func (a *BlocksAPIService) GetBlocksExecute(r ApiGetBlocksRequest) (*models.NakamotoBlockListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.NakamotoBlockListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BlocksAPIService.GetBlocks")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/extended/v2/blocks"

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

type ApiGetBlocksByBurnBlockRequest struct {
	ctx          context.Context
	ApiService   *BlocksAPIService
	heightOrHash models.GetBurnBlockHeightOrHashParameter
	limit        *int32
	offset       *int32
}

// max number of blocks to fetch
func (r ApiGetBlocksByBurnBlockRequest) Limit(limit int32) ApiGetBlocksByBurnBlockRequest {
	r.limit = &limit
	return r
}

// index of first burn block to fetch
func (r ApiGetBlocksByBurnBlockRequest) Offset(offset int32) ApiGetBlocksByBurnBlockRequest {
	r.offset = &offset
	return r
}

func (r ApiGetBlocksByBurnBlockRequest) Execute() (*models.NakamotoBlockListResponse, *http.Response, error) {
	return r.ApiService.GetBlocksByBurnBlockExecute(r)
}

/*
GetBlocksByBurnBlock Get blocks by burn block

Retrieves a list of blocks confirmed by a specific burn block


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param heightOrHash filter by burn block height, hash, or the constant `latest` to filter for the most recent burn block
 @return ApiGetBlocksByBurnBlockRequest
*/
func (a *BlocksAPIService) GetBlocksByBurnBlock(ctx context.Context, heightOrHash models.GetBurnBlockHeightOrHashParameter) ApiGetBlocksByBurnBlockRequest {
	return ApiGetBlocksByBurnBlockRequest{
		ApiService:   a,
		ctx:          ctx,
		heightOrHash: heightOrHash,
	}
}

// Execute executes the request
//  @return NakamotoBlockListResponse
func (a *BlocksAPIService) GetBlocksByBurnBlockExecute(r ApiGetBlocksByBurnBlockRequest) (*models.NakamotoBlockListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.NakamotoBlockListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BlocksAPIService.GetBlocksByBurnBlock")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/extended/v2/burn-blocks/{height_or_hash}/blocks"
	localVarPath = strings.Replace(localVarPath, "{"+"height_or_hash"+"}", url.PathEscape(parameterValueToString(r.heightOrHash, "heightOrHash")), -1)

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
