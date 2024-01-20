package models

import (
	"encoding/json"
)

// SmartContract represents a SmartContract struct.
type SmartContract struct {
	// The Clarity version of the contract, only specified for versioned contract transactions, otherwise null
	ClarityVersion Optional[float64] `json:"clarity_version"`
	// Contract identifier formatted as `<principaladdress>.<contract_name>`
	ContractId string `json:"contract_id"`
	// Clarity code of the smart contract being deployed
	SourceCode string `json:"source_code"`
}

// MarshalJSON implements the json.Marshaler interface for SmartContract.
// It customizes the JSON marshaling process for SmartContract objects.
func (s *SmartContract) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the SmartContract object to a map representation for JSON marshaling.
func (s *SmartContract) toMap() map[string]any {
	structMap := make(map[string]any)
	if s.ClarityVersion.IsValueSet() {
		structMap["clarity_version"] = s.ClarityVersion.Value()
	}
	structMap["contract_id"] = s.ContractId
	structMap["source_code"] = s.SourceCode
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SmartContract.
// It customizes the JSON unmarshaling process for SmartContract objects.
func (s *SmartContract) UnmarshalJSON(input []byte) error {
	temp := &struct {
		ClarityVersion Optional[float64] `json:"clarity_version"`
		ContractId     string            `json:"contract_id"`
		SourceCode     string            `json:"source_code"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	s.ClarityVersion = temp.ClarityVersion
	s.ContractId = temp.ContractId
	s.SourceCode = temp.SourceCode
	return nil
}
