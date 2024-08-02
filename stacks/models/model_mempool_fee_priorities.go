package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the MempoolFeePriorities type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &MempoolFeePriorities{}

// MempoolFeePriorities GET request that returns fee priorities from mempool transactions
type MempoolFeePriorities struct {
	All           MempoolFeePrioritiesAll  `json:"all"`
	TokenTransfer *MempoolFeePrioritiesAll `json:"token_transfer,omitempty"`
	SmartContract *MempoolFeePrioritiesAll `json:"smart_contract,omitempty"`
	ContractCall  *MempoolFeePrioritiesAll `json:"contract_call,omitempty"`
}

type _MempoolFeePriorities MempoolFeePriorities

// NewMempoolFeePriorities instantiates a new MempoolFeePriorities object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMempoolFeePriorities(all MempoolFeePrioritiesAll) *MempoolFeePriorities {
	this := MempoolFeePriorities{}
	this.All = all
	return &this
}

// NewMempoolFeePrioritiesWithDefaults instantiates a new MempoolFeePriorities object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMempoolFeePrioritiesWithDefaults() *MempoolFeePriorities {
	this := MempoolFeePriorities{}
	return &this
}

// GetAll returns the All field value
func (o *MempoolFeePriorities) GetAll() MempoolFeePrioritiesAll {
	if o == nil {
		var ret MempoolFeePrioritiesAll
		return ret
	}

	return o.All
}

// GetAllOk returns a tuple with the All field value
// and a boolean to check if the value has been set.
func (o *MempoolFeePriorities) GetAllOk() (*MempoolFeePrioritiesAll, bool) {
	if o == nil {
		return nil, false
	}
	return &o.All, true
}

// SetAll sets field value
func (o *MempoolFeePriorities) SetAll(v MempoolFeePrioritiesAll) {
	o.All = v
}

// GetTokenTransfer returns the TokenTransfer field value if set, zero value otherwise.
func (o *MempoolFeePriorities) GetTokenTransfer() MempoolFeePrioritiesAll {
	if o == nil || utils.IsNil(o.TokenTransfer) {
		var ret MempoolFeePrioritiesAll
		return ret
	}
	return *o.TokenTransfer
}

// GetTokenTransferOk returns a tuple with the TokenTransfer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MempoolFeePriorities) GetTokenTransferOk() (*MempoolFeePrioritiesAll, bool) {
	if o == nil || utils.IsNil(o.TokenTransfer) {
		return nil, false
	}
	return o.TokenTransfer, true
}

// HasTokenTransfer returns a boolean if a field has been set.
func (o *MempoolFeePriorities) HasTokenTransfer() bool {
	if o != nil && !utils.IsNil(o.TokenTransfer) {
		return true
	}

	return false
}

// SetTokenTransfer gets a reference to the given MempoolFeePrioritiesAll and assigns it to the TokenTransfer field.
func (o *MempoolFeePriorities) SetTokenTransfer(v MempoolFeePrioritiesAll) {
	o.TokenTransfer = &v
}

// GetSmartContract returns the SmartContract field value if set, zero value otherwise.
func (o *MempoolFeePriorities) GetSmartContract() MempoolFeePrioritiesAll {
	if o == nil || utils.IsNil(o.SmartContract) {
		var ret MempoolFeePrioritiesAll
		return ret
	}
	return *o.SmartContract
}

// GetSmartContractOk returns a tuple with the SmartContract field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MempoolFeePriorities) GetSmartContractOk() (*MempoolFeePrioritiesAll, bool) {
	if o == nil || utils.IsNil(o.SmartContract) {
		return nil, false
	}
	return o.SmartContract, true
}

// HasSmartContract returns a boolean if a field has been set.
func (o *MempoolFeePriorities) HasSmartContract() bool {
	if o != nil && !utils.IsNil(o.SmartContract) {
		return true
	}

	return false
}

// SetSmartContract gets a reference to the given MempoolFeePrioritiesAll and assigns it to the SmartContract field.
func (o *MempoolFeePriorities) SetSmartContract(v MempoolFeePrioritiesAll) {
	o.SmartContract = &v
}

// GetContractCall returns the ContractCall field value if set, zero value otherwise.
func (o *MempoolFeePriorities) GetContractCall() MempoolFeePrioritiesAll {
	if o == nil || utils.IsNil(o.ContractCall) {
		var ret MempoolFeePrioritiesAll
		return ret
	}
	return *o.ContractCall
}

// GetContractCallOk returns a tuple with the ContractCall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MempoolFeePriorities) GetContractCallOk() (*MempoolFeePrioritiesAll, bool) {
	if o == nil || utils.IsNil(o.ContractCall) {
		return nil, false
	}
	return o.ContractCall, true
}

// HasContractCall returns a boolean if a field has been set.
func (o *MempoolFeePriorities) HasContractCall() bool {
	if o != nil && !utils.IsNil(o.ContractCall) {
		return true
	}

	return false
}

// SetContractCall gets a reference to the given MempoolFeePrioritiesAll and assigns it to the ContractCall field.
func (o *MempoolFeePriorities) SetContractCall(v MempoolFeePrioritiesAll) {
	o.ContractCall = &v
}

func (o MempoolFeePriorities) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MempoolFeePriorities) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["all"] = o.All
	if !utils.IsNil(o.TokenTransfer) {
		toSerialize["token_transfer"] = o.TokenTransfer
	}
	if !utils.IsNil(o.SmartContract) {
		toSerialize["smart_contract"] = o.SmartContract
	}
	if !utils.IsNil(o.ContractCall) {
		toSerialize["contract_call"] = o.ContractCall
	}
	return toSerialize, nil
}

func (o *MempoolFeePriorities) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"all",
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

	varMempoolFeePriorities := _MempoolFeePriorities{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMempoolFeePriorities)

	if err != nil {
		return err
	}

	*o = MempoolFeePriorities(varMempoolFeePriorities)

	return err
}

type NullableMempoolFeePriorities struct {
	value *MempoolFeePriorities
	isSet bool
}

func (v NullableMempoolFeePriorities) Get() *MempoolFeePriorities {
	return v.value
}

func (v *NullableMempoolFeePriorities) Set(val *MempoolFeePriorities) {
	v.value = val
	v.isSet = true
}

func (v NullableMempoolFeePriorities) IsSet() bool {
	return v.isSet
}

func (v *NullableMempoolFeePriorities) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMempoolFeePriorities(val *MempoolFeePriorities) *NullableMempoolFeePriorities {
	return &NullableMempoolFeePriorities{value: val, isSet: true}
}

func (v NullableMempoolFeePriorities) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMempoolFeePriorities) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
