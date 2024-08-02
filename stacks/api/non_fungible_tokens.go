package api

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"reflect"

	"github.com/cbadawi/go-stacks/stacks/models"
	"github.com/cbadawi/go-stacks/stacks/utils"
)

// NonFungibleTokensAPIService NonFungibleTokensAPI service
type NonFungibleTokensAPIService service

type ApiGetNftHistoryRequest struct {
	ctx             context.Context
	ApiService      *NonFungibleTokensAPIService
	assetIdentifier *string
	value           *string
	limit           *int32
	offset          *int32
	unanchored      *bool
	txMetadata      *bool
}

// token asset class identifier
func (r ApiGetNftHistoryRequest) AssetIdentifier(assetIdentifier string) ApiGetNftHistoryRequest {
	r.assetIdentifier = &assetIdentifier
	return r
}

// hex representation of the token&#39;s unique value
func (r ApiGetNftHistoryRequest) Value(value string) ApiGetNftHistoryRequest {
	r.value = &value
	return r
}

// max number of events to fetch
func (r ApiGetNftHistoryRequest) Limit(limit int32) ApiGetNftHistoryRequest {
	r.limit = &limit
	return r
}

// index of first event to fetch
func (r ApiGetNftHistoryRequest) Offset(offset int32) ApiGetNftHistoryRequest {
	r.offset = &offset
	return r
}

// whether or not to include events from unconfirmed transactions
func (r ApiGetNftHistoryRequest) Unanchored(unanchored bool) ApiGetNftHistoryRequest {
	r.unanchored = &unanchored
	return r
}

// whether or not to include the complete transaction metadata instead of just &#x60;tx_id&#x60;. Enabling this option can affect performance and response times.
func (r ApiGetNftHistoryRequest) TxMetadata(txMetadata bool) ApiGetNftHistoryRequest {
	r.txMetadata = &txMetadata
	return r
}

func (r ApiGetNftHistoryRequest) Execute() (*models.NonFungibleTokenHistoryEventList, *http.Response, error) {
	return r.ApiService.GetNftHistoryExecute(r)
}

