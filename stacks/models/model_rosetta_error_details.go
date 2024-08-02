package models

import (
	"encoding/json"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the RosettaErrorDetails type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RosettaErrorDetails{}

// RosettaErrorDetails Often times it is useful to return context specific to the request that caused the error (i.e. a sample of the stack trace or impacted account) in addition to the standard error message.
type RosettaErrorDetails struct {
	Address *string `json:"address,omitempty"`
	Error   *string `json:"error,omitempty"`
}

// NewRosettaErrorDetails instantiates a new RosettaErrorDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRosettaErrorDetails() *RosettaErrorDetails {
	this := RosettaErrorDetails{}
	return &this
}

// NewRosettaErrorDetailsWithDefaults instantiates a new RosettaErrorDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRosettaErrorDetailsWithDefaults() *RosettaErrorDetails {
	this := RosettaErrorDetails{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *RosettaErrorDetails) GetAddress() string {
	if o == nil || utils.IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaErrorDetails) GetAddressOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *RosettaErrorDetails) HasAddress() bool {
	if o != nil && !utils.IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *RosettaErrorDetails) SetAddress(v string) {
	o.Address = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *RosettaErrorDetails) GetError() string {
	if o == nil || utils.IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaErrorDetails) GetErrorOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *RosettaErrorDetails) HasError() bool {
	if o != nil && !utils.IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *RosettaErrorDetails) SetError(v string) {
	o.Error = &v
}

func (o RosettaErrorDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RosettaErrorDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !utils.IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableRosettaErrorDetails struct {
	value *RosettaErrorDetails
	isSet bool
}

func (v NullableRosettaErrorDetails) Get() *RosettaErrorDetails {
	return v.value
}

func (v *NullableRosettaErrorDetails) Set(val *RosettaErrorDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableRosettaErrorDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableRosettaErrorDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRosettaErrorDetails(val *RosettaErrorDetails) *NullableRosettaErrorDetails {
	return &NullableRosettaErrorDetails{value: val, isSet: true}
}

func (v NullableRosettaErrorDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRosettaErrorDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
