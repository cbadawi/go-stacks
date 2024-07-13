# NonFungibleTokenMintWithTxMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Recipient** | Pointer to **string** |  | [optional] 
**EventIndex** | **int32** |  | 
**Value** | [**NonFungibleTokenHoldingWithTxIdValue**](NonFungibleTokenHoldingWithTxIdValue.md) |  | 
**Tx** | [**Transaction**](Transaction.md) |  | 

## Methods

### NewNonFungibleTokenMintWithTxMetadata

`func NewNonFungibleTokenMintWithTxMetadata(eventIndex int32, value NonFungibleTokenHoldingWithTxIdValue, tx Transaction, ) *NonFungibleTokenMintWithTxMetadata`

NewNonFungibleTokenMintWithTxMetadata instantiates a new NonFungibleTokenMintWithTxMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNonFungibleTokenMintWithTxMetadataWithDefaults

`func NewNonFungibleTokenMintWithTxMetadataWithDefaults() *NonFungibleTokenMintWithTxMetadata`

NewNonFungibleTokenMintWithTxMetadataWithDefaults instantiates a new NonFungibleTokenMintWithTxMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecipient

`func (o *NonFungibleTokenMintWithTxMetadata) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *NonFungibleTokenMintWithTxMetadata) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *NonFungibleTokenMintWithTxMetadata) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *NonFungibleTokenMintWithTxMetadata) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.

### GetEventIndex

`func (o *NonFungibleTokenMintWithTxMetadata) GetEventIndex() int32`

GetEventIndex returns the EventIndex field if non-nil, zero value otherwise.

### GetEventIndexOk

`func (o *NonFungibleTokenMintWithTxMetadata) GetEventIndexOk() (*int32, bool)`

GetEventIndexOk returns a tuple with the EventIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventIndex

`func (o *NonFungibleTokenMintWithTxMetadata) SetEventIndex(v int32)`

SetEventIndex sets EventIndex field to given value.


### GetValue

`func (o *NonFungibleTokenMintWithTxMetadata) GetValue() NonFungibleTokenHoldingWithTxIdValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *NonFungibleTokenMintWithTxMetadata) GetValueOk() (*NonFungibleTokenHoldingWithTxIdValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *NonFungibleTokenMintWithTxMetadata) SetValue(v NonFungibleTokenHoldingWithTxIdValue)`

SetValue sets Value field to given value.


### GetTx

`func (o *NonFungibleTokenMintWithTxMetadata) GetTx() Transaction`

GetTx returns the Tx field if non-nil, zero value otherwise.

### GetTxOk

`func (o *NonFungibleTokenMintWithTxMetadata) GetTxOk() (*Transaction, bool)`

GetTxOk returns a tuple with the Tx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTx

`func (o *NonFungibleTokenMintWithTxMetadata) SetTx(v Transaction)`

SetTx sets Tx field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


