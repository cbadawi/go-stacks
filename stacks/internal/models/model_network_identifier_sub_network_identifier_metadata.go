package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the NetworkIdentifierSubNetworkIdentifierMetadata type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &NetworkIdentifierSubNetworkIdentifierMetadata{}

// NetworkIdentifierSubNetworkIdentifierMetadata Meta data from subnetwork identifier
type NetworkIdentifierSubNetworkIdentifierMetadata struct {
	// producer
	Producer string `json:"producer"`
}

type _NetworkIdentifierSubNetworkIdentifierMetadata NetworkIdentifierSubNetworkIdentifierMetadata

// NewNetworkIdentifierSubNetworkIdentifierMetadata instantiates a new NetworkIdentifierSubNetworkIdentifierMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkIdentifierSubNetworkIdentifierMetadata(producer string) *NetworkIdentifierSubNetworkIdentifierMetadata {
	this := NetworkIdentifierSubNetworkIdentifierMetadata{}
	this.Producer = producer
	return &this
}

// NewNetworkIdentifierSubNetworkIdentifierMetadataWithDefaults instantiates a new NetworkIdentifierSubNetworkIdentifierMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkIdentifierSubNetworkIdentifierMetadataWithDefaults() *NetworkIdentifierSubNetworkIdentifierMetadata {
	this := NetworkIdentifierSubNetworkIdentifierMetadata{}
	return &this
}

// GetProducer returns the Producer field value
func (o *NetworkIdentifierSubNetworkIdentifierMetadata) GetProducer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Producer
}

// GetProducerOk returns a tuple with the Producer field value
// and a boolean to check if the value has been set.
func (o *NetworkIdentifierSubNetworkIdentifierMetadata) GetProducerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Producer, true
}

// SetProducer sets field value
func (o *NetworkIdentifierSubNetworkIdentifierMetadata) SetProducer(v string) {
	o.Producer = v
}

func (o NetworkIdentifierSubNetworkIdentifierMetadata) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkIdentifierSubNetworkIdentifierMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["producer"] = o.Producer
	return toSerialize, nil
}

func (o *NetworkIdentifierSubNetworkIdentifierMetadata) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"producer",
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

	varNetworkIdentifierSubNetworkIdentifierMetadata := _NetworkIdentifierSubNetworkIdentifierMetadata{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNetworkIdentifierSubNetworkIdentifierMetadata)

	if err != nil {
		return err
	}

	*o = NetworkIdentifierSubNetworkIdentifierMetadata(varNetworkIdentifierSubNetworkIdentifierMetadata)

	return err
}

type NullableNetworkIdentifierSubNetworkIdentifierMetadata struct {
	value *NetworkIdentifierSubNetworkIdentifierMetadata
	isSet bool
}

func (v NullableNetworkIdentifierSubNetworkIdentifierMetadata) Get() *NetworkIdentifierSubNetworkIdentifierMetadata {
	return v.value
}

func (v *NullableNetworkIdentifierSubNetworkIdentifierMetadata) Set(val *NetworkIdentifierSubNetworkIdentifierMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkIdentifierSubNetworkIdentifierMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkIdentifierSubNetworkIdentifierMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkIdentifierSubNetworkIdentifierMetadata(val *NetworkIdentifierSubNetworkIdentifierMetadata) *NullableNetworkIdentifierSubNetworkIdentifierMetadata {
	return &NullableNetworkIdentifierSubNetworkIdentifierMetadata{value: val, isSet: true}
}

func (v NullableNetworkIdentifierSubNetworkIdentifierMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkIdentifierSubNetworkIdentifierMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
