package models

import (
	"encoding/json"
)

// ContractInterfaceResponse represents a ContractInterfaceResponse struct.
// GET request to get contract interface
type ContractInterfaceResponse struct {
	// List of defined methods
	Functions []interface{} `json:"functions"`
	// List of defined variables
	Variables []interface{} `json:"variables"`
	// List of defined data-maps
	Maps []interface{} `json:"maps"`
	// List of fungible tokens in the contract
	FungibleTokens []interface{} `json:"fungible_tokens"`
	// List of non-fungible tokens in the contract
	NonFungibleTokens []interface{} `json:"non_fungible_tokens"`
}

// MarshalJSON implements the json.Marshaler interface for ContractInterfaceResponse.
// It customizes the JSON marshaling process for ContractInterfaceResponse objects.
func (c *ContractInterfaceResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the ContractInterfaceResponse object to a map representation for JSON marshaling.
func (c *ContractInterfaceResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["functions"] = c.Functions
	structMap["variables"] = c.Variables
	structMap["maps"] = c.Maps
	structMap["fungible_tokens"] = c.FungibleTokens
	structMap["non_fungible_tokens"] = c.NonFungibleTokens
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ContractInterfaceResponse.
// It customizes the JSON unmarshaling process for ContractInterfaceResponse objects.
func (c *ContractInterfaceResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Functions         []interface{} `json:"functions"`
		Variables         []interface{} `json:"variables"`
		Maps              []interface{} `json:"maps"`
		FungibleTokens    []interface{} `json:"fungible_tokens"`
		NonFungibleTokens []interface{} `json:"non_fungible_tokens"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Functions = temp.Functions
	c.Variables = temp.Variables
	c.Maps = temp.Maps
	c.FungibleTokens = temp.FungibleTokens
	c.NonFungibleTokens = temp.NonFungibleTokens
	return nil
}
