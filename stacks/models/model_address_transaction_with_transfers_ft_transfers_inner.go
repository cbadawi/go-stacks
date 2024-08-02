package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the AddressTransactionWithTransfersFtTransfersInner type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AddressTransactionWithTransfersFtTransfersInner{}

// AddressTransactionWithTransfersFtTransfersInner struct for AddressTransactionWithTransfersFtTransfersInner
type AddressTransactionWithTransfersFtTransfersInner struct {
	// Fungible Token asset identifier.
	AssetIdentifier string `json:"asset_identifier"`
	// Amount transferred as an integer string. This balance does not factor in possible SIP-010 decimals.
	Amount string `json:"amount"`
	// Principal that sent the asset.
	Sender *string `json:"sender,omitempty"`
	// Principal that received the asset.
	Recipient *string `json:"recipient,omitempty"`
}

type _AddressTransactionWithTransfersFtTransfersInner AddressTransactionWithTransfersFtTransfersInner

// NewAddressTransactionWithTransfersFtTransfersInner instantiates a new AddressTransactionWithTransfersFtTransfersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressTransactionWithTransfersFtTransfersInner(assetIdentifier string, amount string) *AddressTransactionWithTransfersFtTransfersInner {
	this := AddressTransactionWithTransfersFtTransfersInner{}
	this.AssetIdentifier = assetIdentifier
	this.Amount = amount
	return &this
}

// NewAddressTransactionWithTransfersFtTransfersInnerWithDefaults instantiates a new AddressTransactionWithTransfersFtTransfersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressTransactionWithTransfersFtTransfersInnerWithDefaults() *AddressTransactionWithTransfersFtTransfersInner {
	this := AddressTransactionWithTransfersFtTransfersInner{}
	return &this
}

// GetAssetIdentifier returns the AssetIdentifier field value
func (o *AddressTransactionWithTransfersFtTransfersInner) GetAssetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetIdentifier
}

// GetAssetIdentifierOk returns a tuple with the AssetIdentifier field value
// and a boolean to check if the value has been set.
func (o *AddressTransactionWithTransfersFtTransfersInner) GetAssetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetIdentifier, true
}

// SetAssetIdentifier sets field value
func (o *AddressTransactionWithTransfersFtTransfersInner) SetAssetIdentifier(v string) {
	o.AssetIdentifier = v
}

// GetAmount returns the Amount field value
func (o *AddressTransactionWithTransfersFtTransfersInner) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *AddressTransactionWithTransfersFtTransfersInner) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *AddressTransactionWithTransfersFtTransfersInner) SetAmount(v string) {
	o.Amount = v
}

// GetSender returns the Sender field value if set, zero value otherwise.
func (o *AddressTransactionWithTransfersFtTransfersInner) GetSender() string {
	if o == nil || utils.IsNil(o.Sender) {
		var ret string
		return ret
	}
	return *o.Sender
}

// GetSenderOk returns a tuple with the Sender field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressTransactionWithTransfersFtTransfersInner) GetSenderOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Sender) {
		return nil, false
	}
	return o.Sender, true
}

// HasSender returns a boolean if a field has been set.
func (o *AddressTransactionWithTransfersFtTransfersInner) HasSender() bool {
	if o != nil && !utils.IsNil(o.Sender) {
		return true
	}

	return false
}

// SetSender gets a reference to the given string and assigns it to the Sender field.
func (o *AddressTransactionWithTransfersFtTransfersInner) SetSender(v string) {
	o.Sender = &v
}

// GetRecipient returns the Recipient field value if set, zero value otherwise.
func (o *AddressTransactionWithTransfersFtTransfersInner) GetRecipient() string {
	if o == nil || utils.IsNil(o.Recipient) {
		var ret string
		return ret
	}
	return *o.Recipient
}

// GetRecipientOk returns a tuple with the Recipient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressTransactionWithTransfersFtTransfersInner) GetRecipientOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Recipient) {
		return nil, false
	}
	return o.Recipient, true
}

// HasRecipient returns a boolean if a field has been set.
func (o *AddressTransactionWithTransfersFtTransfersInner) HasRecipient() bool {
	if o != nil && !utils.IsNil(o.Recipient) {
		return true
	}

	return false
}

// SetRecipient gets a reference to the given string and assigns it to the Recipient field.
func (o *AddressTransactionWithTransfersFtTransfersInner) SetRecipient(v string) {
	o.Recipient = &v
}

func (o AddressTransactionWithTransfersFtTransfersInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressTransactionWithTransfersFtTransfersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["asset_identifier"] = o.AssetIdentifier
	toSerialize["amount"] = o.Amount
	if !utils.IsNil(o.Sender) {
		toSerialize["sender"] = o.Sender
	}
	if !utils.IsNil(o.Recipient) {
		toSerialize["recipient"] = o.Recipient
	}
	return toSerialize, nil
}

func (o *AddressTransactionWithTransfersFtTransfersInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"asset_identifier",
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

	varAddressTransactionWithTransfersFtTransfersInner := _AddressTransactionWithTransfersFtTransfersInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddressTransactionWithTransfersFtTransfersInner)

	if err != nil {
		return err
	}

	*o = AddressTransactionWithTransfersFtTransfersInner(varAddressTransactionWithTransfersFtTransfersInner)

	return err
}

type NullableAddressTransactionWithTransfersFtTransfersInner struct {
	value *AddressTransactionWithTransfersFtTransfersInner
	isSet bool
}

func (v NullableAddressTransactionWithTransfersFtTransfersInner) Get() *AddressTransactionWithTransfersFtTransfersInner {
	return v.value
}

func (v *NullableAddressTransactionWithTransfersFtTransfersInner) Set(val *AddressTransactionWithTransfersFtTransfersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressTransactionWithTransfersFtTransfersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressTransactionWithTransfersFtTransfersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressTransactionWithTransfersFtTransfersInner(val *AddressTransactionWithTransfersFtTransfersInner) *NullableAddressTransactionWithTransfersFtTransfersInner {
	return &NullableAddressTransactionWithTransfersFtTransfersInner{value: val, isSet: true}
}

func (v NullableAddressTransactionWithTransfersFtTransfersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressTransactionWithTransfersFtTransfersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
