package models

import (
	"encoding/json"
)

// RosettaConstructionPreprocessResponse represents a RosettaConstructionPreprocessResponse struct.
// RosettaConstructionPreprocessResponse contains options that will be sent unmodified to /construction/metadata. If it is not necessary to make a request to /construction/metadata, options should be omitted. Some blockchains require the PublicKey of particular AccountIdentifiers to construct a valid transaction. To fetch these PublicKeys, populate required_public_keys with the AccountIdentifiers associated with the desired PublicKeys. If it is not necessary to retrieve any PublicKeys for construction, required_public_keys should be omitted.
type RosettaConstructionPreprocessResponse struct {
	// The options that will be sent directly to /construction/metadata by the caller.
	Options            *RosettaOptions  `json:"options,omitempty"`
	RequiredPublicKeys []RosettaAccount `json:"required_public_keys,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for RosettaConstructionPreprocessResponse.
// It customizes the JSON marshaling process for RosettaConstructionPreprocessResponse objects.
func (r *RosettaConstructionPreprocessResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the RosettaConstructionPreprocessResponse object to a map representation for JSON marshaling.
func (r *RosettaConstructionPreprocessResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if r.Options != nil {
		structMap["options"] = r.Options
	}
	if r.RequiredPublicKeys != nil {
		structMap["required_public_keys"] = r.RequiredPublicKeys
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RosettaConstructionPreprocessResponse.
// It customizes the JSON unmarshaling process for RosettaConstructionPreprocessResponse objects.
func (r *RosettaConstructionPreprocessResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Options            *RosettaOptions  `json:"options,omitempty"`
		RequiredPublicKeys []RosettaAccount `json:"required_public_keys,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	r.Options = temp.Options
	r.RequiredPublicKeys = temp.RequiredPublicKeys
	return nil
}
