# TransactionEventFungibleAsset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventIndex** | **int32** |  | 
**EventType** | **string** |  | 
**TxId** | **string** |  | 
**Asset** | [**TransactionEventFungibleAssetAllOfAsset**](TransactionEventFungibleAssetAllOfAsset.md) |  | 

## Methods

### NewTransactionEventFungibleAsset

`func NewTransactionEventFungibleAsset(eventIndex int32, eventType string, txId string, asset TransactionEventFungibleAssetAllOfAsset, ) *TransactionEventFungibleAsset`

NewTransactionEventFungibleAsset instantiates a new TransactionEventFungibleAsset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionEventFungibleAssetWithDefaults

`func NewTransactionEventFungibleAssetWithDefaults() *TransactionEventFungibleAsset`

NewTransactionEventFungibleAssetWithDefaults instantiates a new TransactionEventFungibleAsset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventIndex

`func (o *TransactionEventFungibleAsset) GetEventIndex() int32`

GetEventIndex returns the EventIndex field if non-nil, zero value otherwise.

### GetEventIndexOk

`func (o *TransactionEventFungibleAsset) GetEventIndexOk() (*int32, bool)`

GetEventIndexOk returns a tuple with the EventIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventIndex

`func (o *TransactionEventFungibleAsset) SetEventIndex(v int32)`

SetEventIndex sets EventIndex field to given value.


### GetEventType

`func (o *TransactionEventFungibleAsset) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *TransactionEventFungibleAsset) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *TransactionEventFungibleAsset) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetTxId

`func (o *TransactionEventFungibleAsset) GetTxId() string`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *TransactionEventFungibleAsset) GetTxIdOk() (*string, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *TransactionEventFungibleAsset) SetTxId(v string)`

SetTxId sets TxId field to given value.


### GetAsset

`func (o *TransactionEventFungibleAsset) GetAsset() TransactionEventFungibleAssetAllOfAsset`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *TransactionEventFungibleAsset) GetAssetOk() (*TransactionEventFungibleAssetAllOfAsset, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *TransactionEventFungibleAsset) SetAsset(v TransactionEventFungibleAssetAllOfAsset)`

SetAsset sets Asset field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


