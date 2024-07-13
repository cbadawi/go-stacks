package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the NonFungibleTokenMintList type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &NonFungibleTokenMintList{}

// NonFungibleTokenMintList List of Non-Fungible Token mint events for an asset identifier
type NonFungibleTokenMintList struct {
	// The number of mint events to return
	Limit int32 `json:"limit"`
	// The number to mint events to skip (starting at `0`)
	Offset int32 `json:"offset"`
	// The number of mint events available
	Total   int32                  `json:"total"`
	Results []NonFungibleTokenMint `json:"results"`
}

type _NonFungibleTokenMintList NonFungibleTokenMintList

// NewNonFungibleTokenMintList instantiates a new NonFungibleTokenMintList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNonFungibleTokenMintList(limit int32, offset int32, total int32, results []NonFungibleTokenMint) *NonFungibleTokenMintList {
	this := NonFungibleTokenMintList{}
	this.Limit = limit
	this.Offset = offset
	this.Total = total
	this.Results = results
	return &this
}

// NewNonFungibleTokenMintListWithDefaults instantiates a new NonFungibleTokenMintList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNonFungibleTokenMintListWithDefaults() *NonFungibleTokenMintList {
	this := NonFungibleTokenMintList{}
	return &this
}

// GetLimit returns the Limit field value
func (o *NonFungibleTokenMintList) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *NonFungibleTokenMintList) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *NonFungibleTokenMintList) SetLimit(v int32) {
	o.Limit = v
}

// GetOffset returns the Offset field value
func (o *NonFungibleTokenMintList) GetOffset() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *NonFungibleTokenMintList) GetOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value
func (o *NonFungibleTokenMintList) SetOffset(v int32) {
	o.Offset = v
}

// GetTotal returns the Total field value
func (o *NonFungibleTokenMintList) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *NonFungibleTokenMintList) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *NonFungibleTokenMintList) SetTotal(v int32) {
	o.Total = v
}

// GetResults returns the Results field value
func (o *NonFungibleTokenMintList) GetResults() []NonFungibleTokenMint {
	if o == nil {
		var ret []NonFungibleTokenMint
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *NonFungibleTokenMintList) GetResultsOk() ([]NonFungibleTokenMint, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *NonFungibleTokenMintList) SetResults(v []NonFungibleTokenMint) {
	o.Results = v
}

func (o NonFungibleTokenMintList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NonFungibleTokenMintList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["limit"] = o.Limit
	toSerialize["offset"] = o.Offset
	toSerialize["total"] = o.Total
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

func (o *NonFungibleTokenMintList) UnmarshalJSON(data []byte) (err error) {
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

	varNonFungibleTokenMintList := _NonFungibleTokenMintList{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNonFungibleTokenMintList)

	if err != nil {
		return err
	}

	*o = NonFungibleTokenMintList(varNonFungibleTokenMintList)

	return err
}

type NullableNonFungibleTokenMintList struct {
	value *NonFungibleTokenMintList
	isSet bool
}

func (v NullableNonFungibleTokenMintList) Get() *NonFungibleTokenMintList {
	return v.value
}

func (v *NullableNonFungibleTokenMintList) Set(val *NonFungibleTokenMintList) {
	v.value = val
	v.isSet = true
}

func (v NullableNonFungibleTokenMintList) IsSet() bool {
	return v.isSet
}

func (v *NullableNonFungibleTokenMintList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNonFungibleTokenMintList(val *NonFungibleTokenMintList) *NullableNonFungibleTokenMintList {
	return &NullableNonFungibleTokenMintList{value: val, isSet: true}
}

func (v NullableNonFungibleTokenMintList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNonFungibleTokenMintList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
