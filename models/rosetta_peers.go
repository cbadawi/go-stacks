package models

import (
	"encoding/json"
)

// RosettaPeers represents a RosettaPeers struct.
// A Peer is a representation of a node's peer.
type RosettaPeers struct {
	// peer id
	PeerId string `json:"peer_id"`
	// meta data
	Metadata *interface{} `json:"metadata,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for RosettaPeers.
// It customizes the JSON marshaling process for RosettaPeers objects.
func (r *RosettaPeers) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the RosettaPeers object to a map representation for JSON marshaling.
func (r *RosettaPeers) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["peer_id"] = r.PeerId
	if r.Metadata != nil {
		structMap["metadata"] = r.Metadata
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RosettaPeers.
// It customizes the JSON unmarshaling process for RosettaPeers objects.
func (r *RosettaPeers) UnmarshalJSON(input []byte) error {
	temp := &struct {
		PeerId   string       `json:"peer_id"`
		Metadata *interface{} `json:"metadata,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	r.PeerId = temp.PeerId
	r.Metadata = temp.Metadata
	return nil
}
