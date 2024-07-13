# MempoolTxSearchResultResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | **string** | The id used to search this query. | 
**EntityType** | **string** |  | 
**TxData** | [**MempoolTxSearchResultResultTxData**](MempoolTxSearchResultResultTxData.md) |  | 
**Metadata** | Pointer to [**MempoolTransaction**](MempoolTransaction.md) |  | [optional] 

## Methods

### NewMempoolTxSearchResultResult

`func NewMempoolTxSearchResultResult(entityId string, entityType string, txData MempoolTxSearchResultResultTxData, ) *MempoolTxSearchResultResult`

NewMempoolTxSearchResultResult instantiates a new MempoolTxSearchResultResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMempoolTxSearchResultResultWithDefaults

`func NewMempoolTxSearchResultResultWithDefaults() *MempoolTxSearchResultResult`

NewMempoolTxSearchResultResultWithDefaults instantiates a new MempoolTxSearchResultResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityId

`func (o *MempoolTxSearchResultResult) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *MempoolTxSearchResultResult) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *MempoolTxSearchResultResult) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.


### GetEntityType

`func (o *MempoolTxSearchResultResult) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *MempoolTxSearchResultResult) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *MempoolTxSearchResultResult) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetTxData

`func (o *MempoolTxSearchResultResult) GetTxData() MempoolTxSearchResultResultTxData`

GetTxData returns the TxData field if non-nil, zero value otherwise.

### GetTxDataOk

`func (o *MempoolTxSearchResultResult) GetTxDataOk() (*MempoolTxSearchResultResultTxData, bool)`

GetTxDataOk returns a tuple with the TxData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxData

`func (o *MempoolTxSearchResultResult) SetTxData(v MempoolTxSearchResultResultTxData)`

SetTxData sets TxData field to given value.


### GetMetadata

`func (o *MempoolTxSearchResultResult) GetMetadata() MempoolTransaction`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *MempoolTxSearchResultResult) GetMetadataOk() (*MempoolTransaction, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *MempoolTxSearchResultResult) SetMetadata(v MempoolTransaction)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *MempoolTxSearchResultResult) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


