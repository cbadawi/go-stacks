package models

import (
	"encoding/json"
)

// BurnchainRewardSlotHolder represents a BurnchainRewardSlotHolder struct.
// Reward slot holder on the burnchain
type BurnchainRewardSlotHolder struct {
	// Set to `true` if block corresponds to the canonical burchchain tip
	Canonical bool `json:"canonical"`
	// The hash representing the burnchain block
	BurnBlockHash string `json:"burn_block_hash"`
	// Height of the burnchain block
	BurnBlockHeight int `json:"burn_block_height"`
	// The recipient address that validly received PoX commitments, in the format native to the burnchain (e.g. B58 encoded for Bitcoin)
	Address string `json:"address"`
	// The index position of the reward entry, useful for ordering when there's more than one slot per burnchain block
	SlotIndex int `json:"slot_index"`
}

// MarshalJSON implements the json.Marshaler interface for BurnchainRewardSlotHolder.
// It customizes the JSON marshaling process for BurnchainRewardSlotHolder objects.
func (b *BurnchainRewardSlotHolder) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(b.toMap())
}

// toMap converts the BurnchainRewardSlotHolder object to a map representation for JSON marshaling.
func (b *BurnchainRewardSlotHolder) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["canonical"] = b.Canonical
	structMap["burn_block_hash"] = b.BurnBlockHash
	structMap["burn_block_height"] = b.BurnBlockHeight
	structMap["address"] = b.Address
	structMap["slot_index"] = b.SlotIndex
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for BurnchainRewardSlotHolder.
// It customizes the JSON unmarshaling process for BurnchainRewardSlotHolder objects.
func (b *BurnchainRewardSlotHolder) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Canonical       bool   `json:"canonical"`
		BurnBlockHash   string `json:"burn_block_hash"`
		BurnBlockHeight int    `json:"burn_block_height"`
		Address         string `json:"address"`
		SlotIndex       int    `json:"slot_index"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	b.Canonical = temp.Canonical
	b.BurnBlockHash = temp.BurnBlockHash
	b.BurnBlockHeight = temp.BurnBlockHeight
	b.Address = temp.Address
	b.SlotIndex = temp.SlotIndex
	return nil
}
