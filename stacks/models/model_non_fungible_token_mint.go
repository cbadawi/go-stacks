package models

import (
	"encoding/json"
	"fmt"
)

// NonFungibleTokenMint Describes the minting of a Non-Fungible Token
type NonFungibleTokenMint struct {
	NonFungibleTokenMintWithTxId       *NonFungibleTokenMintWithTxId
	NonFungibleTokenMintWithTxMetadata *NonFungibleTokenMintWithTxMetadata
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *NonFungibleTokenMint) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into NonFungibleTokenMintWithTxId
	err = json.Unmarshal(data, &dst.NonFungibleTokenMintWithTxId)
	if err == nil {
		jsonNonFungibleTokenMintWithTxId, _ := json.Marshal(dst.NonFungibleTokenMintWithTxId)
		if string(jsonNonFungibleTokenMintWithTxId) == "{}" { // empty struct
			dst.NonFungibleTokenMintWithTxId = nil
		} else {
			return nil // data stored in dst.NonFungibleTokenMintWithTxId, return on the first match
		}
	} else {
		dst.NonFungibleTokenMintWithTxId = nil
	}

	// try to unmarshal JSON data into NonFungibleTokenMintWithTxMetadata
	err = json.Unmarshal(data, &dst.NonFungibleTokenMintWithTxMetadata)
	if err == nil {
		jsonNonFungibleTokenMintWithTxMetadata, _ := json.Marshal(dst.NonFungibleTokenMintWithTxMetadata)
		if string(jsonNonFungibleTokenMintWithTxMetadata) == "{}" { // empty struct
			dst.NonFungibleTokenMintWithTxMetadata = nil
		} else {
			return nil // data stored in dst.NonFungibleTokenMintWithTxMetadata, return on the first match
		}
	} else {
		dst.NonFungibleTokenMintWithTxMetadata = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(NonFungibleTokenMint)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *NonFungibleTokenMint) MarshalJSON() ([]byte, error) {
	if src.NonFungibleTokenMintWithTxId != nil {
		return json.Marshal(&src.NonFungibleTokenMintWithTxId)
	}

	if src.NonFungibleTokenMintWithTxMetadata != nil {
		return json.Marshal(&src.NonFungibleTokenMintWithTxMetadata)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableNonFungibleTokenMint struct {
	value *NonFungibleTokenMint
	isSet bool
}

func (v NullableNonFungibleTokenMint) Get() *NonFungibleTokenMint {
	return v.value
}

func (v *NullableNonFungibleTokenMint) Set(val *NonFungibleTokenMint) {
	v.value = val
	v.isSet = true
}

func (v NullableNonFungibleTokenMint) IsSet() bool {
	return v.isSet
}

func (v *NullableNonFungibleTokenMint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNonFungibleTokenMint(val *NonFungibleTokenMint) *NullableNonFungibleTokenMint {
	return &NullableNonFungibleTokenMint{value: val, isSet: true}
}

func (v NullableNonFungibleTokenMint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNonFungibleTokenMint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
