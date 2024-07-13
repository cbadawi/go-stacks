package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the Block type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Block{}

// Block A block
type Block struct {
	// Set to `true` if block corresponds to the canonical chain tip
	Canonical bool `json:"canonical"`
	// Height of the block
	Height int32 `json:"height"`
	// Hash representing the block
	Hash string `json:"hash"`
	// Unix timestamp (in seconds) indicating when this block was mined.
	BlockTime float32 `json:"block_time"`
	// An ISO 8601 (YYYY-MM-DDTHH:mm:ss.sssZ) indicating when this block was mined.
	BlockTimeIso string `json:"block_time_iso"`
	// The only hash that can uniquely identify an anchored block or an unconfirmed state trie
	IndexBlockHash string `json:"index_block_hash"`
	// Hash of the parent block
	ParentBlockHash string `json:"parent_block_hash"`
	// Unix timestamp (in seconds) indicating when this block was mined.
	BurnBlockTime float32 `json:"burn_block_time"`
	// An ISO 8601 (YYYY-MM-DDTHH:mm:ss.sssZ) indicating when this block was mined.
	BurnBlockTimeIso string `json:"burn_block_time_iso"`
	// Hash of the anchor chain block
	BurnBlockHash string `json:"burn_block_hash"`
	// Height of the anchor chain block
	BurnBlockHeight int32 `json:"burn_block_height"`
	// Anchor chain transaction ID
	MinerTxid string `json:"miner_txid"`
	// The hash of the last streamed block that precedes this block to which this block is to be appended. Not every anchored block will have a parent microblock stream. An anchored block that does not have a parent microblock stream has the parent microblock hash set to an empty string, and the parent microblock sequence number set to -1.
	ParentMicroblockHash string `json:"parent_microblock_hash"`
	// The hash of the last streamed block that precedes this block to which this block is to be appended. Not every anchored block will have a parent microblock stream. An anchored block that does not have a parent microblock stream has the parent microblock hash set to an empty string, and the parent microblock sequence number set to -1.
	ParentMicroblockSequence int32 `json:"parent_microblock_sequence"`
	// List of transactions included in the block
	Txs []string `json:"txs"`
	// List of microblocks that were accepted in this anchor block. Not every anchored block will have a accepted all (or any) of the previously streamed microblocks. Microblocks that were orphaned are not included in this list.
	MicroblocksAccepted []string `json:"microblocks_accepted"`
	// List of microblocks that were streamed/produced by this anchor block's miner. This list only includes microblocks that were accepted in the following anchor block. Microblocks that were orphaned are not included in this list.
	MicroblocksStreamed []string `json:"microblocks_streamed"`
	// Execution cost read count.
	ExecutionCostReadCount int32 `json:"execution_cost_read_count"`
	// Execution cost read length.
	ExecutionCostReadLength int32 `json:"execution_cost_read_length"`
	// Execution cost runtime.
	ExecutionCostRuntime int32 `json:"execution_cost_runtime"`
	// Execution cost write count.
	ExecutionCostWriteCount int32 `json:"execution_cost_write_count"`
	// Execution cost write length.
	ExecutionCostWriteLength int32 `json:"execution_cost_write_length"`
	// List of txs counts in each accepted microblock
	MicroblockTxCount map[string]float32 `json:"microblock_tx_count"`
}

type _Block Block

// NewBlock instantiates a new Block object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlock(canonical bool, height int32, hash string, blockTime float32, blockTimeIso string, indexBlockHash string, parentBlockHash string, burnBlockTime float32, burnBlockTimeIso string, burnBlockHash string, burnBlockHeight int32, minerTxid string, parentMicroblockHash string, parentMicroblockSequence int32, txs []string, microblocksAccepted []string, microblocksStreamed []string, executionCostReadCount int32, executionCostReadLength int32, executionCostRuntime int32, executionCostWriteCount int32, executionCostWriteLength int32, microblockTxCount map[string]float32) *Block {
	this := Block{}
	this.Canonical = canonical
	this.Height = height
	this.Hash = hash
	this.BlockTime = blockTime
	this.BlockTimeIso = blockTimeIso
	this.IndexBlockHash = indexBlockHash
	this.ParentBlockHash = parentBlockHash
	this.BurnBlockTime = burnBlockTime
	this.BurnBlockTimeIso = burnBlockTimeIso
	this.BurnBlockHash = burnBlockHash
	this.BurnBlockHeight = burnBlockHeight
	this.MinerTxid = minerTxid
	this.ParentMicroblockHash = parentMicroblockHash
	this.ParentMicroblockSequence = parentMicroblockSequence
	this.Txs = txs
	this.MicroblocksAccepted = microblocksAccepted
	this.MicroblocksStreamed = microblocksStreamed
	this.ExecutionCostReadCount = executionCostReadCount
	this.ExecutionCostReadLength = executionCostReadLength
	this.ExecutionCostRuntime = executionCostRuntime
	this.ExecutionCostWriteCount = executionCostWriteCount
	this.ExecutionCostWriteLength = executionCostWriteLength
	this.MicroblockTxCount = microblockTxCount
	return &this
}

