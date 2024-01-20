package models

import (
	"encoding/json"
)

// SmartContract1 represents a SmartContract1 struct.
type SmartContract1 struct {
	NoPriority     int `json:"no_priority"`
	LowPriority    int `json:"low_priority"`
	MediumPriority int `json:"medium_priority"`
	HighPriority   int `json:"high_priority"`
}

// MarshalJSON implements the json.Marshaler interface for SmartContract1.
// It customizes the JSON marshaling process for SmartContract1 objects.
func (s *SmartContract1) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the SmartContract1 object to a map representation for JSON marshaling.
func (s *SmartContract1) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["no_priority"] = s.NoPriority
	structMap["low_priority"] = s.LowPriority
	structMap["medium_priority"] = s.MediumPriority
	structMap["high_priority"] = s.HighPriority
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SmartContract1.
// It customizes the JSON unmarshaling process for SmartContract1 objects.
func (s *SmartContract1) UnmarshalJSON(input []byte) error {
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

	s.NoPriority = temp.NoPriority
	s.LowPriority = temp.LowPriority
	s.MediumPriority = temp.MediumPriority
	s.HighPriority = temp.HighPriority
	return nil
}
