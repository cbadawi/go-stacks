package models

import (
	"encoding/json"
)

// MempoolTransactionListResponse represents a MempoolTransactionListResponse struct.
// GET request that returns transactions
type MempoolTransactionListResponse struct {
	Limit   int                  `json:"limit"`
	Offset  int                  `json:"offset"`
	Total   int                  `json:"total"`
	Results []MempoolTransaction `json:"results"`
}

// MarshalJSON implements the json.Marshaler interface for MempoolTransactionListResponse.
// It customizes the JSON marshaling process for MempoolTransactionListResponse objects.
func (m *MempoolTransactionListResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(m.toMap())
}

// toMap converts the MempoolTransactionListResponse object to a map representation for JSON marshaling.
func (m *MempoolTransactionListResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["limit"] = m.Limit
	structMap["offset"] = m.Offset
	structMap["total"] = m.Total
	structMap["results"] = m.Results
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MempoolTransactionListResponse.
// It customizes the JSON unmarshaling process for MempoolTransactionListResponse objects.
func (m *MempoolTransactionListResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Limit   int                  `json:"limit"`
		Offset  int                  `json:"offset"`
		Total   int                  `json:"total"`
		Results []MempoolTransaction `json:"results"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	m.Limit = temp.Limit
	m.Offset = temp.Offset
	m.Total = temp.Total
	m.Results = temp.Results
	return nil
}
