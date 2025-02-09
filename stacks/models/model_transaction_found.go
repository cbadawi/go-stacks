package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the TransactionFound type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &TransactionFound{}

// TransactionFound This object returns transaction for found true
type TransactionFound struct {
	Found  bool                   `json:"found"`
	Result TransactionFoundResult `json:"result"`
}

type _TransactionFound TransactionFound

// NewTransactionFound instantiates a new TransactionFound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionFound(found bool, result TransactionFoundResult) *TransactionFound {
	this := TransactionFound{}
	this.Found = found
	this.Result = result
	return &this
}

// NewTransactionFoundWithDefaults instantiates a new TransactionFound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionFoundWithDefaults() *TransactionFound {
	this := TransactionFound{}
	return &this
}

// GetFound returns the Found field value
func (o *TransactionFound) GetFound() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Found
}

// GetFoundOk returns a tuple with the Found field value
// and a boolean to check if the value has been set.
func (o *TransactionFound) GetFoundOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Found, true
}

// SetFound sets field value
func (o *TransactionFound) SetFound(v bool) {
	o.Found = v
}

// GetResult returns the Result field value
func (o *TransactionFound) GetResult() TransactionFoundResult {
	if o == nil {
		var ret TransactionFoundResult
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *TransactionFound) GetResultOk() (*TransactionFoundResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *TransactionFound) SetResult(v TransactionFoundResult) {
	o.Result = v
}

func (o TransactionFound) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionFound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["found"] = o.Found
	toSerialize["result"] = o.Result
	return toSerialize, nil
}

func (o *TransactionFound) UnmarshalJSON(data []byte) (err error) {
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

	varTransactionFound := _TransactionFound{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransactionFound)

	if err != nil {
		return err
	}

	*o = TransactionFound(varTransactionFound)

	return err
}

type NullableTransactionFound struct {
	value *TransactionFound
	isSet bool
}

func (v NullableTransactionFound) Get() *TransactionFound {
	return v.value
}

func (v *NullableTransactionFound) Set(val *TransactionFound) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionFound) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionFound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionFound(val *TransactionFound) *NullableTransactionFound {
	return &NullableTransactionFound{value: val, isSet: true}
}

func (v NullableTransactionFound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionFound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
