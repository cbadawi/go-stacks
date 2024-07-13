# \ProofOfTransferAPI



Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPoxCycle**](ProofOfTransferAPI.md#GetPoxCycle) | **Get** /extended/v2/pox/cycles/{cycle_number} | Get PoX cycle
[**GetPoxCycleSigner**](ProofOfTransferAPI.md#GetPoxCycleSigner) | **Get** /extended/v2/pox/cycles/{cycle_number}/signers/{signer_key} | Get signer in PoX cycle
[**GetPoxCycleSignerStackers**](ProofOfTransferAPI.md#GetPoxCycleSignerStackers) | **Get** /extended/v2/pox/cycles/{cycle_number}/signers/{signer_key}/stackers | Get stackers for signer in PoX cycle
[**GetPoxCycleSigners**](ProofOfTransferAPI.md#GetPoxCycleSigners) | **Get** /extended/v2/pox/cycles/{cycle_number}/signers | Get signers in PoX cycle
[**GetPoxCycles**](ProofOfTransferAPI.md#GetPoxCycles) | **Get** /extended/v2/pox/cycles | Get PoX cycles



## GetPoxCycle

> PoxCycle GetPoxCycle(ctx, cycleNumber).Execute()

Get PoX cycle



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
	cycleNumber := int32(56) // int32 | PoX cycle number

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.ProofOfTransferAPI.GetPoxCycle(context.Background(), cycleNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProofOfTransferAPI.GetPoxCycle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPoxCycle`: PoxCycle
	fmt.Fprintf(os.Stdout, "Response from `ProofOfTransferAPI.GetPoxCycle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cycleNumber** | **int32** | PoX cycle number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPoxCycleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PoxCycle**](PoxCycle.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPoxCycleSigner

> PoxSigner GetPoxCycleSigner(ctx, cycleNumber, signerKey).Execute()

Get signer in PoX cycle



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
	cycleNumber := int32(56) // int32 | PoX cycle number
	signerKey := "0x038e3c4529395611be9abf6fa3b6987e81d402385e3d605a073f42f407565a4a3d" // string | Signer key

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.ProofOfTransferAPI.GetPoxCycleSigner(context.Background(), cycleNumber, signerKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProofOfTransferAPI.GetPoxCycleSigner``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPoxCycleSigner`: PoxSigner
	fmt.Fprintf(os.Stdout, "Response from `ProofOfTransferAPI.GetPoxCycleSigner`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cycleNumber** | **int32** | PoX cycle number | 
**signerKey** | **string** | Signer key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPoxCycleSignerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PoxSigner**](PoxSigner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPoxCycleSignerStackers

> PoxCycleSignerStackersListResponse GetPoxCycleSignerStackers(ctx, cycleNumber, signerKey).Execute()

Get stackers for signer in PoX cycle



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
	cycleNumber := int32(56) // int32 | PoX cycle number
	signerKey := "0x038e3c4529395611be9abf6fa3b6987e81d402385e3d605a073f42f407565a4a3d" // string | Signer key

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.ProofOfTransferAPI.GetPoxCycleSignerStackers(context.Background(), cycleNumber, signerKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProofOfTransferAPI.GetPoxCycleSignerStackers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPoxCycleSignerStackers`: PoxCycleSignerStackersListResponse
	fmt.Fprintf(os.Stdout, "Response from `ProofOfTransferAPI.GetPoxCycleSignerStackers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cycleNumber** | **int32** | PoX cycle number | 
**signerKey** | **string** | Signer key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPoxCycleSignerStackersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PoxCycleSignerStackersListResponse**](PoxCycleSignerStackersListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPoxCycleSigners

> PoxCycleSignersListResponse GetPoxCycleSigners(ctx, cycleNumber).Execute()

Get signers in PoX cycle



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
	cycleNumber := int32(56) // int32 | PoX cycle number

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.ProofOfTransferAPI.GetPoxCycleSigners(context.Background(), cycleNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProofOfTransferAPI.GetPoxCycleSigners``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPoxCycleSigners`: PoxCycleSignersListResponse
	fmt.Fprintf(os.Stdout, "Response from `ProofOfTransferAPI.GetPoxCycleSigners`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cycleNumber** | **int32** | PoX cycle number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPoxCycleSignersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PoxCycleSignersListResponse**](PoxCycleSignersListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPoxCycles

> PoxCycleListResponse GetPoxCycles(ctx).Limit(limit).Offset(offset).Execute()

Get PoX cycles



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
	limit := int32(56) // int32 | max number of cycles to fetch (optional) (default to 20)
	offset := int32(20) // int32 | index of first cycle to fetch (optional)

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.ProofOfTransferAPI.GetPoxCycles(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProofOfTransferAPI.GetPoxCycles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPoxCycles`: PoxCycleListResponse
	fmt.Fprintf(os.Stdout, "Response from `ProofOfTransferAPI.GetPoxCycles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPoxCyclesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | max number of cycles to fetch | [default to 20]
 **offset** | **int32** | index of first cycle to fetch | 

### Return type

[**PoxCycleListResponse**](PoxCycleListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

