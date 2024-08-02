package models

import (
	"encoding/json"
	"fmt"
)

// TransactionList struct for TransactionList
type TransactionList struct {
	TransactionFound    *TransactionFound
	TransactionNotFound *TransactionNotFound
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *TransactionList) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into TransactionFound
	err = json.Unmarshal(data, &dst.TransactionFound)
	if err == nil {
		jsonTransactionFound, _ := json.Marshal(dst.TransactionFound)
		if string(jsonTransactionFound) == "{}" { // empty struct
			dst.TransactionFound = nil
		} else {
			return nil // data stored in dst.TransactionFound, return on the first match
		}
	} else {
		dst.TransactionFound = nil
	}

	// try to unmarshal JSON data into TransactionNotFound
	err = json.Unmarshal(data, &dst.TransactionNotFound)
	if err == nil {
		jsonTransactionNotFound, _ := json.Marshal(dst.TransactionNotFound)
		if string(jsonTransactionNotFound) == "{}" { // empty struct
			dst.TransactionNotFound = nil
		} else {
			return nil // data stored in dst.TransactionNotFound, return on the first match
		}
	} else {
		dst.TransactionNotFound = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(TransactionList)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *TransactionList) MarshalJSON() ([]byte, error) {
	if src.TransactionFound != nil {
		return json.Marshal(&src.TransactionFound)
	}

	if src.TransactionNotFound != nil {
		return json.Marshal(&src.TransactionNotFound)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableTransactionList struct {
	value *TransactionList
	isSet bool
}

func (v NullableTransactionList) Get() *TransactionList {
	return v.value
}

func (v *NullableTransactionList) Set(val *TransactionList) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionList) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionList(val *TransactionList) *NullableTransactionList {
	return &NullableTransactionList{value: val, isSet: true}
}

func (v NullableTransactionList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
