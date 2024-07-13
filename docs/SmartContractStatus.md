# SmartContractStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | Smart contract deployment transaction status | 
**TxId** | **string** | Deployment transaction ID | 
**ContractId** | **string** | Smart contract ID | 
**BlockHeight** | Pointer to **int32** | Height of the transaction confirmation block | [optional] 

## Methods

### NewSmartContractStatus

`func NewSmartContractStatus(status string, txId string, contractId string, ) *SmartContractStatus`

NewSmartContractStatus instantiates a new SmartContractStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmartContractStatusWithDefaults

`func NewSmartContractStatusWithDefaults() *SmartContractStatus`

NewSmartContractStatusWithDefaults instantiates a new SmartContractStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *SmartContractStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SmartContractStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SmartContractStatus) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTxId

`func (o *SmartContractStatus) GetTxId() string`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *SmartContractStatus) GetTxIdOk() (*string, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *SmartContractStatus) SetTxId(v string)`

SetTxId sets TxId field to given value.


### GetContractId

`func (o *SmartContractStatus) GetContractId() string`

GetContractId returns the ContractId field if non-nil, zero value otherwise.

### GetContractIdOk

`func (o *SmartContractStatus) GetContractIdOk() (*string, bool)`

GetContractIdOk returns a tuple with the ContractId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractId

`func (o *SmartContractStatus) SetContractId(v string)`

SetContractId sets ContractId field to given value.


### GetBlockHeight

`func (o *SmartContractStatus) GetBlockHeight() int32`

GetBlockHeight returns the BlockHeight field if non-nil, zero value otherwise.

### GetBlockHeightOk

`func (o *SmartContractStatus) GetBlockHeightOk() (*int32, bool)`

GetBlockHeightOk returns a tuple with the BlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHeight

`func (o *SmartContractStatus) SetBlockHeight(v int32)`

SetBlockHeight sets BlockHeight field to given value.

### HasBlockHeight

`func (o *SmartContractStatus) HasBlockHeight() bool`

HasBlockHeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


