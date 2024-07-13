package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the SmartContractTransactionMetadataSmartContract type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &SmartContractTransactionMetadataSmartContract{}

// SmartContractTransactionMetadataSmartContract struct for SmartContractTransactionMetadataSmartContract
type SmartContractTransactionMetadataSmartContract struct {
	// The Clarity version of the contract, only specified for versioned contract transactions, otherwise null
	ClarityVersion utils.NullableFloat32 `json:"clarity_version,omitempty"`
	// Contract identifier formatted as `<principaladdress>.<contract_name>`
	ContractId string `json:"contract_id"`
	// Clarity code of the smart contract being deployed
	SourceCode string `json:"source_code"`
}

type _SmartContractTransactionMetadataSmartContract SmartContractTransactionMetadataSmartContract

// NewSmartContractTransactionMetadataSmartContract instantiates a new SmartContractTransactionMetadataSmartContract object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmartContractTransactionMetadataSmartContract(contractId string, sourceCode string) *SmartContractTransactionMetadataSmartContract {
	this := SmartContractTransactionMetadataSmartContract{}
	this.ContractId = contractId
	this.SourceCode = sourceCode
	return &this
}

// NewSmartContractTransactionMetadataSmartContractWithDefaults instantiates a new SmartContractTransactionMetadataSmartContract object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmartContractTransactionMetadataSmartContractWithDefaults() *SmartContractTransactionMetadataSmartContract {
	this := SmartContractTransactionMetadataSmartContract{}
	return &this
}

// GetClarityVersion returns the ClarityVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SmartContractTransactionMetadataSmartContract) GetClarityVersion() float32 {
	if o == nil || utils.IsNil(o.ClarityVersion.Get()) {
		var ret float32
		return ret
	}
	return *o.ClarityVersion.Get()
}

// GetClarityVersionOk returns a tuple with the ClarityVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SmartContractTransactionMetadataSmartContract) GetClarityVersionOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClarityVersion.Get(), o.ClarityVersion.IsSet()
}

// HasClarityVersion returns a boolean if a field has been set.
func (o *SmartContractTransactionMetadataSmartContract) HasClarityVersion() bool {
	if o != nil && o.ClarityVersion.IsSet() {
		return true
	}

	return false
}

// SetClarityVersion gets a reference to the given utils.NullableFloat32 and assigns it to the ClarityVersion field.
func (o *SmartContractTransactionMetadataSmartContract) SetClarityVersion(v float32) {
	o.ClarityVersion.Set(&v)
}

// SetClarityVersionNil sets the value for ClarityVersion to be an explicit nil
func (o *SmartContractTransactionMetadataSmartContract) SetClarityVersionNil() {
	o.ClarityVersion.Set(nil)
}

// UnsetClarityVersion ensures that no value is present for ClarityVersion, not even an explicit nil
func (o *SmartContractTransactionMetadataSmartContract) UnsetClarityVersion() {
	o.ClarityVersion.Unset()
}

// GetContractId returns the ContractId field value
func (o *SmartContractTransactionMetadataSmartContract) GetContractId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContractId
}

// GetContractIdOk returns a tuple with the ContractId field value
// and a boolean to check if the value has been set.
func (o *SmartContractTransactionMetadataSmartContract) GetContractIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractId, true
}

// SetContractId sets field value
func (o *SmartContractTransactionMetadataSmartContract) SetContractId(v string) {
	o.ContractId = v
}

// GetSourceCode returns the SourceCode field value
func (o *SmartContractTransactionMetadataSmartContract) GetSourceCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceCode
}

// GetSourceCodeOk returns a tuple with the SourceCode field value
// and a boolean to check if the value has been set.
func (o *SmartContractTransactionMetadataSmartContract) GetSourceCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceCode, true
}

// SetSourceCode sets field value
func (o *SmartContractTransactionMetadataSmartContract) SetSourceCode(v string) {
	o.SourceCode = v
}

func (o SmartContractTransactionMetadataSmartContract) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SmartContractTransactionMetadataSmartContract) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ClarityVersion.IsSet() {
		toSerialize["clarity_version"] = o.ClarityVersion.Get()
	}
	toSerialize["contract_id"] = o.ContractId
	toSerialize["source_code"] = o.SourceCode
	return toSerialize, nil
}

func (o *SmartContractTransactionMetadataSmartContract) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"contract_id",
		"source_code",
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

	varSmartContractTransactionMetadataSmartContract := _SmartContractTransactionMetadataSmartContract{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSmartContractTransactionMetadataSmartContract)

	if err != nil {
		return err
	}

	*o = SmartContractTransactionMetadataSmartContract(varSmartContractTransactionMetadataSmartContract)

	return err
}

type NullableSmartContractTransactionMetadataSmartContract struct {
	value *SmartContractTransactionMetadataSmartContract
	isSet bool
}

func (v NullableSmartContractTransactionMetadataSmartContract) Get() *SmartContractTransactionMetadataSmartContract {
	return v.value
}

func (v *NullableSmartContractTransactionMetadataSmartContract) Set(val *SmartContractTransactionMetadataSmartContract) {
	v.value = val
	v.isSet = true
}

func (v NullableSmartContractTransactionMetadataSmartContract) IsSet() bool {
	return v.isSet
}

func (v *NullableSmartContractTransactionMetadataSmartContract) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmartContractTransactionMetadataSmartContract(val *SmartContractTransactionMetadataSmartContract) *NullableSmartContractTransactionMetadataSmartContract {
	return &NullableSmartContractTransactionMetadataSmartContract{value: val, isSet: true}
}

func (v NullableSmartContractTransactionMetadataSmartContract) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmartContractTransactionMetadataSmartContract) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
