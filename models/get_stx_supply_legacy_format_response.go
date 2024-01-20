package models

import (
	"encoding/json"
)

// GetStxSupplyLegacyFormatResponse represents a GetStxSupplyLegacyFormatResponse struct.
// GET request that returns network target block times
type GetStxSupplyLegacyFormatResponse struct {
	// String quoted decimal number of the percentage of STX that have unlocked
	UnlockedPercent string `json:"unlockedPercent"`
	// String quoted decimal number of the total possible number of STX
	TotalStacks string `json:"totalStacks"`
	// Same as `totalStacks` but formatted with comma thousands separators
	TotalStacksFormatted string `json:"totalStacksFormatted"`
	// String quoted decimal number of the STX that have been mined or unlocked
	UnlockedSupply string `json:"unlockedSupply"`
	// Same as `unlockedSupply` but formatted with comma thousands separators
	UnlockedSupplyFormatted string `json:"unlockedSupplyFormatted"`
	// The block height at which this information was queried
	BlockHeight string `json:"blockHeight"`
}

// MarshalJSON implements the json.Marshaler interface for GetStxSupplyLegacyFormatResponse.
// It customizes the JSON marshaling process for GetStxSupplyLegacyFormatResponse objects.
func (g *GetStxSupplyLegacyFormatResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

// toMap converts the GetStxSupplyLegacyFormatResponse object to a map representation for JSON marshaling.
func (g *GetStxSupplyLegacyFormatResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["unlockedPercent"] = g.UnlockedPercent
	structMap["totalStacks"] = g.TotalStacks
	structMap["totalStacksFormatted"] = g.TotalStacksFormatted
	structMap["unlockedSupply"] = g.UnlockedSupply
	structMap["unlockedSupplyFormatted"] = g.UnlockedSupplyFormatted
	structMap["blockHeight"] = g.BlockHeight
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetStxSupplyLegacyFormatResponse.
// It customizes the JSON unmarshaling process for GetStxSupplyLegacyFormatResponse objects.
func (g *GetStxSupplyLegacyFormatResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		UnlockedPercent         string `json:"unlockedPercent"`
		TotalStacks             string `json:"totalStacks"`
		TotalStacksFormatted    string `json:"totalStacksFormatted"`
		UnlockedSupply          string `json:"unlockedSupply"`
		UnlockedSupplyFormatted string `json:"unlockedSupplyFormatted"`
		BlockHeight             string `json:"blockHeight"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.UnlockedPercent = temp.UnlockedPercent
	g.TotalStacks = temp.TotalStacks
	g.TotalStacksFormatted = temp.TotalStacksFormatted
	g.UnlockedSupply = temp.UnlockedSupply
	g.UnlockedSupplyFormatted = temp.UnlockedSupplyFormatted
	g.BlockHeight = temp.BlockHeight
	return nil
}
