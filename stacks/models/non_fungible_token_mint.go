package models

import (
	"encoding/json"
)

// NonFungibleTokenMint represents a NonFungibleTokenMint struct.
// Describes the minting of a Non-Fungible Token
type NonFungibleTokenMint struct {
	Recipient  *string `json:"recipient,omitempty"`
	EventIndex *int    `json:"event_index,omitempty"`
	// Non-Fungible Token value
	Value *Value18 `json:"value,omitempty"`
	TxId  *string  `json:"tx_id,omitempty"`
	// Describes all transaction types on Stacks 2.0 blockchain
	Tx *Transaction `json:"tx,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for NonFungibleTokenMint.
// It customizes the JSON marshaling process for NonFungibleTokenMint objects.
func (n *NonFungibleTokenMint) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(n.toMap())
}

// toMap converts the NonFungibleTokenMint object to a map representation for JSON marshaling.
func (n *NonFungibleTokenMint) toMap() map[string]any {
	structMap := make(map[string]any)
	if n.Recipient != nil {
		structMap["recipient"] = n.Recipient
	}
	if n.EventIndex != nil {
		structMap["event_index"] = n.EventIndex
	}
	if n.Value != nil {
		structMap["value"] = n.Value
	}
	if n.TxId != nil {
		structMap["tx_id"] = n.TxId
	}
	if n.Tx != nil {
		structMap["tx"] = n.Tx
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for NonFungibleTokenMint.
// It customizes the JSON unmarshaling process for NonFungibleTokenMint objects.
func (n *NonFungibleTokenMint) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		Recipient  *string      `json:"recipient,omitempty"`
		EventIndex *int         `json:"event_index,omitempty"`
		Value      *Value18     `json:"value,omitempty"`
		TxId       *string      `json:"tx_id,omitempty"`
		Tx         *Transaction `json:"tx,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	n.Recipient = temp.Recipient
	n.EventIndex = temp.EventIndex
	n.Value = temp.Value
	n.TxId = temp.TxId
	n.Tx = temp.Tx
	return nil
}
