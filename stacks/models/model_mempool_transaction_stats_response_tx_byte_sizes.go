package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the MempoolTransactionStatsResponseTxByteSizes type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &MempoolTransactionStatsResponseTxByteSizes{}

// MempoolTransactionStatsResponseTxByteSizes The average byte size of transactions in the mempool, broken down by transaction type.
type MempoolTransactionStatsResponseTxByteSizes struct {
	TokenTransfer    MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer `json:"token_transfer"`
	SmartContract    MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer `json:"smart_contract"`
	ContractCall     MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer `json:"contract_call"`
	PoisonMicroblock MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer `json:"poison_microblock"`
}

type _MempoolTransactionStatsResponseTxByteSizes MempoolTransactionStatsResponseTxByteSizes

// NewMempoolTransactionStatsResponseTxByteSizes instantiates a new MempoolTransactionStatsResponseTxByteSizes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMempoolTransactionStatsResponseTxByteSizes(tokenTransfer MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer, smartContract MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer, contractCall MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer, poisonMicroblock MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer) *MempoolTransactionStatsResponseTxByteSizes {
	this := MempoolTransactionStatsResponseTxByteSizes{}
	this.TokenTransfer = tokenTransfer
	this.SmartContract = smartContract
	this.ContractCall = contractCall
	this.PoisonMicroblock = poisonMicroblock
	return &this
}

// NewMempoolTransactionStatsResponseTxByteSizesWithDefaults instantiates a new MempoolTransactionStatsResponseTxByteSizes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMempoolTransactionStatsResponseTxByteSizesWithDefaults() *MempoolTransactionStatsResponseTxByteSizes {
	this := MempoolTransactionStatsResponseTxByteSizes{}
	return &this
}

// GetTokenTransfer returns the TokenTransfer field value
func (o *MempoolTransactionStatsResponseTxByteSizes) GetTokenTransfer() MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer {
	if o == nil {
		var ret MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer
		return ret
	}

	return o.TokenTransfer
}

// GetTokenTransferOk returns a tuple with the TokenTransfer field value
// and a boolean to check if the value has been set.
func (o *MempoolTransactionStatsResponseTxByteSizes) GetTokenTransferOk() (*MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenTransfer, true
}

// SetTokenTransfer sets field value
func (o *MempoolTransactionStatsResponseTxByteSizes) SetTokenTransfer(v MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer) {
	o.TokenTransfer = v
}

// GetSmartContract returns the SmartContract field value
func (o *MempoolTransactionStatsResponseTxByteSizes) GetSmartContract() MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer {
	if o == nil {
		var ret MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer
		return ret
	}

	return o.SmartContract
}

// GetSmartContractOk returns a tuple with the SmartContract field value
// and a boolean to check if the value has been set.
func (o *MempoolTransactionStatsResponseTxByteSizes) GetSmartContractOk() (*MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SmartContract, true
}

// SetSmartContract sets field value
func (o *MempoolTransactionStatsResponseTxByteSizes) SetSmartContract(v MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer) {
	o.SmartContract = v
}

// GetContractCall returns the ContractCall field value
func (o *MempoolTransactionStatsResponseTxByteSizes) GetContractCall() MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer {
	if o == nil {
		var ret MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer
		return ret
	}

	return o.ContractCall
}

// GetContractCallOk returns a tuple with the ContractCall field value
// and a boolean to check if the value has been set.
func (o *MempoolTransactionStatsResponseTxByteSizes) GetContractCallOk() (*MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractCall, true
}

// SetContractCall sets field value
func (o *MempoolTransactionStatsResponseTxByteSizes) SetContractCall(v MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer) {
	o.ContractCall = v
}

// GetPoisonMicroblock returns the PoisonMicroblock field value
func (o *MempoolTransactionStatsResponseTxByteSizes) GetPoisonMicroblock() MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer {
	if o == nil {
		var ret MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer
		return ret
	}

	return o.PoisonMicroblock
}

// GetPoisonMicroblockOk returns a tuple with the PoisonMicroblock field value
// and a boolean to check if the value has been set.
func (o *MempoolTransactionStatsResponseTxByteSizes) GetPoisonMicroblockOk() (*MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PoisonMicroblock, true
}

// SetPoisonMicroblock sets field value
func (o *MempoolTransactionStatsResponseTxByteSizes) SetPoisonMicroblock(v MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer) {
	o.PoisonMicroblock = v
}

func (o MempoolTransactionStatsResponseTxByteSizes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MempoolTransactionStatsResponseTxByteSizes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["token_transfer"] = o.TokenTransfer
	toSerialize["smart_contract"] = o.SmartContract
	toSerialize["contract_call"] = o.ContractCall
	toSerialize["poison_microblock"] = o.PoisonMicroblock
	return toSerialize, nil
}

func (o *MempoolTransactionStatsResponseTxByteSizes) UnmarshalJSON(data []byte) (err error) {
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

	varMempoolTransactionStatsResponseTxByteSizes := _MempoolTransactionStatsResponseTxByteSizes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMempoolTransactionStatsResponseTxByteSizes)

	if err != nil {
		return err
	}

	*o = MempoolTransactionStatsResponseTxByteSizes(varMempoolTransactionStatsResponseTxByteSizes)

	return err
}

type NullableMempoolTransactionStatsResponseTxByteSizes struct {
	value *MempoolTransactionStatsResponseTxByteSizes
	isSet bool
}

func (v NullableMempoolTransactionStatsResponseTxByteSizes) Get() *MempoolTransactionStatsResponseTxByteSizes {
	return v.value
}

func (v *NullableMempoolTransactionStatsResponseTxByteSizes) Set(val *MempoolTransactionStatsResponseTxByteSizes) {
	v.value = val
	v.isSet = true
}

func (v NullableMempoolTransactionStatsResponseTxByteSizes) IsSet() bool {
	return v.isSet
}

func (v *NullableMempoolTransactionStatsResponseTxByteSizes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMempoolTransactionStatsResponseTxByteSizes(val *MempoolTransactionStatsResponseTxByteSizes) *NullableMempoolTransactionStatsResponseTxByteSizes {
	return &NullableMempoolTransactionStatsResponseTxByteSizes{value: val, isSet: true}
}

func (v NullableMempoolTransactionStatsResponseTxByteSizes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMempoolTransactionStatsResponseTxByteSizes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
