package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the RosettaMaxFeeAmount type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RosettaMaxFeeAmount{}

// RosettaMaxFeeAmount Amount is some Value of a Currency. It is considered invalid to specify a Value without a Currency.
type RosettaMaxFeeAmount struct {
	// Value of the transaction in atomic units represented as an arbitrary-sized signed integer. For example, 1 BTC would be represented by a value of 100000000.
	Value    string          `json:"value"`
	Currency RosettaCurrency `json:"currency"`
	//
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

type _RosettaMaxFeeAmount RosettaMaxFeeAmount

// NewRosettaMaxFeeAmount instantiates a new RosettaMaxFeeAmount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRosettaMaxFeeAmount(value string, currency RosettaCurrency) *RosettaMaxFeeAmount {
	this := RosettaMaxFeeAmount{}
	this.Value = value
	this.Currency = currency
	return &this
}

// NewRosettaMaxFeeAmountWithDefaults instantiates a new RosettaMaxFeeAmount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRosettaMaxFeeAmountWithDefaults() *RosettaMaxFeeAmount {
	this := RosettaMaxFeeAmount{}
	return &this
}

// GetValue returns the Value field value
func (o *RosettaMaxFeeAmount) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *RosettaMaxFeeAmount) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *RosettaMaxFeeAmount) SetValue(v string) {
	o.Value = v
}

// GetCurrency returns the Currency field value
func (o *RosettaMaxFeeAmount) GetCurrency() RosettaCurrency {
	if o == nil {
		var ret RosettaCurrency
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *RosettaMaxFeeAmount) GetCurrencyOk() (*RosettaCurrency, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *RosettaMaxFeeAmount) SetCurrency(v RosettaCurrency) {
	o.Currency = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *RosettaMaxFeeAmount) GetMetadata() map[string]interface{} {
	if o == nil || utils.IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaMaxFeeAmount) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || utils.IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *RosettaMaxFeeAmount) HasMetadata() bool {
	if o != nil && !utils.IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *RosettaMaxFeeAmount) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o RosettaMaxFeeAmount) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RosettaMaxFeeAmount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["value"] = o.Value
	toSerialize["currency"] = o.Currency
	if !utils.IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

func (o *RosettaMaxFeeAmount) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"value",
		"currency",
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

	varRosettaMaxFeeAmount := _RosettaMaxFeeAmount{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRosettaMaxFeeAmount)

	if err != nil {
		return err
	}

	*o = RosettaMaxFeeAmount(varRosettaMaxFeeAmount)

	return err
}

type NullableRosettaMaxFeeAmount struct {
	value *RosettaMaxFeeAmount
	isSet bool
}

func (v NullableRosettaMaxFeeAmount) Get() *RosettaMaxFeeAmount {
	return v.value
}

func (v *NullableRosettaMaxFeeAmount) Set(val *RosettaMaxFeeAmount) {
	v.value = val
	v.isSet = true
}

func (v NullableRosettaMaxFeeAmount) IsSet() bool {
	return v.isSet
}

func (v *NullableRosettaMaxFeeAmount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRosettaMaxFeeAmount(val *RosettaMaxFeeAmount) *NullableRosettaMaxFeeAmount {
	return &NullableRosettaMaxFeeAmount{value: val, isSet: true}
}

func (v NullableRosettaMaxFeeAmount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRosettaMaxFeeAmount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
