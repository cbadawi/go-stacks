package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the TokenTransferTransactionMetadataTokenTransfer type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &TokenTransferTransactionMetadataTokenTransfer{}

// TokenTransferTransactionMetadataTokenTransfer struct for TokenTransferTransactionMetadataTokenTransfer
type TokenTransferTransactionMetadataTokenTransfer struct {
	RecipientAddress string `json:"recipient_address"`
	// Transfer amount as Integer string (64-bit unsigned integer)
	Amount string `json:"amount"`
	// Hex encoded arbitrary message, up to 34 bytes length (should try decoding to an ASCII string)
	Memo string `json:"memo"`
}

type _TokenTransferTransactionMetadataTokenTransfer TokenTransferTransactionMetadataTokenTransfer

// NewTokenTransferTransactionMetadataTokenTransfer instantiates a new TokenTransferTransactionMetadataTokenTransfer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenTransferTransactionMetadataTokenTransfer(recipientAddress string, amount string, memo string) *TokenTransferTransactionMetadataTokenTransfer {
	this := TokenTransferTransactionMetadataTokenTransfer{}
	this.RecipientAddress = recipientAddress
	this.Amount = amount
	this.Memo = memo
	return &this
}

// NewTokenTransferTransactionMetadataTokenTransferWithDefaults instantiates a new TokenTransferTransactionMetadataTokenTransfer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenTransferTransactionMetadataTokenTransferWithDefaults() *TokenTransferTransactionMetadataTokenTransfer {
	this := TokenTransferTransactionMetadataTokenTransfer{}
	return &this
}

// GetRecipientAddress returns the RecipientAddress field value
func (o *TokenTransferTransactionMetadataTokenTransfer) GetRecipientAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecipientAddress
}

// GetRecipientAddressOk returns a tuple with the RecipientAddress field value
// and a boolean to check if the value has been set.
func (o *TokenTransferTransactionMetadataTokenTransfer) GetRecipientAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecipientAddress, true
}

// SetRecipientAddress sets field value
func (o *TokenTransferTransactionMetadataTokenTransfer) SetRecipientAddress(v string) {
	o.RecipientAddress = v
}

// GetAmount returns the Amount field value
func (o *TokenTransferTransactionMetadataTokenTransfer) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *TokenTransferTransactionMetadataTokenTransfer) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *TokenTransferTransactionMetadataTokenTransfer) SetAmount(v string) {
	o.Amount = v
}

// GetMemo returns the Memo field value
func (o *TokenTransferTransactionMetadataTokenTransfer) GetMemo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Memo
}

// GetMemoOk returns a tuple with the Memo field value
// and a boolean to check if the value has been set.
func (o *TokenTransferTransactionMetadataTokenTransfer) GetMemoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Memo, true
}

// SetMemo sets field value
func (o *TokenTransferTransactionMetadataTokenTransfer) SetMemo(v string) {
	o.Memo = v
}

func (o TokenTransferTransactionMetadataTokenTransfer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenTransferTransactionMetadataTokenTransfer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["recipient_address"] = o.RecipientAddress
	toSerialize["amount"] = o.Amount
	toSerialize["memo"] = o.Memo
	return toSerialize, nil
}

func (o *TokenTransferTransactionMetadataTokenTransfer) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"recipient_address",
		"amount",
		"memo",
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

	varTokenTransferTransactionMetadataTokenTransfer := _TokenTransferTransactionMetadataTokenTransfer{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTokenTransferTransactionMetadataTokenTransfer)

	if err != nil {
		return err
	}

	*o = TokenTransferTransactionMetadataTokenTransfer(varTokenTransferTransactionMetadataTokenTransfer)

	return err
}

type NullableTokenTransferTransactionMetadataTokenTransfer struct {
	value *TokenTransferTransactionMetadataTokenTransfer
	isSet bool
}

func (v NullableTokenTransferTransactionMetadataTokenTransfer) Get() *TokenTransferTransactionMetadataTokenTransfer {
	return v.value
}

func (v *NullableTokenTransferTransactionMetadataTokenTransfer) Set(val *TokenTransferTransactionMetadataTokenTransfer) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenTransferTransactionMetadataTokenTransfer) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenTransferTransactionMetadataTokenTransfer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenTransferTransactionMetadataTokenTransfer(val *TokenTransferTransactionMetadataTokenTransfer) *NullableTokenTransferTransactionMetadataTokenTransfer {
	return &NullableTokenTransferTransactionMetadataTokenTransfer{value: val, isSet: true}
}

func (v NullableTokenTransferTransactionMetadataTokenTransfer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenTransferTransactionMetadataTokenTransfer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