// NewBlockWithDefaults instantiates a new Block object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockWithDefaults() *Block {
	this := Block{}
	return &this
}

// GetCanonical returns the Canonical field value
func (o *Block) GetCanonical() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Canonical
}

// GetCanonicalOk returns a tuple with the Canonical field value
// and a boolean to check if the value has been set.
func (o *Block) GetCanonicalOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Canonical, true
}

// SetCanonical sets field value
func (o *Block) SetCanonical(v bool) {
	o.Canonical = v
}

// GetHeight returns the Height field value
func (o *Block) GetHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Height
}

// GetHeightOk returns a tuple with the Height field value
// and a boolean to check if the value has been set.
func (o *Block) GetHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Height, true
}

// SetHeight sets field value
func (o *Block) SetHeight(v int32) {
	o.Height = v
}

// GetHash returns the Hash field value
func (o *Block) GetHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hash
}

// GetHashOk returns a tuple with the Hash field value
// and a boolean to check if the value has been set.
func (o *Block) GetHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hash, true
}

// SetHash sets field value
func (o *Block) SetHash(v string) {
	o.Hash = v
}

// GetBlockTime returns the BlockTime field value
func (o *Block) GetBlockTime() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.BlockTime
}

// GetBlockTimeOk returns a tuple with the BlockTime field value
// and a boolean to check if the value has been set.
func (o *Block) GetBlockTimeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockTime, true
}

// SetBlockTime sets field value
func (o *Block) SetBlockTime(v float32) {
	o.BlockTime = v
}

// GetBlockTimeIso returns the BlockTimeIso field value
func (o *Block) GetBlockTimeIso() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BlockTimeIso
}

// GetBlockTimeIsoOk returns a tuple with the BlockTimeIso field value
// and a boolean to check if the value has been set.
func (o *Block) GetBlockTimeIsoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockTimeIso, true
}

// SetBlockTimeIso sets field value
func (o *Block) SetBlockTimeIso(v string) {
	o.BlockTimeIso = v
}

// GetIndexBlockHash returns the IndexBlockHash field value
func (o *Block) GetIndexBlockHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IndexBlockHash
}

// GetIndexBlockHashOk returns a tuple with the IndexBlockHash field value
// and a boolean to check if the value has been set.
func (o *Block) GetIndexBlockHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IndexBlockHash, true
}

// SetIndexBlockHash sets field value
func (o *Block) SetIndexBlockHash(v string) {
	o.IndexBlockHash = v
}

// GetParentBlockHash returns the ParentBlockHash field value
func (o *Block) GetParentBlockHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParentBlockHash
}

// GetParentBlockHashOk returns a tuple with the ParentBlockHash field value
// and a boolean to check if the value has been set.
func (o *Block) GetParentBlockHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentBlockHash, true
}

// SetParentBlockHash sets field value
func (o *Block) SetParentBlockHash(v string) {
	o.ParentBlockHash = v
}

// GetBurnBlockTime returns the BurnBlockTime field value
func (o *Block) GetBurnBlockTime() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.BurnBlockTime
}

// GetBurnBlockTimeOk returns a tuple with the BurnBlockTime field value
// and a boolean to check if the value has been set.
func (o *Block) GetBurnBlockTimeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BurnBlockTime, true
}

// SetBurnBlockTime sets field value
func (o *Block) SetBurnBlockTime(v float32) {
	o.BurnBlockTime = v
}

// GetBurnBlockTimeIso returns the BurnBlockTimeIso field value
func (o *Block) GetBurnBlockTimeIso() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BurnBlockTimeIso
}

// GetBurnBlockTimeIsoOk returns a tuple with the BurnBlockTimeIso field value
// and a boolean to check if the value has been set.
func (o *Block) GetBurnBlockTimeIsoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BurnBlockTimeIso, true
}

// SetBurnBlockTimeIso sets field value
func (o *Block) SetBurnBlockTimeIso(v string) {
	o.BurnBlockTimeIso = v
}

// GetBurnBlockHash returns the BurnBlockHash field value
func (o *Block) GetBurnBlockHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BurnBlockHash
}

// GetBurnBlockHashOk returns a tuple with the BurnBlockHash field value
// and a boolean to check if the value has been set.
func (o *Block) GetBurnBlockHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BurnBlockHash, true
}

// SetBurnBlockHash sets field value
func (o *Block) SetBurnBlockHash(v string) {
	o.BurnBlockHash = v
}

