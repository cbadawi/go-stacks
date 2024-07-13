# SmartContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TxId** | **string** |  | 
**Canonical** | **bool** |  | 
**ContractId** | **string** |  | 
**BlockHeight** | **int32** |  | 
**SourceCode** | **string** |  | 
**Abi** | **string** |  | 

## Methods

### NewSmartContract

`func NewSmartContract(txId string, canonical bool, contractId string, blockHeight int32, sourceCode string, abi string, ) *SmartContract`

NewSmartContract instantiates a new SmartContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmartContractWithDefaults

`func NewSmartContractWithDefaults() *SmartContract`

NewSmartContractWithDefaults instantiates a new SmartContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTxId

`func (o *SmartContract) GetTxId() string`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *SmartContract) GetTxIdOk() (*string, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *SmartContract) SetTxId(v string)`

SetTxId sets TxId field to given value.


### GetCanonical

`func (o *SmartContract) GetCanonical() bool`

GetCanonical returns the Canonical field if non-nil, zero value otherwise.

### GetCanonicalOk

`func (o *SmartContract) GetCanonicalOk() (*bool, bool)`

GetCanonicalOk returns a tuple with the Canonical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonical

`func (o *SmartContract) SetCanonical(v bool)`

SetCanonical sets Canonical field to given value.


### GetContractId

`func (o *SmartContract) GetContractId() string`

GetContractId returns the ContractId field if non-nil, zero value otherwise.

### GetContractIdOk

`func (o *SmartContract) GetContractIdOk() (*string, bool)`

GetContractIdOk returns a tuple with the ContractId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractId

`func (o *SmartContract) SetContractId(v string)`

SetContractId sets ContractId field to given value.


### GetBlockHeight

`func (o *SmartContract) GetBlockHeight() int32`

GetBlockHeight returns the BlockHeight field if non-nil, zero value otherwise.

### GetBlockHeightOk

`func (o *SmartContract) GetBlockHeightOk() (*int32, bool)`

GetBlockHeightOk returns a tuple with the BlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHeight

`func (o *SmartContract) SetBlockHeight(v int32)`

SetBlockHeight sets BlockHeight field to given value.


### GetSourceCode

`func (o *SmartContract) GetSourceCode() string`

GetSourceCode returns the SourceCode field if non-nil, zero value otherwise.

### GetSourceCodeOk

`func (o *SmartContract) GetSourceCodeOk() (*string, bool)`

GetSourceCodeOk returns a tuple with the SourceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCode

`func (o *SmartContract) SetSourceCode(v string)`

SetSourceCode sets SourceCode field to given value.


### GetAbi

`func (o *SmartContract) GetAbi() string`

GetAbi returns the Abi field if non-nil, zero value otherwise.

### GetAbiOk

`func (o *SmartContract) GetAbiOk() (*string, bool)`

GetAbiOk returns a tuple with the Abi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbi

`func (o *SmartContract) SetAbi(v string)`

SetAbi sets Abi field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


