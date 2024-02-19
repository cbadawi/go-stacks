package models

import (
	"encoding/json"
)

// RosettaConstructionHashResponse represents a RosettaConstructionHashResponse struct.
// TransactionIdentifier contains the transaction_identifier of a transaction that was submitted to either /construction/hash or /construction/submit.
type RosettaConstructionHashResponse struct {
	// The transaction_identifier uniquely identifies a transaction in a particular network and block or in the mempool.
	TransactionIdentifier TransactionIdentifier `json:"transaction_identifier"`
	Metadata              *interface{}          `json:"metadata,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for RosettaConstructionHashResponse.
// It customizes the JSON marshaling process for RosettaConstructionHashResponse objects.
func (r *RosettaConstructionHashResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the RosettaConstructionHashResponse object to a map representation for JSON marshaling.
func (r *RosettaConstructionHashResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["transaction_identifier"] = r.TransactionIdentifier
	if r.Metadata != nil {
		structMap["metadata"] = r.Metadata
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RosettaConstructionHashResponse.
// It customizes the JSON unmarshaling process for RosettaConstructionHashResponse objects.
func (r *RosettaConstructionHashResponse) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		TransactionIdentifier TransactionIdentifier `json:"transaction_identifier"`
		Metadata              *interface{}          `json:"metadata,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	r.TransactionIdentifier = temp.TransactionIdentifier
	r.Metadata = temp.Metadata
	return nil
}
