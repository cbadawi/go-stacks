# InboundStxTransfer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sender** | **string** | Principal that sent this transfer | 
**Amount** | **string** | Transfer amount in micro-STX as integer string | 
**Memo** | **string** | Hex encoded memo bytes associated with the transfer | 
**BlockHeight** | **float32** | Block height at which this transfer occurred | 
**TxId** | **string** | The transaction ID in which this transfer occurred | 
**TransferType** | **string** | Indicates if the transfer is from a stx-transfer transaction or a contract-call transaction | 
**TxIndex** | **float32** | Index of the transaction within a block | 

## Methods

### NewInboundStxTransfer

`func NewInboundStxTransfer(sender string, amount string, memo string, blockHeight float32, txId string, transferType string, txIndex float32, ) *InboundStxTransfer`

NewInboundStxTransfer instantiates a new InboundStxTransfer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInboundStxTransferWithDefaults

`func NewInboundStxTransferWithDefaults() *InboundStxTransfer`

NewInboundStxTransferWithDefaults instantiates a new InboundStxTransfer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSender

`func (o *InboundStxTransfer) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *InboundStxTransfer) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *InboundStxTransfer) SetSender(v string)`

SetSender sets Sender field to given value.


### GetAmount

`func (o *InboundStxTransfer) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *InboundStxTransfer) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *InboundStxTransfer) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetMemo

`func (o *InboundStxTransfer) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *InboundStxTransfer) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *InboundStxTransfer) SetMemo(v string)`

SetMemo sets Memo field to given value.


### GetBlockHeight

`func (o *InboundStxTransfer) GetBlockHeight() float32`

GetBlockHeight returns the BlockHeight field if non-nil, zero value otherwise.

### GetBlockHeightOk

`func (o *InboundStxTransfer) GetBlockHeightOk() (*float32, bool)`

GetBlockHeightOk returns a tuple with the BlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHeight

`func (o *InboundStxTransfer) SetBlockHeight(v float32)`

SetBlockHeight sets BlockHeight field to given value.


### GetTxId

`func (o *InboundStxTransfer) GetTxId() string`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *InboundStxTransfer) GetTxIdOk() (*string, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *InboundStxTransfer) SetTxId(v string)`

SetTxId sets TxId field to given value.


### GetTransferType

`func (o *InboundStxTransfer) GetTransferType() string`

GetTransferType returns the TransferType field if non-nil, zero value otherwise.

### GetTransferTypeOk

`func (o *InboundStxTransfer) GetTransferTypeOk() (*string, bool)`

GetTransferTypeOk returns a tuple with the TransferType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferType

`func (o *InboundStxTransfer) SetTransferType(v string)`

SetTransferType sets TransferType field to given value.


### GetTxIndex

`func (o *InboundStxTransfer) GetTxIndex() float32`

GetTxIndex returns the TxIndex field if non-nil, zero value otherwise.

### GetTxIndexOk

`func (o *InboundStxTransfer) GetTxIndexOk() (*float32, bool)`

GetTxIndexOk returns a tuple with the TxIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxIndex

`func (o *InboundStxTransfer) SetTxIndex(v float32)`

SetTxIndex sets TxIndex field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


