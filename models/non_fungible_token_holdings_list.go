package models

import (
	"encoding/json"
)

// NonFungibleTokenHoldingsList represents a NonFungibleTokenHoldingsList struct.
// List of Non-Fungible Token holdings
type NonFungibleTokenHoldingsList struct {
	// The number of Non-Fungible Token holdings to return
	Limit int `json:"limit"`
	// The number to Non-Fungible Token holdings to skip (starting at `0`)
	Offset int `json:"offset"`
	// The number of Non-Fungible Token holdings available
	Total   int                       `json:"total"`
	Results []NonFungibleTokenHolding `json:"results"`
}

// MarshalJSON implements the json.Marshaler interface for NonFungibleTokenHoldingsList.
// It customizes the JSON marshaling process for NonFungibleTokenHoldingsList objects.
func (n *NonFungibleTokenHoldingsList) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(n.toMap())
}

// toMap converts the NonFungibleTokenHoldingsList object to a map representation for JSON marshaling.
func (n *NonFungibleTokenHoldingsList) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["limit"] = n.Limit
	structMap["offset"] = n.Offset
	structMap["total"] = n.Total
	structMap["results"] = n.Results
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for NonFungibleTokenHoldingsList.
// It customizes the JSON unmarshaling process for NonFungibleTokenHoldingsList objects.
func (n *NonFungibleTokenHoldingsList) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Limit   int                       `json:"limit"`
		Offset  int                       `json:"offset"`
		Total   int                       `json:"total"`
		Results []NonFungibleTokenHolding `json:"results"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	n.Limit = temp.Limit
	n.Offset = temp.Offset
	n.Total = temp.Total
	n.Results = temp.Results
	return nil
}
