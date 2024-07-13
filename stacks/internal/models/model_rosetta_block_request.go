package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the RosettaBlockRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RosettaBlockRequest{}

// RosettaBlockRequest A BlockRequest is utilized to make a block request on the /block endpoint.
type RosettaBlockRequest struct {
	NetworkIdentifier NetworkIdentifier             `json:"network_identifier"`
	BlockIdentifier   RosettaPartialBlockIdentifier `json:"block_identifier"`
}

type _RosettaBlockRequest RosettaBlockRequest

// NewRosettaBlockRequest instantiates a new RosettaBlockRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRosettaBlockRequest(networkIdentifier NetworkIdentifier, blockIdentifier RosettaPartialBlockIdentifier) *RosettaBlockRequest {
	this := RosettaBlockRequest{}
	this.NetworkIdentifier = networkIdentifier
	this.BlockIdentifier = blockIdentifier
	return &this
}

// NewRosettaBlockRequestWithDefaults instantiates a new RosettaBlockRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRosettaBlockRequestWithDefaults() *RosettaBlockRequest {
	this := RosettaBlockRequest{}
	return &this
}

// GetNetworkIdentifier returns the NetworkIdentifier field value
func (o *RosettaBlockRequest) GetNetworkIdentifier() NetworkIdentifier {
	if o == nil {
		var ret NetworkIdentifier
		return ret
	}

	return o.NetworkIdentifier
}

// GetNetworkIdentifierOk returns a tuple with the NetworkIdentifier field value
// and a boolean to check if the value has been set.
func (o *RosettaBlockRequest) GetNetworkIdentifierOk() (*NetworkIdentifier, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkIdentifier, true
}

// SetNetworkIdentifier sets field value
func (o *RosettaBlockRequest) SetNetworkIdentifier(v NetworkIdentifier) {
	o.NetworkIdentifier = v
}

// GetBlockIdentifier returns the BlockIdentifier field value
func (o *RosettaBlockRequest) GetBlockIdentifier() RosettaPartialBlockIdentifier {
	if o == nil {
		var ret RosettaPartialBlockIdentifier
		return ret
	}

	return o.BlockIdentifier
}

// GetBlockIdentifierOk returns a tuple with the BlockIdentifier field value
// and a boolean to check if the value has been set.
func (o *RosettaBlockRequest) GetBlockIdentifierOk() (*RosettaPartialBlockIdentifier, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockIdentifier, true
}

// SetBlockIdentifier sets field value
func (o *RosettaBlockRequest) SetBlockIdentifier(v RosettaPartialBlockIdentifier) {
	o.BlockIdentifier = v
}

func (o RosettaBlockRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RosettaBlockRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["network_identifier"] = o.NetworkIdentifier
	toSerialize["block_identifier"] = o.BlockIdentifier
	return toSerialize, nil
}

func (o *RosettaBlockRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"network_identifier",
		"block_identifier",
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

	varRosettaBlockRequest := _RosettaBlockRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRosettaBlockRequest)

	if err != nil {
		return err
	}

	*o = RosettaBlockRequest(varRosettaBlockRequest)

	return err
}

type NullableRosettaBlockRequest struct {
	value *RosettaBlockRequest
	isSet bool
}

func (v NullableRosettaBlockRequest) Get() *RosettaBlockRequest {
	return v.value
}

func (v *NullableRosettaBlockRequest) Set(val *RosettaBlockRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRosettaBlockRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRosettaBlockRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRosettaBlockRequest(val *RosettaBlockRequest) *NullableRosettaBlockRequest {
	return &NullableRosettaBlockRequest{value: val, isSet: true}
}

func (v NullableRosettaBlockRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRosettaBlockRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
