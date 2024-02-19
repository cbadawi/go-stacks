package models

import (
	"encoding/json"
)

// ReadOnlyFunctionArgs represents a ReadOnlyFunctionArgs struct.
// Describes representation of a Type-0 Stacks 2.0 transaction. https://github.com/blockstack/stacks-blockchain/blob/master/sip/sip-005-blocks-and-transactions.md#type-0-transferring-an-asset
type ReadOnlyFunctionArgs struct {
	// The simulated tx-sender
	Sender string `json:"sender"`
	// An array of hex serialized Clarity values
	Arguments []string `json:"arguments"`
}

// MarshalJSON implements the json.Marshaler interface for ReadOnlyFunctionArgs.
// It customizes the JSON marshaling process for ReadOnlyFunctionArgs objects.
func (r *ReadOnlyFunctionArgs) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the ReadOnlyFunctionArgs object to a map representation for JSON marshaling.
func (r *ReadOnlyFunctionArgs) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["sender"] = r.Sender
	structMap["arguments"] = r.Arguments
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ReadOnlyFunctionArgs.
// It customizes the JSON unmarshaling process for ReadOnlyFunctionArgs objects.
func (r *ReadOnlyFunctionArgs) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		Sender    string   `json:"sender"`
		Arguments []string `json:"arguments"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	r.Sender = temp.Sender
	r.Arguments = temp.Arguments
	return nil
}
