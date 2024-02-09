package models

import (
	"encoding/json"
)

// PoisonMicroblock represents a PoisonMicroblock struct.
type PoisonMicroblock struct {
	// Hex encoded microblock header
	MicroblockHeader1 string `json:"microblock_header_1"`
	// Hex encoded microblock header
	MicroblockHeader2 string `json:"microblock_header_2"`
}

// MarshalJSON implements the json.Marshaler interface for PoisonMicroblock.
// It customizes the JSON marshaling process for PoisonMicroblock objects.
func (p *PoisonMicroblock) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(p.toMap())
}

// toMap converts the PoisonMicroblock object to a map representation for JSON marshaling.
func (p *PoisonMicroblock) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["microblock_header_1"] = p.MicroblockHeader1
	structMap["microblock_header_2"] = p.MicroblockHeader2
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PoisonMicroblock.
// It customizes the JSON unmarshaling process for PoisonMicroblock objects.
func (p *PoisonMicroblock) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		MicroblockHeader1 string `json:"microblock_header_1"`
		MicroblockHeader2 string `json:"microblock_header_2"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	p.MicroblockHeader1 = temp.MicroblockHeader1
	p.MicroblockHeader2 = temp.MicroblockHeader2
	return nil
}
