package models

import (
	"encoding/json"
)

// MapEntryResponse represents a MapEntryResponse struct.
// Response of get data map entry request
type MapEntryResponse struct {
	// Hex-encoded string of clarity value. It is always an optional tuple.
	Data string `json:"data"`
	// Hex-encoded string of the MARF proof for the data
	Proof *string `json:"proof,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for MapEntryResponse.
// It customizes the JSON marshaling process for MapEntryResponse objects.
func (m *MapEntryResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(m.toMap())
}

// toMap converts the MapEntryResponse object to a map representation for JSON marshaling.
func (m *MapEntryResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["data"] = m.Data
	if m.Proof != nil {
		structMap["proof"] = m.Proof
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MapEntryResponse.
// It customizes the JSON unmarshaling process for MapEntryResponse objects.
func (m *MapEntryResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Data  string  `json:"data"`
		Proof *string `json:"proof,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	m.Data = temp.Data
	m.Proof = temp.Proof
	return nil
}
