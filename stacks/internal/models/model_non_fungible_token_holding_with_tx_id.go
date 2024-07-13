package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the NonFungibleTokenHoldingWithTxId type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &NonFungibleTokenHoldingWithTxId{}

// NonFungibleTokenHoldingWithTxId Ownership of a Non-Fungible Token
type NonFungibleTokenHoldingWithTxId struct {
	AssetIdentifier string                               `json:"asset_identifier"`
	Value           NonFungibleTokenHoldingWithTxIdValue `json:"value"`
	BlockHeight     float32                              `json:"block_height"`
	TxId            string                               `json:"tx_id"`
}

type _NonFungibleTokenHoldingWithTxId NonFungibleTokenHoldingWithTxId

// NewNonFungibleTokenHoldingWithTxId instantiates a new NonFungibleTokenHoldingWithTxId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNonFungibleTokenHoldingWithTxId(assetIdentifier string, value NonFungibleTokenHoldingWithTxIdValue, blockHeight float32, txId string) *NonFungibleTokenHoldingWithTxId {
	this := NonFungibleTokenHoldingWithTxId{}
	this.AssetIdentifier = assetIdentifier
	this.Value = value
	this.BlockHeight = blockHeight
	this.TxId = txId
	return &this
}

// NewNonFungibleTokenHoldingWithTxIdWithDefaults instantiates a new NonFungibleTokenHoldingWithTxId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNonFungibleTokenHoldingWithTxIdWithDefaults() *NonFungibleTokenHoldingWithTxId {
	this := NonFungibleTokenHoldingWithTxId{}
	return &this
}

// GetAssetIdentifier returns the AssetIdentifier field value
func (o *NonFungibleTokenHoldingWithTxId) GetAssetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetIdentifier
}

// GetAssetIdentifierOk returns a tuple with the AssetIdentifier field value
// and a boolean to check if the value has been set.
func (o *NonFungibleTokenHoldingWithTxId) GetAssetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetIdentifier, true
}

// SetAssetIdentifier sets field value
func (o *NonFungibleTokenHoldingWithTxId) SetAssetIdentifier(v string) {
	o.AssetIdentifier = v
}

// GetValue returns the Value field value
func (o *NonFungibleTokenHoldingWithTxId) GetValue() NonFungibleTokenHoldingWithTxIdValue {
	if o == nil {
		var ret NonFungibleTokenHoldingWithTxIdValue
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *NonFungibleTokenHoldingWithTxId) GetValueOk() (*NonFungibleTokenHoldingWithTxIdValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *NonFungibleTokenHoldingWithTxId) SetValue(v NonFungibleTokenHoldingWithTxIdValue) {
	o.Value = v
}

// GetBlockHeight returns the BlockHeight field value
func (o *NonFungibleTokenHoldingWithTxId) GetBlockHeight() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.BlockHeight
}

// GetBlockHeightOk returns a tuple with the BlockHeight field value
// and a boolean to check if the value has been set.
func (o *NonFungibleTokenHoldingWithTxId) GetBlockHeightOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockHeight, true
}

// SetBlockHeight sets field value
func (o *NonFungibleTokenHoldingWithTxId) SetBlockHeight(v float32) {
	o.BlockHeight = v
}

// GetTxId returns the TxId field value
func (o *NonFungibleTokenHoldingWithTxId) GetTxId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TxId
}

// GetTxIdOk returns a tuple with the TxId field value
// and a boolean to check if the value has been set.
func (o *NonFungibleTokenHoldingWithTxId) GetTxIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxId, true
}

// SetTxId sets field value
func (o *NonFungibleTokenHoldingWithTxId) SetTxId(v string) {
	o.TxId = v
}

func (o NonFungibleTokenHoldingWithTxId) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NonFungibleTokenHoldingWithTxId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["asset_identifier"] = o.AssetIdentifier
	toSerialize["value"] = o.Value
	toSerialize["block_height"] = o.BlockHeight
	toSerialize["tx_id"] = o.TxId
	return toSerialize, nil
}

func (o *NonFungibleTokenHoldingWithTxId) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"asset_identifier",
		"value",
		"block_height",
		"tx_id",
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

	varNonFungibleTokenHoldingWithTxId := _NonFungibleTokenHoldingWithTxId{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNonFungibleTokenHoldingWithTxId)

	if err != nil {
		return err
	}

	*o = NonFungibleTokenHoldingWithTxId(varNonFungibleTokenHoldingWithTxId)

	return err
}

type NullableNonFungibleTokenHoldingWithTxId struct {
	value *NonFungibleTokenHoldingWithTxId
	isSet bool
}

func (v NullableNonFungibleTokenHoldingWithTxId) Get() *NonFungibleTokenHoldingWithTxId {
	return v.value
}

func (v *NullableNonFungibleTokenHoldingWithTxId) Set(val *NonFungibleTokenHoldingWithTxId) {
	v.value = val
	v.isSet = true
}

func (v NullableNonFungibleTokenHoldingWithTxId) IsSet() bool {
	return v.isSet
}

func (v *NullableNonFungibleTokenHoldingWithTxId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNonFungibleTokenHoldingWithTxId(val *NonFungibleTokenHoldingWithTxId) *NullableNonFungibleTokenHoldingWithTxId {
	return &NullableNonFungibleTokenHoldingWithTxId{value: val, isSet: true}
}

func (v NullableNonFungibleTokenHoldingWithTxId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNonFungibleTokenHoldingWithTxId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
