package models

import (
	"encoding/json"
)

// RosettaPartialBlockIdentifier represents a RosettaPartialBlockIdentifier struct.
// When fetching data by BlockIdentifier, it may be possible to only specify the index or hash. If neither property is specified, it is assumed that the client is making a request at the current block.
type RosettaPartialBlockIdentifier struct {
	// This is also known as the block hash.
	Hash *string `json:"hash,omitempty"`
	// This is also known as the block height.
	Index *int `json:"index,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for RosettaPartialBlockIdentifier.
// It customizes the JSON marshaling process for RosettaPartialBlockIdentifier objects.
func (r *RosettaPartialBlockIdentifier) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the RosettaPartialBlockIdentifier object to a map representation for JSON marshaling.
func (r *RosettaPartialBlockIdentifier) toMap() map[string]any {
	structMap := make(map[string]any)
	if r.Hash != nil {
		structMap["hash"] = r.Hash
	}
	if r.Index != nil {
		structMap["index"] = r.Index
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RosettaPartialBlockIdentifier.
// It customizes the JSON unmarshaling process for RosettaPartialBlockIdentifier objects.
func (r *RosettaPartialBlockIdentifier) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		Hash  *string `json:"hash,omitempty"`
		Index *int    `json:"index,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	r.Hash = temp.Hash
	r.Index = temp.Index
	return nil
}
