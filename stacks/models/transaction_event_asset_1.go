package models

import (
	"encoding/json"
)

// TransactionEventAsset1 represents a TransactionEventAsset1 struct.
type TransactionEventAsset1 struct {
	AssetEventType TransactionEventAssetTypeEnum `json:"asset_event_type"`
	AssetId        string                        `json:"asset_id"`
	Sender         string                        `json:"sender"`
	Recipient      string                        `json:"recipient"`
	Amount         string                        `json:"amount"`
	Value          Value                         `json:"value"`
	Memo           *string                       `json:"memo,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for TransactionEventAsset1.
// It customizes the JSON marshaling process for TransactionEventAsset1 objects.
func (t *TransactionEventAsset1) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(t.toMap())
}

// toMap converts the TransactionEventAsset1 object to a map representation for JSON marshaling.
func (t *TransactionEventAsset1) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["asset_event_type"] = t.AssetEventType
	structMap["asset_id"] = t.AssetId
	structMap["sender"] = t.Sender
	structMap["recipient"] = t.Recipient
	structMap["amount"] = t.Amount
	structMap["value"] = t.Value
	if t.Memo != nil {
		structMap["memo"] = t.Memo
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for TransactionEventAsset1.
// It customizes the JSON unmarshaling process for TransactionEventAsset1 objects.
func (t *TransactionEventAsset1) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		AssetEventType TransactionEventAssetTypeEnum `json:"asset_event_type"`
		AssetId        string                        `json:"asset_id"`
		Sender         string                        `json:"sender"`
		Recipient      string                        `json:"recipient"`
		Amount         string                        `json:"amount"`
		Value          Value                         `json:"value"`
		Memo           *string                       `json:"memo,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	t.AssetEventType = temp.AssetEventType
	t.AssetId = temp.AssetId
	t.Sender = temp.Sender
	t.Recipient = temp.Recipient
	t.Amount = temp.Amount
	t.Value = temp.Value
	t.Memo = temp.Memo
	return nil
}
