package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the RosettaPublicKey type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RosettaPublicKey{}

// RosettaPublicKey PublicKey contains a public key byte array for a particular CurveType encoded in hex. Note that there is no PrivateKey struct as this is NEVER the concern of an implementation.
type RosettaPublicKey struct {
	// Hex-encoded public key bytes in the format specified by the CurveType.
	HexBytes string `json:"hex_bytes"`
	// CurveType is the type of cryptographic curve associated with a PublicKey.
	CurveType string `json:"curve_type"`
}

type _RosettaPublicKey RosettaPublicKey

// NewRosettaPublicKey instantiates a new RosettaPublicKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRosettaPublicKey(hexBytes string, curveType string) *RosettaPublicKey {
	this := RosettaPublicKey{}
	this.HexBytes = hexBytes
	this.CurveType = curveType
	return &this
}

// NewRosettaPublicKeyWithDefaults instantiates a new RosettaPublicKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRosettaPublicKeyWithDefaults() *RosettaPublicKey {
	this := RosettaPublicKey{}
	return &this
}

// GetHexBytes returns the HexBytes field value
func (o *RosettaPublicKey) GetHexBytes() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HexBytes
}

// GetHexBytesOk returns a tuple with the HexBytes field value
// and a boolean to check if the value has been set.
func (o *RosettaPublicKey) GetHexBytesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HexBytes, true
}

// SetHexBytes sets field value
func (o *RosettaPublicKey) SetHexBytes(v string) {
	o.HexBytes = v
}

// GetCurveType returns the CurveType field value
func (o *RosettaPublicKey) GetCurveType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurveType
}

// GetCurveTypeOk returns a tuple with the CurveType field value
// and a boolean to check if the value has been set.
func (o *RosettaPublicKey) GetCurveTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurveType, true
}

// SetCurveType sets field value
func (o *RosettaPublicKey) SetCurveType(v string) {
	o.CurveType = v
}

func (o RosettaPublicKey) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RosettaPublicKey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["hex_bytes"] = o.HexBytes
	toSerialize["curve_type"] = o.CurveType
	return toSerialize, nil
}

func (o *RosettaPublicKey) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"hex_bytes",
		"curve_type",
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

	varRosettaPublicKey := _RosettaPublicKey{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRosettaPublicKey)

	if err != nil {
		return err
	}

	*o = RosettaPublicKey(varRosettaPublicKey)

	return err
}

type NullableRosettaPublicKey struct {
	value *RosettaPublicKey
	isSet bool
}

func (v NullableRosettaPublicKey) Get() *RosettaPublicKey {
	return v.value
}

func (v *NullableRosettaPublicKey) Set(val *RosettaPublicKey) {
	v.value = val
	v.isSet = true
}

func (v NullableRosettaPublicKey) IsSet() bool {
	return v.isSet
}

func (v *NullableRosettaPublicKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRosettaPublicKey(val *RosettaPublicKey) *NullableRosettaPublicKey {
	return &NullableRosettaPublicKey{value: val, isSet: true}
}

func (v NullableRosettaPublicKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRosettaPublicKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
