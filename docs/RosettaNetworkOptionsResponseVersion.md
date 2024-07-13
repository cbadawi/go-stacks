# RosettaNetworkOptionsResponseVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RosettaVersion** | **string** | The rosetta_version is the version of the Rosetta interface the implementation adheres to. This can be useful for clients looking to reliably parse responses. | 
**NodeVersion** | **string** | The node_version is the canonical version of the node runtime. This can help clients manage deployments. | 
**MiddlewareVersion** | Pointer to **string** | When a middleware server is used to adhere to the Rosetta interface, it should return its version here. This can help clients manage deployments. | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Any other information that may be useful about versioning of dependent services should be returned here. | [optional] 

## Methods

### NewRosettaNetworkOptionsResponseVersion

`func NewRosettaNetworkOptionsResponseVersion(rosettaVersion string, nodeVersion string, ) *RosettaNetworkOptionsResponseVersion`

NewRosettaNetworkOptionsResponseVersion instantiates a new RosettaNetworkOptionsResponseVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRosettaNetworkOptionsResponseVersionWithDefaults

`func NewRosettaNetworkOptionsResponseVersionWithDefaults() *RosettaNetworkOptionsResponseVersion`

NewRosettaNetworkOptionsResponseVersionWithDefaults instantiates a new RosettaNetworkOptionsResponseVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRosettaVersion

`func (o *RosettaNetworkOptionsResponseVersion) GetRosettaVersion() string`

GetRosettaVersion returns the RosettaVersion field if non-nil, zero value otherwise.

### GetRosettaVersionOk

`func (o *RosettaNetworkOptionsResponseVersion) GetRosettaVersionOk() (*string, bool)`

GetRosettaVersionOk returns a tuple with the RosettaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRosettaVersion

`func (o *RosettaNetworkOptionsResponseVersion) SetRosettaVersion(v string)`

SetRosettaVersion sets RosettaVersion field to given value.


### GetNodeVersion

`func (o *RosettaNetworkOptionsResponseVersion) GetNodeVersion() string`

GetNodeVersion returns the NodeVersion field if non-nil, zero value otherwise.

### GetNodeVersionOk

`func (o *RosettaNetworkOptionsResponseVersion) GetNodeVersionOk() (*string, bool)`

GetNodeVersionOk returns a tuple with the NodeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeVersion

`func (o *RosettaNetworkOptionsResponseVersion) SetNodeVersion(v string)`

SetNodeVersion sets NodeVersion field to given value.


### GetMiddlewareVersion

`func (o *RosettaNetworkOptionsResponseVersion) GetMiddlewareVersion() string`

GetMiddlewareVersion returns the MiddlewareVersion field if non-nil, zero value otherwise.

### GetMiddlewareVersionOk

`func (o *RosettaNetworkOptionsResponseVersion) GetMiddlewareVersionOk() (*string, bool)`

GetMiddlewareVersionOk returns a tuple with the MiddlewareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddlewareVersion

`func (o *RosettaNetworkOptionsResponseVersion) SetMiddlewareVersion(v string)`

SetMiddlewareVersion sets MiddlewareVersion field to given value.

### HasMiddlewareVersion

`func (o *RosettaNetworkOptionsResponseVersion) HasMiddlewareVersion() bool`

HasMiddlewareVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *RosettaNetworkOptionsResponseVersion) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RosettaNetworkOptionsResponseVersion) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RosettaNetworkOptionsResponseVersion) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RosettaNetworkOptionsResponseVersion) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


