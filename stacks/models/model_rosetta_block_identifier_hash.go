package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the RosettaBlockIdentifierHash type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RosettaBlockIdentifierHash{}

// RosettaBlockIdentifierHash This is also known as the block hash.
type RosettaBlockIdentifierHash struct {
	// This is also known as the block hash.
	Hash string `json:"hash"`
}

type _RosettaBlockIdentifierHash RosettaBlockIdentifierHash

// NewRosettaBlockIdentifierHash instantiates a new RosettaBlockIdentifierHash object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRosettaBlockIdentifierHash(hash string) *RosettaBlockIdentifierHash {
	this := RosettaBlockIdentifierHash{}
	this.Hash = hash
	return &this
}

// NewRosettaBlockIdentifierHashWithDefaults instantiates a new RosettaBlockIdentifierHash object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRosettaBlockIdentifierHashWithDefaults() *RosettaBlockIdentifierHash {
	this := RosettaBlockIdentifierHash{}
	return &this
}

// GetHash returns the Hash field value
func (o *RosettaBlockIdentifierHash) GetHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hash
}

// GetHashOk returns a tuple with the Hash field value
// and a boolean to check if the value has been set.
func (o *RosettaBlockIdentifierHash) GetHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hash, true
}

// SetHash sets field value
func (o *RosettaBlockIdentifierHash) SetHash(v string) {
	o.Hash = v
}

func (o RosettaBlockIdentifierHash) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RosettaBlockIdentifierHash) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["hash"] = o.Hash
	return toSerialize, nil
}

func (o *RosettaBlockIdentifierHash) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"hash",
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

	varRosettaBlockIdentifierHash := _RosettaBlockIdentifierHash{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRosettaBlockIdentifierHash)

	if err != nil {
		return err
	}

	*o = RosettaBlockIdentifierHash(varRosettaBlockIdentifierHash)

	return err
}

type NullableRosettaBlockIdentifierHash struct {
	value *RosettaBlockIdentifierHash
	isSet bool
}

func (v NullableRosettaBlockIdentifierHash) Get() *RosettaBlockIdentifierHash {
	return v.value
}

func (v *NullableRosettaBlockIdentifierHash) Set(val *RosettaBlockIdentifierHash) {
	v.value = val
	v.isSet = true
}

func (v NullableRosettaBlockIdentifierHash) IsSet() bool {
	return v.isSet
}

func (v *NullableRosettaBlockIdentifierHash) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRosettaBlockIdentifierHash(val *RosettaBlockIdentifierHash) *NullableRosettaBlockIdentifierHash {
	return &NullableRosettaBlockIdentifierHash{value: val, isSet: true}
}

func (v NullableRosettaBlockIdentifierHash) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRosettaBlockIdentifierHash) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
