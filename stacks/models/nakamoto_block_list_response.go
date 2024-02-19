package models

import (
	"encoding/json"
)

// NakamotoBlockListResponse represents a NakamotoBlockListResponse struct.
// GET request that returns blocks
type NakamotoBlockListResponse struct {
	// The number of blocks to return
	Limit int `json:"limit"`
	// The number to blocks to skip (starting at `0`)
	Offset int `json:"offset"`
	// The number of blocks available
	Total   int             `json:"total"`
	Results []NakamotoBlock `json:"results"`
}

// MarshalJSON implements the json.Marshaler interface for NakamotoBlockListResponse.
// It customizes the JSON marshaling process for NakamotoBlockListResponse objects.
func (n *NakamotoBlockListResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(n.toMap())
}

// toMap converts the NakamotoBlockListResponse object to a map representation for JSON marshaling.
func (n *NakamotoBlockListResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["limit"] = n.Limit
	structMap["offset"] = n.Offset
	structMap["total"] = n.Total
	structMap["results"] = n.Results
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for NakamotoBlockListResponse.
// It customizes the JSON unmarshaling process for NakamotoBlockListResponse objects.
func (n *NakamotoBlockListResponse) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		Limit   int             `json:"limit"`
		Offset  int             `json:"offset"`
		Total   int             `json:"total"`
		Results []NakamotoBlock `json:"results"`
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
