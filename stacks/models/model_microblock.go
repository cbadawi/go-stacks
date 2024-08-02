package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the Microblock type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Microblock{}

// Microblock A microblock
type Microblock struct {
	// Set to `true` if the microblock corresponds to the canonical chain tip.
	Canonical bool `json:"canonical"`
	// Set to `true` if the microblock was not orphaned in a following anchor block. Defaults to `true` if the following anchor block has not yet been created.
	MicroblockCanonical bool `json:"microblock_canonical"`
	// The SHA512/256 hash of this microblock.
	MicroblockHash string `json:"microblock_hash"`
	// A hint to describe how to order a set of microblocks. Starts at 0.
	MicroblockSequence int32 `json:"microblock_sequence"`
	// The SHA512/256 hash of the previous signed microblock in this stream.
	MicroblockParentHash string `json:"microblock_parent_hash"`
	// The anchor block height that confirmed this microblock.
	BlockHeight int32 `json:"block_height"`
	// The height of the anchor block that preceded this microblock.
	ParentBlockHeight int32 `json:"parent_block_height"`
	// The hash of the anchor block that preceded this microblock.
	ParentBlockHash string `json:"parent_block_hash"`
	// The hash of the Bitcoin block that preceded this microblock.
	ParentBurnBlockHash string `json:"parent_burn_block_hash"`
	// The block timestamp of the Bitcoin block that preceded this microblock.
	ParentBurnBlockTime int32 `json:"parent_burn_block_time"`
	// The ISO 8601 (YYYY-MM-DDTHH:mm:ss.sssZ) formatted block time of the bitcoin block that preceded this microblock.
	ParentBurnBlockTimeIso string `json:"parent_burn_block_time_iso"`
	// The height of the Bitcoin block that preceded this microblock.
	ParentBurnBlockHeight int32 `json:"parent_burn_block_height"`
	// The hash of the anchor block that confirmed this microblock. This wil be empty for unanchored microblocks
	BlockHash utils.NullableString `json:"block_hash"`
	// List of transactions included in the microblock
	Txs []string `json:"txs"`
}

type _Microblock Microblock

// NewMicroblock instantiates a new Microblock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicroblock(canonical bool, microblockCanonical bool, microblockHash string, microblockSequence int32, microblockParentHash string, blockHeight int32, parentBlockHeight int32, parentBlockHash string, parentBurnBlockHash string, parentBurnBlockTime int32, parentBurnBlockTimeIso string, parentBurnBlockHeight int32, blockHash utils.NullableString, txs []string) *Microblock {
	this := Microblock{}
	this.Canonical = canonical
	this.MicroblockCanonical = microblockCanonical
	this.MicroblockHash = microblockHash
	this.MicroblockSequence = microblockSequence
	this.MicroblockParentHash = microblockParentHash
	this.BlockHeight = blockHeight
	this.ParentBlockHeight = parentBlockHeight
	this.ParentBlockHash = parentBlockHash
	this.ParentBurnBlockHash = parentBurnBlockHash
	this.ParentBurnBlockTime = parentBurnBlockTime
	this.ParentBurnBlockTimeIso = parentBurnBlockTimeIso
	this.ParentBurnBlockHeight = parentBurnBlockHeight
	this.BlockHash = blockHash
	this.Txs = txs
	return &this
}

// NewMicroblockWithDefaults instantiates a new Microblock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicroblockWithDefaults() *Microblock {
	this := Microblock{}
	return &this
}

// GetCanonical returns the Canonical field value
func (o *Microblock) GetCanonical() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Canonical
}

// GetCanonicalOk returns a tuple with the Canonical field value
// and a boolean to check if the value has been set.
func (o *Microblock) GetCanonicalOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Canonical, true
}

// SetCanonical sets field value
func (o *Microblock) SetCanonical(v bool) {
	o.Canonical = v
}

// GetMicroblockCanonical returns the MicroblockCanonical field value
func (o *Microblock) GetMicroblockCanonical() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.MicroblockCanonical
}

// GetMicroblockCanonicalOk returns a tuple with the MicroblockCanonical field value
// and a boolean to check if the value has been set.
func (o *Microblock) GetMicroblockCanonicalOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MicroblockCanonical, true
}

