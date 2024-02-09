package models

import (
	"encoding/json"
)

// CoinIdentifier represents a CoinIdentifier struct.
// CoinIdentifier uniquely identifies a Coin.
type CoinIdentifier struct {
	// Identifier should be populated with a globally unique identifier of a Coin. In Bitcoin, this identifier would be transaction_hash:index.
	Identifier string `json:"identifier"`
}

// MarshalJSON implements the json.Marshaler interface for CoinIdentifier.
// It customizes the JSON marshaling process for CoinIdentifier objects.
func (c *CoinIdentifier) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CoinIdentifier object to a map representation for JSON marshaling.
func (c *CoinIdentifier) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["identifier"] = c.Identifier
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CoinIdentifier.
// It customizes the JSON unmarshaling process for CoinIdentifier objects.
func (c *CoinIdentifier) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		Identifier string `json:"identifier"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Identifier = temp.Identifier
	return nil
}
