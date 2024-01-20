package models

import (
	"encoding/json"
)

// CoinbasePayload represents a CoinbasePayload struct.
type CoinbasePayload struct {
	// Hex encoded 32-byte scratch space for block leader's use
	Data string `json:"data"`
	// A principal that will receive the miner rewards for this coinbase transaction. Can be either a standard principal or contract principal. Only specified for `coinbase-to-alt-recipient` transaction types, otherwise null.
	AltRecipient Optional[string] `json:"alt_recipient"`
	// Hex encoded 80-byte VRF proof
	VrfProof Optional[string] `json:"vrf_proof"`
}

// MarshalJSON implements the json.Marshaler interface for CoinbasePayload.
// It customizes the JSON marshaling process for CoinbasePayload objects.
func (c *CoinbasePayload) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CoinbasePayload object to a map representation for JSON marshaling.
func (c *CoinbasePayload) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["data"] = c.Data
	if c.AltRecipient.IsValueSet() {
		structMap["alt_recipient"] = c.AltRecipient.Value()
	}
	if c.VrfProof.IsValueSet() {
		structMap["vrf_proof"] = c.VrfProof.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CoinbasePayload.
// It customizes the JSON unmarshaling process for CoinbasePayload objects.
func (c *CoinbasePayload) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Data         string           `json:"data"`
		AltRecipient Optional[string] `json:"alt_recipient"`
		VrfProof     Optional[string] `json:"vrf_proof"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Data = temp.Data
	c.AltRecipient = temp.AltRecipient
	c.VrfProof = temp.VrfProof
	return nil
}
