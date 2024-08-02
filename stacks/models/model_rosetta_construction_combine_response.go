package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the RosettaConstructionCombineResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RosettaConstructionCombineResponse{}

// RosettaConstructionCombineResponse RosettaConstructionCombineResponse is returned by /construction/combine. The network payload will be sent directly to the construction/submit endpoint.
type RosettaConstructionCombineResponse struct {
	// Signed transaction bytes in hex
	SignedTransaction string `json:"signed_transaction"`
}

type _RosettaConstructionCombineResponse RosettaConstructionCombineResponse

// NewRosettaConstructionCombineResponse instantiates a new RosettaConstructionCombineResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRosettaConstructionCombineResponse(signedTransaction string) *RosettaConstructionCombineResponse {
	this := RosettaConstructionCombineResponse{}
	this.SignedTransaction = signedTransaction
	return &this
}

// NewRosettaConstructionCombineResponseWithDefaults instantiates a new RosettaConstructionCombineResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRosettaConstructionCombineResponseWithDefaults() *RosettaConstructionCombineResponse {
	this := RosettaConstructionCombineResponse{}
	return &this
}

// GetSignedTransaction returns the SignedTransaction field value
func (o *RosettaConstructionCombineResponse) GetSignedTransaction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SignedTransaction
}

// GetSignedTransactionOk returns a tuple with the SignedTransaction field value
// and a boolean to check if the value has been set.
func (o *RosettaConstructionCombineResponse) GetSignedTransactionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignedTransaction, true
}

// SetSignedTransaction sets field value
func (o *RosettaConstructionCombineResponse) SetSignedTransaction(v string) {
	o.SignedTransaction = v
}

func (o RosettaConstructionCombineResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RosettaConstructionCombineResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["signed_transaction"] = o.SignedTransaction
	return toSerialize, nil
}

func (o *RosettaConstructionCombineResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"signed_transaction",
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

	varRosettaConstructionCombineResponse := _RosettaConstructionCombineResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRosettaConstructionCombineResponse)

	if err != nil {
		return err
	}

	*o = RosettaConstructionCombineResponse(varRosettaConstructionCombineResponse)

	return err
}

type NullableRosettaConstructionCombineResponse struct {
	value *RosettaConstructionCombineResponse
	isSet bool
}

func (v NullableRosettaConstructionCombineResponse) Get() *RosettaConstructionCombineResponse {
	return v.value
}

func (v *NullableRosettaConstructionCombineResponse) Set(val *RosettaConstructionCombineResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRosettaConstructionCombineResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRosettaConstructionCombineResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRosettaConstructionCombineResponse(val *RosettaConstructionCombineResponse) *NullableRosettaConstructionCombineResponse {
	return &NullableRosettaConstructionCombineResponse{value: val, isSet: true}
}

func (v NullableRosettaConstructionCombineResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRosettaConstructionCombineResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
