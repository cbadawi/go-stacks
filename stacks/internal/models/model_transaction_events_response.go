package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the TransactionEventsResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &TransactionEventsResponse{}

// TransactionEventsResponse GET event for the given transaction
type TransactionEventsResponse struct {
	Limit   int32              `json:"limit"`
	Offset  int32              `json:"offset"`
	Results []TransactionEvent `json:"results"`
}

type _TransactionEventsResponse TransactionEventsResponse

// NewTransactionEventsResponse instantiates a new TransactionEventsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionEventsResponse(limit int32, offset int32, results []TransactionEvent) *TransactionEventsResponse {
	this := TransactionEventsResponse{}
	this.Limit = limit
	this.Offset = offset
	this.Results = results
	return &this
}

// NewTransactionEventsResponseWithDefaults instantiates a new TransactionEventsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionEventsResponseWithDefaults() *TransactionEventsResponse {
	this := TransactionEventsResponse{}
	return &this
}

// GetLimit returns the Limit field value
func (o *TransactionEventsResponse) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *TransactionEventsResponse) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *TransactionEventsResponse) SetLimit(v int32) {
	o.Limit = v
}

// GetOffset returns the Offset field value
func (o *TransactionEventsResponse) GetOffset() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *TransactionEventsResponse) GetOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value
func (o *TransactionEventsResponse) SetOffset(v int32) {
	o.Offset = v
}

// GetResults returns the Results field value
func (o *TransactionEventsResponse) GetResults() []TransactionEvent {
	if o == nil {
		var ret []TransactionEvent
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *TransactionEventsResponse) GetResultsOk() ([]TransactionEvent, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *TransactionEventsResponse) SetResults(v []TransactionEvent) {
	o.Results = v
}

func (o TransactionEventsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionEventsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["limit"] = o.Limit
	toSerialize["offset"] = o.Offset
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

func (o *TransactionEventsResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"limit",
		"offset",
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

	varTransactionEventsResponse := _TransactionEventsResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransactionEventsResponse)

	if err != nil {
		return err
	}

	*o = TransactionEventsResponse(varTransactionEventsResponse)

	return err
}

type NullableTransactionEventsResponse struct {
	value *TransactionEventsResponse
	isSet bool
}

func (v NullableTransactionEventsResponse) Get() *TransactionEventsResponse {
	return v.value
}

func (v *NullableTransactionEventsResponse) Set(val *TransactionEventsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionEventsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionEventsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionEventsResponse(val *TransactionEventsResponse) *NullableTransactionEventsResponse {
	return &NullableTransactionEventsResponse{value: val, isSet: true}
}

func (v NullableTransactionEventsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionEventsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
