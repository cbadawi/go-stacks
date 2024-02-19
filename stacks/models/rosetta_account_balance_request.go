package models

import (
	"encoding/json"
)

// RosettaAccountBalanceRequest represents a RosettaAccountBalanceRequest struct.
// An AccountBalanceRequest is utilized to make a balance request on the /account/balance endpoint. If the block_identifier is populated, a historical balance query should be performed.
type RosettaAccountBalanceRequest struct {
	// The network_identifier specifies which network a particular object is associated with.
	NetworkIdentifier NetworkIdentifier `json:"network_identifier"`
	// The account_identifier uniquely identifies an account within a network. All fields in the account_identifier are utilized to determine this uniqueness (including the metadata field, if populated).
	AccountIdentifier RosettaAccount `json:"account_identifier"`
	// When fetching data by BlockIdentifier, it may be possible to only specify the index or hash. If neither property is specified, it is assumed that the client is making a request at the current block.
	BlockIdentifier *RosettaPartialBlockIdentifier `json:"block_identifier,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for RosettaAccountBalanceRequest.
// It customizes the JSON marshaling process for RosettaAccountBalanceRequest objects.
func (r *RosettaAccountBalanceRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the RosettaAccountBalanceRequest object to a map representation for JSON marshaling.
func (r *RosettaAccountBalanceRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["network_identifier"] = r.NetworkIdentifier
	structMap["account_identifier"] = r.AccountIdentifier
	if r.BlockIdentifier != nil {
		structMap["block_identifier"] = r.BlockIdentifier
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RosettaAccountBalanceRequest.
// It customizes the JSON unmarshaling process for RosettaAccountBalanceRequest objects.
func (r *RosettaAccountBalanceRequest) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		NetworkIdentifier NetworkIdentifier              `json:"network_identifier"`
		AccountIdentifier RosettaAccount                 `json:"account_identifier"`
		BlockIdentifier   *RosettaPartialBlockIdentifier `json:"block_identifier,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	r.NetworkIdentifier = temp.NetworkIdentifier
	r.AccountIdentifier = temp.AccountIdentifier
	r.BlockIdentifier = temp.BlockIdentifier
	return nil
}
