package models

import (
	"encoding/json"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the BnsError type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &BnsError{}

// BnsError Error
type BnsError struct {
	Error *string `json:"error,omitempty"`
}

// NewBnsError instantiates a new BnsError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBnsError() *BnsError {
	this := BnsError{}
	return &this
}

// NewBnsErrorWithDefaults instantiates a new BnsError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBnsErrorWithDefaults() *BnsError {
	this := BnsError{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *BnsError) GetError() string {
	if o == nil || utils.IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BnsError) GetErrorOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *BnsError) HasError() bool {
	if o != nil && !utils.IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *BnsError) SetError(v string) {
	o.Error = &v
}

func (o BnsError) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BnsError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableBnsError struct {
	value *BnsError
	isSet bool
}

func (v NullableBnsError) Get() *BnsError {
	return v.value
}

func (v *NullableBnsError) Set(val *BnsError) {
	v.value = val
	v.isSet = true
}

func (v NullableBnsError) IsSet() bool {
	return v.isSet
}

func (v *NullableBnsError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBnsError(val *BnsError) *NullableBnsError {
	return &NullableBnsError{value: val, isSet: true}
}

func (v NullableBnsError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBnsError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
