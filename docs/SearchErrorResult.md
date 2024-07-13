# SearchErrorResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Found** | **bool** | Indicates if the requested object was found or not | [default to false]
**Result** | [**SearchErrorResultResult**](SearchErrorResultResult.md) |  | 
**Error** | **string** |  | 

## Methods

### NewSearchErrorResult

`func NewSearchErrorResult(found bool, result SearchErrorResultResult, error_ string, ) *SearchErrorResult`

NewSearchErrorResult instantiates a new SearchErrorResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchErrorResultWithDefaults

`func NewSearchErrorResultWithDefaults() *SearchErrorResult`

NewSearchErrorResultWithDefaults instantiates a new SearchErrorResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFound

`func (o *SearchErrorResult) GetFound() bool`

GetFound returns the Found field if non-nil, zero value otherwise.

### GetFoundOk

`func (o *SearchErrorResult) GetFoundOk() (*bool, bool)`

GetFoundOk returns a tuple with the Found field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFound

`func (o *SearchErrorResult) SetFound(v bool)`

SetFound sets Found field to given value.


### GetResult

`func (o *SearchErrorResult) GetResult() SearchErrorResultResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *SearchErrorResult) GetResultOk() (*SearchErrorResultResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *SearchErrorResult) SetResult(v SearchErrorResultResult)`

SetResult sets Result field to given value.


### GetError

`func (o *SearchErrorResult) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SearchErrorResult) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SearchErrorResult) SetError(v string)`

SetError sets Error field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


