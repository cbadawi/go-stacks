package models

import (
	"encoding/json"
)

// ChainTip represents a ChainTip struct.
// Current chain tip information
type ChainTip struct {
	// the current block height
	BlockHeight int `json:"block_height"`
	// the current block hash
	BlockHash string `json:"block_hash"`
	// the current index block hash
	IndexBlockHash string `json:"index_block_hash"`
	// the current microblock hash
	MicroblockHash *string `json:"microblock_hash,omitempty"`
	// the current microblock sequence number
	MicroblockSequence *int `json:"microblock_sequence,omitempty"`
	// the current burn chain block height
	BurnBlockHeight int `json:"burn_block_height"`
}

// MarshalJSON implements the json.Marshaler interface for ChainTip.
// It customizes the JSON marshaling process for ChainTip objects.
func (c *ChainTip) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the ChainTip object to a map representation for JSON marshaling.
func (c *ChainTip) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["block_height"] = c.BlockHeight
	structMap["block_hash"] = c.BlockHash
	structMap["index_block_hash"] = c.IndexBlockHash
	if c.MicroblockHash != nil {
		structMap["microblock_hash"] = c.MicroblockHash
	}
	if c.MicroblockSequence != nil {
		structMap["microblock_sequence"] = c.MicroblockSequence
	}
	structMap["burn_block_height"] = c.BurnBlockHeight
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ChainTip.
// It customizes the JSON unmarshaling process for ChainTip objects.
func (c *ChainTip) UnmarshalJSON(input []byte) error {
	temp := &struct {
		BlockHeight        int     `json:"block_height"`
		BlockHash          string  `json:"block_hash"`
		IndexBlockHash     string  `json:"index_block_hash"`
		MicroblockHash     *string `json:"microblock_hash,omitempty"`
		MicroblockSequence *int    `json:"microblock_sequence,omitempty"`
		BurnBlockHeight    int     `json:"burn_block_height"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.BlockHeight = temp.BlockHeight
	c.BlockHash = temp.BlockHash
	c.IndexBlockHash = temp.IndexBlockHash
	c.MicroblockHash = temp.MicroblockHash
	c.MicroblockSequence = temp.MicroblockSequence
	c.BurnBlockHeight = temp.BurnBlockHeight
	return nil
}
