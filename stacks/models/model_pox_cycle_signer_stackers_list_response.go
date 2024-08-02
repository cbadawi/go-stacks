package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the PoxCycleSignerStackersListResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PoxCycleSignerStackersListResponse{}

// PoxCycleSignerStackersListResponse GET request that returns stackers for a signer in a PoX cycle
type PoxCycleSignerStackersListResponse struct {
	// The number of stackers to return
	Limit int32 `json:"limit"`
	// The number to stackers to skip (starting at `0`)
	Offset int32 `json:"offset"`
	// The total number of stackers
	Total   int32        `json:"total"`
	Results []PoxStacker `json:"results"`
}

type _PoxCycleSignerStackersListResponse PoxCycleSignerStackersListResponse

// NewPoxCycleSignerStackersListResponse instantiates a new PoxCycleSignerStackersListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoxCycleSignerStackersListResponse(limit int32, offset int32, total int32, results []PoxStacker) *PoxCycleSignerStackersListResponse {
	this := PoxCycleSignerStackersListResponse{}
	this.Limit = limit
	this.Offset = offset
	this.Total = total
	this.Results = results
	return &this
}

// NewPoxCycleSignerStackersListResponseWithDefaults instantiates a new PoxCycleSignerStackersListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoxCycleSignerStackersListResponseWithDefaults() *PoxCycleSignerStackersListResponse {
	this := PoxCycleSignerStackersListResponse{}
	var offset int32 = 0
	this.Offset = offset
	return &this
}

// GetLimit returns the Limit field value
func (o *PoxCycleSignerStackersListResponse) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *PoxCycleSignerStackersListResponse) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *PoxCycleSignerStackersListResponse) SetLimit(v int32) {
	o.Limit = v
}

// GetOffset returns the Offset field value
func (o *PoxCycleSignerStackersListResponse) GetOffset() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *PoxCycleSignerStackersListResponse) GetOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value
func (o *PoxCycleSignerStackersListResponse) SetOffset(v int32) {
	o.Offset = v
}

// GetTotal returns the Total field value
func (o *PoxCycleSignerStackersListResponse) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *PoxCycleSignerStackersListResponse) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *PoxCycleSignerStackersListResponse) SetTotal(v int32) {
	o.Total = v
}

// GetResults returns the Results field value
func (o *PoxCycleSignerStackersListResponse) GetResults() []PoxStacker {
	if o == nil {
		var ret []PoxStacker
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PoxCycleSignerStackersListResponse) GetResultsOk() ([]PoxStacker, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PoxCycleSignerStackersListResponse) SetResults(v []PoxStacker) {
	o.Results = v
}

func (o PoxCycleSignerStackersListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PoxCycleSignerStackersListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["limit"] = o.Limit
	toSerialize["offset"] = o.Offset
	toSerialize["total"] = o.Total
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

func (o *PoxCycleSignerStackersListResponse) UnmarshalJSON(data []byte) (err error) {
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

	varPoxCycleSignerStackersListResponse := _PoxCycleSignerStackersListResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPoxCycleSignerStackersListResponse)

	if err != nil {
		return err
	}

	*o = PoxCycleSignerStackersListResponse(varPoxCycleSignerStackersListResponse)

	return err
}

type NullablePoxCycleSignerStackersListResponse struct {
	value *PoxCycleSignerStackersListResponse
	isSet bool
}

func (v NullablePoxCycleSignerStackersListResponse) Get() *PoxCycleSignerStackersListResponse {
	return v.value
}

func (v *NullablePoxCycleSignerStackersListResponse) Set(val *PoxCycleSignerStackersListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePoxCycleSignerStackersListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePoxCycleSignerStackersListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoxCycleSignerStackersListResponse(val *PoxCycleSignerStackersListResponse) *NullablePoxCycleSignerStackersListResponse {
	return &NullablePoxCycleSignerStackersListResponse{value: val, isSet: true}
}

func (v NullablePoxCycleSignerStackersListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoxCycleSignerStackersListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
