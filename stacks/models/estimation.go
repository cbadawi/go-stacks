package models

import (
	"encoding/json"
)

// Estimation represents a Estimation struct.
type Estimation struct {
	FeeRate *float64 `json:"fee_rate,omitempty"`
	Fee     *float64 `json:"fee,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for Estimation.
// It customizes the JSON marshaling process for Estimation objects.
func (e *Estimation) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(e.toMap())
}

// toMap converts the Estimation object to a map representation for JSON marshaling.
func (e *Estimation) toMap() map[string]any {
	structMap := make(map[string]any)
	if e.FeeRate != nil {
		structMap["fee_rate"] = e.FeeRate
	}
	if e.Fee != nil {
		structMap["fee"] = e.Fee
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Estimation.
// It customizes the JSON unmarshaling process for Estimation objects.
func (e *Estimation) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		FeeRate *float64 `json:"fee_rate,omitempty"`
		Fee     *float64 `json:"fee,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	e.FeeRate = temp.FeeRate
	e.Fee = temp.Fee
	return nil
}
