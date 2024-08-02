package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the RosettaOperationIdentifier type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RosettaOperationIdentifier{}

// RosettaOperationIdentifier The operation_identifier uniquely identifies an operation within a transaction.
type RosettaOperationIdentifier struct {
	// The operation index is used to ensure each operation has a unique identifier within a transaction. This index is only relative to the transaction and NOT GLOBAL. The operations in each transaction should start from index 0. To clarify, there may not be any notion of an operation index in the blockchain being described.
	Index int32 `json:"index"`
	// Some blockchains specify an operation index that is essential for client use. For example, Bitcoin uses a network_index to identify which UTXO was used in a transaction. network_index should not be populated if there is no notion of an operation index in a blockchain (typically most account-based blockchains).
	NetworkIndex *int32 `json:"network_index,omitempty"`
}

type _RosettaOperationIdentifier RosettaOperationIdentifier

// NewRosettaOperationIdentifier instantiates a new RosettaOperationIdentifier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRosettaOperationIdentifier(index int32) *RosettaOperationIdentifier {
	this := RosettaOperationIdentifier{}
	this.Index = index
	return &this
}

// NewRosettaOperationIdentifierWithDefaults instantiates a new RosettaOperationIdentifier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRosettaOperationIdentifierWithDefaults() *RosettaOperationIdentifier {
	this := RosettaOperationIdentifier{}
	return &this
}

// GetIndex returns the Index field value
func (o *RosettaOperationIdentifier) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *RosettaOperationIdentifier) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *RosettaOperationIdentifier) SetIndex(v int32) {
	o.Index = v
}

// GetNetworkIndex returns the NetworkIndex field value if set, zero value otherwise.
func (o *RosettaOperationIdentifier) GetNetworkIndex() int32 {
	if o == nil || utils.IsNil(o.NetworkIndex) {
		var ret int32
		return ret
	}
	return *o.NetworkIndex
}

// GetNetworkIndexOk returns a tuple with the NetworkIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaOperationIdentifier) GetNetworkIndexOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.NetworkIndex) {
		return nil, false
	}
	return o.NetworkIndex, true
}

// HasNetworkIndex returns a boolean if a field has been set.
func (o *RosettaOperationIdentifier) HasNetworkIndex() bool {
	if o != nil && !utils.IsNil(o.NetworkIndex) {
		return true
	}

	return false
}

// SetNetworkIndex gets a reference to the given int32 and assigns it to the NetworkIndex field.
func (o *RosettaOperationIdentifier) SetNetworkIndex(v int32) {
	o.NetworkIndex = &v
}

func (o RosettaOperationIdentifier) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RosettaOperationIdentifier) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["index"] = o.Index
	if !utils.IsNil(o.NetworkIndex) {
		toSerialize["network_index"] = o.NetworkIndex
	}
	return toSerialize, nil
}

func (o *RosettaOperationIdentifier) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"index",
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

	varRosettaOperationIdentifier := _RosettaOperationIdentifier{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRosettaOperationIdentifier)

	if err != nil {
		return err
	}

	*o = RosettaOperationIdentifier(varRosettaOperationIdentifier)

	return err
}

type NullableRosettaOperationIdentifier struct {
	value *RosettaOperationIdentifier
	isSet bool
}

func (v NullableRosettaOperationIdentifier) Get() *RosettaOperationIdentifier {
	return v.value
}

func (v *NullableRosettaOperationIdentifier) Set(val *RosettaOperationIdentifier) {
	v.value = val
	v.isSet = true
}

func (v NullableRosettaOperationIdentifier) IsSet() bool {
	return v.isSet
}

func (v *NullableRosettaOperationIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRosettaOperationIdentifier(val *RosettaOperationIdentifier) *NullableRosettaOperationIdentifier {
	return &NullableRosettaOperationIdentifier{value: val, isSet: true}
}

func (v NullableRosettaOperationIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRosettaOperationIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
