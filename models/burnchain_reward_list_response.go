package models

import (
	"encoding/json"
)

// BurnchainRewardListResponse represents a BurnchainRewardListResponse struct.
// GET request that returns blocks
type BurnchainRewardListResponse struct {
	// The number of burnchain rewards to return
	Limit int `json:"limit"`
	// The number to burnchain rewards to skip (starting at `0`)
	Offset  int               `json:"offset"`
	Results []BurnchainReward `json:"results"`
}

// MarshalJSON implements the json.Marshaler interface for BurnchainRewardListResponse.
// It customizes the JSON marshaling process for BurnchainRewardListResponse objects.
func (b *BurnchainRewardListResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(b.toMap())
}

// toMap converts the BurnchainRewardListResponse object to a map representation for JSON marshaling.
func (b *BurnchainRewardListResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["limit"] = b.Limit
	structMap["offset"] = b.Offset
	structMap["results"] = b.Results
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for BurnchainRewardListResponse.
// It customizes the JSON unmarshaling process for BurnchainRewardListResponse objects.
func (b *BurnchainRewardListResponse) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		Limit   int               `json:"limit"`
		Offset  int               `json:"offset"`
		Results []BurnchainReward `json:"results"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	b.Limit = temp.Limit
	b.Offset = temp.Offset
	b.Results = temp.Results
	return nil
}