// SetMicroblockCanonical sets field value
func (o *Microblock) SetMicroblockCanonical(v bool) {
	o.MicroblockCanonical = v
}

// GetMicroblockHash returns the MicroblockHash field value
func (o *Microblock) GetMicroblockHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MicroblockHash
}

// GetMicroblockHashOk returns a tuple with the MicroblockHash field value
// and a boolean to check if the value has been set.
func (o *Microblock) GetMicroblockHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MicroblockHash, true
}

// SetMicroblockHash sets field value
func (o *Microblock) SetMicroblockHash(v string) {
	o.MicroblockHash = v
}

// GetMicroblockSequence returns the MicroblockSequence field value
func (o *Microblock) GetMicroblockSequence() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MicroblockSequence
}

// GetMicroblockSequenceOk returns a tuple with the MicroblockSequence field value
// and a boolean to check if the value has been set.
func (o *Microblock) GetMicroblockSequenceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MicroblockSequence, true
}

// SetMicroblockSequence sets field value
func (o *Microblock) SetMicroblockSequence(v int32) {
	o.MicroblockSequence = v
}

// GetMicroblockParentHash returns the MicroblockParentHash field value
func (o *Microblock) GetMicroblockParentHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MicroblockParentHash
}

// GetMicroblockParentHashOk returns a tuple with the MicroblockParentHash field value
// and a boolean to check if the value has been set.
func (o *Microblock) GetMicroblockParentHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MicroblockParentHash, true
}

// SetMicroblockParentHash sets field value
func (o *Microblock) SetMicroblockParentHash(v string) {
	o.MicroblockParentHash = v
}

// GetBlockHeight returns the BlockHeight field value
func (o *Microblock) GetBlockHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BlockHeight
}

// GetBlockHeightOk returns a tuple with the BlockHeight field value
// and a boolean to check if the value has been set.
func (o *Microblock) GetBlockHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockHeight, true
}

// SetBlockHeight sets field value
func (o *Microblock) SetBlockHeight(v int32) {
	o.BlockHeight = v
}

// GetParentBlockHeight returns the ParentBlockHeight field value
func (o *Microblock) GetParentBlockHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ParentBlockHeight
}

// GetParentBlockHeightOk returns a tuple with the ParentBlockHeight field value
// and a boolean to check if the value has been set.
func (o *Microblock) GetParentBlockHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentBlockHeight, true
}

// SetParentBlockHeight sets field value
func (o *Microblock) SetParentBlockHeight(v int32) {
	o.ParentBlockHeight = v
}

// GetParentBlockHash returns the ParentBlockHash field value
func (o *Microblock) GetParentBlockHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParentBlockHash
}

// GetParentBlockHashOk returns a tuple with the ParentBlockHash field value
// and a boolean to check if the value has been set.
func (o *Microblock) GetParentBlockHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentBlockHash, true
}

// SetParentBlockHash sets field value
func (o *Microblock) SetParentBlockHash(v string) {
	o.ParentBlockHash = v
}

// GetParentBurnBlockHash returns the ParentBurnBlockHash field value
func (o *Microblock) GetParentBurnBlockHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParentBurnBlockHash
}

// GetParentBurnBlockHashOk returns a tuple with the ParentBurnBlockHash field value
// and a boolean to check if the value has been set.
func (o *Microblock) GetParentBurnBlockHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentBurnBlockHash, true
}

// SetParentBurnBlockHash sets field value
func (o *Microblock) SetParentBurnBlockHash(v string) {
	o.ParentBurnBlockHash = v
}

// GetParentBurnBlockTime returns the ParentBurnBlockTime field value
func (o *Microblock) GetParentBurnBlockTime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ParentBurnBlockTime
}

// GetParentBurnBlockTimeOk returns a tuple with the ParentBurnBlockTime field value
// and a boolean to check if the value has been set.
func (o *Microblock) GetParentBurnBlockTimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentBurnBlockTime, true
}

// SetParentBurnBlockTime sets field value
func (o *Microblock) SetParentBurnBlockTime(v int32) {
	o.ParentBurnBlockTime = v
}

