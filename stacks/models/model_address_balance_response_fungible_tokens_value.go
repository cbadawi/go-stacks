package models

import (
	"encoding/json"
	"fmt"
)

// AddressBalanceResponseFungibleTokensValue struct for AddressBalanceResponseFungibleTokensValue
type AddressBalanceResponseFungibleTokensValue struct {
	FtBalance *FtBalance
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AddressBalanceResponseFungibleTokensValue) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into FtBalance
	err = json.Unmarshal(data, &dst.FtBalance)
	if err == nil {
		jsonFtBalance, _ := json.Marshal(dst.FtBalance)
		if string(jsonFtBalance) == "{}" { // empty struct
			dst.FtBalance = nil
		} else {
			return nil // data stored in dst.FtBalance, return on the first match
		}
	} else {
		dst.FtBalance = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(AddressBalanceResponseFungibleTokensValue)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AddressBalanceResponseFungibleTokensValue) MarshalJSON() ([]byte, error) {
	if src.FtBalance != nil {
		return json.Marshal(&src.FtBalance)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAddressBalanceResponseFungibleTokensValue struct {
	value *AddressBalanceResponseFungibleTokensValue
	isSet bool
}

func (v NullableAddressBalanceResponseFungibleTokensValue) Get() *AddressBalanceResponseFungibleTokensValue {
	return v.value
}

func (v *NullableAddressBalanceResponseFungibleTokensValue) Set(val *AddressBalanceResponseFungibleTokensValue) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressBalanceResponseFungibleTokensValue) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressBalanceResponseFungibleTokensValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressBalanceResponseFungibleTokensValue(val *AddressBalanceResponseFungibleTokensValue) *NullableAddressBalanceResponseFungibleTokensValue {
	return &NullableAddressBalanceResponseFungibleTokensValue{value: val, isSet: true}
}

func (v NullableAddressBalanceResponseFungibleTokensValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressBalanceResponseFungibleTokensValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
