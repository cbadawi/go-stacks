package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the RosettaBlockTransactionRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RosettaBlockTransactionRequest{}

// RosettaBlockTransactionRequest A BlockTransactionRequest is used to fetch a Transaction included in a block that is not returned in a BlockResponse.
type RosettaBlockTransactionRequest struct {
	NetworkIdentifier     NetworkIdentifier             `json:"network_identifier"`
	BlockIdentifier       RosettaPartialBlockIdentifier `json:"block_identifier"`
	TransactionIdentifier TransactionIdentifier         `json:"transaction_identifier"`
}

type _RosettaBlockTransactionRequest RosettaBlockTransactionRequest

// NewRosettaBlockTransactionRequest instantiates a new RosettaBlockTransactionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRosettaBlockTransactionRequest(networkIdentifier NetworkIdentifier, blockIdentifier RosettaPartialBlockIdentifier, transactionIdentifier TransactionIdentifier) *RosettaBlockTransactionRequest {
	this := RosettaBlockTransactionRequest{}
	this.NetworkIdentifier = networkIdentifier
	this.BlockIdentifier = blockIdentifier
	this.TransactionIdentifier = transactionIdentifier
	return &this
}

// NewRosettaBlockTransactionRequestWithDefaults instantiates a new RosettaBlockTransactionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRosettaBlockTransactionRequestWithDefaults() *RosettaBlockTransactionRequest {
	this := RosettaBlockTransactionRequest{}
	return &this
}

// GetNetworkIdentifier returns the NetworkIdentifier field value
func (o *RosettaBlockTransactionRequest) GetNetworkIdentifier() NetworkIdentifier {
	if o == nil {
		var ret NetworkIdentifier
		return ret
	}

	return o.NetworkIdentifier
}

// GetNetworkIdentifierOk returns a tuple with the NetworkIdentifier field value
// and a boolean to check if the value has been set.
func (o *RosettaBlockTransactionRequest) GetNetworkIdentifierOk() (*NetworkIdentifier, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkIdentifier, true
}

// SetNetworkIdentifier sets field value
func (o *RosettaBlockTransactionRequest) SetNetworkIdentifier(v NetworkIdentifier) {
	o.NetworkIdentifier = v
}

// GetBlockIdentifier returns the BlockIdentifier field value
func (o *RosettaBlockTransactionRequest) GetBlockIdentifier() RosettaPartialBlockIdentifier {
	if o == nil {
		var ret RosettaPartialBlockIdentifier
		return ret
	}

	return o.BlockIdentifier
}

// GetBlockIdentifierOk returns a tuple with the BlockIdentifier field value
// and a boolean to check if the value has been set.
func (o *RosettaBlockTransactionRequest) GetBlockIdentifierOk() (*RosettaPartialBlockIdentifier, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockIdentifier, true
}

// SetBlockIdentifier sets field value
func (o *RosettaBlockTransactionRequest) SetBlockIdentifier(v RosettaPartialBlockIdentifier) {
	o.BlockIdentifier = v
}

// GetTransactionIdentifier returns the TransactionIdentifier field value
func (o *RosettaBlockTransactionRequest) GetTransactionIdentifier() TransactionIdentifier {
	if o == nil {
		var ret TransactionIdentifier
		return ret
	}

	return o.TransactionIdentifier
}

// GetTransactionIdentifierOk returns a tuple with the TransactionIdentifier field value
// and a boolean to check if the value has been set.
func (o *RosettaBlockTransactionRequest) GetTransactionIdentifierOk() (*TransactionIdentifier, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionIdentifier, true
}

// SetTransactionIdentifier sets field value
func (o *RosettaBlockTransactionRequest) SetTransactionIdentifier(v TransactionIdentifier) {
	o.TransactionIdentifier = v
}

func (o RosettaBlockTransactionRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RosettaBlockTransactionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["network_identifier"] = o.NetworkIdentifier
	toSerialize["block_identifier"] = o.BlockIdentifier
	toSerialize["transaction_identifier"] = o.TransactionIdentifier
	return toSerialize, nil
}

func (o *RosettaBlockTransactionRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"network_identifier",
		"block_identifier",
		"transaction_identifier",
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

	varRosettaBlockTransactionRequest := _RosettaBlockTransactionRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRosettaBlockTransactionRequest)

	if err != nil {
		return err
	}

	*o = RosettaBlockTransactionRequest(varRosettaBlockTransactionRequest)

	return err
}

type NullableRosettaBlockTransactionRequest struct {
	value *RosettaBlockTransactionRequest
	isSet bool
}

func (v NullableRosettaBlockTransactionRequest) Get() *RosettaBlockTransactionRequest {
	return v.value
}

func (v *NullableRosettaBlockTransactionRequest) Set(val *RosettaBlockTransactionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRosettaBlockTransactionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRosettaBlockTransactionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRosettaBlockTransactionRequest(val *RosettaBlockTransactionRequest) *NullableRosettaBlockTransactionRequest {
	return &NullableRosettaBlockTransactionRequest{value: val, isSet: true}
}

func (v NullableRosettaBlockTransactionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRosettaBlockTransactionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
