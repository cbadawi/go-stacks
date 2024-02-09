package models

import (
	"encoding/json"
)

// AddressBalanceResponse represents a AddressBalanceResponse struct.
// GET request that returns address balances
type AddressBalanceResponse struct {
	Stx               StxBalance            `json:"stx"`
	FungibleTokens    map[string]FtBalance  `json:"fungible_tokens"`
	NonFungibleTokens map[string]NftBalance `json:"non_fungible_tokens"`
	// Token Offering Locked
	TokenOfferingLocked *AddressTokenOfferingLocked `json:"token_offering_locked,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for AddressBalanceResponse.
// It customizes the JSON marshaling process for AddressBalanceResponse objects.
func (a *AddressBalanceResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(a.toMap())
}

// toMap converts the AddressBalanceResponse object to a map representation for JSON marshaling.
func (a *AddressBalanceResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["stx"] = a.Stx
	structMap["fungible_tokens"] = a.FungibleTokens
	structMap["non_fungible_tokens"] = a.NonFungibleTokens
	if a.TokenOfferingLocked != nil {
		structMap["token_offering_locked"] = a.TokenOfferingLocked
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AddressBalanceResponse.
// It customizes the JSON unmarshaling process for AddressBalanceResponse objects.
func (a *AddressBalanceResponse) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		Stx                 StxBalance                  `json:"stx"`
		FungibleTokens      map[string]FtBalance        `json:"fungible_tokens"`
		NonFungibleTokens   map[string]NftBalance       `json:"non_fungible_tokens"`
		TokenOfferingLocked *AddressTokenOfferingLocked `json:"token_offering_locked,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	a.Stx = temp.Stx
	a.FungibleTokens = temp.FungibleTokens
	a.NonFungibleTokens = temp.NonFungibleTokens
	a.TokenOfferingLocked = temp.TokenOfferingLocked
	return nil
}
