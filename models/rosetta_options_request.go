package models

import (
	"encoding/json"
)

// RosettaOptionsRequest represents a RosettaOptionsRequest struct.
// This endpoint returns the version information and allowed network-specific types for a NetworkIdentifier. Any NetworkIdentifier returned by /network/list should be accessible here. Because options are retrievable in the context of a NetworkIdentifier, it is possible to define unique options for each network.
type RosettaOptionsRequest struct {
	// The network_identifier specifies which network a particular object is associated with.
	NetworkIdentifier NetworkIdentifier `json:"network_identifier"`
	Metadata          *interface{}      `json:"metadata,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for RosettaOptionsRequest.
// It customizes the JSON marshaling process for RosettaOptionsRequest objects.
func (r *RosettaOptionsRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the RosettaOptionsRequest object to a map representation for JSON marshaling.
func (r *RosettaOptionsRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["network_identifier"] = r.NetworkIdentifier
	if r.Metadata != nil {
		structMap["metadata"] = r.Metadata
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RosettaOptionsRequest.
// It customizes the JSON unmarshaling process for RosettaOptionsRequest objects.
func (r *RosettaOptionsRequest) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		NetworkIdentifier NetworkIdentifier `json:"network_identifier"`
		Metadata          *interface{}      `json:"metadata,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	r.NetworkIdentifier = temp.NetworkIdentifier
	r.Metadata = temp.Metadata
	return nil
}
