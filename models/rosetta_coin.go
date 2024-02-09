package models

import (
	"encoding/json"
)

// RosettaCoin represents a RosettaCoin struct.
// If a blockchain is UTXO-based, all unspent Coins owned by an account_identifier should be returned alongside the balance. It is highly recommended to populate this field so that users of the Rosetta API implementation don't need to maintain their own indexer to track their UTXOs.
type RosettaCoin struct {
	// CoinIdentifier uniquely identifies a Coin.
	CoinIdentifier CoinIdentifier `json:"coin_identifier"`
	// Amount is some Value of a Currency. It is considered invalid to specify a Value without a Currency.
	Amount RosettaAmount `json:"amount"`
}

// MarshalJSON implements the json.Marshaler interface for RosettaCoin.
// It customizes the JSON marshaling process for RosettaCoin objects.
func (r *RosettaCoin) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the RosettaCoin object to a map representation for JSON marshaling.
func (r *RosettaCoin) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["coin_identifier"] = r.CoinIdentifier
	structMap["amount"] = r.Amount
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RosettaCoin.
// It customizes the JSON unmarshaling process for RosettaCoin objects.
func (r *RosettaCoin) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		CoinIdentifier CoinIdentifier `json:"coin_identifier"`
		Amount         RosettaAmount  `json:"amount"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	r.CoinIdentifier = temp.CoinIdentifier
	r.Amount = temp.Amount
	return nil
}
