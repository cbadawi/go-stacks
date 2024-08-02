package models

import (
	"encoding/json"
	"fmt"
)

// BnsFetchFileZoneResponse Fetch a user's raw zone file. This only works for RFC-compliant zone files. This method returns an error for names that have non-standard zone files.
type BnsFetchFileZoneResponse struct {
	BnsFetchFileZoneResponseAnyOf  *BnsFetchFileZoneResponseAnyOf
	BnsFetchFileZoneResponseAnyOf1 *BnsFetchFileZoneResponseAnyOf1
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *BnsFetchFileZoneResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into BnsFetchFileZoneResponseAnyOf
	err = json.Unmarshal(data, &dst.BnsFetchFileZoneResponseAnyOf)
	if err == nil {
		jsonBnsFetchFileZoneResponseAnyOf, _ := json.Marshal(dst.BnsFetchFileZoneResponseAnyOf)
		if string(jsonBnsFetchFileZoneResponseAnyOf) == "{}" { // empty struct
			dst.BnsFetchFileZoneResponseAnyOf = nil
		} else {
			return nil // data stored in dst.BnsFetchFileZoneResponseAnyOf, return on the first match
		}
	} else {
		dst.BnsFetchFileZoneResponseAnyOf = nil
	}

	// try to unmarshal JSON data into BnsFetchFileZoneResponseAnyOf1
	err = json.Unmarshal(data, &dst.BnsFetchFileZoneResponseAnyOf1)
	if err == nil {
		jsonBnsFetchFileZoneResponseAnyOf1, _ := json.Marshal(dst.BnsFetchFileZoneResponseAnyOf1)
		if string(jsonBnsFetchFileZoneResponseAnyOf1) == "{}" { // empty struct
			dst.BnsFetchFileZoneResponseAnyOf1 = nil
		} else {
			return nil // data stored in dst.BnsFetchFileZoneResponseAnyOf1, return on the first match
		}
	} else {
		dst.BnsFetchFileZoneResponseAnyOf1 = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(BnsFetchFileZoneResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *BnsFetchFileZoneResponse) MarshalJSON() ([]byte, error) {
	if src.BnsFetchFileZoneResponseAnyOf != nil {
		return json.Marshal(&src.BnsFetchFileZoneResponseAnyOf)
	}

	if src.BnsFetchFileZoneResponseAnyOf1 != nil {
		return json.Marshal(&src.BnsFetchFileZoneResponseAnyOf1)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableBnsFetchFileZoneResponse struct {
	value *BnsFetchFileZoneResponse
	isSet bool
}

func (v NullableBnsFetchFileZoneResponse) Get() *BnsFetchFileZoneResponse {
	return v.value
}

func (v *NullableBnsFetchFileZoneResponse) Set(val *BnsFetchFileZoneResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBnsFetchFileZoneResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBnsFetchFileZoneResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBnsFetchFileZoneResponse(val *BnsFetchFileZoneResponse) *NullableBnsFetchFileZoneResponse {
	return &NullableBnsFetchFileZoneResponse{value: val, isSet: true}
}

func (v NullableBnsFetchFileZoneResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBnsFetchFileZoneResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
