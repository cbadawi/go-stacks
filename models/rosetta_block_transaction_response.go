package models

import (
	"encoding/json"
)

// RosettaBlockTransactionResponse represents a RosettaBlockTransactionResponse struct.
// A BlockTransactionResponse contains information about a block transaction.
type RosettaBlockTransactionResponse struct {
	// Transactions contain an array of Operations that are attributable to the same TransactionIdentifier.
	Transaction RosettaTransaction `json:"transaction"`
}

// MarshalJSON implements the json.Marshaler interface for RosettaBlockTransactionResponse.
// It customizes the JSON marshaling process for RosettaBlockTransactionResponse objects.
func (r *RosettaBlockTransactionResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the RosettaBlockTransactionResponse object to a map representation for JSON marshaling.
func (r *RosettaBlockTransactionResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["transaction"] = r.Transaction
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RosettaBlockTransactionResponse.
// It customizes the JSON unmarshaling process for RosettaBlockTransactionResponse objects.
func (r *RosettaBlockTransactionResponse) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		Transaction RosettaTransaction `json:"transaction"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	r.Transaction = temp.Transaction
	return nil
}
