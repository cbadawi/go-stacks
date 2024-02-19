package models

import (
	"encoding/json"
)

// RosettaConstructionDeriveRequest represents a RosettaConstructionDeriveRequest struct.
// Network is provided in the request because some blockchains have different address formats for different networks
type RosettaConstructionDeriveRequest struct {
	// The network_identifier specifies which network a particular object is associated with.
	NetworkIdentifier NetworkIdentifier `json:"network_identifier"`
	// PublicKey contains a public key byte array for a particular CurveType encoded in hex. Note that there is no PrivateKey struct as this is NEVER the concern of an implementation.
	PublicKey RosettaPublicKey `json:"public_key"`
	Metadata  *interface{}     `json:"metadata,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for RosettaConstructionDeriveRequest.
// It customizes the JSON marshaling process for RosettaConstructionDeriveRequest objects.
func (r *RosettaConstructionDeriveRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the RosettaConstructionDeriveRequest object to a map representation for JSON marshaling.
func (r *RosettaConstructionDeriveRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["network_identifier"] = r.NetworkIdentifier
	structMap["public_key"] = r.PublicKey
	if r.Metadata != nil {
		structMap["metadata"] = r.Metadata
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RosettaConstructionDeriveRequest.
// It customizes the JSON unmarshaling process for RosettaConstructionDeriveRequest objects.
func (r *RosettaConstructionDeriveRequest) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		NetworkIdentifier NetworkIdentifier `json:"network_identifier"`
		PublicKey         RosettaPublicKey  `json:"public_key"`
		Metadata          *interface{}      `json:"metadata,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	r.NetworkIdentifier = temp.NetworkIdentifier
	r.PublicKey = temp.PublicKey
	r.Metadata = temp.Metadata
	return nil
}
