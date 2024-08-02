package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the SmartContractStatus type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &SmartContractStatus{}

// SmartContractStatus Deployment status of a smart contract
type SmartContractStatus struct {
	// Smart contract deployment transaction status
	Status string `json:"status"`
	// Deployment transaction ID
	TxId string `json:"tx_id"`
	// Smart contract ID
	ContractId string `json:"contract_id"`
	// Height of the transaction confirmation block
	BlockHeight *int32 `json:"block_height,omitempty"`
}

type _SmartContractStatus SmartContractStatus

// NewSmartContractStatus instantiates a new SmartContractStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmartContractStatus(status string, txId string, contractId string) *SmartContractStatus {
	this := SmartContractStatus{}
	this.Status = status
	this.TxId = txId
	this.ContractId = contractId
	return &this
}

// NewSmartContractStatusWithDefaults instantiates a new SmartContractStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmartContractStatusWithDefaults() *SmartContractStatus {
	this := SmartContractStatus{}
	return &this
}

// GetStatus returns the Status field value
func (o *SmartContractStatus) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *SmartContractStatus) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *SmartContractStatus) SetStatus(v string) {
	o.Status = v
}

// GetTxId returns the TxId field value
func (o *SmartContractStatus) GetTxId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TxId
}

// GetTxIdOk returns a tuple with the TxId field value
// and a boolean to check if the value has been set.
func (o *SmartContractStatus) GetTxIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxId, true
}

// SetTxId sets field value
func (o *SmartContractStatus) SetTxId(v string) {
	o.TxId = v
}

// GetContractId returns the ContractId field value
func (o *SmartContractStatus) GetContractId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContractId
}

// GetContractIdOk returns a tuple with the ContractId field value
// and a boolean to check if the value has been set.
func (o *SmartContractStatus) GetContractIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractId, true
}

// SetContractId sets field value
func (o *SmartContractStatus) SetContractId(v string) {
	o.ContractId = v
}

// GetBlockHeight returns the BlockHeight field value if set, zero value otherwise.
func (o *SmartContractStatus) GetBlockHeight() int32 {
	if o == nil || utils.IsNil(o.BlockHeight) {
		var ret int32
		return ret
	}
	return *o.BlockHeight
}

// GetBlockHeightOk returns a tuple with the BlockHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartContractStatus) GetBlockHeightOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.BlockHeight) {
		return nil, false
	}
	return o.BlockHeight, true
}

// HasBlockHeight returns a boolean if a field has been set.
func (o *SmartContractStatus) HasBlockHeight() bool {
	if o != nil && !utils.IsNil(o.BlockHeight) {
		return true
	}

	return false
}

// SetBlockHeight gets a reference to the given int32 and assigns it to the BlockHeight field.
func (o *SmartContractStatus) SetBlockHeight(v int32) {
	o.BlockHeight = &v
}

func (o SmartContractStatus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SmartContractStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	toSerialize["tx_id"] = o.TxId
	toSerialize["contract_id"] = o.ContractId
	if !utils.IsNil(o.BlockHeight) {
		toSerialize["block_height"] = o.BlockHeight
	}
	return toSerialize, nil
}

func (o *SmartContractStatus) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
		"tx_id",
		"contract_id",
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

	varSmartContractStatus := _SmartContractStatus{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSmartContractStatus)

	if err != nil {
		return err
	}

	*o = SmartContractStatus(varSmartContractStatus)

	return err
}

type NullableSmartContractStatus struct {
	value *SmartContractStatus
	isSet bool
}

func (v NullableSmartContractStatus) Get() *SmartContractStatus {
	return v.value
}

func (v *NullableSmartContractStatus) Set(val *SmartContractStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSmartContractStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSmartContractStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmartContractStatus(val *SmartContractStatus) *NullableSmartContractStatus {
	return &NullableSmartContractStatus{value: val, isSet: true}
}

func (v NullableSmartContractStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmartContractStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
