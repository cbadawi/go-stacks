package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the AddressTokenOfferingLocked type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AddressTokenOfferingLocked{}

// AddressTokenOfferingLocked Token Offering Locked
type AddressTokenOfferingLocked struct {
	// Micro-STX amount still locked at current block height.
	TotalLocked string `json:"total_locked"`
	// Micro-STX amount unlocked at current block height.
	TotalUnlocked  string                  `json:"total_unlocked"`
	UnlockSchedule []AddressUnlockSchedule `json:"unlock_schedule"`
}

type _AddressTokenOfferingLocked AddressTokenOfferingLocked

// NewAddressTokenOfferingLocked instantiates a new AddressTokenOfferingLocked object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressTokenOfferingLocked(totalLocked string, totalUnlocked string, unlockSchedule []AddressUnlockSchedule) *AddressTokenOfferingLocked {
	this := AddressTokenOfferingLocked{}
	this.TotalLocked = totalLocked
	this.TotalUnlocked = totalUnlocked
	this.UnlockSchedule = unlockSchedule
	return &this
}

// NewAddressTokenOfferingLockedWithDefaults instantiates a new AddressTokenOfferingLocked object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressTokenOfferingLockedWithDefaults() *AddressTokenOfferingLocked {
	this := AddressTokenOfferingLocked{}
	return &this
}

// GetTotalLocked returns the TotalLocked field value
func (o *AddressTokenOfferingLocked) GetTotalLocked() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TotalLocked
}

// GetTotalLockedOk returns a tuple with the TotalLocked field value
// and a boolean to check if the value has been set.
func (o *AddressTokenOfferingLocked) GetTotalLockedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalLocked, true
}

// SetTotalLocked sets field value
func (o *AddressTokenOfferingLocked) SetTotalLocked(v string) {
	o.TotalLocked = v
}

// GetTotalUnlocked returns the TotalUnlocked field value
func (o *AddressTokenOfferingLocked) GetTotalUnlocked() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TotalUnlocked
}

// GetTotalUnlockedOk returns a tuple with the TotalUnlocked field value
// and a boolean to check if the value has been set.
func (o *AddressTokenOfferingLocked) GetTotalUnlockedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalUnlocked, true
}

// SetTotalUnlocked sets field value
func (o *AddressTokenOfferingLocked) SetTotalUnlocked(v string) {
	o.TotalUnlocked = v
}

// GetUnlockSchedule returns the UnlockSchedule field value
func (o *AddressTokenOfferingLocked) GetUnlockSchedule() []AddressUnlockSchedule {
	if o == nil {
		var ret []AddressUnlockSchedule
		return ret
	}

	return o.UnlockSchedule
}

// GetUnlockScheduleOk returns a tuple with the UnlockSchedule field value
// and a boolean to check if the value has been set.
func (o *AddressTokenOfferingLocked) GetUnlockScheduleOk() ([]AddressUnlockSchedule, bool) {
	if o == nil {
		return nil, false
	}
	return o.UnlockSchedule, true
}

// SetUnlockSchedule sets field value
func (o *AddressTokenOfferingLocked) SetUnlockSchedule(v []AddressUnlockSchedule) {
	o.UnlockSchedule = v
}

func (o AddressTokenOfferingLocked) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressTokenOfferingLocked) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["total_locked"] = o.TotalLocked
	toSerialize["total_unlocked"] = o.TotalUnlocked
	toSerialize["unlock_schedule"] = o.UnlockSchedule
	return toSerialize, nil
}

func (o *AddressTokenOfferingLocked) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"total_locked",
		"total_unlocked",
		"unlock_schedule",
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

	varAddressTokenOfferingLocked := _AddressTokenOfferingLocked{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddressTokenOfferingLocked)

	if err != nil {
		return err
	}

	*o = AddressTokenOfferingLocked(varAddressTokenOfferingLocked)

	return err
}

type NullableAddressTokenOfferingLocked struct {
	value *AddressTokenOfferingLocked
	isSet bool
}

func (v NullableAddressTokenOfferingLocked) Get() *AddressTokenOfferingLocked {
	return v.value
}

func (v *NullableAddressTokenOfferingLocked) Set(val *AddressTokenOfferingLocked) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressTokenOfferingLocked) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressTokenOfferingLocked) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressTokenOfferingLocked(val *AddressTokenOfferingLocked) *NullableAddressTokenOfferingLocked {
	return &NullableAddressTokenOfferingLocked{value: val, isSet: true}
}

func (v NullableAddressTokenOfferingLocked) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressTokenOfferingLocked) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
