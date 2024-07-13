package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the UnanchoredTransactionListResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &UnanchoredTransactionListResponse{}

// UnanchoredTransactionListResponse GET request that returns unanchored transactions
type UnanchoredTransactionListResponse struct {
	// The number of unanchored transactions available
	Total   int32         `json:"total"`
	Results []Transaction `json:"results"`
}

type _UnanchoredTransactionListResponse UnanchoredTransactionListResponse

// NewUnanchoredTransactionListResponse instantiates a new UnanchoredTransactionListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnanchoredTransactionListResponse(total int32, results []Transaction) *UnanchoredTransactionListResponse {
	this := UnanchoredTransactionListResponse{}
	this.Total = total
	this.Results = results
	return &this
}

// NewUnanchoredTransactionListResponseWithDefaults instantiates a new UnanchoredTransactionListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnanchoredTransactionListResponseWithDefaults() *UnanchoredTransactionListResponse {
	this := UnanchoredTransactionListResponse{}
	return &this
}

// GetTotal returns the Total field value
func (o *UnanchoredTransactionListResponse) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *UnanchoredTransactionListResponse) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *UnanchoredTransactionListResponse) SetTotal(v int32) {
	o.Total = v
}

// GetResults returns the Results field value
func (o *UnanchoredTransactionListResponse) GetResults() []Transaction {
	if o == nil {
		var ret []Transaction
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *UnanchoredTransactionListResponse) GetResultsOk() ([]Transaction, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *UnanchoredTransactionListResponse) SetResults(v []Transaction) {
	o.Results = v
}

func (o UnanchoredTransactionListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnanchoredTransactionListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["total"] = o.Total
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

func (o *UnanchoredTransactionListResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varUnanchoredTransactionListResponse := _UnanchoredTransactionListResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnanchoredTransactionListResponse)

	if err != nil {
		return err
	}

	*o = UnanchoredTransactionListResponse(varUnanchoredTransactionListResponse)

	return err
}

type NullableUnanchoredTransactionListResponse struct {
	value *UnanchoredTransactionListResponse
	isSet bool
}

func (v NullableUnanchoredTransactionListResponse) Get() *UnanchoredTransactionListResponse {
	return v.value
}

func (v *NullableUnanchoredTransactionListResponse) Set(val *UnanchoredTransactionListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUnanchoredTransactionListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUnanchoredTransactionListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnanchoredTransactionListResponse(val *UnanchoredTransactionListResponse) *NullableUnanchoredTransactionListResponse {
	return &NullableUnanchoredTransactionListResponse{value: val, isSet: true}
}

func (v NullableUnanchoredTransactionListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnanchoredTransactionListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
