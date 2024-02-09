package models

import (
	"encoding/json"
)

// SearchResult represents a SearchResult struct.
// complete search result for terms
type SearchResult struct {
	// Indicates if the requested object was found or not
	Found *bool `json:"found,omitempty"`
	// This object carries the search result
	Result *Result11 `json:"result,omitempty"`
	Error  *string   `json:"error,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for SearchResult.
// It customizes the JSON marshaling process for SearchResult objects.
func (s *SearchResult) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the SearchResult object to a map representation for JSON marshaling.
func (s *SearchResult) toMap() map[string]any {
	structMap := make(map[string]any)
	if s.Found != nil {
		structMap["found"] = s.Found
	}
	if s.Result != nil {
		structMap["result"] = s.Result
	}
	if s.Error != nil {
		structMap["error"] = s.Error
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SearchResult.
// It customizes the JSON unmarshaling process for SearchResult objects.
func (s *SearchResult) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		Found  *bool     `json:"found,omitempty"`
		Result *Result11 `json:"result,omitempty"`
		Error  *string   `json:"error,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	s.Found = temp.Found
	s.Result = temp.Result
	s.Error = temp.Error
	return nil
}
