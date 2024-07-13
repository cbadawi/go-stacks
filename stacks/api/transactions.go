package api

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strings"

	"github.com/cbadawi/go-stacks/stacks/internal/models"
	"github.com/cbadawi/go-stacks/stacks/utils"
)

// TransactionsAPIService TransactionsAPI service
type TransactionsAPIService service

type ApiGetAddressMempoolTransactionsRequest struct {
	ctx        context.Context
	ApiService *TransactionsAPIService
	address    string
	limit      *int32
	offset     *int32
	unanchored *bool
}

// max number of transactions to fetch
func (r ApiGetAddressMempoolTransactionsRequest) Limit(limit int32) ApiGetAddressMempoolTransactionsRequest {
	r.limit = &limit
	return r
}

// index of first transaction to fetch
func (r ApiGetAddressMempoolTransactionsRequest) Offset(offset int32) ApiGetAddressMempoolTransactionsRequest {
	r.offset = &offset
	return r
}

// Include transaction data from unanchored (i.e. unconfirmed) microblocks
func (r ApiGetAddressMempoolTransactionsRequest) Unanchored(unanchored bool) ApiGetAddressMempoolTransactionsRequest {
	r.unanchored = &unanchored
	return r
}

func (r ApiGetAddressMempoolTransactionsRequest) Execute() (*models.MempoolTransactionListResponse, *http.Response, error) {
	return r.ApiService.GetAddressMempoolTransactionsExecute(r)
}

/*
GetAddressMempoolTransactions Transactions for address

Retrieves all transactions for a given address that are currently in mempool

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param address Transactions for the address
 @return ApiGetAddressMempoolTransactionsRequest
*/
func (a *TransactionsAPIService) GetAddressMempoolTransactions(ctx context.Context, address string) ApiGetAddressMempoolTransactionsRequest {
	return ApiGetAddressMempoolTransactionsRequest{
		ApiService: a,
		ctx:        ctx,
		address:    address,
	}
}

