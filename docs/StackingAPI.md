# \StackingAPI



Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPoolDelegations**](StackingAPI.md#GetPoolDelegations) | **Get** /extended/beta/stacking/{pool_principal}/delegations | Stacking pool members



## GetPoolDelegations

> PoolDelegationsResponse GetPoolDelegations(ctx, poolPrincipal).AfterBlock(afterBlock).Unanchored(unanchored).Limit(limit).Offset(offset).Execute()

Stacking pool members



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
	poolPrincipal := "SPSCWDV3RKV5ZRN1FQD84YE1NQFEDJ9R1F4DYQ11" // string | Address principal of the stacking pool delegator
	afterBlock := int32(56) // int32 | If specified, only delegation events after the given block will be included (optional)
	unanchored := false // bool | whether or not to include Stackers from unconfirmed transactions (optional) (default to false)
	limit := int32(100) // int32 | number of items to return (optional) (default to 100)
	offset := int32(300) // int32 | number of items to skip (optional) (default to 0)

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.StackingAPI.GetPoolDelegations(context.Background(), poolPrincipal).AfterBlock(afterBlock).Unanchored(unanchored).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackingAPI.GetPoolDelegations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPoolDelegations`: PoolDelegationsResponse
	fmt.Fprintf(os.Stdout, "Response from `StackingAPI.GetPoolDelegations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolPrincipal** | **string** | Address principal of the stacking pool delegator | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPoolDelegationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **afterBlock** | **int32** | If specified, only delegation events after the given block will be included | 
 **unanchored** | **bool** | whether or not to include Stackers from unconfirmed transactions | [default to false]
 **limit** | **int32** | number of items to return | [default to 100]
 **offset** | **int32** | number of items to skip | [default to 0]

### Return type

[**PoolDelegationsResponse**](PoolDelegationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

