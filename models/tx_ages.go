package models

import (
	"encoding/json"
)

// TxAges represents a TxAges struct.
// The average time (in blocks) that transactions have lived in the mempool. The start block height is simply the current chain-tip of when the attached Stacks node receives the transaction. This timing can be different across Stacks nodes / API instances due to propagation timing differences in the p2p network.
type TxAges struct {
	TokenTransfer    TokenTransfer2    `json:"token_transfer"`
	SmartContract    SmartContract2    `json:"smart_contract"`
	ContractCall     ContractCall2     `json:"contract_call"`
	PoisonMicroblock PoisonMicroblock1 `json:"poison_microblock"`
}

// MarshalJSON implements the json.Marshaler interface for TxAges.
// It customizes the JSON marshaling process for TxAges objects.
func (t *TxAges) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(t.toMap())
}

// toMap converts the TxAges object to a map representation for JSON marshaling.
func (t *TxAges) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["token_transfer"] = t.TokenTransfer
	structMap["smart_contract"] = t.SmartContract
	structMap["contract_call"] = t.ContractCall
	structMap["poison_microblock"] = t.PoisonMicroblock
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for TxAges.
// It customizes the JSON unmarshaling process for TxAges objects.
func (t *TxAges) UnmarshalJSON(input []byte) error {
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
