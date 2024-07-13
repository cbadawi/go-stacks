package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the ReadOnlyFunctionSuccessResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ReadOnlyFunctionSuccessResponse{}

// ReadOnlyFunctionSuccessResponse GET request to get contract source
type ReadOnlyFunctionSuccessResponse struct {
	Okay   bool    `json:"okay"`
	Result *string `json:"result,omitempty"`
	Cause  *string `json:"cause,omitempty"`
}

type _ReadOnlyFunctionSuccessResponse ReadOnlyFunctionSuccessResponse

// NewReadOnlyFunctionSuccessResponse instantiates a new ReadOnlyFunctionSuccessResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadOnlyFunctionSuccessResponse(okay bool) *ReadOnlyFunctionSuccessResponse {
	this := ReadOnlyFunctionSuccessResponse{}
	this.Okay = okay
	return &this
}

// NewReadOnlyFunctionSuccessResponseWithDefaults instantiates a new ReadOnlyFunctionSuccessResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadOnlyFunctionSuccessResponseWithDefaults() *ReadOnlyFunctionSuccessResponse {
	this := ReadOnlyFunctionSuccessResponse{}
	return &this
}

// GetOkay returns the Okay field value
func (o *ReadOnlyFunctionSuccessResponse) GetOkay() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Okay
}

// GetOkayOk returns a tuple with the Okay field value
// and a boolean to check if the value has been set.
func (o *ReadOnlyFunctionSuccessResponse) GetOkayOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Okay, true
}

// SetOkay sets field value
func (o *ReadOnlyFunctionSuccessResponse) SetOkay(v bool) {
	o.Okay = v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *ReadOnlyFunctionSuccessResponse) GetResult() string {
	if o == nil || utils.IsNil(o.Result) {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadOnlyFunctionSuccessResponse) GetResultOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *ReadOnlyFunctionSuccessResponse) HasResult() bool {
	if o != nil && !utils.IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *ReadOnlyFunctionSuccessResponse) SetResult(v string) {
	o.Result = &v
}

// GetCause returns the Cause field value if set, zero value otherwise.
func (o *ReadOnlyFunctionSuccessResponse) GetCause() string {
	if o == nil || utils.IsNil(o.Cause) {
		var ret string
		return ret
	}
	return *o.Cause
}

// GetCauseOk returns a tuple with the Cause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadOnlyFunctionSuccessResponse) GetCauseOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Cause) {
		return nil, false
	}
	return o.Cause, true
}

// HasCause returns a boolean if a field has been set.
func (o *ReadOnlyFunctionSuccessResponse) HasCause() bool {
	if o != nil && !utils.IsNil(o.Cause) {
		return true
	}

	return false
}

// SetCause gets a reference to the given string and assigns it to the Cause field.
func (o *ReadOnlyFunctionSuccessResponse) SetCause(v string) {
	o.Cause = &v
}

func (o ReadOnlyFunctionSuccessResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReadOnlyFunctionSuccessResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["okay"] = o.Okay
	if !utils.IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	if !utils.IsNil(o.Cause) {
		toSerialize["cause"] = o.Cause
	}
	return toSerialize, nil
}

func (o *ReadOnlyFunctionSuccessResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"okay",
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

	varReadOnlyFunctionSuccessResponse := _ReadOnlyFunctionSuccessResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varReadOnlyFunctionSuccessResponse)

	if err != nil {
		return err
	}

	*o = ReadOnlyFunctionSuccessResponse(varReadOnlyFunctionSuccessResponse)

	return err
}

type NullableReadOnlyFunctionSuccessResponse struct {
	value *ReadOnlyFunctionSuccessResponse
	isSet bool
}

func (v NullableReadOnlyFunctionSuccessResponse) Get() *ReadOnlyFunctionSuccessResponse {
	return v.value
}

func (v *NullableReadOnlyFunctionSuccessResponse) Set(val *ReadOnlyFunctionSuccessResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReadOnlyFunctionSuccessResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReadOnlyFunctionSuccessResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadOnlyFunctionSuccessResponse(val *ReadOnlyFunctionSuccessResponse) *NullableReadOnlyFunctionSuccessResponse {
	return &NullableReadOnlyFunctionSuccessResponse{value: val, isSet: true}
}

func (v NullableReadOnlyFunctionSuccessResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadOnlyFunctionSuccessResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
