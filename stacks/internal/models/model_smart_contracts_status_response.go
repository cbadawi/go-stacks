package models

import (
	"encoding/json"
	"fmt"
)

// SmartContractsStatusResponse struct for SmartContractsStatusResponse
type SmartContractsStatusResponse struct {
	SmartContractFound    *SmartContractFound
	SmartContractNotFound *SmartContractNotFound
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *SmartContractsStatusResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into SmartContractFound
	err = json.Unmarshal(data, &dst.SmartContractFound)
	if err == nil {
		jsonSmartContractFound, _ := json.Marshal(dst.SmartContractFound)
		if string(jsonSmartContractFound) == "{}" { // empty struct
			dst.SmartContractFound = nil
		} else {
			return nil // data stored in dst.SmartContractFound, return on the first match
		}
	} else {
		dst.SmartContractFound = nil
	}

	// try to unmarshal JSON data into SmartContractNotFound
	err = json.Unmarshal(data, &dst.SmartContractNotFound)
	if err == nil {
		jsonSmartContractNotFound, _ := json.Marshal(dst.SmartContractNotFound)
		if string(jsonSmartContractNotFound) == "{}" { // empty struct
			dst.SmartContractNotFound = nil
		} else {
			return nil // data stored in dst.SmartContractNotFound, return on the first match
		}
	} else {
		dst.SmartContractNotFound = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(SmartContractsStatusResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *SmartContractsStatusResponse) MarshalJSON() ([]byte, error) {
	if src.SmartContractFound != nil {
		return json.Marshal(&src.SmartContractFound)
	}

	if src.SmartContractNotFound != nil {
		return json.Marshal(&src.SmartContractNotFound)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableSmartContractsStatusResponse struct {
	value *SmartContractsStatusResponse
	isSet bool
}

func (v NullableSmartContractsStatusResponse) Get() *SmartContractsStatusResponse {
	return v.value
}

func (v *NullableSmartContractsStatusResponse) Set(val *SmartContractsStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSmartContractsStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSmartContractsStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmartContractsStatusResponse(val *SmartContractsStatusResponse) *NullableSmartContractsStatusResponse {
	return &NullableSmartContractsStatusResponse{value: val, isSet: true}
}

func (v NullableSmartContractsStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmartContractsStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
