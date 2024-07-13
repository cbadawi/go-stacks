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

// SearchAPIService SearchAPI service
type SearchAPIService service

type ApiSearchByIdRequest struct {
	ctx             context.Context
	ApiService      *SearchAPIService
	id              string
	includeMetadata *bool
}

// This includes the detailed data for purticular hash in the response
func (r ApiSearchByIdRequest) IncludeMetadata(includeMetadata bool) ApiSearchByIdRequest {
	r.includeMetadata = &includeMetadata
	return r
}

func (r ApiSearchByIdRequest) Execute() (*models.SearchResult, *http.Response, error) {
	return r.ApiService.SearchByIdExecute(r)
}

/*
SearchById Search

Search blocks, transactions, contracts, or accounts by hash/ID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The hex hash string for a block or transaction, account address, or contract address
 @return ApiSearchByIdRequest
*/
func (a *SearchAPIService) SearchById(ctx context.Context, id string) ApiSearchByIdRequest {
	return ApiSearchByIdRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//  @return SearchResult
func (a *SearchAPIService) SearchByIdExecute(r ApiSearchByIdRequest) (*models.SearchResult, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.SearchResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SearchAPIService.SearchById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/extended/v1/search/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.includeMetadata != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include_metadata", r.includeMetadata, "")
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
