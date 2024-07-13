package api

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"github.com/cbadawi/go-stacks/stacks/internal/models"
	"github.com/cbadawi/go-stacks/stacks/utils"
)

// SmartContractsAPIService SmartContractsAPI service
type SmartContractsAPIService service

type ApiCallReadOnlyFunctionRequest struct {
	ctx                  context.Context
	ApiService           *SmartContractsAPIService
	contractAddress      string
	contractName         string
	functionName         string
	readOnlyFunctionArgs *models.ReadOnlyFunctionArgs
	tip                  *string
}

// map of arguments and the simulated tx-sender where sender is either a Contract identifier or a normal Stacks address, and arguments is an array of hex serialized Clarity values.
func (r ApiCallReadOnlyFunctionRequest) ReadOnlyFunctionArgs(readOnlyFunctionArgs models.ReadOnlyFunctionArgs) ApiCallReadOnlyFunctionRequest {
	r.readOnlyFunctionArgs = &readOnlyFunctionArgs
	return r
}

// The Stacks chain tip to query from
func (r ApiCallReadOnlyFunctionRequest) Tip(tip string) ApiCallReadOnlyFunctionRequest {
	r.tip = &tip
	return r
}

func (r ApiCallReadOnlyFunctionRequest) Execute() (*models.ReadOnlyFunctionSuccessResponse, *http.Response, error) {
	return r.ApiService.CallReadOnlyFunctionExecute(r)
}

/*
CallReadOnlyFunction Call read-only function

Call a read-only public function on a given smart contract.

The smart contract and function are specified using the URL path. The arguments and the simulated tx-sender are supplied via the POST body in the following JSON format:


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param contractAddress Stacks address
 @param contractName Contract name
 @param functionName Function name
 @return ApiCallReadOnlyFunctionRequest
*/
func (a *SmartContractsAPIService) CallReadOnlyFunction(ctx context.Context, contractAddress string, contractName string, functionName string) ApiCallReadOnlyFunctionRequest {
	return ApiCallReadOnlyFunctionRequest{
		ApiService:      a,
		ctx:             ctx,
		contractAddress: contractAddress,
		contractName:    contractName,
		functionName:    functionName,
	}
}

