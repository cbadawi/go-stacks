package models

import (
	"encoding/json"
)

// ContractListResponse represents a ContractListResponse struct.
// GET list of contracts
type ContractListResponse struct {
	// The number of contracts to return
	Limit int `json:"limit"`
	// The number to contracts to skip (starting at `0`)
	Offset  int              `json:"offset"`
	Results []SmartContract7 `json:"results"`
}

// MarshalJSON implements the json.Marshaler interface for ContractListResponse.
// It customizes the JSON marshaling process for ContractListResponse objects.
func (c *ContractListResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the ContractListResponse object to a map representation for JSON marshaling.
func (c *ContractListResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["limit"] = c.Limit
	structMap["offset"] = c.Offset
	structMap["results"] = c.Results
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ContractListResponse.
// It customizes the JSON unmarshaling process for ContractListResponse objects.
func (c *ContractListResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Limit   int              `json:"limit"`
		Offset  int              `json:"offset"`
		Results []SmartContract7 `json:"results"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Limit = temp.Limit
	c.Offset = temp.Offset
	c.Results = temp.Results
	return nil
}
