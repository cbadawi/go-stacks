package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the MempoolTransactionStatsResponseTxSimpleFeeAverages type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &MempoolTransactionStatsResponseTxSimpleFeeAverages{}

// MempoolTransactionStatsResponseTxSimpleFeeAverages The simple mean (average) transaction fee, broken down by transaction type. Note that this does not factor in actual execution costs. The average fee is not a reliable metric for calculating a fee for a new transaction.
type MempoolTransactionStatsResponseTxSimpleFeeAverages struct {
	TokenTransfer    MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer `json:"token_transfer"`
	SmartContract    MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer `json:"smart_contract"`
	ContractCall     MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer `json:"contract_call"`
	PoisonMicroblock MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer `json:"poison_microblock"`
}

type _MempoolTransactionStatsResponseTxSimpleFeeAverages MempoolTransactionStatsResponseTxSimpleFeeAverages

// NewMempoolTransactionStatsResponseTxSimpleFeeAverages instantiates a new MempoolTransactionStatsResponseTxSimpleFeeAverages object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMempoolTransactionStatsResponseTxSimpleFeeAverages(tokenTransfer MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer, smartContract MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer, contractCall MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer, poisonMicroblock MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer) *MempoolTransactionStatsResponseTxSimpleFeeAverages {
	this := MempoolTransactionStatsResponseTxSimpleFeeAverages{}
	this.TokenTransfer = tokenTransfer
	this.SmartContract = smartContract
	this.ContractCall = contractCall
	this.PoisonMicroblock = poisonMicroblock
	return &this
}

// NewMempoolTransactionStatsResponseTxSimpleFeeAveragesWithDefaults instantiates a new MempoolTransactionStatsResponseTxSimpleFeeAverages object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMempoolTransactionStatsResponseTxSimpleFeeAveragesWithDefaults() *MempoolTransactionStatsResponseTxSimpleFeeAverages {
	this := MempoolTransactionStatsResponseTxSimpleFeeAverages{}
	return &this
}

// GetTokenTransfer returns the TokenTransfer field value
func (o *MempoolTransactionStatsResponseTxSimpleFeeAverages) GetTokenTransfer() MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer {
	if o == nil {
		var ret MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer
		return ret
	}

	return o.TokenTransfer
}

// GetTokenTransferOk returns a tuple with the TokenTransfer field value
// and a boolean to check if the value has been set.
func (o *MempoolTransactionStatsResponseTxSimpleFeeAverages) GetTokenTransferOk() (*MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenTransfer, true
}

// SetTokenTransfer sets field value
func (o *MempoolTransactionStatsResponseTxSimpleFeeAverages) SetTokenTransfer(v MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer) {
	o.TokenTransfer = v
}

// GetSmartContract returns the SmartContract field value
func (o *MempoolTransactionStatsResponseTxSimpleFeeAverages) GetSmartContract() MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer {
	if o == nil {
		var ret MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer
		return ret
	}

	return o.SmartContract
}

// GetSmartContractOk returns a tuple with the SmartContract field value
// and a boolean to check if the value has been set.
func (o *MempoolTransactionStatsResponseTxSimpleFeeAverages) GetSmartContractOk() (*MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SmartContract, true
}

// SetSmartContract sets field value
func (o *MempoolTransactionStatsResponseTxSimpleFeeAverages) SetSmartContract(v MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer) {
	o.SmartContract = v
}

// GetContractCall returns the ContractCall field value
func (o *MempoolTransactionStatsResponseTxSimpleFeeAverages) GetContractCall() MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer {
	if o == nil {
		var ret MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer
		return ret
	}

	return o.ContractCall
}

// GetContractCallOk returns a tuple with the ContractCall field value
// and a boolean to check if the value has been set.
func (o *MempoolTransactionStatsResponseTxSimpleFeeAverages) GetContractCallOk() (*MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractCall, true
}

// SetContractCall sets field value
func (o *MempoolTransactionStatsResponseTxSimpleFeeAverages) SetContractCall(v MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer) {
	o.ContractCall = v
}

// GetPoisonMicroblock returns the PoisonMicroblock field value
func (o *MempoolTransactionStatsResponseTxSimpleFeeAverages) GetPoisonMicroblock() MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer {
	if o == nil {
		var ret MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer
		return ret
	}

	return o.PoisonMicroblock
}

// GetPoisonMicroblockOk returns a tuple with the PoisonMicroblock field value
// and a boolean to check if the value has been set.
func (o *MempoolTransactionStatsResponseTxSimpleFeeAverages) GetPoisonMicroblockOk() (*MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PoisonMicroblock, true
}

// SetPoisonMicroblock sets field value
func (o *MempoolTransactionStatsResponseTxSimpleFeeAverages) SetPoisonMicroblock(v MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer) {
	o.PoisonMicroblock = v
}

func (o MempoolTransactionStatsResponseTxSimpleFeeAverages) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MempoolTransactionStatsResponseTxSimpleFeeAverages) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["token_transfer"] = o.TokenTransfer
	toSerialize["smart_contract"] = o.SmartContract
	toSerialize["contract_call"] = o.ContractCall
	toSerialize["poison_microblock"] = o.PoisonMicroblock
	return toSerialize, nil
}

func (o *MempoolTransactionStatsResponseTxSimpleFeeAverages) UnmarshalJSON(data []byte) (err error) {
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

	varMempoolTransactionStatsResponseTxSimpleFeeAverages := _MempoolTransactionStatsResponseTxSimpleFeeAverages{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMempoolTransactionStatsResponseTxSimpleFeeAverages)

	if err != nil {
		return err
	}

	*o = MempoolTransactionStatsResponseTxSimpleFeeAverages(varMempoolTransactionStatsResponseTxSimpleFeeAverages)

	return err
}

type NullableMempoolTransactionStatsResponseTxSimpleFeeAverages struct {
	value *MempoolTransactionStatsResponseTxSimpleFeeAverages
	isSet bool
}

func (v NullableMempoolTransactionStatsResponseTxSimpleFeeAverages) Get() *MempoolTransactionStatsResponseTxSimpleFeeAverages {
	return v.value
}

func (v *NullableMempoolTransactionStatsResponseTxSimpleFeeAverages) Set(val *MempoolTransactionStatsResponseTxSimpleFeeAverages) {
	v.value = val
	v.isSet = true
}

func (v NullableMempoolTransactionStatsResponseTxSimpleFeeAverages) IsSet() bool {
	return v.isSet
}

func (v *NullableMempoolTransactionStatsResponseTxSimpleFeeAverages) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMempoolTransactionStatsResponseTxSimpleFeeAverages(val *MempoolTransactionStatsResponseTxSimpleFeeAverages) *NullableMempoolTransactionStatsResponseTxSimpleFeeAverages {
	return &NullableMempoolTransactionStatsResponseTxSimpleFeeAverages{value: val, isSet: true}
}

func (v NullableMempoolTransactionStatsResponseTxSimpleFeeAverages) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMempoolTransactionStatsResponseTxSimpleFeeAverages) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
