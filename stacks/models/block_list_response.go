package models

import (
	"encoding/json"
)

// BlockListResponse represents a BlockListResponse struct.
// GET request that returns blocks
type BlockListResponse struct {
	// The number of blocks to return
	Limit int `json:"limit"`
	// The number to blocks to skip (starting at `0`)
	Offset int `json:"offset"`
	// The number of blocks available
	Total   int     `json:"total"`
	Results []Block `json:"results"`
}

// MarshalJSON implements the json.Marshaler interface for BlockListResponse.
// It customizes the JSON marshaling process for BlockListResponse objects.
func (b *BlockListResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(b.toMap())
}

// toMap converts the BlockListResponse object to a map representation for JSON marshaling.
func (b *BlockListResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["limit"] = b.Limit
	structMap["offset"] = b.Offset
	structMap["total"] = b.Total
	structMap["results"] = b.Results
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for BlockListResponse.
// It customizes the JSON unmarshaling process for BlockListResponse objects.
func (b *BlockListResponse) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		Limit   int     `json:"limit"`
		Offset  int     `json:"offset"`
		Total   int     `json:"total"`
		Results []Block `json:"results"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	b.Limit = temp.Limit
	b.Offset = temp.Offset
	b.Total = temp.Total
	b.Results = temp.Results
	return nil
}
