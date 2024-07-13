package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the PostConditionFungible type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PostConditionFungible{}

// PostConditionFungible struct for PostConditionFungible
type PostConditionFungible struct {
	Principal PostConditionPrincipal `json:"principal"`
	// A fungible condition code encodes a statement being made for either STX or a fungible token, with respect to the originating account.
	ConditionCode string                          `json:"condition_code"`
	Type          string                          `json:"type"`
	Amount        string                          `json:"amount"`
	Asset         PostConditionFungibleAllOfAsset `json:"asset"`
}

type _PostConditionFungible PostConditionFungible

// NewPostConditionFungible instantiates a new PostConditionFungible object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostConditionFungible(principal PostConditionPrincipal, conditionCode string, type_ string, amount string, asset PostConditionFungibleAllOfAsset) *PostConditionFungible {
	this := PostConditionFungible{}
	this.Principal = principal
	this.ConditionCode = conditionCode
	this.Type = type_
	this.Amount = amount
	this.Asset = asset
	return &this
}

// NewPostConditionFungibleWithDefaults instantiates a new PostConditionFungible object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostConditionFungibleWithDefaults() *PostConditionFungible {
	this := PostConditionFungible{}
	return &this
}

// GetPrincipal returns the Principal field value
func (o *PostConditionFungible) GetPrincipal() PostConditionPrincipal {
	if o == nil {
		var ret PostConditionPrincipal
		return ret
	}

	return o.Principal
}

// GetPrincipalOk returns a tuple with the Principal field value
// and a boolean to check if the value has been set.
func (o *PostConditionFungible) GetPrincipalOk() (*PostConditionPrincipal, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Principal, true
}

// SetPrincipal sets field value
func (o *PostConditionFungible) SetPrincipal(v PostConditionPrincipal) {
	o.Principal = v
}

// GetConditionCode returns the ConditionCode field value
func (o *PostConditionFungible) GetConditionCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConditionCode
}

// GetConditionCodeOk returns a tuple with the ConditionCode field value
// and a boolean to check if the value has been set.
func (o *PostConditionFungible) GetConditionCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConditionCode, true
}

// SetConditionCode sets field value
func (o *PostConditionFungible) SetConditionCode(v string) {
	o.ConditionCode = v
}

// GetType returns the Type field value
func (o *PostConditionFungible) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PostConditionFungible) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PostConditionFungible) SetType(v string) {
	o.Type = v
}

// GetAmount returns the Amount field value
func (o *PostConditionFungible) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *PostConditionFungible) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *PostConditionFungible) SetAmount(v string) {
	o.Amount = v
}

// GetAsset returns the Asset field value
func (o *PostConditionFungible) GetAsset() PostConditionFungibleAllOfAsset {
	if o == nil {
		var ret PostConditionFungibleAllOfAsset
		return ret
	}

	return o.Asset
}

// GetAssetOk returns a tuple with the Asset field value
// and a boolean to check if the value has been set.
func (o *PostConditionFungible) GetAssetOk() (*PostConditionFungibleAllOfAsset, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Asset, true
}

// SetAsset sets field value
func (o *PostConditionFungible) SetAsset(v PostConditionFungibleAllOfAsset) {
	o.Asset = v
}

func (o PostConditionFungible) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostConditionFungible) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["principal"] = o.Principal
	toSerialize["condition_code"] = o.ConditionCode
	toSerialize["type"] = o.Type
	toSerialize["amount"] = o.Amount
	toSerialize["asset"] = o.Asset
	return toSerialize, nil
}

func (o *PostConditionFungible) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"principal",
		"condition_code",
		"type",
		"amount",
		"asset",
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

	varPostConditionFungible := _PostConditionFungible{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostConditionFungible)

	if err != nil {
		return err
	}

	*o = PostConditionFungible(varPostConditionFungible)

	return err
}

type NullablePostConditionFungible struct {
	value *PostConditionFungible
	isSet bool
}

func (v NullablePostConditionFungible) Get() *PostConditionFungible {
	return v.value
}

func (v *NullablePostConditionFungible) Set(val *PostConditionFungible) {
	v.value = val
	v.isSet = true
}

func (v NullablePostConditionFungible) IsSet() bool {
	return v.isSet
}

func (v *NullablePostConditionFungible) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostConditionFungible(val *PostConditionFungible) *NullablePostConditionFungible {
	return &NullablePostConditionFungible{value: val, isSet: true}
}

func (v NullablePostConditionFungible) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostConditionFungible) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
