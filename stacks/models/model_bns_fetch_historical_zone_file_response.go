package models

import (
	"encoding/json"
	"fmt"
)

// BnsFetchHistoricalZoneFileResponse Fetches the historical zonefile specified by the username and zone hash.
type BnsFetchHistoricalZoneFileResponse struct {
	BnsFetchHistoricalZoneFileResponseAnyOf  *BnsFetchHistoricalZoneFileResponseAnyOf
	BnsFetchHistoricalZoneFileResponseAnyOf1 *BnsFetchHistoricalZoneFileResponseAnyOf1
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *BnsFetchHistoricalZoneFileResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into BnsFetchHistoricalZoneFileResponseAnyOf
	err = json.Unmarshal(data, &dst.BnsFetchHistoricalZoneFileResponseAnyOf)
	if err == nil {
		jsonBnsFetchHistoricalZoneFileResponseAnyOf, _ := json.Marshal(dst.BnsFetchHistoricalZoneFileResponseAnyOf)
		if string(jsonBnsFetchHistoricalZoneFileResponseAnyOf) == "{}" { // empty struct
			dst.BnsFetchHistoricalZoneFileResponseAnyOf = nil
		} else {
			return nil // data stored in dst.BnsFetchHistoricalZoneFileResponseAnyOf, return on the first match
		}
	} else {
		dst.BnsFetchHistoricalZoneFileResponseAnyOf = nil
	}

	// try to unmarshal JSON data into BnsFetchHistoricalZoneFileResponseAnyOf1
	err = json.Unmarshal(data, &dst.BnsFetchHistoricalZoneFileResponseAnyOf1)
	if err == nil {
		jsonBnsFetchHistoricalZoneFileResponseAnyOf1, _ := json.Marshal(dst.BnsFetchHistoricalZoneFileResponseAnyOf1)
		if string(jsonBnsFetchHistoricalZoneFileResponseAnyOf1) == "{}" { // empty struct
			dst.BnsFetchHistoricalZoneFileResponseAnyOf1 = nil
		} else {
			return nil // data stored in dst.BnsFetchHistoricalZoneFileResponseAnyOf1, return on the first match
		}
	} else {
		dst.BnsFetchHistoricalZoneFileResponseAnyOf1 = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(BnsFetchHistoricalZoneFileResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *BnsFetchHistoricalZoneFileResponse) MarshalJSON() ([]byte, error) {
	if src.BnsFetchHistoricalZoneFileResponseAnyOf != nil {
		return json.Marshal(&src.BnsFetchHistoricalZoneFileResponseAnyOf)
	}

	if src.BnsFetchHistoricalZoneFileResponseAnyOf1 != nil {
		return json.Marshal(&src.BnsFetchHistoricalZoneFileResponseAnyOf1)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableBnsFetchHistoricalZoneFileResponse struct {
	value *BnsFetchHistoricalZoneFileResponse
	isSet bool
}

func (v NullableBnsFetchHistoricalZoneFileResponse) Get() *BnsFetchHistoricalZoneFileResponse {
	return v.value
}

func (v *NullableBnsFetchHistoricalZoneFileResponse) Set(val *BnsFetchHistoricalZoneFileResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBnsFetchHistoricalZoneFileResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBnsFetchHistoricalZoneFileResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBnsFetchHistoricalZoneFileResponse(val *BnsFetchHistoricalZoneFileResponse) *NullableBnsFetchHistoricalZoneFileResponse {
	return &NullableBnsFetchHistoricalZoneFileResponse{value: val, isSet: true}
}

func (v NullableBnsFetchHistoricalZoneFileResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBnsFetchHistoricalZoneFileResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
