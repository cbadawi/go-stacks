package models

import (
	"encoding/json"
)

// StxBalance represents a StxBalance struct.
type StxBalance struct {
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
}

// MarshalJSON implements the json.Marshaler interface for StxBalance.
// It customizes the JSON marshaling process for StxBalance objects.
func (s *StxBalance) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the StxBalance object to a map representation for JSON marshaling.
func (s *StxBalance) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["balance"] = s.Balance
	structMap["total_sent"] = s.TotalSent
	structMap["total_received"] = s.TotalReceived
	structMap["total_fees_sent"] = s.TotalFeesSent
	structMap["total_miner_rewards_received"] = s.TotalMinerRewardsReceived
	structMap["lock_tx_id"] = s.LockTxId
	structMap["locked"] = s.Locked
	structMap["lock_height"] = s.LockHeight
	structMap["burnchain_lock_height"] = s.BurnchainLockHeight
	structMap["burnchain_unlock_height"] = s.BurnchainUnlockHeight
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StxBalance.
// It customizes the JSON unmarshaling process for StxBalance objects.
func (s *StxBalance) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Balance                   string `json:"balance"`
		TotalSent                 string `json:"total_sent"`
		TotalReceived             string `json:"total_received"`
		TotalFeesSent             string `json:"total_fees_sent"`
		TotalMinerRewardsReceived string `json:"total_miner_rewards_received"`
		LockTxId                  string `json:"lock_tx_id"`
		Locked                    string `json:"locked"`
		LockHeight                int    `json:"lock_height"`
		BurnchainLockHeight       int    `json:"burnchain_lock_height"`
		BurnchainUnlockHeight     int    `json:"burnchain_unlock_height"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	s.Balance = temp.Balance
	s.TotalSent = temp.TotalSent
	s.TotalReceived = temp.TotalReceived
	s.TotalFeesSent = temp.TotalFeesSent
	s.TotalMinerRewardsReceived = temp.TotalMinerRewardsReceived
	s.LockTxId = temp.LockTxId
	s.Locked = temp.Locked
	s.LockHeight = temp.LockHeight
	s.BurnchainLockHeight = temp.BurnchainLockHeight
	s.BurnchainUnlockHeight = temp.BurnchainUnlockHeight
	return nil
}
