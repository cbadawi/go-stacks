package models

import (
	"encoding/json"
)

// InboundStxTransfer represents a InboundStxTransfer struct.
// A inbound STX transfer with a memo
type InboundStxTransfer struct {
	// Principal that sent this transfer
	Sender string `json:"sender"`
	// Transfer amount in micro-STX as integer string
	Amount string `json:"amount"`
	// Hex encoded memo bytes associated with the transfer
	Memo string `json:"memo"`
	// Block height at which this transfer occurred
	BlockHeight float64 `json:"block_height"`
	// The transaction ID in which this transfer occurred
	TxId string `json:"tx_id"`
	// Indicates if the transfer is from a stx-transfer transaction or a contract-call transaction
	TransferType TransferTypeEnum `json:"transfer_type"`
	// Index of the transaction within a block
	TxIndex float64 `json:"tx_index"`
}

// MarshalJSON implements the json.Marshaler interface for InboundStxTransfer.
// It customizes the JSON marshaling process for InboundStxTransfer objects.
func (i *InboundStxTransfer) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(i.toMap())
}

// toMap converts the InboundStxTransfer object to a map representation for JSON marshaling.
func (i *InboundStxTransfer) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["sender"] = i.Sender
	structMap["amount"] = i.Amount
	structMap["memo"] = i.Memo
	structMap["block_height"] = i.BlockHeight
	structMap["tx_id"] = i.TxId
	structMap["transfer_type"] = i.TransferType
	structMap["tx_index"] = i.TxIndex
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for InboundStxTransfer.
// It customizes the JSON unmarshaling process for InboundStxTransfer objects.
func (i *InboundStxTransfer) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		Sender       string           `json:"sender"`
		Amount       string           `json:"amount"`
		Memo         string           `json:"memo"`
		BlockHeight  float64          `json:"block_height"`
		TxId         string           `json:"tx_id"`
		TransferType TransferTypeEnum `json:"transfer_type"`
		TxIndex      float64          `json:"tx_index"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	i.Sender = temp.Sender
	i.Amount = temp.Amount
	i.Memo = temp.Memo
	i.BlockHeight = temp.BlockHeight
	i.TxId = temp.TxId
	i.TransferType = temp.TransferType
	i.TxIndex = temp.TxIndex
	return nil
}
