package models

import (
	"encoding/json"
)

// RosettaConstructionParseResponse represents a RosettaConstructionParseResponse struct.
// RosettaConstructionParseResponse contains an array of operations that occur in a transaction blob. This should match the array of operations provided to /construction/preprocess and /construction/payloads.
type RosettaConstructionParseResponse struct {
	Operations []RosettaOperation `json:"operations"`
	// [DEPRECATED by account_identifier_signers in v1.4.4] All signers (addresses) of a particular transaction. If the transaction is unsigned, it should be empty.
	Signers                  []string                   `json:"signers,omitempty"`
	AccountIdentifierSigners []RosettaAccountIdentifier `json:"account_identifier_signers,omitempty"`
	Metadata                 *interface{}               `json:"metadata,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for RosettaConstructionParseResponse.
// It customizes the JSON marshaling process for RosettaConstructionParseResponse objects.
func (r *RosettaConstructionParseResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the RosettaConstructionParseResponse object to a map representation for JSON marshaling.
func (r *RosettaConstructionParseResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["operations"] = r.Operations
	if r.Signers != nil {
		structMap["signers"] = r.Signers
	}
	if r.AccountIdentifierSigners != nil {
		structMap["account_identifier_signers"] = r.AccountIdentifierSigners
	}
	if r.Metadata != nil {
		structMap["metadata"] = r.Metadata
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RosettaConstructionParseResponse.
// It customizes the JSON unmarshaling process for RosettaConstructionParseResponse objects.
func (r *RosettaConstructionParseResponse) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		Operations               []RosettaOperation         `json:"operations"`
		Signers                  []string                   `json:"signers,omitempty"`
		AccountIdentifierSigners []RosettaAccountIdentifier `json:"account_identifier_signers,omitempty"`
		Metadata                 *interface{}               `json:"metadata,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	r.Operations = temp.Operations
	r.Signers = temp.Signers
	r.AccountIdentifierSigners = temp.AccountIdentifierSigners
	r.Metadata = temp.Metadata
	return nil
}
