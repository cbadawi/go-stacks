package models

import (
	"encoding/json"
)

// AddressStxBalanceResponse represents a AddressStxBalanceResponse struct.
// GET request that returns address balances
type AddressStxBalanceResponse struct {
	Balance                   string `json:"balance"`
	TotalSent                 string `json:"total_sent"`
	TotalReceived             string `json:"total_received"`
	TotalFeesSent             string `json:"total_fees_sent"`
	TotalMinerRewardsReceived string `json:"total_miner_rewards_received"`
	// The transaction where the lock event occurred. Empty if no tokens are locked.
	LockTxId string `json:"lock_tx_id"`
	// The amount of locked STX, as string quoted micro-STX. Zero if no tokens are locked.
	Locked string `json:"locked"`
	// The STX chain block height of when the lock event occurred. Zero if no tokens are locked.
	LockHeight int `json:"lock_height"`
	// The burnchain block height of when the lock event occurred. Zero if no tokens are locked.
	BurnchainLockHeight int `json:"burnchain_lock_height"`
	// The burnchain block height of when the tokens unlock. Zero if no tokens are locked.
	BurnchainUnlockHeight int `json:"burnchain_unlock_height"`
	// Token Offering Locked
	TokenOfferingLocked *AddressTokenOfferingLocked `json:"token_offering_locked,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for AddressStxBalanceResponse.
// It customizes the JSON marshaling process for AddressStxBalanceResponse objects.
func (a *AddressStxBalanceResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(a.toMap())
}

// toMap converts the AddressStxBalanceResponse object to a map representation for JSON marshaling.
func (a *AddressStxBalanceResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["balance"] = a.Balance
	structMap["total_sent"] = a.TotalSent
	structMap["total_received"] = a.TotalReceived
	structMap["total_fees_sent"] = a.TotalFeesSent
	structMap["total_miner_rewards_received"] = a.TotalMinerRewardsReceived
	structMap["lock_tx_id"] = a.LockTxId
	structMap["locked"] = a.Locked
	structMap["lock_height"] = a.LockHeight
	structMap["burnchain_lock_height"] = a.BurnchainLockHeight
	structMap["burnchain_unlock_height"] = a.BurnchainUnlockHeight
	if a.TokenOfferingLocked != nil {
		structMap["token_offering_locked"] = a.TokenOfferingLocked
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AddressStxBalanceResponse.
// It customizes the JSON unmarshaling process for AddressStxBalanceResponse objects.
func (a *AddressStxBalanceResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Balance                   string                      `json:"balance"`
		TotalSent                 string                      `json:"total_sent"`
		TotalReceived             string                      `json:"total_received"`
		TotalFeesSent             string                      `json:"total_fees_sent"`
		TotalMinerRewardsReceived string                      `json:"total_miner_rewards_received"`
		LockTxId                  string                      `json:"lock_tx_id"`
		Locked                    string                      `json:"locked"`
		LockHeight                int                         `json:"lock_height"`
		BurnchainLockHeight       int                         `json:"burnchain_lock_height"`
		BurnchainUnlockHeight     int                         `json:"burnchain_unlock_height"`
		TokenOfferingLocked       *AddressTokenOfferingLocked `json:"token_offering_locked,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	a.Balance = temp.Balance
	a.TotalSent = temp.TotalSent
	a.TotalReceived = temp.TotalReceived
	a.TotalFeesSent = temp.TotalFeesSent
	a.TotalMinerRewardsReceived = temp.TotalMinerRewardsReceived
	a.LockTxId = temp.LockTxId
	a.Locked = temp.Locked
	a.LockHeight = temp.LockHeight
	a.BurnchainLockHeight = temp.BurnchainLockHeight
	a.BurnchainUnlockHeight = temp.BurnchainUnlockHeight
	a.TokenOfferingLocked = temp.TokenOfferingLocked
	return nil
}
