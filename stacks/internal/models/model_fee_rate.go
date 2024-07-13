package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the FeeRate type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &FeeRate{}

// FeeRate Get fee rate information.
type FeeRate struct {
	FeeRate int32 `json:"fee_rate"`
}

type _FeeRate FeeRate

// NewFeeRate instantiates a new FeeRate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeeRate(feeRate int32) *FeeRate {
	this := FeeRate{}
	this.FeeRate = feeRate
	return &this
}

// NewFeeRateWithDefaults instantiates a new FeeRate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeeRateWithDefaults() *FeeRate {
	this := FeeRate{}
	return &this
}

// GetFeeRate returns the FeeRate field value
func (o *FeeRate) GetFeeRate() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FeeRate
}

// GetFeeRateOk returns a tuple with the FeeRate field value
// and a boolean to check if the value has been set.
func (o *FeeRate) GetFeeRateOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeeRate, true
}

// SetFeeRate sets field value
func (o *FeeRate) SetFeeRate(v int32) {
	o.FeeRate = v
}

func (o FeeRate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeeRate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fee_rate"] = o.FeeRate
	return toSerialize, nil
}

func (o *FeeRate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fee_rate",
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

	varFeeRate := _FeeRate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFeeRate)

	if err != nil {
		return err
	}

	*o = FeeRate(varFeeRate)

	return err
}

type NullableFeeRate struct {
	value *FeeRate
	isSet bool
}

func (v NullableFeeRate) Get() *FeeRate {
	return v.value
}

func (v *NullableFeeRate) Set(val *FeeRate) {
	v.value = val
	v.isSet = true
}

func (v NullableFeeRate) IsSet() bool {
	return v.isSet
}

func (v *NullableFeeRate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeeRate(val *FeeRate) *NullableFeeRate {
	return &NullableFeeRate{value: val, isSet: true}
}

func (v NullableFeeRate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeeRate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
