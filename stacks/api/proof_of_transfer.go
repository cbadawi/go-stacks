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

// ProofOfTransferAPIService ProofOfTransferAPI service
type ProofOfTransferAPIService service

type ApiGetPoxCycleRequest struct {
	ctx         context.Context
	ApiService  *ProofOfTransferAPIService
	cycleNumber int32
}

func (r ApiGetPoxCycleRequest) Execute() (*models.PoxCycle, *http.Response, error) {
	return r.ApiService.GetPoxCycleExecute(r)
}

/*
GetPoxCycle Get PoX cycle

Retrieves details for a PoX cycle

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param cycleNumber PoX cycle number
 @return ApiGetPoxCycleRequest
*/
func (a *ProofOfTransferAPIService) GetPoxCycle(ctx context.Context, cycleNumber int32) ApiGetPoxCycleRequest {
	return ApiGetPoxCycleRequest{
		ApiService:  a,
		ctx:         ctx,
		cycleNumber: cycleNumber,
	}
}

// Execute executes the request
//  @return PoxCycle
func (a *ProofOfTransferAPIService) GetPoxCycleExecute(r ApiGetPoxCycleRequest) (*models.PoxCycle, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.PoxCycle
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProofOfTransferAPIService.GetPoxCycle")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/extended/v2/pox/cycles/{cycle_number}"
	localVarPath = strings.Replace(localVarPath, "{"+"cycle_number"+"}", url.PathEscape(parameterValueToString(r.cycleNumber, "cycleNumber")), -1)

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

type ApiGetPoxCycleSignerRequest struct {
	ctx         context.Context
	ApiService  *ProofOfTransferAPIService
	cycleNumber int32
	signerKey   string
}

func (r ApiGetPoxCycleSignerRequest) Execute() (*models.PoxSigner, *http.Response, error) {
	return r.ApiService.GetPoxCycleSignerExecute(r)
}

/*
GetPoxCycleSigner Get signer in PoX cycle

Retrieves details for a signer in a PoX cycle

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param cycleNumber PoX cycle number
 @param signerKey Signer key
 @return ApiGetPoxCycleSignerRequest
*/
func (a *ProofOfTransferAPIService) GetPoxCycleSigner(ctx context.Context, cycleNumber int32, signerKey string) ApiGetPoxCycleSignerRequest {
	return ApiGetPoxCycleSignerRequest{
		ApiService:  a,
		ctx:         ctx,
		cycleNumber: cycleNumber,
		signerKey:   signerKey,
	}
}

// Execute executes the request
//  @return PoxSigner
func (a *ProofOfTransferAPIService) GetPoxCycleSignerExecute(r ApiGetPoxCycleSignerRequest) (*models.PoxSigner, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.PoxSigner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProofOfTransferAPIService.GetPoxCycleSigner")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/extended/v2/pox/cycles/{cycle_number}/signers/{signer_key}"
	localVarPath = strings.Replace(localVarPath, "{"+"cycle_number"+"}", url.PathEscape(parameterValueToString(r.cycleNumber, "cycleNumber")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"signer_key"+"}", url.PathEscape(parameterValueToString(r.signerKey, "signerKey")), -1)

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

type ApiGetPoxCycleSignerStackersRequest struct {
	ctx         context.Context
	ApiService  *ProofOfTransferAPIService
	cycleNumber int32
	signerKey   string
}

func (r ApiGetPoxCycleSignerStackersRequest) Execute() (*models.PoxCycleSignerStackersListResponse, *http.Response, error) {
	return r.ApiService.GetPoxCycleSignerStackersExecute(r)
}

/*
GetPoxCycleSignerStackers Get stackers for signer in PoX cycle

Retrieves a list of stackers for a signer in a PoX cycle

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param cycleNumber PoX cycle number
 @param signerKey Signer key
 @return ApiGetPoxCycleSignerStackersRequest
*/
func (a *ProofOfTransferAPIService) GetPoxCycleSignerStackers(ctx context.Context, cycleNumber int32, signerKey string) ApiGetPoxCycleSignerStackersRequest {
	return ApiGetPoxCycleSignerStackersRequest{
		ApiService:  a,
		ctx:         ctx,
		cycleNumber: cycleNumber,
		signerKey:   signerKey,
	}
}

// Execute executes the request
//  @return PoxCycleSignerStackersListResponse
func (a *ProofOfTransferAPIService) GetPoxCycleSignerStackersExecute(r ApiGetPoxCycleSignerStackersRequest) (*models.PoxCycleSignerStackersListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.PoxCycleSignerStackersListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProofOfTransferAPIService.GetPoxCycleSignerStackers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/extended/v2/pox/cycles/{cycle_number}/signers/{signer_key}/stackers"
	localVarPath = strings.Replace(localVarPath, "{"+"cycle_number"+"}", url.PathEscape(parameterValueToString(r.cycleNumber, "cycleNumber")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"signer_key"+"}", url.PathEscape(parameterValueToString(r.signerKey, "signerKey")), -1)

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

type ApiGetPoxCycleSignersRequest struct {
	ctx         context.Context
	ApiService  *ProofOfTransferAPIService
	cycleNumber int32
}

func (r ApiGetPoxCycleSignersRequest) Execute() (*models.PoxCycleSignersListResponse, *http.Response, error) {
	return r.ApiService.GetPoxCycleSignersExecute(r)
}

/*
GetPoxCycleSigners Get signers in PoX cycle

Retrieves a list of signers in a PoX cycle

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param cycleNumber PoX cycle number
 @return ApiGetPoxCycleSignersRequest
*/
func (a *ProofOfTransferAPIService) GetPoxCycleSigners(ctx context.Context, cycleNumber int32) ApiGetPoxCycleSignersRequest {
	return ApiGetPoxCycleSignersRequest{
		ApiService:  a,
		ctx:         ctx,
		cycleNumber: cycleNumber,
	}
}

// Execute executes the request
//  @return PoxCycleSignersListResponse
func (a *ProofOfTransferAPIService) GetPoxCycleSignersExecute(r ApiGetPoxCycleSignersRequest) (*models.PoxCycleSignersListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.PoxCycleSignersListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProofOfTransferAPIService.GetPoxCycleSigners")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/extended/v2/pox/cycles/{cycle_number}/signers"
	localVarPath = strings.Replace(localVarPath, "{"+"cycle_number"+"}", url.PathEscape(parameterValueToString(r.cycleNumber, "cycleNumber")), -1)

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

type ApiGetPoxCyclesRequest struct {
	ctx        context.Context
	ApiService *ProofOfTransferAPIService
	limit      *int32
	offset     *int32
}

// max number of cycles to fetch
func (r ApiGetPoxCyclesRequest) Limit(limit int32) ApiGetPoxCyclesRequest {
	r.limit = &limit
	return r
}

// index of first cycle to fetch
func (r ApiGetPoxCyclesRequest) Offset(offset int32) ApiGetPoxCyclesRequest {
	r.offset = &offset
	return r
}

func (r ApiGetPoxCyclesRequest) Execute() (*models.PoxCycleListResponse, *http.Response, error) {
	return r.ApiService.GetPoxCyclesExecute(r)
}

/*
GetPoxCycles Get PoX cycles

Retrieves a list of PoX cycles

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetPoxCyclesRequest
*/
func (a *ProofOfTransferAPIService) GetPoxCycles(ctx context.Context) ApiGetPoxCyclesRequest {
	return ApiGetPoxCyclesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return PoxCycleListResponse
func (a *ProofOfTransferAPIService) GetPoxCyclesExecute(r ApiGetPoxCyclesRequest) (*models.PoxCycleListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.PoxCycleListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProofOfTransferAPIService.GetPoxCycles")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/extended/v2/pox/cycles"

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
