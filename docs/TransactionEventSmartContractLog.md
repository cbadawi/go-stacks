# TransactionEventSmartContractLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventIndex** | **int32** |  | 
**EventType** | **string** |  | 
**TxId** | **string** |  | 
**ContractLog** | [**TransactionEventSmartContractLogAllOfContractLog**](TransactionEventSmartContractLogAllOfContractLog.md) |  | 

## Methods

### NewTransactionEventSmartContractLog

`func NewTransactionEventSmartContractLog(eventIndex int32, eventType string, txId string, contractLog TransactionEventSmartContractLogAllOfContractLog, ) *TransactionEventSmartContractLog`

NewTransactionEventSmartContractLog instantiates a new TransactionEventSmartContractLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionEventSmartContractLogWithDefaults

`func NewTransactionEventSmartContractLogWithDefaults() *TransactionEventSmartContractLog`

NewTransactionEventSmartContractLogWithDefaults instantiates a new TransactionEventSmartContractLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventIndex

`func (o *TransactionEventSmartContractLog) GetEventIndex() int32`

GetEventIndex returns the EventIndex field if non-nil, zero value otherwise.

### GetEventIndexOk

`func (o *TransactionEventSmartContractLog) GetEventIndexOk() (*int32, bool)`

GetEventIndexOk returns a tuple with the EventIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventIndex

`func (o *TransactionEventSmartContractLog) SetEventIndex(v int32)`

SetEventIndex sets EventIndex field to given value.


### GetEventType

`func (o *TransactionEventSmartContractLog) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *TransactionEventSmartContractLog) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *TransactionEventSmartContractLog) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetTxId

`func (o *TransactionEventSmartContractLog) GetTxId() string`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *TransactionEventSmartContractLog) GetTxIdOk() (*string, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *TransactionEventSmartContractLog) SetTxId(v string)`

SetTxId sets TxId field to given value.


### GetContractLog

`func (o *TransactionEventSmartContractLog) GetContractLog() TransactionEventSmartContractLogAllOfContractLog`

GetContractLog returns the ContractLog field if non-nil, zero value otherwise.

### GetContractLogOk

`func (o *TransactionEventSmartContractLog) GetContractLogOk() (*TransactionEventSmartContractLogAllOfContractLog, bool)`

GetContractLogOk returns a tuple with the ContractLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractLog

`func (o *TransactionEventSmartContractLog) SetContractLog(v TransactionEventSmartContractLogAllOfContractLog)`

SetContractLog sets ContractLog field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


