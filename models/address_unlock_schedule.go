package models

import (
	"encoding/json"
)

// AddressUnlockSchedule represents a AddressUnlockSchedule struct.
// Unlock schedule amount and block height
type AddressUnlockSchedule struct {
	// Micro-STX amount locked at this block height.
	Amount      string  `json:"amount"`
	BlockHeight float64 `json:"block_height"`
}

// MarshalJSON implements the json.Marshaler interface for AddressUnlockSchedule.
// It customizes the JSON marshaling process for AddressUnlockSchedule objects.
func (a *AddressUnlockSchedule) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(a.toMap())
}

// toMap converts the AddressUnlockSchedule object to a map representation for JSON marshaling.
func (a *AddressUnlockSchedule) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["amount"] = a.Amount
	structMap["block_height"] = a.BlockHeight
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AddressUnlockSchedule.
// It customizes the JSON unmarshaling process for AddressUnlockSchedule objects.
func (a *AddressUnlockSchedule) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Amount      string  `json:"amount"`
		BlockHeight float64 `json:"block_height"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	a.Amount = temp.Amount
	a.BlockHeight = temp.BlockHeight
	return nil
}
