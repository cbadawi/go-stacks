package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the RosettaConstructionHashRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RosettaConstructionHashRequest{}

// RosettaConstructionHashRequest TransactionHash returns the network-specific transaction hash for a signed transaction.
type RosettaConstructionHashRequest struct {
	NetworkIdentifier NetworkIdentifier `json:"network_identifier"`
	// Signed transaction
	SignedTransaction string `json:"signed_transaction"`
}

type _RosettaConstructionHashRequest RosettaConstructionHashRequest

// NewRosettaConstructionHashRequest instantiates a new RosettaConstructionHashRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRosettaConstructionHashRequest(networkIdentifier NetworkIdentifier, signedTransaction string) *RosettaConstructionHashRequest {
	this := RosettaConstructionHashRequest{}
	this.NetworkIdentifier = networkIdentifier
	this.SignedTransaction = signedTransaction
	return &this
}

// NewRosettaConstructionHashRequestWithDefaults instantiates a new RosettaConstructionHashRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRosettaConstructionHashRequestWithDefaults() *RosettaConstructionHashRequest {
	this := RosettaConstructionHashRequest{}
	return &this
}

// GetNetworkIdentifier returns the NetworkIdentifier field value
func (o *RosettaConstructionHashRequest) GetNetworkIdentifier() NetworkIdentifier {
	if o == nil {
		var ret NetworkIdentifier
		return ret
	}

	return o.NetworkIdentifier
}

// GetNetworkIdentifierOk returns a tuple with the NetworkIdentifier field value
// and a boolean to check if the value has been set.
func (o *RosettaConstructionHashRequest) GetNetworkIdentifierOk() (*NetworkIdentifier, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkIdentifier, true
}

// SetNetworkIdentifier sets field value
func (o *RosettaConstructionHashRequest) SetNetworkIdentifier(v NetworkIdentifier) {
	o.NetworkIdentifier = v
}

// GetSignedTransaction returns the SignedTransaction field value
func (o *RosettaConstructionHashRequest) GetSignedTransaction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SignedTransaction
}

// GetSignedTransactionOk returns a tuple with the SignedTransaction field value
// and a boolean to check if the value has been set.
func (o *RosettaConstructionHashRequest) GetSignedTransactionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignedTransaction, true
}

// SetSignedTransaction sets field value
func (o *RosettaConstructionHashRequest) SetSignedTransaction(v string) {
	o.SignedTransaction = v
}

func (o RosettaConstructionHashRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RosettaConstructionHashRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["network_identifier"] = o.NetworkIdentifier
	toSerialize["signed_transaction"] = o.SignedTransaction
	return toSerialize, nil
}

func (o *RosettaConstructionHashRequest) UnmarshalJSON(data []byte) (err error) {
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

	varRosettaConstructionHashRequest := _RosettaConstructionHashRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRosettaConstructionHashRequest)

	if err != nil {
		return err
	}

	*o = RosettaConstructionHashRequest(varRosettaConstructionHashRequest)

	return err
}

type NullableRosettaConstructionHashRequest struct {
	value *RosettaConstructionHashRequest
	isSet bool
}

func (v NullableRosettaConstructionHashRequest) Get() *RosettaConstructionHashRequest {
	return v.value
}

func (v *NullableRosettaConstructionHashRequest) Set(val *RosettaConstructionHashRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRosettaConstructionHashRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRosettaConstructionHashRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRosettaConstructionHashRequest(val *RosettaConstructionHashRequest) *NullableRosettaConstructionHashRequest {
	return &NullableRosettaConstructionHashRequest{value: val, isSet: true}
}

func (v NullableRosettaConstructionHashRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRosettaConstructionHashRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
