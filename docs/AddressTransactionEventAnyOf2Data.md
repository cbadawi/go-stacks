# AddressTransactionEventAnyOf2Data

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**AssetIdentifier** | **string** | Non Fungible Token asset identifier. | 
**Value** | [**AddressTransactionEventAnyOf2DataValue**](AddressTransactionEventAnyOf2DataValue.md) |  | 
**Sender** | Pointer to **string** | Principal that sent the asset. | [optional] 
**Recipient** | Pointer to **string** | Principal that received the asset. | [optional] 

## Methods

### NewAddressTransactionEventAnyOf2Data

`func NewAddressTransactionEventAnyOf2Data(type_ string, assetIdentifier string, value AddressTransactionEventAnyOf2DataValue, ) *AddressTransactionEventAnyOf2Data`

NewAddressTransactionEventAnyOf2Data instantiates a new AddressTransactionEventAnyOf2Data object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressTransactionEventAnyOf2DataWithDefaults

`func NewAddressTransactionEventAnyOf2DataWithDefaults() *AddressTransactionEventAnyOf2Data`

NewAddressTransactionEventAnyOf2DataWithDefaults instantiates a new AddressTransactionEventAnyOf2Data object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AddressTransactionEventAnyOf2Data) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddressTransactionEventAnyOf2Data) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddressTransactionEventAnyOf2Data) SetType(v string)`

SetType sets Type field to given value.


### GetAssetIdentifier

`func (o *AddressTransactionEventAnyOf2Data) GetAssetIdentifier() string`

GetAssetIdentifier returns the AssetIdentifier field if non-nil, zero value otherwise.

### GetAssetIdentifierOk

`func (o *AddressTransactionEventAnyOf2Data) GetAssetIdentifierOk() (*string, bool)`

GetAssetIdentifierOk returns a tuple with the AssetIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetIdentifier

`func (o *AddressTransactionEventAnyOf2Data) SetAssetIdentifier(v string)`

SetAssetIdentifier sets AssetIdentifier field to given value.


### GetValue

`func (o *AddressTransactionEventAnyOf2Data) GetValue() AddressTransactionEventAnyOf2DataValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AddressTransactionEventAnyOf2Data) GetValueOk() (*AddressTransactionEventAnyOf2DataValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AddressTransactionEventAnyOf2Data) SetValue(v AddressTransactionEventAnyOf2DataValue)`

SetValue sets Value field to given value.


### GetSender

`func (o *AddressTransactionEventAnyOf2Data) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *AddressTransactionEventAnyOf2Data) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *AddressTransactionEventAnyOf2Data) SetSender(v string)`

SetSender sets Sender field to given value.

### HasSender

`func (o *AddressTransactionEventAnyOf2Data) HasSender() bool`

HasSender returns a boolean if a field has been set.

### GetRecipient

`func (o *AddressTransactionEventAnyOf2Data) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *AddressTransactionEventAnyOf2Data) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *AddressTransactionEventAnyOf2Data) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *AddressTransactionEventAnyOf2Data) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


