package models

import (
	"encoding/json"
)

// RosettaCoinChange represents a RosettaCoinChange struct.
// CoinChange is used to represent a change in state of a some coin identified by a coin_identifier. This object is part of the Operation model and must be populated for UTXO-based blockchains. Coincidentally, this abstraction of UTXOs allows for supporting both account-based transfers and UTXO-based transfers on the same blockchain (when a transfer is account-based, don't populate this model).
type RosettaCoinChange struct {
	// CoinIdentifier uniquely identifies a Coin.
	CoinIdentifier CoinIdentifier `json:"coin_identifier"`
	// CoinActions are different state changes that a Coin can undergo. When a Coin is created, it is coin_created. When a Coin is spent, it is coin_spent. It is assumed that a single Coin cannot be created or spent more than once.
	CoinAction CoinActionEnum `json:"coin_action"`
}

// MarshalJSON implements the json.Marshaler interface for RosettaCoinChange.
// It customizes the JSON marshaling process for RosettaCoinChange objects.
func (r *RosettaCoinChange) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the RosettaCoinChange object to a map representation for JSON marshaling.
func (r *RosettaCoinChange) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["coin_identifier"] = r.CoinIdentifier
	structMap["coin_action"] = r.CoinAction
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RosettaCoinChange.
// It customizes the JSON unmarshaling process for RosettaCoinChange objects.
func (r *RosettaCoinChange) UnmarshalJSON(input []byte) error {
	temp := &struct {
		CoinIdentifier CoinIdentifier `json:"coin_identifier"`
		CoinAction     CoinActionEnum `json:"coin_action"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	r.CoinIdentifier = temp.CoinIdentifier
	r.CoinAction = temp.CoinAction
	return nil
}
