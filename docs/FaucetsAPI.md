# \FaucetsAPI



Method | HTTP request | Description
------------- | ------------- | -------------
[**RunFaucetBtc**](FaucetsAPI.md#RunFaucetBtc) | **Post** /extended/v1/faucets/btc | Add testnet BTC tokens to address
[**RunFaucetStx**](FaucetsAPI.md#RunFaucetStx) | **Post** /extended/v1/faucets/stx | Get STX testnet tokens



## RunFaucetBtc

> RunFaucetResponse RunFaucetBtc(ctx).Address(address).RunFaucetBtcRequest(runFaucetBtcRequest).Execute()

Add testnet BTC tokens to address



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cbadawi/go-stacks"
)

func main() {
	address := "2N4M94S1ZPt8HfxydXzL2P7qyzgVq7MHWts" // string | A valid testnet BTC address
	runFaucetBtcRequest := *openapiclient.NewRunFaucetBtcRequest() // RunFaucetBtcRequest |  (optional)

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.FaucetsAPI.RunFaucetBtc(context.Background()).Address(address).RunFaucetBtcRequest(runFaucetBtcRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FaucetsAPI.RunFaucetBtc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RunFaucetBtc`: RunFaucetResponse
	fmt.Fprintf(os.Stdout, "Response from `FaucetsAPI.RunFaucetBtc`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRunFaucetBtcRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **address** | **string** | A valid testnet BTC address | 
 **runFaucetBtcRequest** | [**RunFaucetBtcRequest**](RunFaucetBtcRequest.md) |  | 

### Return type

[**RunFaucetResponse**](RunFaucetResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RunFaucetStx

> RunFaucetResponse RunFaucetStx(ctx).Address(address).Stacking(stacking).Execute()

Get STX testnet tokens



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cbadawi/go-stacks"
)

func main() {
	address := "ST3M7N9Q9HDRM7RVP1Q26P0EE69358PZZAZD7KMXQ" // string | A valid testnet STX address
	stacking := true // bool | Request the amount of STX tokens needed for individual address stacking (optional) (default to false)

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.FaucetsAPI.RunFaucetStx(context.Background()).Address(address).Stacking(stacking).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FaucetsAPI.RunFaucetStx``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RunFaucetStx`: RunFaucetResponse
	fmt.Fprintf(os.Stdout, "Response from `FaucetsAPI.RunFaucetStx`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRunFaucetStxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **address** | **string** | A valid testnet STX address | 
 **stacking** | **bool** | Request the amount of STX tokens needed for individual address stacking | [default to false]

### Return type

[**RunFaucetResponse**](RunFaucetResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

