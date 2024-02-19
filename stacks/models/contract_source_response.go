package models

import (
	"encoding/json"
)

// ContractSourceResponse represents a ContractSourceResponse struct.
// GET request to get contract source
type ContractSourceResponse struct {
	Source        string `json:"source"`
	PublishHeight int    `json:"publish_height"`
	Proof         string `json:"proof"`
}

// MarshalJSON implements the json.Marshaler interface for ContractSourceResponse.
// It customizes the JSON marshaling process for ContractSourceResponse objects.
func (c *ContractSourceResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the ContractSourceResponse object to a map representation for JSON marshaling.
func (c *ContractSourceResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["source"] = c.Source
	structMap["publish_height"] = c.PublishHeight
	structMap["proof"] = c.Proof
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ContractSourceResponse.
// It customizes the JSON unmarshaling process for ContractSourceResponse objects.
func (c *ContractSourceResponse) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		Source        string `json:"source"`
		PublishHeight int    `json:"publish_height"`
		Proof         string `json:"proof"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Source = temp.Source
	c.PublishHeight = temp.PublishHeight
	c.Proof = temp.Proof
	return nil
}
