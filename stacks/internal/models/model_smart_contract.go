package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the SmartContract type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &SmartContract{}

// SmartContract A Smart Contract Detail
type SmartContract struct {
	TxId        string `json:"tx_id"`
	Canonical   bool   `json:"canonical"`
	ContractId  string `json:"contract_id"`
	BlockHeight int32  `json:"block_height"`
	SourceCode  string `json:"source_code"`
	Abi         string `json:"abi"`
}

type _SmartContract SmartContract

// NewSmartContract instantiates a new SmartContract object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmartContract(txId string, canonical bool, contractId string, blockHeight int32, sourceCode string, abi string) *SmartContract {
	this := SmartContract{}
	this.TxId = txId
	this.Canonical = canonical
	this.ContractId = contractId
	this.BlockHeight = blockHeight
	this.SourceCode = sourceCode
	this.Abi = abi
	return &this
}

// NewSmartContractWithDefaults instantiates a new SmartContract object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmartContractWithDefaults() *SmartContract {
	this := SmartContract{}
	return &this
}

// GetTxId returns the TxId field value
func (o *SmartContract) GetTxId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TxId
}

// GetTxIdOk returns a tuple with the TxId field value
// and a boolean to check if the value has been set.
func (o *SmartContract) GetTxIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxId, true
}

// SetTxId sets field value
func (o *SmartContract) SetTxId(v string) {
	o.TxId = v
}

// GetCanonical returns the Canonical field value
func (o *SmartContract) GetCanonical() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Canonical
}

// GetCanonicalOk returns a tuple with the Canonical field value
// and a boolean to check if the value has been set.
func (o *SmartContract) GetCanonicalOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Canonical, true
}

// SetCanonical sets field value
func (o *SmartContract) SetCanonical(v bool) {
	o.Canonical = v
}

// GetContractId returns the ContractId field value
func (o *SmartContract) GetContractId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContractId
}

// GetContractIdOk returns a tuple with the ContractId field value
// and a boolean to check if the value has been set.
func (o *SmartContract) GetContractIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractId, true
}

// SetContractId sets field value
func (o *SmartContract) SetContractId(v string) {
	o.ContractId = v
}

// GetBlockHeight returns the BlockHeight field value
func (o *SmartContract) GetBlockHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BlockHeight
}

// GetBlockHeightOk returns a tuple with the BlockHeight field value
// and a boolean to check if the value has been set.
func (o *SmartContract) GetBlockHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockHeight, true
}

// SetBlockHeight sets field value
func (o *SmartContract) SetBlockHeight(v int32) {
	o.BlockHeight = v
}

// GetSourceCode returns the SourceCode field value
func (o *SmartContract) GetSourceCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceCode
}

// GetSourceCodeOk returns a tuple with the SourceCode field value
// and a boolean to check if the value has been set.
func (o *SmartContract) GetSourceCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceCode, true
}

// SetSourceCode sets field value
func (o *SmartContract) SetSourceCode(v string) {
	o.SourceCode = v
}

// GetAbi returns the Abi field value
func (o *SmartContract) GetAbi() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Abi
}

// GetAbiOk returns a tuple with the Abi field value
// and a boolean to check if the value has been set.
func (o *SmartContract) GetAbiOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Abi, true
}

// SetAbi sets field value
func (o *SmartContract) SetAbi(v string) {
	o.Abi = v
}

func (o SmartContract) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SmartContract) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tx_id"] = o.TxId
	toSerialize["canonical"] = o.Canonical
	toSerialize["contract_id"] = o.ContractId
	toSerialize["block_height"] = o.BlockHeight
	toSerialize["source_code"] = o.SourceCode
	toSerialize["abi"] = o.Abi
	return toSerialize, nil
}

func (o *SmartContract) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tx_id",
		"canonical",
		"contract_id",
		"block_height",
		"source_code",
		"abi",
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

	varSmartContract := _SmartContract{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSmartContract)

	if err != nil {
		return err
	}

	*o = SmartContract(varSmartContract)

	return err
}

type NullableSmartContract struct {
	value *SmartContract
	isSet bool
}

func (v NullableSmartContract) Get() *SmartContract {
	return v.value
}

func (v *NullableSmartContract) Set(val *SmartContract) {
	v.value = val
	v.isSet = true
}

func (v NullableSmartContract) IsSet() bool {
	return v.isSet
}

func (v *NullableSmartContract) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmartContract(val *SmartContract) *NullableSmartContract {
	return &NullableSmartContract{value: val, isSet: true}
}

func (v NullableSmartContract) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmartContract) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
