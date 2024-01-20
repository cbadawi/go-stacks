package models

import (
	"encoding/json"
)

// BurnchainRewardSlotHolderListResponse represents a BurnchainRewardSlotHolderListResponse struct.
// GET request that returns reward slot holders
type BurnchainRewardSlotHolderListResponse struct {
	// The number of items to return
	Limit int `json:"limit"`
	// The number of items to skip (starting at `0`)
	Offset int `json:"offset"`
	// Total number of available items
	Total   int                         `json:"total"`
	Results []BurnchainRewardSlotHolder `json:"results"`
}

// MarshalJSON implements the json.Marshaler interface for BurnchainRewardSlotHolderListResponse.
// It customizes the JSON marshaling process for BurnchainRewardSlotHolderListResponse objects.
func (b *BurnchainRewardSlotHolderListResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(b.toMap())
}

// toMap converts the BurnchainRewardSlotHolderListResponse object to a map representation for JSON marshaling.
func (b *BurnchainRewardSlotHolderListResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["limit"] = b.Limit
	structMap["offset"] = b.Offset
	structMap["total"] = b.Total
	structMap["results"] = b.Results
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for BurnchainRewardSlotHolderListResponse.
// It customizes the JSON unmarshaling process for BurnchainRewardSlotHolderListResponse objects.
func (b *BurnchainRewardSlotHolderListResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Limit   int                         `json:"limit"`
		Offset  int                         `json:"offset"`
		Total   int                         `json:"total"`
		Results []BurnchainRewardSlotHolder `json:"results"`
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
