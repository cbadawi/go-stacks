package models

import (
	"encoding/json"
)

// Metadata5 represents a Metadata5 struct.
type Metadata5 struct {
	AccountSequence *int    `json:"account_sequence,omitempty"`
	RecentBlockHash *string `json:"recent_block_hash,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for Metadata5.
// It customizes the JSON marshaling process for Metadata5 objects.
func (m *Metadata5) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(m.toMap())
}

// toMap converts the Metadata5 object to a map representation for JSON marshaling.
func (m *Metadata5) toMap() map[string]any {
	structMap := make(map[string]any)
	if m.AccountSequence != nil {
		structMap["account_sequence"] = m.AccountSequence
	}
	if m.RecentBlockHash != nil {
		structMap["recent_block_hash"] = m.RecentBlockHash
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Metadata5.
// It customizes the JSON unmarshaling process for Metadata5 objects.
func (m *Metadata5) UnmarshalJSON(input []byte) error {
	temp := &struct {
		AccountSequence *int    `json:"account_sequence,omitempty"`
		RecentBlockHash *string `json:"recent_block_hash,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	m.AccountSequence = temp.AccountSequence
	m.RecentBlockHash = temp.RecentBlockHash
	return nil
}
