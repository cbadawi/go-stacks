package models

import (
	"encoding/json"
)

// AddressTransactionsWithTransfersListResponse represents a AddressTransactionsWithTransfersListResponse struct.
// GET request that returns account transactions
type AddressTransactionsWithTransfersListResponse struct {
	Limit   int                               `json:"limit"`
	Offset  int                               `json:"offset"`
	Total   int                               `json:"total"`
	Results []AddressTransactionWithTransfers `json:"results"`
}

// MarshalJSON implements the json.Marshaler interface for AddressTransactionsWithTransfersListResponse.
// It customizes the JSON marshaling process for AddressTransactionsWithTransfersListResponse objects.
func (a *AddressTransactionsWithTransfersListResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(a.toMap())
}

// toMap converts the AddressTransactionsWithTransfersListResponse object to a map representation for JSON marshaling.
func (a *AddressTransactionsWithTransfersListResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["limit"] = a.Limit
	structMap["offset"] = a.Offset
	structMap["total"] = a.Total
	structMap["results"] = a.Results
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AddressTransactionsWithTransfersListResponse.
// It customizes the JSON unmarshaling process for AddressTransactionsWithTransfersListResponse objects.
func (a *AddressTransactionsWithTransfersListResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Limit   int                               `json:"limit"`
		Offset  int                               `json:"offset"`
		Total   int                               `json:"total"`
		Results []AddressTransactionWithTransfers `json:"results"`
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
