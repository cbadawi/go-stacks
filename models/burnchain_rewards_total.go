package models

import (
	"encoding/json"
)

// BurnchainRewardsTotal represents a BurnchainRewardsTotal struct.
// Total burnchain rewards made to a recipient
type BurnchainRewardsTotal struct {
	// The recipient address that received the burnchain rewards, in the format native to the burnchain (e.g. B58 encoded for Bitcoin)
	RewardRecipient string `json:"reward_recipient"`
	// The total amount of burnchain tokens rewarded to the recipient, in the smallest unit (e.g. satoshis for Bitcoin)
	RewardAmount string `json:"reward_amount"`
}

// MarshalJSON implements the json.Marshaler interface for BurnchainRewardsTotal.
// It customizes the JSON marshaling process for BurnchainRewardsTotal objects.
func (b *BurnchainRewardsTotal) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(b.toMap())
}

// toMap converts the BurnchainRewardsTotal object to a map representation for JSON marshaling.
func (b *BurnchainRewardsTotal) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["reward_recipient"] = b.RewardRecipient
	structMap["reward_amount"] = b.RewardAmount
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for BurnchainRewardsTotal.
// It customizes the JSON unmarshaling process for BurnchainRewardsTotal objects.
func (b *BurnchainRewardsTotal) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		RewardRecipient string `json:"reward_recipient"`
		RewardAmount    string `json:"reward_amount"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	b.RewardRecipient = temp.RewardRecipient
	b.RewardAmount = temp.RewardAmount
	return nil
}
