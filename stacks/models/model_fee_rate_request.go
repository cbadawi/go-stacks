package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the FeeRateRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &FeeRateRequest{}

// FeeRateRequest Request to fetch fee for a transaction
type FeeRateRequest struct {
	// A serialized transaction
	Transaction string `json:"transaction"`
}

type _FeeRateRequest FeeRateRequest

// NewFeeRateRequest instantiates a new FeeRateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeeRateRequest(transaction string) *FeeRateRequest {
	this := FeeRateRequest{}
	this.Transaction = transaction
	return &this
}

// NewFeeRateRequestWithDefaults instantiates a new FeeRateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeeRateRequestWithDefaults() *FeeRateRequest {
	this := FeeRateRequest{}
	return &this
}

// GetTransaction returns the Transaction field value
func (o *FeeRateRequest) GetTransaction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Transaction
}

// GetTransactionOk returns a tuple with the Transaction field value
// and a boolean to check if the value has been set.
func (o *FeeRateRequest) GetTransactionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Transaction, true
}

// SetTransaction sets field value
func (o *FeeRateRequest) SetTransaction(v string) {
	o.Transaction = v
}

func (o FeeRateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeeRateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["transaction"] = o.Transaction
	return toSerialize, nil
}

func (o *FeeRateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"transaction",
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

	varFeeRateRequest := _FeeRateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFeeRateRequest)

	if err != nil {
		return err
	}

	*o = FeeRateRequest(varFeeRateRequest)

	return err
}

type NullableFeeRateRequest struct {
	value *FeeRateRequest
	isSet bool
}

func (v NullableFeeRateRequest) Get() *FeeRateRequest {
	return v.value
}

func (v *NullableFeeRateRequest) Set(val *FeeRateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFeeRateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFeeRateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeeRateRequest(val *FeeRateRequest) *NullableFeeRateRequest {
	return &NullableFeeRateRequest{value: val, isSet: true}
}

func (v NullableFeeRateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeeRateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
