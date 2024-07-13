package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the RosettaConstructionCombineRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RosettaConstructionCombineRequest{}

// RosettaConstructionCombineRequest RosettaConstructionCombineRequest is the input to the /construction/combine endpoint. It contains the unsigned transaction blob returned by /construction/payloads and all required signatures to create a network transaction.
type RosettaConstructionCombineRequest struct {
	NetworkIdentifier   NetworkIdentifier  `json:"network_identifier"`
	UnsignedTransaction string             `json:"unsigned_transaction"`
	Signatures          []RosettaSignature `json:"signatures"`
}

type _RosettaConstructionCombineRequest RosettaConstructionCombineRequest

// NewRosettaConstructionCombineRequest instantiates a new RosettaConstructionCombineRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRosettaConstructionCombineRequest(networkIdentifier NetworkIdentifier, unsignedTransaction string, signatures []RosettaSignature) *RosettaConstructionCombineRequest {
	this := RosettaConstructionCombineRequest{}
	this.NetworkIdentifier = networkIdentifier
	this.UnsignedTransaction = unsignedTransaction
	this.Signatures = signatures
	return &this
}

// NewRosettaConstructionCombineRequestWithDefaults instantiates a new RosettaConstructionCombineRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRosettaConstructionCombineRequestWithDefaults() *RosettaConstructionCombineRequest {
	this := RosettaConstructionCombineRequest{}
	return &this
}

// GetNetworkIdentifier returns the NetworkIdentifier field value
func (o *RosettaConstructionCombineRequest) GetNetworkIdentifier() NetworkIdentifier {
	if o == nil {
		var ret NetworkIdentifier
		return ret
	}

	return o.NetworkIdentifier
}

// GetNetworkIdentifierOk returns a tuple with the NetworkIdentifier field value
// and a boolean to check if the value has been set.
func (o *RosettaConstructionCombineRequest) GetNetworkIdentifierOk() (*NetworkIdentifier, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkIdentifier, true
}

// SetNetworkIdentifier sets field value
func (o *RosettaConstructionCombineRequest) SetNetworkIdentifier(v NetworkIdentifier) {
	o.NetworkIdentifier = v
}

// GetUnsignedTransaction returns the UnsignedTransaction field value
func (o *RosettaConstructionCombineRequest) GetUnsignedTransaction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnsignedTransaction
}

// GetUnsignedTransactionOk returns a tuple with the UnsignedTransaction field value
// and a boolean to check if the value has been set.
func (o *RosettaConstructionCombineRequest) GetUnsignedTransactionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnsignedTransaction, true
}

// SetUnsignedTransaction sets field value
func (o *RosettaConstructionCombineRequest) SetUnsignedTransaction(v string) {
	o.UnsignedTransaction = v
}

// GetSignatures returns the Signatures field value
func (o *RosettaConstructionCombineRequest) GetSignatures() []RosettaSignature {
	if o == nil {
		var ret []RosettaSignature
		return ret
	}

	return o.Signatures
}

// GetSignaturesOk returns a tuple with the Signatures field value
// and a boolean to check if the value has been set.
func (o *RosettaConstructionCombineRequest) GetSignaturesOk() ([]RosettaSignature, bool) {
	if o == nil {
		return nil, false
	}
	return o.Signatures, true
}

// SetSignatures sets field value
func (o *RosettaConstructionCombineRequest) SetSignatures(v []RosettaSignature) {
	o.Signatures = v
}

func (o RosettaConstructionCombineRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RosettaConstructionCombineRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["network_identifier"] = o.NetworkIdentifier
	toSerialize["unsigned_transaction"] = o.UnsignedTransaction
	toSerialize["signatures"] = o.Signatures
	return toSerialize, nil
}

func (o *RosettaConstructionCombineRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"network_identifier",
		"unsigned_transaction",
		"signatures",
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

	varRosettaConstructionCombineRequest := _RosettaConstructionCombineRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRosettaConstructionCombineRequest)

	if err != nil {
		return err
	}

	*o = RosettaConstructionCombineRequest(varRosettaConstructionCombineRequest)

	return err
}

type NullableRosettaConstructionCombineRequest struct {
	value *RosettaConstructionCombineRequest
	isSet bool
}

func (v NullableRosettaConstructionCombineRequest) Get() *RosettaConstructionCombineRequest {
	return v.value
}

func (v *NullableRosettaConstructionCombineRequest) Set(val *RosettaConstructionCombineRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRosettaConstructionCombineRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRosettaConstructionCombineRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRosettaConstructionCombineRequest(val *RosettaConstructionCombineRequest) *NullableRosettaConstructionCombineRequest {
	return &NullableRosettaConstructionCombineRequest{value: val, isSet: true}
}

func (v NullableRosettaConstructionCombineRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRosettaConstructionCombineRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
