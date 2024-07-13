package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the MempoolFeePrioritiesAll type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &MempoolFeePrioritiesAll{}

// MempoolFeePrioritiesAll struct for MempoolFeePrioritiesAll
type MempoolFeePrioritiesAll struct {
	NoPriority     int32 `json:"no_priority"`
	LowPriority    int32 `json:"low_priority"`
	MediumPriority int32 `json:"medium_priority"`
	HighPriority   int32 `json:"high_priority"`
}

type _MempoolFeePrioritiesAll MempoolFeePrioritiesAll

// NewMempoolFeePrioritiesAll instantiates a new MempoolFeePrioritiesAll object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMempoolFeePrioritiesAll(noPriority int32, lowPriority int32, mediumPriority int32, highPriority int32) *MempoolFeePrioritiesAll {
	this := MempoolFeePrioritiesAll{}
	this.NoPriority = noPriority
	this.LowPriority = lowPriority
	this.MediumPriority = mediumPriority
	this.HighPriority = highPriority
	return &this
}

// NewMempoolFeePrioritiesAllWithDefaults instantiates a new MempoolFeePrioritiesAll object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMempoolFeePrioritiesAllWithDefaults() *MempoolFeePrioritiesAll {
	this := MempoolFeePrioritiesAll{}
	return &this
}

// GetNoPriority returns the NoPriority field value
func (o *MempoolFeePrioritiesAll) GetNoPriority() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NoPriority
}

// GetNoPriorityOk returns a tuple with the NoPriority field value
// and a boolean to check if the value has been set.
func (o *MempoolFeePrioritiesAll) GetNoPriorityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NoPriority, true
}

// SetNoPriority sets field value
func (o *MempoolFeePrioritiesAll) SetNoPriority(v int32) {
	o.NoPriority = v
}

// GetLowPriority returns the LowPriority field value
func (o *MempoolFeePrioritiesAll) GetLowPriority() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.LowPriority
}

// GetLowPriorityOk returns a tuple with the LowPriority field value
// and a boolean to check if the value has been set.
func (o *MempoolFeePrioritiesAll) GetLowPriorityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LowPriority, true
}

// SetLowPriority sets field value
func (o *MempoolFeePrioritiesAll) SetLowPriority(v int32) {
	o.LowPriority = v
}

// GetMediumPriority returns the MediumPriority field value
func (o *MempoolFeePrioritiesAll) GetMediumPriority() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MediumPriority
}

// GetMediumPriorityOk returns a tuple with the MediumPriority field value
// and a boolean to check if the value has been set.
func (o *MempoolFeePrioritiesAll) GetMediumPriorityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MediumPriority, true
}

// SetMediumPriority sets field value
func (o *MempoolFeePrioritiesAll) SetMediumPriority(v int32) {
	o.MediumPriority = v
}

// GetHighPriority returns the HighPriority field value
func (o *MempoolFeePrioritiesAll) GetHighPriority() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.HighPriority
}

// GetHighPriorityOk returns a tuple with the HighPriority field value
// and a boolean to check if the value has been set.
func (o *MempoolFeePrioritiesAll) GetHighPriorityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HighPriority, true
}

// SetHighPriority sets field value
func (o *MempoolFeePrioritiesAll) SetHighPriority(v int32) {
	o.HighPriority = v
}

func (o MempoolFeePrioritiesAll) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MempoolFeePrioritiesAll) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["no_priority"] = o.NoPriority
	toSerialize["low_priority"] = o.LowPriority
	toSerialize["medium_priority"] = o.MediumPriority
	toSerialize["high_priority"] = o.HighPriority
	return toSerialize, nil
}

func (o *MempoolFeePrioritiesAll) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"no_priority",
		"low_priority",
		"medium_priority",
		"high_priority",
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

	varMempoolFeePrioritiesAll := _MempoolFeePrioritiesAll{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMempoolFeePrioritiesAll)

	if err != nil {
		return err
	}

	*o = MempoolFeePrioritiesAll(varMempoolFeePrioritiesAll)

	return err
}

type NullableMempoolFeePrioritiesAll struct {
	value *MempoolFeePrioritiesAll
	isSet bool
}

func (v NullableMempoolFeePrioritiesAll) Get() *MempoolFeePrioritiesAll {
	return v.value
}

func (v *NullableMempoolFeePrioritiesAll) Set(val *MempoolFeePrioritiesAll) {
	v.value = val
	v.isSet = true
}

func (v NullableMempoolFeePrioritiesAll) IsSet() bool {
	return v.isSet
}

func (v *NullableMempoolFeePrioritiesAll) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMempoolFeePrioritiesAll(val *MempoolFeePrioritiesAll) *NullableMempoolFeePrioritiesAll {
	return &NullableMempoolFeePrioritiesAll{value: val, isSet: true}
}

func (v NullableMempoolFeePrioritiesAll) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMempoolFeePrioritiesAll) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
