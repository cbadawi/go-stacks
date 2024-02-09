package models

import (
	"encoding/json"
)

// RosettaNetworkOptionsResponse represents a RosettaNetworkOptionsResponse struct.
// NetworkOptionsResponse contains information about the versioning of the node and the allowed operation statuses, operation types, and errors.
type RosettaNetworkOptionsResponse struct {
	// The Version object is utilized to inform the client of the versions of different components of the Rosetta implementation.
	Version Version `json:"version"`
	// Allow specifies supported Operation status, Operation types, and all possible error statuses. This Allow object is used by clients to validate the correctness of a Rosetta Server implementation. It is expected that these clients will error if they receive some response that contains any of the above information that is not specified here.
	Allow Allow `json:"allow"`
}

// MarshalJSON implements the json.Marshaler interface for RosettaNetworkOptionsResponse.
// It customizes the JSON marshaling process for RosettaNetworkOptionsResponse objects.
func (r *RosettaNetworkOptionsResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the RosettaNetworkOptionsResponse object to a map representation for JSON marshaling.
func (r *RosettaNetworkOptionsResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["version"] = r.Version
	structMap["allow"] = r.Allow
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RosettaNetworkOptionsResponse.
// It customizes the JSON unmarshaling process for RosettaNetworkOptionsResponse objects.
func (r *RosettaNetworkOptionsResponse) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		Version Version `json:"version"`
		Allow   Allow   `json:"allow"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	r.Version = temp.Version
	r.Allow = temp.Allow
	return nil
}
