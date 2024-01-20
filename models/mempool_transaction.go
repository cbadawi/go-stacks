package models

import (
	"encoding/json"
)

// MempoolTransaction represents a MempoolTransaction struct.
// Describes all transaction types on Stacks 2.0 blockchain
type MempoolTransaction struct {
	// Transaction ID
	TxId *string `json:"tx_id,omitempty"`
	// Used for ordering the transactions originating from and paying from an account. The nonce ensures that a transaction is processed at most once. The nonce counts the number of times an account's owner(s) have authorized a transaction. The first transaction from an account will have a nonce value equal to 0, the second will have a nonce value equal to 1, and so on.
	Nonce *int `json:"nonce,omitempty"`
	// Transaction fee as Integer string (64-bit unsigned integer).
	FeeRate *string `json:"fee_rate,omitempty"`
	// Address of the transaction initiator
	SenderAddress *string `json:"sender_address,omitempty"`
	SponsorNonce  *int    `json:"sponsor_nonce,omitempty"`
	// Denotes whether the originating account is the same as the paying account
	Sponsored         *bool                  `json:"sponsored,omitempty"`
	SponsorAddress    *string                `json:"sponsor_address,omitempty"`
	PostConditionMode *PostConditionModeEnum `json:"post_condition_mode,omitempty"`
	PostConditions    []PostCondition1       `json:"post_conditions,omitempty"`
	// `on_chain_only`: the transaction MUST be included in an anchored block, `off_chain_only`: the transaction MUST be included in a microblock, `any`: the leader can choose where to include the transaction.
	AnchorMode *TransactionAnchorModeTypeEnum `json:"anchor_mode,omitempty"`
	// Status of the transaction
	TxStatus *MempoolTransactionStatusEnum `json:"tx_status,omitempty"`
	// A unix timestamp (in seconds) indicating when the transaction broadcast was received by the node.
	ReceiptTime *float64 `json:"receipt_time,omitempty"`
	// An ISO 8601 (YYYY-MM-DDTHH:mm:ss.sssZ) timestamp indicating when the transaction broadcast was received by the node.
	ReceiptTimeIso      *string              `json:"receipt_time_iso,omitempty"`
	TxType              *TxTypeEnum          `json:"tx_type,omitempty"`
	TokenTransfer       *TokenTransfer       `json:"token_transfer,omitempty"`
	SmartContract       *SmartContract       `json:"smart_contract,omitempty"`
	ContractCall        *ContractCall        `json:"contract_call,omitempty"`
	PoisonMicroblock    *PoisonMicroblock    `json:"poison_microblock,omitempty"`
	CoinbasePayload     *CoinbasePayload     `json:"coinbase_payload,omitempty"`
	TenureChangePayload *TenureChangePayload `json:"tenure_change_payload,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for MempoolTransaction.
// It customizes the JSON marshaling process for MempoolTransaction objects.
func (m *MempoolTransaction) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(m.toMap())
}

// toMap converts the MempoolTransaction object to a map representation for JSON marshaling.
func (m *MempoolTransaction) toMap() map[string]any {
	structMap := make(map[string]any)
	if m.TxId != nil {
		structMap["tx_id"] = m.TxId
	}
	if m.Nonce != nil {
		structMap["nonce"] = m.Nonce
	}
	if m.FeeRate != nil {
		structMap["fee_rate"] = m.FeeRate
	}
	if m.SenderAddress != nil {
		structMap["sender_address"] = m.SenderAddress
	}
	if m.SponsorNonce != nil {
		structMap["sponsor_nonce"] = m.SponsorNonce
	}
	if m.Sponsored != nil {
		structMap["sponsored"] = m.Sponsored
	}
	if m.SponsorAddress != nil {
		structMap["sponsor_address"] = m.SponsorAddress
	}
	if m.PostConditionMode != nil {
		structMap["post_condition_mode"] = m.PostConditionMode
	}
	if m.PostConditions != nil {
		structMap["post_conditions"] = m.PostConditions
	}
	if m.AnchorMode != nil {
		structMap["anchor_mode"] = m.AnchorMode
	}
	if m.TxStatus != nil {
		structMap["tx_status"] = m.TxStatus
	}
	if m.ReceiptTime != nil {
		structMap["receipt_time"] = m.ReceiptTime
	}
	if m.ReceiptTimeIso != nil {
		structMap["receipt_time_iso"] = m.ReceiptTimeIso
	}
	if m.TxType != nil {
		structMap["tx_type"] = m.TxType
	}
	if m.TokenTransfer != nil {
		structMap["token_transfer"] = m.TokenTransfer
	}
	if m.SmartContract != nil {
		structMap["smart_contract"] = m.SmartContract
	}
	if m.ContractCall != nil {
		structMap["contract_call"] = m.ContractCall
	}
	if m.PoisonMicroblock != nil {
		structMap["poison_microblock"] = m.PoisonMicroblock
	}
	if m.CoinbasePayload != nil {
		structMap["coinbase_payload"] = m.CoinbasePayload
	}
	if m.TenureChangePayload != nil {
		structMap["tenure_change_payload"] = m.TenureChangePayload
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MempoolTransaction.
// It customizes the JSON unmarshaling process for MempoolTransaction objects.
func (m *MempoolTransaction) UnmarshalJSON(input []byte) error {
	temp := &struct {
		TxId                *string                        `json:"tx_id,omitempty"`
		Nonce               *int                           `json:"nonce,omitempty"`
		FeeRate             *string                        `json:"fee_rate,omitempty"`
		SenderAddress       *string                        `json:"sender_address,omitempty"`
		SponsorNonce        *int                           `json:"sponsor_nonce,omitempty"`
		Sponsored           *bool                          `json:"sponsored,omitempty"`
		SponsorAddress      *string                        `json:"sponsor_address,omitempty"`
		PostConditionMode   *PostConditionModeEnum         `json:"post_condition_mode,omitempty"`
		PostConditions      []PostCondition1               `json:"post_conditions,omitempty"`
		AnchorMode          *TransactionAnchorModeTypeEnum `json:"anchor_mode,omitempty"`
		TxStatus            *MempoolTransactionStatusEnum  `json:"tx_status,omitempty"`
		ReceiptTime         *float64                       `json:"receipt_time,omitempty"`
		ReceiptTimeIso      *string                        `json:"receipt_time_iso,omitempty"`
		TxType              *TxTypeEnum                    `json:"tx_type,omitempty"`
		TokenTransfer       *TokenTransfer                 `json:"token_transfer,omitempty"`
		SmartContract       *SmartContract                 `json:"smart_contract,omitempty"`
		ContractCall        *ContractCall                  `json:"contract_call,omitempty"`
		PoisonMicroblock    *PoisonMicroblock              `json:"poison_microblock,omitempty"`
		CoinbasePayload     *CoinbasePayload               `json:"coinbase_payload,omitempty"`
		TenureChangePayload *TenureChangePayload           `json:"tenure_change_payload,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	m.TxId = temp.TxId
	m.Nonce = temp.Nonce
	m.FeeRate = temp.FeeRate
	m.SenderAddress = temp.SenderAddress
	m.SponsorNonce = temp.SponsorNonce
	m.Sponsored = temp.Sponsored
	m.SponsorAddress = temp.SponsorAddress
	m.PostConditionMode = temp.PostConditionMode
	m.PostConditions = temp.PostConditions
	m.AnchorMode = temp.AnchorMode
	m.TxStatus = temp.TxStatus
	m.ReceiptTime = temp.ReceiptTime
	m.ReceiptTimeIso = temp.ReceiptTimeIso
	m.TxType = temp.TxType
	m.TokenTransfer = temp.TokenTransfer
	m.SmartContract = temp.SmartContract
	m.ContractCall = temp.ContractCall
	m.PoisonMicroblock = temp.PoisonMicroblock
	m.CoinbasePayload = temp.CoinbasePayload
	m.TenureChangePayload = temp.TenureChangePayload
	return nil
}
