package models

import (
	"encoding/json"
)

// RosettaOperationIdentifier represents a RosettaOperationIdentifier struct.
// The operation_identifier uniquely identifies an operation within a transaction.
type RosettaOperationIdentifier struct {
	// The operation index is used to ensure each operation has a unique identifier within a transaction. This index is only relative to the transaction and NOT GLOBAL. The operations in each transaction should start from index 0. To clarify, there may not be any notion of an operation index in the blockchain being described.
	Index int `json:"index"`
	// Some blockchains specify an operation index that is essential for client use. For example, Bitcoin uses a network_index to identify which UTXO was used in a transaction. network_index should not be populated if there is no notion of an operation index in a blockchain (typically most account-based blockchains).
	NetworkIndex *int `json:"network_index,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for RosettaOperationIdentifier.
// It customizes the JSON marshaling process for RosettaOperationIdentifier objects.
func (r *RosettaOperationIdentifier) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the RosettaOperationIdentifier object to a map representation for JSON marshaling.
func (r *RosettaOperationIdentifier) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["index"] = r.Index
	if r.NetworkIndex != nil {
		structMap["network_index"] = r.NetworkIndex
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RosettaOperationIdentifier.
// It customizes the JSON unmarshaling process for RosettaOperationIdentifier objects.
func (r *RosettaOperationIdentifier) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		Index        int  `json:"index"`
		NetworkIndex *int `json:"network_index,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	r.Index = temp.Index
	r.NetworkIndex = temp.NetworkIndex
	return nil
}
