package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the RosettaMempoolTransactionRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RosettaMempoolTransactionRequest{}

// RosettaMempoolTransactionRequest A MempoolTransactionRequest is utilized to retrieve a transaction from the mempool.
type RosettaMempoolTransactionRequest struct {
	NetworkIdentifier     NetworkIdentifier     `json:"network_identifier"`
	TransactionIdentifier TransactionIdentifier `json:"transaction_identifier"`
}

type _RosettaMempoolTransactionRequest RosettaMempoolTransactionRequest

// NewRosettaMempoolTransactionRequest instantiates a new RosettaMempoolTransactionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRosettaMempoolTransactionRequest(networkIdentifier NetworkIdentifier, transactionIdentifier TransactionIdentifier) *RosettaMempoolTransactionRequest {
	this := RosettaMempoolTransactionRequest{}
	this.NetworkIdentifier = networkIdentifier
	this.TransactionIdentifier = transactionIdentifier
	return &this
}

// NewRosettaMempoolTransactionRequestWithDefaults instantiates a new RosettaMempoolTransactionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRosettaMempoolTransactionRequestWithDefaults() *RosettaMempoolTransactionRequest {
	this := RosettaMempoolTransactionRequest{}
	return &this
}

// GetNetworkIdentifier returns the NetworkIdentifier field value
func (o *RosettaMempoolTransactionRequest) GetNetworkIdentifier() NetworkIdentifier {
	if o == nil {
		var ret NetworkIdentifier
		return ret
	}

	return o.NetworkIdentifier
}

// GetNetworkIdentifierOk returns a tuple with the NetworkIdentifier field value
// and a boolean to check if the value has been set.
func (o *RosettaMempoolTransactionRequest) GetNetworkIdentifierOk() (*NetworkIdentifier, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkIdentifier, true
}

// SetNetworkIdentifier sets field value
func (o *RosettaMempoolTransactionRequest) SetNetworkIdentifier(v NetworkIdentifier) {
	o.NetworkIdentifier = v
}

// GetTransactionIdentifier returns the TransactionIdentifier field value
func (o *RosettaMempoolTransactionRequest) GetTransactionIdentifier() TransactionIdentifier {
	if o == nil {
		var ret TransactionIdentifier
		return ret
	}

	return o.TransactionIdentifier
}

// GetTransactionIdentifierOk returns a tuple with the TransactionIdentifier field value
// and a boolean to check if the value has been set.
func (o *RosettaMempoolTransactionRequest) GetTransactionIdentifierOk() (*TransactionIdentifier, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionIdentifier, true
}

// SetTransactionIdentifier sets field value
func (o *RosettaMempoolTransactionRequest) SetTransactionIdentifier(v TransactionIdentifier) {
	o.TransactionIdentifier = v
}

func (o RosettaMempoolTransactionRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RosettaMempoolTransactionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["network_identifier"] = o.NetworkIdentifier
	toSerialize["transaction_identifier"] = o.TransactionIdentifier
	return toSerialize, nil
}

func (o *RosettaMempoolTransactionRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"network_identifier",
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

	varRosettaMempoolTransactionRequest := _RosettaMempoolTransactionRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRosettaMempoolTransactionRequest)

	if err != nil {
		return err
	}

	*o = RosettaMempoolTransactionRequest(varRosettaMempoolTransactionRequest)

	return err
}

type NullableRosettaMempoolTransactionRequest struct {
	value *RosettaMempoolTransactionRequest
	isSet bool
}

func (v NullableRosettaMempoolTransactionRequest) Get() *RosettaMempoolTransactionRequest {
	return v.value
}

func (v *NullableRosettaMempoolTransactionRequest) Set(val *RosettaMempoolTransactionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRosettaMempoolTransactionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRosettaMempoolTransactionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRosettaMempoolTransactionRequest(val *RosettaMempoolTransactionRequest) *NullableRosettaMempoolTransactionRequest {
	return &NullableRosettaMempoolTransactionRequest{value: val, isSet: true}
}

func (v NullableRosettaMempoolTransactionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRosettaMempoolTransactionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
