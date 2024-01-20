package models

import (
	"encoding/json"
)

// RosettaCurrency represents a RosettaCurrency struct.
// Currency is composed of a canonical Symbol and Decimals. This Decimals value is used to convert an Amount.Value from atomic units (Satoshis) to standard units (Bitcoins).
type RosettaCurrency struct {
	// Canonical symbol associated with a currency.
	Symbol string `json:"symbol"`
	// Number of decimal places in the standard unit representation of the amount. For example, BTC has 8 decimals. Note that it is not possible to represent the value of some currency in atomic units that is not base 10.
	Decimals int `json:"decimals"`
	// Any additional information related to the currency itself. For example, it would be useful to populate this object with the contract address of an ERC-20 token.
	Metadata *interface{} `json:"metadata,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for RosettaCurrency.
// It customizes the JSON marshaling process for RosettaCurrency objects.
func (r *RosettaCurrency) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the RosettaCurrency object to a map representation for JSON marshaling.
func (r *RosettaCurrency) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["symbol"] = r.Symbol
	structMap["decimals"] = r.Decimals
	if r.Metadata != nil {
		structMap["metadata"] = r.Metadata
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RosettaCurrency.
// It customizes the JSON unmarshaling process for RosettaCurrency objects.
func (r *RosettaCurrency) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Symbol   string       `json:"symbol"`
		Decimals int          `json:"decimals"`
		Metadata *interface{} `json:"metadata,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	r.Symbol = temp.Symbol
	r.Decimals = temp.Decimals
	r.Metadata = temp.Metadata
	return nil
}
