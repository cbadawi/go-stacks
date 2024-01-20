package models

import (
	"encoding/json"
)

// TransactionEventsResponse represents a TransactionEventsResponse struct.
// GET event for the given transaction
type TransactionEventsResponse struct {
	Limit   int                `json:"limit"`
	Offset  int                `json:"offset"`
	Results []TransactionEvent `json:"results"`
}

// MarshalJSON implements the json.Marshaler interface for TransactionEventsResponse.
// It customizes the JSON marshaling process for TransactionEventsResponse objects.
func (t *TransactionEventsResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(t.toMap())
}

// toMap converts the TransactionEventsResponse object to a map representation for JSON marshaling.
func (t *TransactionEventsResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["limit"] = t.Limit
	structMap["offset"] = t.Offset
	structMap["results"] = t.Results
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for TransactionEventsResponse.
// It customizes the JSON unmarshaling process for TransactionEventsResponse objects.
func (t *TransactionEventsResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Limit   int                `json:"limit"`
		Offset  int                `json:"offset"`
		Results []TransactionEvent `json:"results"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	t.Limit = temp.Limit
	t.Offset = temp.Offset
	t.Results = temp.Results
	return nil
}
