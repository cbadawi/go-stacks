package models

import (
	"encoding/json"
	"fmt"
)

// AddressBalanceResponseNonFungibleTokensValue struct for AddressBalanceResponseNonFungibleTokensValue
type AddressBalanceResponseNonFungibleTokensValue struct {
	NftBalance *NftBalance
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AddressBalanceResponseNonFungibleTokensValue) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into NftBalance
	err = json.Unmarshal(data, &dst.NftBalance)
	if err == nil {
		jsonNftBalance, _ := json.Marshal(dst.NftBalance)
		if string(jsonNftBalance) == "{}" { // empty struct
			dst.NftBalance = nil
		} else {
			return nil // data stored in dst.NftBalance, return on the first match
		}
	} else {
		dst.NftBalance = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(AddressBalanceResponseNonFungibleTokensValue)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AddressBalanceResponseNonFungibleTokensValue) MarshalJSON() ([]byte, error) {
	if src.NftBalance != nil {
		return json.Marshal(&src.NftBalance)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAddressBalanceResponseNonFungibleTokensValue struct {
	value *AddressBalanceResponseNonFungibleTokensValue
	isSet bool
}

func (v NullableAddressBalanceResponseNonFungibleTokensValue) Get() *AddressBalanceResponseNonFungibleTokensValue {
	return v.value
}

func (v *NullableAddressBalanceResponseNonFungibleTokensValue) Set(val *AddressBalanceResponseNonFungibleTokensValue) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressBalanceResponseNonFungibleTokensValue) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressBalanceResponseNonFungibleTokensValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressBalanceResponseNonFungibleTokensValue(val *AddressBalanceResponseNonFungibleTokensValue) *NullableAddressBalanceResponseNonFungibleTokensValue {
	return &NullableAddressBalanceResponseNonFungibleTokensValue{value: val, isSet: true}
}

func (v NullableAddressBalanceResponseNonFungibleTokensValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressBalanceResponseNonFungibleTokensValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
