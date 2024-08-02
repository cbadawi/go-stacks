package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the AddressTransactionEventListResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AddressTransactionEventListResponse{}

// AddressTransactionEventListResponse GET Address Transaction Events
type AddressTransactionEventListResponse struct {
	Limit   int32                     `json:"limit"`
	Offset  int32                     `json:"offset"`
	Total   int32                     `json:"total"`
	Results []AddressTransactionEvent `json:"results"`
}

type _AddressTransactionEventListResponse AddressTransactionEventListResponse

// NewAddressTransactionEventListResponse instantiates a new AddressTransactionEventListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressTransactionEventListResponse(limit int32, offset int32, total int32, results []AddressTransactionEvent) *AddressTransactionEventListResponse {
	this := AddressTransactionEventListResponse{}
	this.Limit = limit
	this.Offset = offset
	this.Total = total
	this.Results = results
	return &this
}

// NewAddressTransactionEventListResponseWithDefaults instantiates a new AddressTransactionEventListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressTransactionEventListResponseWithDefaults() *AddressTransactionEventListResponse {
	this := AddressTransactionEventListResponse{}
	return &this
}

// GetLimit returns the Limit field value
func (o *AddressTransactionEventListResponse) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *AddressTransactionEventListResponse) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *AddressTransactionEventListResponse) SetLimit(v int32) {
	o.Limit = v
}

// GetOffset returns the Offset field value
func (o *AddressTransactionEventListResponse) GetOffset() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *AddressTransactionEventListResponse) GetOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value
func (o *AddressTransactionEventListResponse) SetOffset(v int32) {
	o.Offset = v
}

// GetTotal returns the Total field value
func (o *AddressTransactionEventListResponse) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *AddressTransactionEventListResponse) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *AddressTransactionEventListResponse) SetTotal(v int32) {
	o.Total = v
}

// GetResults returns the Results field value
func (o *AddressTransactionEventListResponse) GetResults() []AddressTransactionEvent {
	if o == nil {
		var ret []AddressTransactionEvent
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *AddressTransactionEventListResponse) GetResultsOk() ([]AddressTransactionEvent, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *AddressTransactionEventListResponse) SetResults(v []AddressTransactionEvent) {
	o.Results = v
}

func (o AddressTransactionEventListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressTransactionEventListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["limit"] = o.Limit
	toSerialize["offset"] = o.Offset
	toSerialize["total"] = o.Total
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

func (o *AddressTransactionEventListResponse) UnmarshalJSON(data []byte) (err error) {
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

	varAddressTransactionEventListResponse := _AddressTransactionEventListResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddressTransactionEventListResponse)

	if err != nil {
		return err
	}

	*o = AddressTransactionEventListResponse(varAddressTransactionEventListResponse)

	return err
}

type NullableAddressTransactionEventListResponse struct {
	value *AddressTransactionEventListResponse
	isSet bool
}

func (v NullableAddressTransactionEventListResponse) Get() *AddressTransactionEventListResponse {
	return v.value
}

func (v *NullableAddressTransactionEventListResponse) Set(val *AddressTransactionEventListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressTransactionEventListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressTransactionEventListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressTransactionEventListResponse(val *AddressTransactionEventListResponse) *NullableAddressTransactionEventListResponse {
	return &NullableAddressTransactionEventListResponse{value: val, isSet: true}
}

func (v NullableAddressTransactionEventListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressTransactionEventListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
