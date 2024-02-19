package models

import (
	"encoding/json"
)

// FtBalance represents a FtBalance struct.
type FtBalance struct {
	Balance       string `json:"balance"`
	TotalSent     string `json:"total_sent"`
	TotalReceived string `json:"total_received"`
}

// MarshalJSON implements the json.Marshaler interface for FtBalance.
// It customizes the JSON marshaling process for FtBalance objects.
func (f *FtBalance) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(f.toMap())
}

// toMap converts the FtBalance object to a map representation for JSON marshaling.
func (f *FtBalance) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["balance"] = f.Balance
	structMap["total_sent"] = f.TotalSent
	structMap["total_received"] = f.TotalReceived
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for FtBalance.
// It customizes the JSON unmarshaling process for FtBalance objects.
func (f *FtBalance) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		Balance       string `json:"balance"`
		TotalSent     string `json:"total_sent"`
		TotalReceived string `json:"total_received"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	f.Balance = temp.Balance
	f.TotalSent = temp.TotalSent
	f.TotalReceived = temp.TotalReceived
	return nil
}
