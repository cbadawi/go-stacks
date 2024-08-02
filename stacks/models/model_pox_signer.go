package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the PoxSigner type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PoxSigner{}

// PoxSigner struct for PoxSigner
type PoxSigner struct {
	SigningKey string `json:"signing_key"`
	// The Stacks address derived from the signing_key.
	SignerAddress        string  `json:"signer_address"`
	Weight               int32   `json:"weight"`
	StackedAmount        string  `json:"stacked_amount"`
	WeightPercent        float32 `json:"weight_percent"`
	StackedAmountPercent float32 `json:"stacked_amount_percent"`
	// The number of solo stackers associated with this signer.
	SoloStackerCount int32 `json:"solo_stacker_count"`
	// The number of pooled stackers associated with this signer.
	PooledStackerCount int32 `json:"pooled_stacker_count"`
}

type _PoxSigner PoxSigner

// NewPoxSigner instantiates a new PoxSigner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoxSigner(signingKey string, signerAddress string, weight int32, stackedAmount string, weightPercent float32, stackedAmountPercent float32, soloStackerCount int32, pooledStackerCount int32) *PoxSigner {
	this := PoxSigner{}
	this.SigningKey = signingKey
	this.SignerAddress = signerAddress
	this.Weight = weight
	this.StackedAmount = stackedAmount
	this.WeightPercent = weightPercent
	this.StackedAmountPercent = stackedAmountPercent
	this.SoloStackerCount = soloStackerCount
	this.PooledStackerCount = pooledStackerCount
	return &this
}

// NewPoxSignerWithDefaults instantiates a new PoxSigner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoxSignerWithDefaults() *PoxSigner {
	this := PoxSigner{}
	return &this
}

// GetSigningKey returns the SigningKey field value
func (o *PoxSigner) GetSigningKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SigningKey
}

// GetSigningKeyOk returns a tuple with the SigningKey field value
// and a boolean to check if the value has been set.
func (o *PoxSigner) GetSigningKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SigningKey, true
}

// SetSigningKey sets field value
func (o *PoxSigner) SetSigningKey(v string) {
	o.SigningKey = v
}

// GetSignerAddress returns the SignerAddress field value
func (o *PoxSigner) GetSignerAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SignerAddress
}

// GetSignerAddressOk returns a tuple with the SignerAddress field value
// and a boolean to check if the value has been set.
func (o *PoxSigner) GetSignerAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignerAddress, true
}

// SetSignerAddress sets field value
func (o *PoxSigner) SetSignerAddress(v string) {
	o.SignerAddress = v
}

// GetWeight returns the Weight field value
func (o *PoxSigner) GetWeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Weight
}

// GetWeightOk returns a tuple with the Weight field value
// and a boolean to check if the value has been set.
func (o *PoxSigner) GetWeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Weight, true
}

// SetWeight sets field value
func (o *PoxSigner) SetWeight(v int32) {
	o.Weight = v
}

// GetStackedAmount returns the StackedAmount field value
func (o *PoxSigner) GetStackedAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StackedAmount
}

// GetStackedAmountOk returns a tuple with the StackedAmount field value
// and a boolean to check if the value has been set.
func (o *PoxSigner) GetStackedAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StackedAmount, true
}

// SetStackedAmount sets field value
func (o *PoxSigner) SetStackedAmount(v string) {
	o.StackedAmount = v
}

// GetWeightPercent returns the WeightPercent field value
func (o *PoxSigner) GetWeightPercent() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.WeightPercent
}

// GetWeightPercentOk returns a tuple with the WeightPercent field value
// and a boolean to check if the value has been set.
func (o *PoxSigner) GetWeightPercentOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WeightPercent, true
}

// SetWeightPercent sets field value
func (o *PoxSigner) SetWeightPercent(v float32) {
	o.WeightPercent = v
}

// GetStackedAmountPercent returns the StackedAmountPercent field value
func (o *PoxSigner) GetStackedAmountPercent() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.StackedAmountPercent
}

// GetStackedAmountPercentOk returns a tuple with the StackedAmountPercent field value
// and a boolean to check if the value has been set.
func (o *PoxSigner) GetStackedAmountPercentOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StackedAmountPercent, true
}

// SetStackedAmountPercent sets field value
func (o *PoxSigner) SetStackedAmountPercent(v float32) {
	o.StackedAmountPercent = v
}

// GetSoloStackerCount returns the SoloStackerCount field value
func (o *PoxSigner) GetSoloStackerCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SoloStackerCount
}

// GetSoloStackerCountOk returns a tuple with the SoloStackerCount field value
// and a boolean to check if the value has been set.
func (o *PoxSigner) GetSoloStackerCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SoloStackerCount, true
}

// SetSoloStackerCount sets field value
func (o *PoxSigner) SetSoloStackerCount(v int32) {
	o.SoloStackerCount = v
}

// GetPooledStackerCount returns the PooledStackerCount field value
func (o *PoxSigner) GetPooledStackerCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PooledStackerCount
}

// GetPooledStackerCountOk returns a tuple with the PooledStackerCount field value
// and a boolean to check if the value has been set.
func (o *PoxSigner) GetPooledStackerCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PooledStackerCount, true
}

// SetPooledStackerCount sets field value
func (o *PoxSigner) SetPooledStackerCount(v int32) {
	o.PooledStackerCount = v
}

func (o PoxSigner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PoxSigner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["signing_key"] = o.SigningKey
	toSerialize["signer_address"] = o.SignerAddress
	toSerialize["weight"] = o.Weight
	toSerialize["stacked_amount"] = o.StackedAmount
	toSerialize["weight_percent"] = o.WeightPercent
	toSerialize["stacked_amount_percent"] = o.StackedAmountPercent
	toSerialize["solo_stacker_count"] = o.SoloStackerCount
	toSerialize["pooled_stacker_count"] = o.PooledStackerCount
	return toSerialize, nil
}

func (o *PoxSigner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"signing_key",
		"signer_address",
		"weight",
		"stacked_amount",
		"weight_percent",
		"stacked_amount_percent",
		"solo_stacker_count",
		"pooled_stacker_count",
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

	varPoxSigner := _PoxSigner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPoxSigner)

	if err != nil {
		return err
	}

	*o = PoxSigner(varPoxSigner)

	return err
}

type NullablePoxSigner struct {
	value *PoxSigner
	isSet bool
}

func (v NullablePoxSigner) Get() *PoxSigner {
	return v.value
}

func (v *NullablePoxSigner) Set(val *PoxSigner) {
	v.value = val
	v.isSet = true
}

func (v NullablePoxSigner) IsSet() bool {
	return v.isSet
}

func (v *NullablePoxSigner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoxSigner(val *PoxSigner) *NullablePoxSigner {
	return &NullablePoxSigner{value: val, isSet: true}
}

func (v NullablePoxSigner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoxSigner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
