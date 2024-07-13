# TransactionEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventIndex** | **int32** |  | 
**EventType** | **string** |  | 
**TxId** | **string** |  | 
**ContractLog** | [**TransactionEventSmartContractLogAllOfContractLog**](TransactionEventSmartContractLogAllOfContractLog.md) |  | 
**StxLockEvent** | [**TransactionEventStxLockAllOfStxLockEvent**](TransactionEventStxLockAllOfStxLockEvent.md) |  | 
**Asset** | [**TransactionEventNonFungibleAssetAllOfAsset**](TransactionEventNonFungibleAssetAllOfAsset.md) |  | 

## Methods

### NewTransactionEvent

`func NewTransactionEvent(eventIndex int32, eventType string, txId string, contractLog TransactionEventSmartContractLogAllOfContractLog, stxLockEvent TransactionEventStxLockAllOfStxLockEvent, asset TransactionEventNonFungibleAssetAllOfAsset, ) *TransactionEvent`

NewTransactionEvent instantiates a new TransactionEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionEventWithDefaults

`func NewTransactionEventWithDefaults() *TransactionEvent`

NewTransactionEventWithDefaults instantiates a new TransactionEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventIndex

`func (o *TransactionEvent) GetEventIndex() int32`

GetEventIndex returns the EventIndex field if non-nil, zero value otherwise.

### GetEventIndexOk

`func (o *TransactionEvent) GetEventIndexOk() (*int32, bool)`

GetEventIndexOk returns a tuple with the EventIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventIndex

`func (o *TransactionEvent) SetEventIndex(v int32)`

SetEventIndex sets EventIndex field to given value.


### GetEventType

`func (o *TransactionEvent) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *TransactionEvent) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *TransactionEvent) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetTxId

`func (o *TransactionEvent) GetTxId() string`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *TransactionEvent) GetTxIdOk() (*string, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *TransactionEvent) SetTxId(v string)`

SetTxId sets TxId field to given value.


### GetContractLog

`func (o *TransactionEvent) GetContractLog() TransactionEventSmartContractLogAllOfContractLog`

GetContractLog returns the ContractLog field if non-nil, zero value otherwise.

### GetContractLogOk

`func (o *TransactionEvent) GetContractLogOk() (*TransactionEventSmartContractLogAllOfContractLog, bool)`

GetContractLogOk returns a tuple with the ContractLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractLog

`func (o *TransactionEvent) SetContractLog(v TransactionEventSmartContractLogAllOfContractLog)`

SetContractLog sets ContractLog field to given value.


### GetStxLockEvent

`func (o *TransactionEvent) GetStxLockEvent() TransactionEventStxLockAllOfStxLockEvent`

GetStxLockEvent returns the StxLockEvent field if non-nil, zero value otherwise.

### GetStxLockEventOk

`func (o *TransactionEvent) GetStxLockEventOk() (*TransactionEventStxLockAllOfStxLockEvent, bool)`

GetStxLockEventOk returns a tuple with the StxLockEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStxLockEvent

`func (o *TransactionEvent) SetStxLockEvent(v TransactionEventStxLockAllOfStxLockEvent)`

SetStxLockEvent sets StxLockEvent field to given value.


### GetAsset

`func (o *TransactionEvent) GetAsset() TransactionEventNonFungibleAssetAllOfAsset`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *TransactionEvent) GetAssetOk() (*TransactionEventNonFungibleAssetAllOfAsset, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *TransactionEvent) SetAsset(v TransactionEventNonFungibleAssetAllOfAsset)`

SetAsset sets Asset field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


