# RosettaBlockIdentifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hash** | **string** | This is also known as the block hash. | 
**Index** | **int32** | This is also known as the block height. | 

## Methods

### NewRosettaBlockIdentifier

`func NewRosettaBlockIdentifier(hash string, index int32, ) *RosettaBlockIdentifier`

NewRosettaBlockIdentifier instantiates a new RosettaBlockIdentifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRosettaBlockIdentifierWithDefaults

`func NewRosettaBlockIdentifierWithDefaults() *RosettaBlockIdentifier`

NewRosettaBlockIdentifierWithDefaults instantiates a new RosettaBlockIdentifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHash

`func (o *RosettaBlockIdentifier) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *RosettaBlockIdentifier) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *RosettaBlockIdentifier) SetHash(v string)`

SetHash sets Hash field to given value.


### GetIndex

`func (o *RosettaBlockIdentifier) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *RosettaBlockIdentifier) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *RosettaBlockIdentifier) SetIndex(v int32)`

SetIndex sets Index field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


