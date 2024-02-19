package models

import (
	"encoding/json"
)

// SigningPayload represents a SigningPayload struct.
// SigningPayload is signed by the client with the keypair associated with an address using the specified SignatureType. SignatureType can be optionally populated if there is a restriction on the signature scheme that can be used to sign the payload.
type SigningPayload struct {
	// [DEPRECATED by account_identifier in v1.4.4] The network-specific address of the account that should sign the payload.
	Address *string `json:"address,omitempty"`
	// The account_identifier uniquely identifies an account within a network. All fields in the account_identifier are utilized to determine this uniqueness (including the metadata field, if populated).
	AccountIdentifier *RosettaAccount `json:"account_identifier,omitempty"`
	HexBytes          string          `json:"hex_bytes"`
	// SignatureType is the type of a cryptographic signature.
	SignatureType *SignatureTypeEnum `json:"signature_type,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for SigningPayload.
// It customizes the JSON marshaling process for SigningPayload objects.
func (s *SigningPayload) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the SigningPayload object to a map representation for JSON marshaling.
func (s *SigningPayload) toMap() map[string]any {
	structMap := make(map[string]any)
	if s.Address != nil {
		structMap["address"] = s.Address
	}
	if s.AccountIdentifier != nil {
		structMap["account_identifier"] = s.AccountIdentifier
	}
	structMap["hex_bytes"] = s.HexBytes
	if s.SignatureType != nil {
		structMap["signature_type"] = s.SignatureType
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SigningPayload.
// It customizes the JSON unmarshaling process for SigningPayload objects.
func (s *SigningPayload) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		Address           *string            `json:"address,omitempty"`
		AccountIdentifier *RosettaAccount    `json:"account_identifier,omitempty"`
		HexBytes          string             `json:"hex_bytes"`
		SignatureType     *SignatureTypeEnum `json:"signature_type,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	s.Address = temp.Address
	s.AccountIdentifier = temp.AccountIdentifier
	s.HexBytes = temp.HexBytes
	s.SignatureType = temp.SignatureType
	return nil
}
