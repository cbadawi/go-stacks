package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the TargetBlockTime type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &TargetBlockTime{}

// TargetBlockTime struct for TargetBlockTime
type TargetBlockTime struct {
	TargetBlockTime int32 `json:"target_block_time"`
}

type _TargetBlockTime TargetBlockTime

// NewTargetBlockTime instantiates a new TargetBlockTime object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTargetBlockTime(targetBlockTime int32) *TargetBlockTime {
	this := TargetBlockTime{}
	this.TargetBlockTime = targetBlockTime
	return &this
}

// NewTargetBlockTimeWithDefaults instantiates a new TargetBlockTime object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTargetBlockTimeWithDefaults() *TargetBlockTime {
	this := TargetBlockTime{}
	return &this
}

// GetTargetBlockTime returns the TargetBlockTime field value
func (o *TargetBlockTime) GetTargetBlockTime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TargetBlockTime
}

// GetTargetBlockTimeOk returns a tuple with the TargetBlockTime field value
// and a boolean to check if the value has been set.
func (o *TargetBlockTime) GetTargetBlockTimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetBlockTime, true
}

// SetTargetBlockTime sets field value
func (o *TargetBlockTime) SetTargetBlockTime(v int32) {
	o.TargetBlockTime = v
}

func (o TargetBlockTime) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TargetBlockTime) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["target_block_time"] = o.TargetBlockTime
	return toSerialize, nil
}

func (o *TargetBlockTime) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"target_block_time",
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

	varTargetBlockTime := _TargetBlockTime{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTargetBlockTime)

	if err != nil {
		return err
	}

	*o = TargetBlockTime(varTargetBlockTime)

	return err
}

type NullableTargetBlockTime struct {
	value *TargetBlockTime
	isSet bool
}

func (v NullableTargetBlockTime) Get() *TargetBlockTime {
	return v.value
}

func (v *NullableTargetBlockTime) Set(val *TargetBlockTime) {
	v.value = val
	v.isSet = true
}

func (v NullableTargetBlockTime) IsSet() bool {
	return v.isSet
}

func (v *NullableTargetBlockTime) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTargetBlockTime(val *TargetBlockTime) *NullableTargetBlockTime {
	return &NullableTargetBlockTime{value: val, isSet: true}
}

func (v NullableTargetBlockTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTargetBlockTime) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
