package models

import (
	"encoding/json"
)

// SmartContract2 represents a SmartContract2 struct.
type SmartContract2 struct {
	P25 *float64 `json:"p25"`
	P50 *float64 `json:"p50"`
	P75 *float64 `json:"p75"`
	P95 *float64 `json:"p95"`
}

// MarshalJSON implements the json.Marshaler interface for SmartContract2.
// It customizes the JSON marshaling process for SmartContract2 objects.
func (s *SmartContract2) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the SmartContract2 object to a map representation for JSON marshaling.
func (s *SmartContract2) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["p25"] = s.P25
	structMap["p50"] = s.P50
	structMap["p75"] = s.P75
	structMap["p95"] = s.P95
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SmartContract2.
// It customizes the JSON unmarshaling process for SmartContract2 objects.
func (s *SmartContract2) UnmarshalJSON(input []byte) error {
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

	s.P25 = temp.P25
	s.P50 = temp.P50
	s.P75 = temp.P75
	s.P95 = temp.P95
	return nil
}
