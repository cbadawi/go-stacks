package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the SmartContractFound type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &SmartContractFound{}

// SmartContractFound struct for SmartContractFound
type SmartContractFound struct {
	Found  bool                `json:"found"`
	Result SmartContractStatus `json:"result"`
}

type _SmartContractFound SmartContractFound

// NewSmartContractFound instantiates a new SmartContractFound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmartContractFound(found bool, result SmartContractStatus) *SmartContractFound {
	this := SmartContractFound{}
	this.Found = found
	this.Result = result
	return &this
}

// NewSmartContractFoundWithDefaults instantiates a new SmartContractFound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmartContractFoundWithDefaults() *SmartContractFound {
	this := SmartContractFound{}
	return &this
}

// GetFound returns the Found field value
func (o *SmartContractFound) GetFound() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Found
}

// GetFoundOk returns a tuple with the Found field value
// and a boolean to check if the value has been set.
func (o *SmartContractFound) GetFoundOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Found, true
}

// SetFound sets field value
func (o *SmartContractFound) SetFound(v bool) {
	o.Found = v
}

// GetResult returns the Result field value
func (o *SmartContractFound) GetResult() SmartContractStatus {
	if o == nil {
		var ret SmartContractStatus
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *SmartContractFound) GetResultOk() (*SmartContractStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *SmartContractFound) SetResult(v SmartContractStatus) {
	o.Result = v
}

func (o SmartContractFound) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SmartContractFound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["found"] = o.Found
	toSerialize["result"] = o.Result
	return toSerialize, nil
}

func (o *SmartContractFound) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"found",
		"result",
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

	varSmartContractFound := _SmartContractFound{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSmartContractFound)

	if err != nil {
		return err
	}

	*o = SmartContractFound(varSmartContractFound)

	return err
}

type NullableSmartContractFound struct {
	value *SmartContractFound
	isSet bool
}

func (v NullableSmartContractFound) Get() *SmartContractFound {
	return v.value
}

func (v *NullableSmartContractFound) Set(val *SmartContractFound) {
	v.value = val
	v.isSet = true
}

func (v NullableSmartContractFound) IsSet() bool {
	return v.isSet
}

func (v *NullableSmartContractFound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmartContractFound(val *SmartContractFound) *NullableSmartContractFound {
	return &NullableSmartContractFound{value: val, isSet: true}
}

func (v NullableSmartContractFound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmartContractFound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
