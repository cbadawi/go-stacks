package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the PostConditionPrincipalAnyOf2 type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PostConditionPrincipalAnyOf2{}

// PostConditionPrincipalAnyOf2 Principal Contract
type PostConditionPrincipalAnyOf2 struct {
	// String literal of type `PostConditionPrincipalType`
	TypeId       string `json:"type_id"`
	Address      string `json:"address"`
	ContractName string `json:"contract_name"`
}

type _PostConditionPrincipalAnyOf2 PostConditionPrincipalAnyOf2

// NewPostConditionPrincipalAnyOf2 instantiates a new PostConditionPrincipalAnyOf2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostConditionPrincipalAnyOf2(typeId string, address string, contractName string) *PostConditionPrincipalAnyOf2 {
	this := PostConditionPrincipalAnyOf2{}
	this.TypeId = typeId
	this.Address = address
	this.ContractName = contractName
	return &this
}

// NewPostConditionPrincipalAnyOf2WithDefaults instantiates a new PostConditionPrincipalAnyOf2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostConditionPrincipalAnyOf2WithDefaults() *PostConditionPrincipalAnyOf2 {
	this := PostConditionPrincipalAnyOf2{}
	return &this
}

// GetTypeId returns the TypeId field value
func (o *PostConditionPrincipalAnyOf2) GetTypeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TypeId
}

// GetTypeIdOk returns a tuple with the TypeId field value
// and a boolean to check if the value has been set.
func (o *PostConditionPrincipalAnyOf2) GetTypeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TypeId, true
}

// SetTypeId sets field value
func (o *PostConditionPrincipalAnyOf2) SetTypeId(v string) {
	o.TypeId = v
}

// GetAddress returns the Address field value
func (o *PostConditionPrincipalAnyOf2) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *PostConditionPrincipalAnyOf2) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *PostConditionPrincipalAnyOf2) SetAddress(v string) {
	o.Address = v
}

// GetContractName returns the ContractName field value
func (o *PostConditionPrincipalAnyOf2) GetContractName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContractName
}

// GetContractNameOk returns a tuple with the ContractName field value
// and a boolean to check if the value has been set.
func (o *PostConditionPrincipalAnyOf2) GetContractNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractName, true
}

// SetContractName sets field value
func (o *PostConditionPrincipalAnyOf2) SetContractName(v string) {
	o.ContractName = v
}

func (o PostConditionPrincipalAnyOf2) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostConditionPrincipalAnyOf2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type_id"] = o.TypeId
	toSerialize["address"] = o.Address
	toSerialize["contract_name"] = o.ContractName
	return toSerialize, nil
}

func (o *PostConditionPrincipalAnyOf2) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type_id",
		"address",
		"contract_name",
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

	varPostConditionPrincipalAnyOf2 := _PostConditionPrincipalAnyOf2{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostConditionPrincipalAnyOf2)

	if err != nil {
		return err
	}

	*o = PostConditionPrincipalAnyOf2(varPostConditionPrincipalAnyOf2)

	return err
}

type NullablePostConditionPrincipalAnyOf2 struct {
	value *PostConditionPrincipalAnyOf2
	isSet bool
}

func (v NullablePostConditionPrincipalAnyOf2) Get() *PostConditionPrincipalAnyOf2 {
	return v.value
}

func (v *NullablePostConditionPrincipalAnyOf2) Set(val *PostConditionPrincipalAnyOf2) {
	v.value = val
	v.isSet = true
}

func (v NullablePostConditionPrincipalAnyOf2) IsSet() bool {
	return v.isSet
}

func (v *NullablePostConditionPrincipalAnyOf2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostConditionPrincipalAnyOf2(val *PostConditionPrincipalAnyOf2) *NullablePostConditionPrincipalAnyOf2 {
	return &NullablePostConditionPrincipalAnyOf2{value: val, isSet: true}
}

func (v NullablePostConditionPrincipalAnyOf2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostConditionPrincipalAnyOf2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
