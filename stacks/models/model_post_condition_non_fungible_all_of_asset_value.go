package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the PostConditionNonFungibleAllOfAssetValue type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PostConditionNonFungibleAllOfAssetValue{}

// PostConditionNonFungibleAllOfAssetValue struct for PostConditionNonFungibleAllOfAssetValue
type PostConditionNonFungibleAllOfAssetValue struct {
	Hex  string `json:"hex"`
	Repr string `json:"repr"`
}

type _PostConditionNonFungibleAllOfAssetValue PostConditionNonFungibleAllOfAssetValue

// NewPostConditionNonFungibleAllOfAssetValue instantiates a new PostConditionNonFungibleAllOfAssetValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostConditionNonFungibleAllOfAssetValue(hex string, repr string) *PostConditionNonFungibleAllOfAssetValue {
	this := PostConditionNonFungibleAllOfAssetValue{}
	this.Hex = hex
	this.Repr = repr
	return &this
}

// NewPostConditionNonFungibleAllOfAssetValueWithDefaults instantiates a new PostConditionNonFungibleAllOfAssetValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostConditionNonFungibleAllOfAssetValueWithDefaults() *PostConditionNonFungibleAllOfAssetValue {
	this := PostConditionNonFungibleAllOfAssetValue{}
	return &this
}

// GetHex returns the Hex field value
func (o *PostConditionNonFungibleAllOfAssetValue) GetHex() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hex
}

// GetHexOk returns a tuple with the Hex field value
// and a boolean to check if the value has been set.
func (o *PostConditionNonFungibleAllOfAssetValue) GetHexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hex, true
}

// SetHex sets field value
func (o *PostConditionNonFungibleAllOfAssetValue) SetHex(v string) {
	o.Hex = v
}

// GetRepr returns the Repr field value
func (o *PostConditionNonFungibleAllOfAssetValue) GetRepr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Repr
}

// GetReprOk returns a tuple with the Repr field value
// and a boolean to check if the value has been set.
func (o *PostConditionNonFungibleAllOfAssetValue) GetReprOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Repr, true
}

// SetRepr sets field value
func (o *PostConditionNonFungibleAllOfAssetValue) SetRepr(v string) {
	o.Repr = v
}

func (o PostConditionNonFungibleAllOfAssetValue) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostConditionNonFungibleAllOfAssetValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["hex"] = o.Hex
	toSerialize["repr"] = o.Repr
	return toSerialize, nil
}

func (o *PostConditionNonFungibleAllOfAssetValue) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"hex",
		"repr",
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

	varPostConditionNonFungibleAllOfAssetValue := _PostConditionNonFungibleAllOfAssetValue{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostConditionNonFungibleAllOfAssetValue)

	if err != nil {
		return err
	}

	*o = PostConditionNonFungibleAllOfAssetValue(varPostConditionNonFungibleAllOfAssetValue)

	return err
}

type NullablePostConditionNonFungibleAllOfAssetValue struct {
	value *PostConditionNonFungibleAllOfAssetValue
	isSet bool
}

func (v NullablePostConditionNonFungibleAllOfAssetValue) Get() *PostConditionNonFungibleAllOfAssetValue {
	return v.value
}

func (v *NullablePostConditionNonFungibleAllOfAssetValue) Set(val *PostConditionNonFungibleAllOfAssetValue) {
	v.value = val
	v.isSet = true
}

func (v NullablePostConditionNonFungibleAllOfAssetValue) IsSet() bool {
	return v.isSet
}

func (v *NullablePostConditionNonFungibleAllOfAssetValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostConditionNonFungibleAllOfAssetValue(val *PostConditionNonFungibleAllOfAssetValue) *NullablePostConditionNonFungibleAllOfAssetValue {
	return &NullablePostConditionNonFungibleAllOfAssetValue{value: val, isSet: true}
}

func (v NullablePostConditionNonFungibleAllOfAssetValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostConditionNonFungibleAllOfAssetValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
