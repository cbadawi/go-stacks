package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the AddressUnlockSchedule type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AddressUnlockSchedule{}

// AddressUnlockSchedule Unlock schedule amount and block height
type AddressUnlockSchedule struct {
	// Micro-STX amount locked at this block height.
	Amount      string  `json:"amount"`
	BlockHeight float32 `json:"block_height"`
}

type _AddressUnlockSchedule AddressUnlockSchedule

// NewAddressUnlockSchedule instantiates a new AddressUnlockSchedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressUnlockSchedule(amount string, blockHeight float32) *AddressUnlockSchedule {
	this := AddressUnlockSchedule{}
	this.Amount = amount
	this.BlockHeight = blockHeight
	return &this
}

// NewAddressUnlockScheduleWithDefaults instantiates a new AddressUnlockSchedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressUnlockScheduleWithDefaults() *AddressUnlockSchedule {
	this := AddressUnlockSchedule{}
	return &this
}

// GetAmount returns the Amount field value
func (o *AddressUnlockSchedule) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *AddressUnlockSchedule) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *AddressUnlockSchedule) SetAmount(v string) {
	o.Amount = v
}

// GetBlockHeight returns the BlockHeight field value
func (o *AddressUnlockSchedule) GetBlockHeight() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.BlockHeight
}

// GetBlockHeightOk returns a tuple with the BlockHeight field value
// and a boolean to check if the value has been set.
func (o *AddressUnlockSchedule) GetBlockHeightOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockHeight, true
}

// SetBlockHeight sets field value
func (o *AddressUnlockSchedule) SetBlockHeight(v float32) {
	o.BlockHeight = v
}

func (o AddressUnlockSchedule) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressUnlockSchedule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	toSerialize["block_height"] = o.BlockHeight
	return toSerialize, nil
}

func (o *AddressUnlockSchedule) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"amount",
		"block_height",
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

	varAddressUnlockSchedule := _AddressUnlockSchedule{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddressUnlockSchedule)

	if err != nil {
		return err
	}

	*o = AddressUnlockSchedule(varAddressUnlockSchedule)

	return err
}

type NullableAddressUnlockSchedule struct {
	value *AddressUnlockSchedule
	isSet bool
}

func (v NullableAddressUnlockSchedule) Get() *AddressUnlockSchedule {
	return v.value
}

func (v *NullableAddressUnlockSchedule) Set(val *AddressUnlockSchedule) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressUnlockSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressUnlockSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressUnlockSchedule(val *AddressUnlockSchedule) *NullableAddressUnlockSchedule {
	return &NullableAddressUnlockSchedule{value: val, isSet: true}
}

func (v NullableAddressUnlockSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressUnlockSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