// GetParentBurnBlockTimeIso returns the ParentBurnBlockTimeIso field value
func (o *Microblock) GetParentBurnBlockTimeIso() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParentBurnBlockTimeIso
}

// GetParentBurnBlockTimeIsoOk returns a tuple with the ParentBurnBlockTimeIso field value
// and a boolean to check if the value has been set.
func (o *Microblock) GetParentBurnBlockTimeIsoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentBurnBlockTimeIso, true
}

// SetParentBurnBlockTimeIso sets field value
func (o *Microblock) SetParentBurnBlockTimeIso(v string) {
	o.ParentBurnBlockTimeIso = v
}

// GetParentBurnBlockHeight returns the ParentBurnBlockHeight field value
func (o *Microblock) GetParentBurnBlockHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ParentBurnBlockHeight
}

// GetParentBurnBlockHeightOk returns a tuple with the ParentBurnBlockHeight field value
// and a boolean to check if the value has been set.
func (o *Microblock) GetParentBurnBlockHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentBurnBlockHeight, true
}

// SetParentBurnBlockHeight sets field value
func (o *Microblock) SetParentBurnBlockHeight(v int32) {
	o.ParentBurnBlockHeight = v
}

// GetBlockHash returns the BlockHash field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Microblock) GetBlockHash() string {
	if o == nil || o.BlockHash.Get() == nil {
		var ret string
		return ret
	}

	return *o.BlockHash.Get()
}

// GetBlockHashOk returns a tuple with the BlockHash field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Microblock) GetBlockHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BlockHash.Get(), o.BlockHash.IsSet()
}

// SetBlockHash sets field value
func (o *Microblock) SetBlockHash(v string) {
	o.BlockHash.Set(&v)
}

// GetTxs returns the Txs field value
func (o *Microblock) GetTxs() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Txs
}

// GetTxsOk returns a tuple with the Txs field value
// and a boolean to check if the value has been set.
func (o *Microblock) GetTxsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Txs, true
}

// SetTxs sets field value
func (o *Microblock) SetTxs(v []string) {
	o.Txs = v
}

func (o Microblock) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Microblock) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["canonical"] = o.Canonical
	toSerialize["microblock_canonical"] = o.MicroblockCanonical
	toSerialize["microblock_hash"] = o.MicroblockHash
	toSerialize["microblock_sequence"] = o.MicroblockSequence
	toSerialize["microblock_parent_hash"] = o.MicroblockParentHash
	toSerialize["block_height"] = o.BlockHeight
	toSerialize["parent_block_height"] = o.ParentBlockHeight
	toSerialize["parent_block_hash"] = o.ParentBlockHash
	toSerialize["parent_burn_block_hash"] = o.ParentBurnBlockHash
	toSerialize["parent_burn_block_time"] = o.ParentBurnBlockTime
	toSerialize["parent_burn_block_time_iso"] = o.ParentBurnBlockTimeIso
	toSerialize["parent_burn_block_height"] = o.ParentBurnBlockHeight
	toSerialize["block_hash"] = o.BlockHash.Get()
	toSerialize["txs"] = o.Txs
	return toSerialize, nil
}

func (o *Microblock) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"canonical",
		"microblock_canonical",
		"microblock_hash",
		"microblock_sequence",
		"microblock_parent_hash",
		"block_height",
		"parent_block_height",
		"parent_block_hash",
		"parent_burn_block_hash",
		"parent_burn_block_time",
		"parent_burn_block_time_iso",
		"parent_burn_block_height",
		"block_hash",
		"txs",
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

	varMicroblock := _Microblock{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMicroblock)

	if err != nil {
		return err
	}

	*o = Microblock(varMicroblock)

	return err
}

type NullableMicroblock struct {
	value *Microblock
	isSet bool
}

func (v NullableMicroblock) Get() *Microblock {
	return v.value
}

func (v *NullableMicroblock) Set(val *Microblock) {
	v.value = val
	v.isSet = true
}

func (v NullableMicroblock) IsSet() bool {
	return v.isSet
}

func (v *NullableMicroblock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicroblock(val *Microblock) *NullableMicroblock {
	return &NullableMicroblock{value: val, isSet: true}
}

func (v NullableMicroblock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicroblock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
