package models

import (
	"encoding/json"
)

// RosettaConstructionMetadataResponse represents a RosettaConstructionMetadataResponse struct.
// The ConstructionMetadataResponse returns network-specific metadata used for transaction construction. Optionally, the implementer can return the suggested fee associated with the transaction being constructed. The caller may use this info to adjust the intent of the transaction or to create a transaction with a different account that can pay the suggested fee. Suggested fee is an array in case fee payment must occur in multiple currencies.
type RosettaConstructionMetadataResponse struct {
	Metadata     Metadata5       `json:"metadata"`
	SuggestedFee []RosettaAmount `json:"suggested_fee,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for RosettaConstructionMetadataResponse.
// It customizes the JSON marshaling process for RosettaConstructionMetadataResponse objects.
func (r *RosettaConstructionMetadataResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the RosettaConstructionMetadataResponse object to a map representation for JSON marshaling.
func (r *RosettaConstructionMetadataResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["metadata"] = r.Metadata
	if r.SuggestedFee != nil {
		structMap["suggested_fee"] = r.SuggestedFee
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RosettaConstructionMetadataResponse.
// It customizes the JSON unmarshaling process for RosettaConstructionMetadataResponse objects.
func (r *RosettaConstructionMetadataResponse) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		Metadata     Metadata5       `json:"metadata"`
		SuggestedFee []RosettaAmount `json:"suggested_fee,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	r.Metadata = temp.Metadata
	r.SuggestedFee = temp.SuggestedFee
	return nil
}
