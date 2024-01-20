package models

import (
	"encoding/json"
)

// ContractLog represents a ContractLog struct.
type ContractLog struct {
	ContractId string `json:"contract_id"`
	Topic      string `json:"topic"`
	Value      Value  `json:"value"`
}

// MarshalJSON implements the json.Marshaler interface for ContractLog.
// It customizes the JSON marshaling process for ContractLog objects.
func (c *ContractLog) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the ContractLog object to a map representation for JSON marshaling.
func (c *ContractLog) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["contract_id"] = c.ContractId
	structMap["topic"] = c.Topic
	structMap["value"] = c.Value
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ContractLog.
// It customizes the JSON unmarshaling process for ContractLog objects.
func (c *ContractLog) UnmarshalJSON(input []byte) error {
	temp := &struct {
		ContractId string `json:"contract_id"`
		Topic      string `json:"topic"`
		Value      Value  `json:"value"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.ContractId = temp.ContractId
	c.Topic = temp.Topic
	c.Value = temp.Value
	return nil
}
