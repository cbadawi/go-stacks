package models

import (
	"encoding/json"
)

// AccountDataResponse represents a AccountDataResponse struct.
// GET request for account data
type AccountDataResponse struct {
	Balance      string `json:"balance"`
	Locked       string `json:"locked"`
	UnlockHeight int    `json:"unlock_height"`
	Nonce        int    `json:"nonce"`
	BalanceProof string `json:"balance_proof"`
	NonceProof   string `json:"nonce_proof"`
}

// MarshalJSON implements the json.Marshaler interface for AccountDataResponse.
// It customizes the JSON marshaling process for AccountDataResponse objects.
func (a *AccountDataResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(a.toMap())
}

// toMap converts the AccountDataResponse object to a map representation for JSON marshaling.
func (a *AccountDataResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["balance"] = a.Balance
	structMap["locked"] = a.Locked
	structMap["unlock_height"] = a.UnlockHeight
	structMap["nonce"] = a.Nonce
	structMap["balance_proof"] = a.BalanceProof
	structMap["nonce_proof"] = a.NonceProof
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AccountDataResponse.
// It customizes the JSON unmarshaling process for AccountDataResponse objects.
func (a *AccountDataResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Balance      string `json:"balance"`
		Locked       string `json:"locked"`
		UnlockHeight int    `json:"unlock_height"`
		Nonce        int    `json:"nonce"`
		BalanceProof string `json:"balance_proof"`
		NonceProof   string `json:"nonce_proof"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	a.Balance = temp.Balance
	a.Locked = temp.Locked
	a.UnlockHeight = temp.UnlockHeight
	a.Nonce = temp.Nonce
	a.BalanceProof = temp.BalanceProof
	a.NonceProof = temp.NonceProof
	return nil
}
