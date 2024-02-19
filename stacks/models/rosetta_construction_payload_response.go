package models

import (
	"encoding/json"
)

// RosettaConstructionPayloadResponse represents a RosettaConstructionPayloadResponse struct.
// RosettaConstructionPayloadResponse is returned by /construction/payloads. It contains an unsigned transaction blob (that is usually needed to construct the a network transaction from a collection of signatures) and an array of payloads that must be signed by the caller.
type RosettaConstructionPayloadResponse struct {
	// This is an unsigned transaction blob (that is usually needed to construct the a network transaction from a collection of signatures)
	UnsignedTransaction string `json:"unsigned_transaction"`
	// An array of payloads that must be signed by the caller
	Payloads []SigningPayload `json:"payloads"`
}

// MarshalJSON implements the json.Marshaler interface for RosettaConstructionPayloadResponse.
// It customizes the JSON marshaling process for RosettaConstructionPayloadResponse objects.
func (r *RosettaConstructionPayloadResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the RosettaConstructionPayloadResponse object to a map representation for JSON marshaling.
func (r *RosettaConstructionPayloadResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["unsigned_transaction"] = r.UnsignedTransaction
	structMap["payloads"] = r.Payloads
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RosettaConstructionPayloadResponse.
// It customizes the JSON unmarshaling process for RosettaConstructionPayloadResponse objects.
func (r *RosettaConstructionPayloadResponse) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		UnsignedTransaction string           `json:"unsigned_transaction"`
		Payloads            []SigningPayload `json:"payloads"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	r.UnsignedTransaction = temp.UnsignedTransaction
	r.Payloads = temp.Payloads
	return nil
}
