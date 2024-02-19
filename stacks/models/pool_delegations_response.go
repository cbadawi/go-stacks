package models

import (
	"encoding/json"
)

// PoolDelegationsResponse represents a PoolDelegationsResponse struct.
// GET request that returns stacking pool member details for a given pool (delegator) principal
type PoolDelegationsResponse struct {
	// The number of Stackers to return
	Limit int `json:"limit"`
	// The number to Stackers to skip (starting at `0`)
	Offset int `json:"offset"`
	// The total number of Stackers
	Total   int              `json:"total"`
	Results []PoolDelegation `json:"results"`
}

// MarshalJSON implements the json.Marshaler interface for PoolDelegationsResponse.
// It customizes the JSON marshaling process for PoolDelegationsResponse objects.
func (p *PoolDelegationsResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(p.toMap())
}

// toMap converts the PoolDelegationsResponse object to a map representation for JSON marshaling.
func (p *PoolDelegationsResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["limit"] = p.Limit
	structMap["offset"] = p.Offset
	structMap["total"] = p.Total
	structMap["results"] = p.Results
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PoolDelegationsResponse.
// It customizes the JSON unmarshaling process for PoolDelegationsResponse objects.
func (p *PoolDelegationsResponse) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		Limit   int              `json:"limit"`
		Offset  int              `json:"offset"`
		Total   int              `json:"total"`
		Results []PoolDelegation `json:"results"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	p.Limit = temp.Limit
	p.Offset = temp.Offset
	p.Total = temp.Total
	p.Results = temp.Results
	return nil
}
