# PostCoreNodeTransactionsError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | **string** | The error | 
**Reason** | **string** | The reason for the error | 
**ReasonData** | **map[string]interface{}** | More details about the reason | 
**Txid** | **string** | The relevant transaction id | 

## Methods

### NewPostCoreNodeTransactionsError

`func NewPostCoreNodeTransactionsError(error_ string, reason string, reasonData map[string]interface{}, txid string, ) *PostCoreNodeTransactionsError`

NewPostCoreNodeTransactionsError instantiates a new PostCoreNodeTransactionsError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostCoreNodeTransactionsErrorWithDefaults

`func NewPostCoreNodeTransactionsErrorWithDefaults() *PostCoreNodeTransactionsError`

NewPostCoreNodeTransactionsErrorWithDefaults instantiates a new PostCoreNodeTransactionsError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *PostCoreNodeTransactionsError) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *PostCoreNodeTransactionsError) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *PostCoreNodeTransactionsError) SetError(v string)`

SetError sets Error field to given value.


### GetReason

`func (o *PostCoreNodeTransactionsError) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *PostCoreNodeTransactionsError) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *PostCoreNodeTransactionsError) SetReason(v string)`

SetReason sets Reason field to given value.


### GetReasonData

`func (o *PostCoreNodeTransactionsError) GetReasonData() map[string]interface{}`

GetReasonData returns the ReasonData field if non-nil, zero value otherwise.

### GetReasonDataOk

`func (o *PostCoreNodeTransactionsError) GetReasonDataOk() (*map[string]interface{}, bool)`

GetReasonDataOk returns a tuple with the ReasonData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonData

`func (o *PostCoreNodeTransactionsError) SetReasonData(v map[string]interface{})`

SetReasonData sets ReasonData field to given value.


### GetTxid

`func (o *PostCoreNodeTransactionsError) GetTxid() string`

GetTxid returns the Txid field if non-nil, zero value otherwise.

### GetTxidOk

`func (o *PostCoreNodeTransactionsError) GetTxidOk() (*string, bool)`

GetTxidOk returns a tuple with the Txid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxid

`func (o *PostCoreNodeTransactionsError) SetTxid(v string)`

SetTxid sets Txid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


