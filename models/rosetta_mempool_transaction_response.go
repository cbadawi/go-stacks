package models

import (
	"encoding/json"
)

// RosettaMempoolTransactionResponse represents a RosettaMempoolTransactionResponse struct.
// A MempoolTransactionResponse contains an estimate of a mempool transaction. It may not be possible to know the full impact of a transaction in the mempool (ex: fee paid).
type RosettaMempoolTransactionResponse struct {
	// Transactions contain an array of Operations that are attributable to the same TransactionIdentifier.
	Transaction RosettaTransaction `json:"transaction"`
	Metadata    *interface{}       `json:"metadata,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for RosettaMempoolTransactionResponse.
// It customizes the JSON marshaling process for RosettaMempoolTransactionResponse objects.
func (r *RosettaMempoolTransactionResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the RosettaMempoolTransactionResponse object to a map representation for JSON marshaling.
func (r *RosettaMempoolTransactionResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["transaction"] = r.Transaction
	if r.Metadata != nil {
		structMap["metadata"] = r.Metadata
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RosettaMempoolTransactionResponse.
// It customizes the JSON unmarshaling process for RosettaMempoolTransactionResponse objects.
func (r *RosettaMempoolTransactionResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Transaction RosettaTransaction `json:"transaction"`
		Metadata    *interface{}       `json:"metadata,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	r.Transaction = temp.Transaction
	r.Metadata = temp.Metadata
	return nil
}
