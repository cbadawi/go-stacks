# RosettaMempoolTransactionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkIdentifier** | [**NetworkIdentifier**](NetworkIdentifier.md) |  | 
**TransactionIdentifier** | [**TransactionIdentifier**](TransactionIdentifier.md) |  | 

## Methods

### NewRosettaMempoolTransactionRequest

`func NewRosettaMempoolTransactionRequest(networkIdentifier NetworkIdentifier, transactionIdentifier TransactionIdentifier, ) *RosettaMempoolTransactionRequest`

NewRosettaMempoolTransactionRequest instantiates a new RosettaMempoolTransactionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRosettaMempoolTransactionRequestWithDefaults

`func NewRosettaMempoolTransactionRequestWithDefaults() *RosettaMempoolTransactionRequest`

NewRosettaMempoolTransactionRequestWithDefaults instantiates a new RosettaMempoolTransactionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkIdentifier

`func (o *RosettaMempoolTransactionRequest) GetNetworkIdentifier() NetworkIdentifier`

GetNetworkIdentifier returns the NetworkIdentifier field if non-nil, zero value otherwise.

### GetNetworkIdentifierOk

`func (o *RosettaMempoolTransactionRequest) GetNetworkIdentifierOk() (*NetworkIdentifier, bool)`

GetNetworkIdentifierOk returns a tuple with the NetworkIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIdentifier

`func (o *RosettaMempoolTransactionRequest) SetNetworkIdentifier(v NetworkIdentifier)`

SetNetworkIdentifier sets NetworkIdentifier field to given value.


### GetTransactionIdentifier

`func (o *RosettaMempoolTransactionRequest) GetTransactionIdentifier() TransactionIdentifier`

GetTransactionIdentifier returns the TransactionIdentifier field if non-nil, zero value otherwise.

### GetTransactionIdentifierOk

`func (o *RosettaMempoolTransactionRequest) GetTransactionIdentifierOk() (*TransactionIdentifier, bool)`

GetTransactionIdentifierOk returns a tuple with the TransactionIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionIdentifier

`func (o *RosettaMempoolTransactionRequest) SetTransactionIdentifier(v TransactionIdentifier)`

SetTransactionIdentifier sets TransactionIdentifier field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


