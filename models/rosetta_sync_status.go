package models

import (
	"encoding/json"
)

// RosettaSyncStatus represents a RosettaSyncStatus struct.
// SyncStatus is used to provide additional context about an implementation's sync status. It is often used to indicate that an implementation is healthy when it cannot be queried until some sync phase occurs. If an implementation is immediately queryable, this model is often not populated.
type RosettaSyncStatus struct {
	// CurrentIndex is the index of the last synced block in the current stage.
	CurrentIndex int `json:"current_index"`
	// TargetIndex is the index of the block that the implementation is attempting to sync to in the current stage.
	TargetIndex *int `json:"target_index,omitempty"`
	// Stage is the phase of the sync process.
	Stage *string `json:"stage,omitempty"`
	// Synced indicates if an implementation has synced up to the most recent block.
	Synced *bool `json:"synced,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for RosettaSyncStatus.
// It customizes the JSON marshaling process for RosettaSyncStatus objects.
func (r *RosettaSyncStatus) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the RosettaSyncStatus object to a map representation for JSON marshaling.
func (r *RosettaSyncStatus) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["current_index"] = r.CurrentIndex
	if r.TargetIndex != nil {
		structMap["target_index"] = r.TargetIndex
	}
	if r.Stage != nil {
		structMap["stage"] = r.Stage
	}
	if r.Synced != nil {
		structMap["synced"] = r.Synced
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RosettaSyncStatus.
// It customizes the JSON unmarshaling process for RosettaSyncStatus objects.
func (r *RosettaSyncStatus) UnmarshalJSON(input []byte) error {
	temp := &struct {
		CurrentIndex int     `json:"current_index"`
		TargetIndex  *int    `json:"target_index,omitempty"`
		Stage        *string `json:"stage,omitempty"`
		Synced       *bool   `json:"synced,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	r.CurrentIndex = temp.CurrentIndex
	r.TargetIndex = temp.TargetIndex
	r.Stage = temp.Stage
	r.Synced = temp.Synced
	return nil
}