// Execute executes the request
//  @return MempoolTransactionListResponse
func (a *TransactionsAPIService) GetAddressMempoolTransactionsExecute(r ApiGetAddressMempoolTransactionsRequest) (*models.MempoolTransactionListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.MempoolTransactionListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransactionsAPIService.GetAddressMempoolTransactions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/extended/v1/address/{address}/mempool"
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

type ApiGetAddressTransactionEventsRequest struct {
	ctx        context.Context
	ApiService *TransactionsAPIService
	address    string
	txId       string
	limit      *int32
	offset     *int32
}

// Number of events to fetch
func (r ApiGetAddressTransactionEventsRequest) Limit(limit int32) ApiGetAddressTransactionEventsRequest {
	r.limit = &limit
	return r
}

// Index of first event to fetch
func (r ApiGetAddressTransactionEventsRequest) Offset(offset int32) ApiGetAddressTransactionEventsRequest {
	r.offset = &offset
	return r
}

func (r ApiGetAddressTransactionEventsRequest) Execute() (*models.AddressTransactionEventListResponse, *http.Response, error) {
	return r.ApiService.GetAddressTransactionEventsExecute(r)
}

/*
GetAddressTransactionEvents Get events for an address transaction

Retrieves a paginated list of all STX, FT and NFT events concerning a STX address or Smart Contract ID within a specific transaction.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param address STX address or Smart Contract ID
 @param txId Transaction ID
 @return ApiGetAddressTransactionEventsRequest
*/
func (a *TransactionsAPIService) GetAddressTransactionEvents(ctx context.Context, address string, txId string) ApiGetAddressTransactionEventsRequest {
	return ApiGetAddressTransactionEventsRequest{
		ApiService: a,
		ctx:        ctx,
		address:    address,
		txId:       txId,
	}
}

// Execute executes the request
//  @return AddressTransactionEventListResponse
func (a *TransactionsAPIService) GetAddressTransactionEventsExecute(r ApiGetAddressTransactionEventsRequest) (*models.AddressTransactionEventListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.AddressTransactionEventListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransactionsAPIService.GetAddressTransactionEvents")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/extended/v2/addresses/{address}/transactions/{tx_id}/events"
	localVarPath = strings.Replace(localVarPath, "{"+"address"+"}", url.PathEscape(parameterValueToString(r.address, "address")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"tx_id"+"}", url.PathEscape(parameterValueToString(r.txId, "txId")), -1)

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

type ApiGetAddressTransactionsRequest struct {
	ctx        context.Context
	ApiService *TransactionsAPIService
	address    string
	limit      *int32
	offset     *int32
}

// Number of transactions to fetch
func (r ApiGetAddressTransactionsRequest) Limit(limit int32) ApiGetAddressTransactionsRequest {
	r.limit = &limit
	return r
}

// Index of first transaction to fetch
func (r ApiGetAddressTransactionsRequest) Offset(offset int32) ApiGetAddressTransactionsRequest {
	r.offset = &offset
	return r
}

func (r ApiGetAddressTransactionsRequest) Execute() (*models.AddressTransactionsV2ListResponse, *http.Response, error) {
	return r.ApiService.GetAddressTransactionsExecute(r)
}

/*
GetAddressTransactions Get address transactions

Retrieves a paginated list of confirmed transactions sent or received by a STX address or Smart Contract ID, alongside the total amount of STX sent or received and the number of STX, FT and NFT transfers contained within each transaction.

More information on Transaction types can be found [here](https://docs.stacks.co/understand-stacks/transactions#types).


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param address STX address or Smart Contract ID
 @return ApiGetAddressTransactionsRequest
*/
func (a *TransactionsAPIService) GetAddressTransactions(ctx context.Context, address string) ApiGetAddressTransactionsRequest {
	return ApiGetAddressTransactionsRequest{
		ApiService: a,
		ctx:        ctx,
		address:    address,
	}
}

// Execute executes the request
//  @return AddressTransactionsV2ListResponse
func (a *TransactionsAPIService) GetAddressTransactionsExecute(r ApiGetAddressTransactionsRequest) (*models.AddressTransactionsV2ListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.AddressTransactionsV2ListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransactionsAPIService.GetAddressTransactions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/extended/v2/addresses/{address}/transactions"
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

type ApiGetDroppedMempoolTransactionListRequest struct {
	ctx        context.Context
	ApiService *TransactionsAPIService
	limit      *int32
	offset     *int32
}

// max number of mempool transactions to fetch
func (r ApiGetDroppedMempoolTransactionListRequest) Limit(limit int32) ApiGetDroppedMempoolTransactionListRequest {
	r.limit = &limit
	return r
}

// index of first mempool transaction to fetch
func (r ApiGetDroppedMempoolTransactionListRequest) Offset(offset int32) ApiGetDroppedMempoolTransactionListRequest {
	r.offset = &offset
	return r
}

func (r ApiGetDroppedMempoolTransactionListRequest) Execute() (*models.MempoolTransactionListResponse, *http.Response, error) {
	return r.ApiService.GetDroppedMempoolTransactionListExecute(r)
}

/*
GetDroppedMempoolTransactionList Get dropped mempool transactions

Retrieves all recently-broadcast transactions that have been dropped from the mempool.

Transactions are dropped from the mempool if:
 * they were stale and awaiting garbage collection or,
 * were expensive,  or
 * were replaced with a new fee


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetDroppedMempoolTransactionListRequest
*/
func (a *TransactionsAPIService) GetDroppedMempoolTransactionList(ctx context.Context) ApiGetDroppedMempoolTransactionListRequest {
	return ApiGetDroppedMempoolTransactionListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return models.MempoolTransactionListResponse
func (a *TransactionsAPIService) GetDroppedMempoolTransactionListExecute(r ApiGetDroppedMempoolTransactionListRequest) (*models.MempoolTransactionListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.MempoolTransactionListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransactionsAPIService.GetDroppedMempoolTransactionList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/extended/v1/tx/mempool/dropped"

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

type ApiGetFilteredEventsRequest struct {
	ctx        context.Context
	ApiService *TransactionsAPIService
	txId       *string
	address    *string
	limit      *int32
	offset     *int32
	type_      *[]string
}

// Hash of transaction
func (r ApiGetFilteredEventsRequest) TxId(txId string) ApiGetFilteredEventsRequest {
	r.txId = &txId
	return r
}

// Stacks address or a Contract identifier
func (r ApiGetFilteredEventsRequest) Address(address string) ApiGetFilteredEventsRequest {
	r.address = &address
	return r
}

// number of items to return
func (r ApiGetFilteredEventsRequest) Limit(limit int32) ApiGetFilteredEventsRequest {
	r.limit = &limit
	return r
}

// number of items to skip
func (r ApiGetFilteredEventsRequest) Offset(offset int32) ApiGetFilteredEventsRequest {
	r.offset = &offset
	return r
}

// Filter the events on event type
func (r ApiGetFilteredEventsRequest) Type_(type_ []string) ApiGetFilteredEventsRequest {
	r.type_ = &type_
	return r
}

func (r ApiGetFilteredEventsRequest) Execute() (*models.TransactionEventsResponse, *http.Response, error) {
	return r.ApiService.GetFilteredEventsExecute(r)
}

/*
GetFilteredEvents Transaction Events

Retrieves the list of events filtered by principal (STX address or Smart Contract ID), transaction id or event types. The list of event types is ('smart_contract_log', 'stx_lock', 'stx_asset', 'fungible_token_asset', 'non_fungible_token_asset').

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetFilteredEventsRequest
*/
func (a *TransactionsAPIService) GetFilteredEvents(ctx context.Context) ApiGetFilteredEventsRequest {
	return ApiGetFilteredEventsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return TransactionEventsResponse
func (a *TransactionsAPIService) GetFilteredEventsExecute(r ApiGetFilteredEventsRequest) (*models.TransactionEventsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.TransactionEventsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransactionsAPIService.GetFilteredEvents")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/extended/v1/tx/events"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.txId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "tx_id", r.txId, "")
	}
	if r.address != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "address", r.address, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	}
	if r.type_ != nil {
		t := *r.type_
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "type", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "type", t, "multi")
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

type ApiGetMempoolTransactionListRequest struct {
	ctx              context.Context
	ApiService       *TransactionsAPIService
	senderAddress    *string
	recipientAddress *string
	address          *string
	orderBy          *string
	order            *string
	limit            *int32
	offset           *int32
	unanchored       *bool
}

// Filter to only return transactions with this sender address.
func (r ApiGetMempoolTransactionListRequest) SenderAddress(senderAddress string) ApiGetMempoolTransactionListRequest {
	r.senderAddress = &senderAddress
	return r
}

// Filter to only return transactions with this recipient address (only applicable for STX transfer tx types).
func (r ApiGetMempoolTransactionListRequest) RecipientAddress(recipientAddress string) ApiGetMempoolTransactionListRequest {
	r.recipientAddress = &recipientAddress
	return r
}

// Filter to only return transactions with this address as the sender or recipient (recipient only applicable for STX transfer tx types).
func (r ApiGetMempoolTransactionListRequest) Address(address string) ApiGetMempoolTransactionListRequest {
	r.address = &address
	return r
}

// Option to sort results by transaction age, size, or fee rate.
func (r ApiGetMempoolTransactionListRequest) OrderBy(orderBy string) ApiGetMempoolTransactionListRequest {
	r.orderBy = &orderBy
	return r
}

// Option to sort results in ascending or descending order.
func (r ApiGetMempoolTransactionListRequest) Order(order string) ApiGetMempoolTransactionListRequest {
	r.order = &order
	return r
}

// max number of mempool transactions to fetch
func (r ApiGetMempoolTransactionListRequest) Limit(limit int32) ApiGetMempoolTransactionListRequest {
	r.limit = &limit
	return r
}

// index of first mempool transaction to fetch
func (r ApiGetMempoolTransactionListRequest) Offset(offset int32) ApiGetMempoolTransactionListRequest {
	r.offset = &offset
	return r
}

// Include transaction data from unanchored (i.e. unconfirmed) microblocks
func (r ApiGetMempoolTransactionListRequest) Unanchored(unanchored bool) ApiGetMempoolTransactionListRequest {
	r.unanchored = &unanchored
	return r
}

func (r ApiGetMempoolTransactionListRequest) Execute() (*models.MempoolTransactionListResponse, *http.Response, error) {
	return r.ApiService.GetMempoolTransactionListExecute(r)
}

/*
GetMempoolTransactionList Get mempool transactions

Retrieves all transactions that have been recently broadcast to the mempool. These are pending transactions awaiting confirmation.

If you need to monitor new transactions, we highly recommend subscribing to [WebSockets or Socket.io](https://github.com/hirosystems/stacks-blockchain-api/tree/master/client) for real-time updates.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetMempoolTransactionListRequest
*/
func (a *TransactionsAPIService) GetMempoolTransactionList(ctx context.Context) ApiGetMempoolTransactionListRequest {
	return ApiGetMempoolTransactionListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return MempoolTransactionListResponse
func (a *TransactionsAPIService) GetMempoolTransactionListExecute(r ApiGetMempoolTransactionListRequest) (*models.MempoolTransactionListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.MempoolTransactionListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransactionsAPIService.GetMempoolTransactionList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/extended/v1/tx/mempool"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.senderAddress != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sender_address", r.senderAddress, "")
	}
	if r.recipientAddress != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "recipient_address", r.recipientAddress, "")
	}
	if r.address != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "address", r.address, "")
	}
	if r.orderBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order_by", r.orderBy, "")
	}
	if r.order != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order", r.order, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	} else {
		var defaultValue int32 = 20
		r.limit = &defaultValue
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

