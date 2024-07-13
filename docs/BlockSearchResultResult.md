# BlockSearchResultResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | **string** | The id used to search this query. | 
**EntityType** | **string** |  | 
**BlockData** | [**BlockSearchResultResultBlockData**](BlockSearchResultResultBlockData.md) |  | 
**Metadata** | Pointer to [**Block**](Block.md) |  | [optional] 

## Methods

### NewBlockSearchResultResult

`func NewBlockSearchResultResult(entityId string, entityType string, blockData BlockSearchResultResultBlockData, ) *BlockSearchResultResult`

NewBlockSearchResultResult instantiates a new BlockSearchResultResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockSearchResultResultWithDefaults

`func NewBlockSearchResultResultWithDefaults() *BlockSearchResultResult`

NewBlockSearchResultResultWithDefaults instantiates a new BlockSearchResultResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityId

`func (o *BlockSearchResultResult) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *BlockSearchResultResult) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *BlockSearchResultResult) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.


### GetEntityType

`func (o *BlockSearchResultResult) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *BlockSearchResultResult) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *BlockSearchResultResult) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetBlockData

`func (o *BlockSearchResultResult) GetBlockData() BlockSearchResultResultBlockData`

GetBlockData returns the BlockData field if non-nil, zero value otherwise.

### GetBlockDataOk

`func (o *BlockSearchResultResult) GetBlockDataOk() (*BlockSearchResultResultBlockData, bool)`

GetBlockDataOk returns a tuple with the BlockData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockData

`func (o *BlockSearchResultResult) SetBlockData(v BlockSearchResultResultBlockData)`

SetBlockData sets BlockData field to given value.


### GetMetadata

`func (o *BlockSearchResultResult) GetMetadata() Block`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BlockSearchResultResult) GetMetadataOk() (*Block, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BlockSearchResultResult) SetMetadata(v Block)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *BlockSearchResultResult) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


