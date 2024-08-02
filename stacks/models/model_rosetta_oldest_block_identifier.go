package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the RosettaOldestBlockIdentifier type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RosettaOldestBlockIdentifier{}

// RosettaOldestBlockIdentifier The block_identifier uniquely identifies a block in a particular network.
type RosettaOldestBlockIdentifier struct {
	// This is also known as the block height.
	Index int32 `json:"index"`
	// Block hash
	Hash string `json:"hash"`
}

type _RosettaOldestBlockIdentifier RosettaOldestBlockIdentifier

// NewRosettaOldestBlockIdentifier instantiates a new RosettaOldestBlockIdentifier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRosettaOldestBlockIdentifier(index int32, hash string) *RosettaOldestBlockIdentifier {
	this := RosettaOldestBlockIdentifier{}
	this.Index = index
	this.Hash = hash
	return &this
}

// NewRosettaOldestBlockIdentifierWithDefaults instantiates a new RosettaOldestBlockIdentifier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRosettaOldestBlockIdentifierWithDefaults() *RosettaOldestBlockIdentifier {
	this := RosettaOldestBlockIdentifier{}
	return &this
}

// GetIndex returns the Index field value
func (o *RosettaOldestBlockIdentifier) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *RosettaOldestBlockIdentifier) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *RosettaOldestBlockIdentifier) SetIndex(v int32) {
	o.Index = v
}

// GetHash returns the Hash field value
func (o *RosettaOldestBlockIdentifier) GetHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hash
}

// GetHashOk returns a tuple with the Hash field value
// and a boolean to check if the value has been set.
func (o *RosettaOldestBlockIdentifier) GetHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hash, true
}

// SetHash sets field value
func (o *RosettaOldestBlockIdentifier) SetHash(v string) {
	o.Hash = v
}

func (o RosettaOldestBlockIdentifier) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RosettaOldestBlockIdentifier) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["index"] = o.Index
	toSerialize["hash"] = o.Hash
	return toSerialize, nil
}

func (o *RosettaOldestBlockIdentifier) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"index",
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

	varRosettaOldestBlockIdentifier := _RosettaOldestBlockIdentifier{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRosettaOldestBlockIdentifier)

	if err != nil {
		return err
	}

	*o = RosettaOldestBlockIdentifier(varRosettaOldestBlockIdentifier)

	return err
}

type NullableRosettaOldestBlockIdentifier struct {
	value *RosettaOldestBlockIdentifier
	isSet bool
}

func (v NullableRosettaOldestBlockIdentifier) Get() *RosettaOldestBlockIdentifier {
	return v.value
}

func (v *NullableRosettaOldestBlockIdentifier) Set(val *RosettaOldestBlockIdentifier) {
	v.value = val
	v.isSet = true
}

func (v NullableRosettaOldestBlockIdentifier) IsSet() bool {
	return v.isSet
}

func (v *NullableRosettaOldestBlockIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRosettaOldestBlockIdentifier(val *RosettaOldestBlockIdentifier) *NullableRosettaOldestBlockIdentifier {
	return &NullableRosettaOldestBlockIdentifier{value: val, isSet: true}
}

func (v NullableRosettaOldestBlockIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRosettaOldestBlockIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
