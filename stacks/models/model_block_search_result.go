package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the BlockSearchResult type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &BlockSearchResult{}

// BlockSearchResult Block search result
type BlockSearchResult struct {
	// Indicates if the requested object was found or not
	Found  bool                    `json:"found"`
	Result BlockSearchResultResult `json:"result"`
}

type _BlockSearchResult BlockSearchResult

// NewBlockSearchResult instantiates a new BlockSearchResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockSearchResult(found bool, result BlockSearchResultResult) *BlockSearchResult {
	this := BlockSearchResult{}
	this.Found = found
	this.Result = result
	return &this
}

// NewBlockSearchResultWithDefaults instantiates a new BlockSearchResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockSearchResultWithDefaults() *BlockSearchResult {
	this := BlockSearchResult{}
	var found bool = true
	this.Found = found
	return &this
}

// GetFound returns the Found field value
func (o *BlockSearchResult) GetFound() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Found
}

// GetFoundOk returns a tuple with the Found field value
// and a boolean to check if the value has been set.
func (o *BlockSearchResult) GetFoundOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Found, true
}

// SetFound sets field value
func (o *BlockSearchResult) SetFound(v bool) {
	o.Found = v
}

// GetResult returns the Result field value
func (o *BlockSearchResult) GetResult() BlockSearchResultResult {
	if o == nil {
		var ret BlockSearchResultResult
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *BlockSearchResult) GetResultOk() (*BlockSearchResultResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *BlockSearchResult) SetResult(v BlockSearchResultResult) {
	o.Result = v
}

func (o BlockSearchResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlockSearchResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["found"] = o.Found
	toSerialize["result"] = o.Result
	return toSerialize, nil
}

func (o *BlockSearchResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"found",
		"result",
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

	varBlockSearchResult := _BlockSearchResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBlockSearchResult)

	if err != nil {
		return err
	}

	*o = BlockSearchResult(varBlockSearchResult)

	return err
}

type NullableBlockSearchResult struct {
	value *BlockSearchResult
	isSet bool
}

func (v NullableBlockSearchResult) Get() *BlockSearchResult {
	return v.value
}

func (v *NullableBlockSearchResult) Set(val *BlockSearchResult) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockSearchResult) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockSearchResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockSearchResult(val *BlockSearchResult) *NullableBlockSearchResult {
	return &NullableBlockSearchResult{value: val, isSet: true}
}

func (v NullableBlockSearchResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockSearchResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
