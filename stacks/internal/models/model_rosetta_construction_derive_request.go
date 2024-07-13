package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the RosettaConstructionDeriveRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RosettaConstructionDeriveRequest{}

// RosettaConstructionDeriveRequest Network is provided in the request because some blockchains have different address formats for different networks
type RosettaConstructionDeriveRequest struct {
	NetworkIdentifier NetworkIdentifier      `json:"network_identifier"`
	PublicKey         RosettaPublicKey       `json:"public_key"`
	Metadata          map[string]interface{} `json:"metadata,omitempty"`
}

type _RosettaConstructionDeriveRequest RosettaConstructionDeriveRequest

// NewRosettaConstructionDeriveRequest instantiates a new RosettaConstructionDeriveRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRosettaConstructionDeriveRequest(networkIdentifier NetworkIdentifier, publicKey RosettaPublicKey) *RosettaConstructionDeriveRequest {
	this := RosettaConstructionDeriveRequest{}
	this.NetworkIdentifier = networkIdentifier
	this.PublicKey = publicKey
	return &this
}

// NewRosettaConstructionDeriveRequestWithDefaults instantiates a new RosettaConstructionDeriveRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRosettaConstructionDeriveRequestWithDefaults() *RosettaConstructionDeriveRequest {
	this := RosettaConstructionDeriveRequest{}
	return &this
}

// GetNetworkIdentifier returns the NetworkIdentifier field value
func (o *RosettaConstructionDeriveRequest) GetNetworkIdentifier() NetworkIdentifier {
	if o == nil {
		var ret NetworkIdentifier
		return ret
	}

	return o.NetworkIdentifier
}

// GetNetworkIdentifierOk returns a tuple with the NetworkIdentifier field value
// and a boolean to check if the value has been set.
func (o *RosettaConstructionDeriveRequest) GetNetworkIdentifierOk() (*NetworkIdentifier, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkIdentifier, true
}

// SetNetworkIdentifier sets field value
func (o *RosettaConstructionDeriveRequest) SetNetworkIdentifier(v NetworkIdentifier) {
	o.NetworkIdentifier = v
}

// GetPublicKey returns the PublicKey field value
func (o *RosettaConstructionDeriveRequest) GetPublicKey() RosettaPublicKey {
	if o == nil {
		var ret RosettaPublicKey
		return ret
	}

	return o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value
// and a boolean to check if the value has been set.
func (o *RosettaConstructionDeriveRequest) GetPublicKeyOk() (*RosettaPublicKey, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicKey, true
}

// SetPublicKey sets field value
func (o *RosettaConstructionDeriveRequest) SetPublicKey(v RosettaPublicKey) {
	o.PublicKey = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *RosettaConstructionDeriveRequest) GetMetadata() map[string]interface{} {
	if o == nil || utils.IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaConstructionDeriveRequest) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || utils.IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *RosettaConstructionDeriveRequest) HasMetadata() bool {
	if o != nil && !utils.IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *RosettaConstructionDeriveRequest) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o RosettaConstructionDeriveRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RosettaConstructionDeriveRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["network_identifier"] = o.NetworkIdentifier
	toSerialize["public_key"] = o.PublicKey
	if !utils.IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

func (o *RosettaConstructionDeriveRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"network_identifier",
		"public_key",
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

	varRosettaConstructionDeriveRequest := _RosettaConstructionDeriveRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRosettaConstructionDeriveRequest)

	if err != nil {
		return err
	}

	*o = RosettaConstructionDeriveRequest(varRosettaConstructionDeriveRequest)

	return err
}

type NullableRosettaConstructionDeriveRequest struct {
	value *RosettaConstructionDeriveRequest
	isSet bool
}

func (v NullableRosettaConstructionDeriveRequest) Get() *RosettaConstructionDeriveRequest {
	return v.value
}

func (v *NullableRosettaConstructionDeriveRequest) Set(val *RosettaConstructionDeriveRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRosettaConstructionDeriveRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRosettaConstructionDeriveRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRosettaConstructionDeriveRequest(val *RosettaConstructionDeriveRequest) *NullableRosettaConstructionDeriveRequest {
	return &NullableRosettaConstructionDeriveRequest{value: val, isSet: true}
}

func (v NullableRosettaConstructionDeriveRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRosettaConstructionDeriveRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
