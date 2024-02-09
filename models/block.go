package models

import (
	"encoding/json"
)

// Block represents a Block struct.
// A block
type Block struct {
	// Set to `true` if block corresponds to the canonical chain tip
	Canonical bool `json:"canonical"`
	// Height of the block
	Height int `json:"height"`
	// Hash representing the block
	Hash string `json:"hash"`
	// The only hash that can uniquely identify an anchored block or an unconfirmed state trie
	IndexBlockHash string `json:"index_block_hash"`
	// Hash of the parent block
	ParentBlockHash string `json:"parent_block_hash"`
	// Unix timestamp (in seconds) indicating when this block was mined.
	BurnBlockTime float64 `json:"burn_block_time"`
	// An ISO 8601 (YYYY-MM-DDTHH:mm:ss.sssZ) indicating when this block was mined.
	BurnBlockTimeIso string `json:"burn_block_time_iso"`
	// Hash of the anchor chain block
	BurnBlockHash string `json:"burn_block_hash"`
	// Height of the anchor chain block
	BurnBlockHeight int `json:"burn_block_height"`
	// Anchor chain transaction ID
	MinerTxid string `json:"miner_txid"`
	// The hash of the last streamed block that precedes this block to which this block is to be appended. Not every anchored block will have a parent microblock stream. An anchored block that does not have a parent microblock stream has the parent microblock hash set to an empty string, and the parent microblock sequence number set to -1.
	ParentMicroblockHash string `json:"parent_microblock_hash"`
	// The hash of the last streamed block that precedes this block to which this block is to be appended. Not every anchored block will have a parent microblock stream. An anchored block that does not have a parent microblock stream has the parent microblock hash set to an empty string, and the parent microblock sequence number set to -1.
	ParentMicroblockSequence int `json:"parent_microblock_sequence"`
	// List of transactions included in the block
	Txs []string `json:"txs"`
	// List of microblocks that were accepted in this anchor block. Not every anchored block will have a accepted all (or any) of the previously streamed microblocks. Microblocks that were orphaned are not included in this list.
	MicroblocksAccepted []string `json:"microblocks_accepted"`
	// List of microblocks that were streamed/produced by this anchor block's miner. This list only includes microblocks that were accepted in the following anchor block. Microblocks that were orphaned are not included in this list.
	MicroblocksStreamed []string `json:"microblocks_streamed"`
	// Execution cost read count.
	ExecutionCostReadCount int `json:"execution_cost_read_count"`
	// Execution cost read length.
	ExecutionCostReadLength int `json:"execution_cost_read_length"`
	// Execution cost runtime.
	ExecutionCostRuntime int `json:"execution_cost_runtime"`
	// Execution cost write count.
	ExecutionCostWriteCount int `json:"execution_cost_write_count"`
	// Execution cost write length.
	ExecutionCostWriteLength int `json:"execution_cost_write_length"`
	// List of txs counts in each accepted microblock
	MicroblockTxCount map[string]float64 `json:"microblock_tx_count"`
}

// MarshalJSON implements the json.Marshaler interface for Block.
// It customizes the JSON marshaling process for Block objects.
func (b *Block) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(b.toMap())
}

// toMap converts the Block object to a map representation for JSON marshaling.
func (b *Block) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["canonical"] = b.Canonical
	structMap["height"] = b.Height
	structMap["hash"] = b.Hash
	structMap["index_block_hash"] = b.IndexBlockHash
	structMap["parent_block_hash"] = b.ParentBlockHash
	structMap["burn_block_time"] = b.BurnBlockTime
	structMap["burn_block_time_iso"] = b.BurnBlockTimeIso
	structMap["burn_block_hash"] = b.BurnBlockHash
	structMap["burn_block_height"] = b.BurnBlockHeight
	structMap["miner_txid"] = b.MinerTxid
	structMap["parent_microblock_hash"] = b.ParentMicroblockHash
	structMap["parent_microblock_sequence"] = b.ParentMicroblockSequence
	structMap["txs"] = b.Txs
	structMap["microblocks_accepted"] = b.MicroblocksAccepted
	structMap["microblocks_streamed"] = b.MicroblocksStreamed
	structMap["execution_cost_read_count"] = b.ExecutionCostReadCount
	structMap["execution_cost_read_length"] = b.ExecutionCostReadLength
	structMap["execution_cost_runtime"] = b.ExecutionCostRuntime
	structMap["execution_cost_write_count"] = b.ExecutionCostWriteCount
	structMap["execution_cost_write_length"] = b.ExecutionCostWriteLength
	structMap["microblock_tx_count"] = b.MicroblockTxCount
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Block.
// It customizes the JSON unmarshaling process for Block objects.
func (b *Block) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		Canonical                bool               `json:"canonical"`
		Height                   int                `json:"height"`
		Hash                     string             `json:"hash"`
		IndexBlockHash           string             `json:"index_block_hash"`
		ParentBlockHash          string             `json:"parent_block_hash"`
		BurnBlockTime            float64            `json:"burn_block_time"`
		BurnBlockTimeIso         string             `json:"burn_block_time_iso"`
		BurnBlockHash            string             `json:"burn_block_hash"`
		BurnBlockHeight          int                `json:"burn_block_height"`
		MinerTxid                string             `json:"miner_txid"`
		ParentMicroblockHash     string             `json:"parent_microblock_hash"`
		ParentMicroblockSequence int                `json:"parent_microblock_sequence"`
		Txs                      []string           `json:"txs"`
		MicroblocksAccepted      []string           `json:"microblocks_accepted"`
		MicroblocksStreamed      []string           `json:"microblocks_streamed"`
		ExecutionCostReadCount   int                `json:"execution_cost_read_count"`
		ExecutionCostReadLength  int                `json:"execution_cost_read_length"`
		ExecutionCostRuntime     int                `json:"execution_cost_runtime"`
		ExecutionCostWriteCount  int                `json:"execution_cost_write_count"`
		ExecutionCostWriteLength int                `json:"execution_cost_write_length"`
		MicroblockTxCount        map[string]float64 `json:"microblock_tx_count"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	b.Canonical = temp.Canonical
	b.Height = temp.Height
	b.Hash = temp.Hash
	b.IndexBlockHash = temp.IndexBlockHash
	b.ParentBlockHash = temp.ParentBlockHash
	b.BurnBlockTime = temp.BurnBlockTime
	b.BurnBlockTimeIso = temp.BurnBlockTimeIso
	b.BurnBlockHash = temp.BurnBlockHash
	b.BurnBlockHeight = temp.BurnBlockHeight
	b.MinerTxid = temp.MinerTxid
	b.ParentMicroblockHash = temp.ParentMicroblockHash
	b.ParentMicroblockSequence = temp.ParentMicroblockSequence
	b.Txs = temp.Txs
	b.MicroblocksAccepted = temp.MicroblocksAccepted
	b.MicroblocksStreamed = temp.MicroblocksStreamed
	b.ExecutionCostReadCount = temp.ExecutionCostReadCount
	b.ExecutionCostReadLength = temp.ExecutionCostReadLength
	b.ExecutionCostRuntime = temp.ExecutionCostRuntime
	b.ExecutionCostWriteCount = temp.ExecutionCostWriteCount
	b.ExecutionCostWriteLength = temp.ExecutionCostWriteLength
	b.MicroblockTxCount = temp.MicroblockTxCount
	return nil
}
