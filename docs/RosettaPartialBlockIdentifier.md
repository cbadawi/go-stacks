# RosettaPartialBlockIdentifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hash** | **string** | This is also known as the block hash. | 
**Index** | **int32** | This is also known as the block height. | 

## Methods

### NewRosettaPartialBlockIdentifier

`func NewRosettaPartialBlockIdentifier(hash string, index int32, ) *RosettaPartialBlockIdentifier`

NewRosettaPartialBlockIdentifier instantiates a new RosettaPartialBlockIdentifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRosettaPartialBlockIdentifierWithDefaults

`func NewRosettaPartialBlockIdentifierWithDefaults() *RosettaPartialBlockIdentifier`

NewRosettaPartialBlockIdentifierWithDefaults instantiates a new RosettaPartialBlockIdentifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHash

`func (o *RosettaPartialBlockIdentifier) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *RosettaPartialBlockIdentifier) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *RosettaPartialBlockIdentifier) SetHash(v string)`

SetHash sets Hash field to given value.


### GetIndex

`func (o *RosettaPartialBlockIdentifier) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *RosettaPartialBlockIdentifier) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *RosettaPartialBlockIdentifier) SetIndex(v int32)`

SetIndex sets Index field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


