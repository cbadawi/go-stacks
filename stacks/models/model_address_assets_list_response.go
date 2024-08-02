package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the AddressAssetsListResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AddressAssetsListResponse{}

// AddressAssetsListResponse GET request that returns address assets
type AddressAssetsListResponse struct {
	Limit   int32              `json:"limit"`
	Offset  int32              `json:"offset"`
	Total   int32              `json:"total"`
	Results []TransactionEvent `json:"results"`
}

type _AddressAssetsListResponse AddressAssetsListResponse

// NewAddressAssetsListResponse instantiates a new AddressAssetsListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressAssetsListResponse(limit int32, offset int32, total int32, results []TransactionEvent) *AddressAssetsListResponse {
	this := AddressAssetsListResponse{}
	this.Limit = limit
	this.Offset = offset
	this.Total = total
	this.Results = results
	return &this
}

// NewAddressAssetsListResponseWithDefaults instantiates a new AddressAssetsListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressAssetsListResponseWithDefaults() *AddressAssetsListResponse {
	this := AddressAssetsListResponse{}
	return &this
}

// GetLimit returns the Limit field value
func (o *AddressAssetsListResponse) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *AddressAssetsListResponse) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *AddressAssetsListResponse) SetLimit(v int32) {
	o.Limit = v
}

// GetOffset returns the Offset field value
func (o *AddressAssetsListResponse) GetOffset() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *AddressAssetsListResponse) GetOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value
func (o *AddressAssetsListResponse) SetOffset(v int32) {
	o.Offset = v
}

// GetTotal returns the Total field value
func (o *AddressAssetsListResponse) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *AddressAssetsListResponse) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *AddressAssetsListResponse) SetTotal(v int32) {
	o.Total = v
}

// GetResults returns the Results field value
func (o *AddressAssetsListResponse) GetResults() []TransactionEvent {
	if o == nil {
		var ret []TransactionEvent
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *AddressAssetsListResponse) GetResultsOk() ([]TransactionEvent, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *AddressAssetsListResponse) SetResults(v []TransactionEvent) {
	o.Results = v
}

func (o AddressAssetsListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressAssetsListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["limit"] = o.Limit
	toSerialize["offset"] = o.Offset
	toSerialize["total"] = o.Total
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

func (o *AddressAssetsListResponse) UnmarshalJSON(data []byte) (err error) {
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

	varAddressAssetsListResponse := _AddressAssetsListResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddressAssetsListResponse)

	if err != nil {
		return err
	}

	*o = AddressAssetsListResponse(varAddressAssetsListResponse)

	return err
}

type NullableAddressAssetsListResponse struct {
	value *AddressAssetsListResponse
	isSet bool
}

func (v NullableAddressAssetsListResponse) Get() *AddressAssetsListResponse {
	return v.value
}

func (v *NullableAddressAssetsListResponse) Set(val *AddressAssetsListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressAssetsListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressAssetsListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressAssetsListResponse(val *AddressAssetsListResponse) *NullableAddressAssetsListResponse {
	return &NullableAddressAssetsListResponse{value: val, isSet: true}
}

func (v NullableAddressAssetsListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressAssetsListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
