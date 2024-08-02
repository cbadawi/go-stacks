package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the AddressTransactionEventsStx type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AddressTransactionEventsStx{}

// AddressTransactionEventsStx struct for AddressTransactionEventsStx
type AddressTransactionEventsStx struct {
	Transfer int32 `json:"transfer"`
	Mint     int32 `json:"mint"`
	Burn     int32 `json:"burn"`
}

type _AddressTransactionEventsStx AddressTransactionEventsStx

// NewAddressTransactionEventsStx instantiates a new AddressTransactionEventsStx object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressTransactionEventsStx(transfer int32, mint int32, burn int32) *AddressTransactionEventsStx {
	this := AddressTransactionEventsStx{}
	this.Transfer = transfer
	this.Mint = mint
	this.Burn = burn
	return &this
}

// NewAddressTransactionEventsStxWithDefaults instantiates a new AddressTransactionEventsStx object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressTransactionEventsStxWithDefaults() *AddressTransactionEventsStx {
	this := AddressTransactionEventsStx{}
	return &this
}

// GetTransfer returns the Transfer field value
func (o *AddressTransactionEventsStx) GetTransfer() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Transfer
}

// GetTransferOk returns a tuple with the Transfer field value
// and a boolean to check if the value has been set.
func (o *AddressTransactionEventsStx) GetTransferOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Transfer, true
}

// SetTransfer sets field value
func (o *AddressTransactionEventsStx) SetTransfer(v int32) {
	o.Transfer = v
}

// GetMint returns the Mint field value
func (o *AddressTransactionEventsStx) GetMint() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Mint
}

// GetMintOk returns a tuple with the Mint field value
// and a boolean to check if the value has been set.
func (o *AddressTransactionEventsStx) GetMintOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mint, true
}

// SetMint sets field value
func (o *AddressTransactionEventsStx) SetMint(v int32) {
	o.Mint = v
}

// GetBurn returns the Burn field value
func (o *AddressTransactionEventsStx) GetBurn() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Burn
}

// GetBurnOk returns a tuple with the Burn field value
// and a boolean to check if the value has been set.
func (o *AddressTransactionEventsStx) GetBurnOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Burn, true
}

// SetBurn sets field value
func (o *AddressTransactionEventsStx) SetBurn(v int32) {
	o.Burn = v
}

func (o AddressTransactionEventsStx) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressTransactionEventsStx) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["transfer"] = o.Transfer
	toSerialize["mint"] = o.Mint
	toSerialize["burn"] = o.Burn
	return toSerialize, nil
}

func (o *AddressTransactionEventsStx) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"transfer",
		"mint",
		"burn",
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

	varAddressTransactionEventsStx := _AddressTransactionEventsStx{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddressTransactionEventsStx)

	if err != nil {
		return err
	}

	*o = AddressTransactionEventsStx(varAddressTransactionEventsStx)

	return err
}

type NullableAddressTransactionEventsStx struct {
	value *AddressTransactionEventsStx
	isSet bool
}

func (v NullableAddressTransactionEventsStx) Get() *AddressTransactionEventsStx {
	return v.value
}

func (v *NullableAddressTransactionEventsStx) Set(val *AddressTransactionEventsStx) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressTransactionEventsStx) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressTransactionEventsStx) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressTransactionEventsStx(val *AddressTransactionEventsStx) *NullableAddressTransactionEventsStx {
	return &NullableAddressTransactionEventsStx{value: val, isSet: true}
}

func (v NullableAddressTransactionEventsStx) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressTransactionEventsStx) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
