package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the AddressTransactionsV2ListResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AddressTransactionsV2ListResponse{}

// AddressTransactionsV2ListResponse GET Address Transactions
type AddressTransactionsV2ListResponse struct {
	Limit   int32                `json:"limit"`
	Offset  int32                `json:"offset"`
	Total   int32                `json:"total"`
	Results []AddressTransaction `json:"results"`
}

type _AddressTransactionsV2ListResponse AddressTransactionsV2ListResponse

// NewAddressTransactionsV2ListResponse instantiates a new AddressTransactionsV2ListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressTransactionsV2ListResponse(limit int32, offset int32, total int32, results []AddressTransaction) *AddressTransactionsV2ListResponse {
	this := AddressTransactionsV2ListResponse{}
	this.Limit = limit
	this.Offset = offset
	this.Total = total
	this.Results = results
	return &this
}

// NewAddressTransactionsV2ListResponseWithDefaults instantiates a new AddressTransactionsV2ListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressTransactionsV2ListResponseWithDefaults() *AddressTransactionsV2ListResponse {
	this := AddressTransactionsV2ListResponse{}
	return &this
}

// GetLimit returns the Limit field value
func (o *AddressTransactionsV2ListResponse) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *AddressTransactionsV2ListResponse) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *AddressTransactionsV2ListResponse) SetLimit(v int32) {
	o.Limit = v
}

// GetOffset returns the Offset field value
func (o *AddressTransactionsV2ListResponse) GetOffset() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *AddressTransactionsV2ListResponse) GetOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value
func (o *AddressTransactionsV2ListResponse) SetOffset(v int32) {
	o.Offset = v
}

// GetTotal returns the Total field value
func (o *AddressTransactionsV2ListResponse) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *AddressTransactionsV2ListResponse) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *AddressTransactionsV2ListResponse) SetTotal(v int32) {
	o.Total = v
}

// GetResults returns the Results field value
func (o *AddressTransactionsV2ListResponse) GetResults() []AddressTransaction {
	if o == nil {
		var ret []AddressTransaction
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *AddressTransactionsV2ListResponse) GetResultsOk() ([]AddressTransaction, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *AddressTransactionsV2ListResponse) SetResults(v []AddressTransaction) {
	o.Results = v
}

func (o AddressTransactionsV2ListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressTransactionsV2ListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["limit"] = o.Limit
	toSerialize["offset"] = o.Offset
	toSerialize["total"] = o.Total
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

func (o *AddressTransactionsV2ListResponse) UnmarshalJSON(data []byte) (err error) {
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

	varAddressTransactionsV2ListResponse := _AddressTransactionsV2ListResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddressTransactionsV2ListResponse)

	if err != nil {
		return err
	}

	*o = AddressTransactionsV2ListResponse(varAddressTransactionsV2ListResponse)

	return err
}

type NullableAddressTransactionsV2ListResponse struct {
	value *AddressTransactionsV2ListResponse
	isSet bool
}

func (v NullableAddressTransactionsV2ListResponse) Get() *AddressTransactionsV2ListResponse {
	return v.value
}

func (v *NullableAddressTransactionsV2ListResponse) Set(val *AddressTransactionsV2ListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressTransactionsV2ListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressTransactionsV2ListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressTransactionsV2ListResponse(val *AddressTransactionsV2ListResponse) *NullableAddressTransactionsV2ListResponse {
	return &NullableAddressTransactionsV2ListResponse{value: val, isSet: true}
}

func (v NullableAddressTransactionsV2ListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressTransactionsV2ListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
