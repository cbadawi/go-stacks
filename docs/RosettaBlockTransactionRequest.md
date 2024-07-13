# RosettaBlockTransactionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkIdentifier** | [**NetworkIdentifier**](NetworkIdentifier.md) |  | 
**BlockIdentifier** | [**RosettaPartialBlockIdentifier**](RosettaPartialBlockIdentifier.md) |  | 
**TransactionIdentifier** | [**TransactionIdentifier**](TransactionIdentifier.md) |  | 

## Methods

### NewRosettaBlockTransactionRequest

`func NewRosettaBlockTransactionRequest(networkIdentifier NetworkIdentifier, blockIdentifier RosettaPartialBlockIdentifier, transactionIdentifier TransactionIdentifier, ) *RosettaBlockTransactionRequest`

NewRosettaBlockTransactionRequest instantiates a new RosettaBlockTransactionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRosettaBlockTransactionRequestWithDefaults

`func NewRosettaBlockTransactionRequestWithDefaults() *RosettaBlockTransactionRequest`

NewRosettaBlockTransactionRequestWithDefaults instantiates a new RosettaBlockTransactionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkIdentifier

`func (o *RosettaBlockTransactionRequest) GetNetworkIdentifier() NetworkIdentifier`

GetNetworkIdentifier returns the NetworkIdentifier field if non-nil, zero value otherwise.

### GetNetworkIdentifierOk

`func (o *RosettaBlockTransactionRequest) GetNetworkIdentifierOk() (*NetworkIdentifier, bool)`

GetNetworkIdentifierOk returns a tuple with the NetworkIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIdentifier

`func (o *RosettaBlockTransactionRequest) SetNetworkIdentifier(v NetworkIdentifier)`

SetNetworkIdentifier sets NetworkIdentifier field to given value.


### GetBlockIdentifier

`func (o *RosettaBlockTransactionRequest) GetBlockIdentifier() RosettaPartialBlockIdentifier`

GetBlockIdentifier returns the BlockIdentifier field if non-nil, zero value otherwise.

### GetBlockIdentifierOk

`func (o *RosettaBlockTransactionRequest) GetBlockIdentifierOk() (*RosettaPartialBlockIdentifier, bool)`

GetBlockIdentifierOk returns a tuple with the BlockIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockIdentifier

`func (o *RosettaBlockTransactionRequest) SetBlockIdentifier(v RosettaPartialBlockIdentifier)`

SetBlockIdentifier sets BlockIdentifier field to given value.


### GetTransactionIdentifier

`func (o *RosettaBlockTransactionRequest) GetTransactionIdentifier() TransactionIdentifier`

GetTransactionIdentifier returns the TransactionIdentifier field if non-nil, zero value otherwise.

### GetTransactionIdentifierOk

`func (o *RosettaBlockTransactionRequest) GetTransactionIdentifierOk() (*TransactionIdentifier, bool)`

GetTransactionIdentifierOk returns a tuple with the TransactionIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionIdentifier

`func (o *RosettaBlockTransactionRequest) SetTransactionIdentifier(v TransactionIdentifier)`

SetTransactionIdentifier sets TransactionIdentifier field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


