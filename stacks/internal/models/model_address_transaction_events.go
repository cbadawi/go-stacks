package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the AddressTransactionEvents type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AddressTransactionEvents{}

// AddressTransactionEvents struct for AddressTransactionEvents
type AddressTransactionEvents struct {
	Stx AddressTransactionEventsStx `json:"stx"`
	Ft  AddressTransactionEventsStx `json:"ft"`
	Nft AddressTransactionEventsStx `json:"nft"`
}

type _AddressTransactionEvents AddressTransactionEvents

// NewAddressTransactionEvents instantiates a new AddressTransactionEvents object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressTransactionEvents(stx AddressTransactionEventsStx, ft AddressTransactionEventsStx, nft AddressTransactionEventsStx) *AddressTransactionEvents {
	this := AddressTransactionEvents{}
	this.Stx = stx
	this.Ft = ft
	this.Nft = nft
	return &this
}

// NewAddressTransactionEventsWithDefaults instantiates a new AddressTransactionEvents object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressTransactionEventsWithDefaults() *AddressTransactionEvents {
	this := AddressTransactionEvents{}
	return &this
}

// GetStx returns the Stx field value
func (o *AddressTransactionEvents) GetStx() AddressTransactionEventsStx {
	if o == nil {
		var ret AddressTransactionEventsStx
		return ret
	}

	return o.Stx
}

// GetStxOk returns a tuple with the Stx field value
// and a boolean to check if the value has been set.
func (o *AddressTransactionEvents) GetStxOk() (*AddressTransactionEventsStx, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stx, true
}

// SetStx sets field value
func (o *AddressTransactionEvents) SetStx(v AddressTransactionEventsStx) {
	o.Stx = v
}

// GetFt returns the Ft field value
func (o *AddressTransactionEvents) GetFt() AddressTransactionEventsStx {
	if o == nil {
		var ret AddressTransactionEventsStx
		return ret
	}

	return o.Ft
}

// GetFtOk returns a tuple with the Ft field value
// and a boolean to check if the value has been set.
func (o *AddressTransactionEvents) GetFtOk() (*AddressTransactionEventsStx, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ft, true
}

// SetFt sets field value
func (o *AddressTransactionEvents) SetFt(v AddressTransactionEventsStx) {
	o.Ft = v
}

// GetNft returns the Nft field value
func (o *AddressTransactionEvents) GetNft() AddressTransactionEventsStx {
	if o == nil {
		var ret AddressTransactionEventsStx
		return ret
	}

	return o.Nft
}

// GetNftOk returns a tuple with the Nft field value
// and a boolean to check if the value has been set.
func (o *AddressTransactionEvents) GetNftOk() (*AddressTransactionEventsStx, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nft, true
}

// SetNft sets field value
func (o *AddressTransactionEvents) SetNft(v AddressTransactionEventsStx) {
	o.Nft = v
}

func (o AddressTransactionEvents) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressTransactionEvents) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["stx"] = o.Stx
	toSerialize["ft"] = o.Ft
	toSerialize["nft"] = o.Nft
	return toSerialize, nil
}

func (o *AddressTransactionEvents) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"stx",
		"ft",
		"nft",
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

	varAddressTransactionEvents := _AddressTransactionEvents{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddressTransactionEvents)

	if err != nil {
		return err
	}

	*o = AddressTransactionEvents(varAddressTransactionEvents)

	return err
}

type NullableAddressTransactionEvents struct {
	value *AddressTransactionEvents
	isSet bool
}

func (v NullableAddressTransactionEvents) Get() *AddressTransactionEvents {
	return v.value
}

func (v *NullableAddressTransactionEvents) Set(val *AddressTransactionEvents) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressTransactionEvents) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressTransactionEvents) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressTransactionEvents(val *AddressTransactionEvents) *NullableAddressTransactionEvents {
	return &NullableAddressTransactionEvents{value: val, isSet: true}
}

func (v NullableAddressTransactionEvents) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressTransactionEvents) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
