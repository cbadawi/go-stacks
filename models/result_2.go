package models

import (
	"encoding/json"
)

// Result2 represents a Result2 struct.
type Result2 struct {
	TxId string `json:"tx_id"`
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
	// Hash of the blocked this transactions was associated with
	BlockHash *string `json:"block_hash,omitempty"`
	// Height of the block this transactions was associated with
	BlockHeight *int `json:"block_height,omitempty"`
	// Unix timestamp (in seconds) indicating when this block was mined
	BurnBlockTime *int `json:"burn_block_time,omitempty"`
	// An ISO 8601 (YYYY-MM-DDTHH:mm:ss.sssZ) timestamp indicating when this block was mined.
	BurnBlockTimeIso *string `json:"burn_block_time_iso,omitempty"`
	// Unix timestamp (in seconds) indicating when this parent block was mined
	ParentBurnBlockTime *int `json:"parent_burn_block_time,omitempty"`
	// An ISO 8601 (YYYY-MM-DDTHH:mm:ss.sssZ) timestamp indicating when this parent block was mined.
	ParentBurnBlockTimeIso *string `json:"parent_burn_block_time_iso,omitempty"`
	// Set to `true` if block corresponds to the canonical chain tip
	Canonical *bool `json:"canonical,omitempty"`
	// Index of the transaction, indicating the order. Starts at `0` and increases with each transaction
	TxIndex *int `json:"tx_index,omitempty"`
	// Result of the transaction. For contract calls, this will show the value returned by the call. For other transaction types, this will return a boolean indicating the success of the transaction.
	TxResult *TxResult `json:"tx_result,omitempty"`
	// Number of transaction events
	EventCount *int `json:"event_count,omitempty"`
	// Hash of the previous block.
	ParentBlockHash *string `json:"parent_block_hash,omitempty"`
	// True if the transaction is included in a microblock that has not been confirmed by an anchor block.
	IsUnanchored *bool `json:"is_unanchored,omitempty"`
	// The microblock hash that this transaction was streamed in. If the transaction was batched in an anchor block (not included within a microblock) then this value will be an empty string.
	MicroblockHash *string `json:"microblock_hash,omitempty"`
	// The microblock sequence number that this transaction was streamed in. If the transaction was batched in an anchor block (not included within a microblock) then this value will be 2147483647 (0x7fffffff, the max int32 value), this value preserves logical transaction ordering on (block_height, microblock_sequence, tx_index).
	MicroblockSequence *int `json:"microblock_sequence,omitempty"`
	// Set to `true` if microblock is anchored in the canonical chain tip, `false` if the transaction was orphaned in a micro-fork.
	MicroblockCanonical *bool `json:"microblock_canonical,omitempty"`
	// Execution cost read count.
	ExecutionCostReadCount *int `json:"execution_cost_read_count,omitempty"`
	// Execution cost read length.
	ExecutionCostReadLength *int `json:"execution_cost_read_length,omitempty"`
	// Execution cost runtime.
	ExecutionCostRuntime *int `json:"execution_cost_runtime,omitempty"`
	// Execution cost write count.
	ExecutionCostWriteCount *int `json:"execution_cost_write_count,omitempty"`
	// Execution cost write length.
	ExecutionCostWriteLength *int `json:"execution_cost_write_length,omitempty"`
	// List of transaction events
	Events []TransactionEvent1 `json:"events,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for Result2.
// It customizes the JSON marshaling process for Result2 objects.
func (r *Result2) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the Result2 object to a map representation for JSON marshaling.
func (r *Result2) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["tx_id"] = r.TxId
	if r.Nonce != nil {
		structMap["nonce"] = r.Nonce
	}
	if r.FeeRate != nil {
		structMap["fee_rate"] = r.FeeRate
	}
	if r.SenderAddress != nil {
		structMap["sender_address"] = r.SenderAddress
	}
	if r.SponsorNonce != nil {
		structMap["sponsor_nonce"] = r.SponsorNonce
	}
	if r.Sponsored != nil {
		structMap["sponsored"] = r.Sponsored
	}
	if r.SponsorAddress != nil {
		structMap["sponsor_address"] = r.SponsorAddress
	}
	if r.PostConditionMode != nil {
		structMap["post_condition_mode"] = r.PostConditionMode
	}
	if r.PostConditions != nil {
		structMap["post_conditions"] = r.PostConditions
	}
	if r.AnchorMode != nil {
		structMap["anchor_mode"] = r.AnchorMode
	}
	if r.TxStatus != nil {
		structMap["tx_status"] = r.TxStatus
	}
	if r.ReceiptTime != nil {
		structMap["receipt_time"] = r.ReceiptTime
	}
	if r.ReceiptTimeIso != nil {
		structMap["receipt_time_iso"] = r.ReceiptTimeIso
	}
	if r.TxType != nil {
		structMap["tx_type"] = r.TxType
	}
	if r.TokenTransfer != nil {
		structMap["token_transfer"] = r.TokenTransfer
	}
	if r.SmartContract != nil {
		structMap["smart_contract"] = r.SmartContract
	}
	if r.ContractCall != nil {
		structMap["contract_call"] = r.ContractCall
	}
	if r.PoisonMicroblock != nil {
		structMap["poison_microblock"] = r.PoisonMicroblock
	}
	if r.CoinbasePayload != nil {
		structMap["coinbase_payload"] = r.CoinbasePayload
	}
	if r.TenureChangePayload != nil {
		structMap["tenure_change_payload"] = r.TenureChangePayload
	}
	if r.BlockHash != nil {
		structMap["block_hash"] = r.BlockHash
	}
	if r.BlockHeight != nil {
		structMap["block_height"] = r.BlockHeight
	}
	if r.BurnBlockTime != nil {
		structMap["burn_block_time"] = r.BurnBlockTime
	}
	if r.BurnBlockTimeIso != nil {
		structMap["burn_block_time_iso"] = r.BurnBlockTimeIso
	}
	if r.ParentBurnBlockTime != nil {
		structMap["parent_burn_block_time"] = r.ParentBurnBlockTime
	}
	if r.ParentBurnBlockTimeIso != nil {
		structMap["parent_burn_block_time_iso"] = r.ParentBurnBlockTimeIso
	}
	if r.Canonical != nil {
		structMap["canonical"] = r.Canonical
	}
	if r.TxIndex != nil {
		structMap["tx_index"] = r.TxIndex
	}
	if r.TxResult != nil {
		structMap["tx_result"] = r.TxResult
	}
	if r.EventCount != nil {
		structMap["event_count"] = r.EventCount
	}
	if r.ParentBlockHash != nil {
		structMap["parent_block_hash"] = r.ParentBlockHash
	}
	if r.IsUnanchored != nil {
		structMap["is_unanchored"] = r.IsUnanchored
	}
	if r.MicroblockHash != nil {
		structMap["microblock_hash"] = r.MicroblockHash
	}
	if r.MicroblockSequence != nil {
		structMap["microblock_sequence"] = r.MicroblockSequence
	}
	if r.MicroblockCanonical != nil {
		structMap["microblock_canonical"] = r.MicroblockCanonical
	}
	if r.ExecutionCostReadCount != nil {
		structMap["execution_cost_read_count"] = r.ExecutionCostReadCount
	}
	if r.ExecutionCostReadLength != nil {
		structMap["execution_cost_read_length"] = r.ExecutionCostReadLength
	}
	if r.ExecutionCostRuntime != nil {
		structMap["execution_cost_runtime"] = r.ExecutionCostRuntime
	}
	if r.ExecutionCostWriteCount != nil {
		structMap["execution_cost_write_count"] = r.ExecutionCostWriteCount
	}
	if r.ExecutionCostWriteLength != nil {
		structMap["execution_cost_write_length"] = r.ExecutionCostWriteLength
	}
	if r.Events != nil {
		structMap["events"] = r.Events
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Result2.
// It customizes the JSON unmarshaling process for Result2 objects.
func (r *Result2) UnmarshalJSON(input []byte) error {
	temp := &struct {
		TxId                     string                         `json:"tx_id"`
		Nonce                    *int                           `json:"nonce,omitempty"`
		FeeRate                  *string                        `json:"fee_rate,omitempty"`
		SenderAddress            *string                        `json:"sender_address,omitempty"`
		SponsorNonce             *int                           `json:"sponsor_nonce,omitempty"`
		Sponsored                *bool                          `json:"sponsored,omitempty"`
		SponsorAddress           *string                        `json:"sponsor_address,omitempty"`
		PostConditionMode        *PostConditionModeEnum         `json:"post_condition_mode,omitempty"`
		PostConditions           []PostCondition1               `json:"post_conditions,omitempty"`
		AnchorMode               *TransactionAnchorModeTypeEnum `json:"anchor_mode,omitempty"`
		TxStatus                 *MempoolTransactionStatusEnum  `json:"tx_status,omitempty"`
		ReceiptTime              *float64                       `json:"receipt_time,omitempty"`
		ReceiptTimeIso           *string                        `json:"receipt_time_iso,omitempty"`
		TxType                   *TxTypeEnum                    `json:"tx_type,omitempty"`
		TokenTransfer            *TokenTransfer                 `json:"token_transfer,omitempty"`
		SmartContract            *SmartContract                 `json:"smart_contract,omitempty"`
		ContractCall             *ContractCall                  `json:"contract_call,omitempty"`
		PoisonMicroblock         *PoisonMicroblock              `json:"poison_microblock,omitempty"`
		CoinbasePayload          *CoinbasePayload               `json:"coinbase_payload,omitempty"`
		TenureChangePayload      *TenureChangePayload           `json:"tenure_change_payload,omitempty"`
		BlockHash                *string                        `json:"block_hash,omitempty"`
		BlockHeight              *int                           `json:"block_height,omitempty"`
		BurnBlockTime            *int                           `json:"burn_block_time,omitempty"`
		BurnBlockTimeIso         *string                        `json:"burn_block_time_iso,omitempty"`
		ParentBurnBlockTime      *int                           `json:"parent_burn_block_time,omitempty"`
		ParentBurnBlockTimeIso   *string                        `json:"parent_burn_block_time_iso,omitempty"`
		Canonical                *bool                          `json:"canonical,omitempty"`
		TxIndex                  *int                           `json:"tx_index,omitempty"`
		TxResult                 *TxResult                      `json:"tx_result,omitempty"`
		EventCount               *int                           `json:"event_count,omitempty"`
		ParentBlockHash          *string                        `json:"parent_block_hash,omitempty"`
		IsUnanchored             *bool                          `json:"is_unanchored,omitempty"`
		MicroblockHash           *string                        `json:"microblock_hash,omitempty"`
		MicroblockSequence       *int                           `json:"microblock_sequence,omitempty"`
		MicroblockCanonical      *bool                          `json:"microblock_canonical,omitempty"`
		ExecutionCostReadCount   *int                           `json:"execution_cost_read_count,omitempty"`
		ExecutionCostReadLength  *int                           `json:"execution_cost_read_length,omitempty"`
		ExecutionCostRuntime     *int                           `json:"execution_cost_runtime,omitempty"`
		ExecutionCostWriteCount  *int                           `json:"execution_cost_write_count,omitempty"`
		ExecutionCostWriteLength *int                           `json:"execution_cost_write_length,omitempty"`
		Events                   []TransactionEvent1            `json:"events,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	r.TxId = temp.TxId
	r.Nonce = temp.Nonce
	r.FeeRate = temp.FeeRate
	r.SenderAddress = temp.SenderAddress
	r.SponsorNonce = temp.SponsorNonce
	r.Sponsored = temp.Sponsored
	r.SponsorAddress = temp.SponsorAddress
	r.PostConditionMode = temp.PostConditionMode
	r.PostConditions = temp.PostConditions
	r.AnchorMode = temp.AnchorMode
	r.TxStatus = temp.TxStatus
	r.ReceiptTime = temp.ReceiptTime
	r.ReceiptTimeIso = temp.ReceiptTimeIso
	r.TxType = temp.TxType
	r.TokenTransfer = temp.TokenTransfer
	r.SmartContract = temp.SmartContract
	r.ContractCall = temp.ContractCall
	r.PoisonMicroblock = temp.PoisonMicroblock
	r.CoinbasePayload = temp.CoinbasePayload
	r.TenureChangePayload = temp.TenureChangePayload
	r.BlockHash = temp.BlockHash
	r.BlockHeight = temp.BlockHeight
	r.BurnBlockTime = temp.BurnBlockTime
	r.BurnBlockTimeIso = temp.BurnBlockTimeIso
	r.ParentBurnBlockTime = temp.ParentBurnBlockTime
	r.ParentBurnBlockTimeIso = temp.ParentBurnBlockTimeIso
	r.Canonical = temp.Canonical
	r.TxIndex = temp.TxIndex
	r.TxResult = temp.TxResult
	r.EventCount = temp.EventCount
	r.ParentBlockHash = temp.ParentBlockHash
	r.IsUnanchored = temp.IsUnanchored
	r.MicroblockHash = temp.MicroblockHash
	r.MicroblockSequence = temp.MicroblockSequence
	r.MicroblockCanonical = temp.MicroblockCanonical
	r.ExecutionCostReadCount = temp.ExecutionCostReadCount
	r.ExecutionCostReadLength = temp.ExecutionCostReadLength
	r.ExecutionCostRuntime = temp.ExecutionCostRuntime
	r.ExecutionCostWriteCount = temp.ExecutionCostWriteCount
	r.ExecutionCostWriteLength = temp.ExecutionCostWriteLength
	r.Events = temp.Events
	return nil
}
