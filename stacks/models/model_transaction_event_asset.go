package models

import (
	"encoding/json"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the TransactionEventAsset type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &TransactionEventAsset{}

// TransactionEventAsset struct for TransactionEventAsset
type TransactionEventAsset struct {
	AssetEventType *string `json:"asset_event_type,omitempty"`
	AssetId        *string `json:"asset_id,omitempty"`
	Sender         *string `json:"sender,omitempty"`
	Recipient      *string `json:"recipient,omitempty"`
	Amount         *string `json:"amount,omitempty"`
	Value          *string `json:"value,omitempty"`
	Memo           *string `json:"memo,omitempty"`
}

// NewTransactionEventAsset instantiates a new TransactionEventAsset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionEventAsset() *TransactionEventAsset {
	this := TransactionEventAsset{}
	return &this
}

// NewTransactionEventAssetWithDefaults instantiates a new TransactionEventAsset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionEventAssetWithDefaults() *TransactionEventAsset {
	this := TransactionEventAsset{}
	return &this
}

// GetAssetEventType returns the AssetEventType field value if set, zero value otherwise.
func (o *TransactionEventAsset) GetAssetEventType() string {
	if o == nil || utils.IsNil(o.AssetEventType) {
		var ret string
		return ret
	}
	return *o.AssetEventType
}

// GetAssetEventTypeOk returns a tuple with the AssetEventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionEventAsset) GetAssetEventTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.AssetEventType) {
		return nil, false
	}
	return o.AssetEventType, true
}

// HasAssetEventType returns a boolean if a field has been set.
func (o *TransactionEventAsset) HasAssetEventType() bool {
	if o != nil && !utils.IsNil(o.AssetEventType) {
		return true
	}

	return false
}

// SetAssetEventType gets a reference to the given string and assigns it to the AssetEventType field.
func (o *TransactionEventAsset) SetAssetEventType(v string) {
	o.AssetEventType = &v
}

// GetAssetId returns the AssetId field value if set, zero value otherwise.
func (o *TransactionEventAsset) GetAssetId() string {
	if o == nil || utils.IsNil(o.AssetId) {
		var ret string
		return ret
	}
	return *o.AssetId
}

// GetAssetIdOk returns a tuple with the AssetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionEventAsset) GetAssetIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.AssetId) {
		return nil, false
	}
	return o.AssetId, true
}

// HasAssetId returns a boolean if a field has been set.
func (o *TransactionEventAsset) HasAssetId() bool {
	if o != nil && !utils.IsNil(o.AssetId) {
		return true
	}

	return false
}

// SetAssetId gets a reference to the given string and assigns it to the AssetId field.
func (o *TransactionEventAsset) SetAssetId(v string) {
	o.AssetId = &v
}

// GetSender returns the Sender field value if set, zero value otherwise.
func (o *TransactionEventAsset) GetSender() string {
	if o == nil || utils.IsNil(o.Sender) {
		var ret string
		return ret
	}
	return *o.Sender
}

// GetSenderOk returns a tuple with the Sender field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionEventAsset) GetSenderOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Sender) {
		return nil, false
	}
	return o.Sender, true
}

// HasSender returns a boolean if a field has been set.
func (o *TransactionEventAsset) HasSender() bool {
	if o != nil && !utils.IsNil(o.Sender) {
		return true
	}

	return false
}

// SetSender gets a reference to the given string and assigns it to the Sender field.
func (o *TransactionEventAsset) SetSender(v string) {
	o.Sender = &v
}

// GetRecipient returns the Recipient field value if set, zero value otherwise.
func (o *TransactionEventAsset) GetRecipient() string {
	if o == nil || utils.IsNil(o.Recipient) {
		var ret string
		return ret
	}
	return *o.Recipient
}

// GetRecipientOk returns a tuple with the Recipient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionEventAsset) GetRecipientOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Recipient) {
		return nil, false
	}
	return o.Recipient, true
}

// HasRecipient returns a boolean if a field has been set.
func (o *TransactionEventAsset) HasRecipient() bool {
	if o != nil && !utils.IsNil(o.Recipient) {
		return true
	}

	return false
}

// SetRecipient gets a reference to the given string and assigns it to the Recipient field.
func (o *TransactionEventAsset) SetRecipient(v string) {
	o.Recipient = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *TransactionEventAsset) GetAmount() string {
	if o == nil || utils.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionEventAsset) GetAmountOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *TransactionEventAsset) HasAmount() bool {
	if o != nil && !utils.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *TransactionEventAsset) SetAmount(v string) {
	o.Amount = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *TransactionEventAsset) GetValue() string {
	if o == nil || utils.IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionEventAsset) GetValueOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *TransactionEventAsset) HasValue() bool {
	if o != nil && !utils.IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *TransactionEventAsset) SetValue(v string) {
	o.Value = &v
}

// GetMemo returns the Memo field value if set, zero value otherwise.
func (o *TransactionEventAsset) GetMemo() string {
	if o == nil || utils.IsNil(o.Memo) {
		var ret string
		return ret
	}
	return *o.Memo
}

// GetMemoOk returns a tuple with the Memo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionEventAsset) GetMemoOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Memo) {
		return nil, false
	}
	return o.Memo, true
}

// HasMemo returns a boolean if a field has been set.
func (o *TransactionEventAsset) HasMemo() bool {
	if o != nil && !utils.IsNil(o.Memo) {
		return true
	}

	return false
}

// SetMemo gets a reference to the given string and assigns it to the Memo field.
func (o *TransactionEventAsset) SetMemo(v string) {
	o.Memo = &v
}

func (o TransactionEventAsset) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionEventAsset) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.AssetEventType) {
		toSerialize["asset_event_type"] = o.AssetEventType
	}
	if !utils.IsNil(o.AssetId) {
		toSerialize["asset_id"] = o.AssetId
	}
	if !utils.IsNil(o.Sender) {
		toSerialize["sender"] = o.Sender
	}
	if !utils.IsNil(o.Recipient) {
		toSerialize["recipient"] = o.Recipient
	}
	if !utils.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !utils.IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !utils.IsNil(o.Memo) {
		toSerialize["memo"] = o.Memo
	}
	return toSerialize, nil
}

type NullableTransactionEventAsset struct {
	value *TransactionEventAsset
	isSet bool
}

func (v NullableTransactionEventAsset) Get() *TransactionEventAsset {
	return v.value
}

func (v *NullableTransactionEventAsset) Set(val *TransactionEventAsset) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionEventAsset) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionEventAsset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionEventAsset(val *TransactionEventAsset) *NullableTransactionEventAsset {
	return &NullableTransactionEventAsset{value: val, isSet: true}
}

func (v NullableTransactionEventAsset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionEventAsset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
