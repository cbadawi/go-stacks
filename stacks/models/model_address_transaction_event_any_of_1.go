package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the AddressTransactionEventAnyOf1 type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AddressTransactionEventAnyOf1{}

// AddressTransactionEventAnyOf1 struct for AddressTransactionEventAnyOf1
type AddressTransactionEventAnyOf1 struct {
	Type       string                            `json:"type"`
	EventIndex int32                             `json:"event_index"`
	Data       AddressTransactionEventAnyOf1Data `json:"data"`
}

type _AddressTransactionEventAnyOf1 AddressTransactionEventAnyOf1

// NewAddressTransactionEventAnyOf1 instantiates a new AddressTransactionEventAnyOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressTransactionEventAnyOf1(type_ string, eventIndex int32, data AddressTransactionEventAnyOf1Data) *AddressTransactionEventAnyOf1 {
	this := AddressTransactionEventAnyOf1{}
	this.Type = type_
	this.EventIndex = eventIndex
	this.Data = data
	return &this
}

// NewAddressTransactionEventAnyOf1WithDefaults instantiates a new AddressTransactionEventAnyOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressTransactionEventAnyOf1WithDefaults() *AddressTransactionEventAnyOf1 {
	this := AddressTransactionEventAnyOf1{}
	return &this
}

// GetType returns the Type field value
func (o *AddressTransactionEventAnyOf1) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AddressTransactionEventAnyOf1) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AddressTransactionEventAnyOf1) SetType(v string) {
	o.Type = v
}

// GetEventIndex returns the EventIndex field value
func (o *AddressTransactionEventAnyOf1) GetEventIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EventIndex
}

// GetEventIndexOk returns a tuple with the EventIndex field value
// and a boolean to check if the value has been set.
func (o *AddressTransactionEventAnyOf1) GetEventIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventIndex, true
}

// SetEventIndex sets field value
func (o *AddressTransactionEventAnyOf1) SetEventIndex(v int32) {
	o.EventIndex = v
}

// GetData returns the Data field value
func (o *AddressTransactionEventAnyOf1) GetData() AddressTransactionEventAnyOf1Data {
	if o == nil {
		var ret AddressTransactionEventAnyOf1Data
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AddressTransactionEventAnyOf1) GetDataOk() (*AddressTransactionEventAnyOf1Data, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AddressTransactionEventAnyOf1) SetData(v AddressTransactionEventAnyOf1Data) {
	o.Data = v
}

func (o AddressTransactionEventAnyOf1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressTransactionEventAnyOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["event_index"] = o.EventIndex
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *AddressTransactionEventAnyOf1) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"event_index",
		"data",
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

	varAddressTransactionEventAnyOf1 := _AddressTransactionEventAnyOf1{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddressTransactionEventAnyOf1)

	if err != nil {
		return err
	}

	*o = AddressTransactionEventAnyOf1(varAddressTransactionEventAnyOf1)

	return err
}

type NullableAddressTransactionEventAnyOf1 struct {
	value *AddressTransactionEventAnyOf1
	isSet bool
}

func (v NullableAddressTransactionEventAnyOf1) Get() *AddressTransactionEventAnyOf1 {
	return v.value
}

func (v *NullableAddressTransactionEventAnyOf1) Set(val *AddressTransactionEventAnyOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressTransactionEventAnyOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressTransactionEventAnyOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressTransactionEventAnyOf1(val *AddressTransactionEventAnyOf1) *NullableAddressTransactionEventAnyOf1 {
	return &NullableAddressTransactionEventAnyOf1{value: val, isSet: true}
}

func (v NullableAddressTransactionEventAnyOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressTransactionEventAnyOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
