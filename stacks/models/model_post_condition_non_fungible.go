package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the PostConditionNonFungible type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PostConditionNonFungible{}

// PostConditionNonFungible struct for PostConditionNonFungible
type PostConditionNonFungible struct {
	Principal PostConditionPrincipal `json:"principal"`
	// A non-fungible condition code encodes a statement being made about a non-fungible token, with respect to whether or not the particular non-fungible token is owned by the account.
	ConditionCode string                                  `json:"condition_code"`
	Type          string                                  `json:"type"`
	AssetValue    PostConditionNonFungibleAllOfAssetValue `json:"asset_value"`
	Asset         PostConditionFungibleAllOfAsset         `json:"asset"`
}

type _PostConditionNonFungible PostConditionNonFungible

// NewPostConditionNonFungible instantiates a new PostConditionNonFungible object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostConditionNonFungible(principal PostConditionPrincipal, conditionCode string, type_ string, assetValue PostConditionNonFungibleAllOfAssetValue, asset PostConditionFungibleAllOfAsset) *PostConditionNonFungible {
	this := PostConditionNonFungible{}
	this.Principal = principal
	this.ConditionCode = conditionCode
	this.Type = type_
	this.AssetValue = assetValue
	this.Asset = asset
	return &this
}

// NewPostConditionNonFungibleWithDefaults instantiates a new PostConditionNonFungible object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostConditionNonFungibleWithDefaults() *PostConditionNonFungible {
	this := PostConditionNonFungible{}
	return &this
}

// GetPrincipal returns the Principal field value
func (o *PostConditionNonFungible) GetPrincipal() PostConditionPrincipal {
	if o == nil {
		var ret PostConditionPrincipal
		return ret
	}

	return o.Principal
}

// GetPrincipalOk returns a tuple with the Principal field value
// and a boolean to check if the value has been set.
func (o *PostConditionNonFungible) GetPrincipalOk() (*PostConditionPrincipal, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Principal, true
}

// SetPrincipal sets field value
func (o *PostConditionNonFungible) SetPrincipal(v PostConditionPrincipal) {
	o.Principal = v
}

// GetConditionCode returns the ConditionCode field value
func (o *PostConditionNonFungible) GetConditionCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConditionCode
}

// GetConditionCodeOk returns a tuple with the ConditionCode field value
// and a boolean to check if the value has been set.
func (o *PostConditionNonFungible) GetConditionCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConditionCode, true
}

// SetConditionCode sets field value
func (o *PostConditionNonFungible) SetConditionCode(v string) {
	o.ConditionCode = v
}

// GetType returns the Type field value
func (o *PostConditionNonFungible) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PostConditionNonFungible) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PostConditionNonFungible) SetType(v string) {
	o.Type = v
}

// GetAssetValue returns the AssetValue field value
func (o *PostConditionNonFungible) GetAssetValue() PostConditionNonFungibleAllOfAssetValue {
	if o == nil {
		var ret PostConditionNonFungibleAllOfAssetValue
		return ret
	}

	return o.AssetValue
}

// GetAssetValueOk returns a tuple with the AssetValue field value
// and a boolean to check if the value has been set.
func (o *PostConditionNonFungible) GetAssetValueOk() (*PostConditionNonFungibleAllOfAssetValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetValue, true
}

// SetAssetValue sets field value
func (o *PostConditionNonFungible) SetAssetValue(v PostConditionNonFungibleAllOfAssetValue) {
	o.AssetValue = v
}

// GetAsset returns the Asset field value
func (o *PostConditionNonFungible) GetAsset() PostConditionFungibleAllOfAsset {
	if o == nil {
		var ret PostConditionFungibleAllOfAsset
		return ret
	}

	return o.Asset
}

// GetAssetOk returns a tuple with the Asset field value
// and a boolean to check if the value has been set.
func (o *PostConditionNonFungible) GetAssetOk() (*PostConditionFungibleAllOfAsset, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Asset, true
}

// SetAsset sets field value
func (o *PostConditionNonFungible) SetAsset(v PostConditionFungibleAllOfAsset) {
	o.Asset = v
}

func (o PostConditionNonFungible) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostConditionNonFungible) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["principal"] = o.Principal
	toSerialize["condition_code"] = o.ConditionCode
	toSerialize["type"] = o.Type
	toSerialize["asset_value"] = o.AssetValue
	toSerialize["asset"] = o.Asset
	return toSerialize, nil
}

func (o *PostConditionNonFungible) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"principal",
		"condition_code",
		"type",
		"asset_value",
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

	varPostConditionNonFungible := _PostConditionNonFungible{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostConditionNonFungible)

	if err != nil {
		return err
	}

	*o = PostConditionNonFungible(varPostConditionNonFungible)

	return err
}

type NullablePostConditionNonFungible struct {
	value *PostConditionNonFungible
	isSet bool
}

func (v NullablePostConditionNonFungible) Get() *PostConditionNonFungible {
	return v.value
}

func (v *NullablePostConditionNonFungible) Set(val *PostConditionNonFungible) {
	v.value = val
	v.isSet = true
}

func (v NullablePostConditionNonFungible) IsSet() bool {
	return v.isSet
}

func (v *NullablePostConditionNonFungible) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostConditionNonFungible(val *PostConditionNonFungible) *NullablePostConditionNonFungible {
	return &NullablePostConditionNonFungible{value: val, isSet: true}
}

func (v NullablePostConditionNonFungible) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostConditionNonFungible) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
