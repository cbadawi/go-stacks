# \SearchAPI



Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchById**](SearchAPI.md#SearchById) | **Get** /extended/v1/search/{id} | Search



## SearchById

> SearchResult SearchById(ctx, id).IncludeMetadata(includeMetadata).Execute()

Search



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
	id := "0xcf8b233f19f6c07d2dc1963302d2436efd36e9afac127bf6582824a13961c06d" // string | The hex hash string for a block or transaction, account address, or contract address
	includeMetadata := true // bool | This includes the detailed data for purticular hash in the response (optional)

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.SearchById(context.Background(), id).IncludeMetadata(includeMetadata).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchById`: SearchResult
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The hex hash string for a block or transaction, account address, or contract address | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeMetadata** | **bool** | This includes the detailed data for purticular hash in the response | 

### Return type

[**SearchResult**](SearchResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

