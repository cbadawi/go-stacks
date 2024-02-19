package models

import (
	"encoding/json"
)

// RosettaConstructionCombineResponse represents a RosettaConstructionCombineResponse struct.
// RosettaConstructionCombineResponse is returned by /construction/combine. The network payload will be sent directly to the construction/submit endpoint.
type RosettaConstructionCombineResponse struct {
	// Signed transaction bytes in hex
	SignedTransaction string `json:"signed_transaction"`
}

// MarshalJSON implements the json.Marshaler interface for RosettaConstructionCombineResponse.
// It customizes the JSON marshaling process for RosettaConstructionCombineResponse objects.
func (r *RosettaConstructionCombineResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the RosettaConstructionCombineResponse object to a map representation for JSON marshaling.
func (r *RosettaConstructionCombineResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["signed_transaction"] = r.SignedTransaction
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RosettaConstructionCombineResponse.
// It customizes the JSON unmarshaling process for RosettaConstructionCombineResponse objects.
func (r *RosettaConstructionCombineResponse) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		SignedTransaction string `json:"signed_transaction"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	r.SignedTransaction = temp.SignedTransaction
	return nil
}
