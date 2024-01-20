package models

import (
	"encoding/json"
)

// Metadata1 represents a Metadata1 struct.
// Meta data from subnetwork identifier
type Metadata1 struct {
	// producer
	Producer string `json:"producer"`
}

// MarshalJSON implements the json.Marshaler interface for Metadata1.
// It customizes the JSON marshaling process for Metadata1 objects.
func (m *Metadata1) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(m.toMap())
}

// toMap converts the Metadata1 object to a map representation for JSON marshaling.
func (m *Metadata1) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["producer"] = m.Producer
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Metadata1.
// It customizes the JSON unmarshaling process for Metadata1 objects.
func (m *Metadata1) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Producer string `json:"producer"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	m.Producer = temp.Producer
	return nil
}
