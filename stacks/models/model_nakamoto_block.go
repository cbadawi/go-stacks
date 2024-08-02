package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the NakamotoBlock type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &NakamotoBlock{}

// NakamotoBlock A block
type NakamotoBlock struct {
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
	// Index block hash of the parent block
	ParentIndexBlockHash string `json:"parent_index_block_hash"`
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
	// Number of transactions included in the block
	TxCount int32 `json:"tx_count"`
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
}

type _NakamotoBlock NakamotoBlock

// NewNakamotoBlock instantiates a new NakamotoBlock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNakamotoBlock(canonical bool, height int32, hash string, blockTime float32, blockTimeIso string, indexBlockHash string, parentBlockHash string, parentIndexBlockHash string, burnBlockTime float32, burnBlockTimeIso string, burnBlockHash string, burnBlockHeight int32, minerTxid string, txCount int32, executionCostReadCount int32, executionCostReadLength int32, executionCostRuntime int32, executionCostWriteCount int32, executionCostWriteLength int32) *NakamotoBlock {
	this := NakamotoBlock{}
	this.Canonical = canonical
	this.Height = height
	this.Hash = hash
	this.BlockTime = blockTime
	this.BlockTimeIso = blockTimeIso
	this.IndexBlockHash = indexBlockHash
	this.ParentBlockHash = parentBlockHash
	this.ParentIndexBlockHash = parentIndexBlockHash
	this.BurnBlockTime = burnBlockTime
	this.BurnBlockTimeIso = burnBlockTimeIso
	this.BurnBlockHash = burnBlockHash
	this.BurnBlockHeight = burnBlockHeight
	this.MinerTxid = minerTxid
	this.TxCount = txCount
	this.ExecutionCostReadCount = executionCostReadCount
	this.ExecutionCostReadLength = executionCostReadLength
	this.ExecutionCostRuntime = executionCostRuntime
	this.ExecutionCostWriteCount = executionCostWriteCount
	this.ExecutionCostWriteLength = executionCostWriteLength
	return &this
}

// NewNakamotoBlockWithDefaults instantiates a new NakamotoBlock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNakamotoBlockWithDefaults() *NakamotoBlock {
	this := NakamotoBlock{}
	return &this
}

// GetCanonical returns the Canonical field value
func (o *NakamotoBlock) GetCanonical() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Canonical
}

// GetCanonicalOk returns a tuple with the Canonical field value
// and a boolean to check if the value has been set.
func (o *NakamotoBlock) GetCanonicalOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Canonical, true
}

// SetCanonical sets field value
func (o *NakamotoBlock) SetCanonical(v bool) {
	o.Canonical = v
}

// GetHeight returns the Height field value
func (o *NakamotoBlock) GetHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Height
}

// GetHeightOk returns a tuple with the Height field value
// and a boolean to check if the value has been set.
func (o *NakamotoBlock) GetHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Height, true
}

// SetHeight sets field value
func (o *NakamotoBlock) SetHeight(v int32) {
	o.Height = v
}

// GetHash returns the Hash field value
func (o *NakamotoBlock) GetHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hash
}

// GetHashOk returns a tuple with the Hash field value
// and a boolean to check if the value has been set.
func (o *NakamotoBlock) GetHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hash, true
}

// SetHash sets field value
func (o *NakamotoBlock) SetHash(v string) {
	o.Hash = v
}

// GetBlockTime returns the BlockTime field value
func (o *NakamotoBlock) GetBlockTime() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.BlockTime
}

// GetBlockTimeOk returns a tuple with the BlockTime field value
// and a boolean to check if the value has been set.
func (o *NakamotoBlock) GetBlockTimeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockTime, true
}

// SetBlockTime sets field value
func (o *NakamotoBlock) SetBlockTime(v float32) {
	o.BlockTime = v
}

// GetBlockTimeIso returns the BlockTimeIso field value
func (o *NakamotoBlock) GetBlockTimeIso() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BlockTimeIso
}

// GetBlockTimeIsoOk returns a tuple with the BlockTimeIso field value
// and a boolean to check if the value has been set.
func (o *NakamotoBlock) GetBlockTimeIsoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockTimeIso, true
}

// SetBlockTimeIso sets field value
func (o *NakamotoBlock) SetBlockTimeIso(v string) {
	o.BlockTimeIso = v
}

// GetIndexBlockHash returns the IndexBlockHash field value
func (o *NakamotoBlock) GetIndexBlockHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IndexBlockHash
}

// GetIndexBlockHashOk returns a tuple with the IndexBlockHash field value
// and a boolean to check if the value has been set.
func (o *NakamotoBlock) GetIndexBlockHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IndexBlockHash, true
}

// SetIndexBlockHash sets field value
func (o *NakamotoBlock) SetIndexBlockHash(v string) {
	o.IndexBlockHash = v
}

