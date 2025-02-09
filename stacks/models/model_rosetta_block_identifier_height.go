package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the RosettaBlockIdentifierHeight type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RosettaBlockIdentifierHeight{}

// RosettaBlockIdentifierHeight This is also known as the block height.
type RosettaBlockIdentifierHeight struct {
	// This is also known as the block height.
	Index int32 `json:"index"`
}

type _RosettaBlockIdentifierHeight RosettaBlockIdentifierHeight

// NewRosettaBlockIdentifierHeight instantiates a new RosettaBlockIdentifierHeight object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRosettaBlockIdentifierHeight(index int32) *RosettaBlockIdentifierHeight {
	this := RosettaBlockIdentifierHeight{}
	this.Index = index
	return &this
}

// NewRosettaBlockIdentifierHeightWithDefaults instantiates a new RosettaBlockIdentifierHeight object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRosettaBlockIdentifierHeightWithDefaults() *RosettaBlockIdentifierHeight {
	this := RosettaBlockIdentifierHeight{}
	return &this
}

// GetIndex returns the Index field value
func (o *RosettaBlockIdentifierHeight) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *RosettaBlockIdentifierHeight) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *RosettaBlockIdentifierHeight) SetIndex(v int32) {
	o.Index = v
}

func (o RosettaBlockIdentifierHeight) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RosettaBlockIdentifierHeight) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["index"] = o.Index
	return toSerialize, nil
}

func (o *RosettaBlockIdentifierHeight) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varRosettaBlockIdentifierHeight := _RosettaBlockIdentifierHeight{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRosettaBlockIdentifierHeight)

	if err != nil {
		return err
	}

	*o = RosettaBlockIdentifierHeight(varRosettaBlockIdentifierHeight)

	return err
}

type NullableRosettaBlockIdentifierHeight struct {
	value *RosettaBlockIdentifierHeight
	isSet bool
}

func (v NullableRosettaBlockIdentifierHeight) Get() *RosettaBlockIdentifierHeight {
	return v.value
}

func (v *NullableRosettaBlockIdentifierHeight) Set(val *RosettaBlockIdentifierHeight) {
	v.value = val
	v.isSet = true
}

func (v NullableRosettaBlockIdentifierHeight) IsSet() bool {
	return v.isSet
}

func (v *NullableRosettaBlockIdentifierHeight) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRosettaBlockIdentifierHeight(val *RosettaBlockIdentifierHeight) *NullableRosettaBlockIdentifierHeight {
	return &NullableRosettaBlockIdentifierHeight{value: val, isSet: true}
}

func (v NullableRosettaBlockIdentifierHeight) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRosettaBlockIdentifierHeight) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