/*
GetNftHistory Non-Fungible Token history

Retrieves all events relevant to a Non-Fungible Token. Useful to determine the ownership history of a particular asset.

More information on Non-Fungible Tokens on the Stacks blockchain can be found [here](https://docs.stacks.co/write-smart-contracts/tokens#non-fungible-tokens-nfts).


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetNftHistoryRequest
*/
func (a *NonFungibleTokensAPIService) GetNftHistory(ctx context.Context) ApiGetNftHistoryRequest {
	return ApiGetNftHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return NonFungibleTokenHistoryEventList
func (a *NonFungibleTokensAPIService) GetNftHistoryExecute(r ApiGetNftHistoryRequest) (*models.NonFungibleTokenHistoryEventList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.NonFungibleTokenHistoryEventList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NonFungibleTokensAPIService.GetNftHistory")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/extended/v1/tokens/nft/history"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.assetIdentifier == nil {
		return localVarReturnValue, nil, utils.ReportError("assetIdentifier is required and must be specified", a.client.cfg.Logger)
	}
	if r.value == nil {
		return localVarReturnValue, nil, utils.ReportError("value is required and must be specified", a.client.cfg.Logger)
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "asset_identifier", r.assetIdentifier, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "value", r.value, "")
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	} else {
		var defaultValue int32 = 50
		r.limit = &defaultValue
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	} else {
		var defaultValue int32 = 0
		r.offset = &defaultValue
	}
	if r.unanchored != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "unanchored", r.unanchored, "")
	} else {
		var defaultValue bool = false
		r.unanchored = &defaultValue
	}
	if r.txMetadata != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "tx_metadata", r.txMetadata, "")
	} else {
		var defaultValue bool = false
		r.txMetadata = &defaultValue
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

type ApiGetNftHoldingsRequest struct {
	ctx              context.Context
	ApiService       *NonFungibleTokensAPIService
	principal        *string
	assetIdentifiers *[]string
	limit            *int32
	offset           *int32
	unanchored       *bool
	txMetadata       *bool
}

// token owner&#39;s STX address or Smart Contract ID
func (r ApiGetNftHoldingsRequest) Principal(principal string) ApiGetNftHoldingsRequest {
	r.principal = &principal
	return r
}

// identifiers of the token asset classes to filter for
func (r ApiGetNftHoldingsRequest) AssetIdentifiers(assetIdentifiers []string) ApiGetNftHoldingsRequest {
	r.assetIdentifiers = &assetIdentifiers
	return r
}

// max number of tokens to fetch
func (r ApiGetNftHoldingsRequest) Limit(limit int32) ApiGetNftHoldingsRequest {
	r.limit = &limit
	return r
}

// index of first tokens to fetch
func (r ApiGetNftHoldingsRequest) Offset(offset int32) ApiGetNftHoldingsRequest {
	r.offset = &offset
	return r
}

// whether or not to include tokens from unconfirmed transactions
func (r ApiGetNftHoldingsRequest) Unanchored(unanchored bool) ApiGetNftHoldingsRequest {
	r.unanchored = &unanchored
	return r
}

// whether or not to include the complete transaction metadata instead of just &#x60;tx_id&#x60;. Enabling this option can affect performance and response times.
func (r ApiGetNftHoldingsRequest) TxMetadata(txMetadata bool) ApiGetNftHoldingsRequest {
	r.txMetadata = &txMetadata
	return r
}

func (r ApiGetNftHoldingsRequest) Execute() (*models.NonFungibleTokenHoldingsList, *http.Response, error) {
	return r.ApiService.GetNftHoldingsExecute(r)
}

/*
GetNftHoldings Non-Fungible Token holdings

Retrieves the list of Non-Fungible Tokens owned by the given principal (STX address or Smart Contract ID).
Results can be filtered by one or more asset identifiers and can include metadata about the transaction that made the principal own each token.

More information on Non-Fungible Tokens on the Stacks blockchain can be found [here](https://docs.stacks.co/write-smart-contracts/tokens#non-fungible-tokens-nfts).


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetNftHoldingsRequest
*/
func (a *NonFungibleTokensAPIService) GetNftHoldings(ctx context.Context) ApiGetNftHoldingsRequest {
	return ApiGetNftHoldingsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return NonFungibleTokenHoldingsList
func (a *NonFungibleTokensAPIService) GetNftHoldingsExecute(r ApiGetNftHoldingsRequest) (*models.NonFungibleTokenHoldingsList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.NonFungibleTokenHoldingsList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NonFungibleTokensAPIService.GetNftHoldings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/extended/v1/tokens/nft/holdings"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.principal == nil {
		return localVarReturnValue, nil, utils.ReportError("principal is required and must be specified", a.client.cfg.Logger)
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "principal", r.principal, "")
	if r.assetIdentifiers != nil {
		t := *r.assetIdentifiers
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "asset_identifiers", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "asset_identifiers", t, "multi")
		}
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	} else {
		var defaultValue int32 = 50
		r.limit = &defaultValue
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	} else {
		var defaultValue int32 = 0
		r.offset = &defaultValue
	}
	if r.unanchored != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "unanchored", r.unanchored, "")
	} else {
		var defaultValue bool = false
		r.unanchored = &defaultValue
	}
	if r.txMetadata != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "tx_metadata", r.txMetadata, "")
	} else {
		var defaultValue bool = false
		r.txMetadata = &defaultValue
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

type ApiGetNftMintsRequest struct {
	ctx             context.Context
	ApiService      *NonFungibleTokensAPIService
	assetIdentifier *string
	limit           *int32
	offset          *int32
	unanchored      *bool
	txMetadata      *bool
}

// token asset class identifier
func (r ApiGetNftMintsRequest) AssetIdentifier(assetIdentifier string) ApiGetNftMintsRequest {
	r.assetIdentifier = &assetIdentifier
	return r
}

// max number of events to fetch
func (r ApiGetNftMintsRequest) Limit(limit int32) ApiGetNftMintsRequest {
	r.limit = &limit
	return r
}

// index of first event to fetch
func (r ApiGetNftMintsRequest) Offset(offset int32) ApiGetNftMintsRequest {
	r.offset = &offset
	return r
}

// whether or not to include events from unconfirmed transactions
func (r ApiGetNftMintsRequest) Unanchored(unanchored bool) ApiGetNftMintsRequest {
	r.unanchored = &unanchored
	return r
}

// whether or not to include the complete transaction metadata instead of just &#x60;tx_id&#x60;. Enabling this option can affect performance and response times.
func (r ApiGetNftMintsRequest) TxMetadata(txMetadata bool) ApiGetNftMintsRequest {
	r.txMetadata = &txMetadata
	return r
}

func (r ApiGetNftMintsRequest) Execute() (*models.NonFungibleTokenMintList, *http.Response, error) {
	return r.ApiService.GetNftMintsExecute(r)
}

/*
GetNftMints Non-Fungible Token mints

Retrieves all mint events for a Non-Fungible Token asset class. Useful to determine which NFTs of a particular collection have been claimed.

More information on Non-Fungible Tokens on the Stacks blockchain can be found [here](https://docs.stacks.co/write-smart-contracts/tokens#non-fungible-tokens-nfts).


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetNftMintsRequest
*/
func (a *NonFungibleTokensAPIService) GetNftMints(ctx context.Context) ApiGetNftMintsRequest {
	return ApiGetNftMintsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return NonFungibleTokenMintList
func (a *NonFungibleTokensAPIService) GetNftMintsExecute(r ApiGetNftMintsRequest) (*models.NonFungibleTokenMintList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.NonFungibleTokenMintList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NonFungibleTokensAPIService.GetNftMints")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/extended/v1/tokens/nft/mints"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.assetIdentifier == nil {
		return localVarReturnValue, nil, utils.ReportError("assetIdentifier is required and must be specified", a.client.cfg.Logger)
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "asset_identifier", r.assetIdentifier, "")
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	} else {
		var defaultValue int32 = 50
		r.limit = &defaultValue
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	} else {
		var defaultValue int32 = 0
		r.offset = &defaultValue
	}
	if r.unanchored != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "unanchored", r.unanchored, "")
	} else {
		var defaultValue bool = false
		r.unanchored = &defaultValue
	}
	if r.txMetadata != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "tx_metadata", r.txMetadata, "")
	} else {
		var defaultValue bool = false
		r.txMetadata = &defaultValue
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
