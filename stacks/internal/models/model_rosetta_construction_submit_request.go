package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the RosettaConstructionSubmitRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RosettaConstructionSubmitRequest{}

// RosettaConstructionSubmitRequest Submit the transaction in blockchain
type RosettaConstructionSubmitRequest struct {
	NetworkIdentifier NetworkIdentifier `json:"network_identifier"`
	// Signed transaction
	SignedTransaction string `json:"signed_transaction"`
}

type _RosettaConstructionSubmitRequest RosettaConstructionSubmitRequest

// NewRosettaConstructionSubmitRequest instantiates a new RosettaConstructionSubmitRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRosettaConstructionSubmitRequest(networkIdentifier NetworkIdentifier, signedTransaction string) *RosettaConstructionSubmitRequest {
	this := RosettaConstructionSubmitRequest{}
	this.NetworkIdentifier = networkIdentifier
	this.SignedTransaction = signedTransaction
	return &this
}

// NewRosettaConstructionSubmitRequestWithDefaults instantiates a new RosettaConstructionSubmitRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRosettaConstructionSubmitRequestWithDefaults() *RosettaConstructionSubmitRequest {
	this := RosettaConstructionSubmitRequest{}
	return &this
}

// GetNetworkIdentifier returns the NetworkIdentifier field value
func (o *RosettaConstructionSubmitRequest) GetNetworkIdentifier() NetworkIdentifier {
	if o == nil {
		var ret NetworkIdentifier
		return ret
	}

	return o.NetworkIdentifier
}

// GetNetworkIdentifierOk returns a tuple with the NetworkIdentifier field value
// and a boolean to check if the value has been set.
func (o *RosettaConstructionSubmitRequest) GetNetworkIdentifierOk() (*NetworkIdentifier, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkIdentifier, true
}

// SetNetworkIdentifier sets field value
func (o *RosettaConstructionSubmitRequest) SetNetworkIdentifier(v NetworkIdentifier) {
	o.NetworkIdentifier = v
}

// GetSignedTransaction returns the SignedTransaction field value
func (o *RosettaConstructionSubmitRequest) GetSignedTransaction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SignedTransaction
}

// GetSignedTransactionOk returns a tuple with the SignedTransaction field value
// and a boolean to check if the value has been set.
func (o *RosettaConstructionSubmitRequest) GetSignedTransactionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignedTransaction, true
}

// SetSignedTransaction sets field value
func (o *RosettaConstructionSubmitRequest) SetSignedTransaction(v string) {
	o.SignedTransaction = v
}

func (o RosettaConstructionSubmitRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RosettaConstructionSubmitRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["network_identifier"] = o.NetworkIdentifier
	toSerialize["signed_transaction"] = o.SignedTransaction
	return toSerialize, nil
}

func (o *RosettaConstructionSubmitRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"network_identifier",
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

	varRosettaConstructionSubmitRequest := _RosettaConstructionSubmitRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRosettaConstructionSubmitRequest)

	if err != nil {
		return err
	}

	*o = RosettaConstructionSubmitRequest(varRosettaConstructionSubmitRequest)

	return err
}

type NullableRosettaConstructionSubmitRequest struct {
	value *RosettaConstructionSubmitRequest
	isSet bool
}

func (v NullableRosettaConstructionSubmitRequest) Get() *RosettaConstructionSubmitRequest {
	return v.value
}

func (v *NullableRosettaConstructionSubmitRequest) Set(val *RosettaConstructionSubmitRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRosettaConstructionSubmitRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRosettaConstructionSubmitRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRosettaConstructionSubmitRequest(val *RosettaConstructionSubmitRequest) *NullableRosettaConstructionSubmitRequest {
	return &NullableRosettaConstructionSubmitRequest{value: val, isSet: true}
}

func (v NullableRosettaConstructionSubmitRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRosettaConstructionSubmitRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
