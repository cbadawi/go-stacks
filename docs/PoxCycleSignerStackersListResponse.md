# PoxCycleSignerStackersListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | **int32** | The number of stackers to return | 
**Offset** | **int32** | The number to stackers to skip (starting at &#x60;0&#x60;) | [default to 0]
**Total** | **int32** | The total number of stackers | 
**Results** | [**[]PoxStacker**](PoxStacker.md) |  | 

## Methods

### NewPoxCycleSignerStackersListResponse

`func NewPoxCycleSignerStackersListResponse(limit int32, offset int32, total int32, results []PoxStacker, ) *PoxCycleSignerStackersListResponse`

NewPoxCycleSignerStackersListResponse instantiates a new PoxCycleSignerStackersListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoxCycleSignerStackersListResponseWithDefaults

`func NewPoxCycleSignerStackersListResponseWithDefaults() *PoxCycleSignerStackersListResponse`

NewPoxCycleSignerStackersListResponseWithDefaults instantiates a new PoxCycleSignerStackersListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *PoxCycleSignerStackersListResponse) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *PoxCycleSignerStackersListResponse) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *PoxCycleSignerStackersListResponse) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetOffset

`func (o *PoxCycleSignerStackersListResponse) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *PoxCycleSignerStackersListResponse) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *PoxCycleSignerStackersListResponse) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetTotal

`func (o *PoxCycleSignerStackersListResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PoxCycleSignerStackersListResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PoxCycleSignerStackersListResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetResults

`func (o *PoxCycleSignerStackersListResponse) GetResults() []PoxStacker`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PoxCycleSignerStackersListResponse) GetResultsOk() (*[]PoxStacker, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PoxCycleSignerStackersListResponse) SetResults(v []PoxStacker)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


