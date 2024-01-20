package models

import (
	"encoding/json"
)

// NftBalance represents a NftBalance struct.
type NftBalance struct {
	Count         string `json:"count"`
	TotalSent     string `json:"total_sent"`
	TotalReceived string `json:"total_received"`
}

// MarshalJSON implements the json.Marshaler interface for NftBalance.
// It customizes the JSON marshaling process for NftBalance objects.
func (n *NftBalance) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(n.toMap())
}

// toMap converts the NftBalance object to a map representation for JSON marshaling.
func (n *NftBalance) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["count"] = n.Count
	structMap["total_sent"] = n.TotalSent
	structMap["total_received"] = n.TotalReceived
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for NftBalance.
// It customizes the JSON unmarshaling process for NftBalance objects.
func (n *NftBalance) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Count         string `json:"count"`
		TotalSent     string `json:"total_sent"`
		TotalReceived string `json:"total_received"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	n.Count = temp.Count
	n.TotalSent = temp.TotalSent
	n.TotalReceived = temp.TotalReceived
	return nil
}
