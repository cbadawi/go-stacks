package models

import (
	"encoding/json"
)

// RosettaStatusRequest represents a RosettaStatusRequest struct.
// This endpoint returns the current status of the network requested. Any NetworkIdentifier returned by /network/list should be accessible here.
type RosettaStatusRequest struct {
	// The network_identifier specifies which network a particular object is associated with.
	NetworkIdentifier NetworkIdentifier `json:"network_identifier"`
	Metadata          *interface{}      `json:"metadata,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for RosettaStatusRequest.
// It customizes the JSON marshaling process for RosettaStatusRequest objects.
func (r *RosettaStatusRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the RosettaStatusRequest object to a map representation for JSON marshaling.
func (r *RosettaStatusRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["network_identifier"] = r.NetworkIdentifier
	if r.Metadata != nil {
		structMap["metadata"] = r.Metadata
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RosettaStatusRequest.
// It customizes the JSON unmarshaling process for RosettaStatusRequest objects.
func (r *RosettaStatusRequest) UnmarshalJSON(input []byte) error {
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
