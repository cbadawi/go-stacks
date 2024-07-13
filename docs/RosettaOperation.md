# RosettaOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OperationIdentifier** | [**RosettaOperationIdentifier**](RosettaOperationIdentifier.md) |  | 
**RelatedOperations** | Pointer to [**[]RosettaRelatedOperation**](RosettaRelatedOperation.md) | Restrict referenced related_operations to identifier indexes &lt; the current operation_identifier.index. This ensures there exists a clear DAG-structure of relations. Since operations are one-sided, one could imagine relating operations in a single transfer or linking operations in a call tree. | [optional] 
**Type** | **string** | The network-specific type of the operation. Ensure that any type that can be returned here is also specified in the NetworkStatus. This can be very useful to downstream consumers that parse all block data. | 
**Status** | Pointer to **string** | The network-specific status of the operation. Status is not defined on the transaction object because blockchains with smart contracts may have transactions that partially apply. Blockchains with atomic transactions (all operations succeed or all operations fail) will have the same status for each operation. | [optional] 
**Account** | Pointer to [**RosettaAccount**](RosettaAccount.md) |  | [optional] 
**Amount** | Pointer to [**RosettaAmount**](RosettaAmount.md) |  | [optional] 
**CoinChange** | Pointer to [**RosettaCoinChange**](RosettaCoinChange.md) |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Operations Meta Data | [optional] 

## Methods

### NewRosettaOperation

`func NewRosettaOperation(operationIdentifier RosettaOperationIdentifier, type_ string, ) *RosettaOperation`

NewRosettaOperation instantiates a new RosettaOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRosettaOperationWithDefaults

`func NewRosettaOperationWithDefaults() *RosettaOperation`

NewRosettaOperationWithDefaults instantiates a new RosettaOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperationIdentifier

`func (o *RosettaOperation) GetOperationIdentifier() RosettaOperationIdentifier`

GetOperationIdentifier returns the OperationIdentifier field if non-nil, zero value otherwise.

### GetOperationIdentifierOk

`func (o *RosettaOperation) GetOperationIdentifierOk() (*RosettaOperationIdentifier, bool)`

GetOperationIdentifierOk returns a tuple with the OperationIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationIdentifier

`func (o *RosettaOperation) SetOperationIdentifier(v RosettaOperationIdentifier)`

SetOperationIdentifier sets OperationIdentifier field to given value.


### GetRelatedOperations

`func (o *RosettaOperation) GetRelatedOperations() []RosettaRelatedOperation`

GetRelatedOperations returns the RelatedOperations field if non-nil, zero value otherwise.

### GetRelatedOperationsOk

`func (o *RosettaOperation) GetRelatedOperationsOk() (*[]RosettaRelatedOperation, bool)`

GetRelatedOperationsOk returns a tuple with the RelatedOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedOperations

`func (o *RosettaOperation) SetRelatedOperations(v []RosettaRelatedOperation)`

SetRelatedOperations sets RelatedOperations field to given value.

### HasRelatedOperations

`func (o *RosettaOperation) HasRelatedOperations() bool`

HasRelatedOperations returns a boolean if a field has been set.

### GetType

`func (o *RosettaOperation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RosettaOperation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RosettaOperation) SetType(v string)`

SetType sets Type field to given value.


### GetStatus

`func (o *RosettaOperation) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RosettaOperation) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RosettaOperation) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RosettaOperation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAccount

`func (o *RosettaOperation) GetAccount() RosettaAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *RosettaOperation) GetAccountOk() (*RosettaAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *RosettaOperation) SetAccount(v RosettaAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *RosettaOperation) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAmount

`func (o *RosettaOperation) GetAmount() RosettaAmount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *RosettaOperation) GetAmountOk() (*RosettaAmount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *RosettaOperation) SetAmount(v RosettaAmount)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *RosettaOperation) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCoinChange

`func (o *RosettaOperation) GetCoinChange() RosettaCoinChange`

GetCoinChange returns the CoinChange field if non-nil, zero value otherwise.

### GetCoinChangeOk

`func (o *RosettaOperation) GetCoinChangeOk() (*RosettaCoinChange, bool)`

GetCoinChangeOk returns a tuple with the CoinChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoinChange

`func (o *RosettaOperation) SetCoinChange(v RosettaCoinChange)`

SetCoinChange sets CoinChange field to given value.

### HasCoinChange

`func (o *RosettaOperation) HasCoinChange() bool`

HasCoinChange returns a boolean if a field has been set.

### GetMetadata

`func (o *RosettaOperation) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RosettaOperation) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RosettaOperation) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RosettaOperation) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