// GetBurnBlockHeight returns the BurnBlockHeight field value
func (o *Block) GetBurnBlockHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BurnBlockHeight
}

// GetBurnBlockHeightOk returns a tuple with the BurnBlockHeight field value
// and a boolean to check if the value has been set.
func (o *Block) GetBurnBlockHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BurnBlockHeight, true
}

// SetBurnBlockHeight sets field value
func (o *Block) SetBurnBlockHeight(v int32) {
	o.BurnBlockHeight = v
}

// GetMinerTxid returns the MinerTxid field value
func (o *Block) GetMinerTxid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MinerTxid
}

// GetMinerTxidOk returns a tuple with the MinerTxid field value
// and a boolean to check if the value has been set.
func (o *Block) GetMinerTxidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinerTxid, true
}

// SetMinerTxid sets field value
func (o *Block) SetMinerTxid(v string) {
	o.MinerTxid = v
}

// GetParentMicroblockHash returns the ParentMicroblockHash field value
func (o *Block) GetParentMicroblockHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParentMicroblockHash
}

// GetParentMicroblockHashOk returns a tuple with the ParentMicroblockHash field value
// and a boolean to check if the value has been set.
func (o *Block) GetParentMicroblockHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentMicroblockHash, true
}

// SetParentMicroblockHash sets field value
func (o *Block) SetParentMicroblockHash(v string) {
	o.ParentMicroblockHash = v
}

// GetParentMicroblockSequence returns the ParentMicroblockSequence field value
func (o *Block) GetParentMicroblockSequence() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ParentMicroblockSequence
}

// GetParentMicroblockSequenceOk returns a tuple with the ParentMicroblockSequence field value
// and a boolean to check if the value has been set.
func (o *Block) GetParentMicroblockSequenceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentMicroblockSequence, true
}

// SetParentMicroblockSequence sets field value
func (o *Block) SetParentMicroblockSequence(v int32) {
	o.ParentMicroblockSequence = v
}

// GetTxs returns the Txs field value
func (o *Block) GetTxs() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Txs
}

// GetTxsOk returns a tuple with the Txs field value
// and a boolean to check if the value has been set.
func (o *Block) GetTxsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Txs, true
}

// SetTxs sets field value
func (o *Block) SetTxs(v []string) {
	o.Txs = v
}

// GetMicroblocksAccepted returns the MicroblocksAccepted field value
func (o *Block) GetMicroblocksAccepted() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.MicroblocksAccepted
}

// GetMicroblocksAcceptedOk returns a tuple with the MicroblocksAccepted field value
// and a boolean to check if the value has been set.
func (o *Block) GetMicroblocksAcceptedOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MicroblocksAccepted, true
}

// SetMicroblocksAccepted sets field value
func (o *Block) SetMicroblocksAccepted(v []string) {
	o.MicroblocksAccepted = v
}

// GetMicroblocksStreamed returns the MicroblocksStreamed field value
func (o *Block) GetMicroblocksStreamed() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.MicroblocksStreamed
}

// GetMicroblocksStreamedOk returns a tuple with the MicroblocksStreamed field value
// and a boolean to check if the value has been set.
func (o *Block) GetMicroblocksStreamedOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MicroblocksStreamed, true
}

// SetMicroblocksStreamed sets field value
func (o *Block) SetMicroblocksStreamed(v []string) {
	o.MicroblocksStreamed = v
}

// GetExecutionCostReadCount returns the ExecutionCostReadCount field value
func (o *Block) GetExecutionCostReadCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExecutionCostReadCount
}

// GetExecutionCostReadCountOk returns a tuple with the ExecutionCostReadCount field value
// and a boolean to check if the value has been set.
func (o *Block) GetExecutionCostReadCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExecutionCostReadCount, true
}

// SetExecutionCostReadCount sets field value
func (o *Block) SetExecutionCostReadCount(v int32) {
	o.ExecutionCostReadCount = v
}

// GetExecutionCostReadLength returns the ExecutionCostReadLength field value
func (o *Block) GetExecutionCostReadLength() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExecutionCostReadLength
}

// GetExecutionCostReadLengthOk returns a tuple with the ExecutionCostReadLength field value
// and a boolean to check if the value has been set.
func (o *Block) GetExecutionCostReadLengthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExecutionCostReadLength, true
}

// SetExecutionCostReadLength sets field value
func (o *Block) SetExecutionCostReadLength(v int32) {
	o.ExecutionCostReadLength = v
}

// GetExecutionCostRuntime returns the ExecutionCostRuntime field value
func (o *Block) GetExecutionCostRuntime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExecutionCostRuntime
}

// GetExecutionCostRuntimeOk returns a tuple with the ExecutionCostRuntime field value
// and a boolean to check if the value has been set.
func (o *Block) GetExecutionCostRuntimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExecutionCostRuntime, true
}

