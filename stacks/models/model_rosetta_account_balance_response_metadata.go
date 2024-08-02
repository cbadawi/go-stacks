package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the RosettaAccountBalanceResponseMetadata type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RosettaAccountBalanceResponseMetadata{}

// RosettaAccountBalanceResponseMetadata Account-based blockchains that utilize a nonce or sequence number should include that number in the metadata. This number could be unique to the identifier or global across the account address.
type RosettaAccountBalanceResponseMetadata struct {
	SequenceNumber int32 `json:"sequence_number"`
}

type _RosettaAccountBalanceResponseMetadata RosettaAccountBalanceResponseMetadata

// NewRosettaAccountBalanceResponseMetadata instantiates a new RosettaAccountBalanceResponseMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRosettaAccountBalanceResponseMetadata(sequenceNumber int32) *RosettaAccountBalanceResponseMetadata {
	this := RosettaAccountBalanceResponseMetadata{}
	this.SequenceNumber = sequenceNumber
	return &this
}

// NewRosettaAccountBalanceResponseMetadataWithDefaults instantiates a new RosettaAccountBalanceResponseMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRosettaAccountBalanceResponseMetadataWithDefaults() *RosettaAccountBalanceResponseMetadata {
	this := RosettaAccountBalanceResponseMetadata{}
	return &this
}

// GetSequenceNumber returns the SequenceNumber field value
func (o *RosettaAccountBalanceResponseMetadata) GetSequenceNumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SequenceNumber
}

// GetSequenceNumberOk returns a tuple with the SequenceNumber field value
// and a boolean to check if the value has been set.
func (o *RosettaAccountBalanceResponseMetadata) GetSequenceNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SequenceNumber, true
}

// SetSequenceNumber sets field value
func (o *RosettaAccountBalanceResponseMetadata) SetSequenceNumber(v int32) {
	o.SequenceNumber = v
}

func (o RosettaAccountBalanceResponseMetadata) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RosettaAccountBalanceResponseMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sequence_number"] = o.SequenceNumber
	return toSerialize, nil
}

func (o *RosettaAccountBalanceResponseMetadata) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sequence_number",
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

	varRosettaAccountBalanceResponseMetadata := _RosettaAccountBalanceResponseMetadata{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRosettaAccountBalanceResponseMetadata)

	if err != nil {
		return err
	}

	*o = RosettaAccountBalanceResponseMetadata(varRosettaAccountBalanceResponseMetadata)

	return err
}

type NullableRosettaAccountBalanceResponseMetadata struct {
	value *RosettaAccountBalanceResponseMetadata
	isSet bool
}

func (v NullableRosettaAccountBalanceResponseMetadata) Get() *RosettaAccountBalanceResponseMetadata {
	return v.value
}

func (v *NullableRosettaAccountBalanceResponseMetadata) Set(val *RosettaAccountBalanceResponseMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableRosettaAccountBalanceResponseMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableRosettaAccountBalanceResponseMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRosettaAccountBalanceResponseMetadata(val *RosettaAccountBalanceResponseMetadata) *NullableRosettaAccountBalanceResponseMetadata {
	return &NullableRosettaAccountBalanceResponseMetadata{value: val, isSet: true}
}

func (v NullableRosettaAccountBalanceResponseMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRosettaAccountBalanceResponseMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
