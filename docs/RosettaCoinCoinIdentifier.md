# RosettaCoinCoinIdentifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | **string** | Identifier should be populated with a globally unique identifier of a Coin. In Bitcoin, this identifier would be transaction_hash:index. | 

## Methods

### NewRosettaCoinCoinIdentifier

`func NewRosettaCoinCoinIdentifier(identifier string, ) *RosettaCoinCoinIdentifier`

NewRosettaCoinCoinIdentifier instantiates a new RosettaCoinCoinIdentifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRosettaCoinCoinIdentifierWithDefaults

`func NewRosettaCoinCoinIdentifierWithDefaults() *RosettaCoinCoinIdentifier`

NewRosettaCoinCoinIdentifierWithDefaults instantiates a new RosettaCoinCoinIdentifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *RosettaCoinCoinIdentifier) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *RosettaCoinCoinIdentifier) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *RosettaCoinCoinIdentifier) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


