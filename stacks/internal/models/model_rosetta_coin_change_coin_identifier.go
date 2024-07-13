package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the RosettaCoinChangeCoinIdentifier type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RosettaCoinChangeCoinIdentifier{}

// RosettaCoinChangeCoinIdentifier CoinIdentifier uniquely identifies a Coin.
type RosettaCoinChangeCoinIdentifier struct {
	// Identifier should be populated with a globally unique identifier of a Coin. In Bitcoin, this identifier would be transaction_hash:index.
	Identifier string `json:"identifier"`
}

type _RosettaCoinChangeCoinIdentifier RosettaCoinChangeCoinIdentifier

// NewRosettaCoinChangeCoinIdentifier instantiates a new RosettaCoinChangeCoinIdentifier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRosettaCoinChangeCoinIdentifier(identifier string) *RosettaCoinChangeCoinIdentifier {
	this := RosettaCoinChangeCoinIdentifier{}
	this.Identifier = identifier
	return &this
}

// NewRosettaCoinChangeCoinIdentifierWithDefaults instantiates a new RosettaCoinChangeCoinIdentifier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRosettaCoinChangeCoinIdentifierWithDefaults() *RosettaCoinChangeCoinIdentifier {
	this := RosettaCoinChangeCoinIdentifier{}
	return &this
}

// GetIdentifier returns the Identifier field value
func (o *RosettaCoinChangeCoinIdentifier) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *RosettaCoinChangeCoinIdentifier) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value
func (o *RosettaCoinChangeCoinIdentifier) SetIdentifier(v string) {
	o.Identifier = v
}

func (o RosettaCoinChangeCoinIdentifier) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RosettaCoinChangeCoinIdentifier) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["identifier"] = o.Identifier
	return toSerialize, nil
}

func (o *RosettaCoinChangeCoinIdentifier) UnmarshalJSON(data []byte) (err error) {
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

	varRosettaCoinChangeCoinIdentifier := _RosettaCoinChangeCoinIdentifier{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRosettaCoinChangeCoinIdentifier)

	if err != nil {
		return err
	}

	*o = RosettaCoinChangeCoinIdentifier(varRosettaCoinChangeCoinIdentifier)

	return err
}

type NullableRosettaCoinChangeCoinIdentifier struct {
	value *RosettaCoinChangeCoinIdentifier
	isSet bool
}

func (v NullableRosettaCoinChangeCoinIdentifier) Get() *RosettaCoinChangeCoinIdentifier {
	return v.value
}

func (v *NullableRosettaCoinChangeCoinIdentifier) Set(val *RosettaCoinChangeCoinIdentifier) {
	v.value = val
	v.isSet = true
}

func (v NullableRosettaCoinChangeCoinIdentifier) IsSet() bool {
	return v.isSet
}

func (v *NullableRosettaCoinChangeCoinIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRosettaCoinChangeCoinIdentifier(val *RosettaCoinChangeCoinIdentifier) *NullableRosettaCoinChangeCoinIdentifier {
	return &NullableRosettaCoinChangeCoinIdentifier{value: val, isSet: true}
}

func (v NullableRosettaCoinChangeCoinIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRosettaCoinChangeCoinIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
