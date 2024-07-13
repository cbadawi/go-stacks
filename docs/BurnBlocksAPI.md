# \BurnBlocksAPI


Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBurnBlock**](BurnBlocksAPI.md#GetBurnBlock) | **Get** /extended/v2/burn-blocks/{height_or_hash} | Get burn block
[**GetBurnBlocks**](BurnBlocksAPI.md#GetBurnBlocks) | **Get** /extended/v2/burn-blocks | Get burn blocks



## GetBurnBlock

> BurnBlock GetBurnBlock(ctx, heightOrHash).Execute()

Get burn block



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
	heightOrHash := openapiclient.get_burn_block_height_or_hash_parameter{Int32: new(int32)} // GetBurnBlockHeightOrHashParameter | filter by burn block height, hash, or the constant `latest` to filter for the most recent burn block

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.BurnBlocksAPI.GetBurnBlock(context.Background(), heightOrHash).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BurnBlocksAPI.GetBurnBlock``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBurnBlock`: BurnBlock
	fmt.Fprintf(os.Stdout, "Response from `BurnBlocksAPI.GetBurnBlock`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**heightOrHash** | [**GetBurnBlockHeightOrHashParameter**](.md) | filter by burn block height, hash, or the constant &#x60;latest&#x60; to filter for the most recent burn block | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBurnBlockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BurnBlock**](BurnBlock.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBurnBlocks

> BurnBlockListResponse GetBurnBlocks(ctx).Limit(limit).Offset(offset).Execute()

Get burn blocks



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
	limit := int32(56) // int32 | max number of burn blocks to fetch (optional) (default to 20)
	offset := int32(42000) // int32 | index of first burn block to fetch (optional)

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.BurnBlocksAPI.GetBurnBlocks(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BurnBlocksAPI.GetBurnBlocks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBurnBlocks`: BurnBlockListResponse
	fmt.Fprintf(os.Stdout, "Response from `BurnBlocksAPI.GetBurnBlocks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBurnBlocksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | max number of burn blocks to fetch | [default to 20]
 **offset** | **int32** | index of first burn block to fetch | 

### Return type

[**BurnBlockListResponse**](BurnBlockListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

