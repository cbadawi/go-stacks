package models

import (
	"encoding/json"
)

// TokenTransfer1 represents a TokenTransfer1 struct.
type TokenTransfer1 struct {
	NoPriority     int `json:"no_priority"`
	LowPriority    int `json:"low_priority"`
	MediumPriority int `json:"medium_priority"`
	HighPriority   int `json:"high_priority"`
}

// MarshalJSON implements the json.Marshaler interface for TokenTransfer1.
// It customizes the JSON marshaling process for TokenTransfer1 objects.
func (t *TokenTransfer1) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(t.toMap())
}

// toMap converts the TokenTransfer1 object to a map representation for JSON marshaling.
func (t *TokenTransfer1) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["no_priority"] = t.NoPriority
	structMap["low_priority"] = t.LowPriority
	structMap["medium_priority"] = t.MediumPriority
	structMap["high_priority"] = t.HighPriority
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for TokenTransfer1.
// It customizes the JSON unmarshaling process for TokenTransfer1 objects.
func (t *TokenTransfer1) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
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

	t.NoPriority = temp.NoPriority
	t.LowPriority = temp.LowPriority
	t.MediumPriority = temp.MediumPriority
	t.HighPriority = temp.HighPriority
	return nil
}
