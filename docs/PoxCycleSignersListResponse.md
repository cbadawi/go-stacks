# PoxCycleSignersListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | **int32** | The number of signers to return | 
**Offset** | **int32** | The number to signers to skip (starting at &#x60;0&#x60;) | [default to 0]
**Total** | **int32** | The total number of signers | 
**Results** | [**[]PoxSigner**](PoxSigner.md) |  | 

## Methods

### NewPoxCycleSignersListResponse

`func NewPoxCycleSignersListResponse(limit int32, offset int32, total int32, results []PoxSigner, ) *PoxCycleSignersListResponse`

NewPoxCycleSignersListResponse instantiates a new PoxCycleSignersListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoxCycleSignersListResponseWithDefaults

`func NewPoxCycleSignersListResponseWithDefaults() *PoxCycleSignersListResponse`

NewPoxCycleSignersListResponseWithDefaults instantiates a new PoxCycleSignersListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *PoxCycleSignersListResponse) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *PoxCycleSignersListResponse) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *PoxCycleSignersListResponse) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetOffset

`func (o *PoxCycleSignersListResponse) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *PoxCycleSignersListResponse) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *PoxCycleSignersListResponse) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetTotal

`func (o *PoxCycleSignersListResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PoxCycleSignersListResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PoxCycleSignersListResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetResults

`func (o *PoxCycleSignersListResponse) GetResults() []PoxSigner`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PoxCycleSignersListResponse) GetResultsOk() (*[]PoxSigner, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PoxCycleSignersListResponse) SetResults(v []PoxSigner)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


