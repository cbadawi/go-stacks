# TokenTransferTransactionMetadataTokenTransfer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecipientAddress** | **string** |  | 
**Amount** | **string** | Transfer amount as Integer string (64-bit unsigned integer) | 
**Memo** | **string** | Hex encoded arbitrary message, up to 34 bytes length (should try decoding to an ASCII string) | 

## Methods

### NewTokenTransferTransactionMetadataTokenTransfer

`func NewTokenTransferTransactionMetadataTokenTransfer(recipientAddress string, amount string, memo string, ) *TokenTransferTransactionMetadataTokenTransfer`

NewTokenTransferTransactionMetadataTokenTransfer instantiates a new TokenTransferTransactionMetadataTokenTransfer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenTransferTransactionMetadataTokenTransferWithDefaults

`func NewTokenTransferTransactionMetadataTokenTransferWithDefaults() *TokenTransferTransactionMetadataTokenTransfer`

NewTokenTransferTransactionMetadataTokenTransferWithDefaults instantiates a new TokenTransferTransactionMetadataTokenTransfer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecipientAddress

`func (o *TokenTransferTransactionMetadataTokenTransfer) GetRecipientAddress() string`

GetRecipientAddress returns the RecipientAddress field if non-nil, zero value otherwise.

### GetRecipientAddressOk

`func (o *TokenTransferTransactionMetadataTokenTransfer) GetRecipientAddressOk() (*string, bool)`

GetRecipientAddressOk returns a tuple with the RecipientAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientAddress

`func (o *TokenTransferTransactionMetadataTokenTransfer) SetRecipientAddress(v string)`

SetRecipientAddress sets RecipientAddress field to given value.


### GetAmount

`func (o *TokenTransferTransactionMetadataTokenTransfer) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TokenTransferTransactionMetadataTokenTransfer) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TokenTransferTransactionMetadataTokenTransfer) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetMemo

`func (o *TokenTransferTransactionMetadataTokenTransfer) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *TokenTransferTransactionMetadataTokenTransfer) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *TokenTransferTransactionMetadataTokenTransfer) SetMemo(v string)`

SetMemo sets Memo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


