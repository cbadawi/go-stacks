# TransactionList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Found** | **bool** |  | 
**Result** | [**TransactionNotFoundResult**](TransactionNotFoundResult.md) |  | 

## Methods

### NewTransactionList

`func NewTransactionList(found bool, result TransactionNotFoundResult, ) *TransactionList`

NewTransactionList instantiates a new TransactionList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionListWithDefaults

`func NewTransactionListWithDefaults() *TransactionList`

NewTransactionListWithDefaults instantiates a new TransactionList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFound

`func (o *TransactionList) GetFound() bool`

GetFound returns the Found field if non-nil, zero value otherwise.

### GetFoundOk

`func (o *TransactionList) GetFoundOk() (*bool, bool)`

GetFoundOk returns a tuple with the Found field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFound

`func (o *TransactionList) SetFound(v bool)`

SetFound sets Found field to given value.


### GetResult

`func (o *TransactionList) GetResult() TransactionNotFoundResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *TransactionList) GetResultOk() (*TransactionNotFoundResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *TransactionList) SetResult(v TransactionNotFoundResult)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


