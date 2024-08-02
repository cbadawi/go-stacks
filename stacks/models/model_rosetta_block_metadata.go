package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the RosettaBlockMetadata type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RosettaBlockMetadata{}

// RosettaBlockMetadata meta data
type RosettaBlockMetadata struct {
	//
	BurnBlockHeight float32 `json:"burn_block_height"`
}

type _RosettaBlockMetadata RosettaBlockMetadata

// NewRosettaBlockMetadata instantiates a new RosettaBlockMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRosettaBlockMetadata(burnBlockHeight float32) *RosettaBlockMetadata {
	this := RosettaBlockMetadata{}
	this.BurnBlockHeight = burnBlockHeight
	return &this
}

// NewRosettaBlockMetadataWithDefaults instantiates a new RosettaBlockMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRosettaBlockMetadataWithDefaults() *RosettaBlockMetadata {
	this := RosettaBlockMetadata{}
	return &this
}

// GetBurnBlockHeight returns the BurnBlockHeight field value
func (o *RosettaBlockMetadata) GetBurnBlockHeight() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.BurnBlockHeight
}

// GetBurnBlockHeightOk returns a tuple with the BurnBlockHeight field value
// and a boolean to check if the value has been set.
func (o *RosettaBlockMetadata) GetBurnBlockHeightOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BurnBlockHeight, true
}

// SetBurnBlockHeight sets field value
func (o *RosettaBlockMetadata) SetBurnBlockHeight(v float32) {
	o.BurnBlockHeight = v
}

func (o RosettaBlockMetadata) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RosettaBlockMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["burn_block_height"] = o.BurnBlockHeight
	return toSerialize, nil
}

func (o *RosettaBlockMetadata) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"burn_block_height",
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

	varRosettaBlockMetadata := _RosettaBlockMetadata{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRosettaBlockMetadata)

	if err != nil {
		return err
	}

	*o = RosettaBlockMetadata(varRosettaBlockMetadata)

	return err
}

type NullableRosettaBlockMetadata struct {
	value *RosettaBlockMetadata
	isSet bool
}

func (v NullableRosettaBlockMetadata) Get() *RosettaBlockMetadata {
	return v.value
}

func (v *NullableRosettaBlockMetadata) Set(val *RosettaBlockMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableRosettaBlockMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableRosettaBlockMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRosettaBlockMetadata(val *RosettaBlockMetadata) *NullableRosettaBlockMetadata {
	return &NullableRosettaBlockMetadata{value: val, isSet: true}
}

func (v NullableRosettaBlockMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRosettaBlockMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
