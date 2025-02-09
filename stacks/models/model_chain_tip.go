package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the ChainTip type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ChainTip{}

// ChainTip Current chain tip information
type ChainTip struct {
	// the current block height
	BlockHeight int32 `json:"block_height"`
	// the current block hash
	BlockHash string `json:"block_hash"`
	// the current index block hash
	IndexBlockHash string `json:"index_block_hash"`
	// the current microblock hash
	MicroblockHash *string `json:"microblock_hash,omitempty"`
	// the current microblock sequence number
	MicroblockSequence *int32 `json:"microblock_sequence,omitempty"`
	// the current burn chain block height
	BurnBlockHeight int32 `json:"burn_block_height"`
}

type _ChainTip ChainTip

// NewChainTip instantiates a new ChainTip object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChainTip(blockHeight int32, blockHash string, indexBlockHash string, burnBlockHeight int32) *ChainTip {
	this := ChainTip{}
	this.BlockHeight = blockHeight
	this.BlockHash = blockHash
	this.IndexBlockHash = indexBlockHash
	this.BurnBlockHeight = burnBlockHeight
	return &this
}

// NewChainTipWithDefaults instantiates a new ChainTip object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChainTipWithDefaults() *ChainTip {
	this := ChainTip{}
	return &this
}

// GetBlockHeight returns the BlockHeight field value
func (o *ChainTip) GetBlockHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BlockHeight
}

// GetBlockHeightOk returns a tuple with the BlockHeight field value
// and a boolean to check if the value has been set.
func (o *ChainTip) GetBlockHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockHeight, true
}

// SetBlockHeight sets field value
func (o *ChainTip) SetBlockHeight(v int32) {
	o.BlockHeight = v
}

// GetBlockHash returns the BlockHash field value
func (o *ChainTip) GetBlockHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BlockHash
}

// GetBlockHashOk returns a tuple with the BlockHash field value
// and a boolean to check if the value has been set.
func (o *ChainTip) GetBlockHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockHash, true
}

// SetBlockHash sets field value
func (o *ChainTip) SetBlockHash(v string) {
	o.BlockHash = v
}

// GetIndexBlockHash returns the IndexBlockHash field value
func (o *ChainTip) GetIndexBlockHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IndexBlockHash
}

// GetIndexBlockHashOk returns a tuple with the IndexBlockHash field value
// and a boolean to check if the value has been set.
func (o *ChainTip) GetIndexBlockHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IndexBlockHash, true
}

// SetIndexBlockHash sets field value
func (o *ChainTip) SetIndexBlockHash(v string) {
	o.IndexBlockHash = v
}

// GetMicroblockHash returns the MicroblockHash field value if set, zero value otherwise.
func (o *ChainTip) GetMicroblockHash() string {
	if o == nil || utils.IsNil(o.MicroblockHash) {
		var ret string
		return ret
	}
	return *o.MicroblockHash
}

// GetMicroblockHashOk returns a tuple with the MicroblockHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChainTip) GetMicroblockHashOk() (*string, bool) {
	if o == nil || utils.IsNil(o.MicroblockHash) {
		return nil, false
	}
	return o.MicroblockHash, true
}

// HasMicroblockHash returns a boolean if a field has been set.
func (o *ChainTip) HasMicroblockHash() bool {
	if o != nil && !utils.IsNil(o.MicroblockHash) {
		return true
	}

	return false
}

// SetMicroblockHash gets a reference to the given string and assigns it to the MicroblockHash field.
func (o *ChainTip) SetMicroblockHash(v string) {
	o.MicroblockHash = &v
}

// GetMicroblockSequence returns the MicroblockSequence field value if set, zero value otherwise.
func (o *ChainTip) GetMicroblockSequence() int32 {
	if o == nil || utils.IsNil(o.MicroblockSequence) {
		var ret int32
		return ret
	}
	return *o.MicroblockSequence
}

// GetMicroblockSequenceOk returns a tuple with the MicroblockSequence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChainTip) GetMicroblockSequenceOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.MicroblockSequence) {
		return nil, false
	}
	return o.MicroblockSequence, true
}

// HasMicroblockSequence returns a boolean if a field has been set.
func (o *ChainTip) HasMicroblockSequence() bool {
	if o != nil && !utils.IsNil(o.MicroblockSequence) {
		return true
	}

	return false
}

// SetMicroblockSequence gets a reference to the given int32 and assigns it to the MicroblockSequence field.
func (o *ChainTip) SetMicroblockSequence(v int32) {
	o.MicroblockSequence = &v
}

// GetBurnBlockHeight returns the BurnBlockHeight field value
func (o *ChainTip) GetBurnBlockHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BurnBlockHeight
}

// GetBurnBlockHeightOk returns a tuple with the BurnBlockHeight field value
// and a boolean to check if the value has been set.
func (o *ChainTip) GetBurnBlockHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BurnBlockHeight, true
}

// SetBurnBlockHeight sets field value
func (o *ChainTip) SetBurnBlockHeight(v int32) {
	o.BurnBlockHeight = v
}

func (o ChainTip) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChainTip) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["block_height"] = o.BlockHeight
	toSerialize["block_hash"] = o.BlockHash
	toSerialize["index_block_hash"] = o.IndexBlockHash
	if !utils.IsNil(o.MicroblockHash) {
		toSerialize["microblock_hash"] = o.MicroblockHash
	}
	if !utils.IsNil(o.MicroblockSequence) {
		toSerialize["microblock_sequence"] = o.MicroblockSequence
	}
	toSerialize["burn_block_height"] = o.BurnBlockHeight
	return toSerialize, nil
}

func (o *ChainTip) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"block_height",
		"block_hash",
		"index_block_hash",
		"burn_block_height",
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

	varChainTip := _ChainTip{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varChainTip)

	if err != nil {
		return err
	}

	*o = ChainTip(varChainTip)

	return err
}

type NullableChainTip struct {
	value *ChainTip
	isSet bool
}

func (v NullableChainTip) Get() *ChainTip {
	return v.value
}

func (v *NullableChainTip) Set(val *ChainTip) {
	v.value = val
	v.isSet = true
}

func (v NullableChainTip) IsSet() bool {
	return v.isSet
}

func (v *NullableChainTip) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChainTip(val *ChainTip) *NullableChainTip {
	return &NullableChainTip{value: val, isSet: true}
}

func (v NullableChainTip) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChainTip) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
