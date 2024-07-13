# ContractCallTransactionMetadataContractCall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContractId** | **string** | Contract identifier formatted as &#x60;&lt;principaladdress&gt;.&lt;contract_name&gt;&#x60; | 
**FunctionName** | **string** | Name of the Clarity function to be invoked | 
**FunctionSignature** | **string** | Function definition, including function name and type as well as parameter names and types | 
**FunctionArgs** | Pointer to [**[]ContractCallTransactionMetadataContractCallFunctionArgsInner**](ContractCallTransactionMetadataContractCallFunctionArgsInner.md) | List of arguments used to invoke the function | [optional] 

## Methods

### NewContractCallTransactionMetadataContractCall

`func NewContractCallTransactionMetadataContractCall(contractId string, functionName string, functionSignature string, ) *ContractCallTransactionMetadataContractCall`

NewContractCallTransactionMetadataContractCall instantiates a new ContractCallTransactionMetadataContractCall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractCallTransactionMetadataContractCallWithDefaults

`func NewContractCallTransactionMetadataContractCallWithDefaults() *ContractCallTransactionMetadataContractCall`

NewContractCallTransactionMetadataContractCallWithDefaults instantiates a new ContractCallTransactionMetadataContractCall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContractId

`func (o *ContractCallTransactionMetadataContractCall) GetContractId() string`

GetContractId returns the ContractId field if non-nil, zero value otherwise.

### GetContractIdOk

`func (o *ContractCallTransactionMetadataContractCall) GetContractIdOk() (*string, bool)`

GetContractIdOk returns a tuple with the ContractId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractId

`func (o *ContractCallTransactionMetadataContractCall) SetContractId(v string)`

SetContractId sets ContractId field to given value.


### GetFunctionName

`func (o *ContractCallTransactionMetadataContractCall) GetFunctionName() string`

GetFunctionName returns the FunctionName field if non-nil, zero value otherwise.

### GetFunctionNameOk

`func (o *ContractCallTransactionMetadataContractCall) GetFunctionNameOk() (*string, bool)`

GetFunctionNameOk returns a tuple with the FunctionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionName

`func (o *ContractCallTransactionMetadataContractCall) SetFunctionName(v string)`

SetFunctionName sets FunctionName field to given value.


### GetFunctionSignature

`func (o *ContractCallTransactionMetadataContractCall) GetFunctionSignature() string`

GetFunctionSignature returns the FunctionSignature field if non-nil, zero value otherwise.

### GetFunctionSignatureOk

`func (o *ContractCallTransactionMetadataContractCall) GetFunctionSignatureOk() (*string, bool)`

GetFunctionSignatureOk returns a tuple with the FunctionSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionSignature

`func (o *ContractCallTransactionMetadataContractCall) SetFunctionSignature(v string)`

SetFunctionSignature sets FunctionSignature field to given value.


### GetFunctionArgs

`func (o *ContractCallTransactionMetadataContractCall) GetFunctionArgs() []ContractCallTransactionMetadataContractCallFunctionArgsInner`

GetFunctionArgs returns the FunctionArgs field if non-nil, zero value otherwise.

### GetFunctionArgsOk

`func (o *ContractCallTransactionMetadataContractCall) GetFunctionArgsOk() (*[]ContractCallTransactionMetadataContractCallFunctionArgsInner, bool)`

GetFunctionArgsOk returns a tuple with the FunctionArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionArgs

`func (o *ContractCallTransactionMetadataContractCall) SetFunctionArgs(v []ContractCallTransactionMetadataContractCallFunctionArgsInner)`

SetFunctionArgs sets FunctionArgs field to given value.

### HasFunctionArgs

`func (o *ContractCallTransactionMetadataContractCall) HasFunctionArgs() bool`

HasFunctionArgs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


