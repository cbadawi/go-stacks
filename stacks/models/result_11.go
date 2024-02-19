package models

import (
	"encoding/json"
)

// Result11 represents a Result11 struct.
// This object carries the search result
type Result11 struct {
	// Shows the currenty category of entity it is searched in.
	EntityType EntityTypeEnum `json:"entity_type"`
	// The id used to search this query.
	EntityId string `json:"entity_id"`
	// A block
	Metadata *Block1 `json:"metadata,omitempty"`
	// Returns basic search result information about the requested id
	BlockData BlockData `json:"block_data"`
	// Returns basic search result information about the requested id
	TxData TxData3 `json:"tx_data"`
}

// MarshalJSON implements the json.Marshaler interface for Result11.
// It customizes the JSON marshaling process for Result11 objects.
func (r *Result11) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the Result11 object to a map representation for JSON marshaling.
func (r *Result11) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["entity_type"] = r.EntityType
	structMap["entity_id"] = r.EntityId
	if r.Metadata != nil {
		structMap["metadata"] = r.Metadata
	}
	structMap["block_data"] = r.BlockData
	structMap["tx_data"] = r.TxData
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Result11.
// It customizes the JSON unmarshaling process for Result11 objects.
func (r *Result11) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		EntityType EntityTypeEnum `json:"entity_type"`
		EntityId   string         `json:"entity_id"`
		Metadata   *Block1        `json:"metadata,omitempty"`
		BlockData  BlockData      `json:"block_data"`
		TxData     TxData3        `json:"tx_data"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	r.EntityType = temp.EntityType
	r.EntityId = temp.EntityId
	r.Metadata = temp.Metadata
	r.BlockData = temp.BlockData
	r.TxData = temp.TxData
	return nil
}
