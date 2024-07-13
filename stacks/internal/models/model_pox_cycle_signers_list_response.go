package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the PoxCycleSignersListResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PoxCycleSignersListResponse{}

// PoxCycleSignersListResponse GET request that returns signers for a PoX cycle
type PoxCycleSignersListResponse struct {
	// The number of signers to return
	Limit int32 `json:"limit"`
	// The number to signers to skip (starting at `0`)
	Offset int32 `json:"offset"`
	// The total number of signers
	Total   int32       `json:"total"`
	Results []PoxSigner `json:"results"`
}

type _PoxCycleSignersListResponse PoxCycleSignersListResponse

// NewPoxCycleSignersListResponse instantiates a new PoxCycleSignersListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoxCycleSignersListResponse(limit int32, offset int32, total int32, results []PoxSigner) *PoxCycleSignersListResponse {
	this := PoxCycleSignersListResponse{}
	this.Limit = limit
	this.Offset = offset
	this.Total = total
	this.Results = results
	return &this
}

// NewPoxCycleSignersListResponseWithDefaults instantiates a new PoxCycleSignersListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoxCycleSignersListResponseWithDefaults() *PoxCycleSignersListResponse {
	this := PoxCycleSignersListResponse{}
	var offset int32 = 0
	this.Offset = offset
	return &this
}

// GetLimit returns the Limit field value
func (o *PoxCycleSignersListResponse) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *PoxCycleSignersListResponse) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *PoxCycleSignersListResponse) SetLimit(v int32) {
	o.Limit = v
}

// GetOffset returns the Offset field value
func (o *PoxCycleSignersListResponse) GetOffset() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *PoxCycleSignersListResponse) GetOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value
func (o *PoxCycleSignersListResponse) SetOffset(v int32) {
	o.Offset = v
}

// GetTotal returns the Total field value
func (o *PoxCycleSignersListResponse) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *PoxCycleSignersListResponse) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *PoxCycleSignersListResponse) SetTotal(v int32) {
	o.Total = v
}

// GetResults returns the Results field value
func (o *PoxCycleSignersListResponse) GetResults() []PoxSigner {
	if o == nil {
		var ret []PoxSigner
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PoxCycleSignersListResponse) GetResultsOk() ([]PoxSigner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PoxCycleSignersListResponse) SetResults(v []PoxSigner) {
	o.Results = v
}

func (o PoxCycleSignersListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PoxCycleSignersListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["limit"] = o.Limit
	toSerialize["offset"] = o.Offset
	toSerialize["total"] = o.Total
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

func (o *PoxCycleSignersListResponse) UnmarshalJSON(data []byte) (err error) {
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

	varPoxCycleSignersListResponse := _PoxCycleSignersListResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPoxCycleSignersListResponse)

	if err != nil {
		return err
	}

	*o = PoxCycleSignersListResponse(varPoxCycleSignersListResponse)

	return err
}

type NullablePoxCycleSignersListResponse struct {
	value *PoxCycleSignersListResponse
	isSet bool
}

func (v NullablePoxCycleSignersListResponse) Get() *PoxCycleSignersListResponse {
	return v.value
}

func (v *NullablePoxCycleSignersListResponse) Set(val *PoxCycleSignersListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePoxCycleSignersListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePoxCycleSignersListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoxCycleSignersListResponse(val *PoxCycleSignersListResponse) *NullablePoxCycleSignersListResponse {
	return &NullablePoxCycleSignersListResponse{value: val, isSet: true}
}

func (v NullablePoxCycleSignersListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoxCycleSignersListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
