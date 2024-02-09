package models

import (
	"encoding/json"
)

// NonFungibleTokenHolding represents a NonFungibleTokenHolding struct.
// Describes the ownership of a Non-Fungible Token
type NonFungibleTokenHolding struct {
	AssetIdentifier *string `json:"asset_identifier,omitempty"`
	// Non-Fungible Token value
	Value       *Value18 `json:"value,omitempty"`
	BlockHeight *float64 `json:"block_height,omitempty"`
	TxId        *string  `json:"tx_id,omitempty"`
	// Describes all transaction types on Stacks 2.0 blockchain
	Tx *Transaction `json:"tx,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for NonFungibleTokenHolding.
// It customizes the JSON marshaling process for NonFungibleTokenHolding objects.
func (n *NonFungibleTokenHolding) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(n.toMap())
}

// toMap converts the NonFungibleTokenHolding object to a map representation for JSON marshaling.
func (n *NonFungibleTokenHolding) toMap() map[string]any {
	structMap := make(map[string]any)
	if n.AssetIdentifier != nil {
		structMap["asset_identifier"] = n.AssetIdentifier
	}
	if n.Value != nil {
		structMap["value"] = n.Value
	}
	if n.BlockHeight != nil {
		structMap["block_height"] = n.BlockHeight
	}
	if n.TxId != nil {
		structMap["tx_id"] = n.TxId
	}
	if n.Tx != nil {
		structMap["tx"] = n.Tx
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for NonFungibleTokenHolding.
// It customizes the JSON unmarshaling process for NonFungibleTokenHolding objects.
func (n *NonFungibleTokenHolding) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		AssetIdentifier *string      `json:"asset_identifier,omitempty"`
		Value           *Value18     `json:"value,omitempty"`
		BlockHeight     *float64     `json:"block_height,omitempty"`
		TxId            *string      `json:"tx_id,omitempty"`
		Tx              *Transaction `json:"tx,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	n.AssetIdentifier = temp.AssetIdentifier
	n.Value = temp.Value
	n.BlockHeight = temp.BlockHeight
	n.TxId = temp.TxId
	n.Tx = temp.Tx
	return nil
}
