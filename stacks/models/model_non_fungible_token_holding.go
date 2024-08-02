package models

import (
	"encoding/json"
	"fmt"
)

// NonFungibleTokenHolding Describes the ownership of a Non-Fungible Token
type NonFungibleTokenHolding struct {
	NonFungibleTokenHoldingWithTxId       *NonFungibleTokenHoldingWithTxId
	NonFungibleTokenHoldingWithTxMetadata *NonFungibleTokenHoldingWithTxMetadata
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *NonFungibleTokenHolding) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into NonFungibleTokenHoldingWithTxId
	err = json.Unmarshal(data, &dst.NonFungibleTokenHoldingWithTxId)
	if err == nil {
		jsonNonFungibleTokenHoldingWithTxId, _ := json.Marshal(dst.NonFungibleTokenHoldingWithTxId)
		if string(jsonNonFungibleTokenHoldingWithTxId) == "{}" { // empty struct
			dst.NonFungibleTokenHoldingWithTxId = nil
		} else {
			return nil // data stored in dst.NonFungibleTokenHoldingWithTxId, return on the first match
		}
	} else {
		dst.NonFungibleTokenHoldingWithTxId = nil
	}

	// try to unmarshal JSON data into NonFungibleTokenHoldingWithTxMetadata
	err = json.Unmarshal(data, &dst.NonFungibleTokenHoldingWithTxMetadata)
	if err == nil {
		jsonNonFungibleTokenHoldingWithTxMetadata, _ := json.Marshal(dst.NonFungibleTokenHoldingWithTxMetadata)
		if string(jsonNonFungibleTokenHoldingWithTxMetadata) == "{}" { // empty struct
			dst.NonFungibleTokenHoldingWithTxMetadata = nil
		} else {
			return nil // data stored in dst.NonFungibleTokenHoldingWithTxMetadata, return on the first match
		}
	} else {
		dst.NonFungibleTokenHoldingWithTxMetadata = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(NonFungibleTokenHolding)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *NonFungibleTokenHolding) MarshalJSON() ([]byte, error) {
	if src.NonFungibleTokenHoldingWithTxId != nil {
		return json.Marshal(&src.NonFungibleTokenHoldingWithTxId)
	}

	if src.NonFungibleTokenHoldingWithTxMetadata != nil {
		return json.Marshal(&src.NonFungibleTokenHoldingWithTxMetadata)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableNonFungibleTokenHolding struct {
	value *NonFungibleTokenHolding
	isSet bool
}

func (v NullableNonFungibleTokenHolding) Get() *NonFungibleTokenHolding {
	return v.value
}

func (v *NullableNonFungibleTokenHolding) Set(val *NonFungibleTokenHolding) {
	v.value = val
	v.isSet = true
}

func (v NullableNonFungibleTokenHolding) IsSet() bool {
	return v.isSet
}

func (v *NullableNonFungibleTokenHolding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNonFungibleTokenHolding(val *NonFungibleTokenHolding) *NullableNonFungibleTokenHolding {
	return &NullableNonFungibleTokenHolding{value: val, isSet: true}
}

func (v NullableNonFungibleTokenHolding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNonFungibleTokenHolding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
