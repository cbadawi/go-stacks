package models

import (
	"encoding/json"
)

// TransactionFeeEstimateResponse represents a TransactionFeeEstimateResponse struct.
// POST response for estimated fee
type TransactionFeeEstimateResponse struct {
	EstimatedCostScalar    int           `json:"estimated_cost_scalar"`
	CostScalarChangeByByte *float64      `json:"cost_scalar_change_by_byte,omitempty"`
	EstimatedCost          EstimatedCost `json:"estimated_cost"`
	Estimations            []Estimation  `json:"estimations,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for TransactionFeeEstimateResponse.
// It customizes the JSON marshaling process for TransactionFeeEstimateResponse objects.
func (t *TransactionFeeEstimateResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(t.toMap())
}

// toMap converts the TransactionFeeEstimateResponse object to a map representation for JSON marshaling.
func (t *TransactionFeeEstimateResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["estimated_cost_scalar"] = t.EstimatedCostScalar
	if t.CostScalarChangeByByte != nil {
		structMap["cost_scalar_change_by_byte"] = t.CostScalarChangeByByte
	}
	structMap["estimated_cost"] = t.EstimatedCost
	if t.Estimations != nil {
		structMap["estimations"] = t.Estimations
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for TransactionFeeEstimateResponse.
// It customizes the JSON unmarshaling process for TransactionFeeEstimateResponse objects.
func (t *TransactionFeeEstimateResponse) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		EstimatedCostScalar    int           `json:"estimated_cost_scalar"`
		CostScalarChangeByByte *float64      `json:"cost_scalar_change_by_byte,omitempty"`
		EstimatedCost          EstimatedCost `json:"estimated_cost"`
		Estimations            []Estimation  `json:"estimations,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	t.EstimatedCostScalar = temp.EstimatedCostScalar
	t.CostScalarChangeByByte = temp.CostScalarChangeByByte
	t.EstimatedCost = temp.EstimatedCost
	t.Estimations = temp.Estimations
	return nil
}
