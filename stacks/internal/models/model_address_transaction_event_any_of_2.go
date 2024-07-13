package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the AddressTransactionEventAnyOf2 type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AddressTransactionEventAnyOf2{}

// AddressTransactionEventAnyOf2 struct for AddressTransactionEventAnyOf2
type AddressTransactionEventAnyOf2 struct {
	Type       string                            `json:"type"`
	EventIndex int32                             `json:"event_index"`
	Data       AddressTransactionEventAnyOf2Data `json:"data"`
}

type _AddressTransactionEventAnyOf2 AddressTransactionEventAnyOf2

// NewAddressTransactionEventAnyOf2 instantiates a new AddressTransactionEventAnyOf2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressTransactionEventAnyOf2(type_ string, eventIndex int32, data AddressTransactionEventAnyOf2Data) *AddressTransactionEventAnyOf2 {
	this := AddressTransactionEventAnyOf2{}
	this.Type = type_
	this.EventIndex = eventIndex
	this.Data = data
	return &this
}

// NewAddressTransactionEventAnyOf2WithDefaults instantiates a new AddressTransactionEventAnyOf2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressTransactionEventAnyOf2WithDefaults() *AddressTransactionEventAnyOf2 {
	this := AddressTransactionEventAnyOf2{}
	return &this
}

// GetType returns the Type field value
func (o *AddressTransactionEventAnyOf2) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AddressTransactionEventAnyOf2) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AddressTransactionEventAnyOf2) SetType(v string) {
	o.Type = v
}

// GetEventIndex returns the EventIndex field value
func (o *AddressTransactionEventAnyOf2) GetEventIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EventIndex
}

// GetEventIndexOk returns a tuple with the EventIndex field value
// and a boolean to check if the value has been set.
func (o *AddressTransactionEventAnyOf2) GetEventIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventIndex, true
}

// SetEventIndex sets field value
func (o *AddressTransactionEventAnyOf2) SetEventIndex(v int32) {
	o.EventIndex = v
}

// GetData returns the Data field value
func (o *AddressTransactionEventAnyOf2) GetData() AddressTransactionEventAnyOf2Data {
	if o == nil {
		var ret AddressTransactionEventAnyOf2Data
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AddressTransactionEventAnyOf2) GetDataOk() (*AddressTransactionEventAnyOf2Data, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AddressTransactionEventAnyOf2) SetData(v AddressTransactionEventAnyOf2Data) {
	o.Data = v
}

func (o AddressTransactionEventAnyOf2) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressTransactionEventAnyOf2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["event_index"] = o.EventIndex
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *AddressTransactionEventAnyOf2) UnmarshalJSON(data []byte) (err error) {
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

	varAddressTransactionEventAnyOf2 := _AddressTransactionEventAnyOf2{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddressTransactionEventAnyOf2)

	if err != nil {
		return err
	}

	*o = AddressTransactionEventAnyOf2(varAddressTransactionEventAnyOf2)

	return err
}

type NullableAddressTransactionEventAnyOf2 struct {
	value *AddressTransactionEventAnyOf2
	isSet bool
}

func (v NullableAddressTransactionEventAnyOf2) Get() *AddressTransactionEventAnyOf2 {
	return v.value
}

func (v *NullableAddressTransactionEventAnyOf2) Set(val *AddressTransactionEventAnyOf2) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressTransactionEventAnyOf2) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressTransactionEventAnyOf2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressTransactionEventAnyOf2(val *AddressTransactionEventAnyOf2) *NullableAddressTransactionEventAnyOf2 {
	return &NullableAddressTransactionEventAnyOf2{value: val, isSet: true}
}

func (v NullableAddressTransactionEventAnyOf2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressTransactionEventAnyOf2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
