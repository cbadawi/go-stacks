package models

import (
	"encoding/json"
	"fmt"
)

// PostCondition Post-conditionscan limit the damage done to a user's assets
type PostCondition struct {
	PostConditionFungible    *PostConditionFungible
	PostConditionNonFungible *PostConditionNonFungible
	PostConditionStx         *PostConditionStx
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *PostCondition) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into PostConditionFungible
	err = json.Unmarshal(data, &dst.PostConditionFungible)
	if err == nil {
		jsonPostConditionFungible, _ := json.Marshal(dst.PostConditionFungible)
		if string(jsonPostConditionFungible) == "{}" { // empty struct
			dst.PostConditionFungible = nil
		} else {
			return nil // data stored in dst.PostConditionFungible, return on the first match
		}
	} else {
		dst.PostConditionFungible = nil
	}

	// try to unmarshal JSON data into PostConditionNonFungible
	err = json.Unmarshal(data, &dst.PostConditionNonFungible)
	if err == nil {
		jsonPostConditionNonFungible, _ := json.Marshal(dst.PostConditionNonFungible)
		if string(jsonPostConditionNonFungible) == "{}" { // empty struct
			dst.PostConditionNonFungible = nil
		} else {
			return nil // data stored in dst.PostConditionNonFungible, return on the first match
		}
	} else {
		dst.PostConditionNonFungible = nil
	}

	// try to unmarshal JSON data into PostConditionStx
	err = json.Unmarshal(data, &dst.PostConditionStx)
	if err == nil {
		jsonPostConditionStx, _ := json.Marshal(dst.PostConditionStx)
		if string(jsonPostConditionStx) == "{}" { // empty struct
			dst.PostConditionStx = nil
		} else {
			return nil // data stored in dst.PostConditionStx, return on the first match
		}
	} else {
		dst.PostConditionStx = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(PostCondition)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *PostCondition) MarshalJSON() ([]byte, error) {
	if src.PostConditionFungible != nil {
		return json.Marshal(&src.PostConditionFungible)
	}

	if src.PostConditionNonFungible != nil {
		return json.Marshal(&src.PostConditionNonFungible)
	}

	if src.PostConditionStx != nil {
		return json.Marshal(&src.PostConditionStx)
	}

	return nil, nil // no data in anyOf schemas
}

type NullablePostCondition struct {
	value *PostCondition
	isSet bool
}

func (v NullablePostCondition) Get() *PostCondition {
	return v.value
}

func (v *NullablePostCondition) Set(val *PostCondition) {
	v.value = val
	v.isSet = true
}

func (v NullablePostCondition) IsSet() bool {
	return v.isSet
}

func (v *NullablePostCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostCondition(val *PostCondition) *NullablePostCondition {
	return &NullablePostCondition{value: val, isSet: true}
}

func (v NullablePostCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
