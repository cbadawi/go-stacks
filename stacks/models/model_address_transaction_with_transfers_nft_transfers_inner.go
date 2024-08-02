package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the AddressTransactionWithTransfersNftTransfersInner type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AddressTransactionWithTransfersNftTransfersInner{}

// AddressTransactionWithTransfersNftTransfersInner struct for AddressTransactionWithTransfersNftTransfersInner
type AddressTransactionWithTransfersNftTransfersInner struct {
	// Non Fungible Token asset identifier.
	AssetIdentifier string                                 `json:"asset_identifier"`
	Value           AddressTransactionEventAnyOf2DataValue `json:"value"`
	// Principal that sent the asset.
	Sender *string `json:"sender,omitempty"`
	// Principal that received the asset.
	Recipient *string `json:"recipient,omitempty"`
}

type _AddressTransactionWithTransfersNftTransfersInner AddressTransactionWithTransfersNftTransfersInner

// NewAddressTransactionWithTransfersNftTransfersInner instantiates a new AddressTransactionWithTransfersNftTransfersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressTransactionWithTransfersNftTransfersInner(assetIdentifier string, value AddressTransactionEventAnyOf2DataValue) *AddressTransactionWithTransfersNftTransfersInner {
	this := AddressTransactionWithTransfersNftTransfersInner{}
	this.AssetIdentifier = assetIdentifier
	this.Value = value
	return &this
}

// NewAddressTransactionWithTransfersNftTransfersInnerWithDefaults instantiates a new AddressTransactionWithTransfersNftTransfersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressTransactionWithTransfersNftTransfersInnerWithDefaults() *AddressTransactionWithTransfersNftTransfersInner {
	this := AddressTransactionWithTransfersNftTransfersInner{}
	return &this
}

// GetAssetIdentifier returns the AssetIdentifier field value
func (o *AddressTransactionWithTransfersNftTransfersInner) GetAssetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetIdentifier
}

// GetAssetIdentifierOk returns a tuple with the AssetIdentifier field value
// and a boolean to check if the value has been set.
func (o *AddressTransactionWithTransfersNftTransfersInner) GetAssetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetIdentifier, true
}

// SetAssetIdentifier sets field value
func (o *AddressTransactionWithTransfersNftTransfersInner) SetAssetIdentifier(v string) {
	o.AssetIdentifier = v
}

// GetValue returns the Value field value
func (o *AddressTransactionWithTransfersNftTransfersInner) GetValue() AddressTransactionEventAnyOf2DataValue {
	if o == nil {
		var ret AddressTransactionEventAnyOf2DataValue
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *AddressTransactionWithTransfersNftTransfersInner) GetValueOk() (*AddressTransactionEventAnyOf2DataValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *AddressTransactionWithTransfersNftTransfersInner) SetValue(v AddressTransactionEventAnyOf2DataValue) {
	o.Value = v
}

// GetSender returns the Sender field value if set, zero value otherwise.
func (o *AddressTransactionWithTransfersNftTransfersInner) GetSender() string {
	if o == nil || utils.IsNil(o.Sender) {
		var ret string
		return ret
	}
	return *o.Sender
}

// GetSenderOk returns a tuple with the Sender field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressTransactionWithTransfersNftTransfersInner) GetSenderOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Sender) {
		return nil, false
	}
	return o.Sender, true
}

// HasSender returns a boolean if a field has been set.
func (o *AddressTransactionWithTransfersNftTransfersInner) HasSender() bool {
	if o != nil && !utils.IsNil(o.Sender) {
		return true
	}

	return false
}

// SetSender gets a reference to the given string and assigns it to the Sender field.
func (o *AddressTransactionWithTransfersNftTransfersInner) SetSender(v string) {
	o.Sender = &v
}

// GetRecipient returns the Recipient field value if set, zero value otherwise.
func (o *AddressTransactionWithTransfersNftTransfersInner) GetRecipient() string {
	if o == nil || utils.IsNil(o.Recipient) {
		var ret string
		return ret
	}
	return *o.Recipient
}

// GetRecipientOk returns a tuple with the Recipient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressTransactionWithTransfersNftTransfersInner) GetRecipientOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Recipient) {
		return nil, false
	}
	return o.Recipient, true
}

// HasRecipient returns a boolean if a field has been set.
func (o *AddressTransactionWithTransfersNftTransfersInner) HasRecipient() bool {
	if o != nil && !utils.IsNil(o.Recipient) {
		return true
	}

	return false
}

// SetRecipient gets a reference to the given string and assigns it to the Recipient field.
func (o *AddressTransactionWithTransfersNftTransfersInner) SetRecipient(v string) {
	o.Recipient = &v
}

func (o AddressTransactionWithTransfersNftTransfersInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressTransactionWithTransfersNftTransfersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["asset_identifier"] = o.AssetIdentifier
	toSerialize["value"] = o.Value
	if !utils.IsNil(o.Sender) {
		toSerialize["sender"] = o.Sender
	}
	if !utils.IsNil(o.Recipient) {
		toSerialize["recipient"] = o.Recipient
	}
	return toSerialize, nil
}

func (o *AddressTransactionWithTransfersNftTransfersInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"asset_identifier",
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

	varAddressTransactionWithTransfersNftTransfersInner := _AddressTransactionWithTransfersNftTransfersInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddressTransactionWithTransfersNftTransfersInner)

	if err != nil {
		return err
	}

	*o = AddressTransactionWithTransfersNftTransfersInner(varAddressTransactionWithTransfersNftTransfersInner)

	return err
}

type NullableAddressTransactionWithTransfersNftTransfersInner struct {
	value *AddressTransactionWithTransfersNftTransfersInner
	isSet bool
}

func (v NullableAddressTransactionWithTransfersNftTransfersInner) Get() *AddressTransactionWithTransfersNftTransfersInner {
	return v.value
}

func (v *NullableAddressTransactionWithTransfersNftTransfersInner) Set(val *AddressTransactionWithTransfersNftTransfersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressTransactionWithTransfersNftTransfersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressTransactionWithTransfersNftTransfersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressTransactionWithTransfersNftTransfersInner(val *AddressTransactionWithTransfersNftTransfersInner) *NullableAddressTransactionWithTransfersNftTransfersInner {
	return &NullableAddressTransactionWithTransfersNftTransfersInner{value: val, isSet: true}
}

func (v NullableAddressTransactionWithTransfersNftTransfersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressTransactionWithTransfersNftTransfersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
