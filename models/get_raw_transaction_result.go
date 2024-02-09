package models

import (
	"encoding/json"
)

// GetRawTransactionResult represents a GetRawTransactionResult struct.
// GET raw transaction
type GetRawTransactionResult struct {
	// A hex encoded serialized transaction
	RawTx string `json:"raw_tx"`
}

// MarshalJSON implements the json.Marshaler interface for GetRawTransactionResult.
// It customizes the JSON marshaling process for GetRawTransactionResult objects.
func (g *GetRawTransactionResult) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

// toMap converts the GetRawTransactionResult object to a map representation for JSON marshaling.
func (g *GetRawTransactionResult) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["raw_tx"] = g.RawTx
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetRawTransactionResult.
// It customizes the JSON unmarshaling process for GetRawTransactionResult objects.
func (g *GetRawTransactionResult) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		RawTx string `json:"raw_tx"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.RawTx = temp.RawTx
	return nil
}
