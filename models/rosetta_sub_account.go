package models

import (
	"encoding/json"
)

// RosettaSubAccount represents a RosettaSubAccount struct.
// An account may have state specific to a contract address (ERC-20 token) and/or a stake (delegated balance). The sub_account_identifier should specify which state (if applicable) an account instantiation refers to.
type RosettaSubAccount struct {
	// The address may be a cryptographic public key (or some encoding of it) or a provided username.
	Address string `json:"address"`
	// If the SubAccount address is not sufficient to uniquely specify a SubAccount, any other identifying information can be stored here. It is important to note that two SubAccounts with identical addresses but differing metadata will not be considered equal by clients.
	Metadata *interface{} `json:"metadata,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for RosettaSubAccount.
// It customizes the JSON marshaling process for RosettaSubAccount objects.
func (r *RosettaSubAccount) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the RosettaSubAccount object to a map representation for JSON marshaling.
func (r *RosettaSubAccount) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["address"] = r.Address
	if r.Metadata != nil {
		structMap["metadata"] = r.Metadata
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RosettaSubAccount.
// It customizes the JSON unmarshaling process for RosettaSubAccount objects.
func (r *RosettaSubAccount) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Address  string       `json:"address"`
		Metadata *interface{} `json:"metadata,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	r.Address = temp.Address
	r.Metadata = temp.Metadata
	return nil
}
