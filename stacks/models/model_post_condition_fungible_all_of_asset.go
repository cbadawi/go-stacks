package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the PostConditionFungibleAllOfAsset type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PostConditionFungibleAllOfAsset{}

// PostConditionFungibleAllOfAsset struct for PostConditionFungibleAllOfAsset
type PostConditionFungibleAllOfAsset struct {
	AssetName       string `json:"asset_name"`
	ContractAddress string `json:"contract_address"`
	ContractName    string `json:"contract_name"`
}

type _PostConditionFungibleAllOfAsset PostConditionFungibleAllOfAsset

// NewPostConditionFungibleAllOfAsset instantiates a new PostConditionFungibleAllOfAsset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostConditionFungibleAllOfAsset(assetName string, contractAddress string, contractName string) *PostConditionFungibleAllOfAsset {
	this := PostConditionFungibleAllOfAsset{}
	this.AssetName = assetName
	this.ContractAddress = contractAddress
	this.ContractName = contractName
	return &this
}

// NewPostConditionFungibleAllOfAssetWithDefaults instantiates a new PostConditionFungibleAllOfAsset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostConditionFungibleAllOfAssetWithDefaults() *PostConditionFungibleAllOfAsset {
	this := PostConditionFungibleAllOfAsset{}
	return &this
}

// GetAssetName returns the AssetName field value
func (o *PostConditionFungibleAllOfAsset) GetAssetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetName
}

// GetAssetNameOk returns a tuple with the AssetName field value
// and a boolean to check if the value has been set.
func (o *PostConditionFungibleAllOfAsset) GetAssetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetName, true
}

// SetAssetName sets field value
func (o *PostConditionFungibleAllOfAsset) SetAssetName(v string) {
	o.AssetName = v
}

// GetContractAddress returns the ContractAddress field value
func (o *PostConditionFungibleAllOfAsset) GetContractAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContractAddress
}

// GetContractAddressOk returns a tuple with the ContractAddress field value
// and a boolean to check if the value has been set.
func (o *PostConditionFungibleAllOfAsset) GetContractAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractAddress, true
}

// SetContractAddress sets field value
func (o *PostConditionFungibleAllOfAsset) SetContractAddress(v string) {
	o.ContractAddress = v
}

// GetContractName returns the ContractName field value
func (o *PostConditionFungibleAllOfAsset) GetContractName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContractName
}

// GetContractNameOk returns a tuple with the ContractName field value
// and a boolean to check if the value has been set.
func (o *PostConditionFungibleAllOfAsset) GetContractNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractName, true
}

// SetContractName sets field value
func (o *PostConditionFungibleAllOfAsset) SetContractName(v string) {
	o.ContractName = v
}

func (o PostConditionFungibleAllOfAsset) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostConditionFungibleAllOfAsset) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["asset_name"] = o.AssetName
	toSerialize["contract_address"] = o.ContractAddress
	toSerialize["contract_name"] = o.ContractName
	return toSerialize, nil
}

func (o *PostConditionFungibleAllOfAsset) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"asset_name",
		"contract_address",
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

	varPostConditionFungibleAllOfAsset := _PostConditionFungibleAllOfAsset{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostConditionFungibleAllOfAsset)

	if err != nil {
		return err
	}

	*o = PostConditionFungibleAllOfAsset(varPostConditionFungibleAllOfAsset)

	return err
}

type NullablePostConditionFungibleAllOfAsset struct {
	value *PostConditionFungibleAllOfAsset
	isSet bool
}

func (v NullablePostConditionFungibleAllOfAsset) Get() *PostConditionFungibleAllOfAsset {
	return v.value
}

func (v *NullablePostConditionFungibleAllOfAsset) Set(val *PostConditionFungibleAllOfAsset) {
	v.value = val
	v.isSet = true
}

func (v NullablePostConditionFungibleAllOfAsset) IsSet() bool {
	return v.isSet
}

func (v *NullablePostConditionFungibleAllOfAsset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostConditionFungibleAllOfAsset(val *PostConditionFungibleAllOfAsset) *NullablePostConditionFungibleAllOfAsset {
	return &NullablePostConditionFungibleAllOfAsset{value: val, isSet: true}
}

func (v NullablePostConditionFungibleAllOfAsset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostConditionFungibleAllOfAsset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
