package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the RosettaRelatedOperation type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RosettaRelatedOperation{}

// RosettaRelatedOperation Restrict referenced related_operations to identifier indexes < the current operation_identifier.index. This ensures there exists a clear DAG-structure of relations. Since operations are one-sided, one could imagine relating operations in a single transfer or linking operations in a call tree.
type RosettaRelatedOperation struct {
	// Describes the index of related operation.
	Index int32 `json:"index"`
	// Some blockchains specify an operation index that is essential for client use. network_index should not be populated if there is no notion of an operation index in a blockchain (typically most account-based blockchains).
	NetworkIndex *int32 `json:"network_index,omitempty"`
}

type _RosettaRelatedOperation RosettaRelatedOperation

// NewRosettaRelatedOperation instantiates a new RosettaRelatedOperation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRosettaRelatedOperation(index int32) *RosettaRelatedOperation {
	this := RosettaRelatedOperation{}
	this.Index = index
	return &this
}

// NewRosettaRelatedOperationWithDefaults instantiates a new RosettaRelatedOperation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRosettaRelatedOperationWithDefaults() *RosettaRelatedOperation {
	this := RosettaRelatedOperation{}
	return &this
}

// GetIndex returns the Index field value
func (o *RosettaRelatedOperation) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *RosettaRelatedOperation) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *RosettaRelatedOperation) SetIndex(v int32) {
	o.Index = v
}

// GetNetworkIndex returns the NetworkIndex field value if set, zero value otherwise.
func (o *RosettaRelatedOperation) GetNetworkIndex() int32 {
	if o == nil || utils.IsNil(o.NetworkIndex) {
		var ret int32
		return ret
	}
	return *o.NetworkIndex
}

// GetNetworkIndexOk returns a tuple with the NetworkIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaRelatedOperation) GetNetworkIndexOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.NetworkIndex) {
		return nil, false
	}
	return o.NetworkIndex, true
}

// HasNetworkIndex returns a boolean if a field has been set.
func (o *RosettaRelatedOperation) HasNetworkIndex() bool {
	if o != nil && !utils.IsNil(o.NetworkIndex) {
		return true
	}

	return false
}

// SetNetworkIndex gets a reference to the given int32 and assigns it to the NetworkIndex field.
func (o *RosettaRelatedOperation) SetNetworkIndex(v int32) {
	o.NetworkIndex = &v
}

func (o RosettaRelatedOperation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RosettaRelatedOperation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["index"] = o.Index
	if !utils.IsNil(o.NetworkIndex) {
		toSerialize["network_index"] = o.NetworkIndex
	}
	return toSerialize, nil
}

func (o *RosettaRelatedOperation) UnmarshalJSON(data []byte) (err error) {
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

	varRosettaRelatedOperation := _RosettaRelatedOperation{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRosettaRelatedOperation)

	if err != nil {
		return err
	}

	*o = RosettaRelatedOperation(varRosettaRelatedOperation)

	return err
}

type NullableRosettaRelatedOperation struct {
	value *RosettaRelatedOperation
	isSet bool
}

func (v NullableRosettaRelatedOperation) Get() *RosettaRelatedOperation {
	return v.value
}

func (v *NullableRosettaRelatedOperation) Set(val *RosettaRelatedOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableRosettaRelatedOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableRosettaRelatedOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRosettaRelatedOperation(val *RosettaRelatedOperation) *NullableRosettaRelatedOperation {
	return &NullableRosettaRelatedOperation{value: val, isSet: true}
}

func (v NullableRosettaRelatedOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRosettaRelatedOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
