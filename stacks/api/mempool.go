package api

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"

	"github.com/cbadawi/go-stacks/stacks/internal/models"
)

// MempoolAPIService MempoolAPI service
type MempoolAPIService service

type ApiGetMempoolFeePrioritiesRequest struct {
	ctx        context.Context
	ApiService *MempoolAPIService
}

func (r ApiGetMempoolFeePrioritiesRequest) Execute() (*models.MempoolFeePriorities, *http.Response, error) {
	return r.ApiService.GetMempoolFeePrioritiesExecute(r)
}

/*
GetMempoolFeePriorities Get mempool transaction fee priorities

Returns estimated fee priorities (in micro-STX) for all transactions that are currently in the mempool. Also returns priorities separated by transaction type.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetMempoolFeePrioritiesRequest
*/
func (a *MempoolAPIService) GetMempoolFeePriorities(ctx context.Context) ApiGetMempoolFeePrioritiesRequest {
	return ApiGetMempoolFeePrioritiesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return MempoolFeePriorities
func (a *MempoolAPIService) GetMempoolFeePrioritiesExecute(r ApiGetMempoolFeePrioritiesRequest) (*models.MempoolFeePriorities, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.MempoolFeePriorities
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MempoolAPIService.GetMempoolFeePriorities")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/extended/v2/mempool/fees"

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
