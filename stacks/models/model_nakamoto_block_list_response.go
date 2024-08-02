package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the NakamotoBlockListResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &NakamotoBlockListResponse{}

// NakamotoBlockListResponse GET request that returns blocks
type NakamotoBlockListResponse struct {
	// The number of blocks to return
	Limit int32 `json:"limit"`
	// The number to blocks to skip (starting at `0`)
	Offset int32 `json:"offset"`
	// The number of blocks available
	Total   int32           `json:"total"`
	Results []NakamotoBlock `json:"results"`
}

type _NakamotoBlockListResponse NakamotoBlockListResponse

// NewNakamotoBlockListResponse instantiates a new NakamotoBlockListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNakamotoBlockListResponse(limit int32, offset int32, total int32, results []NakamotoBlock) *NakamotoBlockListResponse {
	this := NakamotoBlockListResponse{}
	this.Limit = limit
	this.Offset = offset
	this.Total = total
	this.Results = results
	return &this
}

// NewNakamotoBlockListResponseWithDefaults instantiates a new NakamotoBlockListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNakamotoBlockListResponseWithDefaults() *NakamotoBlockListResponse {
	this := NakamotoBlockListResponse{}
	var offset int32 = 0
	this.Offset = offset
	return &this
}

// GetLimit returns the Limit field value
func (o *NakamotoBlockListResponse) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *NakamotoBlockListResponse) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *NakamotoBlockListResponse) SetLimit(v int32) {
	o.Limit = v
}

// GetOffset returns the Offset field value
func (o *NakamotoBlockListResponse) GetOffset() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *NakamotoBlockListResponse) GetOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value
func (o *NakamotoBlockListResponse) SetOffset(v int32) {
	o.Offset = v
}

// GetTotal returns the Total field value
func (o *NakamotoBlockListResponse) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *NakamotoBlockListResponse) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *NakamotoBlockListResponse) SetTotal(v int32) {
	o.Total = v
}

// GetResults returns the Results field value
func (o *NakamotoBlockListResponse) GetResults() []NakamotoBlock {
	if o == nil {
		var ret []NakamotoBlock
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *NakamotoBlockListResponse) GetResultsOk() ([]NakamotoBlock, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *NakamotoBlockListResponse) SetResults(v []NakamotoBlock) {
	o.Results = v
}

func (o NakamotoBlockListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NakamotoBlockListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["limit"] = o.Limit
	toSerialize["offset"] = o.Offset
	toSerialize["total"] = o.Total
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

func (o *NakamotoBlockListResponse) UnmarshalJSON(data []byte) (err error) {
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

	varNakamotoBlockListResponse := _NakamotoBlockListResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNakamotoBlockListResponse)

	if err != nil {
		return err
	}

	*o = NakamotoBlockListResponse(varNakamotoBlockListResponse)

	return err
}

type NullableNakamotoBlockListResponse struct {
	value *NakamotoBlockListResponse
	isSet bool
}

func (v NullableNakamotoBlockListResponse) Get() *NakamotoBlockListResponse {
	return v.value
}

func (v *NullableNakamotoBlockListResponse) Set(val *NakamotoBlockListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNakamotoBlockListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNakamotoBlockListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNakamotoBlockListResponse(val *NakamotoBlockListResponse) *NullableNakamotoBlockListResponse {
	return &NullableNakamotoBlockListResponse{value: val, isSet: true}
}

func (v NullableNakamotoBlockListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNakamotoBlockListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
