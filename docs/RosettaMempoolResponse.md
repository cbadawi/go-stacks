# RosettaMempoolResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionIdentifiers** | [**[]TransactionIdentifier**](TransactionIdentifier.md) |  | 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewRosettaMempoolResponse

`func NewRosettaMempoolResponse(transactionIdentifiers []TransactionIdentifier, ) *RosettaMempoolResponse`

NewRosettaMempoolResponse instantiates a new RosettaMempoolResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRosettaMempoolResponseWithDefaults

`func NewRosettaMempoolResponseWithDefaults() *RosettaMempoolResponse`

NewRosettaMempoolResponseWithDefaults instantiates a new RosettaMempoolResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionIdentifiers

`func (o *RosettaMempoolResponse) GetTransactionIdentifiers() []TransactionIdentifier`

GetTransactionIdentifiers returns the TransactionIdentifiers field if non-nil, zero value otherwise.

### GetTransactionIdentifiersOk

`func (o *RosettaMempoolResponse) GetTransactionIdentifiersOk() (*[]TransactionIdentifier, bool)`

GetTransactionIdentifiersOk returns a tuple with the TransactionIdentifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionIdentifiers

`func (o *RosettaMempoolResponse) SetTransactionIdentifiers(v []TransactionIdentifier)`

SetTransactionIdentifiers sets TransactionIdentifiers field to given value.


### GetMetadata

`func (o *RosettaMempoolResponse) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RosettaMempoolResponse) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RosettaMempoolResponse) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RosettaMempoolResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


