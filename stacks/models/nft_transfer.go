package models

import (
	"encoding/json"
)

// NftTransfer represents a NftTransfer struct.
type NftTransfer struct {
	// Non Fungible Token asset identifier.
	AssetIdentifier string `json:"asset_identifier"`
	// Non Fungible Token asset value.
	Value Value11 `json:"value"`
	// Principal that sent the asset.
	Sender *string `json:"sender,omitempty"`
	// Principal that received the asset.
	Recipient *string `json:"recipient,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for NftTransfer.
// It customizes the JSON marshaling process for NftTransfer objects.
func (n *NftTransfer) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(n.toMap())
}

// toMap converts the NftTransfer object to a map representation for JSON marshaling.
func (n *NftTransfer) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["asset_identifier"] = n.AssetIdentifier
	structMap["value"] = n.Value
	if n.Sender != nil {
		structMap["sender"] = n.Sender
	}
	if n.Recipient != nil {
		structMap["recipient"] = n.Recipient
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for NftTransfer.
// It customizes the JSON unmarshaling process for NftTransfer objects.
func (n *NftTransfer) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		AssetIdentifier string  `json:"asset_identifier"`
		Value           Value11 `json:"value"`
		Sender          *string `json:"sender,omitempty"`
		Recipient       *string `json:"recipient,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	n.AssetIdentifier = temp.AssetIdentifier
	n.Value = temp.Value
	n.Sender = temp.Sender
	n.Recipient = temp.Recipient
	return nil
}
