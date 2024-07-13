# BurnchainRewardListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | **int32** | The number of burnchain rewards to return | 
**Offset** | **int32** | The number to burnchain rewards to skip (starting at &#x60;0&#x60;) | [default to 0]
**Results** | [**[]BurnchainReward**](BurnchainReward.md) |  | 

## Methods

### NewBurnchainRewardListResponse

`func NewBurnchainRewardListResponse(limit int32, offset int32, results []BurnchainReward, ) *BurnchainRewardListResponse`

NewBurnchainRewardListResponse instantiates a new BurnchainRewardListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBurnchainRewardListResponseWithDefaults

`func NewBurnchainRewardListResponseWithDefaults() *BurnchainRewardListResponse`

NewBurnchainRewardListResponseWithDefaults instantiates a new BurnchainRewardListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *BurnchainRewardListResponse) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *BurnchainRewardListResponse) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *BurnchainRewardListResponse) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetOffset

`func (o *BurnchainRewardListResponse) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *BurnchainRewardListResponse) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *BurnchainRewardListResponse) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetResults

`func (o *BurnchainRewardListResponse) GetResults() []BurnchainReward`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *BurnchainRewardListResponse) GetResultsOk() (*[]BurnchainReward, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *BurnchainRewardListResponse) SetResults(v []BurnchainReward)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


