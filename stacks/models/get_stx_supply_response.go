package models

import (
	"encoding/json"
)

// GetStxSupplyResponse represents a GetStxSupplyResponse struct.
// GET request that returns network target block times
type GetStxSupplyResponse struct {
	// String quoted decimal number of the percentage of STX that have unlocked
	UnlockedPercent string `json:"unlocked_percent"`
	// String quoted decimal number of the total possible number of STX
	TotalStx string `json:"total_stx"`
	// String quoted decimal number of the STX that have been mined or unlocked
	UnlockedStx string `json:"unlocked_stx"`
	// The block height at which this information was queried
	BlockHeight int `json:"block_height"`
}

// MarshalJSON implements the json.Marshaler interface for GetStxSupplyResponse.
// It customizes the JSON marshaling process for GetStxSupplyResponse objects.
func (g *GetStxSupplyResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

// toMap converts the GetStxSupplyResponse object to a map representation for JSON marshaling.
func (g *GetStxSupplyResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["unlocked_percent"] = g.UnlockedPercent
	structMap["total_stx"] = g.TotalStx
	structMap["unlocked_stx"] = g.UnlockedStx
	structMap["block_height"] = g.BlockHeight
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetStxSupplyResponse.
// It customizes the JSON unmarshaling process for GetStxSupplyResponse objects.
func (g *GetStxSupplyResponse) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		UnlockedPercent string `json:"unlocked_percent"`
		TotalStx        string `json:"total_stx"`
		UnlockedStx     string `json:"unlocked_stx"`
		BlockHeight     int    `json:"block_height"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.UnlockedPercent = temp.UnlockedPercent
	g.TotalStx = temp.TotalStx
	g.UnlockedStx = temp.UnlockedStx
	g.BlockHeight = temp.BlockHeight
	return nil
}
