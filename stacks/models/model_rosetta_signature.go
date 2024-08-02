package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the RosettaSignature type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RosettaSignature{}

// RosettaSignature Signature contains the payload that was signed, the public keys of the keypairs used to produce the signature, the signature (encoded in hex), and the SignatureType. PublicKey is often times not known during construction of the signing payloads but may be needed to combine signatures properly.
type RosettaSignature struct {
	SigningPayload SigningPayload   `json:"signing_payload"`
	PublicKey      RosettaPublicKey `json:"public_key"`
	// SignatureType is the type of a cryptographic signature.
	SignatureType string `json:"signature_type"`
	HexBytes      string `json:"hex_bytes"`
}

type _RosettaSignature RosettaSignature

// NewRosettaSignature instantiates a new RosettaSignature object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRosettaSignature(signingPayload SigningPayload, publicKey RosettaPublicKey, signatureType string, hexBytes string) *RosettaSignature {
	this := RosettaSignature{}
	this.SigningPayload = signingPayload
	this.PublicKey = publicKey
	this.SignatureType = signatureType
	this.HexBytes = hexBytes
	return &this
}

// NewRosettaSignatureWithDefaults instantiates a new RosettaSignature object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRosettaSignatureWithDefaults() *RosettaSignature {
	this := RosettaSignature{}
	return &this
}

// GetSigningPayload returns the SigningPayload field value
func (o *RosettaSignature) GetSigningPayload() SigningPayload {
	if o == nil {
		var ret SigningPayload
		return ret
	}

	return o.SigningPayload
}

// GetSigningPayloadOk returns a tuple with the SigningPayload field value
// and a boolean to check if the value has been set.
func (o *RosettaSignature) GetSigningPayloadOk() (*SigningPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SigningPayload, true
}

// SetSigningPayload sets field value
func (o *RosettaSignature) SetSigningPayload(v SigningPayload) {
	o.SigningPayload = v
}

// GetPublicKey returns the PublicKey field value
func (o *RosettaSignature) GetPublicKey() RosettaPublicKey {
	if o == nil {
		var ret RosettaPublicKey
		return ret
	}

	return o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value
// and a boolean to check if the value has been set.
func (o *RosettaSignature) GetPublicKeyOk() (*RosettaPublicKey, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicKey, true
}

// SetPublicKey sets field value
func (o *RosettaSignature) SetPublicKey(v RosettaPublicKey) {
	o.PublicKey = v
}

// GetSignatureType returns the SignatureType field value
func (o *RosettaSignature) GetSignatureType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SignatureType
}

// GetSignatureTypeOk returns a tuple with the SignatureType field value
// and a boolean to check if the value has been set.
func (o *RosettaSignature) GetSignatureTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignatureType, true
}

// SetSignatureType sets field value
func (o *RosettaSignature) SetSignatureType(v string) {
	o.SignatureType = v
}

// GetHexBytes returns the HexBytes field value
func (o *RosettaSignature) GetHexBytes() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HexBytes
}

// GetHexBytesOk returns a tuple with the HexBytes field value
// and a boolean to check if the value has been set.
func (o *RosettaSignature) GetHexBytesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HexBytes, true
}

// SetHexBytes sets field value
func (o *RosettaSignature) SetHexBytes(v string) {
	o.HexBytes = v
}

func (o RosettaSignature) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RosettaSignature) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["signing_payload"] = o.SigningPayload
	toSerialize["public_key"] = o.PublicKey
	toSerialize["signature_type"] = o.SignatureType
	toSerialize["hex_bytes"] = o.HexBytes
	return toSerialize, nil
}

func (o *RosettaSignature) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"signing_payload",
		"public_key",
		"signature_type",
		"hex_bytes",
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

	varRosettaSignature := _RosettaSignature{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRosettaSignature)

	if err != nil {
		return err
	}

	*o = RosettaSignature(varRosettaSignature)

	return err
}

type NullableRosettaSignature struct {
	value *RosettaSignature
	isSet bool
}

func (v NullableRosettaSignature) Get() *RosettaSignature {
	return v.value
}

func (v *NullableRosettaSignature) Set(val *RosettaSignature) {
	v.value = val
	v.isSet = true
}

func (v NullableRosettaSignature) IsSet() bool {
	return v.isSet
}

func (v *NullableRosettaSignature) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRosettaSignature(val *RosettaSignature) *NullableRosettaSignature {
	return &NullableRosettaSignature{value: val, isSet: true}
}

func (v NullableRosettaSignature) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRosettaSignature) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
