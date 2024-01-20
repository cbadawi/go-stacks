package models

import (
	"encoding/json"
)

// TxSimpleFeeAverages represents a TxSimpleFeeAverages struct.
// The simple mean (average) transaction fee, broken down by transaction type. Note that this does not factor in actual execution costs. The average fee is not a reliable metric for calculating a fee for a new transaction.
type TxSimpleFeeAverages struct {
	TokenTransfer    TokenTransfer2    `json:"token_transfer"`
	SmartContract    SmartContract2    `json:"smart_contract"`
	ContractCall     ContractCall2     `json:"contract_call"`
	PoisonMicroblock PoisonMicroblock1 `json:"poison_microblock"`
}

// MarshalJSON implements the json.Marshaler interface for TxSimpleFeeAverages.
// It customizes the JSON marshaling process for TxSimpleFeeAverages objects.
func (t *TxSimpleFeeAverages) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(t.toMap())
}

// toMap converts the TxSimpleFeeAverages object to a map representation for JSON marshaling.
func (t *TxSimpleFeeAverages) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["token_transfer"] = t.TokenTransfer
	structMap["smart_contract"] = t.SmartContract
	structMap["contract_call"] = t.ContractCall
	structMap["poison_microblock"] = t.PoisonMicroblock
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for TxSimpleFeeAverages.
// It customizes the JSON unmarshaling process for TxSimpleFeeAverages objects.
func (t *TxSimpleFeeAverages) UnmarshalJSON(input []byte) error {
	temp := &struct {
		TokenTransfer    TokenTransfer2    `json:"token_transfer"`
		SmartContract    SmartContract2    `json:"smart_contract"`
		ContractCall     ContractCall2     `json:"contract_call"`
		PoisonMicroblock PoisonMicroblock1 `json:"poison_microblock"`
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
