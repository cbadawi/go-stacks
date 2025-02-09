package models

import (
	"encoding/json"
	"fmt"
)

// SearchSuccessResult Search success result
type SearchSuccessResult struct {
	AddressSearchResult   *AddressSearchResult
	BlockSearchResult     *BlockSearchResult
	ContractSearchResult  *ContractSearchResult
	MempoolTxSearchResult *MempoolTxSearchResult
	TxSearchResult        *TxSearchResult
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *SearchSuccessResult) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AddressSearchResult
	err = json.Unmarshal(data, &dst.AddressSearchResult)
	if err == nil {
		jsonAddressSearchResult, _ := json.Marshal(dst.AddressSearchResult)
		if string(jsonAddressSearchResult) == "{}" { // empty struct
			dst.AddressSearchResult = nil
		} else {
			return nil // data stored in dst.AddressSearchResult, return on the first match
		}
	} else {
		dst.AddressSearchResult = nil
	}

	// try to unmarshal JSON data into BlockSearchResult
	err = json.Unmarshal(data, &dst.BlockSearchResult)
	if err == nil {
		jsonBlockSearchResult, _ := json.Marshal(dst.BlockSearchResult)
		if string(jsonBlockSearchResult) == "{}" { // empty struct
			dst.BlockSearchResult = nil
		} else {
			return nil // data stored in dst.BlockSearchResult, return on the first match
		}
	} else {
		dst.BlockSearchResult = nil
	}

	// try to unmarshal JSON data into ContractSearchResult
	err = json.Unmarshal(data, &dst.ContractSearchResult)
	if err == nil {
		jsonContractSearchResult, _ := json.Marshal(dst.ContractSearchResult)
		if string(jsonContractSearchResult) == "{}" { // empty struct
			dst.ContractSearchResult = nil
		} else {
			return nil // data stored in dst.ContractSearchResult, return on the first match
		}
	} else {
		dst.ContractSearchResult = nil
	}

	// try to unmarshal JSON data into MempoolTxSearchResult
	err = json.Unmarshal(data, &dst.MempoolTxSearchResult)
	if err == nil {
		jsonMempoolTxSearchResult, _ := json.Marshal(dst.MempoolTxSearchResult)
		if string(jsonMempoolTxSearchResult) == "{}" { // empty struct
			dst.MempoolTxSearchResult = nil
		} else {
			return nil // data stored in dst.MempoolTxSearchResult, return on the first match
		}
	} else {
		dst.MempoolTxSearchResult = nil
	}

	// try to unmarshal JSON data into TxSearchResult
	err = json.Unmarshal(data, &dst.TxSearchResult)
	if err == nil {
		jsonTxSearchResult, _ := json.Marshal(dst.TxSearchResult)
		if string(jsonTxSearchResult) == "{}" { // empty struct
			dst.TxSearchResult = nil
		} else {
			return nil // data stored in dst.TxSearchResult, return on the first match
		}
	} else {
		dst.TxSearchResult = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(SearchSuccessResult)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *SearchSuccessResult) MarshalJSON() ([]byte, error) {
	if src.AddressSearchResult != nil {
		return json.Marshal(&src.AddressSearchResult)
	}

	if src.BlockSearchResult != nil {
		return json.Marshal(&src.BlockSearchResult)
	}

	if src.ContractSearchResult != nil {
		return json.Marshal(&src.ContractSearchResult)
	}

	if src.MempoolTxSearchResult != nil {
		return json.Marshal(&src.MempoolTxSearchResult)
	}

	if src.TxSearchResult != nil {
		return json.Marshal(&src.TxSearchResult)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableSearchSuccessResult struct {
	value *SearchSuccessResult
	isSet bool
}

func (v NullableSearchSuccessResult) Get() *SearchSuccessResult {
	return v.value
}

func (v *NullableSearchSuccessResult) Set(val *SearchSuccessResult) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchSuccessResult) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchSuccessResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchSuccessResult(val *SearchSuccessResult) *NullableSearchSuccessResult {
	return &NullableSearchSuccessResult{value: val, isSet: true}
}

func (v NullableSearchSuccessResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchSuccessResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
