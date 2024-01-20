package models

import (
	"encoding/json"
)

// PostConditionPrincipal4 represents a PostConditionPrincipal4 struct.
type PostConditionPrincipal4 struct {
	// String literal of type `PostConditionPrincipalType`
	TypeId       string `json:"type_id"`
	Address      string `json:"address"`
	ContractName string `json:"contract_name"`
}

// MarshalJSON implements the json.Marshaler interface for PostConditionPrincipal4.
// It customizes the JSON marshaling process for PostConditionPrincipal4 objects.
func (p *PostConditionPrincipal4) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(p.toMap())
}

// toMap converts the PostConditionPrincipal4 object to a map representation for JSON marshaling.
func (p *PostConditionPrincipal4) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["type_id"] = p.TypeId
	structMap["address"] = p.Address
	structMap["contract_name"] = p.ContractName
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PostConditionPrincipal4.
// It customizes the JSON unmarshaling process for PostConditionPrincipal4 objects.
func (p *PostConditionPrincipal4) UnmarshalJSON(input []byte) error {
	temp := &struct {
		TypeId       string `json:"type_id"`
		Address      string `json:"address"`
		ContractName string `json:"contract_name"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	p.TypeId = temp.TypeId
	p.Address = temp.Address
	p.ContractName = temp.ContractName
	return nil
}
