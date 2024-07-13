# BlockSearchResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Found** | **bool** | Indicates if the requested object was found or not | [default to true]
**Result** | [**BlockSearchResultResult**](BlockSearchResultResult.md) |  | 

## Methods

### NewBlockSearchResult

`func NewBlockSearchResult(found bool, result BlockSearchResultResult, ) *BlockSearchResult`

NewBlockSearchResult instantiates a new BlockSearchResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockSearchResultWithDefaults

`func NewBlockSearchResultWithDefaults() *BlockSearchResult`

NewBlockSearchResultWithDefaults instantiates a new BlockSearchResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFound

`func (o *BlockSearchResult) GetFound() bool`

GetFound returns the Found field if non-nil, zero value otherwise.

### GetFoundOk

`func (o *BlockSearchResult) GetFoundOk() (*bool, bool)`

GetFoundOk returns a tuple with the Found field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFound

`func (o *BlockSearchResult) SetFound(v bool)`

SetFound sets Found field to given value.


### GetResult

`func (o *BlockSearchResult) GetResult() BlockSearchResultResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *BlockSearchResult) GetResultOk() (*BlockSearchResultResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *BlockSearchResult) SetResult(v BlockSearchResultResult)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


