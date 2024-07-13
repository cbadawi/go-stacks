# ContractSearchResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Found** | **bool** | Indicates if the requested object was found or not | [default to true]
**Result** | [**ContractSearchResultResult**](ContractSearchResultResult.md) |  | 

## Methods

### NewContractSearchResult

`func NewContractSearchResult(found bool, result ContractSearchResultResult, ) *ContractSearchResult`

NewContractSearchResult instantiates a new ContractSearchResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractSearchResultWithDefaults

`func NewContractSearchResultWithDefaults() *ContractSearchResult`

NewContractSearchResultWithDefaults instantiates a new ContractSearchResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFound

`func (o *ContractSearchResult) GetFound() bool`

GetFound returns the Found field if non-nil, zero value otherwise.

### GetFoundOk

`func (o *ContractSearchResult) GetFoundOk() (*bool, bool)`

GetFoundOk returns a tuple with the Found field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFound

`func (o *ContractSearchResult) SetFound(v bool)`

SetFound sets Found field to given value.


### GetResult

`func (o *ContractSearchResult) GetResult() ContractSearchResultResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ContractSearchResult) GetResultOk() (*ContractSearchResultResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ContractSearchResult) SetResult(v ContractSearchResultResult)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


