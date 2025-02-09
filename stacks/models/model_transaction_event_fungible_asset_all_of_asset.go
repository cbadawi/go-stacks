package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the TransactionEventFungibleAssetAllOfAsset type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &TransactionEventFungibleAssetAllOfAsset{}

// TransactionEventFungibleAssetAllOfAsset struct for TransactionEventFungibleAssetAllOfAsset
type TransactionEventFungibleAssetAllOfAsset struct {
	AssetEventType string `json:"asset_event_type"`
	AssetId        string `json:"asset_id"`
	Sender         string `json:"sender"`
	Recipient      string `json:"recipient"`
	Amount         string `json:"amount"`
}

type _TransactionEventFungibleAssetAllOfAsset TransactionEventFungibleAssetAllOfAsset

// NewTransactionEventFungibleAssetAllOfAsset instantiates a new TransactionEventFungibleAssetAllOfAsset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionEventFungibleAssetAllOfAsset(assetEventType string, assetId string, sender string, recipient string, amount string) *TransactionEventFungibleAssetAllOfAsset {
	this := TransactionEventFungibleAssetAllOfAsset{}
	this.AssetEventType = assetEventType
	this.AssetId = assetId
	this.Sender = sender
	this.Recipient = recipient
	this.Amount = amount
	return &this
}

// NewTransactionEventFungibleAssetAllOfAssetWithDefaults instantiates a new TransactionEventFungibleAssetAllOfAsset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionEventFungibleAssetAllOfAssetWithDefaults() *TransactionEventFungibleAssetAllOfAsset {
	this := TransactionEventFungibleAssetAllOfAsset{}
	return &this
}

// GetAssetEventType returns the AssetEventType field value
func (o *TransactionEventFungibleAssetAllOfAsset) GetAssetEventType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetEventType
}

// GetAssetEventTypeOk returns a tuple with the AssetEventType field value
// and a boolean to check if the value has been set.
func (o *TransactionEventFungibleAssetAllOfAsset) GetAssetEventTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetEventType, true
}

// SetAssetEventType sets field value
func (o *TransactionEventFungibleAssetAllOfAsset) SetAssetEventType(v string) {
	o.AssetEventType = v
}

// GetAssetId returns the AssetId field value
func (o *TransactionEventFungibleAssetAllOfAsset) GetAssetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetId
}

// GetAssetIdOk returns a tuple with the AssetId field value
// and a boolean to check if the value has been set.
func (o *TransactionEventFungibleAssetAllOfAsset) GetAssetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetId, true
}

// SetAssetId sets field value
func (o *TransactionEventFungibleAssetAllOfAsset) SetAssetId(v string) {
	o.AssetId = v
}

// GetSender returns the Sender field value
func (o *TransactionEventFungibleAssetAllOfAsset) GetSender() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sender
}

// GetSenderOk returns a tuple with the Sender field value
// and a boolean to check if the value has been set.
func (o *TransactionEventFungibleAssetAllOfAsset) GetSenderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sender, true
}

// SetSender sets field value
func (o *TransactionEventFungibleAssetAllOfAsset) SetSender(v string) {
	o.Sender = v
}

// GetRecipient returns the Recipient field value
func (o *TransactionEventFungibleAssetAllOfAsset) GetRecipient() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Recipient
}

// GetRecipientOk returns a tuple with the Recipient field value
// and a boolean to check if the value has been set.
func (o *TransactionEventFungibleAssetAllOfAsset) GetRecipientOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Recipient, true
}

// SetRecipient sets field value
func (o *TransactionEventFungibleAssetAllOfAsset) SetRecipient(v string) {
	o.Recipient = v
}

// GetAmount returns the Amount field value
func (o *TransactionEventFungibleAssetAllOfAsset) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *TransactionEventFungibleAssetAllOfAsset) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *TransactionEventFungibleAssetAllOfAsset) SetAmount(v string) {
	o.Amount = v
}

func (o TransactionEventFungibleAssetAllOfAsset) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionEventFungibleAssetAllOfAsset) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["asset_event_type"] = o.AssetEventType
	toSerialize["asset_id"] = o.AssetId
	toSerialize["sender"] = o.Sender
	toSerialize["recipient"] = o.Recipient
	toSerialize["amount"] = o.Amount
	return toSerialize, nil
}

func (o *TransactionEventFungibleAssetAllOfAsset) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"asset_event_type",
		"asset_id",
		"sender",
		"recipient",
		"amount",
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

	varTransactionEventFungibleAssetAllOfAsset := _TransactionEventFungibleAssetAllOfAsset{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransactionEventFungibleAssetAllOfAsset)

	if err != nil {
		return err
	}

	*o = TransactionEventFungibleAssetAllOfAsset(varTransactionEventFungibleAssetAllOfAsset)

	return err
}

type NullableTransactionEventFungibleAssetAllOfAsset struct {
	value *TransactionEventFungibleAssetAllOfAsset
	isSet bool
}

func (v NullableTransactionEventFungibleAssetAllOfAsset) Get() *TransactionEventFungibleAssetAllOfAsset {
	return v.value
}

func (v *NullableTransactionEventFungibleAssetAllOfAsset) Set(val *TransactionEventFungibleAssetAllOfAsset) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionEventFungibleAssetAllOfAsset) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionEventFungibleAssetAllOfAsset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionEventFungibleAssetAllOfAsset(val *TransactionEventFungibleAssetAllOfAsset) *NullableTransactionEventFungibleAssetAllOfAsset {
	return &NullableTransactionEventFungibleAssetAllOfAsset{value: val, isSet: true}
}

func (v NullableTransactionEventFungibleAssetAllOfAsset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionEventFungibleAssetAllOfAsset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
