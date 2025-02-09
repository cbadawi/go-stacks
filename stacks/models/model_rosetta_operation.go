package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the RosettaOperation type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RosettaOperation{}

// RosettaOperation Operations contain all balance-changing information within a transaction. They are always one-sided (only affect 1 AccountIdentifier) and can succeed or fail independently from a Transaction.
type RosettaOperation struct {
	OperationIdentifier RosettaOperationIdentifier `json:"operation_identifier"`
	// Restrict referenced related_operations to identifier indexes < the current operation_identifier.index. This ensures there exists a clear DAG-structure of relations. Since operations are one-sided, one could imagine relating operations in a single transfer or linking operations in a call tree.
	RelatedOperations []RosettaRelatedOperation `json:"related_operations,omitempty"`
	// The network-specific type of the operation. Ensure that any type that can be returned here is also specified in the NetworkStatus. This can be very useful to downstream consumers that parse all block data.
	Type string `json:"type"`
	// The network-specific status of the operation. Status is not defined on the transaction object because blockchains with smart contracts may have transactions that partially apply. Blockchains with atomic transactions (all operations succeed or all operations fail) will have the same status for each operation.
	Status     *string            `json:"status,omitempty"`
	Account    *RosettaAccount    `json:"account,omitempty"`
	Amount     *RosettaAmount     `json:"amount,omitempty"`
	CoinChange *RosettaCoinChange `json:"coin_change,omitempty"`
	// Operations Meta Data
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

type _RosettaOperation RosettaOperation

// NewRosettaOperation instantiates a new RosettaOperation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRosettaOperation(operationIdentifier RosettaOperationIdentifier, type_ string) *RosettaOperation {
	this := RosettaOperation{}
	this.OperationIdentifier = operationIdentifier
	this.Type = type_
	return &this
}

// NewRosettaOperationWithDefaults instantiates a new RosettaOperation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRosettaOperationWithDefaults() *RosettaOperation {
	this := RosettaOperation{}
	return &this
}

// GetOperationIdentifier returns the OperationIdentifier field value
func (o *RosettaOperation) GetOperationIdentifier() RosettaOperationIdentifier {
	if o == nil {
		var ret RosettaOperationIdentifier
		return ret
	}

	return o.OperationIdentifier
}

// GetOperationIdentifierOk returns a tuple with the OperationIdentifier field value
// and a boolean to check if the value has been set.
func (o *RosettaOperation) GetOperationIdentifierOk() (*RosettaOperationIdentifier, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OperationIdentifier, true
}

// SetOperationIdentifier sets field value
func (o *RosettaOperation) SetOperationIdentifier(v RosettaOperationIdentifier) {
	o.OperationIdentifier = v
}

// GetRelatedOperations returns the RelatedOperations field value if set, zero value otherwise.
func (o *RosettaOperation) GetRelatedOperations() []RosettaRelatedOperation {
	if o == nil || utils.IsNil(o.RelatedOperations) {
		var ret []RosettaRelatedOperation
		return ret
	}
	return o.RelatedOperations
}

// GetRelatedOperationsOk returns a tuple with the RelatedOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaOperation) GetRelatedOperationsOk() ([]RosettaRelatedOperation, bool) {
	if o == nil || utils.IsNil(o.RelatedOperations) {
		return nil, false
	}
	return o.RelatedOperations, true
}

// HasRelatedOperations returns a boolean if a field has been set.
func (o *RosettaOperation) HasRelatedOperations() bool {
	if o != nil && !utils.IsNil(o.RelatedOperations) {
		return true
	}

	return false
}

// SetRelatedOperations gets a reference to the given []RosettaRelatedOperation and assigns it to the RelatedOperations field.
func (o *RosettaOperation) SetRelatedOperations(v []RosettaRelatedOperation) {
	o.RelatedOperations = v
}

// GetType returns the Type field value
func (o *RosettaOperation) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RosettaOperation) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RosettaOperation) SetType(v string) {
	o.Type = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *RosettaOperation) GetStatus() string {
	if o == nil || utils.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaOperation) GetStatusOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *RosettaOperation) HasStatus() bool {
	if o != nil && !utils.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *RosettaOperation) SetStatus(v string) {
	o.Status = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *RosettaOperation) GetAccount() RosettaAccount {
	if o == nil || utils.IsNil(o.Account) {
		var ret RosettaAccount
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaOperation) GetAccountOk() (*RosettaAccount, bool) {
	if o == nil || utils.IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *RosettaOperation) HasAccount() bool {
	if o != nil && !utils.IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given RosettaAccount and assigns it to the Account field.
func (o *RosettaOperation) SetAccount(v RosettaAccount) {
	o.Account = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *RosettaOperation) GetAmount() RosettaAmount {
	if o == nil || utils.IsNil(o.Amount) {
		var ret RosettaAmount
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaOperation) GetAmountOk() (*RosettaAmount, bool) {
	if o == nil || utils.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *RosettaOperation) HasAmount() bool {
	if o != nil && !utils.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given RosettaAmount and assigns it to the Amount field.
func (o *RosettaOperation) SetAmount(v RosettaAmount) {
	o.Amount = &v
}

// GetCoinChange returns the CoinChange field value if set, zero value otherwise.
func (o *RosettaOperation) GetCoinChange() RosettaCoinChange {
	if o == nil || utils.IsNil(o.CoinChange) {
		var ret RosettaCoinChange
		return ret
	}
	return *o.CoinChange
}

// GetCoinChangeOk returns a tuple with the CoinChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaOperation) GetCoinChangeOk() (*RosettaCoinChange, bool) {
	if o == nil || utils.IsNil(o.CoinChange) {
		return nil, false
	}
	return o.CoinChange, true
}

// HasCoinChange returns a boolean if a field has been set.
func (o *RosettaOperation) HasCoinChange() bool {
	if o != nil && !utils.IsNil(o.CoinChange) {
		return true
	}

	return false
}

// SetCoinChange gets a reference to the given RosettaCoinChange and assigns it to the CoinChange field.
func (o *RosettaOperation) SetCoinChange(v RosettaCoinChange) {
	o.CoinChange = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *RosettaOperation) GetMetadata() map[string]interface{} {
	if o == nil || utils.IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaOperation) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || utils.IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *RosettaOperation) HasMetadata() bool {
	if o != nil && !utils.IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *RosettaOperation) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o RosettaOperation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RosettaOperation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["operation_identifier"] = o.OperationIdentifier
	if !utils.IsNil(o.RelatedOperations) {
		toSerialize["related_operations"] = o.RelatedOperations
	}
	toSerialize["type"] = o.Type
	if !utils.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !utils.IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	if !utils.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !utils.IsNil(o.CoinChange) {
		toSerialize["coin_change"] = o.CoinChange
	}
	if !utils.IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

func (o *RosettaOperation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"operation_identifier",
		"type",
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

	varRosettaOperation := _RosettaOperation{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRosettaOperation)

	if err != nil {
		return err
	}

	*o = RosettaOperation(varRosettaOperation)

	return err
}

type NullableRosettaOperation struct {
	value *RosettaOperation
	isSet bool
}

func (v NullableRosettaOperation) Get() *RosettaOperation {
	return v.value
}

func (v *NullableRosettaOperation) Set(val *RosettaOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableRosettaOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableRosettaOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRosettaOperation(val *RosettaOperation) *NullableRosettaOperation {
	return &NullableRosettaOperation{value: val, isSet: true}
}

func (v NullableRosettaOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRosettaOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
