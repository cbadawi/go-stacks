package models

import (
	"encoding/json"
)

// RosettaNetworkListResponse represents a RosettaNetworkListResponse struct.
// A NetworkListResponse contains all NetworkIdentifiers that the node can serve information for.
type RosettaNetworkListResponse struct {
	// The network_identifier specifies which network a particular object is associated with.
	NetworkIdentifiers []NetworkIdentifier `json:"network_identifiers"`
}

// MarshalJSON implements the json.Marshaler interface for RosettaNetworkListResponse.
// It customizes the JSON marshaling process for RosettaNetworkListResponse objects.
func (r *RosettaNetworkListResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the RosettaNetworkListResponse object to a map representation for JSON marshaling.
func (r *RosettaNetworkListResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["network_identifiers"] = r.NetworkIdentifiers
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RosettaNetworkListResponse.
// It customizes the JSON unmarshaling process for RosettaNetworkListResponse objects.
func (r *RosettaNetworkListResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		NetworkIdentifiers []NetworkIdentifier `json:"network_identifiers"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	r.NetworkIdentifiers = temp.NetworkIdentifiers
	return nil
}
