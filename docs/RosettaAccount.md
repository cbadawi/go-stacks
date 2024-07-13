# RosettaAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | The address may be a cryptographic public key (or some encoding of it) or a provided username. | 
**SubAccount** | Pointer to [**RosettaSubAccount**](RosettaSubAccount.md) |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Blockchains that utilize a username model (where the address is not a derivative of a cryptographic public key) should specify the public key(s) owned by the address in metadata. | [optional] 

## Methods

### NewRosettaAccount

`func NewRosettaAccount(address string, ) *RosettaAccount`

NewRosettaAccount instantiates a new RosettaAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRosettaAccountWithDefaults

`func NewRosettaAccountWithDefaults() *RosettaAccount`

NewRosettaAccountWithDefaults instantiates a new RosettaAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *RosettaAccount) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *RosettaAccount) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *RosettaAccount) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetSubAccount

`func (o *RosettaAccount) GetSubAccount() RosettaSubAccount`

GetSubAccount returns the SubAccount field if non-nil, zero value otherwise.

### GetSubAccountOk

`func (o *RosettaAccount) GetSubAccountOk() (*RosettaSubAccount, bool)`

GetSubAccountOk returns a tuple with the SubAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubAccount

`func (o *RosettaAccount) SetSubAccount(v RosettaSubAccount)`

SetSubAccount sets SubAccount field to given value.

### HasSubAccount

`func (o *RosettaAccount) HasSubAccount() bool`

HasSubAccount returns a boolean if a field has been set.

### GetMetadata

`func (o *RosettaAccount) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RosettaAccount) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RosettaAccount) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RosettaAccount) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


