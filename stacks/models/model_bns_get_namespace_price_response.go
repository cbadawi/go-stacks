package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the BnsGetNamespacePriceResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &BnsGetNamespacePriceResponse{}

// BnsGetNamespacePriceResponse Fetch price for namespace.
type BnsGetNamespacePriceResponse struct {
	Units  string `json:"units"`
	Amount string `json:"amount" validate:"regexp=^[0-9]+$"`
}

type _BnsGetNamespacePriceResponse BnsGetNamespacePriceResponse

// NewBnsGetNamespacePriceResponse instantiates a new BnsGetNamespacePriceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBnsGetNamespacePriceResponse(units string, amount string) *BnsGetNamespacePriceResponse {
	this := BnsGetNamespacePriceResponse{}
	this.Units = units
	this.Amount = amount
	return &this
}

// NewBnsGetNamespacePriceResponseWithDefaults instantiates a new BnsGetNamespacePriceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBnsGetNamespacePriceResponseWithDefaults() *BnsGetNamespacePriceResponse {
	this := BnsGetNamespacePriceResponse{}
	return &this
}

// GetUnits returns the Units field value
func (o *BnsGetNamespacePriceResponse) GetUnits() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Units
}

// GetUnitsOk returns a tuple with the Units field value
// and a boolean to check if the value has been set.
func (o *BnsGetNamespacePriceResponse) GetUnitsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Units, true
}

// SetUnits sets field value
func (o *BnsGetNamespacePriceResponse) SetUnits(v string) {
	o.Units = v
}

// GetAmount returns the Amount field value
func (o *BnsGetNamespacePriceResponse) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *BnsGetNamespacePriceResponse) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *BnsGetNamespacePriceResponse) SetAmount(v string) {
	o.Amount = v
}

func (o BnsGetNamespacePriceResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BnsGetNamespacePriceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["units"] = o.Units
	toSerialize["amount"] = o.Amount
	return toSerialize, nil
}

func (o *BnsGetNamespacePriceResponse) UnmarshalJSON(data []byte) (err error) {
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

	varBnsGetNamespacePriceResponse := _BnsGetNamespacePriceResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBnsGetNamespacePriceResponse)

	if err != nil {
		return err
	}

	*o = BnsGetNamespacePriceResponse(varBnsGetNamespacePriceResponse)

	return err
}

type NullableBnsGetNamespacePriceResponse struct {
	value *BnsGetNamespacePriceResponse
	isSet bool
}

func (v NullableBnsGetNamespacePriceResponse) Get() *BnsGetNamespacePriceResponse {
	return v.value
}

func (v *NullableBnsGetNamespacePriceResponse) Set(val *BnsGetNamespacePriceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBnsGetNamespacePriceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBnsGetNamespacePriceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBnsGetNamespacePriceResponse(val *BnsGetNamespacePriceResponse) *NullableBnsGetNamespacePriceResponse {
	return &NullableBnsGetNamespacePriceResponse{value: val, isSet: true}
}

func (v NullableBnsGetNamespacePriceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBnsGetNamespacePriceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
