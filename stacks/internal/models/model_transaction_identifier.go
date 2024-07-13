package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the TransactionIdentifier type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &TransactionIdentifier{}

// TransactionIdentifier The transaction_identifier uniquely identifies a transaction in a particular network and block or in the mempool.
type TransactionIdentifier struct {
	// Any transactions that are attributable only to a block (ex: a block event) should use the hash of the block as the identifier.
	Hash string `json:"hash"`
}

type _TransactionIdentifier TransactionIdentifier

// NewTransactionIdentifier instantiates a new TransactionIdentifier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionIdentifier(hash string) *TransactionIdentifier {
	this := TransactionIdentifier{}
	this.Hash = hash
	return &this
}

// NewTransactionIdentifierWithDefaults instantiates a new TransactionIdentifier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionIdentifierWithDefaults() *TransactionIdentifier {
	this := TransactionIdentifier{}
	return &this
}

// GetHash returns the Hash field value
func (o *TransactionIdentifier) GetHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hash
}

// GetHashOk returns a tuple with the Hash field value
// and a boolean to check if the value has been set.
func (o *TransactionIdentifier) GetHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hash, true
}

// SetHash sets field value
func (o *TransactionIdentifier) SetHash(v string) {
	o.Hash = v
}

func (o TransactionIdentifier) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionIdentifier) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["hash"] = o.Hash
	return toSerialize, nil
}

func (o *TransactionIdentifier) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"hash",
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

	varTransactionIdentifier := _TransactionIdentifier{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransactionIdentifier)

	if err != nil {
		return err
	}

	*o = TransactionIdentifier(varTransactionIdentifier)

	return err
}

type NullableTransactionIdentifier struct {
	value *TransactionIdentifier
	isSet bool
}

func (v NullableTransactionIdentifier) Get() *TransactionIdentifier {
	return v.value
}

func (v *NullableTransactionIdentifier) Set(val *TransactionIdentifier) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionIdentifier) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionIdentifier(val *TransactionIdentifier) *NullableTransactionIdentifier {
	return &NullableTransactionIdentifier{value: val, isSet: true}
}

func (v NullableTransactionIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
