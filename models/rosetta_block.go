package models

import (
	"encoding/json"
)

// RosettaBlock represents a RosettaBlock struct.
// Blocks contain an array of Transactions that occurred at a particular BlockIdentifier. A hard requirement for blocks returned by Rosetta implementations is that they MUST be inalterable: once a client has requested and received a block identified by a specific BlockIndentifier, all future calls for that same BlockIdentifier must return the same block contents.
type RosettaBlock struct {
	// The block_identifier uniquely identifies a block in a particular network.
	BlockIdentifier RosettaBlockIdentifier `json:"block_identifier"`
	// The block_identifier uniquely identifies a block in a particular network.
	ParentBlockIdentifier RosettaParentBlockIdentifier `json:"parent_block_identifier"`
	// The timestamp of the block in milliseconds since the Unix Epoch. The timestamp is stored in milliseconds because some blockchains produce blocks more often than once a second.
	Timestamp int `json:"timestamp"`
	// All the transactions in the block
	Transactions []RosettaTransaction `json:"transactions"`
	// meta data
	Metadata *Metadata4 `json:"metadata,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for RosettaBlock.
// It customizes the JSON marshaling process for RosettaBlock objects.
func (r *RosettaBlock) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the RosettaBlock object to a map representation for JSON marshaling.
func (r *RosettaBlock) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["block_identifier"] = r.BlockIdentifier
	structMap["parent_block_identifier"] = r.ParentBlockIdentifier
	structMap["timestamp"] = r.Timestamp
	structMap["transactions"] = r.Transactions
	if r.Metadata != nil {
		structMap["metadata"] = r.Metadata
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RosettaBlock.
// It customizes the JSON unmarshaling process for RosettaBlock objects.
func (r *RosettaBlock) UnmarshalJSON(input []byte) error {
	temp := &struct {
		BlockIdentifier       RosettaBlockIdentifier       `json:"block_identifier"`
		ParentBlockIdentifier RosettaParentBlockIdentifier `json:"parent_block_identifier"`
		Timestamp             int                          `json:"timestamp"`
		Transactions          []RosettaTransaction         `json:"transactions"`
		Metadata              *Metadata4                   `json:"metadata,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	r.BlockIdentifier = temp.BlockIdentifier
	r.ParentBlockIdentifier = temp.ParentBlockIdentifier
	r.Timestamp = temp.Timestamp
	r.Transactions = temp.Transactions
	r.Metadata = temp.Metadata
	return nil
}
