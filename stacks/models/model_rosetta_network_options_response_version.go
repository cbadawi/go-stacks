package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the RosettaNetworkOptionsResponseVersion type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RosettaNetworkOptionsResponseVersion{}

// RosettaNetworkOptionsResponseVersion The Version object is utilized to inform the client of the versions of different components of the Rosetta implementation.
type RosettaNetworkOptionsResponseVersion struct {
	// The rosetta_version is the version of the Rosetta interface the implementation adheres to. This can be useful for clients looking to reliably parse responses.
	RosettaVersion string `json:"rosetta_version"`
	// The node_version is the canonical version of the node runtime. This can help clients manage deployments.
	NodeVersion string `json:"node_version"`
	// When a middleware server is used to adhere to the Rosetta interface, it should return its version here. This can help clients manage deployments.
	MiddlewareVersion *string `json:"middleware_version,omitempty"`
	// Any other information that may be useful about versioning of dependent services should be returned here.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

type _RosettaNetworkOptionsResponseVersion RosettaNetworkOptionsResponseVersion

// NewRosettaNetworkOptionsResponseVersion instantiates a new RosettaNetworkOptionsResponseVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRosettaNetworkOptionsResponseVersion(rosettaVersion string, nodeVersion string) *RosettaNetworkOptionsResponseVersion {
	this := RosettaNetworkOptionsResponseVersion{}
	this.RosettaVersion = rosettaVersion
	this.NodeVersion = nodeVersion
	return &this
}

// NewRosettaNetworkOptionsResponseVersionWithDefaults instantiates a new RosettaNetworkOptionsResponseVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRosettaNetworkOptionsResponseVersionWithDefaults() *RosettaNetworkOptionsResponseVersion {
	this := RosettaNetworkOptionsResponseVersion{}
	return &this
}

// GetRosettaVersion returns the RosettaVersion field value
func (o *RosettaNetworkOptionsResponseVersion) GetRosettaVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RosettaVersion
}

// GetRosettaVersionOk returns a tuple with the RosettaVersion field value
// and a boolean to check if the value has been set.
func (o *RosettaNetworkOptionsResponseVersion) GetRosettaVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RosettaVersion, true
}

// SetRosettaVersion sets field value
func (o *RosettaNetworkOptionsResponseVersion) SetRosettaVersion(v string) {
	o.RosettaVersion = v
}

// GetNodeVersion returns the NodeVersion field value
func (o *RosettaNetworkOptionsResponseVersion) GetNodeVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeVersion
}

// GetNodeVersionOk returns a tuple with the NodeVersion field value
// and a boolean to check if the value has been set.
func (o *RosettaNetworkOptionsResponseVersion) GetNodeVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeVersion, true
}

// SetNodeVersion sets field value
func (o *RosettaNetworkOptionsResponseVersion) SetNodeVersion(v string) {
	o.NodeVersion = v
}

// GetMiddlewareVersion returns the MiddlewareVersion field value if set, zero value otherwise.
func (o *RosettaNetworkOptionsResponseVersion) GetMiddlewareVersion() string {
	if o == nil || utils.IsNil(o.MiddlewareVersion) {
		var ret string
		return ret
	}
	return *o.MiddlewareVersion
}

// GetMiddlewareVersionOk returns a tuple with the MiddlewareVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaNetworkOptionsResponseVersion) GetMiddlewareVersionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.MiddlewareVersion) {
		return nil, false
	}
	return o.MiddlewareVersion, true
}

// HasMiddlewareVersion returns a boolean if a field has been set.
func (o *RosettaNetworkOptionsResponseVersion) HasMiddlewareVersion() bool {
	if o != nil && !utils.IsNil(o.MiddlewareVersion) {
		return true
	}

	return false
}

// SetMiddlewareVersion gets a reference to the given string and assigns it to the MiddlewareVersion field.
func (o *RosettaNetworkOptionsResponseVersion) SetMiddlewareVersion(v string) {
	o.MiddlewareVersion = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *RosettaNetworkOptionsResponseVersion) GetMetadata() map[string]interface{} {
	if o == nil || utils.IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaNetworkOptionsResponseVersion) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || utils.IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *RosettaNetworkOptionsResponseVersion) HasMetadata() bool {
	if o != nil && !utils.IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *RosettaNetworkOptionsResponseVersion) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o RosettaNetworkOptionsResponseVersion) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RosettaNetworkOptionsResponseVersion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["rosetta_version"] = o.RosettaVersion
	toSerialize["node_version"] = o.NodeVersion
	if !utils.IsNil(o.MiddlewareVersion) {
		toSerialize["middleware_version"] = o.MiddlewareVersion
	}
	if !utils.IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

func (o *RosettaNetworkOptionsResponseVersion) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"rosetta_version",
		"node_version",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varRosettaNetworkOptionsResponseVersion := _RosettaNetworkOptionsResponseVersion{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRosettaNetworkOptionsResponseVersion)

	if err != nil {
		return err
	}

	*o = RosettaNetworkOptionsResponseVersion(varRosettaNetworkOptionsResponseVersion)

	return err
}

type NullableRosettaNetworkOptionsResponseVersion struct {
	value *RosettaNetworkOptionsResponseVersion
	isSet bool
}

func (v NullableRosettaNetworkOptionsResponseVersion) Get() *RosettaNetworkOptionsResponseVersion {
	return v.value
}

func (v *NullableRosettaNetworkOptionsResponseVersion) Set(val *RosettaNetworkOptionsResponseVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableRosettaNetworkOptionsResponseVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableRosettaNetworkOptionsResponseVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRosettaNetworkOptionsResponseVersion(val *RosettaNetworkOptionsResponseVersion) *NullableRosettaNetworkOptionsResponseVersion {
	return &NullableRosettaNetworkOptionsResponseVersion{value: val, isSet: true}
}

func (v NullableRosettaNetworkOptionsResponseVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRosettaNetworkOptionsResponseVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
