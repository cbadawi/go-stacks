package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the NonFungibleTokenHoldingsList type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &NonFungibleTokenHoldingsList{}

// NonFungibleTokenHoldingsList List of Non-Fungible Token holdings
type NonFungibleTokenHoldingsList struct {
	// The number of Non-Fungible Token holdings to return
	Limit int32 `json:"limit"`
	// The number to Non-Fungible Token holdings to skip (starting at `0`)
	Offset int32 `json:"offset"`
	// The number of Non-Fungible Token holdings available
	Total   int32                     `json:"total"`
	Results []NonFungibleTokenHolding `json:"results"`
}

type _NonFungibleTokenHoldingsList NonFungibleTokenHoldingsList

// NewNonFungibleTokenHoldingsList instantiates a new NonFungibleTokenHoldingsList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNonFungibleTokenHoldingsList(limit int32, offset int32, total int32, results []NonFungibleTokenHolding) *NonFungibleTokenHoldingsList {
	this := NonFungibleTokenHoldingsList{}
	this.Limit = limit
	this.Offset = offset
	this.Total = total
	this.Results = results
	return &this
}

// NewNonFungibleTokenHoldingsListWithDefaults instantiates a new NonFungibleTokenHoldingsList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNonFungibleTokenHoldingsListWithDefaults() *NonFungibleTokenHoldingsList {
	this := NonFungibleTokenHoldingsList{}
	return &this
}

// GetLimit returns the Limit field value
func (o *NonFungibleTokenHoldingsList) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *NonFungibleTokenHoldingsList) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *NonFungibleTokenHoldingsList) SetLimit(v int32) {
	o.Limit = v
}

// GetOffset returns the Offset field value
func (o *NonFungibleTokenHoldingsList) GetOffset() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *NonFungibleTokenHoldingsList) GetOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value
func (o *NonFungibleTokenHoldingsList) SetOffset(v int32) {
	o.Offset = v
}

// GetTotal returns the Total field value
func (o *NonFungibleTokenHoldingsList) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *NonFungibleTokenHoldingsList) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *NonFungibleTokenHoldingsList) SetTotal(v int32) {
	o.Total = v
}

// GetResults returns the Results field value
func (o *NonFungibleTokenHoldingsList) GetResults() []NonFungibleTokenHolding {
	if o == nil {
		var ret []NonFungibleTokenHolding
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *NonFungibleTokenHoldingsList) GetResultsOk() ([]NonFungibleTokenHolding, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *NonFungibleTokenHoldingsList) SetResults(v []NonFungibleTokenHolding) {
	o.Results = v
}

func (o NonFungibleTokenHoldingsList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NonFungibleTokenHoldingsList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["limit"] = o.Limit
	toSerialize["offset"] = o.Offset
	toSerialize["total"] = o.Total
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

func (o *NonFungibleTokenHoldingsList) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"limit",
		"offset",
		"total",
		"results",
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

	varNonFungibleTokenHoldingsList := _NonFungibleTokenHoldingsList{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNonFungibleTokenHoldingsList)

	if err != nil {
		return err
	}

	*o = NonFungibleTokenHoldingsList(varNonFungibleTokenHoldingsList)

	return err
}

type NullableNonFungibleTokenHoldingsList struct {
	value *NonFungibleTokenHoldingsList
	isSet bool
}

func (v NullableNonFungibleTokenHoldingsList) Get() *NonFungibleTokenHoldingsList {
	return v.value
}

func (v *NullableNonFungibleTokenHoldingsList) Set(val *NonFungibleTokenHoldingsList) {
	v.value = val
	v.isSet = true
}

func (v NullableNonFungibleTokenHoldingsList) IsSet() bool {
	return v.isSet
}

func (v *NullableNonFungibleTokenHoldingsList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNonFungibleTokenHoldingsList(val *NonFungibleTokenHoldingsList) *NullableNonFungibleTokenHoldingsList {
	return &NullableNonFungibleTokenHoldingsList{value: val, isSet: true}
}

func (v NullableNonFungibleTokenHoldingsList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNonFungibleTokenHoldingsList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
