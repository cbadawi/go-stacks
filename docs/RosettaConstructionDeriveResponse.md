# RosettaConstructionDeriveResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | [DEPRECATED by account_identifier in v1.4.4] Address in network-specific format. | [optional] 
**AccountIdentifier** | Pointer to [**RosettaAccountIdentifier**](RosettaAccountIdentifier.md) |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewRosettaConstructionDeriveResponse

`func NewRosettaConstructionDeriveResponse() *RosettaConstructionDeriveResponse`

NewRosettaConstructionDeriveResponse instantiates a new RosettaConstructionDeriveResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRosettaConstructionDeriveResponseWithDefaults

`func NewRosettaConstructionDeriveResponseWithDefaults() *RosettaConstructionDeriveResponse`

NewRosettaConstructionDeriveResponseWithDefaults instantiates a new RosettaConstructionDeriveResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *RosettaConstructionDeriveResponse) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *RosettaConstructionDeriveResponse) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *RosettaConstructionDeriveResponse) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *RosettaConstructionDeriveResponse) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAccountIdentifier

`func (o *RosettaConstructionDeriveResponse) GetAccountIdentifier() RosettaAccountIdentifier`

GetAccountIdentifier returns the AccountIdentifier field if non-nil, zero value otherwise.

### GetAccountIdentifierOk

`func (o *RosettaConstructionDeriveResponse) GetAccountIdentifierOk() (*RosettaAccountIdentifier, bool)`

GetAccountIdentifierOk returns a tuple with the AccountIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountIdentifier

`func (o *RosettaConstructionDeriveResponse) SetAccountIdentifier(v RosettaAccountIdentifier)`

SetAccountIdentifier sets AccountIdentifier field to given value.

### HasAccountIdentifier

`func (o *RosettaConstructionDeriveResponse) HasAccountIdentifier() bool`

HasAccountIdentifier returns a boolean if a field has been set.

### GetMetadata

`func (o *RosettaConstructionDeriveResponse) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RosettaConstructionDeriveResponse) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RosettaConstructionDeriveResponse) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RosettaConstructionDeriveResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


