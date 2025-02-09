package models

import (
	"encoding/json"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the RosettaConstructionMetadataResponseMetadata type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RosettaConstructionMetadataResponseMetadata{}

// RosettaConstructionMetadataResponseMetadata struct for RosettaConstructionMetadataResponseMetadata
type RosettaConstructionMetadataResponseMetadata struct {
	AccountSequence *int32  `json:"account_sequence,omitempty"`
	RecentBlockHash *string `json:"recent_block_hash,omitempty"`
}

// NewRosettaConstructionMetadataResponseMetadata instantiates a new RosettaConstructionMetadataResponseMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRosettaConstructionMetadataResponseMetadata() *RosettaConstructionMetadataResponseMetadata {
	this := RosettaConstructionMetadataResponseMetadata{}
	return &this
}

// NewRosettaConstructionMetadataResponseMetadataWithDefaults instantiates a new RosettaConstructionMetadataResponseMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRosettaConstructionMetadataResponseMetadataWithDefaults() *RosettaConstructionMetadataResponseMetadata {
	this := RosettaConstructionMetadataResponseMetadata{}
	return &this
}

// GetAccountSequence returns the AccountSequence field value if set, zero value otherwise.
func (o *RosettaConstructionMetadataResponseMetadata) GetAccountSequence() int32 {
	if o == nil || utils.IsNil(o.AccountSequence) {
		var ret int32
		return ret
	}
	return *o.AccountSequence
}

// GetAccountSequenceOk returns a tuple with the AccountSequence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaConstructionMetadataResponseMetadata) GetAccountSequenceOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.AccountSequence) {
		return nil, false
	}
	return o.AccountSequence, true
}

// HasAccountSequence returns a boolean if a field has been set.
func (o *RosettaConstructionMetadataResponseMetadata) HasAccountSequence() bool {
	if o != nil && !utils.IsNil(o.AccountSequence) {
		return true
	}

	return false
}

// SetAccountSequence gets a reference to the given int32 and assigns it to the AccountSequence field.
func (o *RosettaConstructionMetadataResponseMetadata) SetAccountSequence(v int32) {
	o.AccountSequence = &v
}

// GetRecentBlockHash returns the RecentBlockHash field value if set, zero value otherwise.
func (o *RosettaConstructionMetadataResponseMetadata) GetRecentBlockHash() string {
	if o == nil || utils.IsNil(o.RecentBlockHash) {
		var ret string
		return ret
	}
	return *o.RecentBlockHash
}

// GetRecentBlockHashOk returns a tuple with the RecentBlockHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaConstructionMetadataResponseMetadata) GetRecentBlockHashOk() (*string, bool) {
	if o == nil || utils.IsNil(o.RecentBlockHash) {
		return nil, false
	}
	return o.RecentBlockHash, true
}

// HasRecentBlockHash returns a boolean if a field has been set.
func (o *RosettaConstructionMetadataResponseMetadata) HasRecentBlockHash() bool {
	if o != nil && !utils.IsNil(o.RecentBlockHash) {
		return true
	}

	return false
}

// SetRecentBlockHash gets a reference to the given string and assigns it to the RecentBlockHash field.
func (o *RosettaConstructionMetadataResponseMetadata) SetRecentBlockHash(v string) {
	o.RecentBlockHash = &v
}

func (o RosettaConstructionMetadataResponseMetadata) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RosettaConstructionMetadataResponseMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.AccountSequence) {
		toSerialize["account_sequence"] = o.AccountSequence
	}
	if !utils.IsNil(o.RecentBlockHash) {
		toSerialize["recent_block_hash"] = o.RecentBlockHash
	}
	return toSerialize, nil
}

type NullableRosettaConstructionMetadataResponseMetadata struct {
	value *RosettaConstructionMetadataResponseMetadata
	isSet bool
}

func (v NullableRosettaConstructionMetadataResponseMetadata) Get() *RosettaConstructionMetadataResponseMetadata {
	return v.value
}

func (v *NullableRosettaConstructionMetadataResponseMetadata) Set(val *RosettaConstructionMetadataResponseMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableRosettaConstructionMetadataResponseMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableRosettaConstructionMetadataResponseMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRosettaConstructionMetadataResponseMetadata(val *RosettaConstructionMetadataResponseMetadata) *NullableRosettaConstructionMetadataResponseMetadata {
	return &NullableRosettaConstructionMetadataResponseMetadata{value: val, isSet: true}
}

func (v NullableRosettaConstructionMetadataResponseMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRosettaConstructionMetadataResponseMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
