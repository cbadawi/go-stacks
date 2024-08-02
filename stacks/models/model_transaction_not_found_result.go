package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the TransactionNotFoundResult type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &TransactionNotFoundResult{}

// TransactionNotFoundResult struct for TransactionNotFoundResult
type TransactionNotFoundResult struct {
	TxId string `json:"tx_id"`
}

type _TransactionNotFoundResult TransactionNotFoundResult

// NewTransactionNotFoundResult instantiates a new TransactionNotFoundResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionNotFoundResult(txId string) *TransactionNotFoundResult {
	this := TransactionNotFoundResult{}
	this.TxId = txId
	return &this
}

// NewTransactionNotFoundResultWithDefaults instantiates a new TransactionNotFoundResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionNotFoundResultWithDefaults() *TransactionNotFoundResult {
	this := TransactionNotFoundResult{}
	return &this
}

// GetTxId returns the TxId field value
func (o *TransactionNotFoundResult) GetTxId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TxId
}

// GetTxIdOk returns a tuple with the TxId field value
// and a boolean to check if the value has been set.
func (o *TransactionNotFoundResult) GetTxIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxId, true
}

// SetTxId sets field value
func (o *TransactionNotFoundResult) SetTxId(v string) {
	o.TxId = v
}

func (o TransactionNotFoundResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionNotFoundResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tx_id"] = o.TxId
	return toSerialize, nil
}

func (o *TransactionNotFoundResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tx_id",
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

	varTransactionNotFoundResult := _TransactionNotFoundResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransactionNotFoundResult)

	if err != nil {
		return err
	}

	*o = TransactionNotFoundResult(varTransactionNotFoundResult)

	return err
}

type NullableTransactionNotFoundResult struct {
	value *TransactionNotFoundResult
	isSet bool
}

func (v NullableTransactionNotFoundResult) Get() *TransactionNotFoundResult {
	return v.value
}

func (v *NullableTransactionNotFoundResult) Set(val *TransactionNotFoundResult) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionNotFoundResult) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionNotFoundResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionNotFoundResult(val *TransactionNotFoundResult) *NullableTransactionNotFoundResult {
	return &NullableTransactionNotFoundResult{value: val, isSet: true}
}

func (v NullableTransactionNotFoundResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionNotFoundResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
