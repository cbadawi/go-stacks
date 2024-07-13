package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the TransactionFeeEstimateResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &TransactionFeeEstimateResponse{}

// TransactionFeeEstimateResponse POST response for estimated fee
type TransactionFeeEstimateResponse struct {
	EstimatedCostScalar    int32                                            `json:"estimated_cost_scalar"`
	CostScalarChangeByByte *float32                                         `json:"cost_scalar_change_by_byte,omitempty"`
	EstimatedCost          TransactionFeeEstimateResponseEstimatedCost      `json:"estimated_cost"`
	Estimations            []TransactionFeeEstimateResponseEstimationsInner `json:"estimations,omitempty"`
}

type _TransactionFeeEstimateResponse TransactionFeeEstimateResponse

// NewTransactionFeeEstimateResponse instantiates a new TransactionFeeEstimateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionFeeEstimateResponse(estimatedCostScalar int32, estimatedCost TransactionFeeEstimateResponseEstimatedCost) *TransactionFeeEstimateResponse {
	this := TransactionFeeEstimateResponse{}
	this.EstimatedCostScalar = estimatedCostScalar
	this.EstimatedCost = estimatedCost
	return &this
}

// NewTransactionFeeEstimateResponseWithDefaults instantiates a new TransactionFeeEstimateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionFeeEstimateResponseWithDefaults() *TransactionFeeEstimateResponse {
	this := TransactionFeeEstimateResponse{}
	return &this
}

// GetEstimatedCostScalar returns the EstimatedCostScalar field value
func (o *TransactionFeeEstimateResponse) GetEstimatedCostScalar() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EstimatedCostScalar
}

// GetEstimatedCostScalarOk returns a tuple with the EstimatedCostScalar field value
// and a boolean to check if the value has been set.
func (o *TransactionFeeEstimateResponse) GetEstimatedCostScalarOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EstimatedCostScalar, true
}

// SetEstimatedCostScalar sets field value
func (o *TransactionFeeEstimateResponse) SetEstimatedCostScalar(v int32) {
	o.EstimatedCostScalar = v
}

// GetCostScalarChangeByByte returns the CostScalarChangeByByte field value if set, zero value otherwise.
func (o *TransactionFeeEstimateResponse) GetCostScalarChangeByByte() float32 {
	if o == nil || utils.IsNil(o.CostScalarChangeByByte) {
		var ret float32
		return ret
	}
	return *o.CostScalarChangeByByte
}

// GetCostScalarChangeByByteOk returns a tuple with the CostScalarChangeByByte field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionFeeEstimateResponse) GetCostScalarChangeByByteOk() (*float32, bool) {
	if o == nil || utils.IsNil(o.CostScalarChangeByByte) {
		return nil, false
	}
	return o.CostScalarChangeByByte, true
}

// HasCostScalarChangeByByte returns a boolean if a field has been set.
func (o *TransactionFeeEstimateResponse) HasCostScalarChangeByByte() bool {
	if o != nil && !utils.IsNil(o.CostScalarChangeByByte) {
		return true
	}

	return false
}

// SetCostScalarChangeByByte gets a reference to the given float32 and assigns it to the CostScalarChangeByByte field.
func (o *TransactionFeeEstimateResponse) SetCostScalarChangeByByte(v float32) {
	o.CostScalarChangeByByte = &v
}

// GetEstimatedCost returns the EstimatedCost field value
func (o *TransactionFeeEstimateResponse) GetEstimatedCost() TransactionFeeEstimateResponseEstimatedCost {
	if o == nil {
		var ret TransactionFeeEstimateResponseEstimatedCost
		return ret
	}

	return o.EstimatedCost
}

// GetEstimatedCostOk returns a tuple with the EstimatedCost field value
// and a boolean to check if the value has been set.
func (o *TransactionFeeEstimateResponse) GetEstimatedCostOk() (*TransactionFeeEstimateResponseEstimatedCost, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EstimatedCost, true
}

// SetEstimatedCost sets field value
func (o *TransactionFeeEstimateResponse) SetEstimatedCost(v TransactionFeeEstimateResponseEstimatedCost) {
	o.EstimatedCost = v
}

// GetEstimations returns the Estimations field value if set, zero value otherwise.
func (o *TransactionFeeEstimateResponse) GetEstimations() []TransactionFeeEstimateResponseEstimationsInner {
	if o == nil || utils.IsNil(o.Estimations) {
		var ret []TransactionFeeEstimateResponseEstimationsInner
		return ret
	}
	return o.Estimations
}

// GetEstimationsOk returns a tuple with the Estimations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionFeeEstimateResponse) GetEstimationsOk() ([]TransactionFeeEstimateResponseEstimationsInner, bool) {
	if o == nil || utils.IsNil(o.Estimations) {
		return nil, false
	}
	return o.Estimations, true
}

// HasEstimations returns a boolean if a field has been set.
func (o *TransactionFeeEstimateResponse) HasEstimations() bool {
	if o != nil && !utils.IsNil(o.Estimations) {
		return true
	}

	return false
}

// SetEstimations gets a reference to the given []TransactionFeeEstimateResponseEstimationsInner and assigns it to the Estimations field.
func (o *TransactionFeeEstimateResponse) SetEstimations(v []TransactionFeeEstimateResponseEstimationsInner) {
	o.Estimations = v
}

func (o TransactionFeeEstimateResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionFeeEstimateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["estimated_cost_scalar"] = o.EstimatedCostScalar
	if !utils.IsNil(o.CostScalarChangeByByte) {
		toSerialize["cost_scalar_change_by_byte"] = o.CostScalarChangeByByte
	}
	toSerialize["estimated_cost"] = o.EstimatedCost
	if !utils.IsNil(o.Estimations) {
		toSerialize["estimations"] = o.Estimations
	}
	return toSerialize, nil
}

func (o *TransactionFeeEstimateResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"estimated_cost_scalar",
		"estimated_cost",
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

	varTransactionFeeEstimateResponse := _TransactionFeeEstimateResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransactionFeeEstimateResponse)

	if err != nil {
		return err
	}

	*o = TransactionFeeEstimateResponse(varTransactionFeeEstimateResponse)

	return err
}

type NullableTransactionFeeEstimateResponse struct {
	value *TransactionFeeEstimateResponse
	isSet bool
}

func (v NullableTransactionFeeEstimateResponse) Get() *TransactionFeeEstimateResponse {
	return v.value
}

func (v *NullableTransactionFeeEstimateResponse) Set(val *TransactionFeeEstimateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionFeeEstimateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionFeeEstimateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionFeeEstimateResponse(val *TransactionFeeEstimateResponse) *NullableTransactionFeeEstimateResponse {
	return &NullableTransactionFeeEstimateResponse{value: val, isSet: true}
}

func (v NullableTransactionFeeEstimateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionFeeEstimateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
