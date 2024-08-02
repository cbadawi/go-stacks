package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the PoisonMicroblockTransactionMetadataPoisonMicroblock type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PoisonMicroblockTransactionMetadataPoisonMicroblock{}

// PoisonMicroblockTransactionMetadataPoisonMicroblock struct for PoisonMicroblockTransactionMetadataPoisonMicroblock
type PoisonMicroblockTransactionMetadataPoisonMicroblock struct {
	// Hex encoded microblock header
	MicroblockHeader1 string `json:"microblock_header_1"`
	// Hex encoded microblock header
	MicroblockHeader2 string `json:"microblock_header_2"`
}

type _PoisonMicroblockTransactionMetadataPoisonMicroblock PoisonMicroblockTransactionMetadataPoisonMicroblock

// NewPoisonMicroblockTransactionMetadataPoisonMicroblock instantiates a new PoisonMicroblockTransactionMetadataPoisonMicroblock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoisonMicroblockTransactionMetadataPoisonMicroblock(microblockHeader1 string, microblockHeader2 string) *PoisonMicroblockTransactionMetadataPoisonMicroblock {
	this := PoisonMicroblockTransactionMetadataPoisonMicroblock{}
	this.MicroblockHeader1 = microblockHeader1
	this.MicroblockHeader2 = microblockHeader2
	return &this
}

// NewPoisonMicroblockTransactionMetadataPoisonMicroblockWithDefaults instantiates a new PoisonMicroblockTransactionMetadataPoisonMicroblock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoisonMicroblockTransactionMetadataPoisonMicroblockWithDefaults() *PoisonMicroblockTransactionMetadataPoisonMicroblock {
	this := PoisonMicroblockTransactionMetadataPoisonMicroblock{}
	return &this
}

// GetMicroblockHeader1 returns the MicroblockHeader1 field value
func (o *PoisonMicroblockTransactionMetadataPoisonMicroblock) GetMicroblockHeader1() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MicroblockHeader1
}

// GetMicroblockHeader1Ok returns a tuple with the MicroblockHeader1 field value
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransactionMetadataPoisonMicroblock) GetMicroblockHeader1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MicroblockHeader1, true
}

// SetMicroblockHeader1 sets field value
func (o *PoisonMicroblockTransactionMetadataPoisonMicroblock) SetMicroblockHeader1(v string) {
	o.MicroblockHeader1 = v
}

// GetMicroblockHeader2 returns the MicroblockHeader2 field value
func (o *PoisonMicroblockTransactionMetadataPoisonMicroblock) GetMicroblockHeader2() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MicroblockHeader2
}

// GetMicroblockHeader2Ok returns a tuple with the MicroblockHeader2 field value
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransactionMetadataPoisonMicroblock) GetMicroblockHeader2Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MicroblockHeader2, true
}

// SetMicroblockHeader2 sets field value
func (o *PoisonMicroblockTransactionMetadataPoisonMicroblock) SetMicroblockHeader2(v string) {
	o.MicroblockHeader2 = v
}

func (o PoisonMicroblockTransactionMetadataPoisonMicroblock) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PoisonMicroblockTransactionMetadataPoisonMicroblock) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["microblock_header_1"] = o.MicroblockHeader1
	toSerialize["microblock_header_2"] = o.MicroblockHeader2
	return toSerialize, nil
}

func (o *PoisonMicroblockTransactionMetadataPoisonMicroblock) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"microblock_header_1",
		"microblock_header_2",
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

	varPoisonMicroblockTransactionMetadataPoisonMicroblock := _PoisonMicroblockTransactionMetadataPoisonMicroblock{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPoisonMicroblockTransactionMetadataPoisonMicroblock)

	if err != nil {
		return err
	}

	*o = PoisonMicroblockTransactionMetadataPoisonMicroblock(varPoisonMicroblockTransactionMetadataPoisonMicroblock)

	return err
}

type NullablePoisonMicroblockTransactionMetadataPoisonMicroblock struct {
	value *PoisonMicroblockTransactionMetadataPoisonMicroblock
	isSet bool
}

func (v NullablePoisonMicroblockTransactionMetadataPoisonMicroblock) Get() *PoisonMicroblockTransactionMetadataPoisonMicroblock {
	return v.value
}

func (v *NullablePoisonMicroblockTransactionMetadataPoisonMicroblock) Set(val *PoisonMicroblockTransactionMetadataPoisonMicroblock) {
	v.value = val
	v.isSet = true
}

func (v NullablePoisonMicroblockTransactionMetadataPoisonMicroblock) IsSet() bool {
	return v.isSet
}

func (v *NullablePoisonMicroblockTransactionMetadataPoisonMicroblock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoisonMicroblockTransactionMetadataPoisonMicroblock(val *PoisonMicroblockTransactionMetadataPoisonMicroblock) *NullablePoisonMicroblockTransactionMetadataPoisonMicroblock {
	return &NullablePoisonMicroblockTransactionMetadataPoisonMicroblock{value: val, isSet: true}
}

func (v NullablePoisonMicroblockTransactionMetadataPoisonMicroblock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoisonMicroblockTransactionMetadataPoisonMicroblock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
