# SearchSuccessResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Found** | **bool** | Indicates if the requested object was found or not | [default to true]
**Result** | [**TxSearchResultResult**](TxSearchResultResult.md) |  | 

## Methods

### NewSearchSuccessResult

`func NewSearchSuccessResult(found bool, result TxSearchResultResult, ) *SearchSuccessResult`

NewSearchSuccessResult instantiates a new SearchSuccessResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchSuccessResultWithDefaults

`func NewSearchSuccessResultWithDefaults() *SearchSuccessResult`

NewSearchSuccessResultWithDefaults instantiates a new SearchSuccessResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFound

`func (o *SearchSuccessResult) GetFound() bool`

GetFound returns the Found field if non-nil, zero value otherwise.

### GetFoundOk

`func (o *SearchSuccessResult) GetFoundOk() (*bool, bool)`

GetFoundOk returns a tuple with the Found field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFound

`func (o *SearchSuccessResult) SetFound(v bool)`

SetFound sets Found field to given value.


### GetResult

`func (o *SearchSuccessResult) GetResult() TxSearchResultResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *SearchSuccessResult) GetResultOk() (*TxSearchResultResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *SearchSuccessResult) SetResult(v TxSearchResultResult)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


