package models

import (
	"encoding/json"
)

// BurnBlock represents a BurnBlock struct.
// A burn block
type BurnBlock struct {
	// Unix timestamp (in seconds) indicating when this block was mined.
	BurnBlockTime float64 `json:"burn_block_time"`
	// An ISO 8601 (YYYY-MM-DDTHH:mm:ss.sssZ) indicating when this block was mined.
	BurnBlockTimeIso string `json:"burn_block_time_iso"`
	// Hash of the anchor chain block
	BurnBlockHash string `json:"burn_block_hash"`
	// Height of the anchor chain block
	BurnBlockHeight int `json:"burn_block_height"`
	// Hashes of the Stacks blocks included in the burn block
	StacksBlocks []string `json:"stacks_blocks"`
}

// MarshalJSON implements the json.Marshaler interface for BurnBlock.
// It customizes the JSON marshaling process for BurnBlock objects.
func (b *BurnBlock) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(b.toMap())
}

// toMap converts the BurnBlock object to a map representation for JSON marshaling.
func (b *BurnBlock) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["burn_block_time"] = b.BurnBlockTime
	structMap["burn_block_time_iso"] = b.BurnBlockTimeIso
	structMap["burn_block_hash"] = b.BurnBlockHash
	structMap["burn_block_height"] = b.BurnBlockHeight
	structMap["stacks_blocks"] = b.StacksBlocks
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for BurnBlock.
// It customizes the JSON unmarshaling process for BurnBlock objects.
func (b *BurnBlock) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		BurnBlockTime    float64  `json:"burn_block_time"`
		BurnBlockTimeIso string   `json:"burn_block_time_iso"`
		BurnBlockHash    string   `json:"burn_block_hash"`
		BurnBlockHeight  int      `json:"burn_block_height"`
		StacksBlocks     []string `json:"stacks_blocks"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	b.BurnBlockTime = temp.BurnBlockTime
	b.BurnBlockTimeIso = temp.BurnBlockTimeIso
	b.BurnBlockHash = temp.BurnBlockHash
	b.BurnBlockHeight = temp.BurnBlockHeight
	b.StacksBlocks = temp.StacksBlocks
	return nil
}
