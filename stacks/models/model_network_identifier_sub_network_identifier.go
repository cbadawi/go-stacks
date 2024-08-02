package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the NetworkIdentifierSubNetworkIdentifier type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &NetworkIdentifierSubNetworkIdentifier{}

// NetworkIdentifierSubNetworkIdentifier In blockchains with sharded state, the SubNetworkIdentifier is required to query some object on a specific shard. This identifier is optional for all non-sharded blockchains.
type NetworkIdentifierSubNetworkIdentifier struct {
	// Network name
	Network  string                                         `json:"network"`
	Metadata *NetworkIdentifierSubNetworkIdentifierMetadata `json:"metadata,omitempty"`
}

type _NetworkIdentifierSubNetworkIdentifier NetworkIdentifierSubNetworkIdentifier

// NewNetworkIdentifierSubNetworkIdentifier instantiates a new NetworkIdentifierSubNetworkIdentifier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkIdentifierSubNetworkIdentifier(network string) *NetworkIdentifierSubNetworkIdentifier {
	this := NetworkIdentifierSubNetworkIdentifier{}
	this.Network = network
	return &this
}

// NewNetworkIdentifierSubNetworkIdentifierWithDefaults instantiates a new NetworkIdentifierSubNetworkIdentifier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkIdentifierSubNetworkIdentifierWithDefaults() *NetworkIdentifierSubNetworkIdentifier {
	this := NetworkIdentifierSubNetworkIdentifier{}
	return &this
}

// GetNetwork returns the Network field value
func (o *NetworkIdentifierSubNetworkIdentifier) GetNetwork() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Network
}

// GetNetworkOk returns a tuple with the Network field value
// and a boolean to check if the value has been set.
func (o *NetworkIdentifierSubNetworkIdentifier) GetNetworkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Network, true
}

// SetNetwork sets field value
func (o *NetworkIdentifierSubNetworkIdentifier) SetNetwork(v string) {
	o.Network = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *NetworkIdentifierSubNetworkIdentifier) GetMetadata() NetworkIdentifierSubNetworkIdentifierMetadata {
	if o == nil || utils.IsNil(o.Metadata) {
		var ret NetworkIdentifierSubNetworkIdentifierMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkIdentifierSubNetworkIdentifier) GetMetadataOk() (*NetworkIdentifierSubNetworkIdentifierMetadata, bool) {
	if o == nil || utils.IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *NetworkIdentifierSubNetworkIdentifier) HasMetadata() bool {
	if o != nil && !utils.IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given NetworkIdentifierSubNetworkIdentifierMetadata and assigns it to the Metadata field.
func (o *NetworkIdentifierSubNetworkIdentifier) SetMetadata(v NetworkIdentifierSubNetworkIdentifierMetadata) {
	o.Metadata = &v
}

func (o NetworkIdentifierSubNetworkIdentifier) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkIdentifierSubNetworkIdentifier) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["network"] = o.Network
	if !utils.IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

func (o *NetworkIdentifierSubNetworkIdentifier) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"network",
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

	varNetworkIdentifierSubNetworkIdentifier := _NetworkIdentifierSubNetworkIdentifier{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNetworkIdentifierSubNetworkIdentifier)

	if err != nil {
		return err
	}

	*o = NetworkIdentifierSubNetworkIdentifier(varNetworkIdentifierSubNetworkIdentifier)

	return err
}

type NullableNetworkIdentifierSubNetworkIdentifier struct {
	value *NetworkIdentifierSubNetworkIdentifier
	isSet bool
}

func (v NullableNetworkIdentifierSubNetworkIdentifier) Get() *NetworkIdentifierSubNetworkIdentifier {
	return v.value
}

func (v *NullableNetworkIdentifierSubNetworkIdentifier) Set(val *NetworkIdentifierSubNetworkIdentifier) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkIdentifierSubNetworkIdentifier) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkIdentifierSubNetworkIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkIdentifierSubNetworkIdentifier(val *NetworkIdentifierSubNetworkIdentifier) *NullableNetworkIdentifierSubNetworkIdentifier {
	return &NullableNetworkIdentifierSubNetworkIdentifier{value: val, isSet: true}
}

func (v NullableNetworkIdentifierSubNetworkIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkIdentifierSubNetworkIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
