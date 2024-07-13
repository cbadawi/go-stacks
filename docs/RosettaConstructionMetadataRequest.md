# RosettaConstructionMetadataRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkIdentifier** | [**NetworkIdentifier**](NetworkIdentifier.md) |  | 
**Options** | [**RosettaOptions**](RosettaOptions.md) |  | 
**PublicKeys** | Pointer to [**[]RosettaPublicKey**](RosettaPublicKey.md) |  | [optional] 

## Methods

### NewRosettaConstructionMetadataRequest

`func NewRosettaConstructionMetadataRequest(networkIdentifier NetworkIdentifier, options RosettaOptions, ) *RosettaConstructionMetadataRequest`

NewRosettaConstructionMetadataRequest instantiates a new RosettaConstructionMetadataRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRosettaConstructionMetadataRequestWithDefaults

`func NewRosettaConstructionMetadataRequestWithDefaults() *RosettaConstructionMetadataRequest`

NewRosettaConstructionMetadataRequestWithDefaults instantiates a new RosettaConstructionMetadataRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkIdentifier

`func (o *RosettaConstructionMetadataRequest) GetNetworkIdentifier() NetworkIdentifier`

GetNetworkIdentifier returns the NetworkIdentifier field if non-nil, zero value otherwise.

### GetNetworkIdentifierOk

`func (o *RosettaConstructionMetadataRequest) GetNetworkIdentifierOk() (*NetworkIdentifier, bool)`

GetNetworkIdentifierOk returns a tuple with the NetworkIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIdentifier

`func (o *RosettaConstructionMetadataRequest) SetNetworkIdentifier(v NetworkIdentifier)`

SetNetworkIdentifier sets NetworkIdentifier field to given value.


### GetOptions

`func (o *RosettaConstructionMetadataRequest) GetOptions() RosettaOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *RosettaConstructionMetadataRequest) GetOptionsOk() (*RosettaOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *RosettaConstructionMetadataRequest) SetOptions(v RosettaOptions)`

SetOptions sets Options field to given value.


### GetPublicKeys

`func (o *RosettaConstructionMetadataRequest) GetPublicKeys() []RosettaPublicKey`

GetPublicKeys returns the PublicKeys field if non-nil, zero value otherwise.

### GetPublicKeysOk

`func (o *RosettaConstructionMetadataRequest) GetPublicKeysOk() (*[]RosettaPublicKey, bool)`

GetPublicKeysOk returns a tuple with the PublicKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeys

`func (o *RosettaConstructionMetadataRequest) SetPublicKeys(v []RosettaPublicKey)`

SetPublicKeys sets PublicKeys field to given value.

### HasPublicKeys

`func (o *RosettaConstructionMetadataRequest) HasPublicKeys() bool`

HasPublicKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


