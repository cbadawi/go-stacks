package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the RosettaNetworkOptionsResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RosettaNetworkOptionsResponse{}

// RosettaNetworkOptionsResponse NetworkOptionsResponse contains information about the versioning of the node and the allowed operation statuses, operation types, and errors.
type RosettaNetworkOptionsResponse struct {
	Version RosettaNetworkOptionsResponseVersion `json:"version"`
	Allow   RosettaNetworkOptionsResponseAllow   `json:"allow"`
}

type _RosettaNetworkOptionsResponse RosettaNetworkOptionsResponse

// NewRosettaNetworkOptionsResponse instantiates a new RosettaNetworkOptionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRosettaNetworkOptionsResponse(version RosettaNetworkOptionsResponseVersion, allow RosettaNetworkOptionsResponseAllow) *RosettaNetworkOptionsResponse {
	this := RosettaNetworkOptionsResponse{}
	this.Version = version
	this.Allow = allow
	return &this
}

// NewRosettaNetworkOptionsResponseWithDefaults instantiates a new RosettaNetworkOptionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRosettaNetworkOptionsResponseWithDefaults() *RosettaNetworkOptionsResponse {
	this := RosettaNetworkOptionsResponse{}
	return &this
}

// GetVersion returns the Version field value
func (o *RosettaNetworkOptionsResponse) GetVersion() RosettaNetworkOptionsResponseVersion {
	if o == nil {
		var ret RosettaNetworkOptionsResponseVersion
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *RosettaNetworkOptionsResponse) GetVersionOk() (*RosettaNetworkOptionsResponseVersion, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *RosettaNetworkOptionsResponse) SetVersion(v RosettaNetworkOptionsResponseVersion) {
	o.Version = v
}

// GetAllow returns the Allow field value
func (o *RosettaNetworkOptionsResponse) GetAllow() RosettaNetworkOptionsResponseAllow {
	if o == nil {
		var ret RosettaNetworkOptionsResponseAllow
		return ret
	}

	return o.Allow
}

// GetAllowOk returns a tuple with the Allow field value
// and a boolean to check if the value has been set.
func (o *RosettaNetworkOptionsResponse) GetAllowOk() (*RosettaNetworkOptionsResponseAllow, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Allow, true
}

// SetAllow sets field value
func (o *RosettaNetworkOptionsResponse) SetAllow(v RosettaNetworkOptionsResponseAllow) {
	o.Allow = v
}

func (o RosettaNetworkOptionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RosettaNetworkOptionsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["version"] = o.Version
	toSerialize["allow"] = o.Allow
	return toSerialize, nil
}

func (o *RosettaNetworkOptionsResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"version",
		"allow",
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

	varRosettaNetworkOptionsResponse := _RosettaNetworkOptionsResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRosettaNetworkOptionsResponse)

	if err != nil {
		return err
	}

	*o = RosettaNetworkOptionsResponse(varRosettaNetworkOptionsResponse)

	return err
}

type NullableRosettaNetworkOptionsResponse struct {
	value *RosettaNetworkOptionsResponse
	isSet bool
}

func (v NullableRosettaNetworkOptionsResponse) Get() *RosettaNetworkOptionsResponse {
	return v.value
}

func (v *NullableRosettaNetworkOptionsResponse) Set(val *RosettaNetworkOptionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRosettaNetworkOptionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRosettaNetworkOptionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRosettaNetworkOptionsResponse(val *RosettaNetworkOptionsResponse) *NullableRosettaNetworkOptionsResponse {
	return &NullableRosettaNetworkOptionsResponse{value: val, isSet: true}
}

func (v NullableRosettaNetworkOptionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRosettaNetworkOptionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
