package models

import (
	"encoding/json"
)

// AssetValue represents a AssetValue struct.
type AssetValue struct {
	Hex  string `json:"hex"`
	Repr string `json:"repr"`
}

// MarshalJSON implements the json.Marshaler interface for AssetValue.
// It customizes the JSON marshaling process for AssetValue objects.
func (a *AssetValue) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(a.toMap())
}

// toMap converts the AssetValue object to a map representation for JSON marshaling.
func (a *AssetValue) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["hex"] = a.Hex
	structMap["repr"] = a.Repr
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AssetValue.
// It customizes the JSON unmarshaling process for AssetValue objects.
func (a *AssetValue) UnmarshalJSON(input []byte) error {
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

	a.Hex = temp.Hex
	a.Repr = temp.Repr
	return nil
}
