# PoolDelegationsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | **int32** | The number of Stackers to return | 
**Offset** | **int32** | The number to Stackers to skip (starting at &#x60;0&#x60;) | [default to 0]
**Total** | **int32** | The total number of Stackers | 
**Results** | [**[]PoolDelegation**](PoolDelegation.md) |  | 

## Methods

### NewPoolDelegationsResponse

`func NewPoolDelegationsResponse(limit int32, offset int32, total int32, results []PoolDelegation, ) *PoolDelegationsResponse`

NewPoolDelegationsResponse instantiates a new PoolDelegationsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolDelegationsResponseWithDefaults

`func NewPoolDelegationsResponseWithDefaults() *PoolDelegationsResponse`

NewPoolDelegationsResponseWithDefaults instantiates a new PoolDelegationsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *PoolDelegationsResponse) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *PoolDelegationsResponse) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *PoolDelegationsResponse) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetOffset

`func (o *PoolDelegationsResponse) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *PoolDelegationsResponse) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *PoolDelegationsResponse) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetTotal

`func (o *PoolDelegationsResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PoolDelegationsResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PoolDelegationsResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetResults

`func (o *PoolDelegationsResponse) GetResults() []PoolDelegation`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PoolDelegationsResponse) GetResultsOk() (*[]PoolDelegation, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PoolDelegationsResponse) SetResults(v []PoolDelegation)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


