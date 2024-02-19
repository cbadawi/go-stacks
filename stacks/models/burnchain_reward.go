package models

import (
	"encoding/json"
)

// BurnchainReward represents a BurnchainReward struct.
// Reward payment made on the burnchain
type BurnchainReward struct {
	// Set to `true` if block corresponds to the canonical burchchain tip
	Canonical bool `json:"canonical"`
	// The hash representing the burnchain block
	BurnBlockHash string `json:"burn_block_hash"`
	// Height of the burnchain block
	BurnBlockHeight int `json:"burn_block_height"`
	// The total amount of burnchain tokens burned for this burnchain block, in the smallest unit (e.g. satoshis for Bitcoin)
	BurnAmount string `json:"burn_amount"`
	// The recipient address that received the burnchain rewards, in the format native to the burnchain (e.g. B58 encoded for Bitcoin)
	RewardRecipient string `json:"reward_recipient"`
	// The amount of burnchain tokens rewarded to the recipient, in the smallest unit (e.g. satoshis for Bitcoin)
	RewardAmount string `json:"reward_amount"`
	// The index position of the reward entry, useful for ordering when there's more than one recipient per burnchain block
	RewardIndex int `json:"reward_index"`
}

// MarshalJSON implements the json.Marshaler interface for BurnchainReward.
// It customizes the JSON marshaling process for BurnchainReward objects.
func (b *BurnchainReward) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(b.toMap())
}

// toMap converts the BurnchainReward object to a map representation for JSON marshaling.
func (b *BurnchainReward) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["canonical"] = b.Canonical
	structMap["burn_block_hash"] = b.BurnBlockHash
	structMap["burn_block_height"] = b.BurnBlockHeight
	structMap["burn_amount"] = b.BurnAmount
	structMap["reward_recipient"] = b.RewardRecipient
	structMap["reward_amount"] = b.RewardAmount
	structMap["reward_index"] = b.RewardIndex
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for BurnchainReward.
// It customizes the JSON unmarshaling process for BurnchainReward objects.
func (b *BurnchainReward) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		Canonical       bool   `json:"canonical"`
		BurnBlockHash   string `json:"burn_block_hash"`
		BurnBlockHeight int    `json:"burn_block_height"`
		BurnAmount      string `json:"burn_amount"`
		RewardRecipient string `json:"reward_recipient"`
		RewardAmount    string `json:"reward_amount"`
		RewardIndex     int    `json:"reward_index"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	b.Canonical = temp.Canonical
	b.BurnBlockHash = temp.BurnBlockHash
	b.BurnBlockHeight = temp.BurnBlockHeight
	b.BurnAmount = temp.BurnAmount
	b.RewardRecipient = temp.RewardRecipient
	b.RewardAmount = temp.RewardAmount
	b.RewardIndex = temp.RewardIndex
	return nil
}
