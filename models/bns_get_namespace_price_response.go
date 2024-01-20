package models

import (
	"encoding/json"
)

// BnsGetNamespacePriceResponse represents a BnsGetNamespacePriceResponse struct.
// Fetch price for namespace.
type BnsGetNamespacePriceResponse struct {
	Units  string `json:"units"`
	Amount string `json:"amount"`
}

// MarshalJSON implements the json.Marshaler interface for BnsGetNamespacePriceResponse.
// It customizes the JSON marshaling process for BnsGetNamespacePriceResponse objects.
func (b *BnsGetNamespacePriceResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(b.toMap())
}

// toMap converts the BnsGetNamespacePriceResponse object to a map representation for JSON marshaling.
func (b *BnsGetNamespacePriceResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["units"] = b.Units
	structMap["amount"] = b.Amount
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for BnsGetNamespacePriceResponse.
// It customizes the JSON unmarshaling process for BnsGetNamespacePriceResponse objects.
func (b *BnsGetNamespacePriceResponse) UnmarshalJSON(input []byte) error {
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
