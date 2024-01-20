package models

import (
	"encoding/json"
)

// FtTransfer represents a FtTransfer struct.
type FtTransfer struct {
	// Fungible Token asset identifier.
	AssetIdentifier string `json:"asset_identifier"`
	// Amount transferred as an integer string. This balance does not factor in possible SIP-010 decimals.
	Amount string `json:"amount"`
	// Principal that sent the asset.
	Sender *string `json:"sender,omitempty"`
	// Principal that received the asset.
	Recipient *string `json:"recipient,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for FtTransfer.
// It customizes the JSON marshaling process for FtTransfer objects.
func (f *FtTransfer) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(f.toMap())
}

// toMap converts the FtTransfer object to a map representation for JSON marshaling.
func (f *FtTransfer) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["asset_identifier"] = f.AssetIdentifier
	structMap["amount"] = f.Amount
	if f.Sender != nil {
		structMap["sender"] = f.Sender
	}
	if f.Recipient != nil {
		structMap["recipient"] = f.Recipient
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for FtTransfer.
// It customizes the JSON unmarshaling process for FtTransfer objects.
func (f *FtTransfer) UnmarshalJSON(input []byte) error {
	temp := &struct {
		AssetIdentifier string  `json:"asset_identifier"`
		Amount          string  `json:"amount"`
		Sender          *string `json:"sender,omitempty"`
		Recipient       *string `json:"recipient,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	f.AssetIdentifier = temp.AssetIdentifier
	f.Amount = temp.Amount
	f.Sender = temp.Sender
	f.Recipient = temp.Recipient
	return nil
}
