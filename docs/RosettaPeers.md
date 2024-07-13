# RosettaPeers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PeerId** | **string** | peer id | 
**Metadata** | Pointer to **map[string]interface{}** | meta data | [optional] 

## Methods

### NewRosettaPeers

`func NewRosettaPeers(peerId string, ) *RosettaPeers`

NewRosettaPeers instantiates a new RosettaPeers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRosettaPeersWithDefaults

`func NewRosettaPeersWithDefaults() *RosettaPeers`

NewRosettaPeersWithDefaults instantiates a new RosettaPeers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeerId

`func (o *RosettaPeers) GetPeerId() string`

GetPeerId returns the PeerId field if non-nil, zero value otherwise.

### GetPeerIdOk

`func (o *RosettaPeers) GetPeerIdOk() (*string, bool)`

GetPeerIdOk returns a tuple with the PeerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerId

`func (o *RosettaPeers) SetPeerId(v string)`

SetPeerId sets PeerId field to given value.


### GetMetadata

`func (o *RosettaPeers) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RosettaPeers) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RosettaPeers) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RosettaPeers) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


