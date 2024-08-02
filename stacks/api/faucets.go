package api

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"

	"github.com/cbadawi/go-stacks/stacks/models"
	"github.com/cbadawi/go-stacks/stacks/utils"
)

// FaucetsAPIService FaucetsAPI service
type FaucetsAPIService service

type ApiRunFaucetBtcRequest struct {
	ctx                 context.Context
	ApiService          *FaucetsAPIService
	address             *string
	runFaucetBtcRequest *models.RunFaucetBtcRequest
}

// A valid testnet BTC address
func (r ApiRunFaucetBtcRequest) Address(address string) ApiRunFaucetBtcRequest {
	r.address = &address
	return r
}

func (r ApiRunFaucetBtcRequest) RunFaucetBtcRequest(runFaucetBtcRequest models.RunFaucetBtcRequest) ApiRunFaucetBtcRequest {
	r.runFaucetBtcRequest = &runFaucetBtcRequest
	return r
}

func (r ApiRunFaucetBtcRequest) Execute() (*models.RunFaucetResponse, *http.Response, error) {
	return r.ApiService.RunFaucetBtcExecute(r)
}

/*
RunFaucetBtc Add testnet BTC tokens to address

Add 1 BTC token to the specified testnet BTC address.

The endpoint returns the transaction ID, which you can use to view the transaction in a testnet Bitcoin block
explorer. The tokens are delivered once the transaction has been included in a block.

**Note:** This is a testnet only endpoint. This endpoint will not work on the mainnet.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRunFaucetBtcRequest
*/
func (a *FaucetsAPIService) RunFaucetBtc(ctx context.Context) ApiRunFaucetBtcRequest {
	return ApiRunFaucetBtcRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return RunFaucetResponse
func (a *FaucetsAPIService) RunFaucetBtcExecute(r ApiRunFaucetBtcRequest) (*models.RunFaucetResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.RunFaucetResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FaucetsAPIService.RunFaucetBtc")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/extended/v1/faucets/btc"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.address == nil {
		return localVarReturnValue, nil, utils.ReportError("address is required and must be specified", a.client.cfg.Logger)
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "address", r.address, "")
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
	localVarPostBody = r.runFaucetBtcRequest
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
		if localVarHTTPResponse.StatusCode == 403 {
			var v models.RunFaucetBtc403Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
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

type ApiRunFaucetStxRequest struct {
	ctx        context.Context
	ApiService *FaucetsAPIService
	address    *string
	stacking   *bool
}

// A valid testnet STX address
func (r ApiRunFaucetStxRequest) Address(address string) ApiRunFaucetStxRequest {
	r.address = &address
	return r
}

// Request the amount of STX tokens needed for individual address stacking
func (r ApiRunFaucetStxRequest) Stacking(stacking bool) ApiRunFaucetStxRequest {
	r.stacking = &stacking
	return r
}

func (r ApiRunFaucetStxRequest) Execute() (*models.RunFaucetResponse, *http.Response, error) {
	return r.ApiService.RunFaucetStxExecute(r)
}

/*
RunFaucetStx Get STX testnet tokens

Add 500 STX tokens to the specified testnet address. Testnet STX addresses begin with `ST`. If the `stacking`
parameter is set to `true`, the faucet will add the required number of tokens for individual stacking to the
specified testnet address.

The endpoint returns the transaction ID, which you can use to view the transaction in the
[Stacks Explorer](https://explorer.hiro.so/?chain=testnet). The tokens are delivered once the transaction has
been included in an anchor block.

A common reason for failed faucet transactions is that the faucet has run out of tokens. If you are experiencing
failed faucet transactions to a testnet address, you can get help in [Discord](https://stacks.chat).

**Note:** This is a testnet only endpoint. This endpoint will not work on the mainnet.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRunFaucetStxRequest
*/
func (a *FaucetsAPIService) RunFaucetStx(ctx context.Context) ApiRunFaucetStxRequest {
	return ApiRunFaucetStxRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return RunFaucetResponse
func (a *FaucetsAPIService) RunFaucetStxExecute(r ApiRunFaucetStxRequest) (*models.RunFaucetResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.RunFaucetResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FaucetsAPIService.RunFaucetStx")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/extended/v1/faucets/stx"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.address == nil {
		return localVarReturnValue, nil, utils.ReportError("address is required and must be specified", a.client.cfg.Logger)
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "address", r.address, "")
	if r.stacking != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "stacking", r.stacking, "")
	} else {
		var defaultValue bool = false
		r.stacking = &defaultValue
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
