package models

import (
	"encoding/json"
)

// ContractCall1 represents a ContractCall1 struct.
type ContractCall1 struct {
	NoPriority     int `json:"no_priority"`
	LowPriority    int `json:"low_priority"`
	MediumPriority int `json:"medium_priority"`
	HighPriority   int `json:"high_priority"`
}

// MarshalJSON implements the json.Marshaler interface for ContractCall1.
// It customizes the JSON marshaling process for ContractCall1 objects.
func (c *ContractCall1) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the ContractCall1 object to a map representation for JSON marshaling.
func (c *ContractCall1) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["no_priority"] = c.NoPriority
	structMap["low_priority"] = c.LowPriority
	structMap["medium_priority"] = c.MediumPriority
	structMap["high_priority"] = c.HighPriority
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ContractCall1.
// It customizes the JSON unmarshaling process for ContractCall1 objects.
func (c *ContractCall1) UnmarshalJSON(input []byte) error {
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

	c.NoPriority = temp.NoPriority
	c.LowPriority = temp.LowPriority
	c.MediumPriority = temp.MediumPriority
	c.HighPriority = temp.HighPriority
	return nil
}
