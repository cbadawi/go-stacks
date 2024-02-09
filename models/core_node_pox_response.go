package models

import (
	"encoding/json"
)

// CoreNodePoxResponse represents a CoreNodePoxResponse struct.
// Get Proof of Transfer (PoX) information
type CoreNodePoxResponse struct {
	ContractId                 string `json:"contract_id"`
	FirstBurnchainBlockHeight  int    `json:"first_burnchain_block_height"`
	MinAmountUstx              int    `json:"min_amount_ustx"`
	RegistrationWindowLength   int    `json:"registration_window_length"`
	RejectionFraction          int    `json:"rejection_fraction"`
	RewardCycleId              int    `json:"reward_cycle_id"`
	RewardCycleLength          int    `json:"reward_cycle_length"`
	RejectionVotesLeftRequired int    `json:"rejection_votes_left_required"`
	TotalLiquidSupplyUstx      int    `json:"total_liquid_supply_ustx"`
}

// MarshalJSON implements the json.Marshaler interface for CoreNodePoxResponse.
// It customizes the JSON marshaling process for CoreNodePoxResponse objects.
func (c *CoreNodePoxResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CoreNodePoxResponse object to a map representation for JSON marshaling.
func (c *CoreNodePoxResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["contract_id"] = c.ContractId
	structMap["first_burnchain_block_height"] = c.FirstBurnchainBlockHeight
	structMap["min_amount_ustx"] = c.MinAmountUstx
	structMap["registration_window_length"] = c.RegistrationWindowLength
	structMap["rejection_fraction"] = c.RejectionFraction
	structMap["reward_cycle_id"] = c.RewardCycleId
	structMap["reward_cycle_length"] = c.RewardCycleLength
	structMap["rejection_votes_left_required"] = c.RejectionVotesLeftRequired
	structMap["total_liquid_supply_ustx"] = c.TotalLiquidSupplyUstx
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CoreNodePoxResponse.
// It customizes the JSON unmarshaling process for CoreNodePoxResponse objects.
func (c *CoreNodePoxResponse) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		ContractId                 string `json:"contract_id"`
		FirstBurnchainBlockHeight  int    `json:"first_burnchain_block_height"`
		MinAmountUstx              int    `json:"min_amount_ustx"`
		RegistrationWindowLength   int    `json:"registration_window_length"`
		RejectionFraction          int    `json:"rejection_fraction"`
		RewardCycleId              int    `json:"reward_cycle_id"`
		RewardCycleLength          int    `json:"reward_cycle_length"`
		RejectionVotesLeftRequired int    `json:"rejection_votes_left_required"`
		TotalLiquidSupplyUstx      int    `json:"total_liquid_supply_ustx"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.ContractId = temp.ContractId
	c.FirstBurnchainBlockHeight = temp.FirstBurnchainBlockHeight
	c.MinAmountUstx = temp.MinAmountUstx
	c.RegistrationWindowLength = temp.RegistrationWindowLength
	c.RejectionFraction = temp.RejectionFraction
	c.RewardCycleId = temp.RewardCycleId
	c.RewardCycleLength = temp.RewardCycleLength
	c.RejectionVotesLeftRequired = temp.RejectionVotesLeftRequired
	c.TotalLiquidSupplyUstx = temp.TotalLiquidSupplyUstx
	return nil
}
