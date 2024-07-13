package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the MempoolTransactionStatsResponseTxAges type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &MempoolTransactionStatsResponseTxAges{}

// MempoolTransactionStatsResponseTxAges The average time (in blocks) that transactions have lived in the mempool. The start block height is simply the current chain-tip of when the attached Stacks node receives the transaction. This timing can be different across Stacks nodes / API instances due to propagation timing differences in the p2p network.
type MempoolTransactionStatsResponseTxAges struct {
	TokenTransfer    MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer `json:"token_transfer"`
	SmartContract    MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer `json:"smart_contract"`
	ContractCall     MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer `json:"contract_call"`
	PoisonMicroblock MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer `json:"poison_microblock"`
}

type _MempoolTransactionStatsResponseTxAges MempoolTransactionStatsResponseTxAges

// NewMempoolTransactionStatsResponseTxAges instantiates a new MempoolTransactionStatsResponseTxAges object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMempoolTransactionStatsResponseTxAges(tokenTransfer MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer, smartContract MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer, contractCall MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer, poisonMicroblock MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer) *MempoolTransactionStatsResponseTxAges {
	this := MempoolTransactionStatsResponseTxAges{}
	this.TokenTransfer = tokenTransfer
	this.SmartContract = smartContract
	this.ContractCall = contractCall
	this.PoisonMicroblock = poisonMicroblock
	return &this
}

// NewMempoolTransactionStatsResponseTxAgesWithDefaults instantiates a new MempoolTransactionStatsResponseTxAges object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMempoolTransactionStatsResponseTxAgesWithDefaults() *MempoolTransactionStatsResponseTxAges {
	this := MempoolTransactionStatsResponseTxAges{}
	return &this
}

// GetTokenTransfer returns the TokenTransfer field value
func (o *MempoolTransactionStatsResponseTxAges) GetTokenTransfer() MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer {
	if o == nil {
		var ret MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer
		return ret
	}

	return o.TokenTransfer
}

// GetTokenTransferOk returns a tuple with the TokenTransfer field value
// and a boolean to check if the value has been set.
func (o *MempoolTransactionStatsResponseTxAges) GetTokenTransferOk() (*MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenTransfer, true
}

// SetTokenTransfer sets field value
func (o *MempoolTransactionStatsResponseTxAges) SetTokenTransfer(v MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer) {
	o.TokenTransfer = v
}

// GetSmartContract returns the SmartContract field value
func (o *MempoolTransactionStatsResponseTxAges) GetSmartContract() MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer {
	if o == nil {
		var ret MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer
		return ret
	}

	return o.SmartContract
}

// GetSmartContractOk returns a tuple with the SmartContract field value
// and a boolean to check if the value has been set.
func (o *MempoolTransactionStatsResponseTxAges) GetSmartContractOk() (*MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SmartContract, true
}

// SetSmartContract sets field value
func (o *MempoolTransactionStatsResponseTxAges) SetSmartContract(v MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer) {
	o.SmartContract = v
}

// GetContractCall returns the ContractCall field value
func (o *MempoolTransactionStatsResponseTxAges) GetContractCall() MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer {
	if o == nil {
		var ret MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer
		return ret
	}

	return o.ContractCall
}

// GetContractCallOk returns a tuple with the ContractCall field value
// and a boolean to check if the value has been set.
func (o *MempoolTransactionStatsResponseTxAges) GetContractCallOk() (*MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractCall, true
}

// SetContractCall sets field value
func (o *MempoolTransactionStatsResponseTxAges) SetContractCall(v MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer) {
	o.ContractCall = v
}

// GetPoisonMicroblock returns the PoisonMicroblock field value
func (o *MempoolTransactionStatsResponseTxAges) GetPoisonMicroblock() MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer {
	if o == nil {
		var ret MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer
		return ret
	}

	return o.PoisonMicroblock
}

// GetPoisonMicroblockOk returns a tuple with the PoisonMicroblock field value
// and a boolean to check if the value has been set.
func (o *MempoolTransactionStatsResponseTxAges) GetPoisonMicroblockOk() (*MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PoisonMicroblock, true
}

// SetPoisonMicroblock sets field value
func (o *MempoolTransactionStatsResponseTxAges) SetPoisonMicroblock(v MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer) {
	o.PoisonMicroblock = v
}

func (o MempoolTransactionStatsResponseTxAges) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MempoolTransactionStatsResponseTxAges) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["token_transfer"] = o.TokenTransfer
	toSerialize["smart_contract"] = o.SmartContract
	toSerialize["contract_call"] = o.ContractCall
	toSerialize["poison_microblock"] = o.PoisonMicroblock
	return toSerialize, nil
}

func (o *MempoolTransactionStatsResponseTxAges) UnmarshalJSON(data []byte) (err error) {
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

	varMempoolTransactionStatsResponseTxAges := _MempoolTransactionStatsResponseTxAges{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMempoolTransactionStatsResponseTxAges)

	if err != nil {
		return err
	}

	*o = MempoolTransactionStatsResponseTxAges(varMempoolTransactionStatsResponseTxAges)

	return err
}

type NullableMempoolTransactionStatsResponseTxAges struct {
	value *MempoolTransactionStatsResponseTxAges
	isSet bool
}

func (v NullableMempoolTransactionStatsResponseTxAges) Get() *MempoolTransactionStatsResponseTxAges {
	return v.value
}

func (v *NullableMempoolTransactionStatsResponseTxAges) Set(val *MempoolTransactionStatsResponseTxAges) {
	v.value = val
	v.isSet = true
}

func (v NullableMempoolTransactionStatsResponseTxAges) IsSet() bool {
	return v.isSet
}

func (v *NullableMempoolTransactionStatsResponseTxAges) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMempoolTransactionStatsResponseTxAges(val *MempoolTransactionStatsResponseTxAges) *NullableMempoolTransactionStatsResponseTxAges {
	return &NullableMempoolTransactionStatsResponseTxAges{value: val, isSet: true}
}

func (v NullableMempoolTransactionStatsResponseTxAges) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMempoolTransactionStatsResponseTxAges) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
