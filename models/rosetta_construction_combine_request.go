package models

import (
	"encoding/json"
)

// RosettaConstructionCombineRequest represents a RosettaConstructionCombineRequest struct.
// RosettaConstructionCombineRequest is the input to the /construction/combine endpoint. It contains the unsigned transaction blob returned by /construction/payloads and all required signatures to create a network transaction.
type RosettaConstructionCombineRequest struct {
	// The network_identifier specifies which network a particular object is associated with.
	NetworkIdentifier   NetworkIdentifier  `json:"network_identifier"`
	UnsignedTransaction string             `json:"unsigned_transaction"`
	Signatures          []RosettaSignature `json:"signatures"`
}

// MarshalJSON implements the json.Marshaler interface for RosettaConstructionCombineRequest.
// It customizes the JSON marshaling process for RosettaConstructionCombineRequest objects.
func (r *RosettaConstructionCombineRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the RosettaConstructionCombineRequest object to a map representation for JSON marshaling.
func (r *RosettaConstructionCombineRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["network_identifier"] = r.NetworkIdentifier
	structMap["unsigned_transaction"] = r.UnsignedTransaction
	structMap["signatures"] = r.Signatures
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RosettaConstructionCombineRequest.
// It customizes the JSON unmarshaling process for RosettaConstructionCombineRequest objects.
func (r *RosettaConstructionCombineRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		NetworkIdentifier   NetworkIdentifier  `json:"network_identifier"`
		UnsignedTransaction string             `json:"unsigned_transaction"`
		Signatures          []RosettaSignature `json:"signatures"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	r.NetworkIdentifier = temp.NetworkIdentifier
	r.UnsignedTransaction = temp.UnsignedTransaction
	r.Signatures = temp.Signatures
	return nil
}
