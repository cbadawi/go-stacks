# SigningPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | [DEPRECATED by account_identifier in v1.4.4] The network-specific address of the account that should sign the payload. | [optional] 
**AccountIdentifier** | Pointer to [**RosettaAccount**](RosettaAccount.md) |  | [optional] 
**HexBytes** | **string** |  | 
**SignatureType** | Pointer to **string** | SignatureType is the type of a cryptographic signature. | [optional] 

## Methods

### NewSigningPayload

`func NewSigningPayload(hexBytes string, ) *SigningPayload`

NewSigningPayload instantiates a new SigningPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSigningPayloadWithDefaults

`func NewSigningPayloadWithDefaults() *SigningPayload`

NewSigningPayloadWithDefaults instantiates a new SigningPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *SigningPayload) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *SigningPayload) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *SigningPayload) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *SigningPayload) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAccountIdentifier

`func (o *SigningPayload) GetAccountIdentifier() RosettaAccount`

GetAccountIdentifier returns the AccountIdentifier field if non-nil, zero value otherwise.

### GetAccountIdentifierOk

`func (o *SigningPayload) GetAccountIdentifierOk() (*RosettaAccount, bool)`

GetAccountIdentifierOk returns a tuple with the AccountIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountIdentifier

`func (o *SigningPayload) SetAccountIdentifier(v RosettaAccount)`

SetAccountIdentifier sets AccountIdentifier field to given value.

### HasAccountIdentifier

`func (o *SigningPayload) HasAccountIdentifier() bool`

HasAccountIdentifier returns a boolean if a field has been set.

### GetHexBytes

`func (o *SigningPayload) GetHexBytes() string`

GetHexBytes returns the HexBytes field if non-nil, zero value otherwise.

### GetHexBytesOk

`func (o *SigningPayload) GetHexBytesOk() (*string, bool)`

GetHexBytesOk returns a tuple with the HexBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHexBytes

`func (o *SigningPayload) SetHexBytes(v string)`

SetHexBytes sets HexBytes field to given value.


### GetSignatureType

`func (o *SigningPayload) GetSignatureType() string`

GetSignatureType returns the SignatureType field if non-nil, zero value otherwise.

### GetSignatureTypeOk

`func (o *SigningPayload) GetSignatureTypeOk() (*string, bool)`

GetSignatureTypeOk returns a tuple with the SignatureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureType

`func (o *SigningPayload) SetSignatureType(v string)`

SetSignatureType sets SignatureType field to given value.

### HasSignatureType

`func (o *SigningPayload) HasSignatureType() bool`

HasSignatureType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


