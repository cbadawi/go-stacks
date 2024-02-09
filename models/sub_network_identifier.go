package models

import (
	"encoding/json"
)

// SubNetworkIdentifier represents a SubNetworkIdentifier struct.
// In blockchains with sharded state, the SubNetworkIdentifier is required to query some object on a specific shard. This identifier is optional for all non-sharded blockchains.
type SubNetworkIdentifier struct {
	// Network name
	Network string `json:"network"`
	// Meta data from subnetwork identifier
	Metadata *Metadata1 `json:"metadata,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for SubNetworkIdentifier.
// It customizes the JSON marshaling process for SubNetworkIdentifier objects.
func (s *SubNetworkIdentifier) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the SubNetworkIdentifier object to a map representation for JSON marshaling.
func (s *SubNetworkIdentifier) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["network"] = s.Network
	if s.Metadata != nil {
		structMap["metadata"] = s.Metadata
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubNetworkIdentifier.
// It customizes the JSON unmarshaling process for SubNetworkIdentifier objects.
func (s *SubNetworkIdentifier) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		Network  string     `json:"network"`
		Metadata *Metadata1 `json:"metadata,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	s.Network = temp.Network
	s.Metadata = temp.Metadata
	return nil
}
