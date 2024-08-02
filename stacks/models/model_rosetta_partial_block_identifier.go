package models

import (
	"encoding/json"
	"fmt"
)

// RosettaPartialBlockIdentifier When fetching data by BlockIdentifier, it may be possible to only specify the index or hash. If neither property is specified, it is assumed that the client is making a request at the current block.
type RosettaPartialBlockIdentifier struct {
	RosettaBlockIdentifierHash   *RosettaBlockIdentifierHash
	RosettaBlockIdentifierHeight *RosettaBlockIdentifierHeight
	MapmapOfStringAny            *map[string]interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *RosettaPartialBlockIdentifier) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into RosettaBlockIdentifierHash
	err = json.Unmarshal(data, &dst.RosettaBlockIdentifierHash)
	if err == nil {
		jsonRosettaBlockIdentifierHash, _ := json.Marshal(dst.RosettaBlockIdentifierHash)
		if string(jsonRosettaBlockIdentifierHash) == "{}" { // empty struct
			dst.RosettaBlockIdentifierHash = nil
		} else {
			return nil // data stored in dst.RosettaBlockIdentifierHash, return on the first match
		}
	} else {
		dst.RosettaBlockIdentifierHash = nil
	}

	// try to unmarshal JSON data into RosettaBlockIdentifierHeight
	err = json.Unmarshal(data, &dst.RosettaBlockIdentifierHeight)
	if err == nil {
		jsonRosettaBlockIdentifierHeight, _ := json.Marshal(dst.RosettaBlockIdentifierHeight)
		if string(jsonRosettaBlockIdentifierHeight) == "{}" { // empty struct
			dst.RosettaBlockIdentifierHeight = nil
		} else {
			return nil // data stored in dst.RosettaBlockIdentifierHeight, return on the first match
		}
	} else {
		dst.RosettaBlockIdentifierHeight = nil
	}

	// try to unmarshal JSON data into MapmapOfStringAny
	err = json.Unmarshal(data, &dst.MapmapOfStringAny)
	if err == nil {
		jsonMapmapOfStringAny, _ := json.Marshal(dst.MapmapOfStringAny)
		if string(jsonMapmapOfStringAny) == "{}" { // empty struct
			dst.MapmapOfStringAny = nil
		} else {
			return nil // data stored in dst.MapmapOfStringAny, return on the first match
		}
	} else {
		dst.MapmapOfStringAny = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(RosettaPartialBlockIdentifier)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *RosettaPartialBlockIdentifier) MarshalJSON() ([]byte, error) {
	if src.RosettaBlockIdentifierHash != nil {
		return json.Marshal(&src.RosettaBlockIdentifierHash)
	}

	if src.RosettaBlockIdentifierHeight != nil {
		return json.Marshal(&src.RosettaBlockIdentifierHeight)
	}

	if src.MapmapOfStringAny != nil {
		return json.Marshal(&src.MapmapOfStringAny)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableRosettaPartialBlockIdentifier struct {
	value *RosettaPartialBlockIdentifier
	isSet bool
}

func (v NullableRosettaPartialBlockIdentifier) Get() *RosettaPartialBlockIdentifier {
	return v.value
}

func (v *NullableRosettaPartialBlockIdentifier) Set(val *RosettaPartialBlockIdentifier) {
	v.value = val
	v.isSet = true
}

func (v NullableRosettaPartialBlockIdentifier) IsSet() bool {
	return v.isSet
}

func (v *NullableRosettaPartialBlockIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRosettaPartialBlockIdentifier(val *RosettaPartialBlockIdentifier) *NullableRosettaPartialBlockIdentifier {
	return &NullableRosettaPartialBlockIdentifier{value: val, isSet: true}
}

func (v NullableRosettaPartialBlockIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRosettaPartialBlockIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
