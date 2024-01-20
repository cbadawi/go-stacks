package models

import (
	"encoding/json"
)

// TxTypeCounts represents a TxTypeCounts struct.
// Number of tranasction in the mempool, broken down by transaction type.
type TxTypeCounts struct {
	TokenTransfer    float64 `json:"token_transfer"`
	SmartContract    float64 `json:"smart_contract"`
	ContractCall     float64 `json:"contract_call"`
	PoisonMicroblock float64 `json:"poison_microblock"`
}

// MarshalJSON implements the json.Marshaler interface for TxTypeCounts.
// It customizes the JSON marshaling process for TxTypeCounts objects.
func (t *TxTypeCounts) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(t.toMap())
}

// toMap converts the TxTypeCounts object to a map representation for JSON marshaling.
func (t *TxTypeCounts) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["token_transfer"] = t.TokenTransfer
	structMap["smart_contract"] = t.SmartContract
	structMap["contract_call"] = t.ContractCall
	structMap["poison_microblock"] = t.PoisonMicroblock
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for TxTypeCounts.
// It customizes the JSON unmarshaling process for TxTypeCounts objects.
func (t *TxTypeCounts) UnmarshalJSON(input []byte) error {
	temp := &struct {
		TokenTransfer    float64 `json:"token_transfer"`
		SmartContract    float64 `json:"smart_contract"`
		ContractCall     float64 `json:"contract_call"`
		PoisonMicroblock float64 `json:"poison_microblock"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	t.TokenTransfer = temp.TokenTransfer
	t.SmartContract = temp.SmartContract
	t.ContractCall = temp.ContractCall
	t.PoisonMicroblock = temp.PoisonMicroblock
	return nil
}
