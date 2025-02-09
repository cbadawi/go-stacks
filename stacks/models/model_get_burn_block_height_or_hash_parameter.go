package models

import (
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
	validator "gopkg.in/validator.v2"
)

// GetBurnBlockHeightOrHashParameter - struct for GetBurnBlockHeightOrHashParameter
type GetBurnBlockHeightOrHashParameter struct {
	Int32  *int32
	String *string
}

// int32AsGetBurnBlockHeightOrHashParameter is a convenience function that returns int32 wrapped in GetBurnBlockHeightOrHashParameter
func Int32AsGetBurnBlockHeightOrHashParameter(v *int32) GetBurnBlockHeightOrHashParameter {
	return GetBurnBlockHeightOrHashParameter{
		Int32: v,
	}
}

// stringAsGetBurnBlockHeightOrHashParameter is a convenience function that returns string wrapped in GetBurnBlockHeightOrHashParameter
func StringAsGetBurnBlockHeightOrHashParameter(v *string) GetBurnBlockHeightOrHashParameter {
	return GetBurnBlockHeightOrHashParameter{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetBurnBlockHeightOrHashParameter) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Int32
	err = utils.NewStrictDecoder(data).Decode(&dst.Int32)
	if err == nil {
		jsonInt32, _ := json.Marshal(dst.Int32)
		if string(jsonInt32) == "{}" { // empty struct
			dst.Int32 = nil
		} else {
			if err = validator.Validate(dst.Int32); err != nil {
				dst.Int32 = nil
			} else {
				match++
			}
		}
	} else {
		dst.Int32 = nil
	}

	// try to unmarshal data into String
	err = utils.NewStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			if err = validator.Validate(dst.String); err != nil {
				dst.String = nil
			} else {
				match++
			}
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Int32 = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetBurnBlockHeightOrHashParameter)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetBurnBlockHeightOrHashParameter)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetBurnBlockHeightOrHashParameter) MarshalJSON() ([]byte, error) {
	if src.Int32 != nil {
		return json.Marshal(&src.Int32)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetBurnBlockHeightOrHashParameter) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.Int32 != nil {
		return obj.Int32
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableGetBurnBlockHeightOrHashParameter struct {
	value *GetBurnBlockHeightOrHashParameter
	isSet bool
}

func (v NullableGetBurnBlockHeightOrHashParameter) Get() *GetBurnBlockHeightOrHashParameter {
	return v.value
}

func (v *NullableGetBurnBlockHeightOrHashParameter) Set(val *GetBurnBlockHeightOrHashParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBurnBlockHeightOrHashParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBurnBlockHeightOrHashParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBurnBlockHeightOrHashParameter(val *GetBurnBlockHeightOrHashParameter) *NullableGetBurnBlockHeightOrHashParameter {
	return &NullableGetBurnBlockHeightOrHashParameter{value: val, isSet: true}
}

func (v NullableGetBurnBlockHeightOrHashParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBurnBlockHeightOrHashParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