// GetParentBlockHash returns the ParentBlockHash field value
func (o *NakamotoBlock) GetParentBlockHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParentBlockHash
}

// GetParentBlockHashOk returns a tuple with the ParentBlockHash field value
// and a boolean to check if the value has been set.
func (o *NakamotoBlock) GetParentBlockHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentBlockHash, true
}

// SetParentBlockHash sets field value
func (o *NakamotoBlock) SetParentBlockHash(v string) {
	o.ParentBlockHash = v
}

// GetParentIndexBlockHash returns the ParentIndexBlockHash field value
func (o *NakamotoBlock) GetParentIndexBlockHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParentIndexBlockHash
}

// GetParentIndexBlockHashOk returns a tuple with the ParentIndexBlockHash field value
// and a boolean to check if the value has been set.
func (o *NakamotoBlock) GetParentIndexBlockHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentIndexBlockHash, true
}

// SetParentIndexBlockHash sets field value
func (o *NakamotoBlock) SetParentIndexBlockHash(v string) {
	o.ParentIndexBlockHash = v
}

// GetBurnBlockTime returns the BurnBlockTime field value
func (o *NakamotoBlock) GetBurnBlockTime() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.BurnBlockTime
}

// GetBurnBlockTimeOk returns a tuple with the BurnBlockTime field value
// and a boolean to check if the value has been set.
func (o *NakamotoBlock) GetBurnBlockTimeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BurnBlockTime, true
}

// SetBurnBlockTime sets field value
func (o *NakamotoBlock) SetBurnBlockTime(v float32) {
	o.BurnBlockTime = v
}

// GetBurnBlockTimeIso returns the BurnBlockTimeIso field value
func (o *NakamotoBlock) GetBurnBlockTimeIso() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BurnBlockTimeIso
}

// GetBurnBlockTimeIsoOk returns a tuple with the BurnBlockTimeIso field value
// and a boolean to check if the value has been set.
func (o *NakamotoBlock) GetBurnBlockTimeIsoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BurnBlockTimeIso, true
}

// SetBurnBlockTimeIso sets field value
func (o *NakamotoBlock) SetBurnBlockTimeIso(v string) {
	o.BurnBlockTimeIso = v
}

// GetBurnBlockHash returns the BurnBlockHash field value
func (o *NakamotoBlock) GetBurnBlockHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BurnBlockHash
}

// GetBurnBlockHashOk returns a tuple with the BurnBlockHash field value
// and a boolean to check if the value has been set.
func (o *NakamotoBlock) GetBurnBlockHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BurnBlockHash, true
}

// SetBurnBlockHash sets field value
func (o *NakamotoBlock) SetBurnBlockHash(v string) {
	o.BurnBlockHash = v
}

// GetBurnBlockHeight returns the BurnBlockHeight field value
func (o *NakamotoBlock) GetBurnBlockHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BurnBlockHeight
}

// GetBurnBlockHeightOk returns a tuple with the BurnBlockHeight field value
// and a boolean to check if the value has been set.
func (o *NakamotoBlock) GetBurnBlockHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BurnBlockHeight, true
}

// SetBurnBlockHeight sets field value
func (o *NakamotoBlock) SetBurnBlockHeight(v int32) {
	o.BurnBlockHeight = v
}

// GetMinerTxid returns the MinerTxid field value
func (o *NakamotoBlock) GetMinerTxid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MinerTxid
}

// GetMinerTxidOk returns a tuple with the MinerTxid field value
// and a boolean to check if the value has been set.
func (o *NakamotoBlock) GetMinerTxidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinerTxid, true
}

// SetMinerTxid sets field value
func (o *NakamotoBlock) SetMinerTxid(v string) {
	o.MinerTxid = v
}

// GetTxCount returns the TxCount field value
func (o *NakamotoBlock) GetTxCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TxCount
}

// GetTxCountOk returns a tuple with the TxCount field value
// and a boolean to check if the value has been set.
func (o *NakamotoBlock) GetTxCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxCount, true
}

// SetTxCount sets field value
func (o *NakamotoBlock) SetTxCount(v int32) {
	o.TxCount = v
}

// GetExecutionCostReadCount returns the ExecutionCostReadCount field value
func (o *NakamotoBlock) GetExecutionCostReadCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExecutionCostReadCount
}

// GetExecutionCostReadCountOk returns a tuple with the ExecutionCostReadCount field value
// and a boolean to check if the value has been set.
func (o *NakamotoBlock) GetExecutionCostReadCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExecutionCostReadCount, true
}

// SetExecutionCostReadCount sets field value
func (o *NakamotoBlock) SetExecutionCostReadCount(v int32) {
	o.ExecutionCostReadCount = v
}

// GetExecutionCostReadLength returns the ExecutionCostReadLength field value
func (o *NakamotoBlock) GetExecutionCostReadLength() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExecutionCostReadLength
}

