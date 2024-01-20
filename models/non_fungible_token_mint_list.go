package models

import (
	"encoding/json"
)

// NonFungibleTokenMintList represents a NonFungibleTokenMintList struct.
// List of Non-Fungible Token mint events for an asset identifier
type NonFungibleTokenMintList struct {
	// The number of mint events to return
	Limit int `json:"limit"`
	// The number to mint events to skip (starting at `0`)
	Offset int `json:"offset"`
	// The number of mint events available
	Total   int                    `json:"total"`
	Results []NonFungibleTokenMint `json:"results"`
}

// MarshalJSON implements the json.Marshaler interface for NonFungibleTokenMintList.
// It customizes the JSON marshaling process for NonFungibleTokenMintList objects.
func (n *NonFungibleTokenMintList) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(n.toMap())
}

// toMap converts the NonFungibleTokenMintList object to a map representation for JSON marshaling.
func (n *NonFungibleTokenMintList) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["limit"] = n.Limit
	structMap["offset"] = n.Offset
	structMap["total"] = n.Total
	structMap["results"] = n.Results
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for NonFungibleTokenMintList.
// It customizes the JSON unmarshaling process for NonFungibleTokenMintList objects.
func (n *NonFungibleTokenMintList) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Limit   int                    `json:"limit"`
		Offset  int                    `json:"offset"`
		Total   int                    `json:"total"`
		Results []NonFungibleTokenMint `json:"results"`
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
