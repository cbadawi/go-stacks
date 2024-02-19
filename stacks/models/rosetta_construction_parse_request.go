package models

import (
	"encoding/json"
)

// RosettaConstructionParseRequest represents a RosettaConstructionParseRequest struct.
// Parse is called on both unsigned and signed transactions to understand the intent of the formulated transaction. This is run as a sanity check before signing (after /construction/payloads) and before broadcast (after /construction/combine).
type RosettaConstructionParseRequest struct {
	// The network_identifier specifies which network a particular object is associated with.
	NetworkIdentifier NetworkIdentifier `json:"network_identifier"`
	// Signed is a boolean indicating whether the transaction is signed.
	Signed bool `json:"signed"`
	// This must be either the unsigned transaction blob returned by /construction/payloads or the signed transaction blob returned by /construction/combine.
	Transaction string `json:"transaction"`
}

// MarshalJSON implements the json.Marshaler interface for RosettaConstructionParseRequest.
// It customizes the JSON marshaling process for RosettaConstructionParseRequest objects.
func (r *RosettaConstructionParseRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the RosettaConstructionParseRequest object to a map representation for JSON marshaling.
func (r *RosettaConstructionParseRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["network_identifier"] = r.NetworkIdentifier
	structMap["signed"] = r.Signed
	structMap["transaction"] = r.Transaction
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RosettaConstructionParseRequest.
// It customizes the JSON unmarshaling process for RosettaConstructionParseRequest objects.
func (r *RosettaConstructionParseRequest) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		NetworkIdentifier NetworkIdentifier `json:"network_identifier"`
		Signed            bool              `json:"signed"`
		Transaction       string            `json:"transaction"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	r.NetworkIdentifier = temp.NetworkIdentifier
	r.Signed = temp.Signed
	r.Transaction = temp.Transaction
	return nil
}
