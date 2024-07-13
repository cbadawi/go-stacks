# \NonFungibleTokensAPI



Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNftHistory**](NonFungibleTokensAPI.md#GetNftHistory) | **Get** /extended/v1/tokens/nft/history | Non-Fungible Token history
[**GetNftHoldings**](NonFungibleTokensAPI.md#GetNftHoldings) | **Get** /extended/v1/tokens/nft/holdings | Non-Fungible Token holdings
[**GetNftMints**](NonFungibleTokensAPI.md#GetNftMints) | **Get** /extended/v1/tokens/nft/mints | Non-Fungible Token mints



## GetNftHistory

> NonFungibleTokenHistoryEventList GetNftHistory(ctx).AssetIdentifier(assetIdentifier).Value(value).Limit(limit).Offset(offset).Unanchored(unanchored).TxMetadata(txMetadata).Execute()

Non-Fungible Token history



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
	assetIdentifier := "SP2X0TZ59D5SZ8ACQ6YMCHHNR2ZN51Z32E2CJ173.the-explorer-guild::The-Explorer-Guild" // string | token asset class identifier
	value := "0x0100000000000000000000000000000803" // string | hex representation of the token's unique value
	limit := int32(56) // int32 | max number of events to fetch (optional) (default to 50)
	offset := int32(42000) // int32 | index of first event to fetch (optional) (default to 0)
	unanchored := false // bool | whether or not to include events from unconfirmed transactions (optional) (default to false)
	txMetadata := true // bool | whether or not to include the complete transaction metadata instead of just `tx_id`. Enabling this option can affect performance and response times. (optional) (default to false)

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.NonFungibleTokensAPI.GetNftHistory(context.Background()).AssetIdentifier(assetIdentifier).Value(value).Limit(limit).Offset(offset).Unanchored(unanchored).TxMetadata(txMetadata).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NonFungibleTokensAPI.GetNftHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNftHistory`: NonFungibleTokenHistoryEventList
	fmt.Fprintf(os.Stdout, "Response from `NonFungibleTokensAPI.GetNftHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNftHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assetIdentifier** | **string** | token asset class identifier | 
 **value** | **string** | hex representation of the token&#39;s unique value | 
 **limit** | **int32** | max number of events to fetch | [default to 50]
 **offset** | **int32** | index of first event to fetch | [default to 0]
 **unanchored** | **bool** | whether or not to include events from unconfirmed transactions | [default to false]
 **txMetadata** | **bool** | whether or not to include the complete transaction metadata instead of just &#x60;tx_id&#x60;. Enabling this option can affect performance and response times. | [default to false]

### Return type

[**NonFungibleTokenHistoryEventList**](NonFungibleTokenHistoryEventList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNftHoldings

> NonFungibleTokenHoldingsList GetNftHoldings(ctx).Principal(principal).AssetIdentifiers(assetIdentifiers).Limit(limit).Offset(offset).Unanchored(unanchored).TxMetadata(txMetadata).Execute()

Non-Fungible Token holdings



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
	principal := "SPNWZ5V2TPWGQGVDR6T7B6RQ4XMGZ4PXTEE0VQ0S.marketplace-v3" // string | token owner's STX address or Smart Contract ID
	assetIdentifiers := []string{"Inner_example"} // []string | identifiers of the token asset classes to filter for (optional)
	limit := int32(56) // int32 | max number of tokens to fetch (optional) (default to 50)
	offset := int32(42000) // int32 | index of first tokens to fetch (optional) (default to 0)
	unanchored := false // bool | whether or not to include tokens from unconfirmed transactions (optional) (default to false)
	txMetadata := true // bool | whether or not to include the complete transaction metadata instead of just `tx_id`. Enabling this option can affect performance and response times. (optional) (default to false)

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.NonFungibleTokensAPI.GetNftHoldings(context.Background()).Principal(principal).AssetIdentifiers(assetIdentifiers).Limit(limit).Offset(offset).Unanchored(unanchored).TxMetadata(txMetadata).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NonFungibleTokensAPI.GetNftHoldings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNftHoldings`: NonFungibleTokenHoldingsList
	fmt.Fprintf(os.Stdout, "Response from `NonFungibleTokensAPI.GetNftHoldings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNftHoldingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **principal** | **string** | token owner&#39;s STX address or Smart Contract ID | 
 **assetIdentifiers** | **[]string** | identifiers of the token asset classes to filter for | 
 **limit** | **int32** | max number of tokens to fetch | [default to 50]
 **offset** | **int32** | index of first tokens to fetch | [default to 0]
 **unanchored** | **bool** | whether or not to include tokens from unconfirmed transactions | [default to false]
 **txMetadata** | **bool** | whether or not to include the complete transaction metadata instead of just &#x60;tx_id&#x60;. Enabling this option can affect performance and response times. | [default to false]

### Return type

[**NonFungibleTokenHoldingsList**](NonFungibleTokenHoldingsList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNftMints

> NonFungibleTokenMintList GetNftMints(ctx).AssetIdentifier(assetIdentifier).Limit(limit).Offset(offset).Unanchored(unanchored).TxMetadata(txMetadata).Execute()

Non-Fungible Token mints



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
	assetIdentifier := "SP2X0TZ59D5SZ8ACQ6YMCHHNR2ZN51Z32E2CJ173.the-explorer-guild::The-Explorer-Guild" // string | token asset class identifier
	limit := int32(56) // int32 | max number of events to fetch (optional) (default to 50)
	offset := int32(42000) // int32 | index of first event to fetch (optional) (default to 0)
	unanchored := false // bool | whether or not to include events from unconfirmed transactions (optional) (default to false)
	txMetadata := true // bool | whether or not to include the complete transaction metadata instead of just `tx_id`. Enabling this option can affect performance and response times. (optional) (default to false)

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.NonFungibleTokensAPI.GetNftMints(context.Background()).AssetIdentifier(assetIdentifier).Limit(limit).Offset(offset).Unanchored(unanchored).TxMetadata(txMetadata).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NonFungibleTokensAPI.GetNftMints``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNftMints`: NonFungibleTokenMintList
	fmt.Fprintf(os.Stdout, "Response from `NonFungibleTokensAPI.GetNftMints`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNftMintsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assetIdentifier** | **string** | token asset class identifier | 
 **limit** | **int32** | max number of events to fetch | [default to 50]
 **offset** | **int32** | index of first event to fetch | [default to 0]
 **unanchored** | **bool** | whether or not to include events from unconfirmed transactions | [default to false]
 **txMetadata** | **bool** | whether or not to include the complete transaction metadata instead of just &#x60;tx_id&#x60;. Enabling this option can affect performance and response times. | [default to false]

### Return type

[**NonFungibleTokenMintList**](NonFungibleTokenMintList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

