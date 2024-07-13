# \BlocksAPI



Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAverageBlockTimes**](BlocksAPI.md#GetAverageBlockTimes) | **Get** /extended/v2/blocks/average-times | Get average block times
[**GetBlock**](BlocksAPI.md#GetBlock) | **Get** /extended/v2/blocks/{height_or_hash} | Get block
[**GetBlockByBurnBlockHash**](BlocksAPI.md#GetBlockByBurnBlockHash) | **Get** /extended/v1/block/by_burn_block_hash/{burn_block_hash} | Get block by burnchain block hash
[**GetBlockByBurnBlockHeight**](BlocksAPI.md#GetBlockByBurnBlockHeight) | **Get** /extended/v1/block/by_burn_block_height/{burn_block_height} | Get block by burnchain height
[**GetBlockByHash**](BlocksAPI.md#GetBlockByHash) | **Get** /extended/v1/block/{hash} | Get block by hash
[**GetBlockByHeight**](BlocksAPI.md#GetBlockByHeight) | **Get** /extended/v1/block/by_height/{height} | Get block by height
[**GetBlockList**](BlocksAPI.md#GetBlockList) | **Get** /extended/v1/block | Get recent blocks
[**GetBlocks**](BlocksAPI.md#GetBlocks) | **Get** /extended/v2/blocks | Get blocks
[**GetBlocksByBurnBlock**](BlocksAPI.md#GetBlocksByBurnBlock) | **Get** /extended/v2/burn-blocks/{height_or_hash}/blocks | Get blocks by burn block



## GetAverageBlockTimes

> AverageBlockTimesResponse GetAverageBlockTimes(ctx).Execute()

Get average block times



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
	resp, r, err := apiClient.BlocksAPI.GetAverageBlockTimes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlocksAPI.GetAverageBlockTimes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAverageBlockTimes`: AverageBlockTimesResponse
	fmt.Fprintf(os.Stdout, "Response from `BlocksAPI.GetAverageBlockTimes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAverageBlockTimesRequest struct via the builder pattern


### Return type

[**AverageBlockTimesResponse**](AverageBlockTimesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlock

> NakamotoBlock GetBlock(ctx, heightOrHash).Execute()

Get block



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
	heightOrHash := openapiclient.get_burn_block_height_or_hash_parameter{Int32: new(int32)} // GetBurnBlockHeightOrHashParameter | filter by block height, hash, index block hash or the constant `latest` to filter for the most recent block

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.BlocksAPI.GetBlock(context.Background(), heightOrHash).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlocksAPI.GetBlock``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlock`: NakamotoBlock
	fmt.Fprintf(os.Stdout, "Response from `BlocksAPI.GetBlock`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**heightOrHash** | [**GetBurnBlockHeightOrHashParameter**](.md) | filter by block height, hash, index block hash or the constant &#x60;latest&#x60; to filter for the most recent block | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NakamotoBlock**](NakamotoBlock.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlockByBurnBlockHash

> Block GetBlockByBurnBlockHash(ctx, burnBlockHash).Execute()

Get block by burnchain block hash



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
	burnBlockHash := "0x00000000000000000002bba732926cf68b6eda3e2cdbc2a85af79f10efeeeb10" // string | Hash of the burnchain block

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.BlocksAPI.GetBlockByBurnBlockHash(context.Background(), burnBlockHash).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlocksAPI.GetBlockByBurnBlockHash``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlockByBurnBlockHash`: Block
	fmt.Fprintf(os.Stdout, "Response from `BlocksAPI.GetBlockByBurnBlockHash`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**burnBlockHash** | **string** | Hash of the burnchain block | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockByBurnBlockHashRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Block**](Block.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlockByBurnBlockHeight

> Block GetBlockByBurnBlockHeight(ctx, burnBlockHeight).Execute()

Get block by burnchain height



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
	burnBlockHeight := float32(744603) // float32 | Height of the burn chain block

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.BlocksAPI.GetBlockByBurnBlockHeight(context.Background(), burnBlockHeight).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlocksAPI.GetBlockByBurnBlockHeight``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlockByBurnBlockHeight`: Block
	fmt.Fprintf(os.Stdout, "Response from `BlocksAPI.GetBlockByBurnBlockHeight`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**burnBlockHeight** | **float32** | Height of the burn chain block | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockByBurnBlockHeightRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Block**](Block.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlockByHash

> Block GetBlockByHash(ctx, hash).Execute()

Get block by hash



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
	hash := "0x4839a8b01cfb39ffcc0d07d3db31e848d5adf5279d529ed5062300b9f353ff79" // string | Hash of the block

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.BlocksAPI.GetBlockByHash(context.Background(), hash).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlocksAPI.GetBlockByHash``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlockByHash`: Block
	fmt.Fprintf(os.Stdout, "Response from `BlocksAPI.GetBlockByHash`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hash** | **string** | Hash of the block | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockByHashRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Block**](Block.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlockByHeight

> Block GetBlockByHeight(ctx, height).Execute()

Get block by height



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
	height := float32(10000) // float32 | Height of the block

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.BlocksAPI.GetBlockByHeight(context.Background(), height).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlocksAPI.GetBlockByHeight``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlockByHeight`: Block
	fmt.Fprintf(os.Stdout, "Response from `BlocksAPI.GetBlockByHeight`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**height** | **float32** | Height of the block | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockByHeightRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Block**](Block.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlockList

> BlockListResponse GetBlockList(ctx).Limit(limit).Offset(offset).Execute()

Get recent blocks



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
	limit := int32(56) // int32 | max number of blocks to fetch (optional) (default to 20)
	offset := int32(42000) // int32 | index of first block to fetch (optional)

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.BlocksAPI.GetBlockList(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlocksAPI.GetBlockList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlockList`: BlockListResponse
	fmt.Fprintf(os.Stdout, "Response from `BlocksAPI.GetBlockList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | max number of blocks to fetch | [default to 20]
 **offset** | **int32** | index of first block to fetch | 

### Return type

[**BlockListResponse**](BlockListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlocks

> NakamotoBlockListResponse GetBlocks(ctx).Limit(limit).Offset(offset).Execute()

Get blocks



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
	limit := int32(20) // int32 | max number of blocks to fetch (optional)
	offset := int32(0) // int32 | index of first burn block to fetch (optional)

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.BlocksAPI.GetBlocks(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlocksAPI.GetBlocks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlocks`: NakamotoBlockListResponse
	fmt.Fprintf(os.Stdout, "Response from `BlocksAPI.GetBlocks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBlocksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | max number of blocks to fetch | 
 **offset** | **int32** | index of first burn block to fetch | 

### Return type

[**NakamotoBlockListResponse**](NakamotoBlockListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlocksByBurnBlock

> NakamotoBlockListResponse GetBlocksByBurnBlock(ctx, heightOrHash).Limit(limit).Offset(offset).Execute()

Get blocks by burn block



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
	limit := int32(20) // int32 | max number of blocks to fetch (optional)
	offset := int32(0) // int32 | index of first burn block to fetch (optional)

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.BlocksAPI.GetBlocksByBurnBlock(context.Background(), heightOrHash).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlocksAPI.GetBlocksByBurnBlock``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlocksByBurnBlock`: NakamotoBlockListResponse
	fmt.Fprintf(os.Stdout, "Response from `BlocksAPI.GetBlocksByBurnBlock`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**heightOrHash** | [**GetBurnBlockHeightOrHashParameter**](.md) | filter by burn block height, hash, or the constant &#x60;latest&#x60; to filter for the most recent burn block | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlocksByBurnBlockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | max number of blocks to fetch | 
 **offset** | **int32** | index of first burn block to fetch | 

### Return type

[**NakamotoBlockListResponse**](NakamotoBlockListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

