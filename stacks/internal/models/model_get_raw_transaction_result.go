package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the GetRawTransactionResult type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &GetRawTransactionResult{}

// GetRawTransactionResult GET raw transaction
type GetRawTransactionResult struct {
	// A hex encoded serialized transaction
	RawTx string `json:"raw_tx"`
}

type _GetRawTransactionResult GetRawTransactionResult

// NewGetRawTransactionResult instantiates a new GetRawTransactionResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRawTransactionResult(rawTx string) *GetRawTransactionResult {
	this := GetRawTransactionResult{}
	this.RawTx = rawTx
	return &this
}

// NewGetRawTransactionResultWithDefaults instantiates a new GetRawTransactionResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRawTransactionResultWithDefaults() *GetRawTransactionResult {
	this := GetRawTransactionResult{}
	return &this
}

// GetRawTx returns the RawTx field value
func (o *GetRawTransactionResult) GetRawTx() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RawTx
}

// GetRawTxOk returns a tuple with the RawTx field value
// and a boolean to check if the value has been set.
func (o *GetRawTransactionResult) GetRawTxOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RawTx, true
}

// SetRawTx sets field value
func (o *GetRawTransactionResult) SetRawTx(v string) {
	o.RawTx = v
}

func (o GetRawTransactionResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRawTransactionResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["raw_tx"] = o.RawTx
	return toSerialize, nil
}

func (o *GetRawTransactionResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"raw_tx",
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

	varGetRawTransactionResult := _GetRawTransactionResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetRawTransactionResult)

	if err != nil {
		return err
	}

	*o = GetRawTransactionResult(varGetRawTransactionResult)

	return err
}

type NullableGetRawTransactionResult struct {
	value *GetRawTransactionResult
	isSet bool
}

func (v NullableGetRawTransactionResult) Get() *GetRawTransactionResult {
	return v.value
}

func (v *NullableGetRawTransactionResult) Set(val *GetRawTransactionResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRawTransactionResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRawTransactionResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRawTransactionResult(val *GetRawTransactionResult) *NullableGetRawTransactionResult {
	return &NullableGetRawTransactionResult{value: val, isSet: true}
}

func (v NullableGetRawTransactionResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRawTransactionResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
