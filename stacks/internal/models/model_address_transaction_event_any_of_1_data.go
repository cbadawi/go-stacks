package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the AddressTransactionEventAnyOf1Data type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AddressTransactionEventAnyOf1Data{}

// AddressTransactionEventAnyOf1Data struct for AddressTransactionEventAnyOf1Data
type AddressTransactionEventAnyOf1Data struct {
	Type string `json:"type"`
	// Fungible Token asset identifier.
	AssetIdentifier string `json:"asset_identifier"`
	// Amount transferred as an integer string. This balance does not factor in possible SIP-010 decimals.
	Amount string `json:"amount"`
	// Principal that sent the asset.
	Sender *string `json:"sender,omitempty"`
	// Principal that received the asset.
	Recipient *string `json:"recipient,omitempty"`
}

type _AddressTransactionEventAnyOf1Data AddressTransactionEventAnyOf1Data

// NewAddressTransactionEventAnyOf1Data instantiates a new AddressTransactionEventAnyOf1Data object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressTransactionEventAnyOf1Data(type_ string, assetIdentifier string, amount string) *AddressTransactionEventAnyOf1Data {
	this := AddressTransactionEventAnyOf1Data{}
	this.Type = type_
	this.AssetIdentifier = assetIdentifier
	this.Amount = amount
	return &this
}

// NewAddressTransactionEventAnyOf1DataWithDefaults instantiates a new AddressTransactionEventAnyOf1Data object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressTransactionEventAnyOf1DataWithDefaults() *AddressTransactionEventAnyOf1Data {
	this := AddressTransactionEventAnyOf1Data{}
	return &this
}

// GetType returns the Type field value
func (o *AddressTransactionEventAnyOf1Data) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AddressTransactionEventAnyOf1Data) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AddressTransactionEventAnyOf1Data) SetType(v string) {
	o.Type = v
}

// GetAssetIdentifier returns the AssetIdentifier field value
func (o *AddressTransactionEventAnyOf1Data) GetAssetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetIdentifier
}

// GetAssetIdentifierOk returns a tuple with the AssetIdentifier field value
// and a boolean to check if the value has been set.
func (o *AddressTransactionEventAnyOf1Data) GetAssetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetIdentifier, true
}

// SetAssetIdentifier sets field value
func (o *AddressTransactionEventAnyOf1Data) SetAssetIdentifier(v string) {
	o.AssetIdentifier = v
}

// GetAmount returns the Amount field value
func (o *AddressTransactionEventAnyOf1Data) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *AddressTransactionEventAnyOf1Data) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *AddressTransactionEventAnyOf1Data) SetAmount(v string) {
	o.Amount = v
}

// GetSender returns the Sender field value if set, zero value otherwise.
func (o *AddressTransactionEventAnyOf1Data) GetSender() string {
	if o == nil || utils.IsNil(o.Sender) {
		var ret string
		return ret
	}
	return *o.Sender
}

// GetSenderOk returns a tuple with the Sender field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressTransactionEventAnyOf1Data) GetSenderOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Sender) {
		return nil, false
	}
	return o.Sender, true
}

// HasSender returns a boolean if a field has been set.
func (o *AddressTransactionEventAnyOf1Data) HasSender() bool {
	if o != nil && !utils.IsNil(o.Sender) {
		return true
	}

	return false
}

// SetSender gets a reference to the given string and assigns it to the Sender field.
func (o *AddressTransactionEventAnyOf1Data) SetSender(v string) {
	o.Sender = &v
}

// GetRecipient returns the Recipient field value if set, zero value otherwise.
func (o *AddressTransactionEventAnyOf1Data) GetRecipient() string {
	if o == nil || utils.IsNil(o.Recipient) {
		var ret string
		return ret
	}
	return *o.Recipient
}

// GetRecipientOk returns a tuple with the Recipient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressTransactionEventAnyOf1Data) GetRecipientOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Recipient) {
		return nil, false
	}
	return o.Recipient, true
}

// HasRecipient returns a boolean if a field has been set.
func (o *AddressTransactionEventAnyOf1Data) HasRecipient() bool {
	if o != nil && !utils.IsNil(o.Recipient) {
		return true
	}

	return false
}

// SetRecipient gets a reference to the given string and assigns it to the Recipient field.
func (o *AddressTransactionEventAnyOf1Data) SetRecipient(v string) {
	o.Recipient = &v
}

func (o AddressTransactionEventAnyOf1Data) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressTransactionEventAnyOf1Data) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
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

func (o *AddressTransactionEventAnyOf1Data) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varAddressTransactionEventAnyOf1Data := _AddressTransactionEventAnyOf1Data{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddressTransactionEventAnyOf1Data)

	if err != nil {
		return err
	}

	*o = AddressTransactionEventAnyOf1Data(varAddressTransactionEventAnyOf1Data)

	return err
}

type NullableAddressTransactionEventAnyOf1Data struct {
	value *AddressTransactionEventAnyOf1Data
	isSet bool
}

func (v NullableAddressTransactionEventAnyOf1Data) Get() *AddressTransactionEventAnyOf1Data {
	return v.value
}

func (v *NullableAddressTransactionEventAnyOf1Data) Set(val *AddressTransactionEventAnyOf1Data) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressTransactionEventAnyOf1Data) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressTransactionEventAnyOf1Data) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressTransactionEventAnyOf1Data(val *AddressTransactionEventAnyOf1Data) *NullableAddressTransactionEventAnyOf1Data {
	return &NullableAddressTransactionEventAnyOf1Data{value: val, isSet: true}
}

func (v NullableAddressTransactionEventAnyOf1Data) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressTransactionEventAnyOf1Data) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
