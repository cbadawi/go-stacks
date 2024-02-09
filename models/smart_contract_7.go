package models

import (
	"encoding/json"
)

// SmartContract7 represents a SmartContract7 struct.
// A Smart Contract Detail
type SmartContract7 struct {
	TxId        string `json:"tx_id"`
	Canonical   bool   `json:"canonical"`
	ContractId  string `json:"contract_id"`
	BlockHeight int    `json:"block_height"`
	SourceCode  string `json:"source_code"`
	Abi         string `json:"abi"`
}

// MarshalJSON implements the json.Marshaler interface for SmartContract7.
// It customizes the JSON marshaling process for SmartContract7 objects.
func (s *SmartContract7) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the SmartContract7 object to a map representation for JSON marshaling.
func (s *SmartContract7) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["tx_id"] = s.TxId
	structMap["canonical"] = s.Canonical
	structMap["contract_id"] = s.ContractId
	structMap["block_height"] = s.BlockHeight
	structMap["source_code"] = s.SourceCode
	structMap["abi"] = s.Abi
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SmartContract7.
// It customizes the JSON unmarshaling process for SmartContract7 objects.
func (s *SmartContract7) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		TxId        string `json:"tx_id"`
		Canonical   bool   `json:"canonical"`
		ContractId  string `json:"contract_id"`
		BlockHeight int    `json:"block_height"`
		SourceCode  string `json:"source_code"`
		Abi         string `json:"abi"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	s.TxId = temp.TxId
	s.Canonical = temp.Canonical
	s.ContractId = temp.ContractId
	s.BlockHeight = temp.BlockHeight
	s.SourceCode = temp.SourceCode
	s.Abi = temp.Abi
	return nil
}
