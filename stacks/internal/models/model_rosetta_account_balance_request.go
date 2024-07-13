package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the RosettaAccountBalanceRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RosettaAccountBalanceRequest{}

// RosettaAccountBalanceRequest An AccountBalanceRequest is utilized to make a balance request on the /account/balance endpoint. If the block_identifier is populated, a historical balance query should be performed.
type RosettaAccountBalanceRequest struct {
	NetworkIdentifier NetworkIdentifier              `json:"network_identifier"`
	AccountIdentifier RosettaAccount                 `json:"account_identifier"`
	BlockIdentifier   *RosettaPartialBlockIdentifier `json:"block_identifier,omitempty"`
}

type _RosettaAccountBalanceRequest RosettaAccountBalanceRequest

// NewRosettaAccountBalanceRequest instantiates a new RosettaAccountBalanceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRosettaAccountBalanceRequest(networkIdentifier NetworkIdentifier, accountIdentifier RosettaAccount) *RosettaAccountBalanceRequest {
	this := RosettaAccountBalanceRequest{}
	this.NetworkIdentifier = networkIdentifier
	this.AccountIdentifier = accountIdentifier
	return &this
}

// NewRosettaAccountBalanceRequestWithDefaults instantiates a new RosettaAccountBalanceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRosettaAccountBalanceRequestWithDefaults() *RosettaAccountBalanceRequest {
	this := RosettaAccountBalanceRequest{}
	return &this
}

// GetNetworkIdentifier returns the NetworkIdentifier field value
func (o *RosettaAccountBalanceRequest) GetNetworkIdentifier() NetworkIdentifier {
	if o == nil {
		var ret NetworkIdentifier
		return ret
	}

	return o.NetworkIdentifier
}

// GetNetworkIdentifierOk returns a tuple with the NetworkIdentifier field value
// and a boolean to check if the value has been set.
func (o *RosettaAccountBalanceRequest) GetNetworkIdentifierOk() (*NetworkIdentifier, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkIdentifier, true
}

// SetNetworkIdentifier sets field value
func (o *RosettaAccountBalanceRequest) SetNetworkIdentifier(v NetworkIdentifier) {
	o.NetworkIdentifier = v
}

// GetAccountIdentifier returns the AccountIdentifier field value
func (o *RosettaAccountBalanceRequest) GetAccountIdentifier() RosettaAccount {
	if o == nil {
		var ret RosettaAccount
		return ret
	}

	return o.AccountIdentifier
}

// GetAccountIdentifierOk returns a tuple with the AccountIdentifier field value
// and a boolean to check if the value has been set.
func (o *RosettaAccountBalanceRequest) GetAccountIdentifierOk() (*RosettaAccount, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountIdentifier, true
}

// SetAccountIdentifier sets field value
func (o *RosettaAccountBalanceRequest) SetAccountIdentifier(v RosettaAccount) {
	o.AccountIdentifier = v
}

// GetBlockIdentifier returns the BlockIdentifier field value if set, zero value otherwise.
func (o *RosettaAccountBalanceRequest) GetBlockIdentifier() RosettaPartialBlockIdentifier {
	if o == nil || utils.IsNil(o.BlockIdentifier) {
		var ret RosettaPartialBlockIdentifier
		return ret
	}
	return *o.BlockIdentifier
}

// GetBlockIdentifierOk returns a tuple with the BlockIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaAccountBalanceRequest) GetBlockIdentifierOk() (*RosettaPartialBlockIdentifier, bool) {
	if o == nil || utils.IsNil(o.BlockIdentifier) {
		return nil, false
	}
	return o.BlockIdentifier, true
}

// HasBlockIdentifier returns a boolean if a field has been set.
func (o *RosettaAccountBalanceRequest) HasBlockIdentifier() bool {
	if o != nil && !utils.IsNil(o.BlockIdentifier) {
		return true
	}

	return false
}

// SetBlockIdentifier gets a reference to the given RosettaPartialBlockIdentifier and assigns it to the BlockIdentifier field.
func (o *RosettaAccountBalanceRequest) SetBlockIdentifier(v RosettaPartialBlockIdentifier) {
	o.BlockIdentifier = &v
}

func (o RosettaAccountBalanceRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RosettaAccountBalanceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["network_identifier"] = o.NetworkIdentifier
	toSerialize["account_identifier"] = o.AccountIdentifier
	if !utils.IsNil(o.BlockIdentifier) {
		toSerialize["block_identifier"] = o.BlockIdentifier
	}
	return toSerialize, nil
}

func (o *RosettaAccountBalanceRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"network_identifier",
		"account_identifier",
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

	varRosettaAccountBalanceRequest := _RosettaAccountBalanceRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRosettaAccountBalanceRequest)

	if err != nil {
		return err
	}

	*o = RosettaAccountBalanceRequest(varRosettaAccountBalanceRequest)

	return err
}

type NullableRosettaAccountBalanceRequest struct {
	value *RosettaAccountBalanceRequest
	isSet bool
}

func (v NullableRosettaAccountBalanceRequest) Get() *RosettaAccountBalanceRequest {
	return v.value
}

func (v *NullableRosettaAccountBalanceRequest) Set(val *RosettaAccountBalanceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRosettaAccountBalanceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRosettaAccountBalanceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRosettaAccountBalanceRequest(val *RosettaAccountBalanceRequest) *NullableRosettaAccountBalanceRequest {
	return &NullableRosettaAccountBalanceRequest{value: val, isSet: true}
}

func (v NullableRosettaAccountBalanceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRosettaAccountBalanceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
