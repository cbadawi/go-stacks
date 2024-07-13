package models

import (
	"encoding/json"
	"fmt"
)

// PostConditionPrincipal struct for PostConditionPrincipal
type PostConditionPrincipal struct {
	PostConditionPrincipalAnyOf  *PostConditionPrincipalAnyOf
	PostConditionPrincipalAnyOf1 *PostConditionPrincipalAnyOf1
	PostConditionPrincipalAnyOf2 *PostConditionPrincipalAnyOf2
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *PostConditionPrincipal) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into PostConditionPrincipalAnyOf
	err = json.Unmarshal(data, &dst.PostConditionPrincipalAnyOf)
	if err == nil {
		jsonPostConditionPrincipalAnyOf, _ := json.Marshal(dst.PostConditionPrincipalAnyOf)
		if string(jsonPostConditionPrincipalAnyOf) == "{}" { // empty struct
			dst.PostConditionPrincipalAnyOf = nil
		} else {
			return nil // data stored in dst.PostConditionPrincipalAnyOf, return on the first match
		}
	} else {
		dst.PostConditionPrincipalAnyOf = nil
	}

	// try to unmarshal JSON data into PostConditionPrincipalAnyOf1
	err = json.Unmarshal(data, &dst.PostConditionPrincipalAnyOf1)
	if err == nil {
		jsonPostConditionPrincipalAnyOf1, _ := json.Marshal(dst.PostConditionPrincipalAnyOf1)
		if string(jsonPostConditionPrincipalAnyOf1) == "{}" { // empty struct
			dst.PostConditionPrincipalAnyOf1 = nil
		} else {
			return nil // data stored in dst.PostConditionPrincipalAnyOf1, return on the first match
		}
	} else {
		dst.PostConditionPrincipalAnyOf1 = nil
	}

	// try to unmarshal JSON data into PostConditionPrincipalAnyOf2
	err = json.Unmarshal(data, &dst.PostConditionPrincipalAnyOf2)
	if err == nil {
		jsonPostConditionPrincipalAnyOf2, _ := json.Marshal(dst.PostConditionPrincipalAnyOf2)
		if string(jsonPostConditionPrincipalAnyOf2) == "{}" { // empty struct
			dst.PostConditionPrincipalAnyOf2 = nil
		} else {
			return nil // data stored in dst.PostConditionPrincipalAnyOf2, return on the first match
		}
	} else {
		dst.PostConditionPrincipalAnyOf2 = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(PostConditionPrincipal)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *PostConditionPrincipal) MarshalJSON() ([]byte, error) {
	if src.PostConditionPrincipalAnyOf != nil {
		return json.Marshal(&src.PostConditionPrincipalAnyOf)
	}

	if src.PostConditionPrincipalAnyOf1 != nil {
		return json.Marshal(&src.PostConditionPrincipalAnyOf1)
	}

	if src.PostConditionPrincipalAnyOf2 != nil {
		return json.Marshal(&src.PostConditionPrincipalAnyOf2)
	}

	return nil, nil // no data in anyOf schemas
}

type NullablePostConditionPrincipal struct {
	value *PostConditionPrincipal
	isSet bool
}

func (v NullablePostConditionPrincipal) Get() *PostConditionPrincipal {
	return v.value
}

func (v *NullablePostConditionPrincipal) Set(val *PostConditionPrincipal) {
	v.value = val
	v.isSet = true
}

func (v NullablePostConditionPrincipal) IsSet() bool {
	return v.isSet
}

func (v *NullablePostConditionPrincipal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostConditionPrincipal(val *PostConditionPrincipal) *NullablePostConditionPrincipal {
	return &NullablePostConditionPrincipal{value: val, isSet: true}
}

func (v NullablePostConditionPrincipal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostConditionPrincipal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
