package models

import (
	"encoding/json"
	"fmt"
)

// SearchResult complete search result for terms
type SearchResult struct {
	SearchErrorResult   *SearchErrorResult
	SearchSuccessResult *SearchSuccessResult
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *SearchResult) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into SearchErrorResult
	err = json.Unmarshal(data, &dst.SearchErrorResult)
	if err == nil {
		jsonSearchErrorResult, _ := json.Marshal(dst.SearchErrorResult)
		if string(jsonSearchErrorResult) == "{}" { // empty struct
			dst.SearchErrorResult = nil
		} else {
			return nil // data stored in dst.SearchErrorResult, return on the first match
		}
	} else {
		dst.SearchErrorResult = nil
	}

	// try to unmarshal JSON data into SearchSuccessResult
	err = json.Unmarshal(data, &dst.SearchSuccessResult)
	if err == nil {
		jsonSearchSuccessResult, _ := json.Marshal(dst.SearchSuccessResult)
		if string(jsonSearchSuccessResult) == "{}" { // empty struct
			dst.SearchSuccessResult = nil
		} else {
			return nil // data stored in dst.SearchSuccessResult, return on the first match
		}
	} else {
		dst.SearchSuccessResult = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(SearchResult)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *SearchResult) MarshalJSON() ([]byte, error) {
	if src.SearchErrorResult != nil {
		return json.Marshal(&src.SearchErrorResult)
	}

	if src.SearchSuccessResult != nil {
		return json.Marshal(&src.SearchSuccessResult)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableSearchResult struct {
	value *SearchResult
	isSet bool
}

func (v NullableSearchResult) Get() *SearchResult {
	return v.value
}

func (v *NullableSearchResult) Set(val *SearchResult) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchResult) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchResult(val *SearchResult) *NullableSearchResult {
	return &NullableSearchResult{value: val, isSet: true}
}

func (v NullableSearchResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