// SetExecutionCostRuntime sets field value
func (o *Block) SetExecutionCostRuntime(v int32) {
	o.ExecutionCostRuntime = v
}

// GetExecutionCostWriteCount returns the ExecutionCostWriteCount field value
func (o *Block) GetExecutionCostWriteCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExecutionCostWriteCount
}

// GetExecutionCostWriteCountOk returns a tuple with the ExecutionCostWriteCount field value
// and a boolean to check if the value has been set.
func (o *Block) GetExecutionCostWriteCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExecutionCostWriteCount, true
}

// SetExecutionCostWriteCount sets field value
func (o *Block) SetExecutionCostWriteCount(v int32) {
	o.ExecutionCostWriteCount = v
}

// GetExecutionCostWriteLength returns the ExecutionCostWriteLength field value
func (o *Block) GetExecutionCostWriteLength() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExecutionCostWriteLength
}

// GetExecutionCostWriteLengthOk returns a tuple with the ExecutionCostWriteLength field value
// and a boolean to check if the value has been set.
func (o *Block) GetExecutionCostWriteLengthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExecutionCostWriteLength, true
}

// SetExecutionCostWriteLength sets field value
func (o *Block) SetExecutionCostWriteLength(v int32) {
	o.ExecutionCostWriteLength = v
}

// GetMicroblockTxCount returns the MicroblockTxCount field value
func (o *Block) GetMicroblockTxCount() map[string]float32 {
	if o == nil {
		var ret map[string]float32
		return ret
	}

	return o.MicroblockTxCount
}

// GetMicroblockTxCountOk returns a tuple with the MicroblockTxCount field value
// and a boolean to check if the value has been set.
func (o *Block) GetMicroblockTxCountOk() (*map[string]float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MicroblockTxCount, true
}

// SetMicroblockTxCount sets field value
func (o *Block) SetMicroblockTxCount(v map[string]float32) {
	o.MicroblockTxCount = v
}

func (o Block) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Block) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["canonical"] = o.Canonical
	toSerialize["height"] = o.Height
	toSerialize["hash"] = o.Hash
	toSerialize["block_time"] = o.BlockTime
	toSerialize["block_time_iso"] = o.BlockTimeIso
	toSerialize["index_block_hash"] = o.IndexBlockHash
	toSerialize["parent_block_hash"] = o.ParentBlockHash
	toSerialize["burn_block_time"] = o.BurnBlockTime
	toSerialize["burn_block_time_iso"] = o.BurnBlockTimeIso
	toSerialize["burn_block_hash"] = o.BurnBlockHash
	toSerialize["burn_block_height"] = o.BurnBlockHeight
	toSerialize["miner_txid"] = o.MinerTxid
	toSerialize["parent_microblock_hash"] = o.ParentMicroblockHash
	toSerialize["parent_microblock_sequence"] = o.ParentMicroblockSequence
	toSerialize["txs"] = o.Txs
	toSerialize["microblocks_accepted"] = o.MicroblocksAccepted
	toSerialize["microblocks_streamed"] = o.MicroblocksStreamed
	toSerialize["execution_cost_read_count"] = o.ExecutionCostReadCount
	toSerialize["execution_cost_read_length"] = o.ExecutionCostReadLength
	toSerialize["execution_cost_runtime"] = o.ExecutionCostRuntime
	toSerialize["execution_cost_write_count"] = o.ExecutionCostWriteCount
	toSerialize["execution_cost_write_length"] = o.ExecutionCostWriteLength
	toSerialize["microblock_tx_count"] = o.MicroblockTxCount
	return toSerialize, nil
}

func (o *Block) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"canonical",
		"height",
		"hash",
		"block_time",
		"block_time_iso",
		"index_block_hash",
		"parent_block_hash",
		"burn_block_time",
		"burn_block_time_iso",
		"burn_block_hash",
		"burn_block_height",
		"miner_txid",
		"parent_microblock_hash",
		"parent_microblock_sequence",
		"txs",
		"microblocks_accepted",
		"microblocks_streamed",
		"execution_cost_read_count",
		"execution_cost_read_length",
		"execution_cost_runtime",
		"execution_cost_write_count",
		"execution_cost_write_length",
		"microblock_tx_count",
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

	varBlock := _Block{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBlock)

	if err != nil {
		return err
	}

	*o = Block(varBlock)

	return err
}

type NullableBlock struct {
	value *Block
	isSet bool
}

func (v NullableBlock) Get() *Block {
	return v.value
}

func (v *NullableBlock) Set(val *Block) {
	v.value = val
	v.isSet = true
}

func (v NullableBlock) IsSet() bool {
	return v.isSet
}

func (v *NullableBlock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlock(val *Block) *NullableBlock {
	return &NullableBlock{value: val, isSet: true}
}

func (v NullableBlock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
