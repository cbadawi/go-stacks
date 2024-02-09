package models

import (
	"encoding/json"
)

// AddressTransactionsListResponse represents a AddressTransactionsListResponse struct.
// GET request that returns account transactions
type AddressTransactionsListResponse struct {
	Limit   int      `json:"limit"`
	Offset  int      `json:"offset"`
	Total   int      `json:"total"`
	Results []Result `json:"results"`
}

// MarshalJSON implements the json.Marshaler interface for AddressTransactionsListResponse.
// It customizes the JSON marshaling process for AddressTransactionsListResponse objects.
func (a *AddressTransactionsListResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(a.toMap())
}

// toMap converts the AddressTransactionsListResponse object to a map representation for JSON marshaling.
func (a *AddressTransactionsListResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["limit"] = a.Limit
	structMap["offset"] = a.Offset
	structMap["total"] = a.Total
	structMap["results"] = a.Results
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AddressTransactionsListResponse.
// It customizes the JSON unmarshaling process for AddressTransactionsListResponse objects.
func (a *AddressTransactionsListResponse) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		Limit   int      `json:"limit"`
		Offset  int      `json:"offset"`
		Total   int      `json:"total"`
		Results []Result `json:"results"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	a.Limit = temp.Limit
	a.Offset = temp.Offset
	a.Total = temp.Total
	a.Results = temp.Results
	return nil
}
