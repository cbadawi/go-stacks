package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the RosettaSubAccount type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RosettaSubAccount{}

// RosettaSubAccount An account may have state specific to a contract address (ERC-20 token) and/or a stake (delegated balance). The sub_account_identifier should specify which state (if applicable) an account instantiation refers to.
type RosettaSubAccount struct {
	// The address may be a cryptographic public key (or some encoding of it) or a provided username.
	Address string `json:"address"`
	// If the SubAccount address is not sufficient to uniquely specify a SubAccount, any other identifying information can be stored here. It is important to note that two SubAccounts with identical addresses but differing metadata will not be considered equal by clients.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

type _RosettaSubAccount RosettaSubAccount

// NewRosettaSubAccount instantiates a new RosettaSubAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRosettaSubAccount(address string) *RosettaSubAccount {
	this := RosettaSubAccount{}
	this.Address = address
	return &this
}

// NewRosettaSubAccountWithDefaults instantiates a new RosettaSubAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRosettaSubAccountWithDefaults() *RosettaSubAccount {
	this := RosettaSubAccount{}
	return &this
}

// GetAddress returns the Address field value
func (o *RosettaSubAccount) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *RosettaSubAccount) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *RosettaSubAccount) SetAddress(v string) {
	o.Address = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *RosettaSubAccount) GetMetadata() map[string]interface{} {
	if o == nil || utils.IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaSubAccount) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || utils.IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *RosettaSubAccount) HasMetadata() bool {
	if o != nil && !utils.IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *RosettaSubAccount) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o RosettaSubAccount) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RosettaSubAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address"] = o.Address
	if !utils.IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

func (o *RosettaSubAccount) UnmarshalJSON(data []byte) (err error) {
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

	varRosettaSubAccount := _RosettaSubAccount{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRosettaSubAccount)

	if err != nil {
		return err
	}

	*o = RosettaSubAccount(varRosettaSubAccount)

	return err
}

type NullableRosettaSubAccount struct {
	value *RosettaSubAccount
	isSet bool
}

func (v NullableRosettaSubAccount) Get() *RosettaSubAccount {
	return v.value
}

func (v *NullableRosettaSubAccount) Set(val *RosettaSubAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableRosettaSubAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableRosettaSubAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRosettaSubAccount(val *RosettaSubAccount) *NullableRosettaSubAccount {
	return &NullableRosettaSubAccount{value: val, isSet: true}
}

func (v NullableRosettaSubAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRosettaSubAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
