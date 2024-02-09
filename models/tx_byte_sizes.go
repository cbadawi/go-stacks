package models

import (
	"encoding/json"
)

// TxByteSizes represents a TxByteSizes struct.
// The average byte size of transactions in the mempool, broken down by transaction type.
type TxByteSizes struct {
	TokenTransfer    TokenTransfer2    `json:"token_transfer"`
	SmartContract    SmartContract2    `json:"smart_contract"`
	ContractCall     ContractCall2     `json:"contract_call"`
	PoisonMicroblock PoisonMicroblock1 `json:"poison_microblock"`
}

// MarshalJSON implements the json.Marshaler interface for TxByteSizes.
// It customizes the JSON marshaling process for TxByteSizes objects.
func (t *TxByteSizes) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(t.toMap())
}

// toMap converts the TxByteSizes object to a map representation for JSON marshaling.
func (t *TxByteSizes) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["token_transfer"] = t.TokenTransfer
	structMap["smart_contract"] = t.SmartContract
	structMap["contract_call"] = t.ContractCall
	structMap["poison_microblock"] = t.PoisonMicroblock
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for TxByteSizes.
// It customizes the JSON unmarshaling process for TxByteSizes objects.
func (t *TxByteSizes) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
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
