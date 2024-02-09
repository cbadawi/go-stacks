package models

import (
	"encoding/json"
)

// RosettaAmount represents a RosettaAmount struct.
// Amount is some Value of a Currency. It is considered invalid to specify a Value without a Currency.
type RosettaAmount struct {
	// Value of the transaction in atomic units represented as an arbitrary-sized signed integer. For example, 1 BTC would be represented by a value of 100000000.
	Value string `json:"value"`
	// Currency is composed of a canonical Symbol and Decimals. This Decimals value is used to convert an Amount.Value from atomic units (Satoshis) to standard units (Bitcoins).
	Currency RosettaCurrency `json:"currency"`
	Metadata *interface{}    `json:"metadata,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for RosettaAmount.
// It customizes the JSON marshaling process for RosettaAmount objects.
func (r *RosettaAmount) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the RosettaAmount object to a map representation for JSON marshaling.
func (r *RosettaAmount) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["value"] = r.Value
	structMap["currency"] = r.Currency
	if r.Metadata != nil {
		structMap["metadata"] = r.Metadata
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RosettaAmount.
// It customizes the JSON unmarshaling process for RosettaAmount objects.
func (r *RosettaAmount) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		Value    string          `json:"value"`
		Currency RosettaCurrency `json:"currency"`
		Metadata *interface{}    `json:"metadata,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	r.Value = temp.Value
	r.Currency = temp.Currency
	r.Metadata = temp.Metadata
	return nil
}
