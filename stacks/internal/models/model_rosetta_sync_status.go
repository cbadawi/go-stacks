package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the RosettaSyncStatus type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RosettaSyncStatus{}

// RosettaSyncStatus SyncStatus is used to provide additional context about an implementation's sync status. It is often used to indicate that an implementation is healthy when it cannot be queried until some sync phase occurs. If an implementation is immediately queryable, this model is often not populated.
type RosettaSyncStatus struct {
	// CurrentIndex is the index of the last synced block in the current stage.
	CurrentIndex int32 `json:"current_index"`
	// TargetIndex is the index of the block that the implementation is attempting to sync to in the current stage.
	TargetIndex *int32 `json:"target_index,omitempty"`
	// Stage is the phase of the sync process.
	Stage *string `json:"stage,omitempty"`
	// Synced indicates if an implementation has synced up to the most recent block.
	Synced *bool `json:"synced,omitempty"`
}

type _RosettaSyncStatus RosettaSyncStatus

// NewRosettaSyncStatus instantiates a new RosettaSyncStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRosettaSyncStatus(currentIndex int32) *RosettaSyncStatus {
	this := RosettaSyncStatus{}
	this.CurrentIndex = currentIndex
	return &this
}

// NewRosettaSyncStatusWithDefaults instantiates a new RosettaSyncStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRosettaSyncStatusWithDefaults() *RosettaSyncStatus {
	this := RosettaSyncStatus{}
	return &this
}

// GetCurrentIndex returns the CurrentIndex field value
func (o *RosettaSyncStatus) GetCurrentIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CurrentIndex
}

// GetCurrentIndexOk returns a tuple with the CurrentIndex field value
// and a boolean to check if the value has been set.
func (o *RosettaSyncStatus) GetCurrentIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentIndex, true
}

// SetCurrentIndex sets field value
func (o *RosettaSyncStatus) SetCurrentIndex(v int32) {
	o.CurrentIndex = v
}

// GetTargetIndex returns the TargetIndex field value if set, zero value otherwise.
func (o *RosettaSyncStatus) GetTargetIndex() int32 {
	if o == nil || utils.IsNil(o.TargetIndex) {
		var ret int32
		return ret
	}
	return *o.TargetIndex
}

// GetTargetIndexOk returns a tuple with the TargetIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaSyncStatus) GetTargetIndexOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.TargetIndex) {
		return nil, false
	}
	return o.TargetIndex, true
}

// HasTargetIndex returns a boolean if a field has been set.
func (o *RosettaSyncStatus) HasTargetIndex() bool {
	if o != nil && !utils.IsNil(o.TargetIndex) {
		return true
	}

	return false
}

// SetTargetIndex gets a reference to the given int32 and assigns it to the TargetIndex field.
func (o *RosettaSyncStatus) SetTargetIndex(v int32) {
	o.TargetIndex = &v
}

// GetStage returns the Stage field value if set, zero value otherwise.
func (o *RosettaSyncStatus) GetStage() string {
	if o == nil || utils.IsNil(o.Stage) {
		var ret string
		return ret
	}
	return *o.Stage
}

// GetStageOk returns a tuple with the Stage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaSyncStatus) GetStageOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Stage) {
		return nil, false
	}
	return o.Stage, true
}

// HasStage returns a boolean if a field has been set.
func (o *RosettaSyncStatus) HasStage() bool {
	if o != nil && !utils.IsNil(o.Stage) {
		return true
	}

	return false
}

// SetStage gets a reference to the given string and assigns it to the Stage field.
func (o *RosettaSyncStatus) SetStage(v string) {
	o.Stage = &v
}

// GetSynced returns the Synced field value if set, zero value otherwise.
func (o *RosettaSyncStatus) GetSynced() bool {
	if o == nil || utils.IsNil(o.Synced) {
		var ret bool
		return ret
	}
	return *o.Synced
}

// GetSyncedOk returns a tuple with the Synced field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaSyncStatus) GetSyncedOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.Synced) {
		return nil, false
	}
	return o.Synced, true
}

// HasSynced returns a boolean if a field has been set.
func (o *RosettaSyncStatus) HasSynced() bool {
	if o != nil && !utils.IsNil(o.Synced) {
		return true
	}

	return false
}

// SetSynced gets a reference to the given bool and assigns it to the Synced field.
func (o *RosettaSyncStatus) SetSynced(v bool) {
	o.Synced = &v
}

func (o RosettaSyncStatus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RosettaSyncStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["current_index"] = o.CurrentIndex
	if !utils.IsNil(o.TargetIndex) {
		toSerialize["target_index"] = o.TargetIndex
	}
	if !utils.IsNil(o.Stage) {
		toSerialize["stage"] = o.Stage
	}
	if !utils.IsNil(o.Synced) {
		toSerialize["synced"] = o.Synced
	}
	return toSerialize, nil
}

func (o *RosettaSyncStatus) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"current_index",
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

	varRosettaSyncStatus := _RosettaSyncStatus{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRosettaSyncStatus)

	if err != nil {
		return err
	}

	*o = RosettaSyncStatus(varRosettaSyncStatus)

	return err
}

type NullableRosettaSyncStatus struct {
	value *RosettaSyncStatus
	isSet bool
}

func (v NullableRosettaSyncStatus) Get() *RosettaSyncStatus {
	return v.value
}

func (v *NullableRosettaSyncStatus) Set(val *RosettaSyncStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableRosettaSyncStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableRosettaSyncStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRosettaSyncStatus(val *RosettaSyncStatus) *NullableRosettaSyncStatus {
	return &NullableRosettaSyncStatus{value: val, isSet: true}
}

func (v NullableRosettaSyncStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRosettaSyncStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
