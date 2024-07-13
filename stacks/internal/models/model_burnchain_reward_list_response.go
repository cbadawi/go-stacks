package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the BurnchainRewardListResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &BurnchainRewardListResponse{}

// BurnchainRewardListResponse GET request that returns blocks
type BurnchainRewardListResponse struct {
	// The number of burnchain rewards to return
	Limit int32 `json:"limit"`
	// The number to burnchain rewards to skip (starting at `0`)
	Offset  int32             `json:"offset"`
	Results []BurnchainReward `json:"results"`
}

type _BurnchainRewardListResponse BurnchainRewardListResponse

// NewBurnchainRewardListResponse instantiates a new BurnchainRewardListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBurnchainRewardListResponse(limit int32, offset int32, results []BurnchainReward) *BurnchainRewardListResponse {
	this := BurnchainRewardListResponse{}
	this.Limit = limit
	this.Offset = offset
	this.Results = results
	return &this
}

// NewBurnchainRewardListResponseWithDefaults instantiates a new BurnchainRewardListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBurnchainRewardListResponseWithDefaults() *BurnchainRewardListResponse {
	this := BurnchainRewardListResponse{}
	var offset int32 = 0
	this.Offset = offset
	return &this
}

// GetLimit returns the Limit field value
func (o *BurnchainRewardListResponse) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *BurnchainRewardListResponse) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *BurnchainRewardListResponse) SetLimit(v int32) {
	o.Limit = v
}

// GetOffset returns the Offset field value
func (o *BurnchainRewardListResponse) GetOffset() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *BurnchainRewardListResponse) GetOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value
func (o *BurnchainRewardListResponse) SetOffset(v int32) {
	o.Offset = v
}

// GetResults returns the Results field value
func (o *BurnchainRewardListResponse) GetResults() []BurnchainReward {
	if o == nil {
		var ret []BurnchainReward
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *BurnchainRewardListResponse) GetResultsOk() ([]BurnchainReward, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *BurnchainRewardListResponse) SetResults(v []BurnchainReward) {
	o.Results = v
}

func (o BurnchainRewardListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BurnchainRewardListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["limit"] = o.Limit
	toSerialize["offset"] = o.Offset
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

func (o *BurnchainRewardListResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"limit",
		"offset",
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

	varBurnchainRewardListResponse := _BurnchainRewardListResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBurnchainRewardListResponse)

	if err != nil {
		return err
	}

	*o = BurnchainRewardListResponse(varBurnchainRewardListResponse)

	return err
}

type NullableBurnchainRewardListResponse struct {
	value *BurnchainRewardListResponse
	isSet bool
}

func (v NullableBurnchainRewardListResponse) Get() *BurnchainRewardListResponse {
	return v.value
}

func (v *NullableBurnchainRewardListResponse) Set(val *BurnchainRewardListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBurnchainRewardListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBurnchainRewardListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBurnchainRewardListResponse(val *BurnchainRewardListResponse) *NullableBurnchainRewardListResponse {
	return &NullableBurnchainRewardListResponse{value: val, isSet: true}
}

func (v NullableBurnchainRewardListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBurnchainRewardListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
