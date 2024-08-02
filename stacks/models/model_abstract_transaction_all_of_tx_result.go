package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the AbstractTransactionAllOfTxResult type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AbstractTransactionAllOfTxResult{}

// AbstractTransactionAllOfTxResult Result of the transaction. For contract calls, this will show the value returned by the call. For other transaction types, this will return a boolean indicating the success of the transaction.
type AbstractTransactionAllOfTxResult struct {
	// Hex string representing the value fo the transaction result
	Hex string `json:"hex"`
	// Readable string of the transaction result
	Repr string `json:"repr"`
}

type _AbstractTransactionAllOfTxResult AbstractTransactionAllOfTxResult

// NewAbstractTransactionAllOfTxResult instantiates a new AbstractTransactionAllOfTxResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAbstractTransactionAllOfTxResult(hex string, repr string) *AbstractTransactionAllOfTxResult {
	this := AbstractTransactionAllOfTxResult{}
	this.Hex = hex
	this.Repr = repr
	return &this
}

// NewAbstractTransactionAllOfTxResultWithDefaults instantiates a new AbstractTransactionAllOfTxResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAbstractTransactionAllOfTxResultWithDefaults() *AbstractTransactionAllOfTxResult {
	this := AbstractTransactionAllOfTxResult{}
	return &this
}

// GetHex returns the Hex field value
func (o *AbstractTransactionAllOfTxResult) GetHex() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hex
}

// GetHexOk returns a tuple with the Hex field value
// and a boolean to check if the value has been set.
func (o *AbstractTransactionAllOfTxResult) GetHexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hex, true
}

// SetHex sets field value
func (o *AbstractTransactionAllOfTxResult) SetHex(v string) {
	o.Hex = v
}

// GetRepr returns the Repr field value
func (o *AbstractTransactionAllOfTxResult) GetRepr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Repr
}

// GetReprOk returns a tuple with the Repr field value
// and a boolean to check if the value has been set.
func (o *AbstractTransactionAllOfTxResult) GetReprOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Repr, true
}

// SetRepr sets field value
func (o *AbstractTransactionAllOfTxResult) SetRepr(v string) {
	o.Repr = v
}

func (o AbstractTransactionAllOfTxResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AbstractTransactionAllOfTxResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["hex"] = o.Hex
	toSerialize["repr"] = o.Repr
	return toSerialize, nil
}

func (o *AbstractTransactionAllOfTxResult) UnmarshalJSON(data []byte) (err error) {
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

	varAbstractTransactionAllOfTxResult := _AbstractTransactionAllOfTxResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAbstractTransactionAllOfTxResult)

	if err != nil {
		return err
	}

	*o = AbstractTransactionAllOfTxResult(varAbstractTransactionAllOfTxResult)

	return err
}

type NullableAbstractTransactionAllOfTxResult struct {
	value *AbstractTransactionAllOfTxResult
	isSet bool
}

func (v NullableAbstractTransactionAllOfTxResult) Get() *AbstractTransactionAllOfTxResult {
	return v.value
}

func (v *NullableAbstractTransactionAllOfTxResult) Set(val *AbstractTransactionAllOfTxResult) {
	v.value = val
	v.isSet = true
}

func (v NullableAbstractTransactionAllOfTxResult) IsSet() bool {
	return v.isSet
}

func (v *NullableAbstractTransactionAllOfTxResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAbstractTransactionAllOfTxResult(val *AbstractTransactionAllOfTxResult) *NullableAbstractTransactionAllOfTxResult {
	return &NullableAbstractTransactionAllOfTxResult{value: val, isSet: true}
}

func (v NullableAbstractTransactionAllOfTxResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAbstractTransactionAllOfTxResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
