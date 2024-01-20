package models

import (
	"encoding/json"
)

// RosettaMempoolRequest represents a RosettaMempoolRequest struct.
// Get all Transaction Identifiers in the mempool
type RosettaMempoolRequest struct {
	// The network_identifier specifies which network a particular object is associated with.
	NetworkIdentifier NetworkIdentifier `json:"network_identifier"`
	Metadata          *interface{}      `json:"metadata,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for RosettaMempoolRequest.
// It customizes the JSON marshaling process for RosettaMempoolRequest objects.
func (r *RosettaMempoolRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the RosettaMempoolRequest object to a map representation for JSON marshaling.
func (r *RosettaMempoolRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["network_identifier"] = r.NetworkIdentifier
	if r.Metadata != nil {
		structMap["metadata"] = r.Metadata
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RosettaMempoolRequest.
// It customizes the JSON unmarshaling process for RosettaMempoolRequest objects.
func (r *RosettaMempoolRequest) UnmarshalJSON(input []byte) error {
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
