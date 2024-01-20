package models

import (
	"encoding/json"
)

// BnsGetNamePriceResponse represents a BnsGetNamePriceResponse struct.
// Fetch price for name.
type BnsGetNamePriceResponse struct {
	Units  string `json:"units"`
	Amount string `json:"amount"`
}

// MarshalJSON implements the json.Marshaler interface for BnsGetNamePriceResponse.
// It customizes the JSON marshaling process for BnsGetNamePriceResponse objects.
func (b *BnsGetNamePriceResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(b.toMap())
}

// toMap converts the BnsGetNamePriceResponse object to a map representation for JSON marshaling.
func (b *BnsGetNamePriceResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["units"] = b.Units
	structMap["amount"] = b.Amount
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for BnsGetNamePriceResponse.
// It customizes the JSON unmarshaling process for BnsGetNamePriceResponse objects.
func (b *BnsGetNamePriceResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Units  string `json:"units"`
		Amount string `json:"amount"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	b.Units = temp.Units
	b.Amount = temp.Amount
	return nil
}
