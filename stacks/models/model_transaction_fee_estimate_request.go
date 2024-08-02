package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the TransactionFeeEstimateRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &TransactionFeeEstimateRequest{}

// TransactionFeeEstimateRequest POST request for estimated fee
type TransactionFeeEstimateRequest struct {
	TransactionPayload string `json:"transaction_payload"`
	EstimatedLen       *int32 `json:"estimated_len,omitempty"`
}

type _TransactionFeeEstimateRequest TransactionFeeEstimateRequest

// NewTransactionFeeEstimateRequest instantiates a new TransactionFeeEstimateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionFeeEstimateRequest(transactionPayload string) *TransactionFeeEstimateRequest {
	this := TransactionFeeEstimateRequest{}
	this.TransactionPayload = transactionPayload
	return &this
}

// NewTransactionFeeEstimateRequestWithDefaults instantiates a new TransactionFeeEstimateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionFeeEstimateRequestWithDefaults() *TransactionFeeEstimateRequest {
	this := TransactionFeeEstimateRequest{}
	return &this
}

// GetTransactionPayload returns the TransactionPayload field value
func (o *TransactionFeeEstimateRequest) GetTransactionPayload() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionPayload
}

// GetTransactionPayloadOk returns a tuple with the TransactionPayload field value
// and a boolean to check if the value has been set.
func (o *TransactionFeeEstimateRequest) GetTransactionPayloadOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionPayload, true
}

// SetTransactionPayload sets field value
func (o *TransactionFeeEstimateRequest) SetTransactionPayload(v string) {
	o.TransactionPayload = v
}

// GetEstimatedLen returns the EstimatedLen field value if set, zero value otherwise.
func (o *TransactionFeeEstimateRequest) GetEstimatedLen() int32 {
	if o == nil || utils.IsNil(o.EstimatedLen) {
		var ret int32
		return ret
	}
	return *o.EstimatedLen
}

// GetEstimatedLenOk returns a tuple with the EstimatedLen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionFeeEstimateRequest) GetEstimatedLenOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.EstimatedLen) {
		return nil, false
	}
	return o.EstimatedLen, true
}

// HasEstimatedLen returns a boolean if a field has been set.
func (o *TransactionFeeEstimateRequest) HasEstimatedLen() bool {
	if o != nil && !utils.IsNil(o.EstimatedLen) {
		return true
	}

	return false
}

// SetEstimatedLen gets a reference to the given int32 and assigns it to the EstimatedLen field.
func (o *TransactionFeeEstimateRequest) SetEstimatedLen(v int32) {
	o.EstimatedLen = &v
}

func (o TransactionFeeEstimateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionFeeEstimateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["transaction_payload"] = o.TransactionPayload
	if !utils.IsNil(o.EstimatedLen) {
		toSerialize["estimated_len"] = o.EstimatedLen
	}
	return toSerialize, nil
}

func (o *TransactionFeeEstimateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"transaction_payload",
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

	varTransactionFeeEstimateRequest := _TransactionFeeEstimateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransactionFeeEstimateRequest)

	if err != nil {
		return err
	}

	*o = TransactionFeeEstimateRequest(varTransactionFeeEstimateRequest)

	return err
}

type NullableTransactionFeeEstimateRequest struct {
	value *TransactionFeeEstimateRequest
	isSet bool
}

func (v NullableTransactionFeeEstimateRequest) Get() *TransactionFeeEstimateRequest {
	return v.value
}

func (v *NullableTransactionFeeEstimateRequest) Set(val *TransactionFeeEstimateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionFeeEstimateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionFeeEstimateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionFeeEstimateRequest(val *TransactionFeeEstimateRequest) *NullableTransactionFeeEstimateRequest {
	return &NullableTransactionFeeEstimateRequest{value: val, isSet: true}
}

func (v NullableTransactionFeeEstimateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionFeeEstimateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
