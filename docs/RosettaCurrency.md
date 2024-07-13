# RosettaCurrency

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** | Canonical symbol associated with a currency. | 
**Decimals** | **int32** | Number of decimal places in the standard unit representation of the amount. For example, BTC has 8 decimals. Note that it is not possible to represent the value of some currency in atomic units that is not base 10. | 
**Metadata** | Pointer to **map[string]interface{}** | Any additional information related to the currency itself. For example, it would be useful to populate this object with the contract address of an ERC-20 token. | [optional] 

## Methods

### NewRosettaCurrency

`func NewRosettaCurrency(symbol string, decimals int32, ) *RosettaCurrency`

NewRosettaCurrency instantiates a new RosettaCurrency object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRosettaCurrencyWithDefaults

`func NewRosettaCurrencyWithDefaults() *RosettaCurrency`

NewRosettaCurrencyWithDefaults instantiates a new RosettaCurrency object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *RosettaCurrency) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *RosettaCurrency) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *RosettaCurrency) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetDecimals

`func (o *RosettaCurrency) GetDecimals() int32`

GetDecimals returns the Decimals field if non-nil, zero value otherwise.

### GetDecimalsOk

`func (o *RosettaCurrency) GetDecimalsOk() (*int32, bool)`

GetDecimalsOk returns a tuple with the Decimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimals

`func (o *RosettaCurrency) SetDecimals(v int32)`

SetDecimals sets Decimals field to given value.


### GetMetadata

`func (o *RosettaCurrency) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RosettaCurrency) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RosettaCurrency) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RosettaCurrency) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


