package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the ContractCallTransactionMetadataContractCallFunctionArgsInner type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ContractCallTransactionMetadataContractCallFunctionArgsInner{}

// ContractCallTransactionMetadataContractCallFunctionArgsInner struct for ContractCallTransactionMetadataContractCallFunctionArgsInner
type ContractCallTransactionMetadataContractCallFunctionArgsInner struct {
	Hex  string `json:"hex"`
	Repr string `json:"repr"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type _ContractCallTransactionMetadataContractCallFunctionArgsInner ContractCallTransactionMetadataContractCallFunctionArgsInner

// NewContractCallTransactionMetadataContractCallFunctionArgsInner instantiates a new ContractCallTransactionMetadataContractCallFunctionArgsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContractCallTransactionMetadataContractCallFunctionArgsInner(hex string, repr string, name string, type_ string) *ContractCallTransactionMetadataContractCallFunctionArgsInner {
	this := ContractCallTransactionMetadataContractCallFunctionArgsInner{}
	this.Hex = hex
	this.Repr = repr
	this.Name = name
	this.Type = type_
	return &this
}

// NewContractCallTransactionMetadataContractCallFunctionArgsInnerWithDefaults instantiates a new ContractCallTransactionMetadataContractCallFunctionArgsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractCallTransactionMetadataContractCallFunctionArgsInnerWithDefaults() *ContractCallTransactionMetadataContractCallFunctionArgsInner {
	this := ContractCallTransactionMetadataContractCallFunctionArgsInner{}
	return &this
}

// GetHex returns the Hex field value
func (o *ContractCallTransactionMetadataContractCallFunctionArgsInner) GetHex() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hex
}

// GetHexOk returns a tuple with the Hex field value
// and a boolean to check if the value has been set.
func (o *ContractCallTransactionMetadataContractCallFunctionArgsInner) GetHexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hex, true
}

// SetHex sets field value
func (o *ContractCallTransactionMetadataContractCallFunctionArgsInner) SetHex(v string) {
	o.Hex = v
}

// GetRepr returns the Repr field value
func (o *ContractCallTransactionMetadataContractCallFunctionArgsInner) GetRepr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Repr
}

// GetReprOk returns a tuple with the Repr field value
// and a boolean to check if the value has been set.
func (o *ContractCallTransactionMetadataContractCallFunctionArgsInner) GetReprOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Repr, true
}

// SetRepr sets field value
func (o *ContractCallTransactionMetadataContractCallFunctionArgsInner) SetRepr(v string) {
	o.Repr = v
}

// GetName returns the Name field value
func (o *ContractCallTransactionMetadataContractCallFunctionArgsInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ContractCallTransactionMetadataContractCallFunctionArgsInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ContractCallTransactionMetadataContractCallFunctionArgsInner) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *ContractCallTransactionMetadataContractCallFunctionArgsInner) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ContractCallTransactionMetadataContractCallFunctionArgsInner) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ContractCallTransactionMetadataContractCallFunctionArgsInner) SetType(v string) {
	o.Type = v
}

func (o ContractCallTransactionMetadataContractCallFunctionArgsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContractCallTransactionMetadataContractCallFunctionArgsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["hex"] = o.Hex
	toSerialize["repr"] = o.Repr
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *ContractCallTransactionMetadataContractCallFunctionArgsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"hex",
		"repr",
		"name",
		"type",
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

	varContractCallTransactionMetadataContractCallFunctionArgsInner := _ContractCallTransactionMetadataContractCallFunctionArgsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varContractCallTransactionMetadataContractCallFunctionArgsInner)

	if err != nil {
		return err
	}

	*o = ContractCallTransactionMetadataContractCallFunctionArgsInner(varContractCallTransactionMetadataContractCallFunctionArgsInner)

	return err
}

type NullableContractCallTransactionMetadataContractCallFunctionArgsInner struct {
	value *ContractCallTransactionMetadataContractCallFunctionArgsInner
	isSet bool
}

func (v NullableContractCallTransactionMetadataContractCallFunctionArgsInner) Get() *ContractCallTransactionMetadataContractCallFunctionArgsInner {
	return v.value
}

func (v *NullableContractCallTransactionMetadataContractCallFunctionArgsInner) Set(val *ContractCallTransactionMetadataContractCallFunctionArgsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableContractCallTransactionMetadataContractCallFunctionArgsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableContractCallTransactionMetadataContractCallFunctionArgsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractCallTransactionMetadataContractCallFunctionArgsInner(val *ContractCallTransactionMetadataContractCallFunctionArgsInner) *NullableContractCallTransactionMetadataContractCallFunctionArgsInner {
	return &NullableContractCallTransactionMetadataContractCallFunctionArgsInner{value: val, isSet: true}
}

func (v NullableContractCallTransactionMetadataContractCallFunctionArgsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractCallTransactionMetadataContractCallFunctionArgsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
