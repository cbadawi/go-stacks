package models

import (
	"encoding/json"
)

// ReadOnlyFunctionSuccessResponse represents a ReadOnlyFunctionSuccessResponse struct.
// GET request to get contract source
type ReadOnlyFunctionSuccessResponse struct {
	Okay   bool    `json:"okay"`
	Result *string `json:"result,omitempty"`
	Cause  *string `json:"cause,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for ReadOnlyFunctionSuccessResponse.
// It customizes the JSON marshaling process for ReadOnlyFunctionSuccessResponse objects.
func (r *ReadOnlyFunctionSuccessResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the ReadOnlyFunctionSuccessResponse object to a map representation for JSON marshaling.
func (r *ReadOnlyFunctionSuccessResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["okay"] = r.Okay
	if r.Result != nil {
		structMap["result"] = r.Result
	}
	if r.Cause != nil {
		structMap["cause"] = r.Cause
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ReadOnlyFunctionSuccessResponse.
// It customizes the JSON unmarshaling process for ReadOnlyFunctionSuccessResponse objects.
func (r *ReadOnlyFunctionSuccessResponse) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		Okay   bool    `json:"okay"`
		Result *string `json:"result,omitempty"`
		Cause  *string `json:"cause,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	r.Okay = temp.Okay
	r.Result = temp.Result
	r.Cause = temp.Cause
	return nil
}
