# NakamotoBlockListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | **int32** | The number of blocks to return | 
**Offset** | **int32** | The number to blocks to skip (starting at &#x60;0&#x60;) | [default to 0]
**Total** | **int32** | The number of blocks available | 
**Results** | [**[]NakamotoBlock**](NakamotoBlock.md) |  | 

## Methods

### NewNakamotoBlockListResponse

`func NewNakamotoBlockListResponse(limit int32, offset int32, total int32, results []NakamotoBlock, ) *NakamotoBlockListResponse`

NewNakamotoBlockListResponse instantiates a new NakamotoBlockListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNakamotoBlockListResponseWithDefaults

`func NewNakamotoBlockListResponseWithDefaults() *NakamotoBlockListResponse`

NewNakamotoBlockListResponseWithDefaults instantiates a new NakamotoBlockListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *NakamotoBlockListResponse) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *NakamotoBlockListResponse) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *NakamotoBlockListResponse) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetOffset

`func (o *NakamotoBlockListResponse) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *NakamotoBlockListResponse) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *NakamotoBlockListResponse) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetTotal

`func (o *NakamotoBlockListResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *NakamotoBlockListResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *NakamotoBlockListResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetResults

`func (o *NakamotoBlockListResponse) GetResults() []NakamotoBlock`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *NakamotoBlockListResponse) GetResultsOk() (*[]NakamotoBlock, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *NakamotoBlockListResponse) SetResults(v []NakamotoBlock)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


