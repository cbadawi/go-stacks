package models

import (
	"encoding/json"
)

// Metadata3 represents a Metadata3 struct.
// Transactions that are related to other transactions (like a cross-shard transaction) should include the tranaction_identifier of these transactions in the metadata.
type Metadata3 struct {
	// STX token transfer memo.
	Memo *string `json:"memo,omitempty"`
	// The Size
	Size *int `json:"size,omitempty"`
	// The locktime
	LockTime *int `json:"lockTime,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for Metadata3.
// It customizes the JSON marshaling process for Metadata3 objects.
func (m *Metadata3) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(m.toMap())
}

// toMap converts the Metadata3 object to a map representation for JSON marshaling.
func (m *Metadata3) toMap() map[string]any {
	structMap := make(map[string]any)
	if m.Memo != nil {
		structMap["memo"] = m.Memo
	}
	if m.Size != nil {
		structMap["size"] = m.Size
	}
	if m.LockTime != nil {
		structMap["lockTime"] = m.LockTime
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Metadata3.
// It customizes the JSON unmarshaling process for Metadata3 objects.
func (m *Metadata3) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Memo     *string `json:"memo,omitempty"`
		Size     *int    `json:"size,omitempty"`
		LockTime *int    `json:"lockTime,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	m.Memo = temp.Memo
	m.Size = temp.Size
	m.LockTime = temp.LockTime
	return nil
}
