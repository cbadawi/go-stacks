package models

import (
	"encoding/json"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the RosettaTransactionMetadata type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RosettaTransactionMetadata{}

// RosettaTransactionMetadata Transactions that are related to other transactions (like a cross-shard transaction) should include the tranaction_identifier of these transactions in the metadata.
type RosettaTransactionMetadata struct {
	// STX token transfer memo.
	Memo *string `json:"memo,omitempty"`
	// The Size
	Size *int32 `json:"size,omitempty"`
	// The locktime
	LockTime *int32 `json:"lockTime,omitempty"`
}

// NewRosettaTransactionMetadata instantiates a new RosettaTransactionMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRosettaTransactionMetadata() *RosettaTransactionMetadata {
	this := RosettaTransactionMetadata{}
	return &this
}

// NewRosettaTransactionMetadataWithDefaults instantiates a new RosettaTransactionMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRosettaTransactionMetadataWithDefaults() *RosettaTransactionMetadata {
	this := RosettaTransactionMetadata{}
	return &this
}

// GetMemo returns the Memo field value if set, zero value otherwise.
func (o *RosettaTransactionMetadata) GetMemo() string {
	if o == nil || utils.IsNil(o.Memo) {
		var ret string
		return ret
	}
	return *o.Memo
}

// GetMemoOk returns a tuple with the Memo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaTransactionMetadata) GetMemoOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Memo) {
		return nil, false
	}
	return o.Memo, true
}

// HasMemo returns a boolean if a field has been set.
func (o *RosettaTransactionMetadata) HasMemo() bool {
	if o != nil && !utils.IsNil(o.Memo) {
		return true
	}

	return false
}

// SetMemo gets a reference to the given string and assigns it to the Memo field.
func (o *RosettaTransactionMetadata) SetMemo(v string) {
	o.Memo = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *RosettaTransactionMetadata) GetSize() int32 {
	if o == nil || utils.IsNil(o.Size) {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaTransactionMetadata) GetSizeOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *RosettaTransactionMetadata) HasSize() bool {
	if o != nil && !utils.IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *RosettaTransactionMetadata) SetSize(v int32) {
	o.Size = &v
}

// GetLockTime returns the LockTime field value if set, zero value otherwise.
func (o *RosettaTransactionMetadata) GetLockTime() int32 {
	if o == nil || utils.IsNil(o.LockTime) {
		var ret int32
		return ret
	}
	return *o.LockTime
}

// GetLockTimeOk returns a tuple with the LockTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaTransactionMetadata) GetLockTimeOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.LockTime) {
		return nil, false
	}
	return o.LockTime, true
}

// HasLockTime returns a boolean if a field has been set.
func (o *RosettaTransactionMetadata) HasLockTime() bool {
	if o != nil && !utils.IsNil(o.LockTime) {
		return true
	}

	return false
}

// SetLockTime gets a reference to the given int32 and assigns it to the LockTime field.
func (o *RosettaTransactionMetadata) SetLockTime(v int32) {
	o.LockTime = &v
}

func (o RosettaTransactionMetadata) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RosettaTransactionMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Memo) {
		toSerialize["memo"] = o.Memo
	}
	if !utils.IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !utils.IsNil(o.LockTime) {
		toSerialize["lockTime"] = o.LockTime
	}
	return toSerialize, nil
}

type NullableRosettaTransactionMetadata struct {
	value *RosettaTransactionMetadata
	isSet bool
}

func (v NullableRosettaTransactionMetadata) Get() *RosettaTransactionMetadata {
	return v.value
}

func (v *NullableRosettaTransactionMetadata) Set(val *RosettaTransactionMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableRosettaTransactionMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableRosettaTransactionMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRosettaTransactionMetadata(val *RosettaTransactionMetadata) *NullableRosettaTransactionMetadata {
	return &NullableRosettaTransactionMetadata{value: val, isSet: true}
}

func (v NullableRosettaTransactionMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRosettaTransactionMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
