# RosettaRelatedOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | **int32** | Describes the index of related operation. | 
**NetworkIndex** | Pointer to **int32** | Some blockchains specify an operation index that is essential for client use. network_index should not be populated if there is no notion of an operation index in a blockchain (typically most account-based blockchains). | [optional] 

## Methods

### NewRosettaRelatedOperation

`func NewRosettaRelatedOperation(index int32, ) *RosettaRelatedOperation`

NewRosettaRelatedOperation instantiates a new RosettaRelatedOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRosettaRelatedOperationWithDefaults

`func NewRosettaRelatedOperationWithDefaults() *RosettaRelatedOperation`

NewRosettaRelatedOperationWithDefaults instantiates a new RosettaRelatedOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *RosettaRelatedOperation) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *RosettaRelatedOperation) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *RosettaRelatedOperation) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetNetworkIndex

`func (o *RosettaRelatedOperation) GetNetworkIndex() int32`

GetNetworkIndex returns the NetworkIndex field if non-nil, zero value otherwise.

### GetNetworkIndexOk

`func (o *RosettaRelatedOperation) GetNetworkIndexOk() (*int32, bool)`

GetNetworkIndexOk returns a tuple with the NetworkIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIndex

`func (o *RosettaRelatedOperation) SetNetworkIndex(v int32)`

SetNetworkIndex sets NetworkIndex field to given value.

### HasNetworkIndex

`func (o *RosettaRelatedOperation) HasNetworkIndex() bool`

HasNetworkIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


