package models

import (
	"encoding/json"
)

// UnanchoredTransactionListResponse represents a UnanchoredTransactionListResponse struct.
// GET request that returns unanchored transactions
type UnanchoredTransactionListResponse struct {
	// The number of unanchored transactions available
	Total   int           `json:"total"`
	Results []Transaction `json:"results"`
}

// MarshalJSON implements the json.Marshaler interface for UnanchoredTransactionListResponse.
// It customizes the JSON marshaling process for UnanchoredTransactionListResponse objects.
func (u *UnanchoredTransactionListResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

// toMap converts the UnanchoredTransactionListResponse object to a map representation for JSON marshaling.
func (u *UnanchoredTransactionListResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["total"] = u.Total
	structMap["results"] = u.Results
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UnanchoredTransactionListResponse.
// It customizes the JSON unmarshaling process for UnanchoredTransactionListResponse objects.
func (u *UnanchoredTransactionListResponse) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		Total   int           `json:"total"`
		Results []Transaction `json:"results"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	u.Total = temp.Total
	u.Results = temp.Results
	return nil
}
