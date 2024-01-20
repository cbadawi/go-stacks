package models

import (
	"encoding/json"
)

// TokenTransfer2 represents a TokenTransfer2 struct.
type TokenTransfer2 struct {
	P25 *float64 `json:"p25"`
	P50 *float64 `json:"p50"`
	P75 *float64 `json:"p75"`
	P95 *float64 `json:"p95"`
}

// MarshalJSON implements the json.Marshaler interface for TokenTransfer2.
// It customizes the JSON marshaling process for TokenTransfer2 objects.
func (t *TokenTransfer2) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(t.toMap())
}

// toMap converts the TokenTransfer2 object to a map representation for JSON marshaling.
func (t *TokenTransfer2) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["p25"] = t.P25
	structMap["p50"] = t.P50
	structMap["p75"] = t.P75
	structMap["p95"] = t.P95
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for TokenTransfer2.
// It customizes the JSON unmarshaling process for TokenTransfer2 objects.
func (t *TokenTransfer2) UnmarshalJSON(input []byte) error {
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

	t.P25 = temp.P25
	t.P50 = temp.P50
	t.P75 = temp.P75
	t.P95 = temp.P95
	return nil
}
