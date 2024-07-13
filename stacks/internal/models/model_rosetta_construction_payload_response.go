package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the RosettaConstructionPayloadResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RosettaConstructionPayloadResponse{}

// RosettaConstructionPayloadResponse RosettaConstructionPayloadResponse is returned by /construction/payloads. It contains an unsigned transaction blob (that is usually needed to construct the a network transaction from a collection of signatures) and an array of payloads that must be signed by the caller.
type RosettaConstructionPayloadResponse struct {
	// This is an unsigned transaction blob (that is usually needed to construct the a network transaction from a collection of signatures)
	UnsignedTransaction string `json:"unsigned_transaction"`
	// An array of payloads that must be signed by the caller
	Payloads []SigningPayload `json:"payloads"`
}

type _RosettaConstructionPayloadResponse RosettaConstructionPayloadResponse

// NewRosettaConstructionPayloadResponse instantiates a new RosettaConstructionPayloadResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRosettaConstructionPayloadResponse(unsignedTransaction string, payloads []SigningPayload) *RosettaConstructionPayloadResponse {
	this := RosettaConstructionPayloadResponse{}
	this.UnsignedTransaction = unsignedTransaction
	this.Payloads = payloads
	return &this
}

// NewRosettaConstructionPayloadResponseWithDefaults instantiates a new RosettaConstructionPayloadResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRosettaConstructionPayloadResponseWithDefaults() *RosettaConstructionPayloadResponse {
	this := RosettaConstructionPayloadResponse{}
	return &this
}

// GetUnsignedTransaction returns the UnsignedTransaction field value
func (o *RosettaConstructionPayloadResponse) GetUnsignedTransaction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnsignedTransaction
}

// GetUnsignedTransactionOk returns a tuple with the UnsignedTransaction field value
// and a boolean to check if the value has been set.
func (o *RosettaConstructionPayloadResponse) GetUnsignedTransactionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnsignedTransaction, true
}

// SetUnsignedTransaction sets field value
func (o *RosettaConstructionPayloadResponse) SetUnsignedTransaction(v string) {
	o.UnsignedTransaction = v
}

// GetPayloads returns the Payloads field value
func (o *RosettaConstructionPayloadResponse) GetPayloads() []SigningPayload {
	if o == nil {
		var ret []SigningPayload
		return ret
	}

	return o.Payloads
}

// GetPayloadsOk returns a tuple with the Payloads field value
// and a boolean to check if the value has been set.
func (o *RosettaConstructionPayloadResponse) GetPayloadsOk() ([]SigningPayload, bool) {
	if o == nil {
		return nil, false
	}
	return o.Payloads, true
}

// SetPayloads sets field value
func (o *RosettaConstructionPayloadResponse) SetPayloads(v []SigningPayload) {
	o.Payloads = v
}

func (o RosettaConstructionPayloadResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RosettaConstructionPayloadResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["unsigned_transaction"] = o.UnsignedTransaction
	toSerialize["payloads"] = o.Payloads
	return toSerialize, nil
}

func (o *RosettaConstructionPayloadResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"unsigned_transaction",
		"payloads",
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

	varRosettaConstructionPayloadResponse := _RosettaConstructionPayloadResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRosettaConstructionPayloadResponse)

	if err != nil {
		return err
	}

	*o = RosettaConstructionPayloadResponse(varRosettaConstructionPayloadResponse)

	return err
}

type NullableRosettaConstructionPayloadResponse struct {
	value *RosettaConstructionPayloadResponse
	isSet bool
}

func (v NullableRosettaConstructionPayloadResponse) Get() *RosettaConstructionPayloadResponse {
	return v.value
}

func (v *NullableRosettaConstructionPayloadResponse) Set(val *RosettaConstructionPayloadResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRosettaConstructionPayloadResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRosettaConstructionPayloadResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRosettaConstructionPayloadResponse(val *RosettaConstructionPayloadResponse) *NullableRosettaConstructionPayloadResponse {
	return &NullableRosettaConstructionPayloadResponse{value: val, isSet: true}
}

func (v NullableRosettaConstructionPayloadResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRosettaConstructionPayloadResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
