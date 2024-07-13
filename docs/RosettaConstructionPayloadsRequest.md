# RosettaConstructionPayloadsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkIdentifier** | [**NetworkIdentifier**](NetworkIdentifier.md) |  | 
**Operations** | [**[]RosettaOperation**](RosettaOperation.md) |  | 
**PublicKeys** | Pointer to [**[]RosettaPublicKey**](RosettaPublicKey.md) |  | [optional] 
**Metadata** | Pointer to [**RosettaConstructionMetadataResponseMetadata**](RosettaConstructionMetadataResponseMetadata.md) |  | [optional] 

## Methods

### NewRosettaConstructionPayloadsRequest

`func NewRosettaConstructionPayloadsRequest(networkIdentifier NetworkIdentifier, operations []RosettaOperation, ) *RosettaConstructionPayloadsRequest`

NewRosettaConstructionPayloadsRequest instantiates a new RosettaConstructionPayloadsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRosettaConstructionPayloadsRequestWithDefaults

`func NewRosettaConstructionPayloadsRequestWithDefaults() *RosettaConstructionPayloadsRequest`

NewRosettaConstructionPayloadsRequestWithDefaults instantiates a new RosettaConstructionPayloadsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkIdentifier

`func (o *RosettaConstructionPayloadsRequest) GetNetworkIdentifier() NetworkIdentifier`

GetNetworkIdentifier returns the NetworkIdentifier field if non-nil, zero value otherwise.

### GetNetworkIdentifierOk

`func (o *RosettaConstructionPayloadsRequest) GetNetworkIdentifierOk() (*NetworkIdentifier, bool)`

GetNetworkIdentifierOk returns a tuple with the NetworkIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIdentifier

`func (o *RosettaConstructionPayloadsRequest) SetNetworkIdentifier(v NetworkIdentifier)`

SetNetworkIdentifier sets NetworkIdentifier field to given value.


### GetOperations

`func (o *RosettaConstructionPayloadsRequest) GetOperations() []RosettaOperation`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *RosettaConstructionPayloadsRequest) GetOperationsOk() (*[]RosettaOperation, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *RosettaConstructionPayloadsRequest) SetOperations(v []RosettaOperation)`

SetOperations sets Operations field to given value.


### GetPublicKeys

`func (o *RosettaConstructionPayloadsRequest) GetPublicKeys() []RosettaPublicKey`

GetPublicKeys returns the PublicKeys field if non-nil, zero value otherwise.

### GetPublicKeysOk

`func (o *RosettaConstructionPayloadsRequest) GetPublicKeysOk() (*[]RosettaPublicKey, bool)`

GetPublicKeysOk returns a tuple with the PublicKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeys

`func (o *RosettaConstructionPayloadsRequest) SetPublicKeys(v []RosettaPublicKey)`

SetPublicKeys sets PublicKeys field to given value.

### HasPublicKeys

`func (o *RosettaConstructionPayloadsRequest) HasPublicKeys() bool`

HasPublicKeys returns a boolean if a field has been set.

### GetMetadata

`func (o *RosettaConstructionPayloadsRequest) GetMetadata() RosettaConstructionMetadataResponseMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RosettaConstructionPayloadsRequest) GetMetadataOk() (*RosettaConstructionMetadataResponseMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RosettaConstructionPayloadsRequest) SetMetadata(v RosettaConstructionMetadataResponseMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RosettaConstructionPayloadsRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


