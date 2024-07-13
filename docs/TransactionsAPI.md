# \TransactionsAPI



Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAddressMempoolTransactions**](TransactionsAPI.md#GetAddressMempoolTransactions) | **Get** /extended/v1/address/{address}/mempool | Transactions for address
[**GetAddressTransactionEvents**](TransactionsAPI.md#GetAddressTransactionEvents) | **Get** /extended/v2/addresses/{address}/transactions/{tx_id}/events | Get events for an address transaction
[**GetAddressTransactions**](TransactionsAPI.md#GetAddressTransactions) | **Get** /extended/v2/addresses/{address}/transactions | Get address transactions
[**GetDroppedMempoolTransactionList**](TransactionsAPI.md#GetDroppedMempoolTransactionList) | **Get** /extended/v1/tx/mempool/dropped | Get dropped mempool transactions
[**GetFilteredEvents**](TransactionsAPI.md#GetFilteredEvents) | **Get** /extended/v1/tx/events | Transaction Events
[**GetMempoolTransactionList**](TransactionsAPI.md#GetMempoolTransactionList) | **Get** /extended/v1/tx/mempool | Get mempool transactions
[**GetMempoolTransactionStats**](TransactionsAPI.md#GetMempoolTransactionStats) | **Get** /extended/v1/tx/mempool/stats | Get statistics for mempool transactions
[**GetRawTransactionById**](TransactionsAPI.md#GetRawTransactionById) | **Get** /extended/v1/tx/{tx_id}/raw | Get Raw Transaction
[**GetTransactionById**](TransactionsAPI.md#GetTransactionById) | **Get** /extended/v1/tx/{tx_id} | Get transaction
[**GetTransactionList**](TransactionsAPI.md#GetTransactionList) | **Get** /extended/v1/tx | Get recent transactions
[**GetTransactionsByBlock**](TransactionsAPI.md#GetTransactionsByBlock) | **Get** /extended/v2/blocks/{height_or_hash}/transactions | Get transactions by block
[**GetTransactionsByBlockHash**](TransactionsAPI.md#GetTransactionsByBlockHash) | **Get** /extended/v1/tx/block/{block_hash} | Transactions by block hash
[**GetTransactionsByBlockHeight**](TransactionsAPI.md#GetTransactionsByBlockHeight) | **Get** /extended/v1/tx/block_height/{height} | Transactions by block height
[**GetTxListDetails**](TransactionsAPI.md#GetTxListDetails) | **Get** /extended/v1/tx/multiple | Get list of details for transactions
[**PostCoreNodeTransactions**](TransactionsAPI.md#PostCoreNodeTransactions) | **Post** /v2/transactions | Broadcast raw transaction



## GetAddressMempoolTransactions

> MempoolTransactionListResponse GetAddressMempoolTransactions(ctx, address).Limit(limit).Offset(offset).Unanchored(unanchored).Execute()

Transactions for address



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
	address := "SP197DVH8KTJGX4STM61QN0WJV8Y9QJWXV83ZGNR9" // string | Transactions for the address
	limit := int32(90) // int32 | max number of transactions to fetch (optional)
	offset := int32(42000) // int32 | index of first transaction to fetch (optional)
	unanchored := false // bool | Include transaction data from unanchored (i.e. unconfirmed) microblocks (optional) (default to false)

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.GetAddressMempoolTransactions(context.Background(), address).Limit(limit).Offset(offset).Unanchored(unanchored).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.GetAddressMempoolTransactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAddressMempoolTransactions`: MempoolTransactionListResponse
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.GetAddressMempoolTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | Transactions for the address | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAddressMempoolTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | max number of transactions to fetch | 
 **offset** | **int32** | index of first transaction to fetch | 
 **unanchored** | **bool** | Include transaction data from unanchored (i.e. unconfirmed) microblocks | [default to false]

### Return type

[**MempoolTransactionListResponse**](MempoolTransactionListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAddressTransactionEvents

> AddressTransactionEventListResponse GetAddressTransactionEvents(ctx, address, txId).Limit(limit).Offset(offset).Execute()

Get events for an address transaction



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
	address := "SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0" // string | STX address or Smart Contract ID
	txId := "0x0a411719e3bfde95f9e227a2d7f8fac3d6c646b1e6cc186db0e2838a2c6cd9c0" // string | Transaction ID
	limit := int32(20) // int32 | Number of events to fetch (optional)
	offset := int32(10) // int32 | Index of first event to fetch (optional)

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.GetAddressTransactionEvents(context.Background(), address, txId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.GetAddressTransactionEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAddressTransactionEvents`: AddressTransactionEventListResponse
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.GetAddressTransactionEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | STX address or Smart Contract ID | 
**txId** | **string** | Transaction ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAddressTransactionEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | Number of events to fetch | 
 **offset** | **int32** | Index of first event to fetch | 

### Return type

[**AddressTransactionEventListResponse**](AddressTransactionEventListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAddressTransactions

> AddressTransactionsV2ListResponse GetAddressTransactions(ctx, address).Limit(limit).Offset(offset).Execute()

Get address transactions



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
	address := "SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0" // string | STX address or Smart Contract ID
	limit := int32(20) // int32 | Number of transactions to fetch (optional)
	offset := int32(10) // int32 | Index of first transaction to fetch (optional)

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.GetAddressTransactions(context.Background(), address).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.GetAddressTransactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAddressTransactions`: AddressTransactionsV2ListResponse
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.GetAddressTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | STX address or Smart Contract ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAddressTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Number of transactions to fetch | 
 **offset** | **int32** | Index of first transaction to fetch | 

### Return type

[**AddressTransactionsV2ListResponse**](AddressTransactionsV2ListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDroppedMempoolTransactionList

> MempoolTransactionListResponse GetDroppedMempoolTransactionList(ctx).Limit(limit).Offset(offset).Execute()

Get dropped mempool transactions



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
	limit := int32(56) // int32 | max number of mempool transactions to fetch (optional) (default to 96)
	offset := int32(42000) // int32 | index of first mempool transaction to fetch (optional)

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.GetDroppedMempoolTransactionList(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.GetDroppedMempoolTransactionList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDroppedMempoolTransactionList`: MempoolTransactionListResponse
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.GetDroppedMempoolTransactionList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDroppedMempoolTransactionListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | max number of mempool transactions to fetch | [default to 96]
 **offset** | **int32** | index of first mempool transaction to fetch | 

### Return type

[**MempoolTransactionListResponse**](MempoolTransactionListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFilteredEvents

> TransactionEventsResponse GetFilteredEvents(ctx).TxId(txId).Address(address).Limit(limit).Offset(offset).Type_(type_).Execute()

Transaction Events



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
	txId := "0x29e25515652dad41ef675bd0670964e3d537b80ec19cf6ca6f1dd65d5bc642c5" // string | Hash of transaction (optional)
	address := "ST1HB64MAJ1MBV4CQ80GF01DZS4T1DSMX20ADCRA4" // string | Stacks address or a Contract identifier (optional)
	limit := int32(100) // int32 | number of items to return (optional)
	offset := int32(42000) // int32 | number of items to skip (optional)
	type_ := []string{"Type_example"} // []string | Filter the events on event type (optional)

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.GetFilteredEvents(context.Background()).TxId(txId).Address(address).Limit(limit).Offset(offset).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.GetFilteredEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFilteredEvents`: TransactionEventsResponse
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.GetFilteredEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFilteredEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **txId** | **string** | Hash of transaction | 
 **address** | **string** | Stacks address or a Contract identifier | 
 **limit** | **int32** | number of items to return | 
 **offset** | **int32** | number of items to skip | 
 **type_** | **[]string** | Filter the events on event type | 

### Return type

[**TransactionEventsResponse**](TransactionEventsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMempoolTransactionList

> MempoolTransactionListResponse GetMempoolTransactionList(ctx).SenderAddress(senderAddress).RecipientAddress(recipientAddress).Address(address).OrderBy(orderBy).Order(order).Limit(limit).Offset(offset).Unanchored(unanchored).Execute()

Get mempool transactions



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
	senderAddress := "SP1GPBP8NBRXDRJBFQBV7KMAZX1Z7W2RFWJEH0V10" // string | Filter to only return transactions with this sender address. (optional)
	recipientAddress := "recipientAddress_example" // string | Filter to only return transactions with this recipient address (only applicable for STX transfer tx types). (optional)
	address := "address_example" // string | Filter to only return transactions with this address as the sender or recipient (recipient only applicable for STX transfer tx types). (optional)
	orderBy := "fee" // string | Option to sort results by transaction age, size, or fee rate. (optional)
	order := "asc" // string | Option to sort results in ascending or descending order. (optional)
	limit := int32(20) // int32 | max number of mempool transactions to fetch (optional) (default to 20)
	offset := int32(42000) // int32 | index of first mempool transaction to fetch (optional)
	unanchored := false // bool | Include transaction data from unanchored (i.e. unconfirmed) microblocks (optional) (default to false)

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.GetMempoolTransactionList(context.Background()).SenderAddress(senderAddress).RecipientAddress(recipientAddress).Address(address).OrderBy(orderBy).Order(order).Limit(limit).Offset(offset).Unanchored(unanchored).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.GetMempoolTransactionList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMempoolTransactionList`: MempoolTransactionListResponse
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.GetMempoolTransactionList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMempoolTransactionListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **senderAddress** | **string** | Filter to only return transactions with this sender address. | 
 **recipientAddress** | **string** | Filter to only return transactions with this recipient address (only applicable for STX transfer tx types). | 
 **address** | **string** | Filter to only return transactions with this address as the sender or recipient (recipient only applicable for STX transfer tx types). | 
 **orderBy** | **string** | Option to sort results by transaction age, size, or fee rate. | 
 **order** | **string** | Option to sort results in ascending or descending order. | 
 **limit** | **int32** | max number of mempool transactions to fetch | [default to 20]
 **offset** | **int32** | index of first mempool transaction to fetch | 
 **unanchored** | **bool** | Include transaction data from unanchored (i.e. unconfirmed) microblocks | [default to false]

### Return type

[**MempoolTransactionListResponse**](MempoolTransactionListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMempoolTransactionStats

> MempoolTransactionStatsResponse GetMempoolTransactionStats(ctx).Execute()

Get statistics for mempool transactions



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
	resp, r, err := apiClient.TransactionsAPI.GetMempoolTransactionStats(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.GetMempoolTransactionStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMempoolTransactionStats`: MempoolTransactionStatsResponse
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.GetMempoolTransactionStats`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMempoolTransactionStatsRequest struct via the builder pattern


### Return type

[**MempoolTransactionStatsResponse**](MempoolTransactionStatsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRawTransactionById

> GetRawTransactionResult GetRawTransactionById(ctx, txId).Execute()

Get Raw Transaction



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
	txId := "0x0a411719e3bfde95f9e227a2d7f8fac3d6c646b1e6cc186db0e2838a2c6cd9c0" // string | Hash of transaction

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.GetRawTransactionById(context.Background(), txId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.GetRawTransactionById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRawTransactionById`: GetRawTransactionResult
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.GetRawTransactionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**txId** | **string** | Hash of transaction | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRawTransactionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetRawTransactionResult**](GetRawTransactionResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransactionById

> Transaction GetTransactionById(ctx, txId).EventOffset(eventOffset).EventLimit(eventLimit).Unanchored(unanchored).Execute()

Get transaction



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
	txId := "0x0a411719e3bfde95f9e227a2d7f8fac3d6c646b1e6cc186db0e2838a2c6cd9c0" // string | Hash of transaction
	eventOffset := int32(56) // int32 | The number of events to skip (optional) (default to 0)
	eventLimit := int32(56) // int32 | The numbers of events to return (optional) (default to 96)
	unanchored := false // bool | Include transaction data from unanchored (i.e. unconfirmed) microblocks (optional) (default to false)

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.GetTransactionById(context.Background(), txId).EventOffset(eventOffset).EventLimit(eventLimit).Unanchored(unanchored).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.GetTransactionById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransactionById`: Transaction
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.GetTransactionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**txId** | **string** | Hash of transaction | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eventOffset** | **int32** | The number of events to skip | [default to 0]
 **eventLimit** | **int32** | The numbers of events to return | [default to 96]
 **unanchored** | **bool** | Include transaction data from unanchored (i.e. unconfirmed) microblocks | [default to false]

### Return type

[**Transaction**](Transaction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransactionList

> TransactionResults GetTransactionList(ctx).Limit(limit).Offset(offset).Type_(type_).FromAddress(fromAddress).ToAddress(toAddress).SortBy(sortBy).StartTime(startTime).EndTime(endTime).ContractId(contractId).FunctionName(functionName).Nonce(nonce).Order(order).Unanchored(unanchored).Execute()

Get recent transactions



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
	limit := int32(100) // int32 | max number of transactions to fetch (optional) (default to 96)
	offset := int32(42000) // int32 | index of first transaction to fetch (optional)
	type_ := []string{"Type_example"} // []string | Filter by transaction type (optional)
	fromAddress := "fromAddress_example" // string | Option to filter results by sender address (optional)
	toAddress := "toAddress_example" // string | Option to filter results by recipient address (optional)
	sortBy := "burn_block_time" // string | Option to sort results by block height, timestamp, or fee (optional) (default to "block_height")
	startTime := int32(1704067200) // int32 | Filter by transactions after this timestamp (unix timestamp in seconds) (optional)
	endTime := int32(1706745599) // int32 | Filter by transactions before this timestamp (unix timestamp in seconds) (optional)
	contractId := "SP000000000000000000002Q6VF78.pox-4" // string | Filter by contract call transactions involving this contract ID (optional)
	functionName := "delegate-stx" // string | Filter by contract call transactions involving this function name (optional)
	nonce := int32(123) // int32 | Filter by transactions with this nonce (optional)
	order := "desc" // string | Option to sort results in ascending or descending order (optional) (default to "desc")
	unanchored := false // bool | Include transaction data from unanchored (i.e. unconfirmed) microblocks (optional) (default to false)

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.GetTransactionList(context.Background()).Limit(limit).Offset(offset).Type_(type_).FromAddress(fromAddress).ToAddress(toAddress).SortBy(sortBy).StartTime(startTime).EndTime(endTime).ContractId(contractId).FunctionName(functionName).Nonce(nonce).Order(order).Unanchored(unanchored).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.GetTransactionList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransactionList`: TransactionResults
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.GetTransactionList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | max number of transactions to fetch | [default to 96]
 **offset** | **int32** | index of first transaction to fetch | 
 **type_** | **[]string** | Filter by transaction type | 
 **fromAddress** | **string** | Option to filter results by sender address | 
 **toAddress** | **string** | Option to filter results by recipient address | 
 **sortBy** | **string** | Option to sort results by block height, timestamp, or fee | [default to &quot;block_height&quot;]
 **startTime** | **int32** | Filter by transactions after this timestamp (unix timestamp in seconds) | 
 **endTime** | **int32** | Filter by transactions before this timestamp (unix timestamp in seconds) | 
 **contractId** | **string** | Filter by contract call transactions involving this contract ID | 
 **functionName** | **string** | Filter by contract call transactions involving this function name | 
 **nonce** | **int32** | Filter by transactions with this nonce | 
 **order** | **string** | Option to sort results in ascending or descending order | [default to &quot;desc&quot;]
 **unanchored** | **bool** | Include transaction data from unanchored (i.e. unconfirmed) microblocks | [default to false]

### Return type

[**TransactionResults**](TransactionResults.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransactionsByBlock

> TransactionResults GetTransactionsByBlock(ctx, heightOrHash).Execute()

Get transactions by block



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
	resp, r, err := apiClient.TransactionsAPI.GetTransactionsByBlock(context.Background(), heightOrHash).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.GetTransactionsByBlock``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransactionsByBlock`: TransactionResults
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.GetTransactionsByBlock`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**heightOrHash** | [**GetBurnBlockHeightOrHashParameter**](.md) | filter by block height, hash, index block hash or the constant &#x60;latest&#x60; to filter for the most recent block | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionsByBlockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TransactionResults**](TransactionResults.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransactionsByBlockHash

> TransactionResults GetTransactionsByBlockHash(ctx, blockHash).Limit(limit).Offset(offset).Execute()

Transactions by block hash



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
	blockHash := "0x0a83d82a65460a9e711f85a44616350280040b75317dbe486a923c1131b5ff99" // string | Hash of block
	limit := int32(10) // int32 | max number of transactions to fetch (optional)
	offset := int32(42000) // int32 | index of first transaction to fetch (optional)

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.GetTransactionsByBlockHash(context.Background(), blockHash).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.GetTransactionsByBlockHash``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransactionsByBlockHash`: TransactionResults
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.GetTransactionsByBlockHash`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockHash** | **string** | Hash of block | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionsByBlockHashRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | max number of transactions to fetch | 
 **offset** | **int32** | index of first transaction to fetch | 

### Return type

[**TransactionResults**](TransactionResults.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransactionsByBlockHeight

> TransactionResults GetTransactionsByBlockHeight(ctx, height).Limit(limit).Offset(offset).Unanchored(unanchored).Execute()

Transactions by block height



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
	height := int32(66119) // int32 | Height of block
	limit := int32(10) // int32 | max number of transactions to fetch (optional)
	offset := int32(42000) // int32 | index of first transaction to fetch (optional)
	unanchored := false // bool | Include transaction data from unanchored (i.e. unconfirmed) microblocks (optional) (default to false)

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.GetTransactionsByBlockHeight(context.Background(), height).Limit(limit).Offset(offset).Unanchored(unanchored).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.GetTransactionsByBlockHeight``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransactionsByBlockHeight`: TransactionResults
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.GetTransactionsByBlockHeight`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**height** | **int32** | Height of block | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionsByBlockHeightRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | max number of transactions to fetch | 
 **offset** | **int32** | index of first transaction to fetch | 
 **unanchored** | **bool** | Include transaction data from unanchored (i.e. unconfirmed) microblocks | [default to false]

### Return type

[**TransactionResults**](TransactionResults.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTxListDetails

> map[string]TransactionList GetTxListDetails(ctx).TxId(txId).EventOffset(eventOffset).EventLimit(eventLimit).Unanchored(unanchored).Execute()

Get list of details for transactions



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
	txId := []string{"Inner_example"} // []string | Array of transaction ids
	eventOffset := int32(56) // int32 | The number of events to skip (optional) (default to 0)
	eventLimit := int32(56) // int32 | The numbers of events to return (optional) (default to 96)
	unanchored := false // bool | Include transaction data from unanchored (i.e. unconfirmed) microblocks (optional) (default to false)

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.GetTxListDetails(context.Background()).TxId(txId).EventOffset(eventOffset).EventLimit(eventLimit).Unanchored(unanchored).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.GetTxListDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTxListDetails`: map[string]TransactionList
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.GetTxListDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTxListDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **txId** | **[]string** | Array of transaction ids | 
 **eventOffset** | **int32** | The number of events to skip | [default to 0]
 **eventLimit** | **int32** | The numbers of events to return | [default to 96]
 **unanchored** | **bool** | Include transaction data from unanchored (i.e. unconfirmed) microblocks | [default to false]

### Return type

[**map[string]TransactionList**](TransactionList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCoreNodeTransactions

> string PostCoreNodeTransactions(ctx).Body(body).Execute()

Broadcast raw transaction



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
	body := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.PostCoreNodeTransactions(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.PostCoreNodeTransactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCoreNodeTransactions`: string
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.PostCoreNodeTransactions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCoreNodeTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | ***os.File** |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

