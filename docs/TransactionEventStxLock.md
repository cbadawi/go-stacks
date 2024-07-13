# TransactionEventStxLock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventIndex** | **int32** |  | 
**EventType** | **string** |  | 
**TxId** | **string** |  | 
**StxLockEvent** | [**TransactionEventStxLockAllOfStxLockEvent**](TransactionEventStxLockAllOfStxLockEvent.md) |  | 

## Methods

### NewTransactionEventStxLock

`func NewTransactionEventStxLock(eventIndex int32, eventType string, txId string, stxLockEvent TransactionEventStxLockAllOfStxLockEvent, ) *TransactionEventStxLock`

NewTransactionEventStxLock instantiates a new TransactionEventStxLock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionEventStxLockWithDefaults

`func NewTransactionEventStxLockWithDefaults() *TransactionEventStxLock`

NewTransactionEventStxLockWithDefaults instantiates a new TransactionEventStxLock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventIndex

`func (o *TransactionEventStxLock) GetEventIndex() int32`

GetEventIndex returns the EventIndex field if non-nil, zero value otherwise.

### GetEventIndexOk

`func (o *TransactionEventStxLock) GetEventIndexOk() (*int32, bool)`

GetEventIndexOk returns a tuple with the EventIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventIndex

`func (o *TransactionEventStxLock) SetEventIndex(v int32)`

SetEventIndex sets EventIndex field to given value.


### GetEventType

`func (o *TransactionEventStxLock) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *TransactionEventStxLock) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *TransactionEventStxLock) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetTxId

`func (o *TransactionEventStxLock) GetTxId() string`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *TransactionEventStxLock) GetTxIdOk() (*string, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *TransactionEventStxLock) SetTxId(v string)`

SetTxId sets TxId field to given value.


### GetStxLockEvent

`func (o *TransactionEventStxLock) GetStxLockEvent() TransactionEventStxLockAllOfStxLockEvent`

GetStxLockEvent returns the StxLockEvent field if non-nil, zero value otherwise.

### GetStxLockEventOk

`func (o *TransactionEventStxLock) GetStxLockEventOk() (*TransactionEventStxLockAllOfStxLockEvent, bool)`

GetStxLockEventOk returns a tuple with the StxLockEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStxLockEvent

`func (o *TransactionEventStxLock) SetStxLockEvent(v TransactionEventStxLockAllOfStxLockEvent)`

SetStxLockEvent sets StxLockEvent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


