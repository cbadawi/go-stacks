package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the AddressSearchResult type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AddressSearchResult{}

// AddressSearchResult Address search result
type AddressSearchResult struct {
	// Indicates if the requested object was found or not
	Found  bool                      `json:"found"`
	Result AddressSearchResultResult `json:"result"`
}

type _AddressSearchResult AddressSearchResult

// NewAddressSearchResult instantiates a new AddressSearchResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressSearchResult(found bool, result AddressSearchResultResult) *AddressSearchResult {
	this := AddressSearchResult{}
	this.Found = found
	this.Result = result
	return &this
}

// NewAddressSearchResultWithDefaults instantiates a new AddressSearchResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressSearchResultWithDefaults() *AddressSearchResult {
	this := AddressSearchResult{}
	var found bool = true
	this.Found = found
	return &this
}

// GetFound returns the Found field value
func (o *AddressSearchResult) GetFound() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Found
}

// GetFoundOk returns a tuple with the Found field value
// and a boolean to check if the value has been set.
func (o *AddressSearchResult) GetFoundOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Found, true
}

// SetFound sets field value
func (o *AddressSearchResult) SetFound(v bool) {
	o.Found = v
}

// GetResult returns the Result field value
func (o *AddressSearchResult) GetResult() AddressSearchResultResult {
	if o == nil {
		var ret AddressSearchResultResult
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *AddressSearchResult) GetResultOk() (*AddressSearchResultResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *AddressSearchResult) SetResult(v AddressSearchResultResult) {
	o.Result = v
}

func (o AddressSearchResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressSearchResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["found"] = o.Found
	toSerialize["result"] = o.Result
	return toSerialize, nil
}

func (o *AddressSearchResult) UnmarshalJSON(data []byte) (err error) {
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

	varAddressSearchResult := _AddressSearchResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddressSearchResult)

	if err != nil {
		return err
	}

	*o = AddressSearchResult(varAddressSearchResult)

	return err
}

type NullableAddressSearchResult struct {
	value *AddressSearchResult
	isSet bool
}

func (v NullableAddressSearchResult) Get() *AddressSearchResult {
	return v.value
}

func (v *NullableAddressSearchResult) Set(val *AddressSearchResult) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressSearchResult) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressSearchResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressSearchResult(val *AddressSearchResult) *NullableAddressSearchResult {
	return &NullableAddressSearchResult{value: val, isSet: true}
}

func (v NullableAddressSearchResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressSearchResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
