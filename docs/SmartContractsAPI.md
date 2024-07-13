# \SmartContractsAPI



Method | HTTP request | Description
------------- | ------------- | -------------
[**CallReadOnlyFunction**](SmartContractsAPI.md#CallReadOnlyFunction) | **Post** /v2/contracts/call-read/{contract_address}/{contract_name}/{function_name} | Call read-only function
[**GetContractById**](SmartContractsAPI.md#GetContractById) | **Get** /extended/v1/contract/{contract_id} | Get contract info
[**GetContractDataMapEntry**](SmartContractsAPI.md#GetContractDataMapEntry) | **Post** /v2/map_entry/{contract_address}/{contract_name}/{map_name} | Get specific data-map inside a contract
[**GetContractEventsById**](SmartContractsAPI.md#GetContractEventsById) | **Get** /extended/v1/contract/{contract_id}/events | Get contract events
[**GetContractInterface**](SmartContractsAPI.md#GetContractInterface) | **Get** /v2/contracts/interface/{contract_address}/{contract_name} | Get contract interface
[**GetContractSource**](SmartContractsAPI.md#GetContractSource) | **Get** /v2/contracts/source/{contract_address}/{contract_name} | Get contract source
[**GetContractsByTrait**](SmartContractsAPI.md#GetContractsByTrait) | **Get** /extended/v1/contract/by_trait | Get contracts by trait
[**GetSmartContractsStatus**](SmartContractsAPI.md#GetSmartContractsStatus) | **Get** /extended/v2/smart-contracts/status | Get smart contracts status



## CallReadOnlyFunction

> ReadOnlyFunctionSuccessResponse CallReadOnlyFunction(ctx, contractAddress, contractName, functionName).ReadOnlyFunctionArgs(readOnlyFunctionArgs).Tip(tip).Execute()

Call read-only function



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
	contractAddress := "SP187Y7NRSG3T9Z9WTSWNEN3XRV1YSJWS81C7JKV7" // string | Stacks address
	contractName := "imaginary-friends-zebras" // string | Contract name
	functionName := "get-token-uri" // string | Function name
	readOnlyFunctionArgs := *openapiclient.NewReadOnlyFunctionArgs("Sender_example", []string{"Arguments_example"}) // ReadOnlyFunctionArgs | map of arguments and the simulated tx-sender where sender is either a Contract identifier or a normal Stacks address, and arguments is an array of hex serialized Clarity values.
	tip := "tip_example" // string | The Stacks chain tip to query from (optional)

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.SmartContractsAPI.CallReadOnlyFunction(context.Background(), contractAddress, contractName, functionName).ReadOnlyFunctionArgs(readOnlyFunctionArgs).Tip(tip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmartContractsAPI.CallReadOnlyFunction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CallReadOnlyFunction`: ReadOnlyFunctionSuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `SmartContractsAPI.CallReadOnlyFunction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractAddress** | **string** | Stacks address | 
**contractName** | **string** | Contract name | 
**functionName** | **string** | Function name | 

### Other Parameters

Other parameters are passed through a pointer to a apiCallReadOnlyFunctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **readOnlyFunctionArgs** | [**ReadOnlyFunctionArgs**](ReadOnlyFunctionArgs.md) | map of arguments and the simulated tx-sender where sender is either a Contract identifier or a normal Stacks address, and arguments is an array of hex serialized Clarity values. | 
 **tip** | **string** | The Stacks chain tip to query from | 

### Return type

[**ReadOnlyFunctionSuccessResponse**](ReadOnlyFunctionSuccessResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContractById

> SmartContract GetContractById(ctx, contractId).Unanchored(unanchored).Execute()

Get contract info



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
	contractId := "SP6P4EJF0VG8V0RB3TQQKJBHDQKEF6NVRD1KZE3C.satoshibles" // string | Contract identifier formatted as `<contract_address>.<contract_name>`
	unanchored := false // bool | Include transaction data from unanchored (i.e. unconfirmed) microblocks (optional) (default to false)

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.SmartContractsAPI.GetContractById(context.Background(), contractId).Unanchored(unanchored).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmartContractsAPI.GetContractById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContractById`: SmartContract
	fmt.Fprintf(os.Stdout, "Response from `SmartContractsAPI.GetContractById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **string** | Contract identifier formatted as &#x60;&lt;contract_address&gt;.&lt;contract_name&gt;&#x60; | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContractByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **unanchored** | **bool** | Include transaction data from unanchored (i.e. unconfirmed) microblocks | [default to false]

### Return type

[**SmartContract**](SmartContract.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContractDataMapEntry

> MapEntryResponse GetContractDataMapEntry(ctx, contractAddress, contractName, mapName).Key(key).Proof(proof).Tip(tip).Execute()

Get specific data-map inside a contract



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
	contractAddress := "SPSCWDV3RKV5ZRN1FQD84YE1NQFEDJ9R1F4DYQ11" // string | Stacks address
	contractName := "newyorkcitycoin-core-v2" // string | Contract name
	mapName := "approved-contracts" // string | Map name
	key := "{"$ref":"./entities/contracts/get-specific-data-map-inside-contract.example.json"}" // string | Hex string serialization of the lookup key (which should be a Clarity value)
	proof := int32(56) // int32 | Returns object without the proof field when set to 0 (optional)
	tip := "tip_example" // string | The Stacks chain tip to query from (optional)

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.SmartContractsAPI.GetContractDataMapEntry(context.Background(), contractAddress, contractName, mapName).Key(key).Proof(proof).Tip(tip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmartContractsAPI.GetContractDataMapEntry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContractDataMapEntry`: MapEntryResponse
	fmt.Fprintf(os.Stdout, "Response from `SmartContractsAPI.GetContractDataMapEntry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractAddress** | **string** | Stacks address | 
**contractName** | **string** | Contract name | 
**mapName** | **string** | Map name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContractDataMapEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **key** | **string** | Hex string serialization of the lookup key (which should be a Clarity value) | 
 **proof** | **int32** | Returns object without the proof field when set to 0 | 
 **tip** | **string** | The Stacks chain tip to query from | 

### Return type

[**MapEntryResponse**](MapEntryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContractEventsById

> TransactionEvent GetContractEventsById(ctx, contractId).Limit(limit).Offset(offset).Unanchored(unanchored).Execute()

Get contract events



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
	contractId := "SP6P4EJF0VG8V0RB3TQQKJBHDQKEF6NVRD1KZE3C.satoshibles" // string | Contract identifier formatted as `<contract_address>.<contract_name>`
	limit := int32(56) // int32 | max number of contract events to fetch (optional)
	offset := int32(42000) // int32 | index of first contract event to fetch (optional)
	unanchored := false // bool | Include transaction data from unanchored (i.e. unconfirmed) microblocks (optional) (default to false)

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.SmartContractsAPI.GetContractEventsById(context.Background(), contractId).Limit(limit).Offset(offset).Unanchored(unanchored).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmartContractsAPI.GetContractEventsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContractEventsById`: TransactionEvent
	fmt.Fprintf(os.Stdout, "Response from `SmartContractsAPI.GetContractEventsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **string** | Contract identifier formatted as &#x60;&lt;contract_address&gt;.&lt;contract_name&gt;&#x60; | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContractEventsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | max number of contract events to fetch | 
 **offset** | **int32** | index of first contract event to fetch | 
 **unanchored** | **bool** | Include transaction data from unanchored (i.e. unconfirmed) microblocks | [default to false]

### Return type

[**TransactionEvent**](TransactionEvent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContractInterface

> ContractInterfaceResponse GetContractInterface(ctx, contractAddress, contractName).Tip(tip).Execute()

Get contract interface



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
	contractAddress := "SP6P4EJF0VG8V0RB3TQQKJBHDQKEF6NVRD1KZE3C" // string | Stacks address
	contractName := "satoshibles" // string | Contract name
	tip := "tip_example" // string | The Stacks chain tip to query from (optional)

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.SmartContractsAPI.GetContractInterface(context.Background(), contractAddress, contractName).Tip(tip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmartContractsAPI.GetContractInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContractInterface`: ContractInterfaceResponse
	fmt.Fprintf(os.Stdout, "Response from `SmartContractsAPI.GetContractInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractAddress** | **string** | Stacks address | 
**contractName** | **string** | Contract name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContractInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tip** | **string** | The Stacks chain tip to query from | 

### Return type

[**ContractInterfaceResponse**](ContractInterfaceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContractSource

> ContractSourceResponse GetContractSource(ctx, contractAddress, contractName).Proof(proof).Tip(tip).Execute()

Get contract source



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
	contractAddress := "SP6P4EJF0VG8V0RB3TQQKJBHDQKEF6NVRD1KZE3C" // string | Stacks address
	contractName := "satoshibles" // string | Contract name
	proof := int32(56) // int32 | Returns object without the proof field if set to 0 (optional)
	tip := "tip_example" // string | The Stacks chain tip to query from (optional)

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.SmartContractsAPI.GetContractSource(context.Background(), contractAddress, contractName).Proof(proof).Tip(tip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmartContractsAPI.GetContractSource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContractSource`: ContractSourceResponse
	fmt.Fprintf(os.Stdout, "Response from `SmartContractsAPI.GetContractSource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractAddress** | **string** | Stacks address | 
**contractName** | **string** | Contract name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContractSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **proof** | **int32** | Returns object without the proof field if set to 0 | 
 **tip** | **string** | The Stacks chain tip to query from | 

### Return type

[**ContractSourceResponse**](ContractSourceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContractsByTrait

> ContractListResponse GetContractsByTrait(ctx).TraitAbi(traitAbi).Limit(limit).Offset(offset).Execute()

Get contracts by trait



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
	traitAbi := "traitAbi_example" // string | JSON abi of the trait.
	limit := int32(56) // int32 | max number of contracts fetch (optional)
	offset := int32(42000) // int32 | index of first contract event to fetch (optional)

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.SmartContractsAPI.GetContractsByTrait(context.Background()).TraitAbi(traitAbi).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmartContractsAPI.GetContractsByTrait``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContractsByTrait`: ContractListResponse
	fmt.Fprintf(os.Stdout, "Response from `SmartContractsAPI.GetContractsByTrait`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetContractsByTraitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **traitAbi** | **string** | JSON abi of the trait. | 
 **limit** | **int32** | max number of contracts fetch | 
 **offset** | **int32** | index of first contract event to fetch | 

### Return type

[**ContractListResponse**](ContractListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSmartContractsStatus

> map[string]SmartContractsStatusResponse GetSmartContractsStatus(ctx).ContractId(contractId).Execute()

Get smart contracts status



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
	contractId := []string{"Inner_example"} // []string | contract ids to fetch status for

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.SmartContractsAPI.GetSmartContractsStatus(context.Background()).ContractId(contractId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmartContractsAPI.GetSmartContractsStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSmartContractsStatus`: map[string]SmartContractsStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `SmartContractsAPI.GetSmartContractsStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSmartContractsStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contractId** | **[]string** | contract ids to fetch status for | 

### Return type

[**map[string]SmartContractsStatusResponse**](SmartContractsStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

