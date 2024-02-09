package models

import (
	"encoding/json"
)

// RosettaMempoolTransactionRequest represents a RosettaMempoolTransactionRequest struct.
// A MempoolTransactionRequest is utilized to retrieve a transaction from the mempool.
type RosettaMempoolTransactionRequest struct {
	// The network_identifier specifies which network a particular object is associated with.
	NetworkIdentifier NetworkIdentifier `json:"network_identifier"`
	// The transaction_identifier uniquely identifies a transaction in a particular network and block or in the mempool.
	TransactionIdentifier TransactionIdentifier `json:"transaction_identifier"`
}

// MarshalJSON implements the json.Marshaler interface for RosettaMempoolTransactionRequest.
// It customizes the JSON marshaling process for RosettaMempoolTransactionRequest objects.
func (r *RosettaMempoolTransactionRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the RosettaMempoolTransactionRequest object to a map representation for JSON marshaling.
func (r *RosettaMempoolTransactionRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["network_identifier"] = r.NetworkIdentifier
	structMap["transaction_identifier"] = r.TransactionIdentifier
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RosettaMempoolTransactionRequest.
// It customizes the JSON unmarshaling process for RosettaMempoolTransactionRequest objects.
func (r *RosettaMempoolTransactionRequest) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		NetworkIdentifier     NetworkIdentifier     `json:"network_identifier"`
		TransactionIdentifier TransactionIdentifier `json:"transaction_identifier"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	r.NetworkIdentifier = temp.NetworkIdentifier
	r.TransactionIdentifier = temp.TransactionIdentifier
	return nil
}
