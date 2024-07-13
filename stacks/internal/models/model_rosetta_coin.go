package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the RosettaCoin type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RosettaCoin{}

// RosettaCoin If a blockchain is UTXO-based, all unspent Coins owned by an account_identifier should be returned alongside the balance. It is highly recommended to populate this field so that users of the Rosetta API implementation don't need to maintain their own indexer to track their UTXOs.
type RosettaCoin struct {
	CoinIdentifier RosettaCoinCoinIdentifier `json:"coin_identifier"`
	Amount         RosettaAmount             `json:"amount"`
}

type _RosettaCoin RosettaCoin

// NewRosettaCoin instantiates a new RosettaCoin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRosettaCoin(coinIdentifier RosettaCoinCoinIdentifier, amount RosettaAmount) *RosettaCoin {
	this := RosettaCoin{}
	this.CoinIdentifier = coinIdentifier
	this.Amount = amount
	return &this
}

// NewRosettaCoinWithDefaults instantiates a new RosettaCoin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRosettaCoinWithDefaults() *RosettaCoin {
	this := RosettaCoin{}
	return &this
}

// GetCoinIdentifier returns the CoinIdentifier field value
func (o *RosettaCoin) GetCoinIdentifier() RosettaCoinCoinIdentifier {
	if o == nil {
		var ret RosettaCoinCoinIdentifier
		return ret
	}

	return o.CoinIdentifier
}

// GetCoinIdentifierOk returns a tuple with the CoinIdentifier field value
// and a boolean to check if the value has been set.
func (o *RosettaCoin) GetCoinIdentifierOk() (*RosettaCoinCoinIdentifier, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CoinIdentifier, true
}

// SetCoinIdentifier sets field value
func (o *RosettaCoin) SetCoinIdentifier(v RosettaCoinCoinIdentifier) {
	o.CoinIdentifier = v
}

// GetAmount returns the Amount field value
func (o *RosettaCoin) GetAmount() RosettaAmount {
	if o == nil {
		var ret RosettaAmount
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *RosettaCoin) GetAmountOk() (*RosettaAmount, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *RosettaCoin) SetAmount(v RosettaAmount) {
	o.Amount = v
}

func (o RosettaCoin) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RosettaCoin) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["coin_identifier"] = o.CoinIdentifier
	toSerialize["amount"] = o.Amount
	return toSerialize, nil
}

func (o *RosettaCoin) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"coin_identifier",
		"amount",
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

	varRosettaCoin := _RosettaCoin{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRosettaCoin)

	if err != nil {
		return err
	}

	*o = RosettaCoin(varRosettaCoin)

	return err
}

type NullableRosettaCoin struct {
	value *RosettaCoin
	isSet bool
}

func (v NullableRosettaCoin) Get() *RosettaCoin {
	return v.value
}

func (v *NullableRosettaCoin) Set(val *RosettaCoin) {
	v.value = val
	v.isSet = true
}

func (v NullableRosettaCoin) IsSet() bool {
	return v.isSet
}

func (v *NullableRosettaCoin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRosettaCoin(val *RosettaCoin) *NullableRosettaCoin {
	return &NullableRosettaCoin{value: val, isSet: true}
}

func (v NullableRosettaCoin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRosettaCoin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
