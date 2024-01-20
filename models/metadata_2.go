package models

import (
	"encoding/json"
)

// Metadata2 represents a Metadata2 struct.
// Account-based blockchains that utilize a nonce or sequence number should include that number in the metadata. This number could be unique to the identifier or global across the account address.
type Metadata2 struct {
	SequenceNumber int `json:"sequence_number"`
}

// MarshalJSON implements the json.Marshaler interface for Metadata2.
// It customizes the JSON marshaling process for Metadata2 objects.
func (m *Metadata2) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(m.toMap())
}

// toMap converts the Metadata2 object to a map representation for JSON marshaling.
func (m *Metadata2) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["sequence_number"] = m.SequenceNumber
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Metadata2.
// It customizes the JSON unmarshaling process for Metadata2 objects.
func (m *Metadata2) UnmarshalJSON(input []byte) error {
	temp := &struct {
		SequenceNumber int `json:"sequence_number"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	m.SequenceNumber = temp.SequenceNumber
	return nil
}
