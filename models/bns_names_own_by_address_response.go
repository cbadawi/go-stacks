package models

import (
	"encoding/json"
)

// BnsNamesOwnByAddressResponse represents a BnsNamesOwnByAddressResponse struct.
// Retrieves a list of names owned by the address provided.
type BnsNamesOwnByAddressResponse struct {
	Names []string `json:"names"`
}

// MarshalJSON implements the json.Marshaler interface for BnsNamesOwnByAddressResponse.
// It customizes the JSON marshaling process for BnsNamesOwnByAddressResponse objects.
func (b *BnsNamesOwnByAddressResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(b.toMap())
}

// toMap converts the BnsNamesOwnByAddressResponse object to a map representation for JSON marshaling.
func (b *BnsNamesOwnByAddressResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["names"] = b.Names
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for BnsNamesOwnByAddressResponse.
// It customizes the JSON unmarshaling process for BnsNamesOwnByAddressResponse objects.
func (b *BnsNamesOwnByAddressResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Names []string `json:"names"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	b.Names = temp.Names
	return nil
}
