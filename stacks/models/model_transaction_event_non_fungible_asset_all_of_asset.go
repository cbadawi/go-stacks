package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the TransactionEventNonFungibleAssetAllOfAsset type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &TransactionEventNonFungibleAssetAllOfAsset{}

// TransactionEventNonFungibleAssetAllOfAsset struct for TransactionEventNonFungibleAssetAllOfAsset
type TransactionEventNonFungibleAssetAllOfAsset struct {
	AssetEventType string                                  `json:"asset_event_type"`
	AssetId        string                                  `json:"asset_id"`
	Sender         string                                  `json:"sender"`
	Recipient      string                                  `json:"recipient"`
	Value          PostConditionNonFungibleAllOfAssetValue `json:"value"`
}

type _TransactionEventNonFungibleAssetAllOfAsset TransactionEventNonFungibleAssetAllOfAsset

// NewTransactionEventNonFungibleAssetAllOfAsset instantiates a new TransactionEventNonFungibleAssetAllOfAsset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionEventNonFungibleAssetAllOfAsset(assetEventType string, assetId string, sender string, recipient string, value PostConditionNonFungibleAllOfAssetValue) *TransactionEventNonFungibleAssetAllOfAsset {
	this := TransactionEventNonFungibleAssetAllOfAsset{}
	this.AssetEventType = assetEventType
	this.AssetId = assetId
	this.Sender = sender
	this.Recipient = recipient
	this.Value = value
	return &this
}

// NewTransactionEventNonFungibleAssetAllOfAssetWithDefaults instantiates a new TransactionEventNonFungibleAssetAllOfAsset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionEventNonFungibleAssetAllOfAssetWithDefaults() *TransactionEventNonFungibleAssetAllOfAsset {
	this := TransactionEventNonFungibleAssetAllOfAsset{}
	return &this
}

// GetAssetEventType returns the AssetEventType field value
func (o *TransactionEventNonFungibleAssetAllOfAsset) GetAssetEventType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetEventType
}

// GetAssetEventTypeOk returns a tuple with the AssetEventType field value
// and a boolean to check if the value has been set.
func (o *TransactionEventNonFungibleAssetAllOfAsset) GetAssetEventTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetEventType, true
}

// SetAssetEventType sets field value
func (o *TransactionEventNonFungibleAssetAllOfAsset) SetAssetEventType(v string) {
	o.AssetEventType = v
}

// GetAssetId returns the AssetId field value
func (o *TransactionEventNonFungibleAssetAllOfAsset) GetAssetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetId
}

// GetAssetIdOk returns a tuple with the AssetId field value
// and a boolean to check if the value has been set.
func (o *TransactionEventNonFungibleAssetAllOfAsset) GetAssetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetId, true
}

// SetAssetId sets field value
func (o *TransactionEventNonFungibleAssetAllOfAsset) SetAssetId(v string) {
	o.AssetId = v
}

// GetSender returns the Sender field value
func (o *TransactionEventNonFungibleAssetAllOfAsset) GetSender() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sender
}

// GetSenderOk returns a tuple with the Sender field value
// and a boolean to check if the value has been set.
func (o *TransactionEventNonFungibleAssetAllOfAsset) GetSenderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sender, true
}

// SetSender sets field value
func (o *TransactionEventNonFungibleAssetAllOfAsset) SetSender(v string) {
	o.Sender = v
}

// GetRecipient returns the Recipient field value
func (o *TransactionEventNonFungibleAssetAllOfAsset) GetRecipient() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Recipient
}

// GetRecipientOk returns a tuple with the Recipient field value
// and a boolean to check if the value has been set.
func (o *TransactionEventNonFungibleAssetAllOfAsset) GetRecipientOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Recipient, true
}

// SetRecipient sets field value
func (o *TransactionEventNonFungibleAssetAllOfAsset) SetRecipient(v string) {
	o.Recipient = v
}

// GetValue returns the Value field value
func (o *TransactionEventNonFungibleAssetAllOfAsset) GetValue() PostConditionNonFungibleAllOfAssetValue {
	if o == nil {
		var ret PostConditionNonFungibleAllOfAssetValue
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *TransactionEventNonFungibleAssetAllOfAsset) GetValueOk() (*PostConditionNonFungibleAllOfAssetValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *TransactionEventNonFungibleAssetAllOfAsset) SetValue(v PostConditionNonFungibleAllOfAssetValue) {
	o.Value = v
}

func (o TransactionEventNonFungibleAssetAllOfAsset) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionEventNonFungibleAssetAllOfAsset) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["asset_event_type"] = o.AssetEventType
	toSerialize["asset_id"] = o.AssetId
	toSerialize["sender"] = o.Sender
	toSerialize["recipient"] = o.Recipient
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *TransactionEventNonFungibleAssetAllOfAsset) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"asset_event_type",
		"asset_id",
		"sender",
		"recipient",
		"value",
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

	varTransactionEventNonFungibleAssetAllOfAsset := _TransactionEventNonFungibleAssetAllOfAsset{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransactionEventNonFungibleAssetAllOfAsset)

	if err != nil {
		return err
	}

	*o = TransactionEventNonFungibleAssetAllOfAsset(varTransactionEventNonFungibleAssetAllOfAsset)

	return err
}

type NullableTransactionEventNonFungibleAssetAllOfAsset struct {
	value *TransactionEventNonFungibleAssetAllOfAsset
	isSet bool
}

func (v NullableTransactionEventNonFungibleAssetAllOfAsset) Get() *TransactionEventNonFungibleAssetAllOfAsset {
	return v.value
}

func (v *NullableTransactionEventNonFungibleAssetAllOfAsset) Set(val *TransactionEventNonFungibleAssetAllOfAsset) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionEventNonFungibleAssetAllOfAsset) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionEventNonFungibleAssetAllOfAsset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionEventNonFungibleAssetAllOfAsset(val *TransactionEventNonFungibleAssetAllOfAsset) *NullableTransactionEventNonFungibleAssetAllOfAsset {
	return &NullableTransactionEventNonFungibleAssetAllOfAsset{value: val, isSet: true}
}

func (v NullableTransactionEventNonFungibleAssetAllOfAsset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionEventNonFungibleAssetAllOfAsset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
