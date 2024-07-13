# NonFungibleTokenHoldingsList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | **int32** | The number of Non-Fungible Token holdings to return | 
**Offset** | **int32** | The number to Non-Fungible Token holdings to skip (starting at &#x60;0&#x60;) | 
**Total** | **int32** | The number of Non-Fungible Token holdings available | 
**Results** | [**[]NonFungibleTokenHolding**](NonFungibleTokenHolding.md) |  | 

## Methods

### NewNonFungibleTokenHoldingsList

`func NewNonFungibleTokenHoldingsList(limit int32, offset int32, total int32, results []NonFungibleTokenHolding, ) *NonFungibleTokenHoldingsList`

NewNonFungibleTokenHoldingsList instantiates a new NonFungibleTokenHoldingsList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNonFungibleTokenHoldingsListWithDefaults

`func NewNonFungibleTokenHoldingsListWithDefaults() *NonFungibleTokenHoldingsList`

NewNonFungibleTokenHoldingsListWithDefaults instantiates a new NonFungibleTokenHoldingsList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *NonFungibleTokenHoldingsList) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *NonFungibleTokenHoldingsList) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *NonFungibleTokenHoldingsList) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetOffset

`func (o *NonFungibleTokenHoldingsList) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *NonFungibleTokenHoldingsList) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *NonFungibleTokenHoldingsList) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetTotal

`func (o *NonFungibleTokenHoldingsList) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *NonFungibleTokenHoldingsList) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *NonFungibleTokenHoldingsList) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetResults

`func (o *NonFungibleTokenHoldingsList) GetResults() []NonFungibleTokenHolding`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *NonFungibleTokenHoldingsList) GetResultsOk() (*[]NonFungibleTokenHolding, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *NonFungibleTokenHoldingsList) SetResults(v []NonFungibleTokenHolding)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


