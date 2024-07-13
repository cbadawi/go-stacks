package models

import (
	"encoding/json"
	"fmt"
)

// TransactionFoundResult struct for TransactionFoundResult
type TransactionFoundResult struct {
	MempoolTransaction *MempoolTransaction
	Transaction        *Transaction
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *TransactionFoundResult) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into MempoolTransaction
	err = json.Unmarshal(data, &dst.MempoolTransaction)
	if err == nil {
		jsonMempoolTransaction, _ := json.Marshal(dst.MempoolTransaction)
		if string(jsonMempoolTransaction) == "{}" { // empty struct
			dst.MempoolTransaction = nil
		} else {
			return nil // data stored in dst.MempoolTransaction, return on the first match
		}
	} else {
		dst.MempoolTransaction = nil
	}

	// try to unmarshal JSON data into Transaction
	err = json.Unmarshal(data, &dst.Transaction)
	if err == nil {
		jsonTransaction, _ := json.Marshal(dst.Transaction)
		if string(jsonTransaction) == "{}" { // empty struct
			dst.Transaction = nil
		} else {
			return nil // data stored in dst.Transaction, return on the first match
		}
	} else {
		dst.Transaction = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(TransactionFoundResult)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *TransactionFoundResult) MarshalJSON() ([]byte, error) {
	if src.MempoolTransaction != nil {
		return json.Marshal(&src.MempoolTransaction)
	}

	if src.Transaction != nil {
		return json.Marshal(&src.Transaction)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableTransactionFoundResult struct {
	value *TransactionFoundResult
	isSet bool
}

func (v NullableTransactionFoundResult) Get() *TransactionFoundResult {
	return v.value
}

func (v *NullableTransactionFoundResult) Set(val *TransactionFoundResult) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionFoundResult) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionFoundResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionFoundResult(val *TransactionFoundResult) *NullableTransactionFoundResult {
	return &NullableTransactionFoundResult{value: val, isSet: true}
}

func (v NullableTransactionFoundResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionFoundResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
