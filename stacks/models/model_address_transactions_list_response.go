package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the AddressTransactionsListResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AddressTransactionsListResponse{}

// AddressTransactionsListResponse GET request that returns account transactions
type AddressTransactionsListResponse struct {
	Limit   int32                                         `json:"limit"`
	Offset  int32                                         `json:"offset"`
	Total   int32                                         `json:"total"`
	Results []AddressTransactionsListResponseResultsInner `json:"results"`
}

type _AddressTransactionsListResponse AddressTransactionsListResponse

// NewAddressTransactionsListResponse instantiates a new AddressTransactionsListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressTransactionsListResponse(limit int32, offset int32, total int32, results []AddressTransactionsListResponseResultsInner) *AddressTransactionsListResponse {
	this := AddressTransactionsListResponse{}
	this.Limit = limit
	this.Offset = offset
	this.Total = total
	this.Results = results
	return &this
}

// NewAddressTransactionsListResponseWithDefaults instantiates a new AddressTransactionsListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressTransactionsListResponseWithDefaults() *AddressTransactionsListResponse {
	this := AddressTransactionsListResponse{}
	return &this
}

// GetLimit returns the Limit field value
func (o *AddressTransactionsListResponse) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *AddressTransactionsListResponse) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *AddressTransactionsListResponse) SetLimit(v int32) {
	o.Limit = v
}

// GetOffset returns the Offset field value
func (o *AddressTransactionsListResponse) GetOffset() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *AddressTransactionsListResponse) GetOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value
func (o *AddressTransactionsListResponse) SetOffset(v int32) {
	o.Offset = v
}

// GetTotal returns the Total field value
func (o *AddressTransactionsListResponse) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *AddressTransactionsListResponse) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *AddressTransactionsListResponse) SetTotal(v int32) {
	o.Total = v
}

// GetResults returns the Results field value
func (o *AddressTransactionsListResponse) GetResults() []AddressTransactionsListResponseResultsInner {
	if o == nil {
		var ret []AddressTransactionsListResponseResultsInner
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *AddressTransactionsListResponse) GetResultsOk() ([]AddressTransactionsListResponseResultsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *AddressTransactionsListResponse) SetResults(v []AddressTransactionsListResponseResultsInner) {
	o.Results = v
}

func (o AddressTransactionsListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressTransactionsListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["limit"] = o.Limit
	toSerialize["offset"] = o.Offset
	toSerialize["total"] = o.Total
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

func (o *AddressTransactionsListResponse) UnmarshalJSON(data []byte) (err error) {
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

	varAddressTransactionsListResponse := _AddressTransactionsListResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddressTransactionsListResponse)

	if err != nil {
		return err
	}

	*o = AddressTransactionsListResponse(varAddressTransactionsListResponse)

	return err
}

type NullableAddressTransactionsListResponse struct {
	value *AddressTransactionsListResponse
	isSet bool
}

func (v NullableAddressTransactionsListResponse) Get() *AddressTransactionsListResponse {
	return v.value
}

func (v *NullableAddressTransactionsListResponse) Set(val *AddressTransactionsListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressTransactionsListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressTransactionsListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressTransactionsListResponse(val *AddressTransactionsListResponse) *NullableAddressTransactionsListResponse {
	return &NullableAddressTransactionsListResponse{value: val, isSet: true}
}

func (v NullableAddressTransactionsListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressTransactionsListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
