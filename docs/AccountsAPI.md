# \AccountsAPI



Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccountAssets**](AccountsAPI.md#GetAccountAssets) | **Get** /extended/v1/address/{principal}/assets | Get account assets
[**GetAccountBalance**](AccountsAPI.md#GetAccountBalance) | **Get** /extended/v1/address/{principal}/balances | Get account balances
[**GetAccountInbound**](AccountsAPI.md#GetAccountInbound) | **Get** /extended/v1/address/{principal}/stx_inbound | Get inbound STX transfers
[**GetAccountInfo**](AccountsAPI.md#GetAccountInfo) | **Get** /v2/accounts/{principal} | Get account info
[**GetAccountNonces**](AccountsAPI.md#GetAccountNonces) | **Get** /extended/v1/address/{principal}/nonces | Get the latest nonce used by an account
[**GetAccountStxBalance**](AccountsAPI.md#GetAccountStxBalance) | **Get** /extended/v1/address/{principal}/stx | Get account STX balance
[**GetAccountTransactions**](AccountsAPI.md#GetAccountTransactions) | **Get** /extended/v1/address/{principal}/transactions | Get account transactions
[**GetAccountTransactionsWithTransfers**](AccountsAPI.md#GetAccountTransactionsWithTransfers) | **Get** /extended/v1/address/{principal}/transactions_with_transfers | Get account transactions including STX transfers for each transaction.
[**GetSingleTransactionWithTransfers**](AccountsAPI.md#GetSingleTransactionWithTransfers) | **Get** /extended/v1/address/{principal}/{tx_id}/with_transfers | Get account transaction information for specific transaction



## GetAccountAssets

> AddressAssetsListResponse GetAccountAssets(ctx, principal).Limit(limit).Offset(offset).Unanchored(unanchored).UntilBlock(untilBlock).Execute()

Get account assets



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
	principal := "SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0" // string | Stacks address or a Contract identifier
	limit := int32(20) // int32 | max number of account assets to fetch (optional)
	offset := int32(42000) // int32 | index of first account assets to fetch (optional)
	unanchored := false // bool | Include transaction data from unanchored (i.e. unconfirmed) microblocks (optional) (default to false)
	untilBlock := "60000" // string | returned data representing the state at that point in time, rather than the current block. Note - Use either of the query parameters but not both at a time. (optional)

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.GetAccountAssets(context.Background(), principal).Limit(limit).Offset(offset).Unanchored(unanchored).UntilBlock(untilBlock).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.GetAccountAssets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountAssets`: AddressAssetsListResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.GetAccountAssets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principal** | **string** | Stacks address or a Contract identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountAssetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | max number of account assets to fetch | 
 **offset** | **int32** | index of first account assets to fetch | 
 **unanchored** | **bool** | Include transaction data from unanchored (i.e. unconfirmed) microblocks | [default to false]
 **untilBlock** | **string** | returned data representing the state at that point in time, rather than the current block. Note - Use either of the query parameters but not both at a time. | 

### Return type

[**AddressAssetsListResponse**](AddressAssetsListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountBalance

> AddressBalanceResponse GetAccountBalance(ctx, principal).Unanchored(unanchored).UntilBlock(untilBlock).Execute()

Get account balances



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
	principal := "SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0" // string | Stacks address or a Contract identifier
	unanchored := false // bool | Include transaction data from unanchored (i.e. unconfirmed) microblocks (optional) (default to false)
	untilBlock := "60000" // string | returned data representing the state up until that point in time, rather than the current block. (optional)

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.GetAccountBalance(context.Background(), principal).Unanchored(unanchored).UntilBlock(untilBlock).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.GetAccountBalance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountBalance`: AddressBalanceResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.GetAccountBalance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principal** | **string** | Stacks address or a Contract identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountBalanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **unanchored** | **bool** | Include transaction data from unanchored (i.e. unconfirmed) microblocks | [default to false]
 **untilBlock** | **string** | returned data representing the state up until that point in time, rather than the current block. | 

### Return type

[**AddressBalanceResponse**](AddressBalanceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountInbound

> AddressStxInboundListResponse GetAccountInbound(ctx, principal).Limit(limit).Offset(offset).Height(height).Unanchored(unanchored).UntilBlock(untilBlock).Execute()

Get inbound STX transfers



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
	principal := "SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0" // string | Stacks address or a Contract identifier
	limit := int32(56) // int32 | number of items to return (optional)
	offset := int32(42000) // int32 | number of items to skip (optional)
	height := float32(8.14) // float32 | Filter for transfers only at this given block height (optional)
	unanchored := false // bool | Include transaction data from unanchored (i.e. unconfirmed) microblocks (optional) (default to false)
	untilBlock := "60000" // string | returned data representing the state up until that point in time, rather than the current block. Note - Use either of the query parameters but not both at a time. (optional)

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.GetAccountInbound(context.Background(), principal).Limit(limit).Offset(offset).Height(height).Unanchored(unanchored).UntilBlock(untilBlock).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.GetAccountInbound``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountInbound`: AddressStxInboundListResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.GetAccountInbound`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principal** | **string** | Stacks address or a Contract identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountInboundRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | number of items to return | 
 **offset** | **int32** | number of items to skip | 
 **height** | **float32** | Filter for transfers only at this given block height | 
 **unanchored** | **bool** | Include transaction data from unanchored (i.e. unconfirmed) microblocks | [default to false]
 **untilBlock** | **string** | returned data representing the state up until that point in time, rather than the current block. Note - Use either of the query parameters but not both at a time. | 

### Return type

[**AddressStxInboundListResponse**](AddressStxInboundListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountInfo

> AccountDataResponse GetAccountInfo(ctx, principal).Proof(proof).Tip(tip).Execute()

Get account info



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
	principal := "SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0" // string | Stacks address or a Contract identifier
	proof := int32(56) // int32 | Returns object without the proof field if set to 0 (optional)
	tip := "tip_example" // string | The Stacks chain tip to query from (optional)

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.GetAccountInfo(context.Background(), principal).Proof(proof).Tip(tip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.GetAccountInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountInfo`: AccountDataResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.GetAccountInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principal** | **string** | Stacks address or a Contract identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **proof** | **int32** | Returns object without the proof field if set to 0 | 
 **tip** | **string** | The Stacks chain tip to query from | 

### Return type

[**AccountDataResponse**](AccountDataResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountNonces

> AddressNonces GetAccountNonces(ctx, principal).BlockHeight(blockHeight).BlockHash(blockHash).Execute()

Get the latest nonce used by an account



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
	principal := "SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0" // string | Stacks address
	blockHeight := float32(66119) // float32 | Optionally get the nonce at a given block height. (optional)
	blockHash := "0x72d53f3cba39e149dcd42708e535bdae03d73e60d2fe853aaf61c0b392f521e9" // string | Optionally get the nonce at a given block hash. Note - Use either of the query parameters but not both at a time. (optional)

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.GetAccountNonces(context.Background(), principal).BlockHeight(blockHeight).BlockHash(blockHash).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.GetAccountNonces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountNonces`: AddressNonces
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.GetAccountNonces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principal** | **string** | Stacks address | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountNoncesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **blockHeight** | **float32** | Optionally get the nonce at a given block height. | 
 **blockHash** | **string** | Optionally get the nonce at a given block hash. Note - Use either of the query parameters but not both at a time. | 

### Return type

[**AddressNonces**](AddressNonces.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountStxBalance

> AddressStxBalanceResponse GetAccountStxBalance(ctx, principal).Unanchored(unanchored).UntilBlock(untilBlock).Execute()

Get account STX balance



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
	principal := "SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0" // string | Stacks address or a Contract identifier.
	unanchored := false // bool | Include transaction data from unanchored (i.e. unconfirmed) microblocks. (optional) (default to false)
	untilBlock := "60000" // string | returned data representing the state up until that point in time, rather than the current block. Note - Use either of the query parameters but not both at a time. (optional)

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.GetAccountStxBalance(context.Background(), principal).Unanchored(unanchored).UntilBlock(untilBlock).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.GetAccountStxBalance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountStxBalance`: AddressStxBalanceResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.GetAccountStxBalance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principal** | **string** | Stacks address or a Contract identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountStxBalanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **unanchored** | **bool** | Include transaction data from unanchored (i.e. unconfirmed) microblocks. | [default to false]
 **untilBlock** | **string** | returned data representing the state up until that point in time, rather than the current block. Note - Use either of the query parameters but not both at a time. | 

### Return type

[**AddressStxBalanceResponse**](AddressStxBalanceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountTransactions

> AddressTransactionsListResponse GetAccountTransactions(ctx, principal).Limit(limit).Offset(offset).Height(height).Unanchored(unanchored).UntilBlock(untilBlock).Execute()

Get account transactions



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
	principal := "SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0" // string | Stacks address or a Contract identifier
	limit := int32(42000) // int32 | max number of account transactions to fetch (optional)
	offset := int32(42000) // int32 | index of first account transaction to fetch (optional)
	height := float32(42000) // float32 | Filter for transactions only at this given block height (optional)
	unanchored := false // bool | Include transaction data from unanchored (i.e. unconfirmed) microblocks (optional) (default to false)
	untilBlock := "60000" // string | returned data representing the state up until that point in time, rather than the current block. Note - Use either of the query parameters but not both at a time. (optional)

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.GetAccountTransactions(context.Background(), principal).Limit(limit).Offset(offset).Height(height).Unanchored(unanchored).UntilBlock(untilBlock).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.GetAccountTransactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountTransactions`: AddressTransactionsListResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.GetAccountTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principal** | **string** | Stacks address or a Contract identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | max number of account transactions to fetch | 
 **offset** | **int32** | index of first account transaction to fetch | 
 **height** | **float32** | Filter for transactions only at this given block height | 
 **unanchored** | **bool** | Include transaction data from unanchored (i.e. unconfirmed) microblocks | [default to false]
 **untilBlock** | **string** | returned data representing the state up until that point in time, rather than the current block. Note - Use either of the query parameters but not both at a time. | 

### Return type

[**AddressTransactionsListResponse**](AddressTransactionsListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountTransactionsWithTransfers

> AddressTransactionsWithTransfersListResponse GetAccountTransactionsWithTransfers(ctx, principal).Limit(limit).Offset(offset).Height(height).Unanchored(unanchored).UntilBlock(untilBlock).Execute()

Get account transactions including STX transfers for each transaction.



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
	principal := "SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0" // string | Stacks address or a Contract identifier
	limit := int32(20) // int32 | max number of account transactions to fetch (optional)
	offset := int32(10) // int32 | index of first account transaction to fetch (optional)
	height := float32(66119) // float32 | Filter for transactions only at this given block height (optional)
	unanchored := false // bool | Include transaction data from unanchored (i.e. unconfirmed) microblocks (optional) (default to false)
	untilBlock := "60000" // string | returned data representing the state up until that point in time, rather than the current block. (optional)

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.GetAccountTransactionsWithTransfers(context.Background(), principal).Limit(limit).Offset(offset).Height(height).Unanchored(unanchored).UntilBlock(untilBlock).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.GetAccountTransactionsWithTransfers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountTransactionsWithTransfers`: AddressTransactionsWithTransfersListResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.GetAccountTransactionsWithTransfers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principal** | **string** | Stacks address or a Contract identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountTransactionsWithTransfersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | max number of account transactions to fetch | 
 **offset** | **int32** | index of first account transaction to fetch | 
 **height** | **float32** | Filter for transactions only at this given block height | 
 **unanchored** | **bool** | Include transaction data from unanchored (i.e. unconfirmed) microblocks | [default to false]
 **untilBlock** | **string** | returned data representing the state up until that point in time, rather than the current block. | 

### Return type

[**AddressTransactionsWithTransfersListResponse**](AddressTransactionsWithTransfersListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSingleTransactionWithTransfers

> AddressTransactionWithTransfers GetSingleTransactionWithTransfers(ctx, principal, txId).Execute()

Get account transaction information for specific transaction



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
	principal := "SP3FBR2AGK5H9QBDH3EEN6DF8EK8JY7RX8QJ5SVTE" // string | Stacks address or a contract identifier
	txId := "0x34d79c7cfc2fe525438736733e501a4bf0308a5556e3e080d1e2c0858aad7448" // string | Transaction id

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.GetSingleTransactionWithTransfers(context.Background(), principal, txId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.GetSingleTransactionWithTransfers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSingleTransactionWithTransfers`: AddressTransactionWithTransfers
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.GetSingleTransactionWithTransfers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principal** | **string** | Stacks address or a contract identifier | 
**txId** | **string** | Transaction id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSingleTransactionWithTransfersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AddressTransactionWithTransfers**](AddressTransactionWithTransfers.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

