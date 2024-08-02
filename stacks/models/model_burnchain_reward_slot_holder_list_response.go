package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the BurnchainRewardSlotHolderListResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &BurnchainRewardSlotHolderListResponse{}

// BurnchainRewardSlotHolderListResponse GET request that returns reward slot holders
type BurnchainRewardSlotHolderListResponse struct {
	// The number of items to return
	Limit int32 `json:"limit"`
	// The number of items to skip (starting at `0`)
	Offset int32 `json:"offset"`
	// Total number of available items
	Total   int32                       `json:"total"`
	Results []BurnchainRewardSlotHolder `json:"results"`
}

type _BurnchainRewardSlotHolderListResponse BurnchainRewardSlotHolderListResponse

// NewBurnchainRewardSlotHolderListResponse instantiates a new BurnchainRewardSlotHolderListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBurnchainRewardSlotHolderListResponse(limit int32, offset int32, total int32, results []BurnchainRewardSlotHolder) *BurnchainRewardSlotHolderListResponse {
	this := BurnchainRewardSlotHolderListResponse{}
	this.Limit = limit
	this.Offset = offset
	this.Total = total
	this.Results = results
	return &this
}

// NewBurnchainRewardSlotHolderListResponseWithDefaults instantiates a new BurnchainRewardSlotHolderListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBurnchainRewardSlotHolderListResponseWithDefaults() *BurnchainRewardSlotHolderListResponse {
	this := BurnchainRewardSlotHolderListResponse{}
	var offset int32 = 0
	this.Offset = offset
	return &this
}

// GetLimit returns the Limit field value
func (o *BurnchainRewardSlotHolderListResponse) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *BurnchainRewardSlotHolderListResponse) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *BurnchainRewardSlotHolderListResponse) SetLimit(v int32) {
	o.Limit = v
}

// GetOffset returns the Offset field value
func (o *BurnchainRewardSlotHolderListResponse) GetOffset() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *BurnchainRewardSlotHolderListResponse) GetOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value
func (o *BurnchainRewardSlotHolderListResponse) SetOffset(v int32) {
	o.Offset = v
}

// GetTotal returns the Total field value
func (o *BurnchainRewardSlotHolderListResponse) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *BurnchainRewardSlotHolderListResponse) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *BurnchainRewardSlotHolderListResponse) SetTotal(v int32) {
	o.Total = v
}

// GetResults returns the Results field value
func (o *BurnchainRewardSlotHolderListResponse) GetResults() []BurnchainRewardSlotHolder {
	if o == nil {
		var ret []BurnchainRewardSlotHolder
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *BurnchainRewardSlotHolderListResponse) GetResultsOk() ([]BurnchainRewardSlotHolder, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *BurnchainRewardSlotHolderListResponse) SetResults(v []BurnchainRewardSlotHolder) {
	o.Results = v
}

func (o BurnchainRewardSlotHolderListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BurnchainRewardSlotHolderListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["limit"] = o.Limit
	toSerialize["offset"] = o.Offset
	toSerialize["total"] = o.Total
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

func (o *BurnchainRewardSlotHolderListResponse) UnmarshalJSON(data []byte) (err error) {
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

	varBurnchainRewardSlotHolderListResponse := _BurnchainRewardSlotHolderListResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBurnchainRewardSlotHolderListResponse)

	if err != nil {
		return err
	}

	*o = BurnchainRewardSlotHolderListResponse(varBurnchainRewardSlotHolderListResponse)

	return err
}

type NullableBurnchainRewardSlotHolderListResponse struct {
	value *BurnchainRewardSlotHolderListResponse
	isSet bool
}

func (v NullableBurnchainRewardSlotHolderListResponse) Get() *BurnchainRewardSlotHolderListResponse {
	return v.value
}

func (v *NullableBurnchainRewardSlotHolderListResponse) Set(val *BurnchainRewardSlotHolderListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBurnchainRewardSlotHolderListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBurnchainRewardSlotHolderListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBurnchainRewardSlotHolderListResponse(val *BurnchainRewardSlotHolderListResponse) *NullableBurnchainRewardSlotHolderListResponse {
	return &NullableBurnchainRewardSlotHolderListResponse{value: val, isSet: true}
}

func (v NullableBurnchainRewardSlotHolderListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBurnchainRewardSlotHolderListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
