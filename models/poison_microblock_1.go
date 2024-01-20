package models

import (
	"encoding/json"
)

// PoisonMicroblock1 represents a PoisonMicroblock1 struct.
type PoisonMicroblock1 struct {
	P25 *float64 `json:"p25"`
	P50 *float64 `json:"p50"`
	P75 *float64 `json:"p75"`
	P95 *float64 `json:"p95"`
}

// MarshalJSON implements the json.Marshaler interface for PoisonMicroblock1.
// It customizes the JSON marshaling process for PoisonMicroblock1 objects.
func (p *PoisonMicroblock1) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(p.toMap())
}

// toMap converts the PoisonMicroblock1 object to a map representation for JSON marshaling.
func (p *PoisonMicroblock1) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["p25"] = p.P25
	structMap["p50"] = p.P50
	structMap["p75"] = p.P75
	structMap["p95"] = p.P95
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PoisonMicroblock1.
// It customizes the JSON unmarshaling process for PoisonMicroblock1 objects.
func (p *PoisonMicroblock1) UnmarshalJSON(input []byte) error {
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

	p.P25 = temp.P25
	p.P50 = temp.P50
	p.P75 = temp.P75
	p.P95 = temp.P95
	return nil
}
