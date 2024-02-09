package models

import (
	"encoding/json"
)

// FeeRate represents a FeeRate struct.
// Get fee rate information.
type FeeRate struct {
	FeeRate int `json:"fee_rate"`
}

// MarshalJSON implements the json.Marshaler interface for FeeRate.
// It customizes the JSON marshaling process for FeeRate objects.
func (f *FeeRate) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(f.toMap())
}

// toMap converts the FeeRate object to a map representation for JSON marshaling.
func (f *FeeRate) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["fee_rate"] = f.FeeRate
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for FeeRate.
// It customizes the JSON unmarshaling process for FeeRate objects.
func (f *FeeRate) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		FeeRate int `json:"fee_rate"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	f.FeeRate = temp.FeeRate
	return nil
}
