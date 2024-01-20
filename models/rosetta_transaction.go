package models

import (
	"encoding/json"
)

// RosettaTransaction represents a RosettaTransaction struct.
// Transactions contain an array of Operations that are attributable to the same TransactionIdentifier.
type RosettaTransaction struct {
	// The transaction_identifier uniquely identifies a transaction in a particular network and block or in the mempool.
	TransactionIdentifier TransactionIdentifier `json:"transaction_identifier"`
	// List of operations
	Operations []RosettaOperation `json:"operations"`
	// Transactions that are related to other transactions (like a cross-shard transaction) should include the tranaction_identifier of these transactions in the metadata.
	Metadata *Metadata3 `json:"metadata,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for RosettaTransaction.
// It customizes the JSON marshaling process for RosettaTransaction objects.
func (r *RosettaTransaction) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the RosettaTransaction object to a map representation for JSON marshaling.
func (r *RosettaTransaction) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["transaction_identifier"] = r.TransactionIdentifier
	structMap["operations"] = r.Operations
	if r.Metadata != nil {
		structMap["metadata"] = r.Metadata
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RosettaTransaction.
// It customizes the JSON unmarshaling process for RosettaTransaction objects.
func (r *RosettaTransaction) UnmarshalJSON(input []byte) error {
	temp := &struct {
		TransactionIdentifier TransactionIdentifier `json:"transaction_identifier"`
		Operations            []RosettaOperation    `json:"operations"`
		Metadata              *Metadata3            `json:"metadata,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	r.TransactionIdentifier = temp.TransactionIdentifier
	r.Operations = temp.Operations
	r.Metadata = temp.Metadata
	return nil
}
