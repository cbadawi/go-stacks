package models

import (
	"encoding/json"
)

// RosettaConstructionHashRequest represents a RosettaConstructionHashRequest struct.
// TransactionHash returns the network-specific transaction hash for a signed transaction.
type RosettaConstructionHashRequest struct {
	// The network_identifier specifies which network a particular object is associated with.
	NetworkIdentifier NetworkIdentifier `json:"network_identifier"`
	// Signed transaction
	SignedTransaction string `json:"signed_transaction"`
}

// MarshalJSON implements the json.Marshaler interface for RosettaConstructionHashRequest.
// It customizes the JSON marshaling process for RosettaConstructionHashRequest objects.
func (r *RosettaConstructionHashRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the RosettaConstructionHashRequest object to a map representation for JSON marshaling.
func (r *RosettaConstructionHashRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["network_identifier"] = r.NetworkIdentifier
	structMap["signed_transaction"] = r.SignedTransaction
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RosettaConstructionHashRequest.
// It customizes the JSON unmarshaling process for RosettaConstructionHashRequest objects.
func (r *RosettaConstructionHashRequest) UnmarshalJSON(input []byte) error {
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
