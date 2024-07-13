package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the RosettaErrorNoDetails type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RosettaErrorNoDetails{}

// RosettaErrorNoDetails Instead of utilizing HTTP status codes to describe node errors (which often do not have a good analog), rich errors are returned using this object. Both the code and message fields can be individually used to correctly identify an error. Implementations MUST use unique values for both fields.
type RosettaErrorNoDetails struct {
	// Code is a network-specific error code. If desired, this code can be equivalent to an HTTP status code.
	Code int32 `json:"code"`
	// Message is a network-specific error message. The message MUST NOT change for a given code. In particular, this means that any contextual information should be included in the details field.
	Message string `json:"message"`
	// An error is retriable if the same request may succeed if submitted again.
	Retriable bool `json:"retriable"`
}

type _RosettaErrorNoDetails RosettaErrorNoDetails

// NewRosettaErrorNoDetails instantiates a new RosettaErrorNoDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRosettaErrorNoDetails(code int32, message string, retriable bool) *RosettaErrorNoDetails {
	this := RosettaErrorNoDetails{}
	this.Code = code
	this.Message = message
	this.Retriable = retriable
	return &this
}

// NewRosettaErrorNoDetailsWithDefaults instantiates a new RosettaErrorNoDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRosettaErrorNoDetailsWithDefaults() *RosettaErrorNoDetails {
	this := RosettaErrorNoDetails{}
	return &this
}

// GetCode returns the Code field value
func (o *RosettaErrorNoDetails) GetCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *RosettaErrorNoDetails) GetCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *RosettaErrorNoDetails) SetCode(v int32) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *RosettaErrorNoDetails) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *RosettaErrorNoDetails) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *RosettaErrorNoDetails) SetMessage(v string) {
	o.Message = v
}

// GetRetriable returns the Retriable field value
func (o *RosettaErrorNoDetails) GetRetriable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Retriable
}

// GetRetriableOk returns a tuple with the Retriable field value
// and a boolean to check if the value has been set.
func (o *RosettaErrorNoDetails) GetRetriableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Retriable, true
}

// SetRetriable sets field value
func (o *RosettaErrorNoDetails) SetRetriable(v bool) {
	o.Retriable = v
}

func (o RosettaErrorNoDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RosettaErrorNoDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	toSerialize["retriable"] = o.Retriable
	return toSerialize, nil
}

func (o *RosettaErrorNoDetails) UnmarshalJSON(data []byte) (err error) {
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

	varRosettaErrorNoDetails := _RosettaErrorNoDetails{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRosettaErrorNoDetails)

	if err != nil {
		return err
	}

	*o = RosettaErrorNoDetails(varRosettaErrorNoDetails)

	return err
}

type NullableRosettaErrorNoDetails struct {
	value *RosettaErrorNoDetails
	isSet bool
}

func (v NullableRosettaErrorNoDetails) Get() *RosettaErrorNoDetails {
	return v.value
}

func (v *NullableRosettaErrorNoDetails) Set(val *RosettaErrorNoDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableRosettaErrorNoDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableRosettaErrorNoDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRosettaErrorNoDetails(val *RosettaErrorNoDetails) *NullableRosettaErrorNoDetails {
	return &NullableRosettaErrorNoDetails{value: val, isSet: true}
}

func (v NullableRosettaErrorNoDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRosettaErrorNoDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
