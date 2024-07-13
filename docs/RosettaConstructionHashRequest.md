# RosettaConstructionHashRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkIdentifier** | [**NetworkIdentifier**](NetworkIdentifier.md) |  | 
**SignedTransaction** | **string** | Signed transaction | 

## Methods

### NewRosettaConstructionHashRequest

`func NewRosettaConstructionHashRequest(networkIdentifier NetworkIdentifier, signedTransaction string, ) *RosettaConstructionHashRequest`

NewRosettaConstructionHashRequest instantiates a new RosettaConstructionHashRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRosettaConstructionHashRequestWithDefaults

`func NewRosettaConstructionHashRequestWithDefaults() *RosettaConstructionHashRequest`

NewRosettaConstructionHashRequestWithDefaults instantiates a new RosettaConstructionHashRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkIdentifier

`func (o *RosettaConstructionHashRequest) GetNetworkIdentifier() NetworkIdentifier`

GetNetworkIdentifier returns the NetworkIdentifier field if non-nil, zero value otherwise.

### GetNetworkIdentifierOk

`func (o *RosettaConstructionHashRequest) GetNetworkIdentifierOk() (*NetworkIdentifier, bool)`

GetNetworkIdentifierOk returns a tuple with the NetworkIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIdentifier

`func (o *RosettaConstructionHashRequest) SetNetworkIdentifier(v NetworkIdentifier)`

SetNetworkIdentifier sets NetworkIdentifier field to given value.


### GetSignedTransaction

`func (o *RosettaConstructionHashRequest) GetSignedTransaction() string`

GetSignedTransaction returns the SignedTransaction field if non-nil, zero value otherwise.

### GetSignedTransactionOk

`func (o *RosettaConstructionHashRequest) GetSignedTransactionOk() (*string, bool)`

GetSignedTransactionOk returns a tuple with the SignedTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedTransaction

`func (o *RosettaConstructionHashRequest) SetSignedTransaction(v string)`

SetSignedTransaction sets SignedTransaction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


