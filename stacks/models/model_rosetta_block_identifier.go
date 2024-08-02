package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the RosettaBlockIdentifier type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RosettaBlockIdentifier{}

// RosettaBlockIdentifier The block_identifier uniquely identifies a block in a particular network.
type RosettaBlockIdentifier struct {
	// This is also known as the block hash.
	Hash string `json:"hash"`
	// This is also known as the block height.
	Index int32 `json:"index"`
}

type _RosettaBlockIdentifier RosettaBlockIdentifier

// NewRosettaBlockIdentifier instantiates a new RosettaBlockIdentifier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRosettaBlockIdentifier(hash string, index int32) *RosettaBlockIdentifier {
	this := RosettaBlockIdentifier{}
	this.Hash = hash
	this.Index = index
	return &this
}

// NewRosettaBlockIdentifierWithDefaults instantiates a new RosettaBlockIdentifier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRosettaBlockIdentifierWithDefaults() *RosettaBlockIdentifier {
	this := RosettaBlockIdentifier{}
	return &this
}

// GetHash returns the Hash field value
func (o *RosettaBlockIdentifier) GetHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hash
}

// GetHashOk returns a tuple with the Hash field value
// and a boolean to check if the value has been set.
func (o *RosettaBlockIdentifier) GetHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hash, true
}

// SetHash sets field value
func (o *RosettaBlockIdentifier) SetHash(v string) {
	o.Hash = v
}

// GetIndex returns the Index field value
func (o *RosettaBlockIdentifier) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *RosettaBlockIdentifier) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *RosettaBlockIdentifier) SetIndex(v int32) {
	o.Index = v
}

func (o RosettaBlockIdentifier) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RosettaBlockIdentifier) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["hash"] = o.Hash
	toSerialize["index"] = o.Index
	return toSerialize, nil
}

func (o *RosettaBlockIdentifier) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"hash",
		"index",
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

	varRosettaBlockIdentifier := _RosettaBlockIdentifier{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRosettaBlockIdentifier)

	if err != nil {
		return err
	}

	*o = RosettaBlockIdentifier(varRosettaBlockIdentifier)

	return err
}

type NullableRosettaBlockIdentifier struct {
	value *RosettaBlockIdentifier
	isSet bool
}

func (v NullableRosettaBlockIdentifier) Get() *RosettaBlockIdentifier {
	return v.value
}

func (v *NullableRosettaBlockIdentifier) Set(val *RosettaBlockIdentifier) {
	v.value = val
	v.isSet = true
}

func (v NullableRosettaBlockIdentifier) IsSet() bool {
	return v.isSet
}

func (v *NullableRosettaBlockIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRosettaBlockIdentifier(val *RosettaBlockIdentifier) *NullableRosettaBlockIdentifier {
	return &NullableRosettaBlockIdentifier{value: val, isSet: true}
}

func (v NullableRosettaBlockIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRosettaBlockIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
