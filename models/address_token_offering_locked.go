package models

import (
	"encoding/json"
)

// AddressTokenOfferingLocked represents a AddressTokenOfferingLocked struct.
// Token Offering Locked
type AddressTokenOfferingLocked struct {
	// Micro-STX amount still locked at current block height.
	TotalLocked string `json:"total_locked"`
	// Micro-STX amount unlocked at current block height.
	TotalUnlocked  string                  `json:"total_unlocked"`
	UnlockSchedule []AddressUnlockSchedule `json:"unlock_schedule"`
}

// MarshalJSON implements the json.Marshaler interface for AddressTokenOfferingLocked.
// It customizes the JSON marshaling process for AddressTokenOfferingLocked objects.
func (a *AddressTokenOfferingLocked) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(a.toMap())
}

// toMap converts the AddressTokenOfferingLocked object to a map representation for JSON marshaling.
func (a *AddressTokenOfferingLocked) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["total_locked"] = a.TotalLocked
	structMap["total_unlocked"] = a.TotalUnlocked
	structMap["unlock_schedule"] = a.UnlockSchedule
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AddressTokenOfferingLocked.
// It customizes the JSON unmarshaling process for AddressTokenOfferingLocked objects.
func (a *AddressTokenOfferingLocked) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		TotalLocked    string                  `json:"total_locked"`
		TotalUnlocked  string                  `json:"total_unlocked"`
		UnlockSchedule []AddressUnlockSchedule `json:"unlock_schedule"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	a.TotalLocked = temp.TotalLocked
	a.TotalUnlocked = temp.TotalUnlocked
	a.UnlockSchedule = temp.UnlockSchedule
	return nil
}
