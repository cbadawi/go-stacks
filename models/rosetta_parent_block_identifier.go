package models

import (
	"encoding/json"
)

// RosettaParentBlockIdentifier represents a RosettaParentBlockIdentifier struct.
// The block_identifier uniquely identifies a block in a particular network.
type RosettaParentBlockIdentifier struct {
	// This is also known as the block height.
	Index int `json:"index"`
	// Block hash
	Hash string `json:"hash"`
}

// MarshalJSON implements the json.Marshaler interface for RosettaParentBlockIdentifier.
// It customizes the JSON marshaling process for RosettaParentBlockIdentifier objects.
func (r *RosettaParentBlockIdentifier) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the RosettaParentBlockIdentifier object to a map representation for JSON marshaling.
func (r *RosettaParentBlockIdentifier) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["index"] = r.Index
	structMap["hash"] = r.Hash
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RosettaParentBlockIdentifier.
// It customizes the JSON unmarshaling process for RosettaParentBlockIdentifier objects.
func (r *RosettaParentBlockIdentifier) UnmarshalJSON(input []byte) error {
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
