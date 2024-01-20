package models

import (
	"encoding/json"
)

// MempoolFeePriorities represents a MempoolFeePriorities struct.
// GET request that returns fee priorities from mempool transactions
type MempoolFeePriorities struct {
	All           All             `json:"all"`
	TokenTransfer *TokenTransfer1 `json:"token_transfer,omitempty"`
	SmartContract *SmartContract1 `json:"smart_contract,omitempty"`
	ContractCall  *ContractCall1  `json:"contract_call,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for MempoolFeePriorities.
// It customizes the JSON marshaling process for MempoolFeePriorities objects.
func (m *MempoolFeePriorities) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(m.toMap())
}

// toMap converts the MempoolFeePriorities object to a map representation for JSON marshaling.
func (m *MempoolFeePriorities) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["all"] = m.All
	if m.TokenTransfer != nil {
		structMap["token_transfer"] = m.TokenTransfer
	}
	if m.SmartContract != nil {
		structMap["smart_contract"] = m.SmartContract
	}
	if m.ContractCall != nil {
		structMap["contract_call"] = m.ContractCall
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MempoolFeePriorities.
// It customizes the JSON unmarshaling process for MempoolFeePriorities objects.
func (m *MempoolFeePriorities) UnmarshalJSON(input []byte) error {
	temp := &struct {
		All           All             `json:"all"`
		TokenTransfer *TokenTransfer1 `json:"token_transfer,omitempty"`
		SmartContract *SmartContract1 `json:"smart_contract,omitempty"`
		ContractCall  *ContractCall1  `json:"contract_call,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	m.All = temp.All
	m.TokenTransfer = temp.TokenTransfer
	m.SmartContract = temp.SmartContract
	m.ContractCall = temp.ContractCall
	return nil
}
