# RosettaOperationIdentifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | **int32** | The operation index is used to ensure each operation has a unique identifier within a transaction. This index is only relative to the transaction and NOT GLOBAL. The operations in each transaction should start from index 0. To clarify, there may not be any notion of an operation index in the blockchain being described. | 
**NetworkIndex** | Pointer to **int32** | Some blockchains specify an operation index that is essential for client use. For example, Bitcoin uses a network_index to identify which UTXO was used in a transaction. network_index should not be populated if there is no notion of an operation index in a blockchain (typically most account-based blockchains). | [optional] 

## Methods

### NewRosettaOperationIdentifier

`func NewRosettaOperationIdentifier(index int32, ) *RosettaOperationIdentifier`

NewRosettaOperationIdentifier instantiates a new RosettaOperationIdentifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRosettaOperationIdentifierWithDefaults

`func NewRosettaOperationIdentifierWithDefaults() *RosettaOperationIdentifier`

NewRosettaOperationIdentifierWithDefaults instantiates a new RosettaOperationIdentifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *RosettaOperationIdentifier) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *RosettaOperationIdentifier) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *RosettaOperationIdentifier) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetNetworkIndex

`func (o *RosettaOperationIdentifier) GetNetworkIndex() int32`

GetNetworkIndex returns the NetworkIndex field if non-nil, zero value otherwise.

### GetNetworkIndexOk

`func (o *RosettaOperationIdentifier) GetNetworkIndexOk() (*int32, bool)`

GetNetworkIndexOk returns a tuple with the NetworkIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIndex

`func (o *RosettaOperationIdentifier) SetNetworkIndex(v int32)`

SetNetworkIndex sets NetworkIndex field to given value.

### HasNetworkIndex

`func (o *RosettaOperationIdentifier) HasNetworkIndex() bool`

HasNetworkIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


