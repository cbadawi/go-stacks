# RosettaMaxFeeAmount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **string** | Value of the transaction in atomic units represented as an arbitrary-sized signed integer. For example, 1 BTC would be represented by a value of 100000000. | 
**Currency** | [**RosettaCurrency**](RosettaCurrency.md) |  | 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewRosettaMaxFeeAmount

`func NewRosettaMaxFeeAmount(value string, currency RosettaCurrency, ) *RosettaMaxFeeAmount`

NewRosettaMaxFeeAmount instantiates a new RosettaMaxFeeAmount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRosettaMaxFeeAmountWithDefaults

`func NewRosettaMaxFeeAmountWithDefaults() *RosettaMaxFeeAmount`

NewRosettaMaxFeeAmountWithDefaults instantiates a new RosettaMaxFeeAmount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *RosettaMaxFeeAmount) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RosettaMaxFeeAmount) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RosettaMaxFeeAmount) SetValue(v string)`

SetValue sets Value field to given value.


### GetCurrency

`func (o *RosettaMaxFeeAmount) GetCurrency() RosettaCurrency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *RosettaMaxFeeAmount) GetCurrencyOk() (*RosettaCurrency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *RosettaMaxFeeAmount) SetCurrency(v RosettaCurrency)`

SetCurrency sets Currency field to given value.


### GetMetadata

`func (o *RosettaMaxFeeAmount) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RosettaMaxFeeAmount) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RosettaMaxFeeAmount) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RosettaMaxFeeAmount) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


