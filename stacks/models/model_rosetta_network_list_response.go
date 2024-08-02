package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the RosettaNetworkListResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RosettaNetworkListResponse{}

// RosettaNetworkListResponse A NetworkListResponse contains all NetworkIdentifiers that the node can serve information for.
type RosettaNetworkListResponse struct {
	// The network_identifier specifies which network a particular object is associated with.
	NetworkIdentifiers []NetworkIdentifier `json:"network_identifiers"`
}

type _RosettaNetworkListResponse RosettaNetworkListResponse

// NewRosettaNetworkListResponse instantiates a new RosettaNetworkListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRosettaNetworkListResponse(networkIdentifiers []NetworkIdentifier) *RosettaNetworkListResponse {
	this := RosettaNetworkListResponse{}
	this.NetworkIdentifiers = networkIdentifiers
	return &this
}

// NewRosettaNetworkListResponseWithDefaults instantiates a new RosettaNetworkListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRosettaNetworkListResponseWithDefaults() *RosettaNetworkListResponse {
	this := RosettaNetworkListResponse{}
	return &this
}

// GetNetworkIdentifiers returns the NetworkIdentifiers field value
func (o *RosettaNetworkListResponse) GetNetworkIdentifiers() []NetworkIdentifier {
	if o == nil {
		var ret []NetworkIdentifier
		return ret
	}

	return o.NetworkIdentifiers
}

// GetNetworkIdentifiersOk returns a tuple with the NetworkIdentifiers field value
// and a boolean to check if the value has been set.
func (o *RosettaNetworkListResponse) GetNetworkIdentifiersOk() ([]NetworkIdentifier, bool) {
	if o == nil {
		return nil, false
	}
	return o.NetworkIdentifiers, true
}

// SetNetworkIdentifiers sets field value
func (o *RosettaNetworkListResponse) SetNetworkIdentifiers(v []NetworkIdentifier) {
	o.NetworkIdentifiers = v
}

func (o RosettaNetworkListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RosettaNetworkListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["network_identifiers"] = o.NetworkIdentifiers
	return toSerialize, nil
}

func (o *RosettaNetworkListResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"network_identifiers",
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

	varRosettaNetworkListResponse := _RosettaNetworkListResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRosettaNetworkListResponse)

	if err != nil {
		return err
	}

	*o = RosettaNetworkListResponse(varRosettaNetworkListResponse)

	return err
}

type NullableRosettaNetworkListResponse struct {
	value *RosettaNetworkListResponse
	isSet bool
}

func (v NullableRosettaNetworkListResponse) Get() *RosettaNetworkListResponse {
	return v.value
}

func (v *NullableRosettaNetworkListResponse) Set(val *RosettaNetworkListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRosettaNetworkListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRosettaNetworkListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRosettaNetworkListResponse(val *RosettaNetworkListResponse) *NullableRosettaNetworkListResponse {
	return &NullableRosettaNetworkListResponse{value: val, isSet: true}
}

func (v NullableRosettaNetworkListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRosettaNetworkListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
