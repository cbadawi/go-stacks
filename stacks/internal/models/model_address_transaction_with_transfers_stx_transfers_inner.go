package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the AddressTransactionWithTransfersStxTransfersInner type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AddressTransactionWithTransfersStxTransfersInner{}

// AddressTransactionWithTransfersStxTransfersInner struct for AddressTransactionWithTransfersStxTransfersInner
type AddressTransactionWithTransfersStxTransfersInner struct {
	// Amount transferred in micro-STX as an integer string.
	Amount string `json:"amount"`
	// Principal that sent STX. This is unspecified if the STX were minted.
	Sender *string `json:"sender,omitempty"`
	// Principal that received STX. This is unspecified if the STX were burned.
	Recipient *string `json:"recipient,omitempty"`
}

type _AddressTransactionWithTransfersStxTransfersInner AddressTransactionWithTransfersStxTransfersInner

// NewAddressTransactionWithTransfersStxTransfersInner instantiates a new AddressTransactionWithTransfersStxTransfersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressTransactionWithTransfersStxTransfersInner(amount string) *AddressTransactionWithTransfersStxTransfersInner {
	this := AddressTransactionWithTransfersStxTransfersInner{}
	this.Amount = amount
	return &this
}

// NewAddressTransactionWithTransfersStxTransfersInnerWithDefaults instantiates a new AddressTransactionWithTransfersStxTransfersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressTransactionWithTransfersStxTransfersInnerWithDefaults() *AddressTransactionWithTransfersStxTransfersInner {
	this := AddressTransactionWithTransfersStxTransfersInner{}
	return &this
}

// GetAmount returns the Amount field value
func (o *AddressTransactionWithTransfersStxTransfersInner) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *AddressTransactionWithTransfersStxTransfersInner) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *AddressTransactionWithTransfersStxTransfersInner) SetAmount(v string) {
	o.Amount = v
}

// GetSender returns the Sender field value if set, zero value otherwise.
func (o *AddressTransactionWithTransfersStxTransfersInner) GetSender() string {
	if o == nil || utils.IsNil(o.Sender) {
		var ret string
		return ret
	}
	return *o.Sender
}

// GetSenderOk returns a tuple with the Sender field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressTransactionWithTransfersStxTransfersInner) GetSenderOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Sender) {
		return nil, false
	}
	return o.Sender, true
}

// HasSender returns a boolean if a field has been set.
func (o *AddressTransactionWithTransfersStxTransfersInner) HasSender() bool {
	if o != nil && !utils.IsNil(o.Sender) {
		return true
	}

	return false
}

// SetSender gets a reference to the given string and assigns it to the Sender field.
func (o *AddressTransactionWithTransfersStxTransfersInner) SetSender(v string) {
	o.Sender = &v
}

// GetRecipient returns the Recipient field value if set, zero value otherwise.
func (o *AddressTransactionWithTransfersStxTransfersInner) GetRecipient() string {
	if o == nil || utils.IsNil(o.Recipient) {
		var ret string
		return ret
	}
	return *o.Recipient
}

// GetRecipientOk returns a tuple with the Recipient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressTransactionWithTransfersStxTransfersInner) GetRecipientOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Recipient) {
		return nil, false
	}
	return o.Recipient, true
}

// HasRecipient returns a boolean if a field has been set.
func (o *AddressTransactionWithTransfersStxTransfersInner) HasRecipient() bool {
	if o != nil && !utils.IsNil(o.Recipient) {
		return true
	}

	return false
}

// SetRecipient gets a reference to the given string and assigns it to the Recipient field.
func (o *AddressTransactionWithTransfersStxTransfersInner) SetRecipient(v string) {
	o.Recipient = &v
}

func (o AddressTransactionWithTransfersStxTransfersInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressTransactionWithTransfersStxTransfersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	if !utils.IsNil(o.Sender) {
		toSerialize["sender"] = o.Sender
	}
	if !utils.IsNil(o.Recipient) {
		toSerialize["recipient"] = o.Recipient
	}
	return toSerialize, nil
}

func (o *AddressTransactionWithTransfersStxTransfersInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varAddressTransactionWithTransfersStxTransfersInner := _AddressTransactionWithTransfersStxTransfersInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddressTransactionWithTransfersStxTransfersInner)

	if err != nil {
		return err
	}

	*o = AddressTransactionWithTransfersStxTransfersInner(varAddressTransactionWithTransfersStxTransfersInner)

	return err
}

type NullableAddressTransactionWithTransfersStxTransfersInner struct {
	value *AddressTransactionWithTransfersStxTransfersInner
	isSet bool
}

func (v NullableAddressTransactionWithTransfersStxTransfersInner) Get() *AddressTransactionWithTransfersStxTransfersInner {
	return v.value
}

func (v *NullableAddressTransactionWithTransfersStxTransfersInner) Set(val *AddressTransactionWithTransfersStxTransfersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressTransactionWithTransfersStxTransfersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressTransactionWithTransfersStxTransfersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressTransactionWithTransfersStxTransfersInner(val *AddressTransactionWithTransfersStxTransfersInner) *NullableAddressTransactionWithTransfersStxTransfersInner {
	return &NullableAddressTransactionWithTransfersStxTransfersInner{value: val, isSet: true}
}

func (v NullableAddressTransactionWithTransfersStxTransfersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressTransactionWithTransfersStxTransfersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
