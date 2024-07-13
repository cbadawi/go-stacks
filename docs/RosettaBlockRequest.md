# RosettaBlockRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkIdentifier** | [**NetworkIdentifier**](NetworkIdentifier.md) |  | 
**BlockIdentifier** | [**RosettaPartialBlockIdentifier**](RosettaPartialBlockIdentifier.md) |  | 

## Methods

### NewRosettaBlockRequest

`func NewRosettaBlockRequest(networkIdentifier NetworkIdentifier, blockIdentifier RosettaPartialBlockIdentifier, ) *RosettaBlockRequest`

NewRosettaBlockRequest instantiates a new RosettaBlockRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRosettaBlockRequestWithDefaults

`func NewRosettaBlockRequestWithDefaults() *RosettaBlockRequest`

NewRosettaBlockRequestWithDefaults instantiates a new RosettaBlockRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkIdentifier

`func (o *RosettaBlockRequest) GetNetworkIdentifier() NetworkIdentifier`

GetNetworkIdentifier returns the NetworkIdentifier field if non-nil, zero value otherwise.

### GetNetworkIdentifierOk

`func (o *RosettaBlockRequest) GetNetworkIdentifierOk() (*NetworkIdentifier, bool)`

GetNetworkIdentifierOk returns a tuple with the NetworkIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIdentifier

`func (o *RosettaBlockRequest) SetNetworkIdentifier(v NetworkIdentifier)`

SetNetworkIdentifier sets NetworkIdentifier field to given value.


### GetBlockIdentifier

`func (o *RosettaBlockRequest) GetBlockIdentifier() RosettaPartialBlockIdentifier`

GetBlockIdentifier returns the BlockIdentifier field if non-nil, zero value otherwise.

### GetBlockIdentifierOk

`func (o *RosettaBlockRequest) GetBlockIdentifierOk() (*RosettaPartialBlockIdentifier, bool)`

GetBlockIdentifierOk returns a tuple with the BlockIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockIdentifier

`func (o *RosettaBlockRequest) SetBlockIdentifier(v RosettaPartialBlockIdentifier)`

SetBlockIdentifier sets BlockIdentifier field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


