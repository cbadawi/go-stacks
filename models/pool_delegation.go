package models

import (
	"encoding/json"
)

// PoolDelegation represents a PoolDelegation struct.
type PoolDelegation struct {
	// The principal of the pool member that issued the delegation
	Stacker string `json:"stacker"`
	// The pox-addr value specified by the stacker in the delegation operation
	PoxAddr *string `json:"pox_addr,omitempty"`
	// The amount of uSTX delegated by the stacker
	AmountUstx string `json:"amount_ustx"`
	// The optional burnchain block unlock height that the stacker may have specified
	BurnBlockUnlockHeight *int `json:"burn_block_unlock_height,omitempty"`
	// The block height at which the stacker delegation transaction was mined at
	BlockHeight int `json:"block_height"`
	// The tx_id of the stacker delegation operation
	TxId string `json:"tx_id"`
}

// MarshalJSON implements the json.Marshaler interface for PoolDelegation.
// It customizes the JSON marshaling process for PoolDelegation objects.
func (p *PoolDelegation) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(p.toMap())
}

// toMap converts the PoolDelegation object to a map representation for JSON marshaling.
func (p *PoolDelegation) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["stacker"] = p.Stacker
	if p.PoxAddr != nil {
		structMap["pox_addr"] = p.PoxAddr
	}
	structMap["amount_ustx"] = p.AmountUstx
	if p.BurnBlockUnlockHeight != nil {
		structMap["burn_block_unlock_height"] = p.BurnBlockUnlockHeight
	}
	structMap["block_height"] = p.BlockHeight
	structMap["tx_id"] = p.TxId
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PoolDelegation.
// It customizes the JSON unmarshaling process for PoolDelegation objects.
func (p *PoolDelegation) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Stacker               string  `json:"stacker"`
		PoxAddr               *string `json:"pox_addr,omitempty"`
		AmountUstx            string  `json:"amount_ustx"`
		BurnBlockUnlockHeight *int    `json:"burn_block_unlock_height,omitempty"`
		BlockHeight           int     `json:"block_height"`
		TxId                  string  `json:"tx_id"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	p.Stacker = temp.Stacker
	p.PoxAddr = temp.PoxAddr
	p.AmountUstx = temp.AmountUstx
	p.BurnBlockUnlockHeight = temp.BurnBlockUnlockHeight
	p.BlockHeight = temp.BlockHeight
	p.TxId = temp.TxId
	return nil
}
