package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the AddressTransactionEventAnyOf2Data type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AddressTransactionEventAnyOf2Data{}

// AddressTransactionEventAnyOf2Data struct for AddressTransactionEventAnyOf2Data
type AddressTransactionEventAnyOf2Data struct {
	Type string `json:"type"`
	// Non Fungible Token asset identifier.
	AssetIdentifier string                                 `json:"asset_identifier"`
	Value           AddressTransactionEventAnyOf2DataValue `json:"value"`
	// Principal that sent the asset.
	Sender *string `json:"sender,omitempty"`
	// Principal that received the asset.
	Recipient *string `json:"recipient,omitempty"`
}

type _AddressTransactionEventAnyOf2Data AddressTransactionEventAnyOf2Data

// NewAddressTransactionEventAnyOf2Data instantiates a new AddressTransactionEventAnyOf2Data object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressTransactionEventAnyOf2Data(type_ string, assetIdentifier string, value AddressTransactionEventAnyOf2DataValue) *AddressTransactionEventAnyOf2Data {
	this := AddressTransactionEventAnyOf2Data{}
	this.Type = type_
	this.AssetIdentifier = assetIdentifier
	this.Value = value
	return &this
}

// NewAddressTransactionEventAnyOf2DataWithDefaults instantiates a new AddressTransactionEventAnyOf2Data object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressTransactionEventAnyOf2DataWithDefaults() *AddressTransactionEventAnyOf2Data {
	this := AddressTransactionEventAnyOf2Data{}
	return &this
}

// GetType returns the Type field value
func (o *AddressTransactionEventAnyOf2Data) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AddressTransactionEventAnyOf2Data) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AddressTransactionEventAnyOf2Data) SetType(v string) {
	o.Type = v
}

// GetAssetIdentifier returns the AssetIdentifier field value
func (o *AddressTransactionEventAnyOf2Data) GetAssetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetIdentifier
}

// GetAssetIdentifierOk returns a tuple with the AssetIdentifier field value
// and a boolean to check if the value has been set.
func (o *AddressTransactionEventAnyOf2Data) GetAssetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetIdentifier, true
}

// SetAssetIdentifier sets field value
func (o *AddressTransactionEventAnyOf2Data) SetAssetIdentifier(v string) {
	o.AssetIdentifier = v
}

// GetValue returns the Value field value
func (o *AddressTransactionEventAnyOf2Data) GetValue() AddressTransactionEventAnyOf2DataValue {
	if o == nil {
		var ret AddressTransactionEventAnyOf2DataValue
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *AddressTransactionEventAnyOf2Data) GetValueOk() (*AddressTransactionEventAnyOf2DataValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *AddressTransactionEventAnyOf2Data) SetValue(v AddressTransactionEventAnyOf2DataValue) {
	o.Value = v
}

// GetSender returns the Sender field value if set, zero value otherwise.
func (o *AddressTransactionEventAnyOf2Data) GetSender() string {
	if o == nil || utils.IsNil(o.Sender) {
		var ret string
		return ret
	}
	return *o.Sender
}

// GetSenderOk returns a tuple with the Sender field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressTransactionEventAnyOf2Data) GetSenderOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Sender) {
		return nil, false
	}
	return o.Sender, true
}

// HasSender returns a boolean if a field has been set.
func (o *AddressTransactionEventAnyOf2Data) HasSender() bool {
	if o != nil && !utils.IsNil(o.Sender) {
		return true
	}

	return false
}

// SetSender gets a reference to the given string and assigns it to the Sender field.
func (o *AddressTransactionEventAnyOf2Data) SetSender(v string) {
	o.Sender = &v
}

// GetRecipient returns the Recipient field value if set, zero value otherwise.
func (o *AddressTransactionEventAnyOf2Data) GetRecipient() string {
	if o == nil || utils.IsNil(o.Recipient) {
		var ret string
		return ret
	}
	return *o.Recipient
}

// GetRecipientOk returns a tuple with the Recipient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressTransactionEventAnyOf2Data) GetRecipientOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Recipient) {
		return nil, false
	}
	return o.Recipient, true
}

// HasRecipient returns a boolean if a field has been set.
func (o *AddressTransactionEventAnyOf2Data) HasRecipient() bool {
	if o != nil && !utils.IsNil(o.Recipient) {
		return true
	}

	return false
}

// SetRecipient gets a reference to the given string and assigns it to the Recipient field.
func (o *AddressTransactionEventAnyOf2Data) SetRecipient(v string) {
	o.Recipient = &v
}

func (o AddressTransactionEventAnyOf2Data) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressTransactionEventAnyOf2Data) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
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

func (o *AddressTransactionEventAnyOf2Data) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varAddressTransactionEventAnyOf2Data := _AddressTransactionEventAnyOf2Data{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddressTransactionEventAnyOf2Data)

	if err != nil {
		return err
	}

	*o = AddressTransactionEventAnyOf2Data(varAddressTransactionEventAnyOf2Data)

	return err
}

type NullableAddressTransactionEventAnyOf2Data struct {
	value *AddressTransactionEventAnyOf2Data
	isSet bool
}

func (v NullableAddressTransactionEventAnyOf2Data) Get() *AddressTransactionEventAnyOf2Data {
	return v.value
}

func (v *NullableAddressTransactionEventAnyOf2Data) Set(val *AddressTransactionEventAnyOf2Data) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressTransactionEventAnyOf2Data) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressTransactionEventAnyOf2Data) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressTransactionEventAnyOf2Data(val *AddressTransactionEventAnyOf2Data) *NullableAddressTransactionEventAnyOf2Data {
	return &NullableAddressTransactionEventAnyOf2Data{value: val, isSet: true}
}

func (v NullableAddressTransactionEventAnyOf2Data) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressTransactionEventAnyOf2Data) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
