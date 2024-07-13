package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the ContractCallTransactionMetadataContractCall type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ContractCallTransactionMetadataContractCall{}

// ContractCallTransactionMetadataContractCall struct for ContractCallTransactionMetadataContractCall
type ContractCallTransactionMetadataContractCall struct {
	// Contract identifier formatted as `<principaladdress>.<contract_name>`
	ContractId string `json:"contract_id"`
	// Name of the Clarity function to be invoked
	FunctionName string `json:"function_name"`
	// Function definition, including function name and type as well as parameter names and types
	FunctionSignature string `json:"function_signature"`
	// List of arguments used to invoke the function
	FunctionArgs []ContractCallTransactionMetadataContractCallFunctionArgsInner `json:"function_args,omitempty"`
}

type _ContractCallTransactionMetadataContractCall ContractCallTransactionMetadataContractCall

// NewContractCallTransactionMetadataContractCall instantiates a new ContractCallTransactionMetadataContractCall object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContractCallTransactionMetadataContractCall(contractId string, functionName string, functionSignature string) *ContractCallTransactionMetadataContractCall {
	this := ContractCallTransactionMetadataContractCall{}
	this.ContractId = contractId
	this.FunctionName = functionName
	this.FunctionSignature = functionSignature
	return &this
}

// NewContractCallTransactionMetadataContractCallWithDefaults instantiates a new ContractCallTransactionMetadataContractCall object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractCallTransactionMetadataContractCallWithDefaults() *ContractCallTransactionMetadataContractCall {
	this := ContractCallTransactionMetadataContractCall{}
	return &this
}

// GetContractId returns the ContractId field value
func (o *ContractCallTransactionMetadataContractCall) GetContractId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContractId
}

// GetContractIdOk returns a tuple with the ContractId field value
// and a boolean to check if the value has been set.
func (o *ContractCallTransactionMetadataContractCall) GetContractIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractId, true
}

// SetContractId sets field value
func (o *ContractCallTransactionMetadataContractCall) SetContractId(v string) {
	o.ContractId = v
}

// GetFunctionName returns the FunctionName field value
func (o *ContractCallTransactionMetadataContractCall) GetFunctionName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FunctionName
}

// GetFunctionNameOk returns a tuple with the FunctionName field value
// and a boolean to check if the value has been set.
func (o *ContractCallTransactionMetadataContractCall) GetFunctionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FunctionName, true
}

// SetFunctionName sets field value
func (o *ContractCallTransactionMetadataContractCall) SetFunctionName(v string) {
	o.FunctionName = v
}

// GetFunctionSignature returns the FunctionSignature field value
func (o *ContractCallTransactionMetadataContractCall) GetFunctionSignature() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FunctionSignature
}

// GetFunctionSignatureOk returns a tuple with the FunctionSignature field value
// and a boolean to check if the value has been set.
func (o *ContractCallTransactionMetadataContractCall) GetFunctionSignatureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FunctionSignature, true
}

// SetFunctionSignature sets field value
func (o *ContractCallTransactionMetadataContractCall) SetFunctionSignature(v string) {
	o.FunctionSignature = v
}

// GetFunctionArgs returns the FunctionArgs field value if set, zero value otherwise.
func (o *ContractCallTransactionMetadataContractCall) GetFunctionArgs() []ContractCallTransactionMetadataContractCallFunctionArgsInner {
	if o == nil || utils.IsNil(o.FunctionArgs) {
		var ret []ContractCallTransactionMetadataContractCallFunctionArgsInner
		return ret
	}
	return o.FunctionArgs
}

// GetFunctionArgsOk returns a tuple with the FunctionArgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractCallTransactionMetadataContractCall) GetFunctionArgsOk() ([]ContractCallTransactionMetadataContractCallFunctionArgsInner, bool) {
	if o == nil || utils.IsNil(o.FunctionArgs) {
		return nil, false
	}
	return o.FunctionArgs, true
}

// HasFunctionArgs returns a boolean if a field has been set.
func (o *ContractCallTransactionMetadataContractCall) HasFunctionArgs() bool {
	if o != nil && !utils.IsNil(o.FunctionArgs) {
		return true
	}

	return false
}

// SetFunctionArgs gets a reference to the given []ContractCallTransactionMetadataContractCallFunctionArgsInner and assigns it to the FunctionArgs field.
func (o *ContractCallTransactionMetadataContractCall) SetFunctionArgs(v []ContractCallTransactionMetadataContractCallFunctionArgsInner) {
	o.FunctionArgs = v
}

func (o ContractCallTransactionMetadataContractCall) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContractCallTransactionMetadataContractCall) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["contract_id"] = o.ContractId
	toSerialize["function_name"] = o.FunctionName
	toSerialize["function_signature"] = o.FunctionSignature
	if !utils.IsNil(o.FunctionArgs) {
		toSerialize["function_args"] = o.FunctionArgs
	}
	return toSerialize, nil
}

func (o *ContractCallTransactionMetadataContractCall) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"contract_id",
		"function_name",
		"function_signature",
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

	varContractCallTransactionMetadataContractCall := _ContractCallTransactionMetadataContractCall{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varContractCallTransactionMetadataContractCall)

	if err != nil {
		return err
	}

	*o = ContractCallTransactionMetadataContractCall(varContractCallTransactionMetadataContractCall)

	return err
}

type NullableContractCallTransactionMetadataContractCall struct {
	value *ContractCallTransactionMetadataContractCall
	isSet bool
}

func (v NullableContractCallTransactionMetadataContractCall) Get() *ContractCallTransactionMetadataContractCall {
	return v.value
}

func (v *NullableContractCallTransactionMetadataContractCall) Set(val *ContractCallTransactionMetadataContractCall) {
	v.value = val
	v.isSet = true
}

func (v NullableContractCallTransactionMetadataContractCall) IsSet() bool {
	return v.isSet
}

func (v *NullableContractCallTransactionMetadataContractCall) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractCallTransactionMetadataContractCall(val *ContractCallTransactionMetadataContractCall) *NullableContractCallTransactionMetadataContractCall {
	return &NullableContractCallTransactionMetadataContractCall{value: val, isSet: true}
}

func (v NullableContractCallTransactionMetadataContractCall) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractCallTransactionMetadataContractCall) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
