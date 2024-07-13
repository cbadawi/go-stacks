package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the PoolDelegationsResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PoolDelegationsResponse{}

// PoolDelegationsResponse GET request that returns stacking pool member details for a given pool (delegator) principal
type PoolDelegationsResponse struct {
	// The number of Stackers to return
	Limit int32 `json:"limit"`
	// The number to Stackers to skip (starting at `0`)
	Offset int32 `json:"offset"`
	// The total number of Stackers
	Total   int32            `json:"total"`
	Results []PoolDelegation `json:"results"`
}

type _PoolDelegationsResponse PoolDelegationsResponse

// NewPoolDelegationsResponse instantiates a new PoolDelegationsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoolDelegationsResponse(limit int32, offset int32, total int32, results []PoolDelegation) *PoolDelegationsResponse {
	this := PoolDelegationsResponse{}
	this.Limit = limit
	this.Offset = offset
	this.Total = total
	this.Results = results
	return &this
}

// NewPoolDelegationsResponseWithDefaults instantiates a new PoolDelegationsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoolDelegationsResponseWithDefaults() *PoolDelegationsResponse {
	this := PoolDelegationsResponse{}
	var offset int32 = 0
	this.Offset = offset
	return &this
}

// GetLimit returns the Limit field value
func (o *PoolDelegationsResponse) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *PoolDelegationsResponse) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *PoolDelegationsResponse) SetLimit(v int32) {
	o.Limit = v
}

// GetOffset returns the Offset field value
func (o *PoolDelegationsResponse) GetOffset() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *PoolDelegationsResponse) GetOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value
func (o *PoolDelegationsResponse) SetOffset(v int32) {
	o.Offset = v
}

// GetTotal returns the Total field value
func (o *PoolDelegationsResponse) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *PoolDelegationsResponse) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *PoolDelegationsResponse) SetTotal(v int32) {
	o.Total = v
}

// GetResults returns the Results field value
func (o *PoolDelegationsResponse) GetResults() []PoolDelegation {
	if o == nil {
		var ret []PoolDelegation
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PoolDelegationsResponse) GetResultsOk() ([]PoolDelegation, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PoolDelegationsResponse) SetResults(v []PoolDelegation) {
	o.Results = v
}

func (o PoolDelegationsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PoolDelegationsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["limit"] = o.Limit
	toSerialize["offset"] = o.Offset
	toSerialize["total"] = o.Total
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

func (o *PoolDelegationsResponse) UnmarshalJSON(data []byte) (err error) {
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

	varPoolDelegationsResponse := _PoolDelegationsResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPoolDelegationsResponse)

	if err != nil {
		return err
	}

	*o = PoolDelegationsResponse(varPoolDelegationsResponse)

	return err
}

type NullablePoolDelegationsResponse struct {
	value *PoolDelegationsResponse
	isSet bool
}

func (v NullablePoolDelegationsResponse) Get() *PoolDelegationsResponse {
	return v.value
}

func (v *NullablePoolDelegationsResponse) Set(val *PoolDelegationsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePoolDelegationsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePoolDelegationsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoolDelegationsResponse(val *PoolDelegationsResponse) *NullablePoolDelegationsResponse {
	return &NullablePoolDelegationsResponse{value: val, isSet: true}
}

func (v NullablePoolDelegationsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoolDelegationsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
