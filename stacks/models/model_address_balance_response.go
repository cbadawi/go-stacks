package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the AddressBalanceResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AddressBalanceResponse{}

// AddressBalanceResponse GET request that returns address balances
type AddressBalanceResponse struct {
	Stx                 StxBalance                                              `json:"stx"`
	FungibleTokens      map[string]AddressBalanceResponseFungibleTokensValue    `json:"fungible_tokens"`
	NonFungibleTokens   map[string]AddressBalanceResponseNonFungibleTokensValue `json:"non_fungible_tokens"`
	TokenOfferingLocked *AddressTokenOfferingLocked                             `json:"token_offering_locked,omitempty"`
}

type _AddressBalanceResponse AddressBalanceResponse

// NewAddressBalanceResponse instantiates a new AddressBalanceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressBalanceResponse(stx StxBalance, fungibleTokens map[string]AddressBalanceResponseFungibleTokensValue, nonFungibleTokens map[string]AddressBalanceResponseNonFungibleTokensValue) *AddressBalanceResponse {
	this := AddressBalanceResponse{}
	this.Stx = stx
	this.FungibleTokens = fungibleTokens
	this.NonFungibleTokens = nonFungibleTokens
	return &this
}

// NewAddressBalanceResponseWithDefaults instantiates a new AddressBalanceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressBalanceResponseWithDefaults() *AddressBalanceResponse {
	this := AddressBalanceResponse{}
	return &this
}

// GetStx returns the Stx field value
func (o *AddressBalanceResponse) GetStx() StxBalance {
	if o == nil {
		var ret StxBalance
		return ret
	}

	return o.Stx
}

// GetStxOk returns a tuple with the Stx field value
// and a boolean to check if the value has been set.
func (o *AddressBalanceResponse) GetStxOk() (*StxBalance, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stx, true
}

// SetStx sets field value
func (o *AddressBalanceResponse) SetStx(v StxBalance) {
	o.Stx = v
}

// GetFungibleTokens returns the FungibleTokens field value
func (o *AddressBalanceResponse) GetFungibleTokens() map[string]AddressBalanceResponseFungibleTokensValue {
	if o == nil {
		var ret map[string]AddressBalanceResponseFungibleTokensValue
		return ret
	}

	return o.FungibleTokens
}

// GetFungibleTokensOk returns a tuple with the FungibleTokens field value
// and a boolean to check if the value has been set.
func (o *AddressBalanceResponse) GetFungibleTokensOk() (*map[string]AddressBalanceResponseFungibleTokensValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FungibleTokens, true
}

// SetFungibleTokens sets field value
func (o *AddressBalanceResponse) SetFungibleTokens(v map[string]AddressBalanceResponseFungibleTokensValue) {
	o.FungibleTokens = v
}

// GetNonFungibleTokens returns the NonFungibleTokens field value
func (o *AddressBalanceResponse) GetNonFungibleTokens() map[string]AddressBalanceResponseNonFungibleTokensValue {
	if o == nil {
		var ret map[string]AddressBalanceResponseNonFungibleTokensValue
		return ret
	}

	return o.NonFungibleTokens
}

// GetNonFungibleTokensOk returns a tuple with the NonFungibleTokens field value
// and a boolean to check if the value has been set.
func (o *AddressBalanceResponse) GetNonFungibleTokensOk() (*map[string]AddressBalanceResponseNonFungibleTokensValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NonFungibleTokens, true
}

// SetNonFungibleTokens sets field value
func (o *AddressBalanceResponse) SetNonFungibleTokens(v map[string]AddressBalanceResponseNonFungibleTokensValue) {
	o.NonFungibleTokens = v
}

// GetTokenOfferingLocked returns the TokenOfferingLocked field value if set, zero value otherwise.
func (o *AddressBalanceResponse) GetTokenOfferingLocked() AddressTokenOfferingLocked {
	if o == nil || utils.IsNil(o.TokenOfferingLocked) {
		var ret AddressTokenOfferingLocked
		return ret
	}
	return *o.TokenOfferingLocked
}

// GetTokenOfferingLockedOk returns a tuple with the TokenOfferingLocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressBalanceResponse) GetTokenOfferingLockedOk() (*AddressTokenOfferingLocked, bool) {
	if o == nil || utils.IsNil(o.TokenOfferingLocked) {
		return nil, false
	}
	return o.TokenOfferingLocked, true
}

// HasTokenOfferingLocked returns a boolean if a field has been set.
func (o *AddressBalanceResponse) HasTokenOfferingLocked() bool {
	if o != nil && !utils.IsNil(o.TokenOfferingLocked) {
		return true
	}

	return false
}

// SetTokenOfferingLocked gets a reference to the given AddressTokenOfferingLocked and assigns it to the TokenOfferingLocked field.
func (o *AddressBalanceResponse) SetTokenOfferingLocked(v AddressTokenOfferingLocked) {
	o.TokenOfferingLocked = &v
}

func (o AddressBalanceResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressBalanceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["stx"] = o.Stx
	toSerialize["fungible_tokens"] = o.FungibleTokens
	toSerialize["non_fungible_tokens"] = o.NonFungibleTokens
	if !utils.IsNil(o.TokenOfferingLocked) {
		toSerialize["token_offering_locked"] = o.TokenOfferingLocked
	}
	return toSerialize, nil
}

func (o *AddressBalanceResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"stx",
		"fungible_tokens",
		"non_fungible_tokens",
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

	varAddressBalanceResponse := _AddressBalanceResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddressBalanceResponse)

	if err != nil {
		return err
	}

	*o = AddressBalanceResponse(varAddressBalanceResponse)

	return err
}

type NullableAddressBalanceResponse struct {
	value *AddressBalanceResponse
	isSet bool
}

func (v NullableAddressBalanceResponse) Get() *AddressBalanceResponse {
	return v.value
}

func (v *NullableAddressBalanceResponse) Set(val *AddressBalanceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressBalanceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressBalanceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressBalanceResponse(val *AddressBalanceResponse) *NullableAddressBalanceResponse {
	return &NullableAddressBalanceResponse{value: val, isSet: true}
}

func (v NullableAddressBalanceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressBalanceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
