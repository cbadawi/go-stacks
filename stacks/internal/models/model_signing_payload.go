package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the SigningPayload type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &SigningPayload{}

// SigningPayload SigningPayload is signed by the client with the keypair associated with an address using the specified SignatureType. SignatureType can be optionally populated if there is a restriction on the signature scheme that can be used to sign the payload.
type SigningPayload struct {
	// [DEPRECATED by account_identifier in v1.4.4] The network-specific address of the account that should sign the payload.
	Address           *string         `json:"address,omitempty"`
	AccountIdentifier *RosettaAccount `json:"account_identifier,omitempty"`
	HexBytes          string          `json:"hex_bytes"`
	// SignatureType is the type of a cryptographic signature.
	SignatureType *string `json:"signature_type,omitempty"`
}

type _SigningPayload SigningPayload

// NewSigningPayload instantiates a new SigningPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSigningPayload(hexBytes string) *SigningPayload {
	this := SigningPayload{}
	this.HexBytes = hexBytes
	return &this
}

// NewSigningPayloadWithDefaults instantiates a new SigningPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSigningPayloadWithDefaults() *SigningPayload {
	this := SigningPayload{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *SigningPayload) GetAddress() string {
	if o == nil || utils.IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SigningPayload) GetAddressOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *SigningPayload) HasAddress() bool {
	if o != nil && !utils.IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *SigningPayload) SetAddress(v string) {
	o.Address = &v
}

// GetAccountIdentifier returns the AccountIdentifier field value if set, zero value otherwise.
func (o *SigningPayload) GetAccountIdentifier() RosettaAccount {
	if o == nil || utils.IsNil(o.AccountIdentifier) {
		var ret RosettaAccount
		return ret
	}
	return *o.AccountIdentifier
}

// GetAccountIdentifierOk returns a tuple with the AccountIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SigningPayload) GetAccountIdentifierOk() (*RosettaAccount, bool) {
	if o == nil || utils.IsNil(o.AccountIdentifier) {
		return nil, false
	}
	return o.AccountIdentifier, true
}

// HasAccountIdentifier returns a boolean if a field has been set.
func (o *SigningPayload) HasAccountIdentifier() bool {
	if o != nil && !utils.IsNil(o.AccountIdentifier) {
		return true
	}

	return false
}

// SetAccountIdentifier gets a reference to the given RosettaAccount and assigns it to the AccountIdentifier field.
func (o *SigningPayload) SetAccountIdentifier(v RosettaAccount) {
	o.AccountIdentifier = &v
}

// GetHexBytes returns the HexBytes field value
func (o *SigningPayload) GetHexBytes() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HexBytes
}

// GetHexBytesOk returns a tuple with the HexBytes field value
// and a boolean to check if the value has been set.
func (o *SigningPayload) GetHexBytesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HexBytes, true
}

// SetHexBytes sets field value
func (o *SigningPayload) SetHexBytes(v string) {
	o.HexBytes = v
}

// GetSignatureType returns the SignatureType field value if set, zero value otherwise.
func (o *SigningPayload) GetSignatureType() string {
	if o == nil || utils.IsNil(o.SignatureType) {
		var ret string
		return ret
	}
	return *o.SignatureType
}

// GetSignatureTypeOk returns a tuple with the SignatureType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SigningPayload) GetSignatureTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.SignatureType) {
		return nil, false
	}
	return o.SignatureType, true
}

// HasSignatureType returns a boolean if a field has been set.
func (o *SigningPayload) HasSignatureType() bool {
	if o != nil && !utils.IsNil(o.SignatureType) {
		return true
	}

	return false
}

// SetSignatureType gets a reference to the given string and assigns it to the SignatureType field.
func (o *SigningPayload) SetSignatureType(v string) {
	o.SignatureType = &v
}

func (o SigningPayload) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SigningPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !utils.IsNil(o.AccountIdentifier) {
		toSerialize["account_identifier"] = o.AccountIdentifier
	}
	toSerialize["hex_bytes"] = o.HexBytes
	if !utils.IsNil(o.SignatureType) {
		toSerialize["signature_type"] = o.SignatureType
	}
	return toSerialize, nil
}

func (o *SigningPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"hex_bytes",
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

	varSigningPayload := _SigningPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSigningPayload)

	if err != nil {
		return err
	}

	*o = SigningPayload(varSigningPayload)

	return err
}

type NullableSigningPayload struct {
	value *SigningPayload
	isSet bool
}

func (v NullableSigningPayload) Get() *SigningPayload {
	return v.value
}

func (v *NullableSigningPayload) Set(val *SigningPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableSigningPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableSigningPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSigningPayload(val *SigningPayload) *NullableSigningPayload {
	return &NullableSigningPayload{value: val, isSet: true}
}

func (v NullableSigningPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSigningPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
