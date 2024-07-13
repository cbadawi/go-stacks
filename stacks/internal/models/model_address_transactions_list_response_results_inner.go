package models

import (
	"encoding/json"
	"fmt"
)

// AddressTransactionsListResponseResultsInner struct for AddressTransactionsListResponseResultsInner
type AddressTransactionsListResponseResultsInner struct {
	MempoolTransaction *MempoolTransaction
	Transaction        *Transaction
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AddressTransactionsListResponseResultsInner) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(AddressTransactionsListResponseResultsInner)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AddressTransactionsListResponseResultsInner) MarshalJSON() ([]byte, error) {
	if src.MempoolTransaction != nil {
		return json.Marshal(&src.MempoolTransaction)
	}

	if src.Transaction != nil {
		return json.Marshal(&src.Transaction)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAddressTransactionsListResponseResultsInner struct {
	value *AddressTransactionsListResponseResultsInner
	isSet bool
}

func (v NullableAddressTransactionsListResponseResultsInner) Get() *AddressTransactionsListResponseResultsInner {
	return v.value
}

func (v *NullableAddressTransactionsListResponseResultsInner) Set(val *AddressTransactionsListResponseResultsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressTransactionsListResponseResultsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressTransactionsListResponseResultsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressTransactionsListResponseResultsInner(val *AddressTransactionsListResponseResultsInner) *NullableAddressTransactionsListResponseResultsInner {
	return &NullableAddressTransactionsListResponseResultsInner{value: val, isSet: true}
}

func (v NullableAddressTransactionsListResponseResultsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressTransactionsListResponseResultsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
