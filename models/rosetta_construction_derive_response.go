package models

import (
	"encoding/json"
)

// RosettaConstructionDeriveResponse represents a RosettaConstructionDeriveResponse struct.
// ConstructionDeriveResponse is returned by the /construction/derive endpoint.
type RosettaConstructionDeriveResponse struct {
	// [DEPRECATED by account_identifier in v1.4.4] Address in network-specific format.
	Address *string `json:"address,omitempty"`
	// The account_identifier uniquely identifies an account within a network. All fields in the account_identifier are utilized to determine this uniqueness (including the metadata field, if populated).
	AccountIdentifier *RosettaAccountIdentifier `json:"account_identifier,omitempty"`
	Metadata          *interface{}              `json:"metadata,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for RosettaConstructionDeriveResponse.
// It customizes the JSON marshaling process for RosettaConstructionDeriveResponse objects.
func (r *RosettaConstructionDeriveResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the RosettaConstructionDeriveResponse object to a map representation for JSON marshaling.
func (r *RosettaConstructionDeriveResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if r.Address != nil {
		structMap["address"] = r.Address
	}
	if r.AccountIdentifier != nil {
		structMap["account_identifier"] = r.AccountIdentifier
	}
	if r.Metadata != nil {
		structMap["metadata"] = r.Metadata
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RosettaConstructionDeriveResponse.
// It customizes the JSON unmarshaling process for RosettaConstructionDeriveResponse objects.
func (r *RosettaConstructionDeriveResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Address           *string                   `json:"address,omitempty"`
		AccountIdentifier *RosettaAccountIdentifier `json:"account_identifier,omitempty"`
		Metadata          *interface{}              `json:"metadata,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	r.Address = temp.Address
	r.AccountIdentifier = temp.AccountIdentifier
	r.Metadata = temp.Metadata
	return nil
}
