package models

import (
	"encoding/json"
)

// BnsGetAllNamespacesResponse represents a BnsGetAllNamespacesResponse struct.
// Fetch a list of all namespaces known to the node.
type BnsGetAllNamespacesResponse struct {
	Namespaces []string `json:"namespaces"`
}

// MarshalJSON implements the json.Marshaler interface for BnsGetAllNamespacesResponse.
// It customizes the JSON marshaling process for BnsGetAllNamespacesResponse objects.
func (b *BnsGetAllNamespacesResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(b.toMap())
}

// toMap converts the BnsGetAllNamespacesResponse object to a map representation for JSON marshaling.
func (b *BnsGetAllNamespacesResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["namespaces"] = b.Namespaces
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for BnsGetAllNamespacesResponse.
// It customizes the JSON unmarshaling process for BnsGetAllNamespacesResponse objects.
func (b *BnsGetAllNamespacesResponse) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		Namespaces []string `json:"namespaces"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	b.Namespaces = temp.Namespaces
	return nil
}
