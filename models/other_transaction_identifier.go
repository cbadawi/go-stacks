package models

import (
	"encoding/json"
)

// OtherTransactionIdentifier represents a OtherTransactionIdentifier struct.
// The transaction_identifier uniquely identifies a transaction in a particular network and block or in the mempool.
type OtherTransactionIdentifier struct {
	// Any transactions that are attributable only to a block (ex: a block event) should use the hash of the block as the identifier.
	Hash string `json:"hash"`
}

// MarshalJSON implements the json.Marshaler interface for OtherTransactionIdentifier.
// It customizes the JSON marshaling process for OtherTransactionIdentifier objects.
func (o *OtherTransactionIdentifier) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(o.toMap())
}

// toMap converts the OtherTransactionIdentifier object to a map representation for JSON marshaling.
func (o *OtherTransactionIdentifier) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["hash"] = o.Hash
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OtherTransactionIdentifier.
// It customizes the JSON unmarshaling process for OtherTransactionIdentifier objects.
func (o *OtherTransactionIdentifier) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Hash string `json:"hash"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	o.Hash = temp.Hash
	return nil
}
