package models

import (
	"encoding/json"
)

// RosettaBlockResponse represents a RosettaBlockResponse struct.
// A BlockResponse includes a fully-populated block or a partially-populated block with a list of other transactions to fetch (other_transactions). As a result of the consensus algorithm of some blockchains, blocks can be omitted (i.e. certain block indexes can be skipped). If a query for one of these omitted indexes is made, the response should not include a Block object. It is VERY important to note that blocks MUST still form a canonical, connected chain of blocks where each block has a unique index. In other words, the PartialBlockIdentifier of a block after an omitted block should reference the last non-omitted block.
type RosettaBlockResponse struct {
	// Blocks contain an array of Transactions that occurred at a particular BlockIdentifier. A hard requirement for blocks returned by Rosetta implementations is that they MUST be inalterable: once a client has requested and received a block identified by a specific BlockIndentifier, all future calls for that same BlockIdentifier must return the same block contents.
	Block *RosettaBlock `json:"block,omitempty"`
	// Some blockchains may require additional transactions to be fetched that weren't returned in the block response (ex: block only returns transaction hashes). For blockchains with a lot of transactions in each block, this can be very useful as consumers can concurrently fetch all transactions returned.
	OtherTransactions []OtherTransactionIdentifier `json:"other_transactions,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for RosettaBlockResponse.
// It customizes the JSON marshaling process for RosettaBlockResponse objects.
func (r *RosettaBlockResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the RosettaBlockResponse object to a map representation for JSON marshaling.
func (r *RosettaBlockResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if r.Block != nil {
		structMap["block"] = r.Block
	}
	if r.OtherTransactions != nil {
		structMap["other_transactions"] = r.OtherTransactions
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RosettaBlockResponse.
// It customizes the JSON unmarshaling process for RosettaBlockResponse objects.
func (r *RosettaBlockResponse) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		Block             *RosettaBlock                `json:"block,omitempty"`
		OtherTransactions []OtherTransactionIdentifier `json:"other_transactions,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	r.Block = temp.Block
	r.OtherTransactions = temp.OtherTransactions
	return nil
}
