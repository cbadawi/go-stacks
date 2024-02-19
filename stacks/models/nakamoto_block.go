package models

import (
	"encoding/json"
)

// NakamotoBlock represents a NakamotoBlock struct.
// A block
type NakamotoBlock struct {
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
	// Index block hash of the parent block
	ParentIndexBlockHash string `json:"parent_index_block_hash"`
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
	// Number of transactions included in the block
	TxCount int `json:"tx_count"`
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
}

// MarshalJSON implements the json.Marshaler interface for NakamotoBlock.
// It customizes the JSON marshaling process for NakamotoBlock objects.
func (n *NakamotoBlock) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(n.toMap())
}

// toMap converts the NakamotoBlock object to a map representation for JSON marshaling.
func (n *NakamotoBlock) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["canonical"] = n.Canonical
	structMap["height"] = n.Height
	structMap["hash"] = n.Hash
	structMap["index_block_hash"] = n.IndexBlockHash
	structMap["parent_block_hash"] = n.ParentBlockHash
	structMap["parent_index_block_hash"] = n.ParentIndexBlockHash
	structMap["burn_block_time"] = n.BurnBlockTime
	structMap["burn_block_time_iso"] = n.BurnBlockTimeIso
	structMap["burn_block_hash"] = n.BurnBlockHash
	structMap["burn_block_height"] = n.BurnBlockHeight
	structMap["miner_txid"] = n.MinerTxid
	structMap["tx_count"] = n.TxCount
	structMap["execution_cost_read_count"] = n.ExecutionCostReadCount
	structMap["execution_cost_read_length"] = n.ExecutionCostReadLength
	structMap["execution_cost_runtime"] = n.ExecutionCostRuntime
	structMap["execution_cost_write_count"] = n.ExecutionCostWriteCount
	structMap["execution_cost_write_length"] = n.ExecutionCostWriteLength
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for NakamotoBlock.
// It customizes the JSON unmarshaling process for NakamotoBlock objects.
func (n *NakamotoBlock) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		Canonical                bool    `json:"canonical"`
		Height                   int     `json:"height"`
		Hash                     string  `json:"hash"`
		IndexBlockHash           string  `json:"index_block_hash"`
		ParentBlockHash          string  `json:"parent_block_hash"`
		ParentIndexBlockHash     string  `json:"parent_index_block_hash"`
		BurnBlockTime            float64 `json:"burn_block_time"`
		BurnBlockTimeIso         string  `json:"burn_block_time_iso"`
		BurnBlockHash            string  `json:"burn_block_hash"`
		BurnBlockHeight          int     `json:"burn_block_height"`
		MinerTxid                string  `json:"miner_txid"`
		TxCount                  int     `json:"tx_count"`
		ExecutionCostReadCount   int     `json:"execution_cost_read_count"`
		ExecutionCostReadLength  int     `json:"execution_cost_read_length"`
		ExecutionCostRuntime     int     `json:"execution_cost_runtime"`
		ExecutionCostWriteCount  int     `json:"execution_cost_write_count"`
		ExecutionCostWriteLength int     `json:"execution_cost_write_length"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	n.Canonical = temp.Canonical
	n.Height = temp.Height
	n.Hash = temp.Hash
	n.IndexBlockHash = temp.IndexBlockHash
	n.ParentBlockHash = temp.ParentBlockHash
	n.ParentIndexBlockHash = temp.ParentIndexBlockHash
	n.BurnBlockTime = temp.BurnBlockTime
	n.BurnBlockTimeIso = temp.BurnBlockTimeIso
	n.BurnBlockHash = temp.BurnBlockHash
	n.BurnBlockHeight = temp.BurnBlockHeight
	n.MinerTxid = temp.MinerTxid
	n.TxCount = temp.TxCount
	n.ExecutionCostReadCount = temp.ExecutionCostReadCount
	n.ExecutionCostReadLength = temp.ExecutionCostReadLength
	n.ExecutionCostRuntime = temp.ExecutionCostRuntime
	n.ExecutionCostWriteCount = temp.ExecutionCostWriteCount
	n.ExecutionCostWriteLength = temp.ExecutionCostWriteLength
	return nil
}
