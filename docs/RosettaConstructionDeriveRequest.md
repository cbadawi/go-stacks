# RosettaConstructionDeriveRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkIdentifier** | [**NetworkIdentifier**](NetworkIdentifier.md) |  | 
**PublicKey** | [**RosettaPublicKey**](RosettaPublicKey.md) |  | 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewRosettaConstructionDeriveRequest

`func NewRosettaConstructionDeriveRequest(networkIdentifier NetworkIdentifier, publicKey RosettaPublicKey, ) *RosettaConstructionDeriveRequest`

NewRosettaConstructionDeriveRequest instantiates a new RosettaConstructionDeriveRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRosettaConstructionDeriveRequestWithDefaults

`func NewRosettaConstructionDeriveRequestWithDefaults() *RosettaConstructionDeriveRequest`

NewRosettaConstructionDeriveRequestWithDefaults instantiates a new RosettaConstructionDeriveRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkIdentifier

`func (o *RosettaConstructionDeriveRequest) GetNetworkIdentifier() NetworkIdentifier`

GetNetworkIdentifier returns the NetworkIdentifier field if non-nil, zero value otherwise.

### GetNetworkIdentifierOk

`func (o *RosettaConstructionDeriveRequest) GetNetworkIdentifierOk() (*NetworkIdentifier, bool)`

GetNetworkIdentifierOk returns a tuple with the NetworkIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIdentifier

`func (o *RosettaConstructionDeriveRequest) SetNetworkIdentifier(v NetworkIdentifier)`

SetNetworkIdentifier sets NetworkIdentifier field to given value.


### GetPublicKey

`func (o *RosettaConstructionDeriveRequest) GetPublicKey() RosettaPublicKey`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *RosettaConstructionDeriveRequest) GetPublicKeyOk() (*RosettaPublicKey, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *RosettaConstructionDeriveRequest) SetPublicKey(v RosettaPublicKey)`

SetPublicKey sets PublicKey field to given value.


### GetMetadata

`func (o *RosettaConstructionDeriveRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RosettaConstructionDeriveRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RosettaConstructionDeriveRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RosettaConstructionDeriveRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


