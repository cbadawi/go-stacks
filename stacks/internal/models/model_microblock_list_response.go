package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the MicroblockListResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &MicroblockListResponse{}

// MicroblockListResponse GET request that returns microblocks
type MicroblockListResponse struct {
	// The number of microblocks to return
	Limit int32 `json:"limit"`
	// The number to microblocks to skip (starting at `0`)
	Offset int32 `json:"offset"`
	// The number of microblocks available
	Total   int32        `json:"total"`
	Results []Microblock `json:"results"`
}

type _MicroblockListResponse MicroblockListResponse

// NewMicroblockListResponse instantiates a new MicroblockListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicroblockListResponse(limit int32, offset int32, total int32, results []Microblock) *MicroblockListResponse {
	this := MicroblockListResponse{}
	this.Limit = limit
	this.Offset = offset
	this.Total = total
	this.Results = results
	return &this
}

// NewMicroblockListResponseWithDefaults instantiates a new MicroblockListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicroblockListResponseWithDefaults() *MicroblockListResponse {
	this := MicroblockListResponse{}
	var offset int32 = 0
	this.Offset = offset
	return &this
}

// GetLimit returns the Limit field value
func (o *MicroblockListResponse) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *MicroblockListResponse) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *MicroblockListResponse) SetLimit(v int32) {
	o.Limit = v
}

// GetOffset returns the Offset field value
func (o *MicroblockListResponse) GetOffset() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *MicroblockListResponse) GetOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value
func (o *MicroblockListResponse) SetOffset(v int32) {
	o.Offset = v
}

// GetTotal returns the Total field value
func (o *MicroblockListResponse) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *MicroblockListResponse) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *MicroblockListResponse) SetTotal(v int32) {
	o.Total = v
}

// GetResults returns the Results field value
func (o *MicroblockListResponse) GetResults() []Microblock {
	if o == nil {
		var ret []Microblock
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *MicroblockListResponse) GetResultsOk() ([]Microblock, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *MicroblockListResponse) SetResults(v []Microblock) {
	o.Results = v
}

func (o MicroblockListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MicroblockListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["limit"] = o.Limit
	toSerialize["offset"] = o.Offset
	toSerialize["total"] = o.Total
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

func (o *MicroblockListResponse) UnmarshalJSON(data []byte) (err error) {
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

	varMicroblockListResponse := _MicroblockListResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMicroblockListResponse)

	if err != nil {
		return err
	}

	*o = MicroblockListResponse(varMicroblockListResponse)

	return err
}

type NullableMicroblockListResponse struct {
	value *MicroblockListResponse
	isSet bool
}

func (v NullableMicroblockListResponse) Get() *MicroblockListResponse {
	return v.value
}

func (v *NullableMicroblockListResponse) Set(val *MicroblockListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMicroblockListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMicroblockListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicroblockListResponse(val *MicroblockListResponse) *NullableMicroblockListResponse {
	return &NullableMicroblockListResponse{value: val, isSet: true}
}

func (v NullableMicroblockListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicroblockListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
