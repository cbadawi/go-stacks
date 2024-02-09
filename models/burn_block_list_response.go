package models

import (
	"encoding/json"
)

// BurnBlockListResponse represents a BurnBlockListResponse struct.
// GET request that returns burn blocks
type BurnBlockListResponse struct {
	// The number of burn blocks to return
	Limit int `json:"limit"`
	// The number to burn blocks to skip (starting at `0`)
	Offset int `json:"offset"`
	// The number of burn blocks available (regardless of filter parameters)
	Total   int         `json:"total"`
	Results []BurnBlock `json:"results"`
}

// MarshalJSON implements the json.Marshaler interface for BurnBlockListResponse.
// It customizes the JSON marshaling process for BurnBlockListResponse objects.
func (b *BurnBlockListResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(b.toMap())
}

// toMap converts the BurnBlockListResponse object to a map representation for JSON marshaling.
func (b *BurnBlockListResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["limit"] = b.Limit
	structMap["offset"] = b.Offset
	structMap["total"] = b.Total
	structMap["results"] = b.Results
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for BurnBlockListResponse.
// It customizes the JSON unmarshaling process for BurnBlockListResponse objects.
func (b *BurnBlockListResponse) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		Limit   int         `json:"limit"`
		Offset  int         `json:"offset"`
		Total   int         `json:"total"`
		Results []BurnBlock `json:"results"`
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
