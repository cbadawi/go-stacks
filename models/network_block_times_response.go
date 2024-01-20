package models

import (
	"encoding/json"
)

// NetworkBlockTimesResponse represents a NetworkBlockTimesResponse struct.
// GET request that returns network target block times
type NetworkBlockTimesResponse struct {
	Mainnet TargetBlockTime `json:"mainnet"`
	Testnet TargetBlockTime `json:"testnet"`
}

// MarshalJSON implements the json.Marshaler interface for NetworkBlockTimesResponse.
// It customizes the JSON marshaling process for NetworkBlockTimesResponse objects.
func (n *NetworkBlockTimesResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(n.toMap())
}

// toMap converts the NetworkBlockTimesResponse object to a map representation for JSON marshaling.
func (n *NetworkBlockTimesResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["mainnet"] = n.Mainnet
	structMap["testnet"] = n.Testnet
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for NetworkBlockTimesResponse.
// It customizes the JSON unmarshaling process for NetworkBlockTimesResponse objects.
func (n *NetworkBlockTimesResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Mainnet TargetBlockTime `json:"mainnet"`
		Testnet TargetBlockTime `json:"testnet"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	n.Mainnet = temp.Mainnet
	n.Testnet = temp.Testnet
	return nil
}
