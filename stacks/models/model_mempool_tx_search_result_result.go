package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the MempoolTxSearchResultResult type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &MempoolTxSearchResultResult{}

// MempoolTxSearchResultResult This object carries the search result
type MempoolTxSearchResultResult struct {
	// The id used to search this query.
	EntityId   string                            `json:"entity_id"`
	EntityType string                            `json:"entity_type"`
	TxData     MempoolTxSearchResultResultTxData `json:"tx_data"`
	Metadata   *MempoolTransaction               `json:"metadata,omitempty"`
}

type _MempoolTxSearchResultResult MempoolTxSearchResultResult

// NewMempoolTxSearchResultResult instantiates a new MempoolTxSearchResultResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMempoolTxSearchResultResult(entityId string, entityType string, txData MempoolTxSearchResultResultTxData) *MempoolTxSearchResultResult {
	this := MempoolTxSearchResultResult{}
	this.EntityId = entityId
	this.EntityType = entityType
	this.TxData = txData
	return &this
}

// NewMempoolTxSearchResultResultWithDefaults instantiates a new MempoolTxSearchResultResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMempoolTxSearchResultResultWithDefaults() *MempoolTxSearchResultResult {
	this := MempoolTxSearchResultResult{}
	return &this
}

// GetEntityId returns the EntityId field value
func (o *MempoolTxSearchResultResult) GetEntityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value
// and a boolean to check if the value has been set.
func (o *MempoolTxSearchResultResult) GetEntityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityId, true
}

// SetEntityId sets field value
func (o *MempoolTxSearchResultResult) SetEntityId(v string) {
	o.EntityId = v
}

// GetEntityType returns the EntityType field value
func (o *MempoolTxSearchResultResult) GetEntityType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value
// and a boolean to check if the value has been set.
func (o *MempoolTxSearchResultResult) GetEntityTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityType, true
}

// SetEntityType sets field value
func (o *MempoolTxSearchResultResult) SetEntityType(v string) {
	o.EntityType = v
}

// GetTxData returns the TxData field value
func (o *MempoolTxSearchResultResult) GetTxData() MempoolTxSearchResultResultTxData {
	if o == nil {
		var ret MempoolTxSearchResultResultTxData
		return ret
	}

	return o.TxData
}

// GetTxDataOk returns a tuple with the TxData field value
// and a boolean to check if the value has been set.
func (o *MempoolTxSearchResultResult) GetTxDataOk() (*MempoolTxSearchResultResultTxData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxData, true
}

// SetTxData sets field value
func (o *MempoolTxSearchResultResult) SetTxData(v MempoolTxSearchResultResultTxData) {
	o.TxData = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *MempoolTxSearchResultResult) GetMetadata() MempoolTransaction {
	if o == nil || utils.IsNil(o.Metadata) {
		var ret MempoolTransaction
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MempoolTxSearchResultResult) GetMetadataOk() (*MempoolTransaction, bool) {
	if o == nil || utils.IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *MempoolTxSearchResultResult) HasMetadata() bool {
	if o != nil && !utils.IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given MempoolTransaction and assigns it to the Metadata field.
func (o *MempoolTxSearchResultResult) SetMetadata(v MempoolTransaction) {
	o.Metadata = &v
}

func (o MempoolTxSearchResultResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MempoolTxSearchResultResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["entity_id"] = o.EntityId
	toSerialize["entity_type"] = o.EntityType
	toSerialize["tx_data"] = o.TxData
	if !utils.IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

func (o *MempoolTxSearchResultResult) UnmarshalJSON(data []byte) (err error) {
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

	varMempoolTxSearchResultResult := _MempoolTxSearchResultResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMempoolTxSearchResultResult)

	if err != nil {
		return err
	}

	*o = MempoolTxSearchResultResult(varMempoolTxSearchResultResult)

	return err
}

type NullableMempoolTxSearchResultResult struct {
	value *MempoolTxSearchResultResult
	isSet bool
}

func (v NullableMempoolTxSearchResultResult) Get() *MempoolTxSearchResultResult {
	return v.value
}

func (v *NullableMempoolTxSearchResultResult) Set(val *MempoolTxSearchResultResult) {
	v.value = val
	v.isSet = true
}

func (v NullableMempoolTxSearchResultResult) IsSet() bool {
	return v.isSet
}

func (v *NullableMempoolTxSearchResultResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMempoolTxSearchResultResult(val *MempoolTxSearchResultResult) *NullableMempoolTxSearchResultResult {
	return &NullableMempoolTxSearchResultResult{value: val, isSet: true}
}

func (v NullableMempoolTxSearchResultResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMempoolTxSearchResultResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
