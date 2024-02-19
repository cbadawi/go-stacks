package models

import (
	"encoding/json"
)

// RosettaBlockIdentifier represents a RosettaBlockIdentifier struct.
// The block_identifier uniquely identifies a block in a particular network.
type RosettaBlockIdentifier struct {
	// This is also known as the block hash.
	Hash string `json:"hash"`
	// This is also known as the block height.
	Index int `json:"index"`
}

// MarshalJSON implements the json.Marshaler interface for RosettaBlockIdentifier.
// It customizes the JSON marshaling process for RosettaBlockIdentifier objects.
func (r *RosettaBlockIdentifier) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the RosettaBlockIdentifier object to a map representation for JSON marshaling.
func (r *RosettaBlockIdentifier) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["hash"] = r.Hash
	structMap["index"] = r.Index
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RosettaBlockIdentifier.
// It customizes the JSON unmarshaling process for RosettaBlockIdentifier objects.
func (r *RosettaBlockIdentifier) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		Hash  string `json:"hash"`
		Index int    `json:"index"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	r.Hash = temp.Hash
	r.Index = temp.Index
	return nil
}
