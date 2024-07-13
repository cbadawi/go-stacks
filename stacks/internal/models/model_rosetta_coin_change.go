package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the RosettaCoinChange type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RosettaCoinChange{}

// RosettaCoinChange CoinChange is used to represent a change in state of a some coin identified by a coin_identifier. This object is part of the Operation model and must be populated for UTXO-based blockchains. Coincidentally, this abstraction of UTXOs allows for supporting both account-based transfers and UTXO-based transfers on the same blockchain (when a transfer is account-based, don't populate this model).
type RosettaCoinChange struct {
	CoinIdentifier RosettaCoinChangeCoinIdentifier `json:"coin_identifier"`
	// CoinActions are different state changes that a Coin can undergo. When a Coin is created, it is coin_created. When a Coin is spent, it is coin_spent. It is assumed that a single Coin cannot be created or spent more than once.
	CoinAction string `json:"coin_action"`
}

type _RosettaCoinChange RosettaCoinChange

// NewRosettaCoinChange instantiates a new RosettaCoinChange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRosettaCoinChange(coinIdentifier RosettaCoinChangeCoinIdentifier, coinAction string) *RosettaCoinChange {
	this := RosettaCoinChange{}
	this.CoinIdentifier = coinIdentifier
	this.CoinAction = coinAction
	return &this
}

// NewRosettaCoinChangeWithDefaults instantiates a new RosettaCoinChange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRosettaCoinChangeWithDefaults() *RosettaCoinChange {
	this := RosettaCoinChange{}
	return &this
}

// GetCoinIdentifier returns the CoinIdentifier field value
func (o *RosettaCoinChange) GetCoinIdentifier() RosettaCoinChangeCoinIdentifier {
	if o == nil {
		var ret RosettaCoinChangeCoinIdentifier
		return ret
	}

	return o.CoinIdentifier
}

// GetCoinIdentifierOk returns a tuple with the CoinIdentifier field value
// and a boolean to check if the value has been set.
func (o *RosettaCoinChange) GetCoinIdentifierOk() (*RosettaCoinChangeCoinIdentifier, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CoinIdentifier, true
}

// SetCoinIdentifier sets field value
func (o *RosettaCoinChange) SetCoinIdentifier(v RosettaCoinChangeCoinIdentifier) {
	o.CoinIdentifier = v
}

// GetCoinAction returns the CoinAction field value
func (o *RosettaCoinChange) GetCoinAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CoinAction
}

// GetCoinActionOk returns a tuple with the CoinAction field value
// and a boolean to check if the value has been set.
func (o *RosettaCoinChange) GetCoinActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CoinAction, true
}

// SetCoinAction sets field value
func (o *RosettaCoinChange) SetCoinAction(v string) {
	o.CoinAction = v
}

func (o RosettaCoinChange) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RosettaCoinChange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["coin_identifier"] = o.CoinIdentifier
	toSerialize["coin_action"] = o.CoinAction
	return toSerialize, nil
}

func (o *RosettaCoinChange) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"coin_identifier",
		"coin_action",
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

	varRosettaCoinChange := _RosettaCoinChange{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRosettaCoinChange)

	if err != nil {
		return err
	}

	*o = RosettaCoinChange(varRosettaCoinChange)

	return err
}

type NullableRosettaCoinChange struct {
	value *RosettaCoinChange
	isSet bool
}

func (v NullableRosettaCoinChange) Get() *RosettaCoinChange {
	return v.value
}

func (v *NullableRosettaCoinChange) Set(val *RosettaCoinChange) {
	v.value = val
	v.isSet = true
}

func (v NullableRosettaCoinChange) IsSet() bool {
	return v.isSet
}

func (v *NullableRosettaCoinChange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRosettaCoinChange(val *RosettaCoinChange) *NullableRosettaCoinChange {
	return &NullableRosettaCoinChange{value: val, isSet: true}
}

func (v NullableRosettaCoinChange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRosettaCoinChange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
