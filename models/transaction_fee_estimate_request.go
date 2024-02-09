package models

import (
	"encoding/json"
)

// TransactionFeeEstimateRequest represents a TransactionFeeEstimateRequest struct.
// POST request for estimated fee
type TransactionFeeEstimateRequest struct {
	TransactionPayload string `json:"transaction_payload"`
	EstimatedLen       *int   `json:"estimated_len,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for TransactionFeeEstimateRequest.
// It customizes the JSON marshaling process for TransactionFeeEstimateRequest objects.
func (t *TransactionFeeEstimateRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(t.toMap())
}

// toMap converts the TransactionFeeEstimateRequest object to a map representation for JSON marshaling.
func (t *TransactionFeeEstimateRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["transaction_payload"] = t.TransactionPayload
	if t.EstimatedLen != nil {
		structMap["estimated_len"] = t.EstimatedLen
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for TransactionFeeEstimateRequest.
// It customizes the JSON unmarshaling process for TransactionFeeEstimateRequest objects.
func (t *TransactionFeeEstimateRequest) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		TransactionPayload string `json:"transaction_payload"`
		EstimatedLen       *int   `json:"estimated_len,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	t.TransactionPayload = temp.TransactionPayload
	t.EstimatedLen = temp.EstimatedLen
	return nil
}
