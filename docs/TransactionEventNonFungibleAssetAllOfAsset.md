# TransactionEventNonFungibleAssetAllOfAsset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetEventType** | **string** |  | 
**AssetId** | **string** |  | 
**Sender** | **string** |  | 
**Recipient** | **string** |  | 
**Value** | [**PostConditionNonFungibleAllOfAssetValue**](PostConditionNonFungibleAllOfAssetValue.md) |  | 

## Methods

### NewTransactionEventNonFungibleAssetAllOfAsset

`func NewTransactionEventNonFungibleAssetAllOfAsset(assetEventType string, assetId string, sender string, recipient string, value PostConditionNonFungibleAllOfAssetValue, ) *TransactionEventNonFungibleAssetAllOfAsset`

NewTransactionEventNonFungibleAssetAllOfAsset instantiates a new TransactionEventNonFungibleAssetAllOfAsset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionEventNonFungibleAssetAllOfAssetWithDefaults

`func NewTransactionEventNonFungibleAssetAllOfAssetWithDefaults() *TransactionEventNonFungibleAssetAllOfAsset`

NewTransactionEventNonFungibleAssetAllOfAssetWithDefaults instantiates a new TransactionEventNonFungibleAssetAllOfAsset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetEventType

`func (o *TransactionEventNonFungibleAssetAllOfAsset) GetAssetEventType() string`

GetAssetEventType returns the AssetEventType field if non-nil, zero value otherwise.

### GetAssetEventTypeOk

`func (o *TransactionEventNonFungibleAssetAllOfAsset) GetAssetEventTypeOk() (*string, bool)`

GetAssetEventTypeOk returns a tuple with the AssetEventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetEventType

`func (o *TransactionEventNonFungibleAssetAllOfAsset) SetAssetEventType(v string)`

SetAssetEventType sets AssetEventType field to given value.


### GetAssetId

`func (o *TransactionEventNonFungibleAssetAllOfAsset) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *TransactionEventNonFungibleAssetAllOfAsset) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *TransactionEventNonFungibleAssetAllOfAsset) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.


### GetSender

`func (o *TransactionEventNonFungibleAssetAllOfAsset) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *TransactionEventNonFungibleAssetAllOfAsset) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *TransactionEventNonFungibleAssetAllOfAsset) SetSender(v string)`

SetSender sets Sender field to given value.


### GetRecipient

`func (o *TransactionEventNonFungibleAssetAllOfAsset) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *TransactionEventNonFungibleAssetAllOfAsset) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *TransactionEventNonFungibleAssetAllOfAsset) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.


### GetValue

`func (o *TransactionEventNonFungibleAssetAllOfAsset) GetValue() PostConditionNonFungibleAllOfAssetValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TransactionEventNonFungibleAssetAllOfAsset) GetValueOk() (*PostConditionNonFungibleAllOfAssetValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TransactionEventNonFungibleAssetAllOfAsset) SetValue(v PostConditionNonFungibleAllOfAssetValue)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


