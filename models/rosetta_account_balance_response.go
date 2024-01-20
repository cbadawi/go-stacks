package models

import (
	"encoding/json"
)

// RosettaAccountBalanceResponse represents a RosettaAccountBalanceResponse struct.
// An AccountBalanceResponse is returned on the /account/balance endpoint. If an account has a balance for each AccountIdentifier describing it (ex: an ERC-20 token balance on a few smart contracts), an account balance request must be made with each AccountIdentifier.
type RosettaAccountBalanceResponse struct {
	// The block_identifier uniquely identifies a block in a particular network.
	BlockIdentifier RosettaBlockIdentifier `json:"block_identifier"`
	// A single account balance may have multiple currencies
	Balances []RosettaAmount `json:"balances"`
	// If a blockchain is UTXO-based, all unspent Coins owned by an account_identifier should be returned alongside the balance. It is highly recommended to populate this field so that users of the Rosetta API implementation don't need to maintain their own indexer to track their UTXOs.
	Coins []RosettaCoin `json:"coins,omitempty"`
	// Account-based blockchains that utilize a nonce or sequence number should include that number in the metadata. This number could be unique to the identifier or global across the account address.
	Metadata *Metadata2 `json:"metadata,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for RosettaAccountBalanceResponse.
// It customizes the JSON marshaling process for RosettaAccountBalanceResponse objects.
func (r *RosettaAccountBalanceResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the RosettaAccountBalanceResponse object to a map representation for JSON marshaling.
func (r *RosettaAccountBalanceResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["block_identifier"] = r.BlockIdentifier
	structMap["balances"] = r.Balances
	if r.Coins != nil {
		structMap["coins"] = r.Coins
	}
	if r.Metadata != nil {
		structMap["metadata"] = r.Metadata
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RosettaAccountBalanceResponse.
// It customizes the JSON unmarshaling process for RosettaAccountBalanceResponse objects.
func (r *RosettaAccountBalanceResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		BlockIdentifier RosettaBlockIdentifier `json:"block_identifier"`
		Balances        []RosettaAmount        `json:"balances"`
		Coins           []RosettaCoin          `json:"coins,omitempty"`
		Metadata        *Metadata2             `json:"metadata,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	r.BlockIdentifier = temp.BlockIdentifier
	r.Balances = temp.Balances
	r.Coins = temp.Coins
	r.Metadata = temp.Metadata
	return nil
}
