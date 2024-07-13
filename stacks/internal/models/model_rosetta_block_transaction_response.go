package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the RosettaBlockTransactionResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RosettaBlockTransactionResponse{}

// RosettaBlockTransactionResponse A BlockTransactionResponse contains information about a block transaction.
type RosettaBlockTransactionResponse struct {
	Transaction RosettaTransaction `json:"transaction"`
}

type _RosettaBlockTransactionResponse RosettaBlockTransactionResponse

// NewRosettaBlockTransactionResponse instantiates a new RosettaBlockTransactionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRosettaBlockTransactionResponse(transaction RosettaTransaction) *RosettaBlockTransactionResponse {
	this := RosettaBlockTransactionResponse{}
	this.Transaction = transaction
	return &this
}

// NewRosettaBlockTransactionResponseWithDefaults instantiates a new RosettaBlockTransactionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRosettaBlockTransactionResponseWithDefaults() *RosettaBlockTransactionResponse {
	this := RosettaBlockTransactionResponse{}
	return &this
}

// GetTransaction returns the Transaction field value
func (o *RosettaBlockTransactionResponse) GetTransaction() RosettaTransaction {
	if o == nil {
		var ret RosettaTransaction
		return ret
	}

	return o.Transaction
}

// GetTransactionOk returns a tuple with the Transaction field value
// and a boolean to check if the value has been set.
func (o *RosettaBlockTransactionResponse) GetTransactionOk() (*RosettaTransaction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Transaction, true
}

// SetTransaction sets field value
func (o *RosettaBlockTransactionResponse) SetTransaction(v RosettaTransaction) {
	o.Transaction = v
}

func (o RosettaBlockTransactionResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RosettaBlockTransactionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["transaction"] = o.Transaction
	return toSerialize, nil
}

func (o *RosettaBlockTransactionResponse) UnmarshalJSON(data []byte) (err error) {
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

	varRosettaBlockTransactionResponse := _RosettaBlockTransactionResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRosettaBlockTransactionResponse)

	if err != nil {
		return err
	}

	*o = RosettaBlockTransactionResponse(varRosettaBlockTransactionResponse)

	return err
}

type NullableRosettaBlockTransactionResponse struct {
	value *RosettaBlockTransactionResponse
	isSet bool
}

func (v NullableRosettaBlockTransactionResponse) Get() *RosettaBlockTransactionResponse {
	return v.value
}

func (v *NullableRosettaBlockTransactionResponse) Set(val *RosettaBlockTransactionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRosettaBlockTransactionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRosettaBlockTransactionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRosettaBlockTransactionResponse(val *RosettaBlockTransactionResponse) *NullableRosettaBlockTransactionResponse {
	return &NullableRosettaBlockTransactionResponse{value: val, isSet: true}
}

func (v NullableRosettaBlockTransactionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRosettaBlockTransactionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
