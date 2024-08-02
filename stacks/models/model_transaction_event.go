package models

import (
	"encoding/json"
	"fmt"
)

// TransactionEvent struct for TransactionEvent
type TransactionEvent struct {
	TransactionEventFungibleAsset    *TransactionEventFungibleAsset
	TransactionEventNonFungibleAsset *TransactionEventNonFungibleAsset
	TransactionEventSmartContractLog *TransactionEventSmartContractLog
	TransactionEventStxAsset         *TransactionEventStxAsset
	TransactionEventStxLock          *TransactionEventStxLock
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *TransactionEvent) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into TransactionEventFungibleAsset
	err = json.Unmarshal(data, &dst.TransactionEventFungibleAsset)
	if err == nil {
		jsonTransactionEventFungibleAsset, _ := json.Marshal(dst.TransactionEventFungibleAsset)
		if string(jsonTransactionEventFungibleAsset) == "{}" { // empty struct
			dst.TransactionEventFungibleAsset = nil
		} else {
			return nil // data stored in dst.TransactionEventFungibleAsset, return on the first match
		}
	} else {
		dst.TransactionEventFungibleAsset = nil
	}

	// try to unmarshal JSON data into TransactionEventNonFungibleAsset
	err = json.Unmarshal(data, &dst.TransactionEventNonFungibleAsset)
	if err == nil {
		jsonTransactionEventNonFungibleAsset, _ := json.Marshal(dst.TransactionEventNonFungibleAsset)
		if string(jsonTransactionEventNonFungibleAsset) == "{}" { // empty struct
			dst.TransactionEventNonFungibleAsset = nil
		} else {
			return nil // data stored in dst.TransactionEventNonFungibleAsset, return on the first match
		}
	} else {
		dst.TransactionEventNonFungibleAsset = nil
	}

	// try to unmarshal JSON data into TransactionEventSmartContractLog
	err = json.Unmarshal(data, &dst.TransactionEventSmartContractLog)
	if err == nil {
		jsonTransactionEventSmartContractLog, _ := json.Marshal(dst.TransactionEventSmartContractLog)
		if string(jsonTransactionEventSmartContractLog) == "{}" { // empty struct
			dst.TransactionEventSmartContractLog = nil
		} else {
			return nil // data stored in dst.TransactionEventSmartContractLog, return on the first match
		}
	} else {
		dst.TransactionEventSmartContractLog = nil
	}

	// try to unmarshal JSON data into TransactionEventStxAsset
	err = json.Unmarshal(data, &dst.TransactionEventStxAsset)
	if err == nil {
		jsonTransactionEventStxAsset, _ := json.Marshal(dst.TransactionEventStxAsset)
		if string(jsonTransactionEventStxAsset) == "{}" { // empty struct
			dst.TransactionEventStxAsset = nil
		} else {
			return nil // data stored in dst.TransactionEventStxAsset, return on the first match
		}
	} else {
		dst.TransactionEventStxAsset = nil
	}

	// try to unmarshal JSON data into TransactionEventStxLock
	err = json.Unmarshal(data, &dst.TransactionEventStxLock)
	if err == nil {
		jsonTransactionEventStxLock, _ := json.Marshal(dst.TransactionEventStxLock)
		if string(jsonTransactionEventStxLock) == "{}" { // empty struct
			dst.TransactionEventStxLock = nil
		} else {
			return nil // data stored in dst.TransactionEventStxLock, return on the first match
		}
	} else {
		dst.TransactionEventStxLock = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(TransactionEvent)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *TransactionEvent) MarshalJSON() ([]byte, error) {
	if src.TransactionEventFungibleAsset != nil {
		return json.Marshal(&src.TransactionEventFungibleAsset)
	}

	if src.TransactionEventNonFungibleAsset != nil {
		return json.Marshal(&src.TransactionEventNonFungibleAsset)
	}

	if src.TransactionEventSmartContractLog != nil {
		return json.Marshal(&src.TransactionEventSmartContractLog)
	}

	if src.TransactionEventStxAsset != nil {
		return json.Marshal(&src.TransactionEventStxAsset)
	}

	if src.TransactionEventStxLock != nil {
		return json.Marshal(&src.TransactionEventStxLock)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableTransactionEvent struct {
	value *TransactionEvent
	isSet bool
}

func (v NullableTransactionEvent) Get() *TransactionEvent {
	return v.value
}

func (v *NullableTransactionEvent) Set(val *TransactionEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionEvent(val *TransactionEvent) *NullableTransactionEvent {
	return &NullableTransactionEvent{value: val, isSet: true}
}

func (v NullableTransactionEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
