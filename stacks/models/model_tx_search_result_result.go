package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the TxSearchResultResult type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &TxSearchResultResult{}

// TxSearchResultResult This object carries the search result
type TxSearchResultResult struct {
	// The id used to search this query.
	EntityId   string                     `json:"entity_id"`
	EntityType string                     `json:"entity_type"`
	TxData     TxSearchResultResultTxData `json:"tx_data"`
	Metadata   *Transaction               `json:"metadata,omitempty"`
}

type _TxSearchResultResult TxSearchResultResult

// NewTxSearchResultResult instantiates a new TxSearchResultResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTxSearchResultResult(entityId string, entityType string, txData TxSearchResultResultTxData) *TxSearchResultResult {
	this := TxSearchResultResult{}
	this.EntityId = entityId
	this.EntityType = entityType
	this.TxData = txData
	return &this
}

// NewTxSearchResultResultWithDefaults instantiates a new TxSearchResultResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTxSearchResultResultWithDefaults() *TxSearchResultResult {
	this := TxSearchResultResult{}
	return &this
}

// GetEntityId returns the EntityId field value
func (o *TxSearchResultResult) GetEntityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value
// and a boolean to check if the value has been set.
func (o *TxSearchResultResult) GetEntityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityId, true
}

// SetEntityId sets field value
func (o *TxSearchResultResult) SetEntityId(v string) {
	o.EntityId = v
}

// GetEntityType returns the EntityType field value
func (o *TxSearchResultResult) GetEntityType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value
// and a boolean to check if the value has been set.
func (o *TxSearchResultResult) GetEntityTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityType, true
}

// SetEntityType sets field value
func (o *TxSearchResultResult) SetEntityType(v string) {
	o.EntityType = v
}

// GetTxData returns the TxData field value
func (o *TxSearchResultResult) GetTxData() TxSearchResultResultTxData {
	if o == nil {
		var ret TxSearchResultResultTxData
		return ret
	}

	return o.TxData
}

// GetTxDataOk returns a tuple with the TxData field value
// and a boolean to check if the value has been set.
func (o *TxSearchResultResult) GetTxDataOk() (*TxSearchResultResultTxData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxData, true
}

// SetTxData sets field value
func (o *TxSearchResultResult) SetTxData(v TxSearchResultResultTxData) {
	o.TxData = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *TxSearchResultResult) GetMetadata() Transaction {
	if o == nil || utils.IsNil(o.Metadata) {
		var ret Transaction
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TxSearchResultResult) GetMetadataOk() (*Transaction, bool) {
	if o == nil || utils.IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *TxSearchResultResult) HasMetadata() bool {
	if o != nil && !utils.IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given Transaction and assigns it to the Metadata field.
func (o *TxSearchResultResult) SetMetadata(v Transaction) {
	o.Metadata = &v
}

func (o TxSearchResultResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TxSearchResultResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["entity_id"] = o.EntityId
	toSerialize["entity_type"] = o.EntityType
	toSerialize["tx_data"] = o.TxData
	if !utils.IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

func (o *TxSearchResultResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"entity_id",
		"entity_type",
		"tx_data",
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

	varTxSearchResultResult := _TxSearchResultResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTxSearchResultResult)

	if err != nil {
		return err
	}

	*o = TxSearchResultResult(varTxSearchResultResult)

	return err
}

type NullableTxSearchResultResult struct {
	value *TxSearchResultResult
	isSet bool
}

func (v NullableTxSearchResultResult) Get() *TxSearchResultResult {
	return v.value
}

func (v *NullableTxSearchResultResult) Set(val *TxSearchResultResult) {
	v.value = val
	v.isSet = true
}

func (v NullableTxSearchResultResult) IsSet() bool {
	return v.isSet
}

func (v *NullableTxSearchResultResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTxSearchResultResult(val *TxSearchResultResult) *NullableTxSearchResultResult {
	return &NullableTxSearchResultResult{value: val, isSet: true}
}

func (v NullableTxSearchResultResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTxSearchResultResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