type ApiGetMempoolTransactionStatsRequest struct {
	ctx        context.Context
	ApiService *TransactionsAPIService
}

func (r ApiGetMempoolTransactionStatsRequest) Execute() (*models.MempoolTransactionStatsResponse, *http.Response, error) {
	return r.ApiService.GetMempoolTransactionStatsExecute(r)
}

/*
GetMempoolTransactionStats Get statistics for mempool transactions

Queries for transactions counts, age (by block height), fees (simple average), and size.
All results broken down by transaction type and percentiles (p25, p50, p75, p95).


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetMempoolTransactionStatsRequest
*/
func (a *TransactionsAPIService) GetMempoolTransactionStats(ctx context.Context) ApiGetMempoolTransactionStatsRequest {
	return ApiGetMempoolTransactionStatsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return MempoolTransactionStatsResponse
func (a *TransactionsAPIService) GetMempoolTransactionStatsExecute(r ApiGetMempoolTransactionStatsRequest) (*models.MempoolTransactionStatsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.MempoolTransactionStatsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransactionsAPIService.GetMempoolTransactionStats")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/extended/v1/tx/mempool/stats"

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

type ApiGetRawTransactionByIdRequest struct {
	ctx        context.Context
	ApiService *TransactionsAPIService
	txId       string
}

func (r ApiGetRawTransactionByIdRequest) Execute() (*models.GetRawTransactionResult, *http.Response, error) {
	return r.ApiService.GetRawTransactionByIdExecute(r)
}

/*
GetRawTransactionById Get Raw Transaction

Retrieves a hex encoded serialized transaction for a given ID


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param txId Hash of transaction
 @return ApiGetRawTransactionByIdRequest
*/
func (a *TransactionsAPIService) GetRawTransactionById(ctx context.Context, txId string) ApiGetRawTransactionByIdRequest {
	return ApiGetRawTransactionByIdRequest{
		ApiService: a,
		ctx:        ctx,
		txId:       txId,
	}
}

// Execute executes the request
//  @return GetRawTransactionResult
func (a *TransactionsAPIService) GetRawTransactionByIdExecute(r ApiGetRawTransactionByIdRequest) (*models.GetRawTransactionResult, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.GetRawTransactionResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransactionsAPIService.GetRawTransactionById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/extended/v1/tx/{tx_id}/raw"
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

type ApiGetTransactionByIdRequest struct {
	ctx         context.Context
	ApiService  *TransactionsAPIService
	txId        string
	eventOffset *int32
	eventLimit  *int32
	unanchored  *bool
}

// The number of events to skip
func (r ApiGetTransactionByIdRequest) EventOffset(eventOffset int32) ApiGetTransactionByIdRequest {
	r.eventOffset = &eventOffset
	return r
}

// The numbers of events to return
func (r ApiGetTransactionByIdRequest) EventLimit(eventLimit int32) ApiGetTransactionByIdRequest {
	r.eventLimit = &eventLimit
	return r
}

// Include transaction data from unanchored (i.e. unconfirmed) microblocks
func (r ApiGetTransactionByIdRequest) Unanchored(unanchored bool) ApiGetTransactionByIdRequest {
	r.unanchored = &unanchored
	return r
}

func (r ApiGetTransactionByIdRequest) Execute() (*models.Transaction, *http.Response, error) {
	return r.ApiService.GetTransactionByIdExecute(r)
}

/*
GetTransactionById Get transaction

Retrieves transaction details for a given transaction ID

`import type { Transaction } from '@stacks/stacks-blockchain-api-types';`


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param txId Hash of transaction
 @return ApiGetTransactionByIdRequest
*/
func (a *TransactionsAPIService) GetTransactionById(ctx context.Context, txId string) ApiGetTransactionByIdRequest {
	return ApiGetTransactionByIdRequest{
		ApiService: a,
		ctx:        ctx,
		txId:       txId,
	}
}

// Execute executes the request
//  @return Transaction
func (a *TransactionsAPIService) GetTransactionByIdExecute(r ApiGetTransactionByIdRequest) (*models.Transaction, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.Transaction
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransactionsAPIService.GetTransactionById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/extended/v1/tx/{tx_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"tx_id"+"}", url.PathEscape(parameterValueToString(r.txId, "txId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.eventOffset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "event_offset", r.eventOffset, "")
	} else {
		var defaultValue int32 = 0
		r.eventOffset = &defaultValue
	}
	if r.eventLimit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "event_limit", r.eventLimit, "")
	} else {
		var defaultValue int32 = 96
		r.eventLimit = &defaultValue
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

type ApiGetTransactionListRequest struct {
	ctx          context.Context
	ApiService   *TransactionsAPIService
	limit        *int32
	offset       *int32
	type_        *[]string
	fromAddress  *string
	toAddress    *string
	sortBy       *string
	startTime    *int32
	endTime      *int32
	contractId   *string
	functionName *string
	nonce        *int32
	order        *string
	unanchored   *bool
}

// max number of transactions to fetch
func (r ApiGetTransactionListRequest) Limit(limit int32) ApiGetTransactionListRequest {
	r.limit = &limit
	return r
}

// index of first transaction to fetch
func (r ApiGetTransactionListRequest) Offset(offset int32) ApiGetTransactionListRequest {
	r.offset = &offset
	return r
}

// Filter by transaction type
func (r ApiGetTransactionListRequest) Type_(type_ []string) ApiGetTransactionListRequest {
	r.type_ = &type_
	return r
}

// Option to filter results by sender address
func (r ApiGetTransactionListRequest) FromAddress(fromAddress string) ApiGetTransactionListRequest {
	r.fromAddress = &fromAddress
	return r
}

// Option to filter results by recipient address
func (r ApiGetTransactionListRequest) ToAddress(toAddress string) ApiGetTransactionListRequest {
	r.toAddress = &toAddress
	return r
}

// Option to sort results by block height, timestamp, or fee
func (r ApiGetTransactionListRequest) SortBy(sortBy string) ApiGetTransactionListRequest {
	r.sortBy = &sortBy
	return r
}

// Filter by transactions after this timestamp (unix timestamp in seconds)
func (r ApiGetTransactionListRequest) StartTime(startTime int32) ApiGetTransactionListRequest {
	r.startTime = &startTime
	return r
}

// Filter by transactions before this timestamp (unix timestamp in seconds)
func (r ApiGetTransactionListRequest) EndTime(endTime int32) ApiGetTransactionListRequest {
	r.endTime = &endTime
	return r
}

// Filter by contract call transactions involving this contract ID
func (r ApiGetTransactionListRequest) ContractId(contractId string) ApiGetTransactionListRequest {
	r.contractId = &contractId
	return r
}

// Filter by contract call transactions involving this function name
func (r ApiGetTransactionListRequest) FunctionName(functionName string) ApiGetTransactionListRequest {
	r.functionName = &functionName
	return r
}

// Filter by transactions with this nonce
func (r ApiGetTransactionListRequest) Nonce(nonce int32) ApiGetTransactionListRequest {
	r.nonce = &nonce
	return r
}

// Option to sort results in ascending or descending order
func (r ApiGetTransactionListRequest) Order(order string) ApiGetTransactionListRequest {
	r.order = &order
	return r
}

// Include transaction data from unanchored (i.e. unconfirmed) microblocks
func (r ApiGetTransactionListRequest) Unanchored(unanchored bool) ApiGetTransactionListRequest {
	r.unanchored = &unanchored
	return r
}

func (r ApiGetTransactionListRequest) Execute() (*models.TransactionResults, *http.Response, error) {
	return r.ApiService.GetTransactionListExecute(r)
}

/*
GetTransactionList Get recent transactions

Retrieves all recently mined transactions

If using TypeScript, import typings for this response from our types package:

`import type { TransactionResults } from '@stacks/stacks-blockchain-api-types';`


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetTransactionListRequest
*/
func (a *TransactionsAPIService) GetTransactionList(ctx context.Context) ApiGetTransactionListRequest {
	return ApiGetTransactionListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return TransactionResults
func (a *TransactionsAPIService) GetTransactionListExecute(r ApiGetTransactionListRequest) (*models.TransactionResults, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.TransactionResults
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransactionsAPIService.GetTransactionList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/extended/v1/tx"

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
	if r.type_ != nil {
		t := *r.type_
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "type", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "type", t, "multi")
		}
	}
	if r.fromAddress != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "from_address", r.fromAddress, "")
	}
	if r.toAddress != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "to_address", r.toAddress, "")
	}
	if r.sortBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort_by", r.sortBy, "")
	} else {
		var defaultValue string = "block_height"
		r.sortBy = &defaultValue
	}
	if r.startTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_time", r.startTime, "")
	}
	if r.endTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_time", r.endTime, "")
	}
	if r.contractId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "contract_id", r.contractId, "")
	}
	if r.functionName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "function_name", r.functionName, "")
	}
	if r.nonce != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nonce", r.nonce, "")
	}
	if r.order != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order", r.order, "")
	} else {
		var defaultValue string = "desc"
		r.order = &defaultValue
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

type ApiGetTransactionsByBlockRequest struct {
	ctx          context.Context
	ApiService   *TransactionsAPIService
	heightOrHash models.GetBurnBlockHeightOrHashParameter
}

func (r ApiGetTransactionsByBlockRequest) Execute() (*models.TransactionResults, *http.Response, error) {
	return r.ApiService.GetTransactionsByBlockExecute(r)
}

/*
GetTransactionsByBlock Get transactions by block

Retrieves transactions confirmed in a single block


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param heightOrHash filter by block height, hash, index block hash or the constant `latest` to filter for the most recent block
 @return ApiGetTransactionsByBlockRequest
*/
func (a *TransactionsAPIService) GetTransactionsByBlock(ctx context.Context, heightOrHash models.GetBurnBlockHeightOrHashParameter) ApiGetTransactionsByBlockRequest {
	return ApiGetTransactionsByBlockRequest{
		ApiService:   a,
		ctx:          ctx,
		heightOrHash: heightOrHash,
	}
}

// Execute executes the request
//  @return TransactionResults
func (a *TransactionsAPIService) GetTransactionsByBlockExecute(r ApiGetTransactionsByBlockRequest) (*models.TransactionResults, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.TransactionResults
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransactionsAPIService.GetTransactionsByBlock")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/extended/v2/blocks/{height_or_hash}/transactions"
	localVarPath = strings.Replace(localVarPath, "{"+"height_or_hash"+"}", url.PathEscape(parameterValueToString(r.heightOrHash, "heightOrHash")), -1)

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

type ApiGetTransactionsByBlockHashRequest struct {
	ctx        context.Context
	ApiService *TransactionsAPIService
	blockHash  string
	limit      *int32
	offset     *int32
}

// max number of transactions to fetch
func (r ApiGetTransactionsByBlockHashRequest) Limit(limit int32) ApiGetTransactionsByBlockHashRequest {
	r.limit = &limit
	return r
}

// index of first transaction to fetch
func (r ApiGetTransactionsByBlockHashRequest) Offset(offset int32) ApiGetTransactionsByBlockHashRequest {
	r.offset = &offset
	return r
}

func (r ApiGetTransactionsByBlockHashRequest) Execute() (*models.TransactionResults, *http.Response, error) {
	return r.ApiService.GetTransactionsByBlockHashExecute(r)
}

/*
GetTransactionsByBlockHash Transactions by block hash

**NOTE:** This endpoint is deprecated in favor of [Get transactions by block](/api/get-transactions-by-block).

Retrieves a list of all transactions within a block for a given block hash.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param blockHash Hash of block
 @return ApiGetTransactionsByBlockHashRequest

Deprecated
*/
func (a *TransactionsAPIService) GetTransactionsByBlockHash(ctx context.Context, blockHash string) ApiGetTransactionsByBlockHashRequest {
	return ApiGetTransactionsByBlockHashRequest{
		ApiService: a,
		ctx:        ctx,
		blockHash:  blockHash,
	}
}

// Execute executes the request
//  @return TransactionResults
// Deprecated
func (a *TransactionsAPIService) GetTransactionsByBlockHashExecute(r ApiGetTransactionsByBlockHashRequest) (*models.TransactionResults, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.TransactionResults
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransactionsAPIService.GetTransactionsByBlockHash")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/extended/v1/tx/block/{block_hash}"
	localVarPath = strings.Replace(localVarPath, "{"+"block_hash"+"}", url.PathEscape(parameterValueToString(r.blockHash, "blockHash")), -1)

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

type ApiGetTransactionsByBlockHeightRequest struct {
	ctx        context.Context
	ApiService *TransactionsAPIService
	height     int32
	limit      *int32
	offset     *int32
	unanchored *bool
}

// max number of transactions to fetch
func (r ApiGetTransactionsByBlockHeightRequest) Limit(limit int32) ApiGetTransactionsByBlockHeightRequest {
	r.limit = &limit
	return r
}

// index of first transaction to fetch
func (r ApiGetTransactionsByBlockHeightRequest) Offset(offset int32) ApiGetTransactionsByBlockHeightRequest {
	r.offset = &offset
	return r
}

// Include transaction data from unanchored (i.e. unconfirmed) microblocks
func (r ApiGetTransactionsByBlockHeightRequest) Unanchored(unanchored bool) ApiGetTransactionsByBlockHeightRequest {
	r.unanchored = &unanchored
	return r
}

func (r ApiGetTransactionsByBlockHeightRequest) Execute() (*models.TransactionResults, *http.Response, error) {
	return r.ApiService.GetTransactionsByBlockHeightExecute(r)
}

/*
GetTransactionsByBlockHeight Transactions by block height

**NOTE:** This endpoint is deprecated in favor of [Get transactions by block](/api/get-transactions-by-block).

Retrieves all transactions within a block at a given height


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param height Height of block
 @return ApiGetTransactionsByBlockHeightRequest

Deprecated
*/
func (a *TransactionsAPIService) GetTransactionsByBlockHeight(ctx context.Context, height int32) ApiGetTransactionsByBlockHeightRequest {
	return ApiGetTransactionsByBlockHeightRequest{
		ApiService: a,
		ctx:        ctx,
		height:     height,
	}
}

// Execute executes the request
//  @return TransactionResults
// Deprecated
func (a *TransactionsAPIService) GetTransactionsByBlockHeightExecute(r ApiGetTransactionsByBlockHeightRequest) (*models.TransactionResults, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.TransactionResults
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransactionsAPIService.GetTransactionsByBlockHeight")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/extended/v1/tx/block_height/{height}"
	localVarPath = strings.Replace(localVarPath, "{"+"height"+"}", url.PathEscape(parameterValueToString(r.height, "height")), -1)

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

type ApiGetTxListDetailsRequest struct {
	ctx         context.Context
	ApiService  *TransactionsAPIService
	txId        *[]string
	eventOffset *int32
	eventLimit  *int32
	unanchored  *bool
}

// Array of transaction ids
func (r ApiGetTxListDetailsRequest) TxId(txId []string) ApiGetTxListDetailsRequest {
	r.txId = &txId
	return r
}

// The number of events to skip
func (r ApiGetTxListDetailsRequest) EventOffset(eventOffset int32) ApiGetTxListDetailsRequest {
	r.eventOffset = &eventOffset
	return r
}

// The numbers of events to return
func (r ApiGetTxListDetailsRequest) EventLimit(eventLimit int32) ApiGetTxListDetailsRequest {
	r.eventLimit = &eventLimit
	return r
}

// Include transaction data from unanchored (i.e. unconfirmed) microblocks
func (r ApiGetTxListDetailsRequest) Unanchored(unanchored bool) ApiGetTxListDetailsRequest {
	r.unanchored = &unanchored
	return r
}

func (r ApiGetTxListDetailsRequest) Execute() (*map[string]models.TransactionList, *http.Response, error) {
	return r.ApiService.GetTxListDetailsExecute(r)
}

/*
GetTxListDetails Get list of details for transactions

Retrieves a list of transactions for a given list of transaction IDs

If using TypeScript, import typings for this response from our types package:

`import type { Transaction } from '@stacks/stacks-blockchain-api-types';`


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetTxListDetailsRequest
*/
func (a *TransactionsAPIService) GetTxListDetails(ctx context.Context) ApiGetTxListDetailsRequest {
	return ApiGetTxListDetailsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return map[string]TransactionList
func (a *TransactionsAPIService) GetTxListDetailsExecute(r ApiGetTxListDetailsRequest) (*map[string]models.TransactionList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *map[string]models.TransactionList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransactionsAPIService.GetTxListDetails")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/extended/v1/tx/multiple"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.txId == nil {
		return localVarReturnValue, nil, utils.ReportError("txId is required and must be specified")
	}

	{
		t := *r.txId
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "tx_id", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "tx_id", t, "multi")
		}
	}
	if r.eventOffset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "event_offset", r.eventOffset, "")
	} else {
		var defaultValue int32 = 0
		r.eventOffset = &defaultValue
	}
	if r.eventLimit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "event_limit", r.eventLimit, "")
	} else {
		var defaultValue int32 = 96
		r.eventLimit = &defaultValue
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

type ApiPostCoreNodeTransactionsRequest struct {
	ctx        context.Context
	ApiService *TransactionsAPIService
	body       *os.File
}

func (r ApiPostCoreNodeTransactionsRequest) Body(body *os.File) ApiPostCoreNodeTransactionsRequest {
	r.body = body
	return r
}

func (r ApiPostCoreNodeTransactionsRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.PostCoreNodeTransactionsExecute(r)
}

/*
PostCoreNodeTransactions Broadcast raw transaction

Broadcasts raw transactions on the network. You can use the [@stacks/transactions](https://github.com/blockstack/stacks.js) project to generate a raw transaction payload.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostCoreNodeTransactionsRequest
*/
func (a *TransactionsAPIService) PostCoreNodeTransactions(ctx context.Context) ApiPostCoreNodeTransactionsRequest {
	return ApiPostCoreNodeTransactionsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return string
func (a *TransactionsAPIService) PostCoreNodeTransactionsExecute(r ApiPostCoreNodeTransactionsRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransactionsAPIService.PostCoreNodeTransactions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/transactions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/octet-stream"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/plain", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v models.PostCoreNodeTransactionsError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
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
