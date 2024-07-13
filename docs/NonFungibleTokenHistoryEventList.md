# NonFungibleTokenHistoryEventList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | **int32** | The number of events to return | 
**Offset** | **int32** | The number to events to skip (starting at &#x60;0&#x60;) | 
**Total** | **int32** | The number of events available | 
**Results** | [**[]NonFungibleTokenHistoryEvent**](NonFungibleTokenHistoryEvent.md) |  | 

## Methods

### NewNonFungibleTokenHistoryEventList

`func NewNonFungibleTokenHistoryEventList(limit int32, offset int32, total int32, results []NonFungibleTokenHistoryEvent, ) *NonFungibleTokenHistoryEventList`

NewNonFungibleTokenHistoryEventList instantiates a new NonFungibleTokenHistoryEventList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNonFungibleTokenHistoryEventListWithDefaults

`func NewNonFungibleTokenHistoryEventListWithDefaults() *NonFungibleTokenHistoryEventList`

NewNonFungibleTokenHistoryEventListWithDefaults instantiates a new NonFungibleTokenHistoryEventList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *NonFungibleTokenHistoryEventList) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *NonFungibleTokenHistoryEventList) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *NonFungibleTokenHistoryEventList) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetOffset

`func (o *NonFungibleTokenHistoryEventList) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *NonFungibleTokenHistoryEventList) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *NonFungibleTokenHistoryEventList) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetTotal

`func (o *NonFungibleTokenHistoryEventList) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *NonFungibleTokenHistoryEventList) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *NonFungibleTokenHistoryEventList) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetResults

`func (o *NonFungibleTokenHistoryEventList) GetResults() []NonFungibleTokenHistoryEvent`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *NonFungibleTokenHistoryEventList) GetResultsOk() (*[]NonFungibleTokenHistoryEvent, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *NonFungibleTokenHistoryEventList) SetResults(v []NonFungibleTokenHistoryEvent)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


