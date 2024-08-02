package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the PostConditionPrincipalAnyOf type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PostConditionPrincipalAnyOf{}

// PostConditionPrincipalAnyOf Principal Origin
type PostConditionPrincipalAnyOf struct {
	// String literal of type `PostConditionPrincipalType`
	TypeId string `json:"type_id"`
}

type _PostConditionPrincipalAnyOf PostConditionPrincipalAnyOf

// NewPostConditionPrincipalAnyOf instantiates a new PostConditionPrincipalAnyOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostConditionPrincipalAnyOf(typeId string) *PostConditionPrincipalAnyOf {
	this := PostConditionPrincipalAnyOf{}
	this.TypeId = typeId
	return &this
}

// NewPostConditionPrincipalAnyOfWithDefaults instantiates a new PostConditionPrincipalAnyOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostConditionPrincipalAnyOfWithDefaults() *PostConditionPrincipalAnyOf {
	this := PostConditionPrincipalAnyOf{}
	return &this
}

// GetTypeId returns the TypeId field value
func (o *PostConditionPrincipalAnyOf) GetTypeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TypeId
}

// GetTypeIdOk returns a tuple with the TypeId field value
// and a boolean to check if the value has been set.
func (o *PostConditionPrincipalAnyOf) GetTypeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TypeId, true
}

// SetTypeId sets field value
func (o *PostConditionPrincipalAnyOf) SetTypeId(v string) {
	o.TypeId = v
}

func (o PostConditionPrincipalAnyOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostConditionPrincipalAnyOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type_id"] = o.TypeId
	return toSerialize, nil
}

func (o *PostConditionPrincipalAnyOf) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type_id",
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

	varPostConditionPrincipalAnyOf := _PostConditionPrincipalAnyOf{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostConditionPrincipalAnyOf)

	if err != nil {
		return err
	}

	*o = PostConditionPrincipalAnyOf(varPostConditionPrincipalAnyOf)

	return err
}

type NullablePostConditionPrincipalAnyOf struct {
	value *PostConditionPrincipalAnyOf
	isSet bool
}

func (v NullablePostConditionPrincipalAnyOf) Get() *PostConditionPrincipalAnyOf {
	return v.value
}

func (v *NullablePostConditionPrincipalAnyOf) Set(val *PostConditionPrincipalAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePostConditionPrincipalAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePostConditionPrincipalAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostConditionPrincipalAnyOf(val *PostConditionPrincipalAnyOf) *NullablePostConditionPrincipalAnyOf {
	return &NullablePostConditionPrincipalAnyOf{value: val, isSet: true}
}

func (v NullablePostConditionPrincipalAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostConditionPrincipalAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
