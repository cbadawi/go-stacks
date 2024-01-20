package models

import (
	"encoding/json"
)

// Value represents a Value struct.
type Value struct {
	Hex  string `json:"hex"`
	Repr string `json:"repr"`
}

// MarshalJSON implements the json.Marshaler interface for Value.
// It customizes the JSON marshaling process for Value objects.
func (v *Value) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(v.toMap())
}

// toMap converts the Value object to a map representation for JSON marshaling.
func (v *Value) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["hex"] = v.Hex
	structMap["repr"] = v.Repr
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Value.
// It customizes the JSON unmarshaling process for Value objects.
func (v *Value) UnmarshalJSON(input []byte) error {
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
