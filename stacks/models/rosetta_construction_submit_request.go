package models

import (
	"encoding/json"
)

// RosettaConstructionSubmitRequest represents a RosettaConstructionSubmitRequest struct.
// Submit the transaction in blockchain
type RosettaConstructionSubmitRequest struct {
	// The network_identifier specifies which network a particular object is associated with.
	NetworkIdentifier NetworkIdentifier `json:"network_identifier"`
	// Signed transaction
	SignedTransaction string `json:"signed_transaction"`
}

// MarshalJSON implements the json.Marshaler interface for RosettaConstructionSubmitRequest.
// It customizes the JSON marshaling process for RosettaConstructionSubmitRequest objects.
func (r *RosettaConstructionSubmitRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the RosettaConstructionSubmitRequest object to a map representation for JSON marshaling.
func (r *RosettaConstructionSubmitRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["network_identifier"] = r.NetworkIdentifier
	structMap["signed_transaction"] = r.SignedTransaction
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RosettaConstructionSubmitRequest.
// It customizes the JSON unmarshaling process for RosettaConstructionSubmitRequest objects.
func (r *RosettaConstructionSubmitRequest) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		NetworkIdentifier NetworkIdentifier `json:"network_identifier"`
		SignedTransaction string            `json:"signed_transaction"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	r.NetworkIdentifier = temp.NetworkIdentifier
	r.SignedTransaction = temp.SignedTransaction
	return nil
}
