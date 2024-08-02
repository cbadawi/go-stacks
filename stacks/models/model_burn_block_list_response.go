package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the BurnBlockListResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &BurnBlockListResponse{}

// BurnBlockListResponse GET request that returns burn blocks
type BurnBlockListResponse struct {
	// The number of burn blocks to return
	Limit int32 `json:"limit"`
	// The number to burn blocks to skip (starting at `0`)
	Offset int32 `json:"offset"`
	// The number of burn blocks available (regardless of filter parameters)
	Total   int32       `json:"total"`
	Results []BurnBlock `json:"results"`
}

type _BurnBlockListResponse BurnBlockListResponse

// NewBurnBlockListResponse instantiates a new BurnBlockListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBurnBlockListResponse(limit int32, offset int32, total int32, results []BurnBlock) *BurnBlockListResponse {
	this := BurnBlockListResponse{}
	this.Limit = limit
	this.Offset = offset
	this.Total = total
	this.Results = results
	return &this
}

// NewBurnBlockListResponseWithDefaults instantiates a new BurnBlockListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBurnBlockListResponseWithDefaults() *BurnBlockListResponse {
	this := BurnBlockListResponse{}
	var offset int32 = 0
	this.Offset = offset
	return &this
}

// GetLimit returns the Limit field value
func (o *BurnBlockListResponse) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *BurnBlockListResponse) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *BurnBlockListResponse) SetLimit(v int32) {
	o.Limit = v
}

// GetOffset returns the Offset field value
func (o *BurnBlockListResponse) GetOffset() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *BurnBlockListResponse) GetOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value
func (o *BurnBlockListResponse) SetOffset(v int32) {
	o.Offset = v
}

// GetTotal returns the Total field value
func (o *BurnBlockListResponse) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *BurnBlockListResponse) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *BurnBlockListResponse) SetTotal(v int32) {
	o.Total = v
}

// GetResults returns the Results field value
func (o *BurnBlockListResponse) GetResults() []BurnBlock {
	if o == nil {
		var ret []BurnBlock
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *BurnBlockListResponse) GetResultsOk() ([]BurnBlock, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *BurnBlockListResponse) SetResults(v []BurnBlock) {
	o.Results = v
}

func (o BurnBlockListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BurnBlockListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["limit"] = o.Limit
	toSerialize["offset"] = o.Offset
	toSerialize["total"] = o.Total
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

func (o *BurnBlockListResponse) UnmarshalJSON(data []byte) (err error) {
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

	varBurnBlockListResponse := _BurnBlockListResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBurnBlockListResponse)

	if err != nil {
		return err
	}

	*o = BurnBlockListResponse(varBurnBlockListResponse)

	return err
}

type NullableBurnBlockListResponse struct {
	value *BurnBlockListResponse
	isSet bool
}

func (v NullableBurnBlockListResponse) Get() *BurnBlockListResponse {
	return v.value
}

func (v *NullableBurnBlockListResponse) Set(val *BurnBlockListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBurnBlockListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBurnBlockListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBurnBlockListResponse(val *BurnBlockListResponse) *NullableBurnBlockListResponse {
	return &NullableBurnBlockListResponse{value: val, isSet: true}
}

func (v NullableBurnBlockListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBurnBlockListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
