package models

import (
	"encoding/json"
)

// FunctionArg represents a FunctionArg struct.
type FunctionArg struct {
	Hex  string `json:"hex"`
	Repr string `json:"repr"`
	Name string `json:"name"`
	Type string `json:"type"`
}

// MarshalJSON implements the json.Marshaler interface for FunctionArg.
// It customizes the JSON marshaling process for FunctionArg objects.
func (f *FunctionArg) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(f.toMap())
}

// toMap converts the FunctionArg object to a map representation for JSON marshaling.
func (f *FunctionArg) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["hex"] = f.Hex
	structMap["repr"] = f.Repr
	structMap["name"] = f.Name
	structMap["type"] = f.Type
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for FunctionArg.
// It customizes the JSON unmarshaling process for FunctionArg objects.
func (f *FunctionArg) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		Hex  string `json:"hex"`
		Repr string `json:"repr"`
		Name string `json:"name"`
		Type string `json:"type"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	f.Hex = temp.Hex
	f.Repr = temp.Repr
	f.Name = temp.Name
	f.Type = temp.Type
	return nil
}
