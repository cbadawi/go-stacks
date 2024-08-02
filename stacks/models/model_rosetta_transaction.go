package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the RosettaTransaction type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RosettaTransaction{}

// RosettaTransaction Transactions contain an array of Operations that are attributable to the same TransactionIdentifier.
type RosettaTransaction struct {
	TransactionIdentifier TransactionIdentifier `json:"transaction_identifier"`
	// List of operations
	Operations []RosettaOperation          `json:"operations"`
	Metadata   *RosettaTransactionMetadata `json:"metadata,omitempty"`
}

type _RosettaTransaction RosettaTransaction

// NewRosettaTransaction instantiates a new RosettaTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRosettaTransaction(transactionIdentifier TransactionIdentifier, operations []RosettaOperation) *RosettaTransaction {
	this := RosettaTransaction{}
	this.TransactionIdentifier = transactionIdentifier
	this.Operations = operations
	return &this
}

// NewRosettaTransactionWithDefaults instantiates a new RosettaTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRosettaTransactionWithDefaults() *RosettaTransaction {
	this := RosettaTransaction{}
	return &this
}

// GetTransactionIdentifier returns the TransactionIdentifier field value
func (o *RosettaTransaction) GetTransactionIdentifier() TransactionIdentifier {
	if o == nil {
		var ret TransactionIdentifier
		return ret
	}

	return o.TransactionIdentifier
}

// GetTransactionIdentifierOk returns a tuple with the TransactionIdentifier field value
// and a boolean to check if the value has been set.
func (o *RosettaTransaction) GetTransactionIdentifierOk() (*TransactionIdentifier, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionIdentifier, true
}

// SetTransactionIdentifier sets field value
func (o *RosettaTransaction) SetTransactionIdentifier(v TransactionIdentifier) {
	o.TransactionIdentifier = v
}

// GetOperations returns the Operations field value
func (o *RosettaTransaction) GetOperations() []RosettaOperation {
	if o == nil {
		var ret []RosettaOperation
		return ret
	}

	return o.Operations
}

// GetOperationsOk returns a tuple with the Operations field value
// and a boolean to check if the value has been set.
func (o *RosettaTransaction) GetOperationsOk() ([]RosettaOperation, bool) {
	if o == nil {
		return nil, false
	}
	return o.Operations, true
}

// SetOperations sets field value
func (o *RosettaTransaction) SetOperations(v []RosettaOperation) {
	o.Operations = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *RosettaTransaction) GetMetadata() RosettaTransactionMetadata {
	if o == nil || utils.IsNil(o.Metadata) {
		var ret RosettaTransactionMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaTransaction) GetMetadataOk() (*RosettaTransactionMetadata, bool) {
	if o == nil || utils.IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *RosettaTransaction) HasMetadata() bool {
	if o != nil && !utils.IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given RosettaTransactionMetadata and assigns it to the Metadata field.
func (o *RosettaTransaction) SetMetadata(v RosettaTransactionMetadata) {
	o.Metadata = &v
}

func (o RosettaTransaction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RosettaTransaction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["transaction_identifier"] = o.TransactionIdentifier
	toSerialize["operations"] = o.Operations
	if !utils.IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

func (o *RosettaTransaction) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"transaction_identifier",
		"operations",
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

	varRosettaTransaction := _RosettaTransaction{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRosettaTransaction)

	if err != nil {
		return err
	}

	*o = RosettaTransaction(varRosettaTransaction)

	return err
}

type NullableRosettaTransaction struct {
	value *RosettaTransaction
	isSet bool
}

func (v NullableRosettaTransaction) Get() *RosettaTransaction {
	return v.value
}

func (v *NullableRosettaTransaction) Set(val *RosettaTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableRosettaTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableRosettaTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRosettaTransaction(val *RosettaTransaction) *NullableRosettaTransaction {
	return &NullableRosettaTransaction{value: val, isSet: true}
}

func (v NullableRosettaTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRosettaTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
