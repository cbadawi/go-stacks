package models

import (
	"encoding/json"
)

// MicroblockListResponse represents a MicroblockListResponse struct.
// GET request that returns microblocks
type MicroblockListResponse struct {
	// The number of microblocks to return
	Limit int `json:"limit"`
	// The number to microblocks to skip (starting at `0`)
	Offset int `json:"offset"`
	// The number of microblocks available
	Total   int          `json:"total"`
	Results []Microblock `json:"results"`
}

// MarshalJSON implements the json.Marshaler interface for MicroblockListResponse.
// It customizes the JSON marshaling process for MicroblockListResponse objects.
func (m *MicroblockListResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(m.toMap())
}

// toMap converts the MicroblockListResponse object to a map representation for JSON marshaling.
func (m *MicroblockListResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["limit"] = m.Limit
	structMap["offset"] = m.Offset
	structMap["total"] = m.Total
	structMap["results"] = m.Results
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MicroblockListResponse.
// It customizes the JSON unmarshaling process for MicroblockListResponse objects.
func (m *MicroblockListResponse) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		Limit   int          `json:"limit"`
		Offset  int          `json:"offset"`
		Total   int          `json:"total"`
		Results []Microblock `json:"results"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	m.Limit = temp.Limit
	m.Offset = temp.Offset
	m.Total = temp.Total
	m.Results = temp.Results
	return nil
}
