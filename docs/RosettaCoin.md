# RosettaCoin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CoinIdentifier** | [**RosettaCoinCoinIdentifier**](RosettaCoinCoinIdentifier.md) |  | 
**Amount** | [**RosettaAmount**](RosettaAmount.md) |  | 

## Methods

### NewRosettaCoin

`func NewRosettaCoin(coinIdentifier RosettaCoinCoinIdentifier, amount RosettaAmount, ) *RosettaCoin`

NewRosettaCoin instantiates a new RosettaCoin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRosettaCoinWithDefaults

`func NewRosettaCoinWithDefaults() *RosettaCoin`

NewRosettaCoinWithDefaults instantiates a new RosettaCoin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoinIdentifier

`func (o *RosettaCoin) GetCoinIdentifier() RosettaCoinCoinIdentifier`

GetCoinIdentifier returns the CoinIdentifier field if non-nil, zero value otherwise.

### GetCoinIdentifierOk

`func (o *RosettaCoin) GetCoinIdentifierOk() (*RosettaCoinCoinIdentifier, bool)`

GetCoinIdentifierOk returns a tuple with the CoinIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoinIdentifier

`func (o *RosettaCoin) SetCoinIdentifier(v RosettaCoinCoinIdentifier)`

SetCoinIdentifier sets CoinIdentifier field to given value.


### GetAmount

`func (o *RosettaCoin) GetAmount() RosettaAmount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *RosettaCoin) GetAmountOk() (*RosettaAmount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *RosettaCoin) SetAmount(v RosettaAmount)`

SetAmount sets Amount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


