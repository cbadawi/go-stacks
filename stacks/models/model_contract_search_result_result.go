package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the ContractSearchResultResult type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ContractSearchResultResult{}

// ContractSearchResultResult This object carries the search result
type ContractSearchResultResult struct {
	// The id used to search this query.
	EntityId   string                                       `json:"entity_id"`
	EntityType string                                       `json:"entity_type"`
	TxData     *ContractSearchResultResultTxData            `json:"tx_data,omitempty"`
	Metadata   *AddressTransactionsListResponseResultsInner `json:"metadata,omitempty"`
}

type _ContractSearchResultResult ContractSearchResultResult

// NewContractSearchResultResult instantiates a new ContractSearchResultResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContractSearchResultResult(entityId string, entityType string) *ContractSearchResultResult {
	this := ContractSearchResultResult{}
	this.EntityId = entityId
	this.EntityType = entityType
	return &this
}

// NewContractSearchResultResultWithDefaults instantiates a new ContractSearchResultResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractSearchResultResultWithDefaults() *ContractSearchResultResult {
	this := ContractSearchResultResult{}
	return &this
}

// GetEntityId returns the EntityId field value
func (o *ContractSearchResultResult) GetEntityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value
// and a boolean to check if the value has been set.
func (o *ContractSearchResultResult) GetEntityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityId, true
}

// SetEntityId sets field value
func (o *ContractSearchResultResult) SetEntityId(v string) {
	o.EntityId = v
}

// GetEntityType returns the EntityType field value
func (o *ContractSearchResultResult) GetEntityType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value
// and a boolean to check if the value has been set.
func (o *ContractSearchResultResult) GetEntityTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityType, true
}

// SetEntityType sets field value
func (o *ContractSearchResultResult) SetEntityType(v string) {
	o.EntityType = v
}

// GetTxData returns the TxData field value if set, zero value otherwise.
func (o *ContractSearchResultResult) GetTxData() ContractSearchResultResultTxData {
	if o == nil || utils.IsNil(o.TxData) {
		var ret ContractSearchResultResultTxData
		return ret
	}
	return *o.TxData
}

// GetTxDataOk returns a tuple with the TxData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractSearchResultResult) GetTxDataOk() (*ContractSearchResultResultTxData, bool) {
	if o == nil || utils.IsNil(o.TxData) {
		return nil, false
	}
	return o.TxData, true
}

// HasTxData returns a boolean if a field has been set.
func (o *ContractSearchResultResult) HasTxData() bool {
	if o != nil && !utils.IsNil(o.TxData) {
		return true
	}

	return false
}

// SetTxData gets a reference to the given ContractSearchResultResultTxData and assigns it to the TxData field.
func (o *ContractSearchResultResult) SetTxData(v ContractSearchResultResultTxData) {
	o.TxData = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ContractSearchResultResult) GetMetadata() AddressTransactionsListResponseResultsInner {
	if o == nil || utils.IsNil(o.Metadata) {
		var ret AddressTransactionsListResponseResultsInner
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractSearchResultResult) GetMetadataOk() (*AddressTransactionsListResponseResultsInner, bool) {
	if o == nil || utils.IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ContractSearchResultResult) HasMetadata() bool {
	if o != nil && !utils.IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given AddressTransactionsListResponseResultsInner and assigns it to the Metadata field.
func (o *ContractSearchResultResult) SetMetadata(v AddressTransactionsListResponseResultsInner) {
	o.Metadata = &v
}

func (o ContractSearchResultResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContractSearchResultResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["entity_id"] = o.EntityId
	toSerialize["entity_type"] = o.EntityType
	if !utils.IsNil(o.TxData) {
		toSerialize["tx_data"] = o.TxData
	}
	if !utils.IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

func (o *ContractSearchResultResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"entity_id",
		"entity_type",
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

	varContractSearchResultResult := _ContractSearchResultResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varContractSearchResultResult)

	if err != nil {
		return err
	}

	*o = ContractSearchResultResult(varContractSearchResultResult)

	return err
}

type NullableContractSearchResultResult struct {
	value *ContractSearchResultResult
	isSet bool
}

func (v NullableContractSearchResultResult) Get() *ContractSearchResultResult {
	return v.value
}

func (v *NullableContractSearchResultResult) Set(val *ContractSearchResultResult) {
	v.value = val
	v.isSet = true
}

func (v NullableContractSearchResultResult) IsSet() bool {
	return v.isSet
}

func (v *NullableContractSearchResultResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractSearchResultResult(val *ContractSearchResultResult) *NullableContractSearchResultResult {
	return &NullableContractSearchResultResult{value: val, isSet: true}
}

func (v NullableContractSearchResultResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractSearchResultResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