// Execute executes the request
//  @return ReadOnlyFunctionSuccessResponse
func (a *SmartContractsAPIService) CallReadOnlyFunctionExecute(r ApiCallReadOnlyFunctionRequest) (*models.ReadOnlyFunctionSuccessResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.ReadOnlyFunctionSuccessResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SmartContractsAPIService.CallReadOnlyFunction")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/contracts/call-read/{contract_address}/{contract_name}/{function_name}"
	localVarPath = strings.Replace(localVarPath, "{"+"contract_address"+"}", url.PathEscape(parameterValueToString(r.contractAddress, "contractAddress")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"contract_name"+"}", url.PathEscape(parameterValueToString(r.contractName, "contractName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"function_name"+"}", url.PathEscape(parameterValueToString(r.functionName, "functionName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.readOnlyFunctionArgs == nil {
		return localVarReturnValue, nil, utils.ReportError("readOnlyFunctionArgs is required and must be specified", a.client.cfg.Logger)
	}

	if r.tip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "tip", r.tip, "")
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
	localVarPostBody = r.readOnlyFunctionArgs
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

type ApiGetContractByIdRequest struct {
	ctx        context.Context
	ApiService *SmartContractsAPIService
	contractId string
	unanchored *bool
}

// Include transaction data from unanchored (i.e. unconfirmed) microblocks
func (r ApiGetContractByIdRequest) Unanchored(unanchored bool) ApiGetContractByIdRequest {
	r.unanchored = &unanchored
	return r
}

func (r ApiGetContractByIdRequest) Execute() (*models.SmartContract, *http.Response, error) {
	return r.ApiService.GetContractByIdExecute(r)
}

/*
GetContractById Get contract info

Retrieves details of a contract with a given `contract_id`

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param contractId Contract identifier formatted as `<contract_address>.<contract_name>`
 @return ApiGetContractByIdRequest
*/
func (a *SmartContractsAPIService) GetContractById(ctx context.Context, contractId string) ApiGetContractByIdRequest {
	return ApiGetContractByIdRequest{
		ApiService: a,
		ctx:        ctx,
		contractId: contractId,
	}
}

// Execute executes the request
//  @return SmartContract
func (a *SmartContractsAPIService) GetContractByIdExecute(r ApiGetContractByIdRequest) (*models.SmartContract, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.SmartContract
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SmartContractsAPIService.GetContractById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/extended/v1/contract/{contract_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"contract_id"+"}", url.PathEscape(parameterValueToString(r.contractId, "contractId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.unanchored != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "unanchored", r.unanchored, "")
	} else {
		var defaultValue bool = false
		r.unanchored = &defaultValue
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

type ApiGetContractDataMapEntryRequest struct {
	ctx             context.Context
	ApiService      *SmartContractsAPIService
	contractAddress string
	contractName    string
	mapName         string
	key             *string
	proof           *int32
	tip             *string
}

// Hex string serialization of the lookup key (which should be a Clarity value)
func (r ApiGetContractDataMapEntryRequest) Key(key string) ApiGetContractDataMapEntryRequest {
	r.key = &key
	return r
}

// Returns object without the proof field when set to 0
func (r ApiGetContractDataMapEntryRequest) Proof(proof int32) ApiGetContractDataMapEntryRequest {
	r.proof = &proof
	return r
}

// The Stacks chain tip to query from
func (r ApiGetContractDataMapEntryRequest) Tip(tip string) ApiGetContractDataMapEntryRequest {
	r.tip = &tip
	return r
}

func (r ApiGetContractDataMapEntryRequest) Execute() (*models.MapEntryResponse, *http.Response, error) {
	return r.ApiService.GetContractDataMapEntryExecute(r)
}

/*
GetContractDataMapEntry Get specific data-map inside a contract

Attempt to fetch data from a contract data map. The contract is identified with Stacks Address `contract_address` and Contract Name `contract_address` in the URL path. The map is identified with [Map Name].

The key to lookup in the map is supplied via the POST body. This should be supplied as the hex string serialization of the key (which should be a Clarity value). Note, this is a JSON string atom.

In the response, `data` is the hex serialization of the map response. Note that map responses are Clarity option types, for non-existent values, this is a serialized none, and for all other responses, it is a serialized (some ...) object.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param contractAddress Stacks address
 @param contractName Contract name
 @param mapName Map name
 @return ApiGetContractDataMapEntryRequest
*/
func (a *SmartContractsAPIService) GetContractDataMapEntry(ctx context.Context, contractAddress string, contractName string, mapName string) ApiGetContractDataMapEntryRequest {
	return ApiGetContractDataMapEntryRequest{
		ApiService:      a,
		ctx:             ctx,
		contractAddress: contractAddress,
		contractName:    contractName,
		mapName:         mapName,
	}
}

// Execute executes the request
//  @return MapEntryResponse
func (a *SmartContractsAPIService) GetContractDataMapEntryExecute(r ApiGetContractDataMapEntryRequest) (*models.MapEntryResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.MapEntryResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SmartContractsAPIService.GetContractDataMapEntry")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/map_entry/{contract_address}/{contract_name}/{map_name}"
	localVarPath = strings.Replace(localVarPath, "{"+"contract_address"+"}", url.PathEscape(parameterValueToString(r.contractAddress, "contractAddress")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"contract_name"+"}", url.PathEscape(parameterValueToString(r.contractName, "contractName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"map_name"+"}", url.PathEscape(parameterValueToString(r.mapName, "mapName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.key == nil {
		return localVarReturnValue, nil, utils.ReportError("key is required and must be specified", a.client.cfg.Logger)
	}

	if r.proof != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "proof", r.proof, "")
	}
	if r.tip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "tip", r.tip, "")
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
	localVarPostBody = r.key
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

type ApiGetContractEventsByIdRequest struct {
	ctx        context.Context
	ApiService *SmartContractsAPIService
	contractId string
	limit      *int32
	offset     *int32
	unanchored *bool
}

// max number of contract events to fetch
func (r ApiGetContractEventsByIdRequest) Limit(limit int32) ApiGetContractEventsByIdRequest {
	r.limit = &limit
	return r
}

// index of first contract event to fetch
func (r ApiGetContractEventsByIdRequest) Offset(offset int32) ApiGetContractEventsByIdRequest {
	r.offset = &offset
	return r
}

// Include transaction data from unanchored (i.e. unconfirmed) microblocks
func (r ApiGetContractEventsByIdRequest) Unanchored(unanchored bool) ApiGetContractEventsByIdRequest {
	r.unanchored = &unanchored
	return r
}

func (r ApiGetContractEventsByIdRequest) Execute() (*models.TransactionEvent, *http.Response, error) {
	return r.ApiService.GetContractEventsByIdExecute(r)
}

/*
GetContractEventsById Get contract events

Retrieves a list of events that have been triggered by a given `contract_id`

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param contractId Contract identifier formatted as `<contract_address>.<contract_name>`
 @return ApiGetContractEventsByIdRequest
*/
func (a *SmartContractsAPIService) GetContractEventsById(ctx context.Context, contractId string) ApiGetContractEventsByIdRequest {
	return ApiGetContractEventsByIdRequest{
		ApiService: a,
		ctx:        ctx,
		contractId: contractId,
	}
}

// Execute executes the request
//  @return TransactionEvent
func (a *SmartContractsAPIService) GetContractEventsByIdExecute(r ApiGetContractEventsByIdRequest) (*models.TransactionEvent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.TransactionEvent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SmartContractsAPIService.GetContractEventsById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/extended/v1/contract/{contract_id}/events"
	localVarPath = strings.Replace(localVarPath, "{"+"contract_id"+"}", url.PathEscape(parameterValueToString(r.contractId, "contractId")), -1)

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

type ApiGetContractInterfaceRequest struct {
	ctx             context.Context
	ApiService      *SmartContractsAPIService
	contractAddress string
	contractName    string
	tip             *string
}

// The Stacks chain tip to query from
func (r ApiGetContractInterfaceRequest) Tip(tip string) ApiGetContractInterfaceRequest {
	r.tip = &tip
	return r
}

func (r ApiGetContractInterfaceRequest) Execute() (*models.ContractInterfaceResponse, *http.Response, error) {
	return r.ApiService.GetContractInterfaceExecute(r)
}

/*
GetContractInterface Get contract interface

Retrieves a contract interface with a given `contract_address` and `contract name`

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param contractAddress Stacks address
 @param contractName Contract name
 @return ApiGetContractInterfaceRequest
*/
func (a *SmartContractsAPIService) GetContractInterface(ctx context.Context, contractAddress string, contractName string) ApiGetContractInterfaceRequest {
	return ApiGetContractInterfaceRequest{
		ApiService:      a,
		ctx:             ctx,
		contractAddress: contractAddress,
		contractName:    contractName,
	}
}

// Execute executes the request
//  @return ContractInterfaceResponse
func (a *SmartContractsAPIService) GetContractInterfaceExecute(r ApiGetContractInterfaceRequest) (*models.ContractInterfaceResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.ContractInterfaceResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SmartContractsAPIService.GetContractInterface")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/contracts/interface/{contract_address}/{contract_name}"
	localVarPath = strings.Replace(localVarPath, "{"+"contract_address"+"}", url.PathEscape(parameterValueToString(r.contractAddress, "contractAddress")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"contract_name"+"}", url.PathEscape(parameterValueToString(r.contractName, "contractName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiGetContractSourceRequest struct {
	ctx             context.Context
	ApiService      *SmartContractsAPIService
	contractAddress string
	contractName    string
	proof           *int32
	tip             *string
}

// Returns object without the proof field if set to 0
func (r ApiGetContractSourceRequest) Proof(proof int32) ApiGetContractSourceRequest {
	r.proof = &proof
	return r
}

// The Stacks chain tip to query from
func (r ApiGetContractSourceRequest) Tip(tip string) ApiGetContractSourceRequest {
	r.tip = &tip
	return r
}

func (r ApiGetContractSourceRequest) Execute() (*models.ContractSourceResponse, *http.Response, error) {
	return r.ApiService.GetContractSourceExecute(r)
}

/*
GetContractSource Get contract source

Retrieves the Clarity source code of a given contract, along with the block height it was published in, and the MARF proof for the data

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param contractAddress Stacks address
 @param contractName Contract name
 @return ApiGetContractSourceRequest
*/
func (a *SmartContractsAPIService) GetContractSource(ctx context.Context, contractAddress string, contractName string) ApiGetContractSourceRequest {
	return ApiGetContractSourceRequest{
		ApiService:      a,
		ctx:             ctx,
		contractAddress: contractAddress,
		contractName:    contractName,
	}
}

// Execute executes the request
//  @return ContractSourceResponse
func (a *SmartContractsAPIService) GetContractSourceExecute(r ApiGetContractSourceRequest) (*models.ContractSourceResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.ContractSourceResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SmartContractsAPIService.GetContractSource")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/contracts/source/{contract_address}/{contract_name}"
	localVarPath = strings.Replace(localVarPath, "{"+"contract_address"+"}", url.PathEscape(parameterValueToString(r.contractAddress, "contractAddress")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"contract_name"+"}", url.PathEscape(parameterValueToString(r.contractName, "contractName")), -1)

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

type ApiGetContractsByTraitRequest struct {
	ctx        context.Context
	ApiService *SmartContractsAPIService
	traitAbi   *string
	limit      *int32
	offset     *int32
}

// JSON abi of the trait.
func (r ApiGetContractsByTraitRequest) TraitAbi(traitAbi string) ApiGetContractsByTraitRequest {
	r.traitAbi = &traitAbi
	return r
}

// max number of contracts fetch
func (r ApiGetContractsByTraitRequest) Limit(limit int32) ApiGetContractsByTraitRequest {
	r.limit = &limit
	return r
}

// index of first contract event to fetch
func (r ApiGetContractsByTraitRequest) Offset(offset int32) ApiGetContractsByTraitRequest {
	r.offset = &offset
	return r
}

func (r ApiGetContractsByTraitRequest) Execute() (*models.ContractListResponse, *http.Response, error) {
	return r.ApiService.GetContractsByTraitExecute(r)
}

/*
GetContractsByTrait Get contracts by trait

Retrieves a list of contracts based on the following traits listed in JSON format -  functions, variables, maps, fungible tokens and non-fungible tokens

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetContractsByTraitRequest
*/
func (a *SmartContractsAPIService) GetContractsByTrait(ctx context.Context) ApiGetContractsByTraitRequest {
	return ApiGetContractsByTraitRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return ContractListResponse
func (a *SmartContractsAPIService) GetContractsByTraitExecute(r ApiGetContractsByTraitRequest) (*models.ContractListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.ContractListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SmartContractsAPIService.GetContractsByTrait")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/extended/v1/contract/by_trait"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.traitAbi == nil {
		return localVarReturnValue, nil, utils.ReportError("traitAbi is required and must be specified", a.client.cfg.Logger)
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "trait_abi", r.traitAbi, "")
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

type ApiGetSmartContractsStatusRequest struct {
	ctx        context.Context
	ApiService *SmartContractsAPIService
	contractId *[]string
}

// contract ids to fetch status for
func (r ApiGetSmartContractsStatusRequest) ContractId(contractId []string) ApiGetSmartContractsStatusRequest {
	r.contractId = &contractId
	return r
}

func (r ApiGetSmartContractsStatusRequest) Execute() (*map[string]models.SmartContractsStatusResponse, *http.Response, error) {
	return r.ApiService.GetSmartContractsStatusExecute(r)
}

/*
GetSmartContractsStatus Get smart contracts status

Retrieves the deployment status of multiple smart contracts.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetSmartContractsStatusRequest
*/
func (a *SmartContractsAPIService) GetSmartContractsStatus(ctx context.Context) ApiGetSmartContractsStatusRequest {
	return ApiGetSmartContractsStatusRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return map[string]SmartContractsStatusResponse
func (a *SmartContractsAPIService) GetSmartContractsStatusExecute(r ApiGetSmartContractsStatusRequest) (*map[string]models.SmartContractsStatusResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *map[string]models.SmartContractsStatusResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SmartContractsAPIService.GetSmartContractsStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/extended/v2/smart-contracts/status"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.contractId == nil {
		return localVarReturnValue, nil, utils.ReportError("contractId is required and must be specified", a.client.cfg.Logger)
	}

	{
		t := *r.contractId
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "contract_id", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "contract_id", t, "multi")
		}
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
