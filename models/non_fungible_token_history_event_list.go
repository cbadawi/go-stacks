package models

import (
	"encoding/json"
)

// NonFungibleTokenHistoryEventList represents a NonFungibleTokenHistoryEventList struct.
// List of Non-Fungible Token history events
type NonFungibleTokenHistoryEventList struct {
	// The number of events to return
	Limit int `json:"limit"`
	// The number to events to skip (starting at `0`)
	Offset int `json:"offset"`
	// The number of events available
	Total   int                            `json:"total"`
	Results []NonFungibleTokenHistoryEvent `json:"results"`
}

// MarshalJSON implements the json.Marshaler interface for NonFungibleTokenHistoryEventList.
// It customizes the JSON marshaling process for NonFungibleTokenHistoryEventList objects.
func (n *NonFungibleTokenHistoryEventList) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(n.toMap())
}

// toMap converts the NonFungibleTokenHistoryEventList object to a map representation for JSON marshaling.
func (n *NonFungibleTokenHistoryEventList) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["limit"] = n.Limit
	structMap["offset"] = n.Offset
	structMap["total"] = n.Total
	structMap["results"] = n.Results
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for NonFungibleTokenHistoryEventList.
// It customizes the JSON unmarshaling process for NonFungibleTokenHistoryEventList objects.
func (n *NonFungibleTokenHistoryEventList) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		Limit   int                            `json:"limit"`
		Offset  int                            `json:"offset"`
		Total   int                            `json:"total"`
		Results []NonFungibleTokenHistoryEvent `json:"results"`
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
