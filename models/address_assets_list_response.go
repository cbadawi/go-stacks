package models

import (
	"encoding/json"
)

// AddressAssetsListResponse represents a AddressAssetsListResponse struct.
// GET request that returns address assets
type AddressAssetsListResponse struct {
	Limit   int                `json:"limit"`
	Offset  int                `json:"offset"`
	Total   int                `json:"total"`
	Results []TransactionEvent `json:"results"`
}

// MarshalJSON implements the json.Marshaler interface for AddressAssetsListResponse.
// It customizes the JSON marshaling process for AddressAssetsListResponse objects.
func (a *AddressAssetsListResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(a.toMap())
}

// toMap converts the AddressAssetsListResponse object to a map representation for JSON marshaling.
func (a *AddressAssetsListResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["limit"] = a.Limit
	structMap["offset"] = a.Offset
	structMap["total"] = a.Total
	structMap["results"] = a.Results
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AddressAssetsListResponse.
// It customizes the JSON unmarshaling process for AddressAssetsListResponse objects.
func (a *AddressAssetsListResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Limit   int                `json:"limit"`
		Offset  int                `json:"offset"`
		Total   int                `json:"total"`
		Results []TransactionEvent `json:"results"`
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
