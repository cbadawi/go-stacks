package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the CoinbaseTransactionMetadataCoinbasePayload type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CoinbaseTransactionMetadataCoinbasePayload{}

// CoinbaseTransactionMetadataCoinbasePayload struct for CoinbaseTransactionMetadataCoinbasePayload
type CoinbaseTransactionMetadataCoinbasePayload struct {
	// Hex encoded 32-byte scratch space for block leader's use
	Data string `json:"data"`
	// A principal that will receive the miner rewards for this coinbase transaction. Can be either a standard principal or contract principal. Only specified for `coinbase-to-alt-recipient` transaction types, otherwise null.
	AltRecipient utils.NullableString `json:"alt_recipient,omitempty"`
	// Hex encoded 80-byte VRF proof
	VrfProof utils.NullableString `json:"vrf_proof,omitempty"`
}

type _CoinbaseTransactionMetadataCoinbasePayload CoinbaseTransactionMetadataCoinbasePayload

// NewCoinbaseTransactionMetadataCoinbasePayload instantiates a new CoinbaseTransactionMetadataCoinbasePayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoinbaseTransactionMetadataCoinbasePayload(data string) *CoinbaseTransactionMetadataCoinbasePayload {
	this := CoinbaseTransactionMetadataCoinbasePayload{}
	this.Data = data
	return &this
}

// NewCoinbaseTransactionMetadataCoinbasePayloadWithDefaults instantiates a new CoinbaseTransactionMetadataCoinbasePayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoinbaseTransactionMetadataCoinbasePayloadWithDefaults() *CoinbaseTransactionMetadataCoinbasePayload {
	this := CoinbaseTransactionMetadataCoinbasePayload{}
	return &this
}

// GetData returns the Data field value
func (o *CoinbaseTransactionMetadataCoinbasePayload) GetData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CoinbaseTransactionMetadataCoinbasePayload) GetDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CoinbaseTransactionMetadataCoinbasePayload) SetData(v string) {
	o.Data = v
}

// GetAltRecipient returns the AltRecipient field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CoinbaseTransactionMetadataCoinbasePayload) GetAltRecipient() string {
	if o == nil || utils.IsNil(o.AltRecipient.Get()) {
		var ret string
		return ret
	}
	return *o.AltRecipient.Get()
}

// GetAltRecipientOk returns a tuple with the AltRecipient field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CoinbaseTransactionMetadataCoinbasePayload) GetAltRecipientOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AltRecipient.Get(), o.AltRecipient.IsSet()
}

// HasAltRecipient returns a boolean if a field has been set.
func (o *CoinbaseTransactionMetadataCoinbasePayload) HasAltRecipient() bool {
	if o != nil && o.AltRecipient.IsSet() {
		return true
	}

	return false
}

// SetAltRecipient gets a reference to the given NullableString and assigns it to the AltRecipient field.
func (o *CoinbaseTransactionMetadataCoinbasePayload) SetAltRecipient(v string) {
	o.AltRecipient.Set(&v)
}

// SetAltRecipientNil sets the value for AltRecipient to be an explicit nil
func (o *CoinbaseTransactionMetadataCoinbasePayload) SetAltRecipientNil() {
	o.AltRecipient.Set(nil)
}

// UnsetAltRecipient ensures that no value is present for AltRecipient, not even an explicit nil
func (o *CoinbaseTransactionMetadataCoinbasePayload) UnsetAltRecipient() {
	o.AltRecipient.Unset()
}

// GetVrfProof returns the VrfProof field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CoinbaseTransactionMetadataCoinbasePayload) GetVrfProof() string {
	if o == nil || utils.IsNil(o.VrfProof.Get()) {
		var ret string
		return ret
	}
	return *o.VrfProof.Get()
}

// GetVrfProofOk returns a tuple with the VrfProof field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CoinbaseTransactionMetadataCoinbasePayload) GetVrfProofOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VrfProof.Get(), o.VrfProof.IsSet()
}

// HasVrfProof returns a boolean if a field has been set.
func (o *CoinbaseTransactionMetadataCoinbasePayload) HasVrfProof() bool {
	if o != nil && o.VrfProof.IsSet() {
		return true
	}

	return false
}

// SetVrfProof gets a reference to the given NullableString and assigns it to the VrfProof field.
func (o *CoinbaseTransactionMetadataCoinbasePayload) SetVrfProof(v string) {
	o.VrfProof.Set(&v)
}

// SetVrfProofNil sets the value for VrfProof to be an explicit nil
func (o *CoinbaseTransactionMetadataCoinbasePayload) SetVrfProofNil() {
	o.VrfProof.Set(nil)
}

// UnsetVrfProof ensures that no value is present for VrfProof, not even an explicit nil
func (o *CoinbaseTransactionMetadataCoinbasePayload) UnsetVrfProof() {
	o.VrfProof.Unset()
}

func (o CoinbaseTransactionMetadataCoinbasePayload) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoinbaseTransactionMetadataCoinbasePayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if o.AltRecipient.IsSet() {
		toSerialize["alt_recipient"] = o.AltRecipient.Get()
	}
	if o.VrfProof.IsSet() {
		toSerialize["vrf_proof"] = o.VrfProof.Get()
	}
	return toSerialize, nil
}

func (o *CoinbaseTransactionMetadataCoinbasePayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varCoinbaseTransactionMetadataCoinbasePayload := _CoinbaseTransactionMetadataCoinbasePayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoinbaseTransactionMetadataCoinbasePayload)

	if err != nil {
		return err
	}

	*o = CoinbaseTransactionMetadataCoinbasePayload(varCoinbaseTransactionMetadataCoinbasePayload)

	return err
}

type NullableCoinbaseTransactionMetadataCoinbasePayload struct {
	value *CoinbaseTransactionMetadataCoinbasePayload
	isSet bool
}

func (v NullableCoinbaseTransactionMetadataCoinbasePayload) Get() *CoinbaseTransactionMetadataCoinbasePayload {
	return v.value
}

func (v *NullableCoinbaseTransactionMetadataCoinbasePayload) Set(val *CoinbaseTransactionMetadataCoinbasePayload) {
	v.value = val
	v.isSet = true
}

func (v NullableCoinbaseTransactionMetadataCoinbasePayload) IsSet() bool {
	return v.isSet
}

func (v *NullableCoinbaseTransactionMetadataCoinbasePayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoinbaseTransactionMetadataCoinbasePayload(val *CoinbaseTransactionMetadataCoinbasePayload) *NullableCoinbaseTransactionMetadataCoinbasePayload {
	return &NullableCoinbaseTransactionMetadataCoinbasePayload{value: val, isSet: true}
}

func (v NullableCoinbaseTransactionMetadataCoinbasePayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoinbaseTransactionMetadataCoinbasePayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
