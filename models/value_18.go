package models

import (
	"encoding/json"
)

// Value18 represents a Value18 struct.
// Non-Fungible Token value
type Value18 struct {
	// Hex string representing the identifier of the Non-Fungible Token
	Hex string `json:"hex"`
	// Readable string of the Non-Fungible Token identifier
	Repr string `json:"repr"`
}

// MarshalJSON implements the json.Marshaler interface for Value18.
// It customizes the JSON marshaling process for Value18 objects.
func (v *Value18) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(v.toMap())
}

// toMap converts the Value18 object to a map representation for JSON marshaling.
func (v *Value18) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["hex"] = v.Hex
	structMap["repr"] = v.Repr
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Value18.
// It customizes the JSON unmarshaling process for Value18 objects.
func (v *Value18) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		Hex  string `json:"hex"`
		Repr string `json:"repr"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	v.Hex = temp.Hex
	v.Repr = temp.Repr
	return nil
}
