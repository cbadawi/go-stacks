package models

import (
	"encoding/json"
)

// TxData3 represents a TxData3 struct.
// Returns basic search result information about the requested id
type TxData3 struct {
	// If the transaction lies within the canonical chain
	Canonical bool `json:"canonical"`
	// Refers to the hash of the block for searched transaction
	BlockHash     string `json:"block_hash"`
	BurnBlockTime int    `json:"burn_block_time"`
	BlockHeight   int    `json:"block_height"`
	TxType        string `json:"tx_type"`
	// Corresponding tx_id for smart_contract
	TxId *string `json:"tx_id,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for TxData3.
// It customizes the JSON marshaling process for TxData3 objects.
func (t *TxData3) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(t.toMap())
}

// toMap converts the TxData3 object to a map representation for JSON marshaling.
func (t *TxData3) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["canonical"] = t.Canonical
	structMap["block_hash"] = t.BlockHash
	structMap["burn_block_time"] = t.BurnBlockTime
	structMap["block_height"] = t.BlockHeight
	structMap["tx_type"] = t.TxType
	if t.TxId != nil {
		structMap["tx_id"] = t.TxId
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for TxData3.
// It customizes the JSON unmarshaling process for TxData3 objects.
func (t *TxData3) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		Canonical     bool    `json:"canonical"`
		BlockHash     string  `json:"block_hash"`
		BurnBlockTime int     `json:"burn_block_time"`
		BlockHeight   int     `json:"block_height"`
		TxType        string  `json:"tx_type"`
		TxId          *string `json:"tx_id,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	t.Canonical = temp.Canonical
	t.BlockHash = temp.BlockHash
	t.BurnBlockTime = temp.BurnBlockTime
	t.BlockHeight = temp.BlockHeight
	t.TxType = temp.TxType
	t.TxId = temp.TxId
	return nil
}
