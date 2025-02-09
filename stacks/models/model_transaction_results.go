package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the TransactionResults type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &TransactionResults{}

// TransactionResults GET request that returns transactions
type TransactionResults struct {
	// The number of transactions to return
	Limit int32 `json:"limit"`
	// The number to transactions to skip (starting at `0`)
	Offset int32 `json:"offset"`
	// The number of transactions available
	Total   int32         `json:"total"`
	Results []Transaction `json:"results"`
}

type _TransactionResults TransactionResults

// NewTransactionResults instantiates a new TransactionResults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionResults(limit int32, offset int32, total int32, results []Transaction) *TransactionResults {
	this := TransactionResults{}
	this.Limit = limit
	this.Offset = offset
	this.Total = total
	this.Results = results
	return &this
}

// NewTransactionResultsWithDefaults instantiates a new TransactionResults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionResultsWithDefaults() *TransactionResults {
	this := TransactionResults{}
	return &this
}

// GetLimit returns the Limit field value
func (o *TransactionResults) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *TransactionResults) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *TransactionResults) SetLimit(v int32) {
	o.Limit = v
}

// GetOffset returns the Offset field value
func (o *TransactionResults) GetOffset() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *TransactionResults) GetOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value
func (o *TransactionResults) SetOffset(v int32) {
	o.Offset = v
}

// GetTotal returns the Total field value
func (o *TransactionResults) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *TransactionResults) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *TransactionResults) SetTotal(v int32) {
	o.Total = v
}

// GetResults returns the Results field value
func (o *TransactionResults) GetResults() []Transaction {
	if o == nil {
		var ret []Transaction
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *TransactionResults) GetResultsOk() ([]Transaction, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *TransactionResults) SetResults(v []Transaction) {
	o.Results = v
}

func (o TransactionResults) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionResults) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["limit"] = o.Limit
	toSerialize["offset"] = o.Offset
	toSerialize["total"] = o.Total
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

func (o *TransactionResults) UnmarshalJSON(data []byte) (err error) {
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

	varTransactionResults := _TransactionResults{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransactionResults)

	if err != nil {
		return err
	}

	*o = TransactionResults(varTransactionResults)

	return err
}

type NullableTransactionResults struct {
	value *TransactionResults
	isSet bool
}

func (v NullableTransactionResults) Get() *TransactionResults {
	return v.value
}

func (v *NullableTransactionResults) Set(val *TransactionResults) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionResults) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionResults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionResults(val *TransactionResults) *NullableTransactionResults {
	return &NullableTransactionResults{value: val, isSet: true}
}

func (v NullableTransactionResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionResults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
