package models

import (
	"encoding/json"
)

// TransactionEvent represents a TransactionEvent struct.
type TransactionEvent struct {
	EventIndex   *int                    `json:"event_index,omitempty"`
	EventType    *EventTypeEnum          `json:"event_type,omitempty"`
	TxId         *string                 `json:"tx_id,omitempty"`
	ContractLog  *ContractLog            `json:"contract_log,omitempty"`
	StxLockEvent *StxLockEvent           `json:"stx_lock_event,omitempty"`
	Asset        *TransactionEventAsset1 `json:"asset,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for TransactionEvent.
// It customizes the JSON marshaling process for TransactionEvent objects.
func (t *TransactionEvent) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(t.toMap())
}

// toMap converts the TransactionEvent object to a map representation for JSON marshaling.
func (t *TransactionEvent) toMap() map[string]any {
	structMap := make(map[string]any)
	if t.EventIndex != nil {
		structMap["event_index"] = t.EventIndex
	}
	if t.EventType != nil {
		structMap["event_type"] = t.EventType
	}
	if t.TxId != nil {
		structMap["tx_id"] = t.TxId
	}
	if t.ContractLog != nil {
		structMap["contract_log"] = t.ContractLog
	}
	if t.StxLockEvent != nil {
		structMap["stx_lock_event"] = t.StxLockEvent
	}
	if t.Asset != nil {
		structMap["asset"] = t.Asset
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for TransactionEvent.
// It customizes the JSON unmarshaling process for TransactionEvent objects.
func (t *TransactionEvent) UnmarshalJSON(input []byte) error {
	temp := &struct {
		EventIndex   *int                    `json:"event_index,omitempty"`
		EventType    *EventTypeEnum          `json:"event_type,omitempty"`
		TxId         *string                 `json:"tx_id,omitempty"`
		ContractLog  *ContractLog            `json:"contract_log,omitempty"`
		StxLockEvent *StxLockEvent           `json:"stx_lock_event,omitempty"`
		Asset        *TransactionEventAsset1 `json:"asset,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	t.EventIndex = temp.EventIndex
	t.EventType = temp.EventType
	t.TxId = temp.TxId
	t.ContractLog = temp.ContractLog
	t.StxLockEvent = temp.StxLockEvent
	t.Asset = temp.Asset
	return nil
}
