package models

import (
	"encoding/json"
)

// ExtendedV1FaucetsBtcRequest represents a ExtendedV1FaucetsBtcRequest struct.
type ExtendedV1FaucetsBtcRequest struct {
	// BTC testnet address
	Address *string `json:"address,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for ExtendedV1FaucetsBtcRequest.
// It customizes the JSON marshaling process for ExtendedV1FaucetsBtcRequest objects.
func (e *ExtendedV1FaucetsBtcRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(e.toMap())
}

// toMap converts the ExtendedV1FaucetsBtcRequest object to a map representation for JSON marshaling.
func (e *ExtendedV1FaucetsBtcRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	if e.Address != nil {
		structMap["address"] = e.Address
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ExtendedV1FaucetsBtcRequest.
// It customizes the JSON unmarshaling process for ExtendedV1FaucetsBtcRequest objects.
func (e *ExtendedV1FaucetsBtcRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Address *string `json:"address,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	e.Address = temp.Address
	return nil
}
