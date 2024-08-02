package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the NonFungibleTokenHistoryEventList type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &NonFungibleTokenHistoryEventList{}

// NonFungibleTokenHistoryEventList List of Non-Fungible Token history events
type NonFungibleTokenHistoryEventList struct {
	// The number of events to return
	Limit int32 `json:"limit"`
	// The number to events to skip (starting at `0`)
	Offset int32 `json:"offset"`
	// The number of events available
	Total   int32                          `json:"total"`
	Results []NonFungibleTokenHistoryEvent `json:"results"`
}

type _NonFungibleTokenHistoryEventList NonFungibleTokenHistoryEventList

// NewNonFungibleTokenHistoryEventList instantiates a new NonFungibleTokenHistoryEventList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNonFungibleTokenHistoryEventList(limit int32, offset int32, total int32, results []NonFungibleTokenHistoryEvent) *NonFungibleTokenHistoryEventList {
	this := NonFungibleTokenHistoryEventList{}
	this.Limit = limit
	this.Offset = offset
	this.Total = total
	this.Results = results
	return &this
}

// NewNonFungibleTokenHistoryEventListWithDefaults instantiates a new NonFungibleTokenHistoryEventList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNonFungibleTokenHistoryEventListWithDefaults() *NonFungibleTokenHistoryEventList {
	this := NonFungibleTokenHistoryEventList{}
	return &this
}

// GetLimit returns the Limit field value
func (o *NonFungibleTokenHistoryEventList) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *NonFungibleTokenHistoryEventList) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *NonFungibleTokenHistoryEventList) SetLimit(v int32) {
	o.Limit = v
}

// GetOffset returns the Offset field value
func (o *NonFungibleTokenHistoryEventList) GetOffset() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *NonFungibleTokenHistoryEventList) GetOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value
func (o *NonFungibleTokenHistoryEventList) SetOffset(v int32) {
	o.Offset = v
}

// GetTotal returns the Total field value
func (o *NonFungibleTokenHistoryEventList) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *NonFungibleTokenHistoryEventList) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *NonFungibleTokenHistoryEventList) SetTotal(v int32) {
	o.Total = v
}

// GetResults returns the Results field value
func (o *NonFungibleTokenHistoryEventList) GetResults() []NonFungibleTokenHistoryEvent {
	if o == nil {
		var ret []NonFungibleTokenHistoryEvent
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *NonFungibleTokenHistoryEventList) GetResultsOk() ([]NonFungibleTokenHistoryEvent, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *NonFungibleTokenHistoryEventList) SetResults(v []NonFungibleTokenHistoryEvent) {
	o.Results = v
}

func (o NonFungibleTokenHistoryEventList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NonFungibleTokenHistoryEventList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["limit"] = o.Limit
	toSerialize["offset"] = o.Offset
	toSerialize["total"] = o.Total
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

func (o *NonFungibleTokenHistoryEventList) UnmarshalJSON(data []byte) (err error) {
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

	varNonFungibleTokenHistoryEventList := _NonFungibleTokenHistoryEventList{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNonFungibleTokenHistoryEventList)

	if err != nil {
		return err
	}

	*o = NonFungibleTokenHistoryEventList(varNonFungibleTokenHistoryEventList)

	return err
}

type NullableNonFungibleTokenHistoryEventList struct {
	value *NonFungibleTokenHistoryEventList
	isSet bool
}

func (v NullableNonFungibleTokenHistoryEventList) Get() *NonFungibleTokenHistoryEventList {
	return v.value
}

func (v *NullableNonFungibleTokenHistoryEventList) Set(val *NonFungibleTokenHistoryEventList) {
	v.value = val
	v.isSet = true
}

func (v NullableNonFungibleTokenHistoryEventList) IsSet() bool {
	return v.isSet
}

func (v *NullableNonFungibleTokenHistoryEventList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNonFungibleTokenHistoryEventList(val *NonFungibleTokenHistoryEventList) *NullableNonFungibleTokenHistoryEventList {
	return &NullableNonFungibleTokenHistoryEventList{value: val, isSet: true}
}

func (v NullableNonFungibleTokenHistoryEventList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNonFungibleTokenHistoryEventList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
