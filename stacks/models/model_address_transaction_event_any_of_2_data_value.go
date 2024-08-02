package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the AddressTransactionEventAnyOf2DataValue type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AddressTransactionEventAnyOf2DataValue{}

// AddressTransactionEventAnyOf2DataValue Non Fungible Token asset value.
type AddressTransactionEventAnyOf2DataValue struct {
	Hex  string `json:"hex"`
	Repr string `json:"repr"`
}

type _AddressTransactionEventAnyOf2DataValue AddressTransactionEventAnyOf2DataValue

// NewAddressTransactionEventAnyOf2DataValue instantiates a new AddressTransactionEventAnyOf2DataValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressTransactionEventAnyOf2DataValue(hex string, repr string) *AddressTransactionEventAnyOf2DataValue {
	this := AddressTransactionEventAnyOf2DataValue{}
	this.Hex = hex
	this.Repr = repr
	return &this
}

// NewAddressTransactionEventAnyOf2DataValueWithDefaults instantiates a new AddressTransactionEventAnyOf2DataValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressTransactionEventAnyOf2DataValueWithDefaults() *AddressTransactionEventAnyOf2DataValue {
	this := AddressTransactionEventAnyOf2DataValue{}
	return &this
}

// GetHex returns the Hex field value
func (o *AddressTransactionEventAnyOf2DataValue) GetHex() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hex
}

// GetHexOk returns a tuple with the Hex field value
// and a boolean to check if the value has been set.
func (o *AddressTransactionEventAnyOf2DataValue) GetHexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hex, true
}

// SetHex sets field value
func (o *AddressTransactionEventAnyOf2DataValue) SetHex(v string) {
	o.Hex = v
}

// GetRepr returns the Repr field value
func (o *AddressTransactionEventAnyOf2DataValue) GetRepr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Repr
}

// GetReprOk returns a tuple with the Repr field value
// and a boolean to check if the value has been set.
func (o *AddressTransactionEventAnyOf2DataValue) GetReprOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Repr, true
}

// SetRepr sets field value
func (o *AddressTransactionEventAnyOf2DataValue) SetRepr(v string) {
	o.Repr = v
}

func (o AddressTransactionEventAnyOf2DataValue) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressTransactionEventAnyOf2DataValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["hex"] = o.Hex
	toSerialize["repr"] = o.Repr
	return toSerialize, nil
}

func (o *AddressTransactionEventAnyOf2DataValue) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"hex",
		"repr",
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

	varAddressTransactionEventAnyOf2DataValue := _AddressTransactionEventAnyOf2DataValue{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddressTransactionEventAnyOf2DataValue)

	if err != nil {
		return err
	}

	*o = AddressTransactionEventAnyOf2DataValue(varAddressTransactionEventAnyOf2DataValue)

	return err
}

type NullableAddressTransactionEventAnyOf2DataValue struct {
	value *AddressTransactionEventAnyOf2DataValue
	isSet bool
}

func (v NullableAddressTransactionEventAnyOf2DataValue) Get() *AddressTransactionEventAnyOf2DataValue {
	return v.value
}

func (v *NullableAddressTransactionEventAnyOf2DataValue) Set(val *AddressTransactionEventAnyOf2DataValue) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressTransactionEventAnyOf2DataValue) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressTransactionEventAnyOf2DataValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressTransactionEventAnyOf2DataValue(val *AddressTransactionEventAnyOf2DataValue) *NullableAddressTransactionEventAnyOf2DataValue {
	return &NullableAddressTransactionEventAnyOf2DataValue{value: val, isSet: true}
}

func (v NullableAddressTransactionEventAnyOf2DataValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressTransactionEventAnyOf2DataValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
