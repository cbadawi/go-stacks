package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the BlockSearchResultResultBlockData type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &BlockSearchResultResultBlockData{}

// BlockSearchResultResultBlockData Returns basic search result information about the requested id
type BlockSearchResultResultBlockData struct {
	// If the block lies within the canonical chain
	Canonical bool `json:"canonical"`
	// Refers to the hash of the block
	Hash            string `json:"hash"`
	ParentBlockHash string `json:"parent_block_hash"`
	BurnBlockTime   int32  `json:"burn_block_time"`
	Height          int32  `json:"height"`
}

type _BlockSearchResultResultBlockData BlockSearchResultResultBlockData

// NewBlockSearchResultResultBlockData instantiates a new BlockSearchResultResultBlockData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockSearchResultResultBlockData(canonical bool, hash string, parentBlockHash string, burnBlockTime int32, height int32) *BlockSearchResultResultBlockData {
	this := BlockSearchResultResultBlockData{}
	this.Canonical = canonical
	this.Hash = hash
	this.ParentBlockHash = parentBlockHash
	this.BurnBlockTime = burnBlockTime
	this.Height = height
	return &this
}

// NewBlockSearchResultResultBlockDataWithDefaults instantiates a new BlockSearchResultResultBlockData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockSearchResultResultBlockDataWithDefaults() *BlockSearchResultResultBlockData {
	this := BlockSearchResultResultBlockData{}
	return &this
}

// GetCanonical returns the Canonical field value
func (o *BlockSearchResultResultBlockData) GetCanonical() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Canonical
}

// GetCanonicalOk returns a tuple with the Canonical field value
// and a boolean to check if the value has been set.
func (o *BlockSearchResultResultBlockData) GetCanonicalOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Canonical, true
}

// SetCanonical sets field value
func (o *BlockSearchResultResultBlockData) SetCanonical(v bool) {
	o.Canonical = v
}

// GetHash returns the Hash field value
func (o *BlockSearchResultResultBlockData) GetHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hash
}

// GetHashOk returns a tuple with the Hash field value
// and a boolean to check if the value has been set.
func (o *BlockSearchResultResultBlockData) GetHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hash, true
}

// SetHash sets field value
func (o *BlockSearchResultResultBlockData) SetHash(v string) {
	o.Hash = v
}

// GetParentBlockHash returns the ParentBlockHash field value
func (o *BlockSearchResultResultBlockData) GetParentBlockHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParentBlockHash
}

// GetParentBlockHashOk returns a tuple with the ParentBlockHash field value
// and a boolean to check if the value has been set.
func (o *BlockSearchResultResultBlockData) GetParentBlockHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentBlockHash, true
}

// SetParentBlockHash sets field value
func (o *BlockSearchResultResultBlockData) SetParentBlockHash(v string) {
	o.ParentBlockHash = v
}

// GetBurnBlockTime returns the BurnBlockTime field value
func (o *BlockSearchResultResultBlockData) GetBurnBlockTime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BurnBlockTime
}

// GetBurnBlockTimeOk returns a tuple with the BurnBlockTime field value
// and a boolean to check if the value has been set.
func (o *BlockSearchResultResultBlockData) GetBurnBlockTimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BurnBlockTime, true
}

// SetBurnBlockTime sets field value
func (o *BlockSearchResultResultBlockData) SetBurnBlockTime(v int32) {
	o.BurnBlockTime = v
}

// GetHeight returns the Height field value
func (o *BlockSearchResultResultBlockData) GetHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Height
}

// GetHeightOk returns a tuple with the Height field value
// and a boolean to check if the value has been set.
func (o *BlockSearchResultResultBlockData) GetHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Height, true
}

// SetHeight sets field value
func (o *BlockSearchResultResultBlockData) SetHeight(v int32) {
	o.Height = v
}

func (o BlockSearchResultResultBlockData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlockSearchResultResultBlockData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["canonical"] = o.Canonical
	toSerialize["hash"] = o.Hash
	toSerialize["parent_block_hash"] = o.ParentBlockHash
	toSerialize["burn_block_time"] = o.BurnBlockTime
	toSerialize["height"] = o.Height
	return toSerialize, nil
}

func (o *BlockSearchResultResultBlockData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"canonical",
		"hash",
		"parent_block_hash",
		"burn_block_time",
		"height",
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

	varBlockSearchResultResultBlockData := _BlockSearchResultResultBlockData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBlockSearchResultResultBlockData)

	if err != nil {
		return err
	}

	*o = BlockSearchResultResultBlockData(varBlockSearchResultResultBlockData)

	return err
}

type NullableBlockSearchResultResultBlockData struct {
	value *BlockSearchResultResultBlockData
	isSet bool
}

func (v NullableBlockSearchResultResultBlockData) Get() *BlockSearchResultResultBlockData {
	return v.value
}

func (v *NullableBlockSearchResultResultBlockData) Set(val *BlockSearchResultResultBlockData) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockSearchResultResultBlockData) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockSearchResultResultBlockData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockSearchResultResultBlockData(val *BlockSearchResultResultBlockData) *NullableBlockSearchResultResultBlockData {
	return &NullableBlockSearchResultResultBlockData{value: val, isSet: true}
}

func (v NullableBlockSearchResultResultBlockData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockSearchResultResultBlockData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
