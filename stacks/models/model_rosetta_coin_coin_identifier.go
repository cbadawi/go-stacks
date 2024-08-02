package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the RosettaCoinCoinIdentifier type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RosettaCoinCoinIdentifier{}

// RosettaCoinCoinIdentifier CoinIdentifier uniquely identifies a Coin.
type RosettaCoinCoinIdentifier struct {
	// Identifier should be populated with a globally unique identifier of a Coin. In Bitcoin, this identifier would be transaction_hash:index.
	Identifier string `json:"identifier"`
}

type _RosettaCoinCoinIdentifier RosettaCoinCoinIdentifier

// NewRosettaCoinCoinIdentifier instantiates a new RosettaCoinCoinIdentifier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRosettaCoinCoinIdentifier(identifier string) *RosettaCoinCoinIdentifier {
	this := RosettaCoinCoinIdentifier{}
	this.Identifier = identifier
	return &this
}

// NewRosettaCoinCoinIdentifierWithDefaults instantiates a new RosettaCoinCoinIdentifier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRosettaCoinCoinIdentifierWithDefaults() *RosettaCoinCoinIdentifier {
	this := RosettaCoinCoinIdentifier{}
	return &this
}

// GetIdentifier returns the Identifier field value
func (o *RosettaCoinCoinIdentifier) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *RosettaCoinCoinIdentifier) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value
func (o *RosettaCoinCoinIdentifier) SetIdentifier(v string) {
	o.Identifier = v
}

func (o RosettaCoinCoinIdentifier) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RosettaCoinCoinIdentifier) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["identifier"] = o.Identifier
	return toSerialize, nil
}

func (o *RosettaCoinCoinIdentifier) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"identifier",
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

	varRosettaCoinCoinIdentifier := _RosettaCoinCoinIdentifier{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRosettaCoinCoinIdentifier)

	if err != nil {
		return err
	}

	*o = RosettaCoinCoinIdentifier(varRosettaCoinCoinIdentifier)

	return err
}

type NullableRosettaCoinCoinIdentifier struct {
	value *RosettaCoinCoinIdentifier
	isSet bool
}

func (v NullableRosettaCoinCoinIdentifier) Get() *RosettaCoinCoinIdentifier {
	return v.value
}

func (v *NullableRosettaCoinCoinIdentifier) Set(val *RosettaCoinCoinIdentifier) {
	v.value = val
	v.isSet = true
}

func (v NullableRosettaCoinCoinIdentifier) IsSet() bool {
	return v.isSet
}

func (v *NullableRosettaCoinCoinIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRosettaCoinCoinIdentifier(val *RosettaCoinCoinIdentifier) *NullableRosettaCoinCoinIdentifier {
	return &NullableRosettaCoinCoinIdentifier{value: val, isSet: true}
}

func (v NullableRosettaCoinCoinIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRosettaCoinCoinIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
