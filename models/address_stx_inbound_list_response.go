package models

import (
	"encoding/json"
)

// AddressStxInboundListResponse represents a AddressStxInboundListResponse struct.
// GET request that returns a list of inbound STX transfers with a memo
type AddressStxInboundListResponse struct {
	Limit   int                  `json:"limit"`
	Offset  int                  `json:"offset"`
	Total   int                  `json:"total"`
	Results []InboundStxTransfer `json:"results"`
}

// MarshalJSON implements the json.Marshaler interface for AddressStxInboundListResponse.
// It customizes the JSON marshaling process for AddressStxInboundListResponse objects.
func (a *AddressStxInboundListResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(a.toMap())
}

// toMap converts the AddressStxInboundListResponse object to a map representation for JSON marshaling.
func (a *AddressStxInboundListResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["limit"] = a.Limit
	structMap["offset"] = a.Offset
	structMap["total"] = a.Total
	structMap["results"] = a.Results
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AddressStxInboundListResponse.
// It customizes the JSON unmarshaling process for AddressStxInboundListResponse objects.
func (a *AddressStxInboundListResponse) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		Limit   int                  `json:"limit"`
		Offset  int                  `json:"offset"`
		Total   int                  `json:"total"`
		Results []InboundStxTransfer `json:"results"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	a.Limit = temp.Limit
	a.Offset = temp.Offset
	a.Total = temp.Total
	a.Results = temp.Results
	return nil
}
