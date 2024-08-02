package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the RosettaParentBlockIdentifier type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RosettaParentBlockIdentifier{}

// RosettaParentBlockIdentifier The block_identifier uniquely identifies a block in a particular network.
type RosettaParentBlockIdentifier struct {
	// This is also known as the block height.
	Index int32 `json:"index"`
	// Block hash
	Hash string `json:"hash"`
}

type _RosettaParentBlockIdentifier RosettaParentBlockIdentifier

// NewRosettaParentBlockIdentifier instantiates a new RosettaParentBlockIdentifier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRosettaParentBlockIdentifier(index int32, hash string) *RosettaParentBlockIdentifier {
	this := RosettaParentBlockIdentifier{}
	this.Index = index
	this.Hash = hash
	return &this
}

// NewRosettaParentBlockIdentifierWithDefaults instantiates a new RosettaParentBlockIdentifier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRosettaParentBlockIdentifierWithDefaults() *RosettaParentBlockIdentifier {
	this := RosettaParentBlockIdentifier{}
	return &this
}

// GetIndex returns the Index field value
func (o *RosettaParentBlockIdentifier) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *RosettaParentBlockIdentifier) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *RosettaParentBlockIdentifier) SetIndex(v int32) {
	o.Index = v
}

// GetHash returns the Hash field value
func (o *RosettaParentBlockIdentifier) GetHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hash
}

// GetHashOk returns a tuple with the Hash field value
// and a boolean to check if the value has been set.
func (o *RosettaParentBlockIdentifier) GetHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hash, true
}

// SetHash sets field value
func (o *RosettaParentBlockIdentifier) SetHash(v string) {
	o.Hash = v
}

func (o RosettaParentBlockIdentifier) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RosettaParentBlockIdentifier) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["index"] = o.Index
	toSerialize["hash"] = o.Hash
	return toSerialize, nil
}

func (o *RosettaParentBlockIdentifier) UnmarshalJSON(data []byte) (err error) {
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

	varRosettaParentBlockIdentifier := _RosettaParentBlockIdentifier{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRosettaParentBlockIdentifier)

	if err != nil {
		return err
	}

	*o = RosettaParentBlockIdentifier(varRosettaParentBlockIdentifier)

	return err
}

type NullableRosettaParentBlockIdentifier struct {
	value *RosettaParentBlockIdentifier
	isSet bool
}

func (v NullableRosettaParentBlockIdentifier) Get() *RosettaParentBlockIdentifier {
	return v.value
}

func (v *NullableRosettaParentBlockIdentifier) Set(val *RosettaParentBlockIdentifier) {
	v.value = val
	v.isSet = true
}

func (v NullableRosettaParentBlockIdentifier) IsSet() bool {
	return v.isSet
}

func (v *NullableRosettaParentBlockIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRosettaParentBlockIdentifier(val *RosettaParentBlockIdentifier) *NullableRosettaParentBlockIdentifier {
	return &NullableRosettaParentBlockIdentifier{value: val, isSet: true}
}

func (v NullableRosettaParentBlockIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRosettaParentBlockIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
