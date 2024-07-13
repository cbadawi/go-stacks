package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the TransactionEventStxLock type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &TransactionEventStxLock{}

// TransactionEventStxLock Only present in `smart_contract` and `contract_call` tx types.
type TransactionEventStxLock struct {
	EventIndex   int32                                    `json:"event_index"`
	EventType    string                                   `json:"event_type"`
	TxId         string                                   `json:"tx_id"`
	StxLockEvent TransactionEventStxLockAllOfStxLockEvent `json:"stx_lock_event"`
}

type _TransactionEventStxLock TransactionEventStxLock

// NewTransactionEventStxLock instantiates a new TransactionEventStxLock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionEventStxLock(eventIndex int32, eventType string, txId string, stxLockEvent TransactionEventStxLockAllOfStxLockEvent) *TransactionEventStxLock {
	this := TransactionEventStxLock{}
	this.EventIndex = eventIndex
	this.EventType = eventType
	this.TxId = txId
	this.StxLockEvent = stxLockEvent
	return &this
}

// NewTransactionEventStxLockWithDefaults instantiates a new TransactionEventStxLock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionEventStxLockWithDefaults() *TransactionEventStxLock {
	this := TransactionEventStxLock{}
	return &this
}

// GetEventIndex returns the EventIndex field value
func (o *TransactionEventStxLock) GetEventIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EventIndex
}

// GetEventIndexOk returns a tuple with the EventIndex field value
// and a boolean to check if the value has been set.
func (o *TransactionEventStxLock) GetEventIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventIndex, true
}

// SetEventIndex sets field value
func (o *TransactionEventStxLock) SetEventIndex(v int32) {
	o.EventIndex = v
}

// GetEventType returns the EventType field value
func (o *TransactionEventStxLock) GetEventType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *TransactionEventStxLock) GetEventTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *TransactionEventStxLock) SetEventType(v string) {
	o.EventType = v
}

// GetTxId returns the TxId field value
func (o *TransactionEventStxLock) GetTxId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TxId
}

// GetTxIdOk returns a tuple with the TxId field value
// and a boolean to check if the value has been set.
func (o *TransactionEventStxLock) GetTxIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxId, true
}

// SetTxId sets field value
func (o *TransactionEventStxLock) SetTxId(v string) {
	o.TxId = v
}

// GetStxLockEvent returns the StxLockEvent field value
func (o *TransactionEventStxLock) GetStxLockEvent() TransactionEventStxLockAllOfStxLockEvent {
	if o == nil {
		var ret TransactionEventStxLockAllOfStxLockEvent
		return ret
	}

	return o.StxLockEvent
}

// GetStxLockEventOk returns a tuple with the StxLockEvent field value
// and a boolean to check if the value has been set.
func (o *TransactionEventStxLock) GetStxLockEventOk() (*TransactionEventStxLockAllOfStxLockEvent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StxLockEvent, true
}

// SetStxLockEvent sets field value
func (o *TransactionEventStxLock) SetStxLockEvent(v TransactionEventStxLockAllOfStxLockEvent) {
	o.StxLockEvent = v
}

func (o TransactionEventStxLock) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionEventStxLock) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["event_index"] = o.EventIndex
	toSerialize["event_type"] = o.EventType
	toSerialize["tx_id"] = o.TxId
	toSerialize["stx_lock_event"] = o.StxLockEvent
	return toSerialize, nil
}

func (o *TransactionEventStxLock) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"event_index",
		"event_type",
		"tx_id",
		"stx_lock_event",
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

	varTransactionEventStxLock := _TransactionEventStxLock{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransactionEventStxLock)

	if err != nil {
		return err
	}

	*o = TransactionEventStxLock(varTransactionEventStxLock)

	return err
}

type NullableTransactionEventStxLock struct {
	value *TransactionEventStxLock
	isSet bool
}

func (v NullableTransactionEventStxLock) Get() *TransactionEventStxLock {
	return v.value
}

func (v *NullableTransactionEventStxLock) Set(val *TransactionEventStxLock) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionEventStxLock) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionEventStxLock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionEventStxLock(val *TransactionEventStxLock) *NullableTransactionEventStxLock {
	return &NullableTransactionEventStxLock{value: val, isSet: true}
}

func (v NullableTransactionEventStxLock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionEventStxLock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
