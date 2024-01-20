package models

import (
	"encoding/json"
)

// BlockData represents a BlockData struct.
// Returns basic search result information about the requested id
type BlockData struct {
	// If the block lies within the canonical chain
	Canonical bool `json:"canonical"`
	// Refers to the hash of the block
	Hash            string `json:"hash"`
	ParentBlockHash string `json:"parent_block_hash"`
	BurnBlockTime   int    `json:"burn_block_time"`
	Height          int    `json:"height"`
}

// MarshalJSON implements the json.Marshaler interface for BlockData.
// It customizes the JSON marshaling process for BlockData objects.
func (b *BlockData) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(b.toMap())
}

// toMap converts the BlockData object to a map representation for JSON marshaling.
func (b *BlockData) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["canonical"] = b.Canonical
	structMap["hash"] = b.Hash
	structMap["parent_block_hash"] = b.ParentBlockHash
	structMap["burn_block_time"] = b.BurnBlockTime
	structMap["height"] = b.Height
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for BlockData.
// It customizes the JSON unmarshaling process for BlockData objects.
func (b *BlockData) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Canonical       bool   `json:"canonical"`
		Hash            string `json:"hash"`
		ParentBlockHash string `json:"parent_block_hash"`
		BurnBlockTime   int    `json:"burn_block_time"`
		Height          int    `json:"height"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	b.Canonical = temp.Canonical
	b.Hash = temp.Hash
	b.ParentBlockHash = temp.ParentBlockHash
	b.BurnBlockTime = temp.BurnBlockTime
	b.Height = temp.Height
	return nil
}
