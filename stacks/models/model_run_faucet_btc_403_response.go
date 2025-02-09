package models

import (
	"encoding/json"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the RunFaucetBtc403Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RunFaucetBtc403Response{}

// RunFaucetBtc403Response struct for RunFaucetBtc403Response
type RunFaucetBtc403Response struct {
	Error   *string `json:"error,omitempty"`
	Success *bool   `json:"success,omitempty"`
}

// NewRunFaucetBtc403Response instantiates a new RunFaucetBtc403Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunFaucetBtc403Response() *RunFaucetBtc403Response {
	this := RunFaucetBtc403Response{}
	return &this
}

// NewRunFaucetBtc403ResponseWithDefaults instantiates a new RunFaucetBtc403Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunFaucetBtc403ResponseWithDefaults() *RunFaucetBtc403Response {
	this := RunFaucetBtc403Response{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *RunFaucetBtc403Response) GetError() string {
	if o == nil || utils.IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunFaucetBtc403Response) GetErrorOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *RunFaucetBtc403Response) HasError() bool {
	if o != nil && !utils.IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *RunFaucetBtc403Response) SetError(v string) {
	o.Error = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *RunFaucetBtc403Response) GetSuccess() bool {
	if o == nil || utils.IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunFaucetBtc403Response) GetSuccessOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *RunFaucetBtc403Response) HasSuccess() bool {
	if o != nil && !utils.IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *RunFaucetBtc403Response) SetSuccess(v bool) {
	o.Success = &v
}

func (o RunFaucetBtc403Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RunFaucetBtc403Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !utils.IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	return toSerialize, nil
}

type NullableRunFaucetBtc403Response struct {
	value *RunFaucetBtc403Response
	isSet bool
}

func (v NullableRunFaucetBtc403Response) Get() *RunFaucetBtc403Response {
	return v.value
}

func (v *NullableRunFaucetBtc403Response) Set(val *RunFaucetBtc403Response) {
	v.value = val
	v.isSet = true
}

func (v NullableRunFaucetBtc403Response) IsSet() bool {
	return v.isSet
}

func (v *NullableRunFaucetBtc403Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunFaucetBtc403Response(val *RunFaucetBtc403Response) *NullableRunFaucetBtc403Response {
	return &NullableRunFaucetBtc403Response{value: val, isSet: true}
}

func (v NullableRunFaucetBtc403Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunFaucetBtc403Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
