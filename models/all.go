package models

import (
	"encoding/json"
)

// All represents a All struct.
type All struct {
	NoPriority     int `json:"no_priority"`
	LowPriority    int `json:"low_priority"`
	MediumPriority int `json:"medium_priority"`
	HighPriority   int `json:"high_priority"`
}

// MarshalJSON implements the json.Marshaler interface for All.
// It customizes the JSON marshaling process for All objects.
func (a *All) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(a.toMap())
}

// toMap converts the All object to a map representation for JSON marshaling.
func (a *All) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["no_priority"] = a.NoPriority
	structMap["low_priority"] = a.LowPriority
	structMap["medium_priority"] = a.MediumPriority
	structMap["high_priority"] = a.HighPriority
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for All.
// It customizes the JSON unmarshaling process for All objects.
func (a *All) UnmarshalJSON(input []byte) error {
	temp := &struct {
		NoPriority     int `json:"no_priority"`
		LowPriority    int `json:"low_priority"`
		MediumPriority int `json:"medium_priority"`
		HighPriority   int `json:"high_priority"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	a.NoPriority = temp.NoPriority
	a.LowPriority = temp.LowPriority
	a.MediumPriority = temp.MediumPriority
	a.HighPriority = temp.HighPriority
	return nil
}
