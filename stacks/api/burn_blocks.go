package api

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"

	"github.com/cbadawi/go-stacks/stacks/internal/models"
)

// BurnBlocksAPIService BurnBlocksAPI service
type BurnBlocksAPIService service

type ApiGetBurnBlockRequest struct {
	ctx          context.Context
	ApiService   *BurnBlocksAPIService
	heightOrHash models.GetBurnBlockHeightOrHashParameter
}

type ApiGetBurnBlocksRequest struct {
	ctx        context.Context
	ApiService *BurnBlocksAPIService
	limit      *int32
	offset     *int32
}

// max number of burn blocks to fetch
func (r ApiGetBurnBlocksRequest) Limit(limit int32) ApiGetBurnBlocksRequest {
	r.limit = &limit
	return r
}

// index of first burn block to fetch
func (r ApiGetBurnBlocksRequest) Offset(offset int32) ApiGetBurnBlocksRequest {
	r.offset = &offset
	return r
}

func (r ApiGetBurnBlocksRequest) Execute() (*models.BurnBlockListResponse, *http.Response, error) {
	return r.ApiService.GetBurnBlocksExecute(r)
}

/*
GetBurnBlocks Get burn blocks

Retrieves a list of recent burn blocks


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetBurnBlocksRequest
*/
func (a *BurnBlocksAPIService) GetBurnBlocks(ctx context.Context) ApiGetBurnBlocksRequest {
	return ApiGetBurnBlocksRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return BurnBlockListResponse
func (a *BurnBlocksAPIService) GetBurnBlocksExecute(r ApiGetBurnBlocksRequest) (*models.BurnBlockListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.BurnBlockListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BurnBlocksAPIService.GetBurnBlocks")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/extended/v2/burn-blocks"

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
