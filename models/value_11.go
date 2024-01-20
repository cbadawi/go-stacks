package models

import (
	"encoding/json"
)

// Value11 represents a Value11 struct.
// Non Fungible Token asset value.
type Value11 struct {
	Hex  string `json:"hex"`
	Repr string `json:"repr"`
}

// MarshalJSON implements the json.Marshaler interface for Value11.
// It customizes the JSON marshaling process for Value11 objects.
func (v *Value11) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(v.toMap())
}

// toMap converts the Value11 object to a map representation for JSON marshaling.
func (v *Value11) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["hex"] = v.Hex
	structMap["repr"] = v.Repr
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Value11.
// It customizes the JSON unmarshaling process for Value11 objects.
func (v *Value11) UnmarshalJSON(input []byte) error {
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
