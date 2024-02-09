package models

import (
	"encoding/json"
)

// RosettaBlockTransactionRequest represents a RosettaBlockTransactionRequest struct.
// A BlockTransactionRequest is used to fetch a Transaction included in a block that is not returned in a BlockResponse.
type RosettaBlockTransactionRequest struct {
	// The network_identifier specifies which network a particular object is associated with.
	NetworkIdentifier NetworkIdentifier `json:"network_identifier"`
	// When fetching data by BlockIdentifier, it may be possible to only specify the index or hash. If neither property is specified, it is assumed that the client is making a request at the current block.
	BlockIdentifier RosettaPartialBlockIdentifier `json:"block_identifier"`
	// The transaction_identifier uniquely identifies a transaction in a particular network and block or in the mempool.
	TransactionIdentifier TransactionIdentifier `json:"transaction_identifier"`
}

// MarshalJSON implements the json.Marshaler interface for RosettaBlockTransactionRequest.
// It customizes the JSON marshaling process for RosettaBlockTransactionRequest objects.
func (r *RosettaBlockTransactionRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the RosettaBlockTransactionRequest object to a map representation for JSON marshaling.
func (r *RosettaBlockTransactionRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["network_identifier"] = r.NetworkIdentifier
	structMap["block_identifier"] = r.BlockIdentifier
	structMap["transaction_identifier"] = r.TransactionIdentifier
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RosettaBlockTransactionRequest.
// It customizes the JSON unmarshaling process for RosettaBlockTransactionRequest objects.
func (r *RosettaBlockTransactionRequest) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		NetworkIdentifier     NetworkIdentifier             `json:"network_identifier"`
		BlockIdentifier       RosettaPartialBlockIdentifier `json:"block_identifier"`
		TransactionIdentifier TransactionIdentifier         `json:"transaction_identifier"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	r.NetworkIdentifier = temp.NetworkIdentifier
	r.BlockIdentifier = temp.BlockIdentifier
	r.TransactionIdentifier = temp.TransactionIdentifier
	return nil
}
