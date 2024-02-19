package models

import (
	"encoding/json"
)

// TxResult represents a TxResult struct.
// Result of the transaction. For contract calls, this will show the value returned by the call. For other transaction types, this will return a boolean indicating the success of the transaction.
type TxResult struct {
	// Hex string representing the value fo the transaction result
	Hex string `json:"hex"`
	// Readable string of the transaction result
	Repr string `json:"repr"`
}

// MarshalJSON implements the json.Marshaler interface for TxResult.
// It customizes the JSON marshaling process for TxResult objects.
func (t *TxResult) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(t.toMap())
}

// toMap converts the TxResult object to a map representation for JSON marshaling.
func (t *TxResult) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["hex"] = t.Hex
	structMap["repr"] = t.Repr
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for TxResult.
// It customizes the JSON unmarshaling process for TxResult objects.
func (t *TxResult) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		Hex  string `json:"hex"`
		Repr string `json:"repr"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	t.Hex = temp.Hex
	t.Repr = temp.Repr
	return nil
}
