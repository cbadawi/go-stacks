# \InfoAPI



Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCoreApiInfo**](InfoAPI.md#GetCoreApiInfo) | **Get** /v2/info | Get Core API info
[**GetNetworkBlockTimeByNetwork**](InfoAPI.md#GetNetworkBlockTimeByNetwork) | **Get** /extended/v1/info/network_block_time/{network} | Get a given network&#39;s target block time
[**GetNetworkBlockTimes**](InfoAPI.md#GetNetworkBlockTimes) | **Get** /extended/v1/info/network_block_times | Get the network target block time
[**GetPoxInfo**](InfoAPI.md#GetPoxInfo) | **Get** /v2/pox | Get Proof-of-Transfer details
[**GetStatus**](InfoAPI.md#GetStatus) | **Get** /extended | API status
[**GetStxSupply**](InfoAPI.md#GetStxSupply) | **Get** /extended/v1/stx_supply | Get total and unlocked STX supply
[**GetStxSupplyCirculatingPlain**](InfoAPI.md#GetStxSupplyCirculatingPlain) | **Get** /extended/v1/stx_supply/circulating/plain | Get circulating STX supply in plain text format
[**GetStxSupplyTotalSupplyPlain**](InfoAPI.md#GetStxSupplyTotalSupplyPlain) | **Get** /extended/v1/stx_supply/total/plain | Get total STX supply in plain text format
[**GetTotalStxSupplyLegacyFormat**](InfoAPI.md#GetTotalStxSupplyLegacyFormat) | **Get** /extended/v1/stx_supply/legacy_format | Get total and unlocked STX supply (results formatted the same as the legacy 1.0 API)



## GetCoreApiInfo

> CoreNodeInfoResponse GetCoreApiInfo(ctx).Execute()

Get Core API info



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

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.InfoAPI.GetCoreApiInfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InfoAPI.GetCoreApiInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCoreApiInfo`: CoreNodeInfoResponse
	fmt.Fprintf(os.Stdout, "Response from `InfoAPI.GetCoreApiInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCoreApiInfoRequest struct via the builder pattern


### Return type

[**CoreNodeInfoResponse**](CoreNodeInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkBlockTimeByNetwork

> TargetBlockTime GetNetworkBlockTimeByNetwork(ctx, network).Execute()

Get a given network's target block time



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
	network := "mainnet" // string | the target block time for a given network (testnet, mainnet).

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.InfoAPI.GetNetworkBlockTimeByNetwork(context.Background(), network).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InfoAPI.GetNetworkBlockTimeByNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkBlockTimeByNetwork`: TargetBlockTime
	fmt.Fprintf(os.Stdout, "Response from `InfoAPI.GetNetworkBlockTimeByNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**network** | **string** | the target block time for a given network (testnet, mainnet). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkBlockTimeByNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TargetBlockTime**](TargetBlockTime.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkBlockTimes

> NetworkBlockTimesResponse GetNetworkBlockTimes(ctx).Execute()

Get the network target block time



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

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.InfoAPI.GetNetworkBlockTimes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InfoAPI.GetNetworkBlockTimes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkBlockTimes`: NetworkBlockTimesResponse
	fmt.Fprintf(os.Stdout, "Response from `InfoAPI.GetNetworkBlockTimes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkBlockTimesRequest struct via the builder pattern


### Return type

[**NetworkBlockTimesResponse**](NetworkBlockTimesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPoxInfo

> CoreNodePoxResponse GetPoxInfo(ctx).Execute()

Get Proof-of-Transfer details



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

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.InfoAPI.GetPoxInfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InfoAPI.GetPoxInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPoxInfo`: CoreNodePoxResponse
	fmt.Fprintf(os.Stdout, "Response from `InfoAPI.GetPoxInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPoxInfoRequest struct via the builder pattern


### Return type

[**CoreNodePoxResponse**](CoreNodePoxResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStatus

> ServerStatusResponse GetStatus(ctx).Execute()

API status



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

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.InfoAPI.GetStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InfoAPI.GetStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStatus`: ServerStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `InfoAPI.GetStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatusRequest struct via the builder pattern


### Return type

[**ServerStatusResponse**](ServerStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStxSupply

> GetStxSupplyResponse GetStxSupply(ctx).Height(height).Execute()

Get total and unlocked STX supply



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
	height := float32(200) // float32 | Supply details are queried from specified block height. If the block height is not specified, the latest block height is taken as default value. Note that the `block height` is referred to the stacks blockchain. (optional)

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.InfoAPI.GetStxSupply(context.Background()).Height(height).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InfoAPI.GetStxSupply``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStxSupply`: GetStxSupplyResponse
	fmt.Fprintf(os.Stdout, "Response from `InfoAPI.GetStxSupply`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStxSupplyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **height** | **float32** | Supply details are queried from specified block height. If the block height is not specified, the latest block height is taken as default value. Note that the &#x60;block height&#x60; is referred to the stacks blockchain. | 

### Return type

[**GetStxSupplyResponse**](GetStxSupplyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStxSupplyCirculatingPlain

> string GetStxSupplyCirculatingPlain(ctx).Execute()

Get circulating STX supply in plain text format



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

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.InfoAPI.GetStxSupplyCirculatingPlain(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InfoAPI.GetStxSupplyCirculatingPlain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStxSupplyCirculatingPlain`: string
	fmt.Fprintf(os.Stdout, "Response from `InfoAPI.GetStxSupplyCirculatingPlain`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStxSupplyCirculatingPlainRequest struct via the builder pattern


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStxSupplyTotalSupplyPlain

> string GetStxSupplyTotalSupplyPlain(ctx).Execute()

Get total STX supply in plain text format



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

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.InfoAPI.GetStxSupplyTotalSupplyPlain(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InfoAPI.GetStxSupplyTotalSupplyPlain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStxSupplyTotalSupplyPlain`: string
	fmt.Fprintf(os.Stdout, "Response from `InfoAPI.GetStxSupplyTotalSupplyPlain`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStxSupplyTotalSupplyPlainRequest struct via the builder pattern


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTotalStxSupplyLegacyFormat

> GetStxSupplyLegacyFormatResponse GetTotalStxSupplyLegacyFormat(ctx).Height(height).Execute()

Get total and unlocked STX supply (results formatted the same as the legacy 1.0 API)



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
	height := float32(200) // float32 | Supply details are queried from specified block height. If the block height is not specified, the latest block height is taken as default value. (optional)

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.InfoAPI.GetTotalStxSupplyLegacyFormat(context.Background()).Height(height).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InfoAPI.GetTotalStxSupplyLegacyFormat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTotalStxSupplyLegacyFormat`: GetStxSupplyLegacyFormatResponse
	fmt.Fprintf(os.Stdout, "Response from `InfoAPI.GetTotalStxSupplyLegacyFormat`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTotalStxSupplyLegacyFormatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **height** | **float32** | Supply details are queried from specified block height. If the block height is not specified, the latest block height is taken as default value. | 

### Return type

[**GetStxSupplyLegacyFormatResponse**](GetStxSupplyLegacyFormatResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

