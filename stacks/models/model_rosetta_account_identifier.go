package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the RosettaAccountIdentifier type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RosettaAccountIdentifier{}

// RosettaAccountIdentifier The account_identifier uniquely identifies an account within a network. All fields in the account_identifier are utilized to determine this uniqueness (including the metadata field, if populated).
type RosettaAccountIdentifier struct {
	// The address may be a cryptographic public key (or some encoding of it) or a provided username.
	Address    string             `json:"address"`
	SubAccount *RosettaSubAccount `json:"sub_account,omitempty"`
	// Blockchains that utilize a username model (where the address is not a derivative of a cryptographic public key) should specify the public key(s) owned by the address in metadata.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

type _RosettaAccountIdentifier RosettaAccountIdentifier

// NewRosettaAccountIdentifier instantiates a new RosettaAccountIdentifier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRosettaAccountIdentifier(address string) *RosettaAccountIdentifier {
	this := RosettaAccountIdentifier{}
	this.Address = address
	return &this
}

// NewRosettaAccountIdentifierWithDefaults instantiates a new RosettaAccountIdentifier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRosettaAccountIdentifierWithDefaults() *RosettaAccountIdentifier {
	this := RosettaAccountIdentifier{}
	return &this
}

// GetAddress returns the Address field value
func (o *RosettaAccountIdentifier) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *RosettaAccountIdentifier) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *RosettaAccountIdentifier) SetAddress(v string) {
	o.Address = v
}

// GetSubAccount returns the SubAccount field value if set, zero value otherwise.
func (o *RosettaAccountIdentifier) GetSubAccount() RosettaSubAccount {
	if o == nil || utils.IsNil(o.SubAccount) {
		var ret RosettaSubAccount
		return ret
	}
	return *o.SubAccount
}

// GetSubAccountOk returns a tuple with the SubAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaAccountIdentifier) GetSubAccountOk() (*RosettaSubAccount, bool) {
	if o == nil || utils.IsNil(o.SubAccount) {
		return nil, false
	}
	return o.SubAccount, true
}

// HasSubAccount returns a boolean if a field has been set.
func (o *RosettaAccountIdentifier) HasSubAccount() bool {
	if o != nil && !utils.IsNil(o.SubAccount) {
		return true
	}

	return false
}

// SetSubAccount gets a reference to the given RosettaSubAccount and assigns it to the SubAccount field.
func (o *RosettaAccountIdentifier) SetSubAccount(v RosettaSubAccount) {
	o.SubAccount = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *RosettaAccountIdentifier) GetMetadata() map[string]interface{} {
	if o == nil || utils.IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaAccountIdentifier) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || utils.IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *RosettaAccountIdentifier) HasMetadata() bool {
	if o != nil && !utils.IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *RosettaAccountIdentifier) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o RosettaAccountIdentifier) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RosettaAccountIdentifier) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address"] = o.Address
	if !utils.IsNil(o.SubAccount) {
		toSerialize["sub_account"] = o.SubAccount
	}
	if !utils.IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

func (o *RosettaAccountIdentifier) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"address",
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

	varRosettaAccountIdentifier := _RosettaAccountIdentifier{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRosettaAccountIdentifier)

	if err != nil {
		return err
	}

	*o = RosettaAccountIdentifier(varRosettaAccountIdentifier)

	return err
}

type NullableRosettaAccountIdentifier struct {
	value *RosettaAccountIdentifier
	isSet bool
}

func (v NullableRosettaAccountIdentifier) Get() *RosettaAccountIdentifier {
	return v.value
}

func (v *NullableRosettaAccountIdentifier) Set(val *RosettaAccountIdentifier) {
	v.value = val
	v.isSet = true
}

func (v NullableRosettaAccountIdentifier) IsSet() bool {
	return v.isSet
}

func (v *NullableRosettaAccountIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRosettaAccountIdentifier(val *RosettaAccountIdentifier) *NullableRosettaAccountIdentifier {
	return &NullableRosettaAccountIdentifier{value: val, isSet: true}
}

func (v NullableRosettaAccountIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRosettaAccountIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
