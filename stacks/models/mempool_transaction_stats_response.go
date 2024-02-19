package models

import (
	"encoding/json"
)

// MempoolTransactionStatsResponse represents a MempoolTransactionStatsResponse struct.
// GET request that returns stats on mempool transactions
type MempoolTransactionStatsResponse struct {
	// Number of tranasction in the mempool, broken down by transaction type.
	TxTypeCounts TxTypeCounts `json:"tx_type_counts"`
	// The simple mean (average) transaction fee, broken down by transaction type. Note that this does not factor in actual execution costs. The average fee is not a reliable metric for calculating a fee for a new transaction.
	TxSimpleFeeAverages TxSimpleFeeAverages `json:"tx_simple_fee_averages"`
	// The average time (in blocks) that transactions have lived in the mempool. The start block height is simply the current chain-tip of when the attached Stacks node receives the transaction. This timing can be different across Stacks nodes / API instances due to propagation timing differences in the p2p network.
	TxAges TxAges `json:"tx_ages"`
	// The average byte size of transactions in the mempool, broken down by transaction type.
	TxByteSizes TxByteSizes `json:"tx_byte_sizes"`
}

// MarshalJSON implements the json.Marshaler interface for MempoolTransactionStatsResponse.
// It customizes the JSON marshaling process for MempoolTransactionStatsResponse objects.
func (m *MempoolTransactionStatsResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(m.toMap())
}

// toMap converts the MempoolTransactionStatsResponse object to a map representation for JSON marshaling.
func (m *MempoolTransactionStatsResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["tx_type_counts"] = m.TxTypeCounts
	structMap["tx_simple_fee_averages"] = m.TxSimpleFeeAverages
	structMap["tx_ages"] = m.TxAges
	structMap["tx_byte_sizes"] = m.TxByteSizes
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MempoolTransactionStatsResponse.
// It customizes the JSON unmarshaling process for MempoolTransactionStatsResponse objects.
func (m *MempoolTransactionStatsResponse) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		TxTypeCounts        TxTypeCounts        `json:"tx_type_counts"`
		TxSimpleFeeAverages TxSimpleFeeAverages `json:"tx_simple_fee_averages"`
		TxAges              TxAges              `json:"tx_ages"`
		TxByteSizes         TxByteSizes         `json:"tx_byte_sizes"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	m.TxTypeCounts = temp.TxTypeCounts
	m.TxSimpleFeeAverages = temp.TxSimpleFeeAverages
	m.TxAges = temp.TxAges
	m.TxByteSizes = temp.TxByteSizes
	return nil
}
