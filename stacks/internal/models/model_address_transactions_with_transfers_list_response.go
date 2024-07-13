package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the AddressTransactionsWithTransfersListResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AddressTransactionsWithTransfersListResponse{}

// AddressTransactionsWithTransfersListResponse GET request that returns account transactions
type AddressTransactionsWithTransfersListResponse struct {
	Limit   int32                             `json:"limit"`
	Offset  int32                             `json:"offset"`
	Total   int32                             `json:"total"`
	Results []AddressTransactionWithTransfers `json:"results"`
}

type _AddressTransactionsWithTransfersListResponse AddressTransactionsWithTransfersListResponse

// NewAddressTransactionsWithTransfersListResponse instantiates a new AddressTransactionsWithTransfersListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressTransactionsWithTransfersListResponse(limit int32, offset int32, total int32, results []AddressTransactionWithTransfers) *AddressTransactionsWithTransfersListResponse {
	this := AddressTransactionsWithTransfersListResponse{}
	this.Limit = limit
	this.Offset = offset
	this.Total = total
	this.Results = results
	return &this
}

// NewAddressTransactionsWithTransfersListResponseWithDefaults instantiates a new AddressTransactionsWithTransfersListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressTransactionsWithTransfersListResponseWithDefaults() *AddressTransactionsWithTransfersListResponse {
	this := AddressTransactionsWithTransfersListResponse{}
	return &this
}

// GetLimit returns the Limit field value
func (o *AddressTransactionsWithTransfersListResponse) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *AddressTransactionsWithTransfersListResponse) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *AddressTransactionsWithTransfersListResponse) SetLimit(v int32) {
	o.Limit = v
}

// GetOffset returns the Offset field value
func (o *AddressTransactionsWithTransfersListResponse) GetOffset() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *AddressTransactionsWithTransfersListResponse) GetOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value
func (o *AddressTransactionsWithTransfersListResponse) SetOffset(v int32) {
	o.Offset = v
}

// GetTotal returns the Total field value
func (o *AddressTransactionsWithTransfersListResponse) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *AddressTransactionsWithTransfersListResponse) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *AddressTransactionsWithTransfersListResponse) SetTotal(v int32) {
	o.Total = v
}

// GetResults returns the Results field value
func (o *AddressTransactionsWithTransfersListResponse) GetResults() []AddressTransactionWithTransfers {
	if o == nil {
		var ret []AddressTransactionWithTransfers
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *AddressTransactionsWithTransfersListResponse) GetResultsOk() ([]AddressTransactionWithTransfers, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *AddressTransactionsWithTransfersListResponse) SetResults(v []AddressTransactionWithTransfers) {
	o.Results = v
}

func (o AddressTransactionsWithTransfersListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressTransactionsWithTransfersListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["limit"] = o.Limit
	toSerialize["offset"] = o.Offset
	toSerialize["total"] = o.Total
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

func (o *AddressTransactionsWithTransfersListResponse) UnmarshalJSON(data []byte) (err error) {
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

	varAddressTransactionsWithTransfersListResponse := _AddressTransactionsWithTransfersListResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddressTransactionsWithTransfersListResponse)

	if err != nil {
		return err
	}

	*o = AddressTransactionsWithTransfersListResponse(varAddressTransactionsWithTransfersListResponse)

	return err
}

type NullableAddressTransactionsWithTransfersListResponse struct {
	value *AddressTransactionsWithTransfersListResponse
	isSet bool
}

func (v NullableAddressTransactionsWithTransfersListResponse) Get() *AddressTransactionsWithTransfersListResponse {
	return v.value
}

func (v *NullableAddressTransactionsWithTransfersListResponse) Set(val *AddressTransactionsWithTransfersListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressTransactionsWithTransfersListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressTransactionsWithTransfersListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressTransactionsWithTransfersListResponse(val *AddressTransactionsWithTransfersListResponse) *NullableAddressTransactionsWithTransfersListResponse {
	return &NullableAddressTransactionsWithTransfersListResponse{value: val, isSet: true}
}

func (v NullableAddressTransactionsWithTransfersListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressTransactionsWithTransfersListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
