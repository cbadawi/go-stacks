package models

import (
	"encoding/json"
)

// TargetBlockTime represents a TargetBlockTime struct.
type TargetBlockTime struct {
	TargetBlockTime int `json:"target_block_time"`
}

// MarshalJSON implements the json.Marshaler interface for TargetBlockTime.
// It customizes the JSON marshaling process for TargetBlockTime objects.
func (t *TargetBlockTime) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(t.toMap())
}

// toMap converts the TargetBlockTime object to a map representation for JSON marshaling.
func (t *TargetBlockTime) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["target_block_time"] = t.TargetBlockTime
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for TargetBlockTime.
// It customizes the JSON unmarshaling process for TargetBlockTime objects.
func (t *TargetBlockTime) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		TargetBlockTime int `json:"target_block_time"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	t.TargetBlockTime = temp.TargetBlockTime
	return nil
}
