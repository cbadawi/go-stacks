package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the RosettaError type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RosettaError{}

// RosettaError Instead of utilizing HTTP status codes to describe node errors (which often do not have a good analog), rich errors are returned using this object. Both the code and message fields can be individually used to correctly identify an error. Implementations MUST use unique values for both fields.
type RosettaError struct {
	// Code is a network-specific error code. If desired, this code can be equivalent to an HTTP status code.
	Code int32 `json:"code"`
	// Message is a network-specific error message. The message MUST NOT change for a given code. In particular, this means that any contextual information should be included in the details field.
	Message string `json:"message"`
	// An error is retriable if the same request may succeed if submitted again.
	Retriable bool                 `json:"retriable"`
	Details   *RosettaErrorDetails `json:"details,omitempty"`
}

type _RosettaError RosettaError

// NewRosettaError instantiates a new RosettaError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRosettaError(code int32, message string, retriable bool) *RosettaError {
	this := RosettaError{}
	this.Code = code
	this.Message = message
	this.Retriable = retriable
	return &this
}

// NewRosettaErrorWithDefaults instantiates a new RosettaError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRosettaErrorWithDefaults() *RosettaError {
	this := RosettaError{}
	return &this
}

// GetCode returns the Code field value
func (o *RosettaError) GetCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *RosettaError) GetCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *RosettaError) SetCode(v int32) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *RosettaError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *RosettaError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *RosettaError) SetMessage(v string) {
	o.Message = v
}

// GetRetriable returns the Retriable field value
func (o *RosettaError) GetRetriable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Retriable
}

// GetRetriableOk returns a tuple with the Retriable field value
// and a boolean to check if the value has been set.
func (o *RosettaError) GetRetriableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Retriable, true
}

// SetRetriable sets field value
func (o *RosettaError) SetRetriable(v bool) {
	o.Retriable = v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *RosettaError) GetDetails() RosettaErrorDetails {
	if o == nil || utils.IsNil(o.Details) {
		var ret RosettaErrorDetails
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaError) GetDetailsOk() (*RosettaErrorDetails, bool) {
	if o == nil || utils.IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *RosettaError) HasDetails() bool {
	if o != nil && !utils.IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given RosettaErrorDetails and assigns it to the Details field.
func (o *RosettaError) SetDetails(v RosettaErrorDetails) {
	o.Details = &v
}

func (o RosettaError) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RosettaError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	toSerialize["retriable"] = o.Retriable
	if !utils.IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	return toSerialize, nil
}

func (o *RosettaError) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"code",
		"message",
		"retriable",
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

	varRosettaError := _RosettaError{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRosettaError)

	if err != nil {
		return err
	}

	*o = RosettaError(varRosettaError)

	return err
}

type NullableRosettaError struct {
	value *RosettaError
	isSet bool
}

func (v NullableRosettaError) Get() *RosettaError {
	return v.value
}

func (v *NullableRosettaError) Set(val *RosettaError) {
	v.value = val
	v.isSet = true
}

func (v NullableRosettaError) IsSet() bool {
	return v.isSet
}

func (v *NullableRosettaError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRosettaError(val *RosettaError) *NullableRosettaError {
	return &NullableRosettaError{value: val, isSet: true}
}

func (v NullableRosettaError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRosettaError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
