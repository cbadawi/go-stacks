package models

import (
	"encoding/json"
)

// Allow represents a Allow struct.
// Allow specifies supported Operation status, Operation types, and all possible error statuses. This Allow object is used by clients to validate the correctness of a Rosetta Server implementation. It is expected that these clients will error if they receive some response that contains any of the above information that is not specified here.
type Allow struct {
	// All Operation.Status this implementation supports. Any status that is returned during parsing that is not listed here will cause client validation to error.
	OperationStatuses []RosettaOperationStatus `json:"operation_statuses"`
	// All Operation.Type this implementation supports. Any type that is returned during parsing that is not listed here will cause client validation to error.
	OperationTypes []string `json:"operation_types"`
	// All Errors that this implementation could return. Any error that is returned during parsing that is not listed here will cause client validation to error.
	Errors []RosettaErrorNoDetails `json:"errors"`
	// Any Rosetta implementation that supports querying the balance of an account at any height in the past should set this to true.
	HistoricalBalanceLookup bool `json:"historical_balance_lookup"`
}

// MarshalJSON implements the json.Marshaler interface for Allow.
// It customizes the JSON marshaling process for Allow objects.
func (a *Allow) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(a.toMap())
}

// toMap converts the Allow object to a map representation for JSON marshaling.
func (a *Allow) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["operation_statuses"] = a.OperationStatuses
	structMap["operation_types"] = a.OperationTypes
	structMap["errors"] = a.Errors
	structMap["historical_balance_lookup"] = a.HistoricalBalanceLookup
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Allow.
// It customizes the JSON unmarshaling process for Allow objects.
func (a *Allow) UnmarshalJSON(input []byte) error {
	temp := &struct {
		OperationStatuses       []RosettaOperationStatus `json:"operation_statuses"`
		OperationTypes          []string                 `json:"operation_types"`
		Errors                  []RosettaErrorNoDetails  `json:"errors"`
		HistoricalBalanceLookup bool                     `json:"historical_balance_lookup"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	a.OperationStatuses = temp.OperationStatuses
	a.OperationTypes = temp.OperationTypes
	a.Errors = temp.Errors
	a.HistoricalBalanceLookup = temp.HistoricalBalanceLookup
	return nil
}
