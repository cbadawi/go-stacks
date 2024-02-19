package models

import (
	"encoding/json"
)

// Microblock represents a Microblock struct.
// A microblock
type Microblock struct {
	// Set to `true` if the microblock corresponds to the canonical chain tip.
	Canonical bool `json:"canonical"`
	// Set to `true` if the microblock was not orphaned in a following anchor block. Defaults to `true` if the following anchor block has not yet been created.
	MicroblockCanonical bool `json:"microblock_canonical"`
	// The SHA512/256 hash of this microblock.
	MicroblockHash string `json:"microblock_hash"`
	// A hint to describe how to order a set of microblocks. Starts at 0.
	MicroblockSequence int `json:"microblock_sequence"`
	// The SHA512/256 hash of the previous signed microblock in this stream.
	MicroblockParentHash string `json:"microblock_parent_hash"`
	// The anchor block height that confirmed this microblock.
	BlockHeight int `json:"block_height"`
	// The height of the anchor block that preceded this microblock.
	ParentBlockHeight int `json:"parent_block_height"`
	// The hash of the anchor block that preceded this microblock.
	ParentBlockHash string `json:"parent_block_hash"`
	// The hash of the Bitcoin block that preceded this microblock.
	ParentBurnBlockHash string `json:"parent_burn_block_hash"`
	// The block timestamp of the Bitcoin block that preceded this microblock.
	ParentBurnBlockTime int `json:"parent_burn_block_time"`
	// The ISO 8601 (YYYY-MM-DDTHH:mm:ss.sssZ) formatted block time of the bitcoin block that preceded this microblock.
	ParentBurnBlockTimeIso string `json:"parent_burn_block_time_iso"`
	// The height of the Bitcoin block that preceded this microblock.
	ParentBurnBlockHeight int `json:"parent_burn_block_height"`
	// The hash of the anchor block that confirmed this microblock. This wil be empty for unanchored microblocks
	BlockHash *string `json:"block_hash"`
	// List of transactions included in the microblock
	Txs []string `json:"txs"`
}

// MarshalJSON implements the json.Marshaler interface for Microblock.
// It customizes the JSON marshaling process for Microblock objects.
func (m *Microblock) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(m.toMap())
}

// toMap converts the Microblock object to a map representation for JSON marshaling.
func (m *Microblock) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["canonical"] = m.Canonical
	structMap["microblock_canonical"] = m.MicroblockCanonical
	structMap["microblock_hash"] = m.MicroblockHash
	structMap["microblock_sequence"] = m.MicroblockSequence
	structMap["microblock_parent_hash"] = m.MicroblockParentHash
	structMap["block_height"] = m.BlockHeight
	structMap["parent_block_height"] = m.ParentBlockHeight
	structMap["parent_block_hash"] = m.ParentBlockHash
	structMap["parent_burn_block_hash"] = m.ParentBurnBlockHash
	structMap["parent_burn_block_time"] = m.ParentBurnBlockTime
	structMap["parent_burn_block_time_iso"] = m.ParentBurnBlockTimeIso
	structMap["parent_burn_block_height"] = m.ParentBurnBlockHeight
	structMap["block_hash"] = m.BlockHash
	structMap["txs"] = m.Txs
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Microblock.
// It customizes the JSON unmarshaling process for Microblock objects.
func (m *Microblock) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Canonical              bool     `json:"canonical"`
		MicroblockCanonical    bool     `json:"microblock_canonical"`
		MicroblockHash         string   `json:"microblock_hash"`
		MicroblockSequence     int      `json:"microblock_sequence"`
		MicroblockParentHash   string   `json:"microblock_parent_hash"`
		BlockHeight            int      `json:"block_height"`
		ParentBlockHeight      int      `json:"parent_block_height"`
		ParentBlockHash        string   `json:"parent_block_hash"`
		ParentBurnBlockHash    string   `json:"parent_burn_block_hash"`
		ParentBurnBlockTime    int      `json:"parent_burn_block_time"`
		ParentBurnBlockTimeIso string   `json:"parent_burn_block_time_iso"`
		ParentBurnBlockHeight  int      `json:"parent_burn_block_height"`
		BlockHash              *string  `json:"block_hash"`
		Txs                    []string `json:"txs"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	m.Canonical = temp.Canonical
	m.MicroblockCanonical = temp.MicroblockCanonical
	m.MicroblockHash = temp.MicroblockHash
	m.MicroblockSequence = temp.MicroblockSequence
	m.MicroblockParentHash = temp.MicroblockParentHash
	m.BlockHeight = temp.BlockHeight
	m.ParentBlockHeight = temp.ParentBlockHeight
	m.ParentBlockHash = temp.ParentBlockHash
	m.ParentBurnBlockHash = temp.ParentBurnBlockHash
	m.ParentBurnBlockTime = temp.ParentBurnBlockTime
	m.ParentBurnBlockTimeIso = temp.ParentBurnBlockTimeIso
	m.ParentBurnBlockHeight = temp.ParentBurnBlockHeight
	m.BlockHash = temp.BlockHash
	m.Txs = temp.Txs
	return nil
}
