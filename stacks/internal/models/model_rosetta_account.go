package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the RosettaAccount type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RosettaAccount{}

// RosettaAccount The account_identifier uniquely identifies an account within a network. All fields in the account_identifier are utilized to determine this uniqueness (including the metadata field, if populated).
type RosettaAccount struct {
	// The address may be a cryptographic public key (or some encoding of it) or a provided username.
	Address    string             `json:"address"`
	SubAccount *RosettaSubAccount `json:"sub_account,omitempty"`
	// Blockchains that utilize a username model (where the address is not a derivative of a cryptographic public key) should specify the public key(s) owned by the address in metadata.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

type _RosettaAccount RosettaAccount

// NewRosettaAccount instantiates a new RosettaAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRosettaAccount(address string) *RosettaAccount {
	this := RosettaAccount{}
	this.Address = address
	return &this
}

// NewRosettaAccountWithDefaults instantiates a new RosettaAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRosettaAccountWithDefaults() *RosettaAccount {
	this := RosettaAccount{}
	return &this
}

// GetAddress returns the Address field value
func (o *RosettaAccount) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *RosettaAccount) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *RosettaAccount) SetAddress(v string) {
	o.Address = v
}

// GetSubAccount returns the SubAccount field value if set, zero value otherwise.
func (o *RosettaAccount) GetSubAccount() RosettaSubAccount {
	if o == nil || utils.IsNil(o.SubAccount) {
		var ret RosettaSubAccount
		return ret
	}
	return *o.SubAccount
}

// GetSubAccountOk returns a tuple with the SubAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaAccount) GetSubAccountOk() (*RosettaSubAccount, bool) {
	if o == nil || utils.IsNil(o.SubAccount) {
		return nil, false
	}
	return o.SubAccount, true
}

// HasSubAccount returns a boolean if a field has been set.
func (o *RosettaAccount) HasSubAccount() bool {
	if o != nil && !utils.IsNil(o.SubAccount) {
		return true
	}

	return false
}

// SetSubAccount gets a reference to the given RosettaSubAccount and assigns it to the SubAccount field.
func (o *RosettaAccount) SetSubAccount(v RosettaSubAccount) {
	o.SubAccount = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *RosettaAccount) GetMetadata() map[string]interface{} {
	if o == nil || utils.IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaAccount) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || utils.IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *RosettaAccount) HasMetadata() bool {
	if o != nil && !utils.IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *RosettaAccount) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o RosettaAccount) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RosettaAccount) ToMap() (map[string]interface{}, error) {
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

func (o *RosettaAccount) UnmarshalJSON(data []byte) (err error) {
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

	varRosettaAccount := _RosettaAccount{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRosettaAccount)

	if err != nil {
		return err
	}

	*o = RosettaAccount(varRosettaAccount)

	return err
}

type NullableRosettaAccount struct {
	value *RosettaAccount
	isSet bool
}

func (v NullableRosettaAccount) Get() *RosettaAccount {
	return v.value
}

func (v *NullableRosettaAccount) Set(val *RosettaAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableRosettaAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableRosettaAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRosettaAccount(val *RosettaAccount) *NullableRosettaAccount {
	return &NullableRosettaAccount{value: val, isSet: true}
}

func (v NullableRosettaAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRosettaAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
