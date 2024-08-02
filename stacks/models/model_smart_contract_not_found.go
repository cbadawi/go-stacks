package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the SmartContractNotFound type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &SmartContractNotFound{}

// SmartContractNotFound struct for SmartContractNotFound
type SmartContractNotFound struct {
	Found bool `json:"found"`
}

type _SmartContractNotFound SmartContractNotFound

// NewSmartContractNotFound instantiates a new SmartContractNotFound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmartContractNotFound(found bool) *SmartContractNotFound {
	this := SmartContractNotFound{}
	this.Found = found
	return &this
}

// NewSmartContractNotFoundWithDefaults instantiates a new SmartContractNotFound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmartContractNotFoundWithDefaults() *SmartContractNotFound {
	this := SmartContractNotFound{}
	return &this
}

// GetFound returns the Found field value
func (o *SmartContractNotFound) GetFound() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Found
}

// GetFoundOk returns a tuple with the Found field value
// and a boolean to check if the value has been set.
func (o *SmartContractNotFound) GetFoundOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Found, true
}

// SetFound sets field value
func (o *SmartContractNotFound) SetFound(v bool) {
	o.Found = v
}

func (o SmartContractNotFound) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SmartContractNotFound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["found"] = o.Found
	return toSerialize, nil
}

func (o *SmartContractNotFound) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"found",
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

	varSmartContractNotFound := _SmartContractNotFound{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSmartContractNotFound)

	if err != nil {
		return err
	}

	*o = SmartContractNotFound(varSmartContractNotFound)

	return err
}

type NullableSmartContractNotFound struct {
	value *SmartContractNotFound
	isSet bool
}

func (v NullableSmartContractNotFound) Get() *SmartContractNotFound {
	return v.value
}

func (v *NullableSmartContractNotFound) Set(val *SmartContractNotFound) {
	v.value = val
	v.isSet = true
}

func (v NullableSmartContractNotFound) IsSet() bool {
	return v.isSet
}

func (v *NullableSmartContractNotFound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmartContractNotFound(val *SmartContractNotFound) *NullableSmartContractNotFound {
	return &NullableSmartContractNotFound{value: val, isSet: true}
}

func (v NullableSmartContractNotFound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmartContractNotFound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
