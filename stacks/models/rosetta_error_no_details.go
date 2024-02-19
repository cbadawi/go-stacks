package models

import (
	"encoding/json"
)

// RosettaErrorNoDetails represents a RosettaErrorNoDetails struct.
// Instead of utilizing HTTP status codes to describe node errors (which often do not have a good analog), rich errors are returned using this object. Both the code and message fields can be individually used to correctly identify an error. Implementations MUST use unique values for both fields.
type RosettaErrorNoDetails struct {
	// Code is a network-specific error code. If desired, this code can be equivalent to an HTTP status code.
	Code int `json:"code"`
	// Message is a network-specific error message. The message MUST NOT change for a given code. In particular, this means that any contextual information should be included in the details field.
	Message string `json:"message"`
	// An error is retriable if the same request may succeed if submitted again.
	Retriable bool `json:"retriable"`
}

// MarshalJSON implements the json.Marshaler interface for RosettaErrorNoDetails.
// It customizes the JSON marshaling process for RosettaErrorNoDetails objects.
func (r *RosettaErrorNoDetails) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the RosettaErrorNoDetails object to a map representation for JSON marshaling.
func (r *RosettaErrorNoDetails) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["code"] = r.Code
	structMap["message"] = r.Message
	structMap["retriable"] = r.Retriable
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RosettaErrorNoDetails.
// It customizes the JSON unmarshaling process for RosettaErrorNoDetails objects.
func (r *RosettaErrorNoDetails) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		Code      int    `json:"code"`
		Message   string `json:"message"`
		Retriable bool   `json:"retriable"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	r.Code = temp.Code
	r.Message = temp.Message
	r.Retriable = temp.Retriable
	return nil
}
