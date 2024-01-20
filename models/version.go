package models

import (
	"encoding/json"
)

// Version represents a Version struct.
// The Version object is utilized to inform the client of the versions of different components of the Rosetta implementation.
type Version struct {
	// The rosetta_version is the version of the Rosetta interface the implementation adheres to. This can be useful for clients looking to reliably parse responses.
	RosettaVersion string `json:"rosetta_version"`
	// The node_version is the canonical version of the node runtime. This can help clients manage deployments.
	NodeVersion string `json:"node_version"`
	// When a middleware server is used to adhere to the Rosetta interface, it should return its version here. This can help clients manage deployments.
	MiddlewareVersion *string `json:"middleware_version,omitempty"`
	// Any other information that may be useful about versioning of dependent services should be returned here.
	Metadata *interface{} `json:"metadata,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for Version.
// It customizes the JSON marshaling process for Version objects.
func (v *Version) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(v.toMap())
}

// toMap converts the Version object to a map representation for JSON marshaling.
func (v *Version) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["rosetta_version"] = v.RosettaVersion
	structMap["node_version"] = v.NodeVersion
	if v.MiddlewareVersion != nil {
		structMap["middleware_version"] = v.MiddlewareVersion
	}
	if v.Metadata != nil {
		structMap["metadata"] = v.Metadata
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Version.
// It customizes the JSON unmarshaling process for Version objects.
func (v *Version) UnmarshalJSON(input []byte) error {
	temp := &struct {
		RosettaVersion    string       `json:"rosetta_version"`
		NodeVersion       string       `json:"node_version"`
		MiddlewareVersion *string      `json:"middleware_version,omitempty"`
		Metadata          *interface{} `json:"metadata,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	v.RosettaVersion = temp.RosettaVersion
	v.NodeVersion = temp.NodeVersion
	v.MiddlewareVersion = temp.MiddlewareVersion
	v.Metadata = temp.Metadata
	return nil
}
