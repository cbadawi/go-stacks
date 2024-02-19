package models

import (
	"encoding/json"
)

// Metadata4 represents a Metadata4 struct.
// meta data
type Metadata4 struct {
	TransactionsRoot string `json:"transactions_root"`
	Difficulty       string `json:"difficulty"`
}

// MarshalJSON implements the json.Marshaler interface for Metadata4.
// It customizes the JSON marshaling process for Metadata4 objects.
func (m *Metadata4) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(m.toMap())
}

// toMap converts the Metadata4 object to a map representation for JSON marshaling.
func (m *Metadata4) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["transactions_root"] = m.TransactionsRoot
	structMap["difficulty"] = m.Difficulty
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Metadata4.
// It customizes the JSON unmarshaling process for Metadata4 objects.
func (m *Metadata4) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		TransactionsRoot string `json:"transactions_root"`
		Difficulty       string `json:"difficulty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	m.TransactionsRoot = temp.TransactionsRoot
	m.Difficulty = temp.Difficulty
	return nil
}
