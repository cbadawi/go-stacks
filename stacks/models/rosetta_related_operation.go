package models

import (
	"encoding/json"
)

// RosettaRelatedOperation represents a RosettaRelatedOperation struct.
// Restrict referenced related_operations to identifier indexes < the current operation_identifier.index. This ensures there exists a clear DAG-structure of relations. Since operations are one-sided, one could imagine relating operations in a single transfer or linking operations in a call tree.
type RosettaRelatedOperation struct {
	// Describes the index of related operation.
	Index int `json:"index"`
	// Some blockchains specify an operation index that is essential for client use. network_index should not be populated if there is no notion of an operation index in a blockchain (typically most account-based blockchains).
	NetworkIndex *int `json:"network_index,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for RosettaRelatedOperation.
// It customizes the JSON marshaling process for RosettaRelatedOperation objects.
func (r *RosettaRelatedOperation) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the RosettaRelatedOperation object to a map representation for JSON marshaling.
func (r *RosettaRelatedOperation) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["index"] = r.Index
	if r.NetworkIndex != nil {
		structMap["network_index"] = r.NetworkIndex
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RosettaRelatedOperation.
// It customizes the JSON unmarshaling process for RosettaRelatedOperation objects.
func (r *RosettaRelatedOperation) UnmarshalJSON(input []byte) error {
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
