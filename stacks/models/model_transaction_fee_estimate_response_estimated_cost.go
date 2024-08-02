package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the TransactionFeeEstimateResponseEstimatedCost type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &TransactionFeeEstimateResponseEstimatedCost{}

// TransactionFeeEstimateResponseEstimatedCost struct for TransactionFeeEstimateResponseEstimatedCost
type TransactionFeeEstimateResponseEstimatedCost struct {
	ReadCount   int32 `json:"read_count"`
	ReadLength  int32 `json:"read_length"`
	Runtime     int32 `json:"runtime"`
	WriteCount  int32 `json:"write_count"`
	WriteLength int32 `json:"write_length"`
}

type _TransactionFeeEstimateResponseEstimatedCost TransactionFeeEstimateResponseEstimatedCost

// NewTransactionFeeEstimateResponseEstimatedCost instantiates a new TransactionFeeEstimateResponseEstimatedCost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionFeeEstimateResponseEstimatedCost(readCount int32, readLength int32, runtime int32, writeCount int32, writeLength int32) *TransactionFeeEstimateResponseEstimatedCost {
	this := TransactionFeeEstimateResponseEstimatedCost{}
	this.ReadCount = readCount
	this.ReadLength = readLength
	this.Runtime = runtime
	this.WriteCount = writeCount
	this.WriteLength = writeLength
	return &this
}

// NewTransactionFeeEstimateResponseEstimatedCostWithDefaults instantiates a new TransactionFeeEstimateResponseEstimatedCost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionFeeEstimateResponseEstimatedCostWithDefaults() *TransactionFeeEstimateResponseEstimatedCost {
	this := TransactionFeeEstimateResponseEstimatedCost{}
	return &this
}

// GetReadCount returns the ReadCount field value
func (o *TransactionFeeEstimateResponseEstimatedCost) GetReadCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ReadCount
}

// GetReadCountOk returns a tuple with the ReadCount field value
// and a boolean to check if the value has been set.
func (o *TransactionFeeEstimateResponseEstimatedCost) GetReadCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReadCount, true
}

// SetReadCount sets field value
func (o *TransactionFeeEstimateResponseEstimatedCost) SetReadCount(v int32) {
	o.ReadCount = v
}

// GetReadLength returns the ReadLength field value
func (o *TransactionFeeEstimateResponseEstimatedCost) GetReadLength() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ReadLength
}

// GetReadLengthOk returns a tuple with the ReadLength field value
// and a boolean to check if the value has been set.
func (o *TransactionFeeEstimateResponseEstimatedCost) GetReadLengthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReadLength, true
}

// SetReadLength sets field value
func (o *TransactionFeeEstimateResponseEstimatedCost) SetReadLength(v int32) {
	o.ReadLength = v
}

// GetRuntime returns the Runtime field value
func (o *TransactionFeeEstimateResponseEstimatedCost) GetRuntime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Runtime
}

// GetRuntimeOk returns a tuple with the Runtime field value
// and a boolean to check if the value has been set.
func (o *TransactionFeeEstimateResponseEstimatedCost) GetRuntimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Runtime, true
}

// SetRuntime sets field value
func (o *TransactionFeeEstimateResponseEstimatedCost) SetRuntime(v int32) {
	o.Runtime = v
}

// GetWriteCount returns the WriteCount field value
func (o *TransactionFeeEstimateResponseEstimatedCost) GetWriteCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.WriteCount
}

// GetWriteCountOk returns a tuple with the WriteCount field value
// and a boolean to check if the value has been set.
func (o *TransactionFeeEstimateResponseEstimatedCost) GetWriteCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WriteCount, true
}

// SetWriteCount sets field value
func (o *TransactionFeeEstimateResponseEstimatedCost) SetWriteCount(v int32) {
	o.WriteCount = v
}

// GetWriteLength returns the WriteLength field value
func (o *TransactionFeeEstimateResponseEstimatedCost) GetWriteLength() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.WriteLength
}

// GetWriteLengthOk returns a tuple with the WriteLength field value
// and a boolean to check if the value has been set.
func (o *TransactionFeeEstimateResponseEstimatedCost) GetWriteLengthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WriteLength, true
}

// SetWriteLength sets field value
func (o *TransactionFeeEstimateResponseEstimatedCost) SetWriteLength(v int32) {
	o.WriteLength = v
}

func (o TransactionFeeEstimateResponseEstimatedCost) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionFeeEstimateResponseEstimatedCost) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["read_count"] = o.ReadCount
	toSerialize["read_length"] = o.ReadLength
	toSerialize["runtime"] = o.Runtime
	toSerialize["write_count"] = o.WriteCount
	toSerialize["write_length"] = o.WriteLength
	return toSerialize, nil
}

func (o *TransactionFeeEstimateResponseEstimatedCost) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"read_count",
		"read_length",
		"runtime",
		"write_count",
		"write_length",
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

	varTransactionFeeEstimateResponseEstimatedCost := _TransactionFeeEstimateResponseEstimatedCost{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransactionFeeEstimateResponseEstimatedCost)

	if err != nil {
		return err
	}

	*o = TransactionFeeEstimateResponseEstimatedCost(varTransactionFeeEstimateResponseEstimatedCost)

	return err
}

type NullableTransactionFeeEstimateResponseEstimatedCost struct {
	value *TransactionFeeEstimateResponseEstimatedCost
	isSet bool
}

func (v NullableTransactionFeeEstimateResponseEstimatedCost) Get() *TransactionFeeEstimateResponseEstimatedCost {
	return v.value
}

func (v *NullableTransactionFeeEstimateResponseEstimatedCost) Set(val *TransactionFeeEstimateResponseEstimatedCost) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionFeeEstimateResponseEstimatedCost) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionFeeEstimateResponseEstimatedCost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionFeeEstimateResponseEstimatedCost(val *TransactionFeeEstimateResponseEstimatedCost) *NullableTransactionFeeEstimateResponseEstimatedCost {
	return &NullableTransactionFeeEstimateResponseEstimatedCost{value: val, isSet: true}
}

func (v NullableTransactionFeeEstimateResponseEstimatedCost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionFeeEstimateResponseEstimatedCost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
