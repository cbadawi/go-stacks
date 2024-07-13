# NetworkIdentifierSubNetworkIdentifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | **string** | Network name | 
**Metadata** | Pointer to [**NetworkIdentifierSubNetworkIdentifierMetadata**](NetworkIdentifierSubNetworkIdentifierMetadata.md) |  | [optional] 

## Methods

### NewNetworkIdentifierSubNetworkIdentifier

`func NewNetworkIdentifierSubNetworkIdentifier(network string, ) *NetworkIdentifierSubNetworkIdentifier`

NewNetworkIdentifierSubNetworkIdentifier instantiates a new NetworkIdentifierSubNetworkIdentifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkIdentifierSubNetworkIdentifierWithDefaults

`func NewNetworkIdentifierSubNetworkIdentifierWithDefaults() *NetworkIdentifierSubNetworkIdentifier`

NewNetworkIdentifierSubNetworkIdentifierWithDefaults instantiates a new NetworkIdentifierSubNetworkIdentifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *NetworkIdentifierSubNetworkIdentifier) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *NetworkIdentifierSubNetworkIdentifier) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *NetworkIdentifierSubNetworkIdentifier) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetMetadata

`func (o *NetworkIdentifierSubNetworkIdentifier) GetMetadata() NetworkIdentifierSubNetworkIdentifierMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NetworkIdentifierSubNetworkIdentifier) GetMetadataOk() (*NetworkIdentifierSubNetworkIdentifierMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NetworkIdentifierSubNetworkIdentifier) SetMetadata(v NetworkIdentifierSubNetworkIdentifierMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *NetworkIdentifierSubNetworkIdentifier) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


