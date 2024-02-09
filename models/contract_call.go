package models

import (
	"encoding/json"
)

// ContractCall represents a ContractCall struct.
type ContractCall struct {
	// Contract identifier formatted as `<principaladdress>.<contract_name>`
	ContractId string `json:"contract_id"`
	// Name of the Clarity function to be invoked
	FunctionName string `json:"function_name"`
	// Function definition, including function name and type as well as parameter names and types
	FunctionSignature string `json:"function_signature"`
	// List of arguments used to invoke the function
	FunctionArgs []FunctionArg `json:"function_args,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for ContractCall.
// It customizes the JSON marshaling process for ContractCall objects.
func (c *ContractCall) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the ContractCall object to a map representation for JSON marshaling.
func (c *ContractCall) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["contract_id"] = c.ContractId
	structMap["function_name"] = c.FunctionName
	structMap["function_signature"] = c.FunctionSignature
	if c.FunctionArgs != nil {
		structMap["function_args"] = c.FunctionArgs
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ContractCall.
// It customizes the JSON unmarshaling process for ContractCall objects.
func (c *ContractCall) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		ContractId        string        `json:"contract_id"`
		FunctionName      string        `json:"function_name"`
		FunctionSignature string        `json:"function_signature"`
		FunctionArgs      []FunctionArg `json:"function_args,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.ContractId = temp.ContractId
	c.FunctionName = temp.FunctionName
	c.FunctionSignature = temp.FunctionSignature
	c.FunctionArgs = temp.FunctionArgs
	return nil
}
