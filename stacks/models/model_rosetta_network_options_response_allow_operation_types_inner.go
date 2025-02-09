package models

import (
	"encoding/json"
	"fmt"
)

// RosettaNetworkOptionsResponseAllowOperationTypesInner struct for RosettaNetworkOptionsResponseAllowOperationTypesInner
type RosettaNetworkOptionsResponseAllowOperationTypesInner struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *RosettaNetworkOptionsResponseAllowOperationTypesInner) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into String
	err = json.Unmarshal(data, &dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			return nil // data stored in dst.String, return on the first match
		}
	} else {
		dst.String = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(RosettaNetworkOptionsResponseAllowOperationTypesInner)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *RosettaNetworkOptionsResponseAllowOperationTypesInner) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableRosettaNetworkOptionsResponseAllowOperationTypesInner struct {
	value *RosettaNetworkOptionsResponseAllowOperationTypesInner
	isSet bool
}

func (v NullableRosettaNetworkOptionsResponseAllowOperationTypesInner) Get() *RosettaNetworkOptionsResponseAllowOperationTypesInner {
	return v.value
}

func (v *NullableRosettaNetworkOptionsResponseAllowOperationTypesInner) Set(val *RosettaNetworkOptionsResponseAllowOperationTypesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableRosettaNetworkOptionsResponseAllowOperationTypesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableRosettaNetworkOptionsResponseAllowOperationTypesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRosettaNetworkOptionsResponseAllowOperationTypesInner(val *RosettaNetworkOptionsResponseAllowOperationTypesInner) *NullableRosettaNetworkOptionsResponseAllowOperationTypesInner {
	return &NullableRosettaNetworkOptionsResponseAllowOperationTypesInner{value: val, isSet: true}
}

func (v NullableRosettaNetworkOptionsResponseAllowOperationTypesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRosettaNetworkOptionsResponseAllowOperationTypesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
