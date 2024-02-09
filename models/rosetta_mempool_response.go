package models

import (
	"encoding/json"
)

// RosettaMempoolResponse represents a RosettaMempoolResponse struct.
// A MempoolResponse contains all transaction identifiers in the mempool for a particular network_identifier.
type RosettaMempoolResponse struct {
	TransactionIdentifiers []TransactionIdentifier `json:"transaction_identifiers"`
	Metadata               *interface{}            `json:"metadata,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for RosettaMempoolResponse.
// It customizes the JSON marshaling process for RosettaMempoolResponse objects.
func (r *RosettaMempoolResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the RosettaMempoolResponse object to a map representation for JSON marshaling.
func (r *RosettaMempoolResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["transaction_identifiers"] = r.TransactionIdentifiers
	if r.Metadata != nil {
		structMap["metadata"] = r.Metadata
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RosettaMempoolResponse.
// It customizes the JSON unmarshaling process for RosettaMempoolResponse objects.
func (r *RosettaMempoolResponse) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		TransactionIdentifiers []TransactionIdentifier `json:"transaction_identifiers"`
		Metadata               *interface{}            `json:"metadata,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	r.TransactionIdentifiers = temp.TransactionIdentifiers
	r.Metadata = temp.Metadata
	return nil
}
