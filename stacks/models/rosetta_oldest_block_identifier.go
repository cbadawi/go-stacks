package models

import (
	"encoding/json"
)

// RosettaOldestBlockIdentifier represents a RosettaOldestBlockIdentifier struct.
// The block_identifier uniquely identifies a block in a particular network.
type RosettaOldestBlockIdentifier struct {
	// This is also known as the block height.
	Index int `json:"index"`
	// Block hash
	Hash string `json:"hash"`
}

// MarshalJSON implements the json.Marshaler interface for RosettaOldestBlockIdentifier.
// It customizes the JSON marshaling process for RosettaOldestBlockIdentifier objects.
func (r *RosettaOldestBlockIdentifier) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the RosettaOldestBlockIdentifier object to a map representation for JSON marshaling.
func (r *RosettaOldestBlockIdentifier) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["index"] = r.Index
	structMap["hash"] = r.Hash
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RosettaOldestBlockIdentifier.
// It customizes the JSON unmarshaling process for RosettaOldestBlockIdentifier objects.
func (r *RosettaOldestBlockIdentifier) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		Index int    `json:"index"`
		Hash  string `json:"hash"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	r.Index = temp.Index
	r.Hash = temp.Hash
	return nil
}
