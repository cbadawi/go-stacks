package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the NonFungibleTokenHoldingWithTxIdValue type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &NonFungibleTokenHoldingWithTxIdValue{}

// NonFungibleTokenHoldingWithTxIdValue Non-Fungible Token value
type NonFungibleTokenHoldingWithTxIdValue struct {
	// Hex string representing the identifier of the Non-Fungible Token
	Hex string `json:"hex"`
	// Readable string of the Non-Fungible Token identifier
	Repr string `json:"repr"`
}

type _NonFungibleTokenHoldingWithTxIdValue NonFungibleTokenHoldingWithTxIdValue

// NewNonFungibleTokenHoldingWithTxIdValue instantiates a new NonFungibleTokenHoldingWithTxIdValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNonFungibleTokenHoldingWithTxIdValue(hex string, repr string) *NonFungibleTokenHoldingWithTxIdValue {
	this := NonFungibleTokenHoldingWithTxIdValue{}
	this.Hex = hex
	this.Repr = repr
	return &this
}

// NewNonFungibleTokenHoldingWithTxIdValueWithDefaults instantiates a new NonFungibleTokenHoldingWithTxIdValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNonFungibleTokenHoldingWithTxIdValueWithDefaults() *NonFungibleTokenHoldingWithTxIdValue {
	this := NonFungibleTokenHoldingWithTxIdValue{}
	return &this
}

// GetHex returns the Hex field value
func (o *NonFungibleTokenHoldingWithTxIdValue) GetHex() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hex
}

// GetHexOk returns a tuple with the Hex field value
// and a boolean to check if the value has been set.
func (o *NonFungibleTokenHoldingWithTxIdValue) GetHexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hex, true
}

// SetHex sets field value
func (o *NonFungibleTokenHoldingWithTxIdValue) SetHex(v string) {
	o.Hex = v
}

// GetRepr returns the Repr field value
func (o *NonFungibleTokenHoldingWithTxIdValue) GetRepr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Repr
}

// GetReprOk returns a tuple with the Repr field value
// and a boolean to check if the value has been set.
func (o *NonFungibleTokenHoldingWithTxIdValue) GetReprOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Repr, true
}

// SetRepr sets field value
func (o *NonFungibleTokenHoldingWithTxIdValue) SetRepr(v string) {
	o.Repr = v
}

func (o NonFungibleTokenHoldingWithTxIdValue) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NonFungibleTokenHoldingWithTxIdValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["hex"] = o.Hex
	toSerialize["repr"] = o.Repr
	return toSerialize, nil
}

func (o *NonFungibleTokenHoldingWithTxIdValue) UnmarshalJSON(data []byte) (err error) {
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

	varNonFungibleTokenHoldingWithTxIdValue := _NonFungibleTokenHoldingWithTxIdValue{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNonFungibleTokenHoldingWithTxIdValue)

	if err != nil {
		return err
	}

	*o = NonFungibleTokenHoldingWithTxIdValue(varNonFungibleTokenHoldingWithTxIdValue)

	return err
}

type NullableNonFungibleTokenHoldingWithTxIdValue struct {
	value *NonFungibleTokenHoldingWithTxIdValue
	isSet bool
}

func (v NullableNonFungibleTokenHoldingWithTxIdValue) Get() *NonFungibleTokenHoldingWithTxIdValue {
	return v.value
}

func (v *NullableNonFungibleTokenHoldingWithTxIdValue) Set(val *NonFungibleTokenHoldingWithTxIdValue) {
	v.value = val
	v.isSet = true
}

func (v NullableNonFungibleTokenHoldingWithTxIdValue) IsSet() bool {
	return v.isSet
}

func (v *NullableNonFungibleTokenHoldingWithTxIdValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNonFungibleTokenHoldingWithTxIdValue(val *NonFungibleTokenHoldingWithTxIdValue) *NullableNonFungibleTokenHoldingWithTxIdValue {
	return &NullableNonFungibleTokenHoldingWithTxIdValue{value: val, isSet: true}
}

func (v NullableNonFungibleTokenHoldingWithTxIdValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNonFungibleTokenHoldingWithTxIdValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
