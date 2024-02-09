package models

import (
	"encoding/json"
)

// StxTransfer represents a StxTransfer struct.
type StxTransfer struct {
	// Amount transferred in micro-STX as an integer string.
	Amount string `json:"amount"`
	// Principal that sent STX. This is unspecified if the STX were minted.
	Sender *string `json:"sender,omitempty"`
	// Principal that received STX. This is unspecified if the STX were burned.
	Recipient *string `json:"recipient,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for StxTransfer.
// It customizes the JSON marshaling process for StxTransfer objects.
func (s *StxTransfer) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the StxTransfer object to a map representation for JSON marshaling.
func (s *StxTransfer) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["amount"] = s.Amount
	if s.Sender != nil {
		structMap["sender"] = s.Sender
	}
	if s.Recipient != nil {
		structMap["recipient"] = s.Recipient
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StxTransfer.
// It customizes the JSON unmarshaling process for StxTransfer objects.
func (s *StxTransfer) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		Amount    string  `json:"amount"`
		Sender    *string `json:"sender,omitempty"`
		Recipient *string `json:"recipient,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	s.Amount = temp.Amount
	s.Sender = temp.Sender
	s.Recipient = temp.Recipient
	return nil
}
