package models

import (
	"encoding/json"
)

// RosettaConstructionMetadataRequest represents a RosettaConstructionMetadataRequest struct.
// A ConstructionMetadataRequest is utilized to get information required to construct a transaction. The Options object used to specify which metadata to return is left purposely unstructured to allow flexibility for implementers. Optionally, the request can also include an array of PublicKeys associated with the AccountIdentifiers returned in ConstructionPreprocessResponse.
type RosettaConstructionMetadataRequest struct {
	// The network_identifier specifies which network a particular object is associated with.
	NetworkIdentifier NetworkIdentifier `json:"network_identifier"`
	// The options that will be sent directly to /construction/metadata by the caller.
	Options    RosettaOptions     `json:"options"`
	PublicKeys []RosettaPublicKey `json:"public_keys,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for RosettaConstructionMetadataRequest.
// It customizes the JSON marshaling process for RosettaConstructionMetadataRequest objects.
func (r *RosettaConstructionMetadataRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the RosettaConstructionMetadataRequest object to a map representation for JSON marshaling.
func (r *RosettaConstructionMetadataRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["network_identifier"] = r.NetworkIdentifier
	structMap["options"] = r.Options
	if r.PublicKeys != nil {
		structMap["public_keys"] = r.PublicKeys
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RosettaConstructionMetadataRequest.
// It customizes the JSON unmarshaling process for RosettaConstructionMetadataRequest objects.
func (r *RosettaConstructionMetadataRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		NetworkIdentifier NetworkIdentifier  `json:"network_identifier"`
		Options           RosettaOptions     `json:"options"`
		PublicKeys        []RosettaPublicKey `json:"public_keys,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	r.NetworkIdentifier = temp.NetworkIdentifier
	r.Options = temp.Options
	r.PublicKeys = temp.PublicKeys
	return nil
}
