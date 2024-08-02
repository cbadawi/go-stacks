package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the ContractInterfaceResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ContractInterfaceResponse{}

// ContractInterfaceResponse GET request to get contract interface
type ContractInterfaceResponse struct {
	// List of defined methods
	Functions []map[string]interface{} `json:"functions"`
	// List of defined variables
	Variables []map[string]interface{} `json:"variables"`
	// List of defined data-maps
	Maps []map[string]interface{} `json:"maps"`
	// List of fungible tokens in the contract
	FungibleTokens []map[string]interface{} `json:"fungible_tokens"`
	// List of non-fungible tokens in the contract
	NonFungibleTokens []map[string]interface{} `json:"non_fungible_tokens"`
}

type _ContractInterfaceResponse ContractInterfaceResponse

// NewContractInterfaceResponse instantiates a new ContractInterfaceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContractInterfaceResponse(functions []map[string]interface{}, variables []map[string]interface{}, maps []map[string]interface{}, fungibleTokens []map[string]interface{}, nonFungibleTokens []map[string]interface{}) *ContractInterfaceResponse {
	this := ContractInterfaceResponse{}
	this.Functions = functions
	this.Variables = variables
	this.Maps = maps
	this.FungibleTokens = fungibleTokens
	this.NonFungibleTokens = nonFungibleTokens
	return &this
}

// NewContractInterfaceResponseWithDefaults instantiates a new ContractInterfaceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractInterfaceResponseWithDefaults() *ContractInterfaceResponse {
	this := ContractInterfaceResponse{}
	return &this
}

// GetFunctions returns the Functions field value
func (o *ContractInterfaceResponse) GetFunctions() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.Functions
}

// GetFunctionsOk returns a tuple with the Functions field value
// and a boolean to check if the value has been set.
func (o *ContractInterfaceResponse) GetFunctionsOk() ([]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Functions, true
}

// SetFunctions sets field value
func (o *ContractInterfaceResponse) SetFunctions(v []map[string]interface{}) {
	o.Functions = v
}

// GetVariables returns the Variables field value
func (o *ContractInterfaceResponse) GetVariables() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.Variables
}

// GetVariablesOk returns a tuple with the Variables field value
// and a boolean to check if the value has been set.
func (o *ContractInterfaceResponse) GetVariablesOk() ([]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Variables, true
}

// SetVariables sets field value
func (o *ContractInterfaceResponse) SetVariables(v []map[string]interface{}) {
	o.Variables = v
}

// GetMaps returns the Maps field value
func (o *ContractInterfaceResponse) GetMaps() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.Maps
}

// GetMapsOk returns a tuple with the Maps field value
// and a boolean to check if the value has been set.
func (o *ContractInterfaceResponse) GetMapsOk() ([]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Maps, true
}

// SetMaps sets field value
func (o *ContractInterfaceResponse) SetMaps(v []map[string]interface{}) {
	o.Maps = v
}

// GetFungibleTokens returns the FungibleTokens field value
func (o *ContractInterfaceResponse) GetFungibleTokens() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.FungibleTokens
}

// GetFungibleTokensOk returns a tuple with the FungibleTokens field value
// and a boolean to check if the value has been set.
func (o *ContractInterfaceResponse) GetFungibleTokensOk() ([]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.FungibleTokens, true
}

// SetFungibleTokens sets field value
func (o *ContractInterfaceResponse) SetFungibleTokens(v []map[string]interface{}) {
	o.FungibleTokens = v
}

// GetNonFungibleTokens returns the NonFungibleTokens field value
func (o *ContractInterfaceResponse) GetNonFungibleTokens() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.NonFungibleTokens
}

// GetNonFungibleTokensOk returns a tuple with the NonFungibleTokens field value
// and a boolean to check if the value has been set.
func (o *ContractInterfaceResponse) GetNonFungibleTokensOk() ([]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.NonFungibleTokens, true
}

// SetNonFungibleTokens sets field value
func (o *ContractInterfaceResponse) SetNonFungibleTokens(v []map[string]interface{}) {
	o.NonFungibleTokens = v
}

func (o ContractInterfaceResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContractInterfaceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["functions"] = o.Functions
	toSerialize["variables"] = o.Variables
	toSerialize["maps"] = o.Maps
	toSerialize["fungible_tokens"] = o.FungibleTokens
	toSerialize["non_fungible_tokens"] = o.NonFungibleTokens
	return toSerialize, nil
}

func (o *ContractInterfaceResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"functions",
		"variables",
		"maps",
		"fungible_tokens",
		"non_fungible_tokens",
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

	varContractInterfaceResponse := _ContractInterfaceResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varContractInterfaceResponse)

	if err != nil {
		return err
	}

	*o = ContractInterfaceResponse(varContractInterfaceResponse)

	return err
}

type NullableContractInterfaceResponse struct {
	value *ContractInterfaceResponse
	isSet bool
}

func (v NullableContractInterfaceResponse) Get() *ContractInterfaceResponse {
	return v.value
}

func (v *NullableContractInterfaceResponse) Set(val *ContractInterfaceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableContractInterfaceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableContractInterfaceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractInterfaceResponse(val *ContractInterfaceResponse) *NullableContractInterfaceResponse {
	return &NullableContractInterfaceResponse{value: val, isSet: true}
}

func (v NullableContractInterfaceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractInterfaceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
