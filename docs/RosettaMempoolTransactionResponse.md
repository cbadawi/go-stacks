# RosettaMempoolTransactionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Transaction** | [**RosettaTransaction**](RosettaTransaction.md) |  | 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewRosettaMempoolTransactionResponse

`func NewRosettaMempoolTransactionResponse(transaction RosettaTransaction, ) *RosettaMempoolTransactionResponse`

NewRosettaMempoolTransactionResponse instantiates a new RosettaMempoolTransactionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRosettaMempoolTransactionResponseWithDefaults

`func NewRosettaMempoolTransactionResponseWithDefaults() *RosettaMempoolTransactionResponse`

NewRosettaMempoolTransactionResponseWithDefaults instantiates a new RosettaMempoolTransactionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransaction

`func (o *RosettaMempoolTransactionResponse) GetTransaction() RosettaTransaction`

GetTransaction returns the Transaction field if non-nil, zero value otherwise.

### GetTransactionOk

`func (o *RosettaMempoolTransactionResponse) GetTransactionOk() (*RosettaTransaction, bool)`

GetTransactionOk returns a tuple with the Transaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransaction

`func (o *RosettaMempoolTransactionResponse) SetTransaction(v RosettaTransaction)`

SetTransaction sets Transaction field to given value.


### GetMetadata

`func (o *RosettaMempoolTransactionResponse) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RosettaMempoolTransactionResponse) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RosettaMempoolTransactionResponse) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RosettaMempoolTransactionResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


