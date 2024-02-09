package models

import (
	"encoding/json"
)

// RosettaConstructionPayloadsRequest represents a RosettaConstructionPayloadsRequest struct.
// ConstructionPayloadsRequest is the request to /construction/payloads. It contains the network, a slice of operations, and arbitrary metadata that was returned by the call to /construction/metadata. Optionally, the request can also include an array of PublicKeys associated with the AccountIdentifiers returned in ConstructionPreprocessResponse.
type RosettaConstructionPayloadsRequest struct {
	// The network_identifier specifies which network a particular object is associated with.
	NetworkIdentifier NetworkIdentifier  `json:"network_identifier"`
	Operations        []RosettaOperation `json:"operations"`
	PublicKeys        []RosettaPublicKey `json:"public_keys,omitempty"`
	Metadata          *Metadata5         `json:"metadata,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for RosettaConstructionPayloadsRequest.
// It customizes the JSON marshaling process for RosettaConstructionPayloadsRequest objects.
func (r *RosettaConstructionPayloadsRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the RosettaConstructionPayloadsRequest object to a map representation for JSON marshaling.
func (r *RosettaConstructionPayloadsRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["network_identifier"] = r.NetworkIdentifier
	structMap["operations"] = r.Operations
	if r.PublicKeys != nil {
		structMap["public_keys"] = r.PublicKeys
	}
	if r.Metadata != nil {
		structMap["metadata"] = r.Metadata
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RosettaConstructionPayloadsRequest.
// It customizes the JSON unmarshaling process for RosettaConstructionPayloadsRequest objects.
func (r *RosettaConstructionPayloadsRequest) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		NetworkIdentifier NetworkIdentifier  `json:"network_identifier"`
		Operations        []RosettaOperation `json:"operations"`
		PublicKeys        []RosettaPublicKey `json:"public_keys,omitempty"`
		Metadata          *Metadata5         `json:"metadata,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	r.NetworkIdentifier = temp.NetworkIdentifier
	r.Operations = temp.Operations
	r.PublicKeys = temp.PublicKeys
	r.Metadata = temp.Metadata
	return nil
}
