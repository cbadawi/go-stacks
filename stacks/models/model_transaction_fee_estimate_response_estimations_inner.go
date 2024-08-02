package models

import (
	"encoding/json"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the TransactionFeeEstimateResponseEstimationsInner type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &TransactionFeeEstimateResponseEstimationsInner{}

// TransactionFeeEstimateResponseEstimationsInner struct for TransactionFeeEstimateResponseEstimationsInner
type TransactionFeeEstimateResponseEstimationsInner struct {
	FeeRate *float32 `json:"fee_rate,omitempty"`
	Fee     *float32 `json:"fee,omitempty"`
}

// NewTransactionFeeEstimateResponseEstimationsInner instantiates a new TransactionFeeEstimateResponseEstimationsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionFeeEstimateResponseEstimationsInner() *TransactionFeeEstimateResponseEstimationsInner {
	this := TransactionFeeEstimateResponseEstimationsInner{}
	return &this
}

// NewTransactionFeeEstimateResponseEstimationsInnerWithDefaults instantiates a new TransactionFeeEstimateResponseEstimationsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionFeeEstimateResponseEstimationsInnerWithDefaults() *TransactionFeeEstimateResponseEstimationsInner {
	this := TransactionFeeEstimateResponseEstimationsInner{}
	return &this
}

// GetFeeRate returns the FeeRate field value if set, zero value otherwise.
func (o *TransactionFeeEstimateResponseEstimationsInner) GetFeeRate() float32 {
	if o == nil || utils.IsNil(o.FeeRate) {
		var ret float32
		return ret
	}
	return *o.FeeRate
}

// GetFeeRateOk returns a tuple with the FeeRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionFeeEstimateResponseEstimationsInner) GetFeeRateOk() (*float32, bool) {
	if o == nil || utils.IsNil(o.FeeRate) {
		return nil, false
	}
	return o.FeeRate, true
}

// HasFeeRate returns a boolean if a field has been set.
func (o *TransactionFeeEstimateResponseEstimationsInner) HasFeeRate() bool {
	if o != nil && !utils.IsNil(o.FeeRate) {
		return true
	}

	return false
}

// SetFeeRate gets a reference to the given float32 and assigns it to the FeeRate field.
func (o *TransactionFeeEstimateResponseEstimationsInner) SetFeeRate(v float32) {
	o.FeeRate = &v
}

// GetFee returns the Fee field value if set, zero value otherwise.
func (o *TransactionFeeEstimateResponseEstimationsInner) GetFee() float32 {
	if o == nil || utils.IsNil(o.Fee) {
		var ret float32
		return ret
	}
	return *o.Fee
}

// GetFeeOk returns a tuple with the Fee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionFeeEstimateResponseEstimationsInner) GetFeeOk() (*float32, bool) {
	if o == nil || utils.IsNil(o.Fee) {
		return nil, false
	}
	return o.Fee, true
}

// HasFee returns a boolean if a field has been set.
func (o *TransactionFeeEstimateResponseEstimationsInner) HasFee() bool {
	if o != nil && !utils.IsNil(o.Fee) {
		return true
	}

	return false
}

// SetFee gets a reference to the given float32 and assigns it to the Fee field.
func (o *TransactionFeeEstimateResponseEstimationsInner) SetFee(v float32) {
	o.Fee = &v
}

func (o TransactionFeeEstimateResponseEstimationsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionFeeEstimateResponseEstimationsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.FeeRate) {
		toSerialize["fee_rate"] = o.FeeRate
	}
	if !utils.IsNil(o.Fee) {
		toSerialize["fee"] = o.Fee
	}
	return toSerialize, nil
}

type NullableTransactionFeeEstimateResponseEstimationsInner struct {
	value *TransactionFeeEstimateResponseEstimationsInner
	isSet bool
}

func (v NullableTransactionFeeEstimateResponseEstimationsInner) Get() *TransactionFeeEstimateResponseEstimationsInner {
	return v.value
}

func (v *NullableTransactionFeeEstimateResponseEstimationsInner) Set(val *TransactionFeeEstimateResponseEstimationsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionFeeEstimateResponseEstimationsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionFeeEstimateResponseEstimationsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionFeeEstimateResponseEstimationsInner(val *TransactionFeeEstimateResponseEstimationsInner) *NullableTransactionFeeEstimateResponseEstimationsInner {
	return &NullableTransactionFeeEstimateResponseEstimationsInner{value: val, isSet: true}
}

func (v NullableTransactionFeeEstimateResponseEstimationsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionFeeEstimateResponseEstimationsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
