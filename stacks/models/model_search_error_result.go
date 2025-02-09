package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the SearchErrorResult type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &SearchErrorResult{}

// SearchErrorResult Error search result
type SearchErrorResult struct {
	// Indicates if the requested object was found or not
	Found  bool                    `json:"found"`
	Result SearchErrorResultResult `json:"result"`
	Error  string                  `json:"error"`
}

type _SearchErrorResult SearchErrorResult

// NewSearchErrorResult instantiates a new SearchErrorResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchErrorResult(found bool, result SearchErrorResultResult, error_ string) *SearchErrorResult {
	this := SearchErrorResult{}
	this.Found = found
	this.Result = result
	this.Error = error_
	return &this
}

// NewSearchErrorResultWithDefaults instantiates a new SearchErrorResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchErrorResultWithDefaults() *SearchErrorResult {
	this := SearchErrorResult{}
	var found bool = false
	this.Found = found
	return &this
}

// GetFound returns the Found field value
func (o *SearchErrorResult) GetFound() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Found
}

// GetFoundOk returns a tuple with the Found field value
// and a boolean to check if the value has been set.
func (o *SearchErrorResult) GetFoundOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Found, true
}

// SetFound sets field value
func (o *SearchErrorResult) SetFound(v bool) {
	o.Found = v
}

// GetResult returns the Result field value
func (o *SearchErrorResult) GetResult() SearchErrorResultResult {
	if o == nil {
		var ret SearchErrorResultResult
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *SearchErrorResult) GetResultOk() (*SearchErrorResultResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *SearchErrorResult) SetResult(v SearchErrorResultResult) {
	o.Result = v
}

// GetError returns the Error field value
func (o *SearchErrorResult) GetError() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *SearchErrorResult) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *SearchErrorResult) SetError(v string) {
	o.Error = v
}

func (o SearchErrorResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchErrorResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["found"] = o.Found
	toSerialize["result"] = o.Result
	toSerialize["error"] = o.Error
	return toSerialize, nil
}

func (o *SearchErrorResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"found",
		"result",
		"error",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSearchErrorResult := _SearchErrorResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSearchErrorResult)

	if err != nil {
		return err
	}

	*o = SearchErrorResult(varSearchErrorResult)

	return err
}

type NullableSearchErrorResult struct {
	value *SearchErrorResult
	isSet bool
}

func (v NullableSearchErrorResult) Get() *SearchErrorResult {
	return v.value
}

func (v *NullableSearchErrorResult) Set(val *SearchErrorResult) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchErrorResult) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchErrorResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchErrorResult(val *SearchErrorResult) *NullableSearchErrorResult {
	return &NullableSearchErrorResult{value: val, isSet: true}
}

func (v NullableSearchErrorResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchErrorResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
