package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the AddressTransactionEventAnyOf type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AddressTransactionEventAnyOf{}

// AddressTransactionEventAnyOf struct for AddressTransactionEventAnyOf
type AddressTransactionEventAnyOf struct {
	Type       string                           `json:"type"`
	EventIndex int32                            `json:"event_index"`
	Data       AddressTransactionEventAnyOfData `json:"data"`
}

type _AddressTransactionEventAnyOf AddressTransactionEventAnyOf

// NewAddressTransactionEventAnyOf instantiates a new AddressTransactionEventAnyOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressTransactionEventAnyOf(type_ string, eventIndex int32, data AddressTransactionEventAnyOfData) *AddressTransactionEventAnyOf {
	this := AddressTransactionEventAnyOf{}
	this.Type = type_
	this.EventIndex = eventIndex
	this.Data = data
	return &this
}

// NewAddressTransactionEventAnyOfWithDefaults instantiates a new AddressTransactionEventAnyOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressTransactionEventAnyOfWithDefaults() *AddressTransactionEventAnyOf {
	this := AddressTransactionEventAnyOf{}
	return &this
}

// GetType returns the Type field value
func (o *AddressTransactionEventAnyOf) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AddressTransactionEventAnyOf) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AddressTransactionEventAnyOf) SetType(v string) {
	o.Type = v
}

// GetEventIndex returns the EventIndex field value
func (o *AddressTransactionEventAnyOf) GetEventIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EventIndex
}

// GetEventIndexOk returns a tuple with the EventIndex field value
// and a boolean to check if the value has been set.
func (o *AddressTransactionEventAnyOf) GetEventIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventIndex, true
}

// SetEventIndex sets field value
func (o *AddressTransactionEventAnyOf) SetEventIndex(v int32) {
	o.EventIndex = v
}

// GetData returns the Data field value
func (o *AddressTransactionEventAnyOf) GetData() AddressTransactionEventAnyOfData {
	if o == nil {
		var ret AddressTransactionEventAnyOfData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AddressTransactionEventAnyOf) GetDataOk() (*AddressTransactionEventAnyOfData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AddressTransactionEventAnyOf) SetData(v AddressTransactionEventAnyOfData) {
	o.Data = v
}

func (o AddressTransactionEventAnyOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressTransactionEventAnyOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["event_index"] = o.EventIndex
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *AddressTransactionEventAnyOf) UnmarshalJSON(data []byte) (err error) {
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

	varAddressTransactionEventAnyOf := _AddressTransactionEventAnyOf{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddressTransactionEventAnyOf)

	if err != nil {
		return err
	}

	*o = AddressTransactionEventAnyOf(varAddressTransactionEventAnyOf)

	return err
}

type NullableAddressTransactionEventAnyOf struct {
	value *AddressTransactionEventAnyOf
	isSet bool
}

func (v NullableAddressTransactionEventAnyOf) Get() *AddressTransactionEventAnyOf {
	return v.value
}

func (v *NullableAddressTransactionEventAnyOf) Set(val *AddressTransactionEventAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressTransactionEventAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressTransactionEventAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressTransactionEventAnyOf(val *AddressTransactionEventAnyOf) *NullableAddressTransactionEventAnyOf {
	return &NullableAddressTransactionEventAnyOf{value: val, isSet: true}
}

func (v NullableAddressTransactionEventAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressTransactionEventAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
