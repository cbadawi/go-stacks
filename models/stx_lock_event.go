package models

import (
	"encoding/json"
)

// StxLockEvent represents a StxLockEvent struct.
type StxLockEvent struct {
	LockedAmount  string `json:"locked_amount"`
	UnlockHeight  int    `json:"unlock_height"`
	LockedAddress string `json:"locked_address"`
}

// MarshalJSON implements the json.Marshaler interface for StxLockEvent.
// It customizes the JSON marshaling process for StxLockEvent objects.
func (s *StxLockEvent) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the StxLockEvent object to a map representation for JSON marshaling.
func (s *StxLockEvent) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["locked_amount"] = s.LockedAmount
	structMap["unlock_height"] = s.UnlockHeight
	structMap["locked_address"] = s.LockedAddress
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StxLockEvent.
// It customizes the JSON unmarshaling process for StxLockEvent objects.
func (s *StxLockEvent) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		LockedAmount  string `json:"locked_amount"`
		UnlockHeight  int    `json:"unlock_height"`
		LockedAddress string `json:"locked_address"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	s.LockedAmount = temp.LockedAmount
	s.UnlockHeight = temp.UnlockHeight
	s.LockedAddress = temp.LockedAddress
	return nil
}