// GetExecutionCostReadLengthOk returns a tuple with the ExecutionCostReadLength field value
// and a boolean to check if the value has been set.
func (o *NakamotoBlock) GetExecutionCostReadLengthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExecutionCostReadLength, true
}

// SetExecutionCostReadLength sets field value
func (o *NakamotoBlock) SetExecutionCostReadLength(v int32) {
	o.ExecutionCostReadLength = v
}

// GetExecutionCostRuntime returns the ExecutionCostRuntime field value
func (o *NakamotoBlock) GetExecutionCostRuntime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExecutionCostRuntime
}

// GetExecutionCostRuntimeOk returns a tuple with the ExecutionCostRuntime field value
// and a boolean to check if the value has been set.
func (o *NakamotoBlock) GetExecutionCostRuntimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExecutionCostRuntime, true
}

// SetExecutionCostRuntime sets field value
func (o *NakamotoBlock) SetExecutionCostRuntime(v int32) {
	o.ExecutionCostRuntime = v
}

// GetExecutionCostWriteCount returns the ExecutionCostWriteCount field value
func (o *NakamotoBlock) GetExecutionCostWriteCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExecutionCostWriteCount
}

// GetExecutionCostWriteCountOk returns a tuple with the ExecutionCostWriteCount field value
// and a boolean to check if the value has been set.
func (o *NakamotoBlock) GetExecutionCostWriteCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExecutionCostWriteCount, true
}

// SetExecutionCostWriteCount sets field value
func (o *NakamotoBlock) SetExecutionCostWriteCount(v int32) {
	o.ExecutionCostWriteCount = v
}

// GetExecutionCostWriteLength returns the ExecutionCostWriteLength field value
func (o *NakamotoBlock) GetExecutionCostWriteLength() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExecutionCostWriteLength
}

// GetExecutionCostWriteLengthOk returns a tuple with the ExecutionCostWriteLength field value
// and a boolean to check if the value has been set.
func (o *NakamotoBlock) GetExecutionCostWriteLengthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExecutionCostWriteLength, true
}

// SetExecutionCostWriteLength sets field value
func (o *NakamotoBlock) SetExecutionCostWriteLength(v int32) {
	o.ExecutionCostWriteLength = v
}

func (o NakamotoBlock) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NakamotoBlock) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["canonical"] = o.Canonical
	toSerialize["height"] = o.Height
	toSerialize["hash"] = o.Hash
	toSerialize["block_time"] = o.BlockTime
	toSerialize["block_time_iso"] = o.BlockTimeIso
	toSerialize["index_block_hash"] = o.IndexBlockHash
	toSerialize["parent_block_hash"] = o.ParentBlockHash
	toSerialize["parent_index_block_hash"] = o.ParentIndexBlockHash
	toSerialize["burn_block_time"] = o.BurnBlockTime
	toSerialize["burn_block_time_iso"] = o.BurnBlockTimeIso
	toSerialize["burn_block_hash"] = o.BurnBlockHash
	toSerialize["burn_block_height"] = o.BurnBlockHeight
	toSerialize["miner_txid"] = o.MinerTxid
	toSerialize["tx_count"] = o.TxCount
	toSerialize["execution_cost_read_count"] = o.ExecutionCostReadCount
	toSerialize["execution_cost_read_length"] = o.ExecutionCostReadLength
	toSerialize["execution_cost_runtime"] = o.ExecutionCostRuntime
	toSerialize["execution_cost_write_count"] = o.ExecutionCostWriteCount
	toSerialize["execution_cost_write_length"] = o.ExecutionCostWriteLength
	return toSerialize, nil
}

func (o *NakamotoBlock) UnmarshalJSON(data []byte) (err error) {
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
		"parent_index_block_hash",
		"burn_block_time",
		"burn_block_time_iso",
		"burn_block_hash",
		"burn_block_height",
		"miner_txid",
		"tx_count",
		"execution_cost_read_count",
		"execution_cost_read_length",
		"execution_cost_runtime",
		"execution_cost_write_count",
		"execution_cost_write_length",
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

	varNakamotoBlock := _NakamotoBlock{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNakamotoBlock)

	if err != nil {
		return err
	}

	*o = NakamotoBlock(varNakamotoBlock)

	return err
}

type NullableNakamotoBlock struct {
	value *NakamotoBlock
	isSet bool
}

func (v NullableNakamotoBlock) Get() *NakamotoBlock {
	return v.value
}

func (v *NullableNakamotoBlock) Set(val *NakamotoBlock) {
	v.value = val
	v.isSet = true
}

func (v NullableNakamotoBlock) IsSet() bool {
	return v.isSet
}

func (v *NullableNakamotoBlock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNakamotoBlock(val *NakamotoBlock) *NullableNakamotoBlock {
	return &NullableNakamotoBlock{value: val, isSet: true}
}

func (v NullableNakamotoBlock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNakamotoBlock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
