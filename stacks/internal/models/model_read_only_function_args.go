package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the ReadOnlyFunctionArgs type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ReadOnlyFunctionArgs{}

// ReadOnlyFunctionArgs Describes representation of a Type-0 Stacks 2.0 transaction. https://github.com/blockstack/stacks-blockchain/blob/master/sip/sip-005-blocks-and-transactions.md#type-0-transferring-an-asset
type ReadOnlyFunctionArgs struct {
	// The simulated tx-sender
	Sender string `json:"sender"`
	// An array of hex serialized Clarity values
	Arguments []string `json:"arguments"`
}

type _ReadOnlyFunctionArgs ReadOnlyFunctionArgs

// NewReadOnlyFunctionArgs instantiates a new ReadOnlyFunctionArgs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadOnlyFunctionArgs(sender string, arguments []string) *ReadOnlyFunctionArgs {
	this := ReadOnlyFunctionArgs{}
	this.Sender = sender
	this.Arguments = arguments
	return &this
}

// NewReadOnlyFunctionArgsWithDefaults instantiates a new ReadOnlyFunctionArgs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadOnlyFunctionArgsWithDefaults() *ReadOnlyFunctionArgs {
	this := ReadOnlyFunctionArgs{}
	return &this
}

// GetSender returns the Sender field value
func (o *ReadOnlyFunctionArgs) GetSender() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sender
}

// GetSenderOk returns a tuple with the Sender field value
// and a boolean to check if the value has been set.
func (o *ReadOnlyFunctionArgs) GetSenderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sender, true
}

// SetSender sets field value
func (o *ReadOnlyFunctionArgs) SetSender(v string) {
	o.Sender = v
}

// GetArguments returns the Arguments field value
func (o *ReadOnlyFunctionArgs) GetArguments() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Arguments
}

// GetArgumentsOk returns a tuple with the Arguments field value
// and a boolean to check if the value has been set.
func (o *ReadOnlyFunctionArgs) GetArgumentsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Arguments, true
}

// SetArguments sets field value
func (o *ReadOnlyFunctionArgs) SetArguments(v []string) {
	o.Arguments = v
}

func (o ReadOnlyFunctionArgs) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReadOnlyFunctionArgs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sender"] = o.Sender
	toSerialize["arguments"] = o.Arguments
	return toSerialize, nil
}

func (o *ReadOnlyFunctionArgs) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sender",
		"arguments",
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

	varReadOnlyFunctionArgs := _ReadOnlyFunctionArgs{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varReadOnlyFunctionArgs)

	if err != nil {
		return err
	}

	*o = ReadOnlyFunctionArgs(varReadOnlyFunctionArgs)

	return err
}

type NullableReadOnlyFunctionArgs struct {
	value *ReadOnlyFunctionArgs
	isSet bool
}

func (v NullableReadOnlyFunctionArgs) Get() *ReadOnlyFunctionArgs {
	return v.value
}

func (v *NullableReadOnlyFunctionArgs) Set(val *ReadOnlyFunctionArgs) {
	v.value = val
	v.isSet = true
}

func (v NullableReadOnlyFunctionArgs) IsSet() bool {
	return v.isSet
}

func (v *NullableReadOnlyFunctionArgs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadOnlyFunctionArgs(val *ReadOnlyFunctionArgs) *NullableReadOnlyFunctionArgs {
	return &NullableReadOnlyFunctionArgs{value: val, isSet: true}
}

func (v NullableReadOnlyFunctionArgs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadOnlyFunctionArgs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
