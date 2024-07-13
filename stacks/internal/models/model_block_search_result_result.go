package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the BlockSearchResultResult type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &BlockSearchResultResult{}

// BlockSearchResultResult This object carries the search result
type BlockSearchResultResult struct {
	// The id used to search this query.
	EntityId   string                           `json:"entity_id"`
	EntityType string                           `json:"entity_type"`
	BlockData  BlockSearchResultResultBlockData `json:"block_data"`
	Metadata   *Block                           `json:"metadata,omitempty"`
}

type _BlockSearchResultResult BlockSearchResultResult

// NewBlockSearchResultResult instantiates a new BlockSearchResultResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockSearchResultResult(entityId string, entityType string, blockData BlockSearchResultResultBlockData) *BlockSearchResultResult {
	this := BlockSearchResultResult{}
	this.EntityId = entityId
	this.EntityType = entityType
	this.BlockData = blockData
	return &this
}

// NewBlockSearchResultResultWithDefaults instantiates a new BlockSearchResultResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockSearchResultResultWithDefaults() *BlockSearchResultResult {
	this := BlockSearchResultResult{}
	return &this
}

// GetEntityId returns the EntityId field value
func (o *BlockSearchResultResult) GetEntityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value
// and a boolean to check if the value has been set.
func (o *BlockSearchResultResult) GetEntityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityId, true
}

// SetEntityId sets field value
func (o *BlockSearchResultResult) SetEntityId(v string) {
	o.EntityId = v
}

// GetEntityType returns the EntityType field value
func (o *BlockSearchResultResult) GetEntityType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value
// and a boolean to check if the value has been set.
func (o *BlockSearchResultResult) GetEntityTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityType, true
}

// SetEntityType sets field value
func (o *BlockSearchResultResult) SetEntityType(v string) {
	o.EntityType = v
}

// GetBlockData returns the BlockData field value
func (o *BlockSearchResultResult) GetBlockData() BlockSearchResultResultBlockData {
	if o == nil {
		var ret BlockSearchResultResultBlockData
		return ret
	}

	return o.BlockData
}

// GetBlockDataOk returns a tuple with the BlockData field value
// and a boolean to check if the value has been set.
func (o *BlockSearchResultResult) GetBlockDataOk() (*BlockSearchResultResultBlockData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockData, true
}

// SetBlockData sets field value
func (o *BlockSearchResultResult) SetBlockData(v BlockSearchResultResultBlockData) {
	o.BlockData = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *BlockSearchResultResult) GetMetadata() Block {
	if o == nil || utils.IsNil(o.Metadata) {
		var ret Block
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockSearchResultResult) GetMetadataOk() (*Block, bool) {
	if o == nil || utils.IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *BlockSearchResultResult) HasMetadata() bool {
	if o != nil && !utils.IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given Block and assigns it to the Metadata field.
func (o *BlockSearchResultResult) SetMetadata(v Block) {
	o.Metadata = &v
}

func (o BlockSearchResultResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlockSearchResultResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["entity_id"] = o.EntityId
	toSerialize["entity_type"] = o.EntityType
	toSerialize["block_data"] = o.BlockData
	if !utils.IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

func (o *BlockSearchResultResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"entity_id",
		"entity_type",
		"block_data",
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

	varBlockSearchResultResult := _BlockSearchResultResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBlockSearchResultResult)

	if err != nil {
		return err
	}

	*o = BlockSearchResultResult(varBlockSearchResultResult)

	return err
}

type NullableBlockSearchResultResult struct {
	value *BlockSearchResultResult
	isSet bool
}

func (v NullableBlockSearchResultResult) Get() *BlockSearchResultResult {
	return v.value
}

func (v *NullableBlockSearchResultResult) Set(val *BlockSearchResultResult) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockSearchResultResult) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockSearchResultResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockSearchResultResult(val *BlockSearchResultResult) *NullableBlockSearchResultResult {
	return &NullableBlockSearchResultResult{value: val, isSet: true}
}

func (v NullableBlockSearchResultResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockSearchResultResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
