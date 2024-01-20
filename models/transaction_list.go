package models

import (
	"encoding/json"
)

// TransactionList represents a TransactionList struct.
type TransactionList struct {
	Found  *bool    `json:"found,omitempty"`
	Result *Result2 `json:"result,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for TransactionList.
// It customizes the JSON marshaling process for TransactionList objects.
func (t *TransactionList) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(t.toMap())
}

// toMap converts the TransactionList object to a map representation for JSON marshaling.
func (t *TransactionList) toMap() map[string]any {
	structMap := make(map[string]any)
	if t.Found != nil {
		structMap["found"] = t.Found
	}
	if t.Result != nil {
		structMap["result"] = t.Result
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for TransactionList.
// It customizes the JSON unmarshaling process for TransactionList objects.
func (t *TransactionList) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Found  *bool    `json:"found,omitempty"`
		Result *Result2 `json:"result,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	t.Found = temp.Found
	t.Result = temp.Result
	return nil
}
