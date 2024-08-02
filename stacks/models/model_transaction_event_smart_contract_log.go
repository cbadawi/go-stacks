package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the TransactionEventSmartContractLog type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &TransactionEventSmartContractLog{}

// TransactionEventSmartContractLog Only present in `smart_contract` and `contract_call` tx types.
type TransactionEventSmartContractLog struct {
	EventIndex  int32                                            `json:"event_index"`
	EventType   string                                           `json:"event_type"`
	TxId        string                                           `json:"tx_id"`
	ContractLog TransactionEventSmartContractLogAllOfContractLog `json:"contract_log"`
}

type _TransactionEventSmartContractLog TransactionEventSmartContractLog

// NewTransactionEventSmartContractLog instantiates a new TransactionEventSmartContractLog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionEventSmartContractLog(eventIndex int32, eventType string, txId string, contractLog TransactionEventSmartContractLogAllOfContractLog) *TransactionEventSmartContractLog {
	this := TransactionEventSmartContractLog{}
	this.EventIndex = eventIndex
	this.EventType = eventType
	this.TxId = txId
	this.ContractLog = contractLog
	return &this
}

// NewTransactionEventSmartContractLogWithDefaults instantiates a new TransactionEventSmartContractLog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionEventSmartContractLogWithDefaults() *TransactionEventSmartContractLog {
	this := TransactionEventSmartContractLog{}
	return &this
}

// GetEventIndex returns the EventIndex field value
func (o *TransactionEventSmartContractLog) GetEventIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EventIndex
}

// GetEventIndexOk returns a tuple with the EventIndex field value
// and a boolean to check if the value has been set.
func (o *TransactionEventSmartContractLog) GetEventIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventIndex, true
}

// SetEventIndex sets field value
func (o *TransactionEventSmartContractLog) SetEventIndex(v int32) {
	o.EventIndex = v
}

// GetEventType returns the EventType field value
func (o *TransactionEventSmartContractLog) GetEventType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *TransactionEventSmartContractLog) GetEventTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *TransactionEventSmartContractLog) SetEventType(v string) {
	o.EventType = v
}

// GetTxId returns the TxId field value
func (o *TransactionEventSmartContractLog) GetTxId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TxId
}

// GetTxIdOk returns a tuple with the TxId field value
// and a boolean to check if the value has been set.
func (o *TransactionEventSmartContractLog) GetTxIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxId, true
}

// SetTxId sets field value
func (o *TransactionEventSmartContractLog) SetTxId(v string) {
	o.TxId = v
}

// GetContractLog returns the ContractLog field value
func (o *TransactionEventSmartContractLog) GetContractLog() TransactionEventSmartContractLogAllOfContractLog {
	if o == nil {
		var ret TransactionEventSmartContractLogAllOfContractLog
		return ret
	}

	return o.ContractLog
}

// GetContractLogOk returns a tuple with the ContractLog field value
// and a boolean to check if the value has been set.
func (o *TransactionEventSmartContractLog) GetContractLogOk() (*TransactionEventSmartContractLogAllOfContractLog, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractLog, true
}

// SetContractLog sets field value
func (o *TransactionEventSmartContractLog) SetContractLog(v TransactionEventSmartContractLogAllOfContractLog) {
	o.ContractLog = v
}

func (o TransactionEventSmartContractLog) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionEventSmartContractLog) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["event_index"] = o.EventIndex
	toSerialize["event_type"] = o.EventType
	toSerialize["tx_id"] = o.TxId
	toSerialize["contract_log"] = o.ContractLog
	return toSerialize, nil
}

func (o *TransactionEventSmartContractLog) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"event_index",
		"event_type",
		"tx_id",
		"contract_log",
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

	varTransactionEventSmartContractLog := _TransactionEventSmartContractLog{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransactionEventSmartContractLog)

	if err != nil {
		return err
	}

	*o = TransactionEventSmartContractLog(varTransactionEventSmartContractLog)

	return err
}

type NullableTransactionEventSmartContractLog struct {
	value *TransactionEventSmartContractLog
	isSet bool
}

func (v NullableTransactionEventSmartContractLog) Get() *TransactionEventSmartContractLog {
	return v.value
}

func (v *NullableTransactionEventSmartContractLog) Set(val *TransactionEventSmartContractLog) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionEventSmartContractLog) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionEventSmartContractLog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionEventSmartContractLog(val *TransactionEventSmartContractLog) *NullableTransactionEventSmartContractLog {
	return &NullableTransactionEventSmartContractLog{value: val, isSet: true}
}

func (v NullableTransactionEventSmartContractLog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionEventSmartContractLog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
