package models

import (
	"encoding/json"
)

// RosettaAccountIdentifier represents a RosettaAccountIdentifier struct.
// The account_identifier uniquely identifies an account within a network. All fields in the account_identifier are utilized to determine this uniqueness (including the metadata field, if populated).
type RosettaAccountIdentifier struct {
	// The address may be a cryptographic public key (or some encoding of it) or a provided username.
	Address string `json:"address"`
	// An account may have state specific to a contract address (ERC-20 token) and/or a stake (delegated balance). The sub_account_identifier should specify which state (if applicable) an account instantiation refers to.
	SubAccount *RosettaSubAccount `json:"sub_account,omitempty"`
	// Blockchains that utilize a username model (where the address is not a derivative of a cryptographic public key) should specify the public key(s) owned by the address in metadata.
	Metadata *interface{} `json:"metadata,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for RosettaAccountIdentifier.
// It customizes the JSON marshaling process for RosettaAccountIdentifier objects.
func (r *RosettaAccountIdentifier) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the RosettaAccountIdentifier object to a map representation for JSON marshaling.
func (r *RosettaAccountIdentifier) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["address"] = r.Address
	if r.SubAccount != nil {
		structMap["sub_account"] = r.SubAccount
	}
	if r.Metadata != nil {
		structMap["metadata"] = r.Metadata
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RosettaAccountIdentifier.
// It customizes the JSON unmarshaling process for RosettaAccountIdentifier objects.
func (r *RosettaAccountIdentifier) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Address    string             `json:"address"`
		SubAccount *RosettaSubAccount `json:"sub_account,omitempty"`
		Metadata   *interface{}       `json:"metadata,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	r.Address = temp.Address
	r.SubAccount = temp.SubAccount
	r.Metadata = temp.Metadata
	return nil
}
