package models

import (
	"encoding/json"
)

// TokenTransfer represents a TokenTransfer struct.
type TokenTransfer struct {
	RecipientAddress string `json:"recipient_address"`
	// Transfer amount as Integer string (64-bit unsigned integer)
	Amount string `json:"amount"`
	// Hex encoded arbitrary message, up to 34 bytes length (should try decoding to an ASCII string)
	Memo string `json:"memo"`
}

// MarshalJSON implements the json.Marshaler interface for TokenTransfer.
// It customizes the JSON marshaling process for TokenTransfer objects.
func (t *TokenTransfer) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(t.toMap())
}

// toMap converts the TokenTransfer object to a map representation for JSON marshaling.
func (t *TokenTransfer) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["recipient_address"] = t.RecipientAddress
	structMap["amount"] = t.Amount
	structMap["memo"] = t.Memo
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for TokenTransfer.
// It customizes the JSON unmarshaling process for TokenTransfer objects.
func (t *TokenTransfer) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		RecipientAddress string `json:"recipient_address"`
		Amount           string `json:"amount"`
		Memo             string `json:"memo"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	t.RecipientAddress = temp.RecipientAddress
	t.Amount = temp.Amount
	t.Memo = temp.Memo
	return nil
}
