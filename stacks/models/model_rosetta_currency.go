package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the RosettaCurrency type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RosettaCurrency{}

// RosettaCurrency Currency is composed of a canonical Symbol and Decimals. This Decimals value is used to convert an Amount.Value from atomic units (Satoshis) to standard units (Bitcoins).
type RosettaCurrency struct {
	// Canonical symbol associated with a currency.
	Symbol string `json:"symbol"`
	// Number of decimal places in the standard unit representation of the amount. For example, BTC has 8 decimals. Note that it is not possible to represent the value of some currency in atomic units that is not base 10.
	Decimals int32 `json:"decimals"`
	// Any additional information related to the currency itself. For example, it would be useful to populate this object with the contract address of an ERC-20 token.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

type _RosettaCurrency RosettaCurrency

// NewRosettaCurrency instantiates a new RosettaCurrency object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRosettaCurrency(symbol string, decimals int32) *RosettaCurrency {
	this := RosettaCurrency{}
	this.Symbol = symbol
	this.Decimals = decimals
	return &this
}

// NewRosettaCurrencyWithDefaults instantiates a new RosettaCurrency object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRosettaCurrencyWithDefaults() *RosettaCurrency {
	this := RosettaCurrency{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *RosettaCurrency) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *RosettaCurrency) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *RosettaCurrency) SetSymbol(v string) {
	o.Symbol = v
}

// GetDecimals returns the Decimals field value
func (o *RosettaCurrency) GetDecimals() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Decimals
}

// GetDecimalsOk returns a tuple with the Decimals field value
// and a boolean to check if the value has been set.
func (o *RosettaCurrency) GetDecimalsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Decimals, true
}

// SetDecimals sets field value
func (o *RosettaCurrency) SetDecimals(v int32) {
	o.Decimals = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *RosettaCurrency) GetMetadata() map[string]interface{} {
	if o == nil || utils.IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaCurrency) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || utils.IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *RosettaCurrency) HasMetadata() bool {
	if o != nil && !utils.IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *RosettaCurrency) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o RosettaCurrency) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RosettaCurrency) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["symbol"] = o.Symbol
	toSerialize["decimals"] = o.Decimals
	if !utils.IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

func (o *RosettaCurrency) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"symbol",
		"decimals",
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

	varRosettaCurrency := _RosettaCurrency{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRosettaCurrency)

	if err != nil {
		return err
	}

	*o = RosettaCurrency(varRosettaCurrency)

	return err
}

type NullableRosettaCurrency struct {
	value *RosettaCurrency
	isSet bool
}

func (v NullableRosettaCurrency) Get() *RosettaCurrency {
	return v.value
}

func (v *NullableRosettaCurrency) Set(val *RosettaCurrency) {
	v.value = val
	v.isSet = true
}

func (v NullableRosettaCurrency) IsSet() bool {
	return v.isSet
}

func (v *NullableRosettaCurrency) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRosettaCurrency(val *RosettaCurrency) *NullableRosettaCurrency {
	return &NullableRosettaCurrency{value: val, isSet: true}
}

func (v NullableRosettaCurrency) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRosettaCurrency) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
