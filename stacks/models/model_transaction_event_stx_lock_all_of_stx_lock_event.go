package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the TransactionEventStxLockAllOfStxLockEvent type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &TransactionEventStxLockAllOfStxLockEvent{}

// TransactionEventStxLockAllOfStxLockEvent struct for TransactionEventStxLockAllOfStxLockEvent
type TransactionEventStxLockAllOfStxLockEvent struct {
	LockedAmount  string `json:"locked_amount"`
	UnlockHeight  int32  `json:"unlock_height"`
	LockedAddress string `json:"locked_address"`
}

type _TransactionEventStxLockAllOfStxLockEvent TransactionEventStxLockAllOfStxLockEvent

// NewTransactionEventStxLockAllOfStxLockEvent instantiates a new TransactionEventStxLockAllOfStxLockEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionEventStxLockAllOfStxLockEvent(lockedAmount string, unlockHeight int32, lockedAddress string) *TransactionEventStxLockAllOfStxLockEvent {
	this := TransactionEventStxLockAllOfStxLockEvent{}
	this.LockedAmount = lockedAmount
	this.UnlockHeight = unlockHeight
	this.LockedAddress = lockedAddress
	return &this
}

// NewTransactionEventStxLockAllOfStxLockEventWithDefaults instantiates a new TransactionEventStxLockAllOfStxLockEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionEventStxLockAllOfStxLockEventWithDefaults() *TransactionEventStxLockAllOfStxLockEvent {
	this := TransactionEventStxLockAllOfStxLockEvent{}
	return &this
}

// GetLockedAmount returns the LockedAmount field value
func (o *TransactionEventStxLockAllOfStxLockEvent) GetLockedAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LockedAmount
}

// GetLockedAmountOk returns a tuple with the LockedAmount field value
// and a boolean to check if the value has been set.
func (o *TransactionEventStxLockAllOfStxLockEvent) GetLockedAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LockedAmount, true
}

// SetLockedAmount sets field value
func (o *TransactionEventStxLockAllOfStxLockEvent) SetLockedAmount(v string) {
	o.LockedAmount = v
}

// GetUnlockHeight returns the UnlockHeight field value
func (o *TransactionEventStxLockAllOfStxLockEvent) GetUnlockHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UnlockHeight
}

// GetUnlockHeightOk returns a tuple with the UnlockHeight field value
// and a boolean to check if the value has been set.
func (o *TransactionEventStxLockAllOfStxLockEvent) GetUnlockHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnlockHeight, true
}

// SetUnlockHeight sets field value
func (o *TransactionEventStxLockAllOfStxLockEvent) SetUnlockHeight(v int32) {
	o.UnlockHeight = v
}

// GetLockedAddress returns the LockedAddress field value
func (o *TransactionEventStxLockAllOfStxLockEvent) GetLockedAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LockedAddress
}

// GetLockedAddressOk returns a tuple with the LockedAddress field value
// and a boolean to check if the value has been set.
func (o *TransactionEventStxLockAllOfStxLockEvent) GetLockedAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LockedAddress, true
}

// SetLockedAddress sets field value
func (o *TransactionEventStxLockAllOfStxLockEvent) SetLockedAddress(v string) {
	o.LockedAddress = v
}

func (o TransactionEventStxLockAllOfStxLockEvent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionEventStxLockAllOfStxLockEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["locked_amount"] = o.LockedAmount
	toSerialize["unlock_height"] = o.UnlockHeight
	toSerialize["locked_address"] = o.LockedAddress
	return toSerialize, nil
}

func (o *TransactionEventStxLockAllOfStxLockEvent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"locked_amount",
		"unlock_height",
		"locked_address",
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

	varTransactionEventStxLockAllOfStxLockEvent := _TransactionEventStxLockAllOfStxLockEvent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransactionEventStxLockAllOfStxLockEvent)

	if err != nil {
		return err
	}

	*o = TransactionEventStxLockAllOfStxLockEvent(varTransactionEventStxLockAllOfStxLockEvent)

	return err
}

type NullableTransactionEventStxLockAllOfStxLockEvent struct {
	value *TransactionEventStxLockAllOfStxLockEvent
	isSet bool
}

func (v NullableTransactionEventStxLockAllOfStxLockEvent) Get() *TransactionEventStxLockAllOfStxLockEvent {
	return v.value
}

func (v *NullableTransactionEventStxLockAllOfStxLockEvent) Set(val *TransactionEventStxLockAllOfStxLockEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionEventStxLockAllOfStxLockEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionEventStxLockAllOfStxLockEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionEventStxLockAllOfStxLockEvent(val *TransactionEventStxLockAllOfStxLockEvent) *NullableTransactionEventStxLockAllOfStxLockEvent {
	return &NullableTransactionEventStxLockAllOfStxLockEvent{value: val, isSet: true}
}

func (v NullableTransactionEventStxLockAllOfStxLockEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionEventStxLockAllOfStxLockEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
