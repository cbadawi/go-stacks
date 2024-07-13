# RosettaCoinChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CoinIdentifier** | [**RosettaCoinChangeCoinIdentifier**](RosettaCoinChangeCoinIdentifier.md) |  | 
**CoinAction** | **string** | CoinActions are different state changes that a Coin can undergo. When a Coin is created, it is coin_created. When a Coin is spent, it is coin_spent. It is assumed that a single Coin cannot be created or spent more than once. | 

## Methods

### NewRosettaCoinChange

`func NewRosettaCoinChange(coinIdentifier RosettaCoinChangeCoinIdentifier, coinAction string, ) *RosettaCoinChange`

NewRosettaCoinChange instantiates a new RosettaCoinChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRosettaCoinChangeWithDefaults

`func NewRosettaCoinChangeWithDefaults() *RosettaCoinChange`

NewRosettaCoinChangeWithDefaults instantiates a new RosettaCoinChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoinIdentifier

`func (o *RosettaCoinChange) GetCoinIdentifier() RosettaCoinChangeCoinIdentifier`

GetCoinIdentifier returns the CoinIdentifier field if non-nil, zero value otherwise.

### GetCoinIdentifierOk

`func (o *RosettaCoinChange) GetCoinIdentifierOk() (*RosettaCoinChangeCoinIdentifier, bool)`

GetCoinIdentifierOk returns a tuple with the CoinIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoinIdentifier

`func (o *RosettaCoinChange) SetCoinIdentifier(v RosettaCoinChangeCoinIdentifier)`

SetCoinIdentifier sets CoinIdentifier field to given value.


### GetCoinAction

`func (o *RosettaCoinChange) GetCoinAction() string`

GetCoinAction returns the CoinAction field if non-nil, zero value otherwise.

### GetCoinActionOk

`func (o *RosettaCoinChange) GetCoinActionOk() (*string, bool)`

GetCoinActionOk returns a tuple with the CoinAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoinAction

`func (o *RosettaCoinChange) SetCoinAction(v string)`

SetCoinAction sets CoinAction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


