# TxSearchResultResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | **string** | The id used to search this query. | 
**EntityType** | **string** |  | 
**TxData** | [**TxSearchResultResultTxData**](TxSearchResultResultTxData.md) |  | 
**Metadata** | Pointer to [**Transaction**](Transaction.md) |  | [optional] 

## Methods

### NewTxSearchResultResult

`func NewTxSearchResultResult(entityId string, entityType string, txData TxSearchResultResultTxData, ) *TxSearchResultResult`

NewTxSearchResultResult instantiates a new TxSearchResultResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTxSearchResultResultWithDefaults

`func NewTxSearchResultResultWithDefaults() *TxSearchResultResult`

NewTxSearchResultResultWithDefaults instantiates a new TxSearchResultResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityId

`func (o *TxSearchResultResult) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *TxSearchResultResult) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *TxSearchResultResult) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.


### GetEntityType

`func (o *TxSearchResultResult) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *TxSearchResultResult) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *TxSearchResultResult) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetTxData

`func (o *TxSearchResultResult) GetTxData() TxSearchResultResultTxData`

GetTxData returns the TxData field if non-nil, zero value otherwise.

### GetTxDataOk

`func (o *TxSearchResultResult) GetTxDataOk() (*TxSearchResultResultTxData, bool)`

GetTxDataOk returns a tuple with the TxData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxData

`func (o *TxSearchResultResult) SetTxData(v TxSearchResultResultTxData)`

SetTxData sets TxData field to given value.


### GetMetadata

`func (o *TxSearchResultResult) GetMetadata() Transaction`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *TxSearchResultResult) GetMetadataOk() (*Transaction, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *TxSearchResultResult) SetMetadata(v Transaction)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *TxSearchResultResult) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


