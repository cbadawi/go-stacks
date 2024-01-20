package models

import (
	"encoding/json"
)

// ContractCall2 represents a ContractCall2 struct.
type ContractCall2 struct {
	P25 *float64 `json:"p25"`
	P50 *float64 `json:"p50"`
	P75 *float64 `json:"p75"`
	P95 *float64 `json:"p95"`
}

// MarshalJSON implements the json.Marshaler interface for ContractCall2.
// It customizes the JSON marshaling process for ContractCall2 objects.
func (c *ContractCall2) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the ContractCall2 object to a map representation for JSON marshaling.
func (c *ContractCall2) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["p25"] = c.P25
	structMap["p50"] = c.P50
	structMap["p75"] = c.P75
	structMap["p95"] = c.P95
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ContractCall2.
// It customizes the JSON unmarshaling process for ContractCall2 objects.
func (c *ContractCall2) UnmarshalJSON(input []byte) error {
	temp := &struct {
		P25 *float64 `json:"p25"`
		P50 *float64 `json:"p50"`
		P75 *float64 `json:"p75"`
		P95 *float64 `json:"p95"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.P25 = temp.P25
	c.P50 = temp.P50
	c.P75 = temp.P75
	c.P95 = temp.P95
	return nil
}
