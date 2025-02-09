package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the MempoolTransactionStatsResponseTxTypeCounts type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &MempoolTransactionStatsResponseTxTypeCounts{}

// MempoolTransactionStatsResponseTxTypeCounts Number of tranasction in the mempool, broken down by transaction type.
type MempoolTransactionStatsResponseTxTypeCounts struct {
	TokenTransfer    float32 `json:"token_transfer"`
	SmartContract    float32 `json:"smart_contract"`
	ContractCall     float32 `json:"contract_call"`
	PoisonMicroblock float32 `json:"poison_microblock"`
}

type _MempoolTransactionStatsResponseTxTypeCounts MempoolTransactionStatsResponseTxTypeCounts

// NewMempoolTransactionStatsResponseTxTypeCounts instantiates a new MempoolTransactionStatsResponseTxTypeCounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMempoolTransactionStatsResponseTxTypeCounts(tokenTransfer float32, smartContract float32, contractCall float32, poisonMicroblock float32) *MempoolTransactionStatsResponseTxTypeCounts {
	this := MempoolTransactionStatsResponseTxTypeCounts{}
	this.TokenTransfer = tokenTransfer
	this.SmartContract = smartContract
	this.ContractCall = contractCall
	this.PoisonMicroblock = poisonMicroblock
	return &this
}

// NewMempoolTransactionStatsResponseTxTypeCountsWithDefaults instantiates a new MempoolTransactionStatsResponseTxTypeCounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMempoolTransactionStatsResponseTxTypeCountsWithDefaults() *MempoolTransactionStatsResponseTxTypeCounts {
	this := MempoolTransactionStatsResponseTxTypeCounts{}
	return &this
}

// GetTokenTransfer returns the TokenTransfer field value
func (o *MempoolTransactionStatsResponseTxTypeCounts) GetTokenTransfer() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TokenTransfer
}

// GetTokenTransferOk returns a tuple with the TokenTransfer field value
// and a boolean to check if the value has been set.
func (o *MempoolTransactionStatsResponseTxTypeCounts) GetTokenTransferOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenTransfer, true
}

// SetTokenTransfer sets field value
func (o *MempoolTransactionStatsResponseTxTypeCounts) SetTokenTransfer(v float32) {
	o.TokenTransfer = v
}

// GetSmartContract returns the SmartContract field value
func (o *MempoolTransactionStatsResponseTxTypeCounts) GetSmartContract() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SmartContract
}

// GetSmartContractOk returns a tuple with the SmartContract field value
// and a boolean to check if the value has been set.
func (o *MempoolTransactionStatsResponseTxTypeCounts) GetSmartContractOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SmartContract, true
}

// SetSmartContract sets field value
func (o *MempoolTransactionStatsResponseTxTypeCounts) SetSmartContract(v float32) {
	o.SmartContract = v
}

// GetContractCall returns the ContractCall field value
func (o *MempoolTransactionStatsResponseTxTypeCounts) GetContractCall() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ContractCall
}

// GetContractCallOk returns a tuple with the ContractCall field value
// and a boolean to check if the value has been set.
func (o *MempoolTransactionStatsResponseTxTypeCounts) GetContractCallOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractCall, true
}

// SetContractCall sets field value
func (o *MempoolTransactionStatsResponseTxTypeCounts) SetContractCall(v float32) {
	o.ContractCall = v
}

// GetPoisonMicroblock returns the PoisonMicroblock field value
func (o *MempoolTransactionStatsResponseTxTypeCounts) GetPoisonMicroblock() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PoisonMicroblock
}

// GetPoisonMicroblockOk returns a tuple with the PoisonMicroblock field value
// and a boolean to check if the value has been set.
func (o *MempoolTransactionStatsResponseTxTypeCounts) GetPoisonMicroblockOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PoisonMicroblock, true
}

// SetPoisonMicroblock sets field value
func (o *MempoolTransactionStatsResponseTxTypeCounts) SetPoisonMicroblock(v float32) {
	o.PoisonMicroblock = v
}

func (o MempoolTransactionStatsResponseTxTypeCounts) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MempoolTransactionStatsResponseTxTypeCounts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["token_transfer"] = o.TokenTransfer
	toSerialize["smart_contract"] = o.SmartContract
	toSerialize["contract_call"] = o.ContractCall
	toSerialize["poison_microblock"] = o.PoisonMicroblock
	return toSerialize, nil
}

func (o *MempoolTransactionStatsResponseTxTypeCounts) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"token_transfer",
		"smart_contract",
		"contract_call",
		"poison_microblock",
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

	varMempoolTransactionStatsResponseTxTypeCounts := _MempoolTransactionStatsResponseTxTypeCounts{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMempoolTransactionStatsResponseTxTypeCounts)

	if err != nil {
		return err
	}

	*o = MempoolTransactionStatsResponseTxTypeCounts(varMempoolTransactionStatsResponseTxTypeCounts)

	return err
}

type NullableMempoolTransactionStatsResponseTxTypeCounts struct {
	value *MempoolTransactionStatsResponseTxTypeCounts
	isSet bool
}

func (v NullableMempoolTransactionStatsResponseTxTypeCounts) Get() *MempoolTransactionStatsResponseTxTypeCounts {
	return v.value
}

func (v *NullableMempoolTransactionStatsResponseTxTypeCounts) Set(val *MempoolTransactionStatsResponseTxTypeCounts) {
	v.value = val
	v.isSet = true
}

func (v NullableMempoolTransactionStatsResponseTxTypeCounts) IsSet() bool {
	return v.isSet
}

func (v *NullableMempoolTransactionStatsResponseTxTypeCounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMempoolTransactionStatsResponseTxTypeCounts(val *MempoolTransactionStatsResponseTxTypeCounts) *NullableMempoolTransactionStatsResponseTxTypeCounts {
	return &NullableMempoolTransactionStatsResponseTxTypeCounts{value: val, isSet: true}
}

func (v NullableMempoolTransactionStatsResponseTxTypeCounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMempoolTransactionStatsResponseTxTypeCounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
