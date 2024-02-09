package models

import (
	"encoding/json"
)

// RosettaOperationStatus represents a RosettaOperationStatus struct.
// OperationStatus is utilized to indicate which Operation status are considered successful.
type RosettaOperationStatus struct {
	// The status is the network-specific status of the operation.
	Status string `json:"status"`
	// An Operation is considered successful if the Operation.Amount should affect the Operation.Account. Some blockchains (like Bitcoin) only include successful operations in blocks but other blockchains (like Ethereum) include unsuccessful operations that incur a fee. To reconcile the computed balance from the stream of Operations, it is critical to understand which Operation.Status indicate an Operation is successful and should affect an Account.
	Successful bool `json:"successful"`
}

// MarshalJSON implements the json.Marshaler interface for RosettaOperationStatus.
// It customizes the JSON marshaling process for RosettaOperationStatus objects.
func (r *RosettaOperationStatus) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the RosettaOperationStatus object to a map representation for JSON marshaling.
func (r *RosettaOperationStatus) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["status"] = r.Status
	structMap["successful"] = r.Successful
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RosettaOperationStatus.
// It customizes the JSON unmarshaling process for RosettaOperationStatus objects.
func (r *RosettaOperationStatus) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		Status     string `json:"status"`
		Successful bool   `json:"successful"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	r.Status = temp.Status
	r.Successful = temp.Successful
	return nil
}
