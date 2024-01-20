package models

import (
	"encoding/json"
)

// TransactionIdentifier represents a TransactionIdentifier struct.
// The transaction_identifier uniquely identifies a transaction in a particular network and block or in the mempool.
type TransactionIdentifier struct {
	// Any transactions that are attributable only to a block (ex: a block event) should use the hash of the block as the identifier.
	Hash string `json:"hash"`
}

// MarshalJSON implements the json.Marshaler interface for TransactionIdentifier.
// It customizes the JSON marshaling process for TransactionIdentifier objects.
func (t *TransactionIdentifier) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(t.toMap())
}

// toMap converts the TransactionIdentifier object to a map representation for JSON marshaling.
func (t *TransactionIdentifier) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["hash"] = t.Hash
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for TransactionIdentifier.
// It customizes the JSON unmarshaling process for TransactionIdentifier objects.
func (t *TransactionIdentifier) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Hash string `json:"hash"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	t.Hash = temp.Hash
	return nil
}
