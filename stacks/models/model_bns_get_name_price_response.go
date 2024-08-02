package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the BnsGetNamePriceResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &BnsGetNamePriceResponse{}

// BnsGetNamePriceResponse Fetch price for name.
type BnsGetNamePriceResponse struct {
	Units  string `json:"units"`
	Amount string `json:"amount" validate:"regexp=^[0-9]+$"`
}

type _BnsGetNamePriceResponse BnsGetNamePriceResponse

// NewBnsGetNamePriceResponse instantiates a new BnsGetNamePriceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBnsGetNamePriceResponse(units string, amount string) *BnsGetNamePriceResponse {
	this := BnsGetNamePriceResponse{}
	this.Units = units
	this.Amount = amount
	return &this
}

// NewBnsGetNamePriceResponseWithDefaults instantiates a new BnsGetNamePriceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBnsGetNamePriceResponseWithDefaults() *BnsGetNamePriceResponse {
	this := BnsGetNamePriceResponse{}
	return &this
}

// GetUnits returns the Units field value
func (o *BnsGetNamePriceResponse) GetUnits() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Units
}

// GetUnitsOk returns a tuple with the Units field value
// and a boolean to check if the value has been set.
func (o *BnsGetNamePriceResponse) GetUnitsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Units, true
}

// SetUnits sets field value
func (o *BnsGetNamePriceResponse) SetUnits(v string) {
	o.Units = v
}

// GetAmount returns the Amount field value
func (o *BnsGetNamePriceResponse) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *BnsGetNamePriceResponse) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *BnsGetNamePriceResponse) SetAmount(v string) {
	o.Amount = v
}

func (o BnsGetNamePriceResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BnsGetNamePriceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["units"] = o.Units
	toSerialize["amount"] = o.Amount
	return toSerialize, nil
}

func (o *BnsGetNamePriceResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"units",
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

	varBnsGetNamePriceResponse := _BnsGetNamePriceResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBnsGetNamePriceResponse)

	if err != nil {
		return err
	}

	*o = BnsGetNamePriceResponse(varBnsGetNamePriceResponse)

	return err
}

type NullableBnsGetNamePriceResponse struct {
	value *BnsGetNamePriceResponse
	isSet bool
}

func (v NullableBnsGetNamePriceResponse) Get() *BnsGetNamePriceResponse {
	return v.value
}

func (v *NullableBnsGetNamePriceResponse) Set(val *BnsGetNamePriceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBnsGetNamePriceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBnsGetNamePriceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBnsGetNamePriceResponse(val *BnsGetNamePriceResponse) *NullableBnsGetNamePriceResponse {
	return &NullableBnsGetNamePriceResponse{value: val, isSet: true}
}

func (v NullableBnsGetNamePriceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBnsGetNamePriceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
