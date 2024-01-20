package models

import (
	"encoding/json"
)

// EstimatedCost represents a EstimatedCost struct.
type EstimatedCost struct {
	ReadCount   int `json:"read_count"`
	ReadLength  int `json:"read_length"`
	Runtime     int `json:"runtime"`
	WriteCount  int `json:"write_count"`
	WriteLength int `json:"write_length"`
}

// MarshalJSON implements the json.Marshaler interface for EstimatedCost.
// It customizes the JSON marshaling process for EstimatedCost objects.
func (e *EstimatedCost) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(e.toMap())
}

// toMap converts the EstimatedCost object to a map representation for JSON marshaling.
func (e *EstimatedCost) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["read_count"] = e.ReadCount
	structMap["read_length"] = e.ReadLength
	structMap["runtime"] = e.Runtime
	structMap["write_count"] = e.WriteCount
	structMap["write_length"] = e.WriteLength
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for EstimatedCost.
// It customizes the JSON unmarshaling process for EstimatedCost objects.
func (e *EstimatedCost) UnmarshalJSON(input []byte) error {
	temp := &struct {
		ReadCount   int `json:"read_count"`
		ReadLength  int `json:"read_length"`
		Runtime     int `json:"runtime"`
		WriteCount  int `json:"write_count"`
		WriteLength int `json:"write_length"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	e.ReadCount = temp.ReadCount
	e.ReadLength = temp.ReadLength
	e.Runtime = temp.Runtime
	e.WriteCount = temp.WriteCount
	e.WriteLength = temp.WriteLength
	return nil
}
