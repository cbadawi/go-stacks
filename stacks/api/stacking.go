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

// StackingAPIService StackingAPI service
type StackingAPIService service

type ApiGetPoolDelegationsRequest struct {
	ctx           context.Context
	ApiService    *StackingAPIService
	poolPrincipal string
	afterBlock    *int32
	unanchored    *bool
	limit         *int32
	offset        *int32
}

// If specified, only delegation events after the given block will be included
func (r ApiGetPoolDelegationsRequest) AfterBlock(afterBlock int32) ApiGetPoolDelegationsRequest {
	r.afterBlock = &afterBlock
	return r
}

// whether or not to include Stackers from unconfirmed transactions
func (r ApiGetPoolDelegationsRequest) Unanchored(unanchored bool) ApiGetPoolDelegationsRequest {
	r.unanchored = &unanchored
	return r
}

// number of items to return
func (r ApiGetPoolDelegationsRequest) Limit(limit int32) ApiGetPoolDelegationsRequest {
	r.limit = &limit
	return r
}

// number of items to skip
func (r ApiGetPoolDelegationsRequest) Offset(offset int32) ApiGetPoolDelegationsRequest {
	r.offset = &offset
	return r
}

func (r ApiGetPoolDelegationsRequest) Execute() (*models.PoolDelegationsResponse, *http.Response, error) {
	return r.ApiService.GetPoolDelegationsExecute(r)
}

/*
GetPoolDelegations Stacking pool members

Retrieves the list of stacking pool members for a given delegator principal.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param poolPrincipal Address principal of the stacking pool delegator
 @return ApiGetPoolDelegationsRequest
*/
func (a *StackingAPIService) GetPoolDelegations(ctx context.Context, poolPrincipal string) ApiGetPoolDelegationsRequest {
	return ApiGetPoolDelegationsRequest{
		ApiService:    a,
		ctx:           ctx,
		poolPrincipal: poolPrincipal,
	}
}

// Execute executes the request
//  @return models.PoolDelegationsResponse
func (a *StackingAPIService) GetPoolDelegationsExecute(r ApiGetPoolDelegationsRequest) (*models.PoolDelegationsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.PoolDelegationsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StackingAPIService.GetPoolDelegations")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/extended/beta/stacking/{pool_principal}/delegations"
	localVarPath = strings.Replace(localVarPath, "{"+"pool_principal"+"}", url.PathEscape(parameterValueToString(r.poolPrincipal, "poolPrincipal")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.afterBlock != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "after_block", r.afterBlock, "")
	}
	if r.unanchored != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "unanchored", r.unanchored, "")
	} else {
		var defaultValue bool = false
		r.unanchored = &defaultValue
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	} else {
		var defaultValue int32 = 100
		r.limit = &defaultValue
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	} else {
		var defaultValue int32 = 0
		r.offset = &defaultValue
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
