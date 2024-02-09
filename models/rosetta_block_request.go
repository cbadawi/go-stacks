package models

import (
	"encoding/json"
)

// RosettaBlockRequest represents a RosettaBlockRequest struct.
// A BlockRequest is utilized to make a block request on the /block endpoint.
type RosettaBlockRequest struct {
	// The network_identifier specifies which network a particular object is associated with.
	NetworkIdentifier NetworkIdentifier `json:"network_identifier"`
	// When fetching data by BlockIdentifier, it may be possible to only specify the index or hash. If neither property is specified, it is assumed that the client is making a request at the current block.
	BlockIdentifier RosettaPartialBlockIdentifier `json:"block_identifier"`
}

// MarshalJSON implements the json.Marshaler interface for RosettaBlockRequest.
// It customizes the JSON marshaling process for RosettaBlockRequest objects.
func (r *RosettaBlockRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the RosettaBlockRequest object to a map representation for JSON marshaling.
func (r *RosettaBlockRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["network_identifier"] = r.NetworkIdentifier
	structMap["block_identifier"] = r.BlockIdentifier
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RosettaBlockRequest.
// It customizes the JSON unmarshaling process for RosettaBlockRequest objects.
func (r *RosettaBlockRequest) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		NetworkIdentifier NetworkIdentifier             `json:"network_identifier"`
		BlockIdentifier   RosettaPartialBlockIdentifier `json:"block_identifier"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	r.NetworkIdentifier = temp.NetworkIdentifier
	r.BlockIdentifier = temp.BlockIdentifier
	return nil
}
