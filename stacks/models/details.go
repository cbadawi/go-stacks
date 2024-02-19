package models

import (
	"encoding/json"
)

// Details represents a Details struct.
// Often times it is useful to return context specific to the request that caused the error (i.e. a sample of the stack trace or impacted account) in addition to the standard error message.
type Details struct {
	Address *string `json:"address,omitempty"`
	Error   *string `json:"error,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for Details.
// It customizes the JSON marshaling process for Details objects.
func (d *Details) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(d.toMap())
}

// toMap converts the Details object to a map representation for JSON marshaling.
func (d *Details) toMap() map[string]any {
	structMap := make(map[string]any)
	if d.Address != nil {
		structMap["address"] = d.Address
	}
	if d.Error != nil {
		structMap["error"] = d.Error
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Details.
// It customizes the JSON unmarshaling process for Details objects.
func (d *Details) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		Address *string `json:"address,omitempty"`
		Error   *string `json:"error,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	d.Address = temp.Address
	d.Error = temp.Error
	return nil
}
