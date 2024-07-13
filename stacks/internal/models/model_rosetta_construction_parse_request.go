package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the RosettaConstructionParseRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RosettaConstructionParseRequest{}

// RosettaConstructionParseRequest Parse is called on both unsigned and signed transactions to understand the intent of the formulated transaction. This is run as a sanity check before signing (after /construction/payloads) and before broadcast (after /construction/combine).
type RosettaConstructionParseRequest struct {
	NetworkIdentifier NetworkIdentifier `json:"network_identifier"`
	// Signed is a boolean indicating whether the transaction is signed.
	Signed bool `json:"signed"`
	// This must be either the unsigned transaction blob returned by /construction/payloads or the signed transaction blob returned by /construction/combine.
	Transaction string `json:"transaction"`
}

type _RosettaConstructionParseRequest RosettaConstructionParseRequest

// NewRosettaConstructionParseRequest instantiates a new RosettaConstructionParseRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRosettaConstructionParseRequest(networkIdentifier NetworkIdentifier, signed bool, transaction string) *RosettaConstructionParseRequest {
	this := RosettaConstructionParseRequest{}
	this.NetworkIdentifier = networkIdentifier
	this.Signed = signed
	this.Transaction = transaction
	return &this
}

// NewRosettaConstructionParseRequestWithDefaults instantiates a new RosettaConstructionParseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRosettaConstructionParseRequestWithDefaults() *RosettaConstructionParseRequest {
	this := RosettaConstructionParseRequest{}
	return &this
}

// GetNetworkIdentifier returns the NetworkIdentifier field value
func (o *RosettaConstructionParseRequest) GetNetworkIdentifier() NetworkIdentifier {
	if o == nil {
		var ret NetworkIdentifier
		return ret
	}

	return o.NetworkIdentifier
}

// GetNetworkIdentifierOk returns a tuple with the NetworkIdentifier field value
// and a boolean to check if the value has been set.
func (o *RosettaConstructionParseRequest) GetNetworkIdentifierOk() (*NetworkIdentifier, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkIdentifier, true
}

// SetNetworkIdentifier sets field value
func (o *RosettaConstructionParseRequest) SetNetworkIdentifier(v NetworkIdentifier) {
	o.NetworkIdentifier = v
}

// GetSigned returns the Signed field value
func (o *RosettaConstructionParseRequest) GetSigned() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Signed
}

// GetSignedOk returns a tuple with the Signed field value
// and a boolean to check if the value has been set.
func (o *RosettaConstructionParseRequest) GetSignedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Signed, true
}

// SetSigned sets field value
func (o *RosettaConstructionParseRequest) SetSigned(v bool) {
	o.Signed = v
}

// GetTransaction returns the Transaction field value
func (o *RosettaConstructionParseRequest) GetTransaction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Transaction
}

// GetTransactionOk returns a tuple with the Transaction field value
// and a boolean to check if the value has been set.
func (o *RosettaConstructionParseRequest) GetTransactionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Transaction, true
}

// SetTransaction sets field value
func (o *RosettaConstructionParseRequest) SetTransaction(v string) {
	o.Transaction = v
}

func (o RosettaConstructionParseRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RosettaConstructionParseRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["network_identifier"] = o.NetworkIdentifier
	toSerialize["signed"] = o.Signed
	toSerialize["transaction"] = o.Transaction
	return toSerialize, nil
}

func (o *RosettaConstructionParseRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"network_identifier",
		"signed",
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

	varRosettaConstructionParseRequest := _RosettaConstructionParseRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRosettaConstructionParseRequest)

	if err != nil {
		return err
	}

	*o = RosettaConstructionParseRequest(varRosettaConstructionParseRequest)

	return err
}

type NullableRosettaConstructionParseRequest struct {
	value *RosettaConstructionParseRequest
	isSet bool
}

func (v NullableRosettaConstructionParseRequest) Get() *RosettaConstructionParseRequest {
	return v.value
}

func (v *NullableRosettaConstructionParseRequest) Set(val *RosettaConstructionParseRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRosettaConstructionParseRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRosettaConstructionParseRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRosettaConstructionParseRequest(val *RosettaConstructionParseRequest) *NullableRosettaConstructionParseRequest {
	return &NullableRosettaConstructionParseRequest{value: val, isSet: true}
}

func (v NullableRosettaConstructionParseRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRosettaConstructionParseRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
