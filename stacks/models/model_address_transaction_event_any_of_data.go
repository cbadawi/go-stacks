package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the AddressTransactionEventAnyOfData type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AddressTransactionEventAnyOfData{}

// AddressTransactionEventAnyOfData struct for AddressTransactionEventAnyOfData
type AddressTransactionEventAnyOfData struct {
	Type string `json:"type"`
	// Amount transferred in micro-STX as an integer string.
	Amount string `json:"amount"`
	// Principal that sent STX. This is unspecified if the STX were minted.
	Sender *string `json:"sender,omitempty"`
	// Principal that received STX. This is unspecified if the STX were burned.
	Recipient *string `json:"recipient,omitempty"`
}

type _AddressTransactionEventAnyOfData AddressTransactionEventAnyOfData

// NewAddressTransactionEventAnyOfData instantiates a new AddressTransactionEventAnyOfData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressTransactionEventAnyOfData(type_ string, amount string) *AddressTransactionEventAnyOfData {
	this := AddressTransactionEventAnyOfData{}
	this.Type = type_
	this.Amount = amount
	return &this
}

// NewAddressTransactionEventAnyOfDataWithDefaults instantiates a new AddressTransactionEventAnyOfData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressTransactionEventAnyOfDataWithDefaults() *AddressTransactionEventAnyOfData {
	this := AddressTransactionEventAnyOfData{}
	return &this
}

// GetType returns the Type field value
func (o *AddressTransactionEventAnyOfData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AddressTransactionEventAnyOfData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AddressTransactionEventAnyOfData) SetType(v string) {
	o.Type = v
}

// GetAmount returns the Amount field value
func (o *AddressTransactionEventAnyOfData) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *AddressTransactionEventAnyOfData) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *AddressTransactionEventAnyOfData) SetAmount(v string) {
	o.Amount = v
}

// GetSender returns the Sender field value if set, zero value otherwise.
func (o *AddressTransactionEventAnyOfData) GetSender() string {
	if o == nil || utils.IsNil(o.Sender) {
		var ret string
		return ret
	}
	return *o.Sender
}

// GetSenderOk returns a tuple with the Sender field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressTransactionEventAnyOfData) GetSenderOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Sender) {
		return nil, false
	}
	return o.Sender, true
}

// HasSender returns a boolean if a field has been set.
func (o *AddressTransactionEventAnyOfData) HasSender() bool {
	if o != nil && !utils.IsNil(o.Sender) {
		return true
	}

	return false
}

// SetSender gets a reference to the given string and assigns it to the Sender field.
func (o *AddressTransactionEventAnyOfData) SetSender(v string) {
	o.Sender = &v
}

// GetRecipient returns the Recipient field value if set, zero value otherwise.
func (o *AddressTransactionEventAnyOfData) GetRecipient() string {
	if o == nil || utils.IsNil(o.Recipient) {
		var ret string
		return ret
	}
	return *o.Recipient
}

// GetRecipientOk returns a tuple with the Recipient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressTransactionEventAnyOfData) GetRecipientOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Recipient) {
		return nil, false
	}
	return o.Recipient, true
}

// HasRecipient returns a boolean if a field has been set.
func (o *AddressTransactionEventAnyOfData) HasRecipient() bool {
	if o != nil && !utils.IsNil(o.Recipient) {
		return true
	}

	return false
}

// SetRecipient gets a reference to the given string and assigns it to the Recipient field.
func (o *AddressTransactionEventAnyOfData) SetRecipient(v string) {
	o.Recipient = &v
}

func (o AddressTransactionEventAnyOfData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressTransactionEventAnyOfData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["amount"] = o.Amount
	if !utils.IsNil(o.Sender) {
		toSerialize["sender"] = o.Sender
	}
	if !utils.IsNil(o.Recipient) {
		toSerialize["recipient"] = o.Recipient
	}
	return toSerialize, nil
}

func (o *AddressTransactionEventAnyOfData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varAddressTransactionEventAnyOfData := _AddressTransactionEventAnyOfData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddressTransactionEventAnyOfData)

	if err != nil {
		return err
	}

	*o = AddressTransactionEventAnyOfData(varAddressTransactionEventAnyOfData)

	return err
}

type NullableAddressTransactionEventAnyOfData struct {
	value *AddressTransactionEventAnyOfData
	isSet bool
}

func (v NullableAddressTransactionEventAnyOfData) Get() *AddressTransactionEventAnyOfData {
	return v.value
}

func (v *NullableAddressTransactionEventAnyOfData) Set(val *AddressTransactionEventAnyOfData) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressTransactionEventAnyOfData) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressTransactionEventAnyOfData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressTransactionEventAnyOfData(val *AddressTransactionEventAnyOfData) *NullableAddressTransactionEventAnyOfData {
	return &NullableAddressTransactionEventAnyOfData{value: val, isSet: true}
}

func (v NullableAddressTransactionEventAnyOfData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressTransactionEventAnyOfData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
