package models

import (
	"encoding/json"
)

// NonFungibleTokenHistoryEvent represents a NonFungibleTokenHistoryEvent struct.
// Describes an event from the history of a Non-Fungible Token
type NonFungibleTokenHistoryEvent struct {
	Sender         Optional[string] `json:"sender"`
	Recipient      *string          `json:"recipient,omitempty"`
	EventIndex     *int             `json:"event_index,omitempty"`
	AssetEventType *string          `json:"asset_event_type,omitempty"`
	TxId           *string          `json:"tx_id,omitempty"`
	// Describes all transaction types on Stacks 2.0 blockchain
	Tx *Transaction `json:"tx,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for NonFungibleTokenHistoryEvent.
// It customizes the JSON marshaling process for NonFungibleTokenHistoryEvent objects.
func (n *NonFungibleTokenHistoryEvent) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(n.toMap())
}

// toMap converts the NonFungibleTokenHistoryEvent object to a map representation for JSON marshaling.
func (n *NonFungibleTokenHistoryEvent) toMap() map[string]any {
	structMap := make(map[string]any)
	if n.Sender.IsValueSet() {
		structMap["sender"] = n.Sender.Value()
	}
	if n.Recipient != nil {
		structMap["recipient"] = n.Recipient
	}
	if n.EventIndex != nil {
		structMap["event_index"] = n.EventIndex
	}
	if n.AssetEventType != nil {
		structMap["asset_event_type"] = n.AssetEventType
	}
	if n.TxId != nil {
		structMap["tx_id"] = n.TxId
	}
	if n.Tx != nil {
		structMap["tx"] = n.Tx
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for NonFungibleTokenHistoryEvent.
// It customizes the JSON unmarshaling process for NonFungibleTokenHistoryEvent objects.
func (n *NonFungibleTokenHistoryEvent) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		Sender         Optional[string] `json:"sender"`
		Recipient      *string          `json:"recipient,omitempty"`
		EventIndex     *int             `json:"event_index,omitempty"`
		AssetEventType *string          `json:"asset_event_type,omitempty"`
		TxId           *string          `json:"tx_id,omitempty"`
		Tx             *Transaction     `json:"tx,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	n.Sender = temp.Sender
	n.Recipient = temp.Recipient
	n.EventIndex = temp.EventIndex
	n.AssetEventType = temp.AssetEventType
	n.TxId = temp.TxId
	n.Tx = temp.Tx
	return nil
}
