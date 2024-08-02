package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the PostConditionPrincipalAnyOf1 type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PostConditionPrincipalAnyOf1{}

// PostConditionPrincipalAnyOf1 Principal Standard
type PostConditionPrincipalAnyOf1 struct {
	// String literal of type `PostConditionPrincipalType`
	TypeId  string `json:"type_id"`
	Address string `json:"address"`
}

type _PostConditionPrincipalAnyOf1 PostConditionPrincipalAnyOf1

// NewPostConditionPrincipalAnyOf1 instantiates a new PostConditionPrincipalAnyOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostConditionPrincipalAnyOf1(typeId string, address string) *PostConditionPrincipalAnyOf1 {
	this := PostConditionPrincipalAnyOf1{}
	this.TypeId = typeId
	this.Address = address
	return &this
}

// NewPostConditionPrincipalAnyOf1WithDefaults instantiates a new PostConditionPrincipalAnyOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostConditionPrincipalAnyOf1WithDefaults() *PostConditionPrincipalAnyOf1 {
	this := PostConditionPrincipalAnyOf1{}
	return &this
}

// GetTypeId returns the TypeId field value
func (o *PostConditionPrincipalAnyOf1) GetTypeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TypeId
}

// GetTypeIdOk returns a tuple with the TypeId field value
// and a boolean to check if the value has been set.
func (o *PostConditionPrincipalAnyOf1) GetTypeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TypeId, true
}

// SetTypeId sets field value
func (o *PostConditionPrincipalAnyOf1) SetTypeId(v string) {
	o.TypeId = v
}

// GetAddress returns the Address field value
func (o *PostConditionPrincipalAnyOf1) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *PostConditionPrincipalAnyOf1) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *PostConditionPrincipalAnyOf1) SetAddress(v string) {
	o.Address = v
}

func (o PostConditionPrincipalAnyOf1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostConditionPrincipalAnyOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type_id"] = o.TypeId
	toSerialize["address"] = o.Address
	return toSerialize, nil
}

func (o *PostConditionPrincipalAnyOf1) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type_id",
		"address",
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

	varPostConditionPrincipalAnyOf1 := _PostConditionPrincipalAnyOf1{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostConditionPrincipalAnyOf1)

	if err != nil {
		return err
	}

	*o = PostConditionPrincipalAnyOf1(varPostConditionPrincipalAnyOf1)

	return err
}

type NullablePostConditionPrincipalAnyOf1 struct {
	value *PostConditionPrincipalAnyOf1
	isSet bool
}

func (v NullablePostConditionPrincipalAnyOf1) Get() *PostConditionPrincipalAnyOf1 {
	return v.value
}

func (v *NullablePostConditionPrincipalAnyOf1) Set(val *PostConditionPrincipalAnyOf1) {
	v.value = val
	v.isSet = true
}

func (v NullablePostConditionPrincipalAnyOf1) IsSet() bool {
	return v.isSet
}

func (v *NullablePostConditionPrincipalAnyOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostConditionPrincipalAnyOf1(val *PostConditionPrincipalAnyOf1) *NullablePostConditionPrincipalAnyOf1 {
	return &NullablePostConditionPrincipalAnyOf1{value: val, isSet: true}
}

func (v NullablePostConditionPrincipalAnyOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostConditionPrincipalAnyOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
