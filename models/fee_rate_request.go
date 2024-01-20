package models

import (
	"encoding/json"
)

// FeeRateRequest represents a FeeRateRequest struct.
// Request to fetch fee for a transaction
type FeeRateRequest struct {
	// A serialized transaction
	Transaction string `json:"transaction"`
}

// MarshalJSON implements the json.Marshaler interface for FeeRateRequest.
// It customizes the JSON marshaling process for FeeRateRequest objects.
func (f *FeeRateRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(f.toMap())
}

// toMap converts the FeeRateRequest object to a map representation for JSON marshaling.
func (f *FeeRateRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["transaction"] = f.Transaction
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for FeeRateRequest.
// It customizes the JSON unmarshaling process for FeeRateRequest objects.
func (f *FeeRateRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Transaction string `json:"transaction"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	f.Transaction = temp.Transaction
	return nil
}
