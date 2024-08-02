package models

import (
	"encoding/json"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the BnsFetchHistoricalZoneFileResponseAnyOf1 type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &BnsFetchHistoricalZoneFileResponseAnyOf1{}

// BnsFetchHistoricalZoneFileResponseAnyOf1 struct for BnsFetchHistoricalZoneFileResponseAnyOf1
type BnsFetchHistoricalZoneFileResponseAnyOf1 struct {
	Error *string `json:"error,omitempty"`
}

// NewBnsFetchHistoricalZoneFileResponseAnyOf1 instantiates a new BnsFetchHistoricalZoneFileResponseAnyOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBnsFetchHistoricalZoneFileResponseAnyOf1() *BnsFetchHistoricalZoneFileResponseAnyOf1 {
	this := BnsFetchHistoricalZoneFileResponseAnyOf1{}
	return &this
}

// NewBnsFetchHistoricalZoneFileResponseAnyOf1WithDefaults instantiates a new BnsFetchHistoricalZoneFileResponseAnyOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBnsFetchHistoricalZoneFileResponseAnyOf1WithDefaults() *BnsFetchHistoricalZoneFileResponseAnyOf1 {
	this := BnsFetchHistoricalZoneFileResponseAnyOf1{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *BnsFetchHistoricalZoneFileResponseAnyOf1) GetError() string {
	if o == nil || utils.IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BnsFetchHistoricalZoneFileResponseAnyOf1) GetErrorOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *BnsFetchHistoricalZoneFileResponseAnyOf1) HasError() bool {
	if o != nil && !utils.IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *BnsFetchHistoricalZoneFileResponseAnyOf1) SetError(v string) {
	o.Error = &v
}

func (o BnsFetchHistoricalZoneFileResponseAnyOf1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BnsFetchHistoricalZoneFileResponseAnyOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableBnsFetchHistoricalZoneFileResponseAnyOf1 struct {
	value *BnsFetchHistoricalZoneFileResponseAnyOf1
	isSet bool
}

func (v NullableBnsFetchHistoricalZoneFileResponseAnyOf1) Get() *BnsFetchHistoricalZoneFileResponseAnyOf1 {
	return v.value
}

func (v *NullableBnsFetchHistoricalZoneFileResponseAnyOf1) Set(val *BnsFetchHistoricalZoneFileResponseAnyOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableBnsFetchHistoricalZoneFileResponseAnyOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableBnsFetchHistoricalZoneFileResponseAnyOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBnsFetchHistoricalZoneFileResponseAnyOf1(val *BnsFetchHistoricalZoneFileResponseAnyOf1) *NullableBnsFetchHistoricalZoneFileResponseAnyOf1 {
	return &NullableBnsFetchHistoricalZoneFileResponseAnyOf1{value: val, isSet: true}
}

func (v NullableBnsFetchHistoricalZoneFileResponseAnyOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBnsFetchHistoricalZoneFileResponseAnyOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
