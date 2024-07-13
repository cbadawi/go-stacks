# CoinbaseTransactionMetadataCoinbasePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | **string** | Hex encoded 32-byte scratch space for block leader&#39;s use | 
**AltRecipient** | Pointer to **NullableString** | A principal that will receive the miner rewards for this coinbase transaction. Can be either a standard principal or contract principal. Only specified for &#x60;coinbase-to-alt-recipient&#x60; transaction types, otherwise null. | [optional] 
**VrfProof** | Pointer to **NullableString** | Hex encoded 80-byte VRF proof | [optional] 

## Methods

### NewCoinbaseTransactionMetadataCoinbasePayload

`func NewCoinbaseTransactionMetadataCoinbasePayload(data string, ) *CoinbaseTransactionMetadataCoinbasePayload`

NewCoinbaseTransactionMetadataCoinbasePayload instantiates a new CoinbaseTransactionMetadataCoinbasePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoinbaseTransactionMetadataCoinbasePayloadWithDefaults

`func NewCoinbaseTransactionMetadataCoinbasePayloadWithDefaults() *CoinbaseTransactionMetadataCoinbasePayload`

NewCoinbaseTransactionMetadataCoinbasePayloadWithDefaults instantiates a new CoinbaseTransactionMetadataCoinbasePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CoinbaseTransactionMetadataCoinbasePayload) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CoinbaseTransactionMetadataCoinbasePayload) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CoinbaseTransactionMetadataCoinbasePayload) SetData(v string)`

SetData sets Data field to given value.


### GetAltRecipient

`func (o *CoinbaseTransactionMetadataCoinbasePayload) GetAltRecipient() string`

GetAltRecipient returns the AltRecipient field if non-nil, zero value otherwise.

### GetAltRecipientOk

`func (o *CoinbaseTransactionMetadataCoinbasePayload) GetAltRecipientOk() (*string, bool)`

GetAltRecipientOk returns a tuple with the AltRecipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltRecipient

`func (o *CoinbaseTransactionMetadataCoinbasePayload) SetAltRecipient(v string)`

SetAltRecipient sets AltRecipient field to given value.

### HasAltRecipient

`func (o *CoinbaseTransactionMetadataCoinbasePayload) HasAltRecipient() bool`

HasAltRecipient returns a boolean if a field has been set.

### SetAltRecipientNil

`func (o *CoinbaseTransactionMetadataCoinbasePayload) SetAltRecipientNil(b bool)`

 SetAltRecipientNil sets the value for AltRecipient to be an explicit nil

### UnsetAltRecipient
`func (o *CoinbaseTransactionMetadataCoinbasePayload) UnsetAltRecipient()`

UnsetAltRecipient ensures that no value is present for AltRecipient, not even an explicit nil
### GetVrfProof

`func (o *CoinbaseTransactionMetadataCoinbasePayload) GetVrfProof() string`

GetVrfProof returns the VrfProof field if non-nil, zero value otherwise.

### GetVrfProofOk

`func (o *CoinbaseTransactionMetadataCoinbasePayload) GetVrfProofOk() (*string, bool)`

GetVrfProofOk returns a tuple with the VrfProof field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfProof

`func (o *CoinbaseTransactionMetadataCoinbasePayload) SetVrfProof(v string)`

SetVrfProof sets VrfProof field to given value.

### HasVrfProof

`func (o *CoinbaseTransactionMetadataCoinbasePayload) HasVrfProof() bool`

HasVrfProof returns a boolean if a field has been set.

### SetVrfProofNil

`func (o *CoinbaseTransactionMetadataCoinbasePayload) SetVrfProofNil(b bool)`

 SetVrfProofNil sets the value for VrfProof to be an explicit nil

### UnsetVrfProof
`func (o *CoinbaseTransactionMetadataCoinbasePayload) UnsetVrfProof()`

UnsetVrfProof ensures that no value is present for VrfProof, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


