package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the AddressStxInboundListResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AddressStxInboundListResponse{}

// AddressStxInboundListResponse GET request that returns a list of inbound STX transfers with a memo
type AddressStxInboundListResponse struct {
	Limit   int32                `json:"limit"`
	Offset  int32                `json:"offset"`
	Total   int32                `json:"total"`
	Results []InboundStxTransfer `json:"results"`
}

type _AddressStxInboundListResponse AddressStxInboundListResponse

// NewAddressStxInboundListResponse instantiates a new AddressStxInboundListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressStxInboundListResponse(limit int32, offset int32, total int32, results []InboundStxTransfer) *AddressStxInboundListResponse {
	this := AddressStxInboundListResponse{}
	this.Limit = limit
	this.Offset = offset
	this.Total = total
	this.Results = results
	return &this
}

// NewAddressStxInboundListResponseWithDefaults instantiates a new AddressStxInboundListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressStxInboundListResponseWithDefaults() *AddressStxInboundListResponse {
	this := AddressStxInboundListResponse{}
	return &this
}

// GetLimit returns the Limit field value
func (o *AddressStxInboundListResponse) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *AddressStxInboundListResponse) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *AddressStxInboundListResponse) SetLimit(v int32) {
	o.Limit = v
}

// GetOffset returns the Offset field value
func (o *AddressStxInboundListResponse) GetOffset() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *AddressStxInboundListResponse) GetOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value
func (o *AddressStxInboundListResponse) SetOffset(v int32) {
	o.Offset = v
}

// GetTotal returns the Total field value
func (o *AddressStxInboundListResponse) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *AddressStxInboundListResponse) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *AddressStxInboundListResponse) SetTotal(v int32) {
	o.Total = v
}

// GetResults returns the Results field value
func (o *AddressStxInboundListResponse) GetResults() []InboundStxTransfer {
	if o == nil {
		var ret []InboundStxTransfer
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *AddressStxInboundListResponse) GetResultsOk() ([]InboundStxTransfer, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *AddressStxInboundListResponse) SetResults(v []InboundStxTransfer) {
	o.Results = v
}

func (o AddressStxInboundListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressStxInboundListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["limit"] = o.Limit
	toSerialize["offset"] = o.Offset
	toSerialize["total"] = o.Total
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

func (o *AddressStxInboundListResponse) UnmarshalJSON(data []byte) (err error) {
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

	varAddressStxInboundListResponse := _AddressStxInboundListResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddressStxInboundListResponse)

	if err != nil {
		return err
	}

	*o = AddressStxInboundListResponse(varAddressStxInboundListResponse)

	return err
}

type NullableAddressStxInboundListResponse struct {
	value *AddressStxInboundListResponse
	isSet bool
}

func (v NullableAddressStxInboundListResponse) Get() *AddressStxInboundListResponse {
	return v.value
}

func (v *NullableAddressStxInboundListResponse) Set(val *AddressStxInboundListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressStxInboundListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressStxInboundListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressStxInboundListResponse(val *AddressStxInboundListResponse) *NullableAddressStxInboundListResponse {
	return &NullableAddressStxInboundListResponse{value: val, isSet: true}
}

func (v NullableAddressStxInboundListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressStxInboundListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
