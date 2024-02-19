package models

import (
	"encoding/json"
)

// TransactionEvent1 represents a TransactionEvent1 struct.
type TransactionEvent1 struct {
	EventIndex   int                    `json:"event_index"`
	EventType    string                 `json:"event_type"`
	TxId         string                 `json:"tx_id"`
	ContractLog  ContractLog            `json:"contract_log"`
	StxLockEvent StxLockEvent           `json:"stx_lock_event"`
	Asset        TransactionEventAsset1 `json:"asset"`
}

// MarshalJSON implements the json.Marshaler interface for TransactionEvent1.
// It customizes the JSON marshaling process for TransactionEvent1 objects.
func (t *TransactionEvent1) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(t.toMap())
}

// toMap converts the TransactionEvent1 object to a map representation for JSON marshaling.
func (t *TransactionEvent1) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["event_index"] = t.EventIndex
	structMap["event_type"] = t.EventType
	structMap["tx_id"] = t.TxId
	structMap["contract_log"] = t.ContractLog
	structMap["stx_lock_event"] = t.StxLockEvent
	structMap["asset"] = t.Asset
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for TransactionEvent1.
// It customizes the JSON unmarshaling process for TransactionEvent1 objects.
func (t *TransactionEvent1) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		EventIndex   int                    `json:"event_index"`
		EventType    string                 `json:"event_type"`
		TxId         string                 `json:"tx_id"`
		ContractLog  ContractLog            `json:"contract_log"`
		StxLockEvent StxLockEvent           `json:"stx_lock_event"`
		Asset        TransactionEventAsset1 `json:"asset"`
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
