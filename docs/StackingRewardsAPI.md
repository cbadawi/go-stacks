# \StackingRewardsAPI



Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBurnchainRewardList**](StackingRewardsAPI.md#GetBurnchainRewardList) | **Get** /extended/v1/burnchain/rewards | Get recent burnchain reward recipients
[**GetBurnchainRewardListByAddress**](StackingRewardsAPI.md#GetBurnchainRewardListByAddress) | **Get** /extended/v1/burnchain/rewards/{address} | Get recent burnchain reward for the given recipient
[**GetBurnchainRewardSlotHolders**](StackingRewardsAPI.md#GetBurnchainRewardSlotHolders) | **Get** /extended/v1/burnchain/reward_slot_holders | Get recent reward slot holders
[**GetBurnchainRewardSlotHoldersByAddress**](StackingRewardsAPI.md#GetBurnchainRewardSlotHoldersByAddress) | **Get** /extended/v1/burnchain/reward_slot_holders/{address} | Get recent reward slot holder entries for the given address
[**GetBurnchainRewardsTotalByAddress**](StackingRewardsAPI.md#GetBurnchainRewardsTotalByAddress) | **Get** /extended/v1/burnchain/rewards/{address}/total | Get total burnchain rewards for the given recipient



## GetBurnchainRewardList

> BurnchainRewardListResponse GetBurnchainRewardList(ctx).Limit(limit).Offset(offset).Execute()

Get recent burnchain reward recipients



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
	limit := int32(56) // int32 | max number of rewards to fetch (optional) (default to 96)
	offset := int32(42000) // int32 | index of first rewards to fetch (optional)

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.StackingRewardsAPI.GetBurnchainRewardList(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackingRewardsAPI.GetBurnchainRewardList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBurnchainRewardList`: BurnchainRewardListResponse
	fmt.Fprintf(os.Stdout, "Response from `StackingRewardsAPI.GetBurnchainRewardList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBurnchainRewardListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | max number of rewards to fetch | [default to 96]
 **offset** | **int32** | index of first rewards to fetch | 

### Return type

[**BurnchainRewardListResponse**](BurnchainRewardListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBurnchainRewardListByAddress

> BurnchainRewardListResponse GetBurnchainRewardListByAddress(ctx, address).Limit(limit).Offset(offset).Execute()

Get recent burnchain reward for the given recipient



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
	address := "36hQtSEXBMevo5chpxhfAGiCTSC34QKgda" // string | Reward recipient address. Should either be in the native burnchain's format (e.g. B58 for Bitcoin), or if a STX principal address is provided it will be encoded as into the equivalent burnchain format
	limit := int32(56) // int32 | max number of rewards to fetch (optional)
	offset := int32(42000) // int32 | index of first rewards to fetch (optional)

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.StackingRewardsAPI.GetBurnchainRewardListByAddress(context.Background(), address).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackingRewardsAPI.GetBurnchainRewardListByAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBurnchainRewardListByAddress`: BurnchainRewardListResponse
	fmt.Fprintf(os.Stdout, "Response from `StackingRewardsAPI.GetBurnchainRewardListByAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | Reward recipient address. Should either be in the native burnchain&#39;s format (e.g. B58 for Bitcoin), or if a STX principal address is provided it will be encoded as into the equivalent burnchain format | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBurnchainRewardListByAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | max number of rewards to fetch | 
 **offset** | **int32** | index of first rewards to fetch | 

### Return type

[**BurnchainRewardListResponse**](BurnchainRewardListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBurnchainRewardSlotHolders

> BurnchainRewardSlotHolderListResponse GetBurnchainRewardSlotHolders(ctx).Limit(limit).Offset(offset).Execute()

Get recent reward slot holders



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
	limit := int32(56) // int32 | max number of items to fetch (optional) (default to 96)
	offset := int32(42000) // int32 | index of the first items to fetch (optional)

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.StackingRewardsAPI.GetBurnchainRewardSlotHolders(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackingRewardsAPI.GetBurnchainRewardSlotHolders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBurnchainRewardSlotHolders`: BurnchainRewardSlotHolderListResponse
	fmt.Fprintf(os.Stdout, "Response from `StackingRewardsAPI.GetBurnchainRewardSlotHolders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBurnchainRewardSlotHoldersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | max number of items to fetch | [default to 96]
 **offset** | **int32** | index of the first items to fetch | 

### Return type

[**BurnchainRewardSlotHolderListResponse**](BurnchainRewardSlotHolderListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBurnchainRewardSlotHoldersByAddress

> BurnchainRewardSlotHolderListResponse GetBurnchainRewardSlotHoldersByAddress(ctx, address).Limit(limit).Offset(offset).Execute()

Get recent reward slot holder entries for the given address



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
	address := "36hQtSEXBMevo5chpxhfAGiCTSC34QKgda" // string | Reward slot holder recipient address. Should either be in the native burnchain's format (e.g. B58 for Bitcoin), or if a STX principal address is provided it will be encoded as into the equivalent burnchain format
	limit := int32(56) // int32 | max number of items to fetch (optional)
	offset := int32(42000) // int32 | index of the first items to fetch (optional)

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.StackingRewardsAPI.GetBurnchainRewardSlotHoldersByAddress(context.Background(), address).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackingRewardsAPI.GetBurnchainRewardSlotHoldersByAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBurnchainRewardSlotHoldersByAddress`: BurnchainRewardSlotHolderListResponse
	fmt.Fprintf(os.Stdout, "Response from `StackingRewardsAPI.GetBurnchainRewardSlotHoldersByAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | Reward slot holder recipient address. Should either be in the native burnchain&#39;s format (e.g. B58 for Bitcoin), or if a STX principal address is provided it will be encoded as into the equivalent burnchain format | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBurnchainRewardSlotHoldersByAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | max number of items to fetch | 
 **offset** | **int32** | index of the first items to fetch | 

### Return type

[**BurnchainRewardSlotHolderListResponse**](BurnchainRewardSlotHolderListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBurnchainRewardsTotalByAddress

> BurnchainRewardsTotal GetBurnchainRewardsTotalByAddress(ctx, address).Execute()

Get total burnchain rewards for the given recipient



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
	address := "36hQtSEXBMevo5chpxhfAGiCTSC34QKgda" // string | Reward recipient address. Should either be in the native burnchain's format (e.g. B58 for Bitcoin), or if a STX principal address is provided it will be encoded as into the equivalent burnchain format

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.StackingRewardsAPI.GetBurnchainRewardsTotalByAddress(context.Background(), address).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackingRewardsAPI.GetBurnchainRewardsTotalByAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBurnchainRewardsTotalByAddress`: BurnchainRewardsTotal
	fmt.Fprintf(os.Stdout, "Response from `StackingRewardsAPI.GetBurnchainRewardsTotalByAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | Reward recipient address. Should either be in the native burnchain&#39;s format (e.g. B58 for Bitcoin), or if a STX principal address is provided it will be encoded as into the equivalent burnchain format | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBurnchainRewardsTotalByAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BurnchainRewardsTotal**](BurnchainRewardsTotal.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

