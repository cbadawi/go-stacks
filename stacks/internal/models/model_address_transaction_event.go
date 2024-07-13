package models

import (
	"encoding/json"
	"fmt"
)

// AddressTransactionEvent Address Transaction Event
type AddressTransactionEvent struct {
	AddressTransactionEventAnyOf  *AddressTransactionEventAnyOf
	AddressTransactionEventAnyOf1 *AddressTransactionEventAnyOf1
	AddressTransactionEventAnyOf2 *AddressTransactionEventAnyOf2
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AddressTransactionEvent) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AddressTransactionEventAnyOf
	err = json.Unmarshal(data, &dst.AddressTransactionEventAnyOf)
	if err == nil {
		jsonAddressTransactionEventAnyOf, _ := json.Marshal(dst.AddressTransactionEventAnyOf)
		if string(jsonAddressTransactionEventAnyOf) == "{}" { // empty struct
			dst.AddressTransactionEventAnyOf = nil
		} else {
			return nil // data stored in dst.AddressTransactionEventAnyOf, return on the first match
		}
	} else {
		dst.AddressTransactionEventAnyOf = nil
	}

	// try to unmarshal JSON data into AddressTransactionEventAnyOf1
	err = json.Unmarshal(data, &dst.AddressTransactionEventAnyOf1)
	if err == nil {
		jsonAddressTransactionEventAnyOf1, _ := json.Marshal(dst.AddressTransactionEventAnyOf1)
		if string(jsonAddressTransactionEventAnyOf1) == "{}" { // empty struct
			dst.AddressTransactionEventAnyOf1 = nil
		} else {
			return nil // data stored in dst.AddressTransactionEventAnyOf1, return on the first match
		}
	} else {
		dst.AddressTransactionEventAnyOf1 = nil
	}

	// try to unmarshal JSON data into AddressTransactionEventAnyOf2
	err = json.Unmarshal(data, &dst.AddressTransactionEventAnyOf2)
	if err == nil {
		jsonAddressTransactionEventAnyOf2, _ := json.Marshal(dst.AddressTransactionEventAnyOf2)
		if string(jsonAddressTransactionEventAnyOf2) == "{}" { // empty struct
			dst.AddressTransactionEventAnyOf2 = nil
		} else {
			return nil // data stored in dst.AddressTransactionEventAnyOf2, return on the first match
		}
	} else {
		dst.AddressTransactionEventAnyOf2 = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(AddressTransactionEvent)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AddressTransactionEvent) MarshalJSON() ([]byte, error) {
	if src.AddressTransactionEventAnyOf != nil {
		return json.Marshal(&src.AddressTransactionEventAnyOf)
	}

	if src.AddressTransactionEventAnyOf1 != nil {
		return json.Marshal(&src.AddressTransactionEventAnyOf1)
	}

	if src.AddressTransactionEventAnyOf2 != nil {
		return json.Marshal(&src.AddressTransactionEventAnyOf2)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAddressTransactionEvent struct {
	value *AddressTransactionEvent
	isSet bool
}

func (v NullableAddressTransactionEvent) Get() *AddressTransactionEvent {
	return v.value
}

func (v *NullableAddressTransactionEvent) Set(val *AddressTransactionEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressTransactionEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressTransactionEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressTransactionEvent(val *AddressTransactionEvent) *NullableAddressTransactionEvent {
	return &NullableAddressTransactionEvent{value: val, isSet: true}
}

func (v NullableAddressTransactionEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressTransactionEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
