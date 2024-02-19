package models

import (
	"encoding/json"
)

// TransactionResults represents a TransactionResults struct.
// GET request that returns transactions
type TransactionResults struct {
	// The number of transactions to return
	Limit int `json:"limit"`
	// The number to transactions to skip (starting at `0`)
	Offset int `json:"offset"`
	// The number of transactions available
	Total   int           `json:"total"`
	Results []Transaction `json:"results"`
}

// MarshalJSON implements the json.Marshaler interface for TransactionResults.
// It customizes the JSON marshaling process for TransactionResults objects.
func (t *TransactionResults) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(t.toMap())
}

// toMap converts the TransactionResults object to a map representation for JSON marshaling.
func (t *TransactionResults) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["limit"] = t.Limit
	structMap["offset"] = t.Offset
	structMap["total"] = t.Total
	structMap["results"] = t.Results
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for TransactionResults.
// It customizes the JSON unmarshaling process for TransactionResults objects.
func (t *TransactionResults) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
	}

	temp := &struct {
		Limit   int           `json:"limit"`
		Offset  int           `json:"offset"`
		Total   int           `json:"total"`
		Results []Transaction `json:"results"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	t.Limit = temp.Limit
	t.Offset = temp.Offset
	t.Total = temp.Total
	t.Results = temp.Results
	return nil
}
