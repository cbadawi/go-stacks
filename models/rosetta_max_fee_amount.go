package models

import (
	"encoding/json"
)

// RosettaMaxFeeAmount represents a RosettaMaxFeeAmount struct.
// Amount is some Value of a Currency. It is considered invalid to specify a Value without a Currency.
type RosettaMaxFeeAmount struct {
	// Value of the transaction in atomic units represented as an arbitrary-sized signed integer. For example, 1 BTC would be represented by a value of 100000000.
	Value string `json:"value"`
	// Currency is composed of a canonical Symbol and Decimals. This Decimals value is used to convert an Amount.Value from atomic units (Satoshis) to standard units (Bitcoins).
	Currency RosettaCurrency `json:"currency"`
	Metadata *interface{}    `json:"metadata,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for RosettaMaxFeeAmount.
// It customizes the JSON marshaling process for RosettaMaxFeeAmount objects.
func (r *RosettaMaxFeeAmount) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the RosettaMaxFeeAmount object to a map representation for JSON marshaling.
func (r *RosettaMaxFeeAmount) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["value"] = r.Value
	structMap["currency"] = r.Currency
	if r.Metadata != nil {
		structMap["metadata"] = r.Metadata
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RosettaMaxFeeAmount.
// It customizes the JSON unmarshaling process for RosettaMaxFeeAmount objects.
func (r *RosettaMaxFeeAmount) UnmarshalJSON(input []byte) error {
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
