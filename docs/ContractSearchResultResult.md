# ContractSearchResultResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | **string** | The id used to search this query. | 
**EntityType** | **string** |  | 
**TxData** | Pointer to [**ContractSearchResultResultTxData**](ContractSearchResultResultTxData.md) |  | [optional] 
**Metadata** | Pointer to [**AddressTransactionsListResponseResultsInner**](AddressTransactionsListResponseResultsInner.md) |  | [optional] 

## Methods

### NewContractSearchResultResult

`func NewContractSearchResultResult(entityId string, entityType string, ) *ContractSearchResultResult`

NewContractSearchResultResult instantiates a new ContractSearchResultResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractSearchResultResultWithDefaults

`func NewContractSearchResultResultWithDefaults() *ContractSearchResultResult`

NewContractSearchResultResultWithDefaults instantiates a new ContractSearchResultResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityId

`func (o *ContractSearchResultResult) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *ContractSearchResultResult) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *ContractSearchResultResult) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.


### GetEntityType

`func (o *ContractSearchResultResult) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *ContractSearchResultResult) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *ContractSearchResultResult) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetTxData

`func (o *ContractSearchResultResult) GetTxData() ContractSearchResultResultTxData`

GetTxData returns the TxData field if non-nil, zero value otherwise.

### GetTxDataOk

`func (o *ContractSearchResultResult) GetTxDataOk() (*ContractSearchResultResultTxData, bool)`

GetTxDataOk returns a tuple with the TxData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxData

`func (o *ContractSearchResultResult) SetTxData(v ContractSearchResultResultTxData)`

SetTxData sets TxData field to given value.

### HasTxData

`func (o *ContractSearchResultResult) HasTxData() bool`

HasTxData returns a boolean if a field has been set.

### GetMetadata

`func (o *ContractSearchResultResult) GetMetadata() AddressTransactionsListResponseResultsInner`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ContractSearchResultResult) GetMetadataOk() (*AddressTransactionsListResponseResultsInner, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ContractSearchResultResult) SetMetadata(v AddressTransactionsListResponseResultsInner)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ContractSearchResultResult) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


