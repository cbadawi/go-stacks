# \RosettaAPI



Method | HTTP request | Description
------------- | ------------- | -------------
[**RosettaAccountBalance**](RosettaAPI.md#RosettaAccountBalance) | **Post** /rosetta/v1/account/balance | Get an Account Balance
[**RosettaBlock**](RosettaAPI.md#RosettaBlock) | **Post** /rosetta/v1/block | Get a Block
[**RosettaBlockTransaction**](RosettaAPI.md#RosettaBlockTransaction) | **Post** /rosetta/v1/block/transaction | Get a Block Transaction
[**RosettaConstructionCombine**](RosettaAPI.md#RosettaConstructionCombine) | **Post** /rosetta/v1/construction/combine | Create Network Transaction from Signatures
[**RosettaConstructionDerive**](RosettaAPI.md#RosettaConstructionDerive) | **Post** /rosetta/v1/construction/derive | Derive an AccountIdentifier from a PublicKey
[**RosettaConstructionHash**](RosettaAPI.md#RosettaConstructionHash) | **Post** /rosetta/v1/construction/hash | Get the Hash of a Signed Transaction
[**RosettaConstructionMetadata**](RosettaAPI.md#RosettaConstructionMetadata) | **Post** /rosetta/v1/construction/metadata | Get Metadata for Transaction Construction
[**RosettaConstructionParse**](RosettaAPI.md#RosettaConstructionParse) | **Post** /rosetta/v1/construction/parse | Parse a Transaction
[**RosettaConstructionPayloads**](RosettaAPI.md#RosettaConstructionPayloads) | **Post** /rosetta/v1/construction/payloads | Generate an Unsigned Transaction and Signing Payloads
[**RosettaConstructionPreprocess**](RosettaAPI.md#RosettaConstructionPreprocess) | **Post** /rosetta/v1/construction/preprocess | Create a Request to Fetch Metadata
[**RosettaConstructionSubmit**](RosettaAPI.md#RosettaConstructionSubmit) | **Post** /rosetta/v1/construction/submit | Submit a Signed Transaction
[**RosettaMempool**](RosettaAPI.md#RosettaMempool) | **Post** /rosetta/v1/mempool | Get All Mempool Transactions
[**RosettaMempoolTransaction**](RosettaAPI.md#RosettaMempoolTransaction) | **Post** /rosetta/v1/mempool/transaction | Get a Mempool Transaction
[**RosettaNetworkList**](RosettaAPI.md#RosettaNetworkList) | **Post** /rosetta/v1/network/list | Get List of Available Networks
[**RosettaNetworkOptions**](RosettaAPI.md#RosettaNetworkOptions) | **Post** /rosetta/v1/network/options | Get Network Options
[**RosettaNetworkStatus**](RosettaAPI.md#RosettaNetworkStatus) | **Post** /rosetta/v1/network/status | Get Network Status



## RosettaAccountBalance

> RosettaAccountBalanceResponse RosettaAccountBalance(ctx).RosettaAccountBalanceRequest(rosettaAccountBalanceRequest).Execute()

Get an Account Balance



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
	rosettaAccountBalanceRequest := *openapiclient.NewRosettaAccountBalanceRequest(*openapiclient.NewNetworkIdentifier("Blockchain_example", "Network_example"), *openapiclient.NewRosettaAccount("Address_example")) // RosettaAccountBalanceRequest | 

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.RosettaAPI.RosettaAccountBalance(context.Background()).RosettaAccountBalanceRequest(rosettaAccountBalanceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RosettaAPI.RosettaAccountBalance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RosettaAccountBalance`: RosettaAccountBalanceResponse
	fmt.Fprintf(os.Stdout, "Response from `RosettaAPI.RosettaAccountBalance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRosettaAccountBalanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rosettaAccountBalanceRequest** | [**RosettaAccountBalanceRequest**](RosettaAccountBalanceRequest.md) |  | 

### Return type

[**RosettaAccountBalanceResponse**](RosettaAccountBalanceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RosettaBlock

> RosettaBlockResponse RosettaBlock(ctx).RosettaBlockRequest(rosettaBlockRequest).Execute()

Get a Block



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
	rosettaBlockRequest := *openapiclient.NewRosettaBlockRequest(*openapiclient.NewNetworkIdentifier("Blockchain_example", "Network_example"), *openapiclient.NewRosettaPartialBlockIdentifier("Hash_example", int32(123))) // RosettaBlockRequest | 

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.RosettaAPI.RosettaBlock(context.Background()).RosettaBlockRequest(rosettaBlockRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RosettaAPI.RosettaBlock``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RosettaBlock`: RosettaBlockResponse
	fmt.Fprintf(os.Stdout, "Response from `RosettaAPI.RosettaBlock`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRosettaBlockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rosettaBlockRequest** | [**RosettaBlockRequest**](RosettaBlockRequest.md) |  | 

### Return type

[**RosettaBlockResponse**](RosettaBlockResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RosettaBlockTransaction

> RosettaBlockTransactionResponse RosettaBlockTransaction(ctx).RosettaBlockTransactionRequest(rosettaBlockTransactionRequest).Execute()

Get a Block Transaction



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
	rosettaBlockTransactionRequest := *openapiclient.NewRosettaBlockTransactionRequest(*openapiclient.NewNetworkIdentifier("Blockchain_example", "Network_example"), *openapiclient.NewRosettaPartialBlockIdentifier("Hash_example", int32(123)), *openapiclient.NewTransactionIdentifier("Hash_example")) // RosettaBlockTransactionRequest | 

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.RosettaAPI.RosettaBlockTransaction(context.Background()).RosettaBlockTransactionRequest(rosettaBlockTransactionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RosettaAPI.RosettaBlockTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RosettaBlockTransaction`: RosettaBlockTransactionResponse
	fmt.Fprintf(os.Stdout, "Response from `RosettaAPI.RosettaBlockTransaction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRosettaBlockTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rosettaBlockTransactionRequest** | [**RosettaBlockTransactionRequest**](RosettaBlockTransactionRequest.md) |  | 

### Return type

[**RosettaBlockTransactionResponse**](RosettaBlockTransactionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RosettaConstructionCombine

> RosettaConstructionCombineResponse RosettaConstructionCombine(ctx).RosettaConstructionCombineRequest(rosettaConstructionCombineRequest).Execute()

Create Network Transaction from Signatures



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
	rosettaConstructionCombineRequest := *openapiclient.NewRosettaConstructionCombineRequest(*openapiclient.NewNetworkIdentifier("Blockchain_example", "Network_example"), "UnsignedTransaction_example", []openapiclient.RosettaSignature{*openapiclient.NewRosettaSignature(*openapiclient.NewSigningPayload("HexBytes_example"), *openapiclient.NewRosettaPublicKey("HexBytes_example", "CurveType_example"), "SignatureType_example", "HexBytes_example")}) // RosettaConstructionCombineRequest | 

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.RosettaAPI.RosettaConstructionCombine(context.Background()).RosettaConstructionCombineRequest(rosettaConstructionCombineRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RosettaAPI.RosettaConstructionCombine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RosettaConstructionCombine`: RosettaConstructionCombineResponse
	fmt.Fprintf(os.Stdout, "Response from `RosettaAPI.RosettaConstructionCombine`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRosettaConstructionCombineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rosettaConstructionCombineRequest** | [**RosettaConstructionCombineRequest**](RosettaConstructionCombineRequest.md) |  | 

### Return type

[**RosettaConstructionCombineResponse**](RosettaConstructionCombineResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RosettaConstructionDerive

> RosettaConstructionDeriveResponse RosettaConstructionDerive(ctx).RosettaConstructionDeriveRequest(rosettaConstructionDeriveRequest).Execute()

Derive an AccountIdentifier from a PublicKey



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
	rosettaConstructionDeriveRequest := *openapiclient.NewRosettaConstructionDeriveRequest(*openapiclient.NewNetworkIdentifier("Blockchain_example", "Network_example"), *openapiclient.NewRosettaPublicKey("HexBytes_example", "CurveType_example")) // RosettaConstructionDeriveRequest | 

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.RosettaAPI.RosettaConstructionDerive(context.Background()).RosettaConstructionDeriveRequest(rosettaConstructionDeriveRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RosettaAPI.RosettaConstructionDerive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RosettaConstructionDerive`: RosettaConstructionDeriveResponse
	fmt.Fprintf(os.Stdout, "Response from `RosettaAPI.RosettaConstructionDerive`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRosettaConstructionDeriveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rosettaConstructionDeriveRequest** | [**RosettaConstructionDeriveRequest**](RosettaConstructionDeriveRequest.md) |  | 

### Return type

[**RosettaConstructionDeriveResponse**](RosettaConstructionDeriveResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RosettaConstructionHash

> RosettaConstructionHashResponse RosettaConstructionHash(ctx).RosettaConstructionHashRequest(rosettaConstructionHashRequest).Execute()

Get the Hash of a Signed Transaction



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
	rosettaConstructionHashRequest := *openapiclient.NewRosettaConstructionHashRequest(*openapiclient.NewNetworkIdentifier("Blockchain_example", "Network_example"), "SignedTransaction_example") // RosettaConstructionHashRequest | 

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.RosettaAPI.RosettaConstructionHash(context.Background()).RosettaConstructionHashRequest(rosettaConstructionHashRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RosettaAPI.RosettaConstructionHash``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RosettaConstructionHash`: RosettaConstructionHashResponse
	fmt.Fprintf(os.Stdout, "Response from `RosettaAPI.RosettaConstructionHash`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRosettaConstructionHashRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rosettaConstructionHashRequest** | [**RosettaConstructionHashRequest**](RosettaConstructionHashRequest.md) |  | 

### Return type

[**RosettaConstructionHashResponse**](RosettaConstructionHashResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RosettaConstructionMetadata

> RosettaConstructionMetadataResponse RosettaConstructionMetadata(ctx).RosettaConstructionMetadataRequest(rosettaConstructionMetadataRequest).Execute()

Get Metadata for Transaction Construction



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
	rosettaConstructionMetadataRequest := *openapiclient.NewRosettaConstructionMetadataRequest(*openapiclient.NewNetworkIdentifier("Blockchain_example", "Network_example"), *openapiclient.NewRosettaOptions()) // RosettaConstructionMetadataRequest | 

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.RosettaAPI.RosettaConstructionMetadata(context.Background()).RosettaConstructionMetadataRequest(rosettaConstructionMetadataRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RosettaAPI.RosettaConstructionMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RosettaConstructionMetadata`: RosettaConstructionMetadataResponse
	fmt.Fprintf(os.Stdout, "Response from `RosettaAPI.RosettaConstructionMetadata`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRosettaConstructionMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rosettaConstructionMetadataRequest** | [**RosettaConstructionMetadataRequest**](RosettaConstructionMetadataRequest.md) |  | 

### Return type

[**RosettaConstructionMetadataResponse**](RosettaConstructionMetadataResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RosettaConstructionParse

> RosettaConstructionParseResponse RosettaConstructionParse(ctx).RosettaConstructionParseRequest(rosettaConstructionParseRequest).Execute()

Parse a Transaction



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
	rosettaConstructionParseRequest := *openapiclient.NewRosettaConstructionParseRequest(*openapiclient.NewNetworkIdentifier("Blockchain_example", "Network_example"), false, "Transaction_example") // RosettaConstructionParseRequest | 

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.RosettaAPI.RosettaConstructionParse(context.Background()).RosettaConstructionParseRequest(rosettaConstructionParseRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RosettaAPI.RosettaConstructionParse``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RosettaConstructionParse`: RosettaConstructionParseResponse
	fmt.Fprintf(os.Stdout, "Response from `RosettaAPI.RosettaConstructionParse`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRosettaConstructionParseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rosettaConstructionParseRequest** | [**RosettaConstructionParseRequest**](RosettaConstructionParseRequest.md) |  | 

### Return type

[**RosettaConstructionParseResponse**](RosettaConstructionParseResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RosettaConstructionPayloads

> RosettaConstructionPayloadResponse RosettaConstructionPayloads(ctx).RosettaConstructionPayloadsRequest(rosettaConstructionPayloadsRequest).Execute()

Generate an Unsigned Transaction and Signing Payloads



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
	rosettaConstructionPayloadsRequest := *openapiclient.NewRosettaConstructionPayloadsRequest(*openapiclient.NewNetworkIdentifier("Blockchain_example", "Network_example"), []openapiclient.RosettaOperation{*openapiclient.NewRosettaOperation(*openapiclient.NewRosettaOperationIdentifier(int32(123)), "Type_example")}) // RosettaConstructionPayloadsRequest | 

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.RosettaAPI.RosettaConstructionPayloads(context.Background()).RosettaConstructionPayloadsRequest(rosettaConstructionPayloadsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RosettaAPI.RosettaConstructionPayloads``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RosettaConstructionPayloads`: RosettaConstructionPayloadResponse
	fmt.Fprintf(os.Stdout, "Response from `RosettaAPI.RosettaConstructionPayloads`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRosettaConstructionPayloadsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rosettaConstructionPayloadsRequest** | [**RosettaConstructionPayloadsRequest**](RosettaConstructionPayloadsRequest.md) |  | 

### Return type

[**RosettaConstructionPayloadResponse**](RosettaConstructionPayloadResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RosettaConstructionPreprocess

> RosettaConstructionPreprocessResponse RosettaConstructionPreprocess(ctx).RosettaConstructionPreprocessRequest(rosettaConstructionPreprocessRequest).Execute()

Create a Request to Fetch Metadata



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
	rosettaConstructionPreprocessRequest := *openapiclient.NewRosettaConstructionPreprocessRequest(*openapiclient.NewNetworkIdentifier("Blockchain_example", "Network_example"), []openapiclient.RosettaOperation{*openapiclient.NewRosettaOperation(*openapiclient.NewRosettaOperationIdentifier(int32(123)), "Type_example")}) // RosettaConstructionPreprocessRequest | 

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.RosettaAPI.RosettaConstructionPreprocess(context.Background()).RosettaConstructionPreprocessRequest(rosettaConstructionPreprocessRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RosettaAPI.RosettaConstructionPreprocess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RosettaConstructionPreprocess`: RosettaConstructionPreprocessResponse
	fmt.Fprintf(os.Stdout, "Response from `RosettaAPI.RosettaConstructionPreprocess`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRosettaConstructionPreprocessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rosettaConstructionPreprocessRequest** | [**RosettaConstructionPreprocessRequest**](RosettaConstructionPreprocessRequest.md) |  | 

### Return type

[**RosettaConstructionPreprocessResponse**](RosettaConstructionPreprocessResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RosettaConstructionSubmit

> RosettaConstructionSubmitResponse RosettaConstructionSubmit(ctx).RosettaConstructionSubmitRequest(rosettaConstructionSubmitRequest).Execute()

Submit a Signed Transaction



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
	rosettaConstructionSubmitRequest := *openapiclient.NewRosettaConstructionSubmitRequest(*openapiclient.NewNetworkIdentifier("Blockchain_example", "Network_example"), "SignedTransaction_example") // RosettaConstructionSubmitRequest | 

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.RosettaAPI.RosettaConstructionSubmit(context.Background()).RosettaConstructionSubmitRequest(rosettaConstructionSubmitRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RosettaAPI.RosettaConstructionSubmit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RosettaConstructionSubmit`: RosettaConstructionSubmitResponse
	fmt.Fprintf(os.Stdout, "Response from `RosettaAPI.RosettaConstructionSubmit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRosettaConstructionSubmitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rosettaConstructionSubmitRequest** | [**RosettaConstructionSubmitRequest**](RosettaConstructionSubmitRequest.md) |  | 

### Return type

[**RosettaConstructionSubmitResponse**](RosettaConstructionSubmitResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RosettaMempool

> RosettaMempoolResponse RosettaMempool(ctx).RosettaMempoolRequest(rosettaMempoolRequest).Execute()

Get All Mempool Transactions



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
	rosettaMempoolRequest := *openapiclient.NewRosettaMempoolRequest(*openapiclient.NewNetworkIdentifier("Blockchain_example", "Network_example")) // RosettaMempoolRequest | 

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.RosettaAPI.RosettaMempool(context.Background()).RosettaMempoolRequest(rosettaMempoolRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RosettaAPI.RosettaMempool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RosettaMempool`: RosettaMempoolResponse
	fmt.Fprintf(os.Stdout, "Response from `RosettaAPI.RosettaMempool`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRosettaMempoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rosettaMempoolRequest** | [**RosettaMempoolRequest**](RosettaMempoolRequest.md) |  | 

### Return type

[**RosettaMempoolResponse**](RosettaMempoolResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RosettaMempoolTransaction

> RosettaMempoolTransactionResponse RosettaMempoolTransaction(ctx).RosettaMempoolTransactionRequest(rosettaMempoolTransactionRequest).Execute()

Get a Mempool Transaction



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
	rosettaMempoolTransactionRequest := *openapiclient.NewRosettaMempoolTransactionRequest(*openapiclient.NewNetworkIdentifier("Blockchain_example", "Network_example"), *openapiclient.NewTransactionIdentifier("Hash_example")) // RosettaMempoolTransactionRequest | 

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.RosettaAPI.RosettaMempoolTransaction(context.Background()).RosettaMempoolTransactionRequest(rosettaMempoolTransactionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RosettaAPI.RosettaMempoolTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RosettaMempoolTransaction`: RosettaMempoolTransactionResponse
	fmt.Fprintf(os.Stdout, "Response from `RosettaAPI.RosettaMempoolTransaction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRosettaMempoolTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rosettaMempoolTransactionRequest** | [**RosettaMempoolTransactionRequest**](RosettaMempoolTransactionRequest.md) |  | 

### Return type

[**RosettaMempoolTransactionResponse**](RosettaMempoolTransactionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RosettaNetworkList

> RosettaNetworkListResponse RosettaNetworkList(ctx).Execute()

Get List of Available Networks



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
	resp, r, err := apiClient.RosettaAPI.RosettaNetworkList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RosettaAPI.RosettaNetworkList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RosettaNetworkList`: RosettaNetworkListResponse
	fmt.Fprintf(os.Stdout, "Response from `RosettaAPI.RosettaNetworkList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRosettaNetworkListRequest struct via the builder pattern


### Return type

[**RosettaNetworkListResponse**](RosettaNetworkListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RosettaNetworkOptions

> RosettaNetworkOptionsResponse RosettaNetworkOptions(ctx).RosettaOptionsRequest(rosettaOptionsRequest).Execute()

Get Network Options



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
	rosettaOptionsRequest := *openapiclient.NewRosettaOptionsRequest(*openapiclient.NewNetworkIdentifier("Blockchain_example", "Network_example")) // RosettaOptionsRequest | 

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.RosettaAPI.RosettaNetworkOptions(context.Background()).RosettaOptionsRequest(rosettaOptionsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RosettaAPI.RosettaNetworkOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RosettaNetworkOptions`: RosettaNetworkOptionsResponse
	fmt.Fprintf(os.Stdout, "Response from `RosettaAPI.RosettaNetworkOptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRosettaNetworkOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rosettaOptionsRequest** | [**RosettaOptionsRequest**](RosettaOptionsRequest.md) |  | 

### Return type

[**RosettaNetworkOptionsResponse**](RosettaNetworkOptionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RosettaNetworkStatus

> RosettaNetworkStatusResponse RosettaNetworkStatus(ctx).RosettaStatusRequest(rosettaStatusRequest).Execute()

Get Network Status



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
	rosettaStatusRequest := *openapiclient.NewRosettaStatusRequest(*openapiclient.NewNetworkIdentifier("Blockchain_example", "Network_example")) // RosettaStatusRequest | 

	configuration :=config.NewConfiguration()
	apiClient := NewAPIClient(configuration)
	resp, r, err := apiClient.RosettaAPI.RosettaNetworkStatus(context.Background()).RosettaStatusRequest(rosettaStatusRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RosettaAPI.RosettaNetworkStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RosettaNetworkStatus`: RosettaNetworkStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `RosettaAPI.RosettaNetworkStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRosettaNetworkStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rosettaStatusRequest** | [**RosettaStatusRequest**](RosettaStatusRequest.md) |  | 

### Return type

[**RosettaNetworkStatusResponse**](RosettaNetworkStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

