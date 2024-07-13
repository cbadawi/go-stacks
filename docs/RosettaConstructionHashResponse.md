# RosettaConstructionHashResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionIdentifier** | [**TransactionIdentifier**](TransactionIdentifier.md) |  | 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewRosettaConstructionHashResponse

`func NewRosettaConstructionHashResponse(transactionIdentifier TransactionIdentifier, ) *RosettaConstructionHashResponse`

NewRosettaConstructionHashResponse instantiates a new RosettaConstructionHashResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRosettaConstructionHashResponseWithDefaults

`func NewRosettaConstructionHashResponseWithDefaults() *RosettaConstructionHashResponse`

NewRosettaConstructionHashResponseWithDefaults instantiates a new RosettaConstructionHashResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionIdentifier

`func (o *RosettaConstructionHashResponse) GetTransactionIdentifier() TransactionIdentifier`

GetTransactionIdentifier returns the TransactionIdentifier field if non-nil, zero value otherwise.

### GetTransactionIdentifierOk

`func (o *RosettaConstructionHashResponse) GetTransactionIdentifierOk() (*TransactionIdentifier, bool)`

GetTransactionIdentifierOk returns a tuple with the TransactionIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionIdentifier

`func (o *RosettaConstructionHashResponse) SetTransactionIdentifier(v TransactionIdentifier)`

SetTransactionIdentifier sets TransactionIdentifier field to given value.


### GetMetadata

`func (o *RosettaConstructionHashResponse) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RosettaConstructionHashResponse) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RosettaConstructionHashResponse) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RosettaConstructionHashResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


