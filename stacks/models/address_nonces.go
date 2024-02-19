package models

import (
	"encoding/json"
)

// AddressNonces represents a AddressNonces struct.
// The latest nonce values used by an account by inspecting the mempool, microblock transactions, and anchored transactions
type AddressNonces struct {
	// The latest nonce found within mempool transactions sent by this address. Will be null if there are no current mempool transactions for this address.
	LastMempoolTxNonce *int `json:"last_mempool_tx_nonce"`
	// The latest nonce found within transactions sent by this address, including unanchored microblock transactions. Will be null if there are no current transactions for this address.
	LastExecutedTxNonce *int `json:"last_executed_tx_nonce"`
	// The likely nonce required for creating the next transaction, based on the last nonces seen by the API. This can be incorrect if the API's mempool or transactions aren't fully synchronized, even by a small amount, or if a previous transaction is still propagating through the Stacks blockchain network when this endpoint is called.
	PossibleNextNonce int `json:"possible_next_nonce"`
	// Nonces that appear to be missing and likely causing a mempool transaction to be stuck.
	DetectedMissingNonces []int `json:"detected_missing_nonces"`
	// Nonces currently in mempool for this address.
	DetectedMempoolNonces []int `json:"detected_mempool_nonces,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for AddressNonces.
// It customizes the JSON marshaling process for AddressNonces objects.
func (a *AddressNonces) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(a.toMap())
}

// toMap converts the AddressNonces object to a map representation for JSON marshaling.
func (a *AddressNonces) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["last_mempool_tx_nonce"] = a.LastMempoolTxNonce
	structMap["last_executed_tx_nonce"] = a.LastExecutedTxNonce
	structMap["possible_next_nonce"] = a.PossibleNextNonce
	structMap["detected_missing_nonces"] = a.DetectedMissingNonces
	if a.DetectedMempoolNonces != nil {
		structMap["detected_mempool_nonces"] = a.DetectedMempoolNonces
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AddressNonces.
// It customizes the JSON unmarshaling process for AddressNonces objects.
func (a *AddressNonces) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		LastMempoolTxNonce    *int  `json:"last_mempool_tx_nonce"`
		LastExecutedTxNonce   *int  `json:"last_executed_tx_nonce"`
		PossibleNextNonce     int   `json:"possible_next_nonce"`
		DetectedMissingNonces []int `json:"detected_missing_nonces"`
		DetectedMempoolNonces []int `json:"detected_mempool_nonces,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	a.LastMempoolTxNonce = temp.LastMempoolTxNonce
	a.LastExecutedTxNonce = temp.LastExecutedTxNonce
	a.PossibleNextNonce = temp.PossibleNextNonce
	a.DetectedMissingNonces = temp.DetectedMissingNonces
	a.DetectedMempoolNonces = temp.DetectedMempoolNonces
	return nil
}
