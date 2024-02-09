package models

import (
	"encoding/json"
)

// Asset represents a Asset struct.
type Asset struct {
	AssetName       string `json:"asset_name"`
	ContractAddress string `json:"contract_address"`
	ContractName    string `json:"contract_name"`
}

// MarshalJSON implements the json.Marshaler interface for Asset.
// It customizes the JSON marshaling process for Asset objects.
func (a *Asset) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(a.toMap())
}

// toMap converts the Asset object to a map representation for JSON marshaling.
func (a *Asset) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["asset_name"] = a.AssetName
	structMap["contract_address"] = a.ContractAddress
	structMap["contract_name"] = a.ContractName
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Asset.
// It customizes the JSON unmarshaling process for Asset objects.
func (a *Asset) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		AssetName       string `json:"asset_name"`
		ContractAddress string `json:"contract_address"`
		ContractName    string `json:"contract_name"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	a.AssetName = temp.AssetName
	a.ContractAddress = temp.ContractAddress
	a.ContractName = temp.ContractName
	return nil
}
