package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the TransactionEventSmartContractLogAllOfContractLog type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &TransactionEventSmartContractLogAllOfContractLog{}

// TransactionEventSmartContractLogAllOfContractLog struct for TransactionEventSmartContractLogAllOfContractLog
type TransactionEventSmartContractLogAllOfContractLog struct {
	ContractId string                                  `json:"contract_id"`
	Topic      string                                  `json:"topic"`
	Value      PostConditionNonFungibleAllOfAssetValue `json:"value"`
}

type _TransactionEventSmartContractLogAllOfContractLog TransactionEventSmartContractLogAllOfContractLog

// NewTransactionEventSmartContractLogAllOfContractLog instantiates a new TransactionEventSmartContractLogAllOfContractLog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionEventSmartContractLogAllOfContractLog(contractId string, topic string, value PostConditionNonFungibleAllOfAssetValue) *TransactionEventSmartContractLogAllOfContractLog {
	this := TransactionEventSmartContractLogAllOfContractLog{}
	this.ContractId = contractId
	this.Topic = topic
	this.Value = value
	return &this
}

// NewTransactionEventSmartContractLogAllOfContractLogWithDefaults instantiates a new TransactionEventSmartContractLogAllOfContractLog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionEventSmartContractLogAllOfContractLogWithDefaults() *TransactionEventSmartContractLogAllOfContractLog {
	this := TransactionEventSmartContractLogAllOfContractLog{}
	return &this
}

// GetContractId returns the ContractId field value
func (o *TransactionEventSmartContractLogAllOfContractLog) GetContractId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContractId
}

// GetContractIdOk returns a tuple with the ContractId field value
// and a boolean to check if the value has been set.
func (o *TransactionEventSmartContractLogAllOfContractLog) GetContractIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractId, true
}

// SetContractId sets field value
func (o *TransactionEventSmartContractLogAllOfContractLog) SetContractId(v string) {
	o.ContractId = v
}

// GetTopic returns the Topic field value
func (o *TransactionEventSmartContractLogAllOfContractLog) GetTopic() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Topic
}

// GetTopicOk returns a tuple with the Topic field value
// and a boolean to check if the value has been set.
func (o *TransactionEventSmartContractLogAllOfContractLog) GetTopicOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Topic, true
}

// SetTopic sets field value
func (o *TransactionEventSmartContractLogAllOfContractLog) SetTopic(v string) {
	o.Topic = v
}

// GetValue returns the Value field value
func (o *TransactionEventSmartContractLogAllOfContractLog) GetValue() PostConditionNonFungibleAllOfAssetValue {
	if o == nil {
		var ret PostConditionNonFungibleAllOfAssetValue
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *TransactionEventSmartContractLogAllOfContractLog) GetValueOk() (*PostConditionNonFungibleAllOfAssetValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *TransactionEventSmartContractLogAllOfContractLog) SetValue(v PostConditionNonFungibleAllOfAssetValue) {
	o.Value = v
}

func (o TransactionEventSmartContractLogAllOfContractLog) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionEventSmartContractLogAllOfContractLog) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["contract_id"] = o.ContractId
	toSerialize["topic"] = o.Topic
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *TransactionEventSmartContractLogAllOfContractLog) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"contract_id",
		"topic",
		"value",
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

	varTransactionEventSmartContractLogAllOfContractLog := _TransactionEventSmartContractLogAllOfContractLog{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransactionEventSmartContractLogAllOfContractLog)

	if err != nil {
		return err
	}

	*o = TransactionEventSmartContractLogAllOfContractLog(varTransactionEventSmartContractLogAllOfContractLog)

	return err
}

type NullableTransactionEventSmartContractLogAllOfContractLog struct {
	value *TransactionEventSmartContractLogAllOfContractLog
	isSet bool
}

func (v NullableTransactionEventSmartContractLogAllOfContractLog) Get() *TransactionEventSmartContractLogAllOfContractLog {
	return v.value
}

func (v *NullableTransactionEventSmartContractLogAllOfContractLog) Set(val *TransactionEventSmartContractLogAllOfContractLog) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionEventSmartContractLogAllOfContractLog) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionEventSmartContractLogAllOfContractLog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionEventSmartContractLogAllOfContractLog(val *TransactionEventSmartContractLogAllOfContractLog) *NullableTransactionEventSmartContractLogAllOfContractLog {
	return &NullableTransactionEventSmartContractLogAllOfContractLog{value: val, isSet: true}
}

func (v NullableTransactionEventSmartContractLogAllOfContractLog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionEventSmartContractLogAllOfContractLog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
