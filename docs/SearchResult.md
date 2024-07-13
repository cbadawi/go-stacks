# SearchResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Found** | **bool** | Indicates if the requested object was found or not | [default to true]
**Result** | [**TxSearchResultResult**](TxSearchResultResult.md) |  | 
**Error** | **string** |  | 

## Methods

### NewSearchResult

`func NewSearchResult(found bool, result TxSearchResultResult, error_ string, ) *SearchResult`

NewSearchResult instantiates a new SearchResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchResultWithDefaults

`func NewSearchResultWithDefaults() *SearchResult`

NewSearchResultWithDefaults instantiates a new SearchResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFound

`func (o *SearchResult) GetFound() bool`

GetFound returns the Found field if non-nil, zero value otherwise.

### GetFoundOk

`func (o *SearchResult) GetFoundOk() (*bool, bool)`

GetFoundOk returns a tuple with the Found field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFound

`func (o *SearchResult) SetFound(v bool)`

SetFound sets Found field to given value.


### GetResult

`func (o *SearchResult) GetResult() TxSearchResultResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *SearchResult) GetResultOk() (*TxSearchResultResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *SearchResult) SetResult(v TxSearchResultResult)`

SetResult sets Result field to given value.


### GetError

`func (o *SearchResult) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SearchResult) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SearchResult) SetError(v string)`

SetError sets Error field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


