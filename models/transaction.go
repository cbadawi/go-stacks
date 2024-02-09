package models

import (
	"encoding/json"
)

// Transaction represents a Transaction struct.
// Describes all transaction types on Stacks 2.0 blockchain
type Transaction struct {
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
	// Status of the transaction
	TxStatus *TransactionStatusEnum `json:"tx_status,omitempty"`
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
	Events              []TransactionEvent1  `json:"events,omitempty"`
	TxType              *TxTypeEnum          `json:"tx_type,omitempty"`
	TokenTransfer       *TokenTransfer       `json:"token_transfer,omitempty"`
	SmartContract       *SmartContract       `json:"smart_contract,omitempty"`
	ContractCall        *ContractCall        `json:"contract_call,omitempty"`
	PoisonMicroblock    *PoisonMicroblock    `json:"poison_microblock,omitempty"`
	CoinbasePayload     *CoinbasePayload     `json:"coinbase_payload,omitempty"`
	TenureChangePayload *TenureChangePayload `json:"tenure_change_payload,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for Transaction.
// It customizes the JSON marshaling process for Transaction objects.
func (t *Transaction) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(t.toMap())
}

// toMap converts the Transaction object to a map representation for JSON marshaling.
func (t *Transaction) toMap() map[string]any {
	structMap := make(map[string]any)
	if t.TxId != nil {
		structMap["tx_id"] = t.TxId
	}
	if t.Nonce != nil {
		structMap["nonce"] = t.Nonce
	}
	if t.FeeRate != nil {
		structMap["fee_rate"] = t.FeeRate
	}
	if t.SenderAddress != nil {
		structMap["sender_address"] = t.SenderAddress
	}
	if t.SponsorNonce != nil {
		structMap["sponsor_nonce"] = t.SponsorNonce
	}
	if t.Sponsored != nil {
		structMap["sponsored"] = t.Sponsored
	}
	if t.SponsorAddress != nil {
		structMap["sponsor_address"] = t.SponsorAddress
	}
	if t.PostConditionMode != nil {
		structMap["post_condition_mode"] = t.PostConditionMode
	}
	if t.PostConditions != nil {
		structMap["post_conditions"] = t.PostConditions
	}
	if t.AnchorMode != nil {
		structMap["anchor_mode"] = t.AnchorMode
	}
	if t.BlockHash != nil {
		structMap["block_hash"] = t.BlockHash
	}
	if t.BlockHeight != nil {
		structMap["block_height"] = t.BlockHeight
	}
	if t.BurnBlockTime != nil {
		structMap["burn_block_time"] = t.BurnBlockTime
	}
	if t.BurnBlockTimeIso != nil {
		structMap["burn_block_time_iso"] = t.BurnBlockTimeIso
	}
	if t.ParentBurnBlockTime != nil {
		structMap["parent_burn_block_time"] = t.ParentBurnBlockTime
	}
	if t.ParentBurnBlockTimeIso != nil {
		structMap["parent_burn_block_time_iso"] = t.ParentBurnBlockTimeIso
	}
	if t.Canonical != nil {
		structMap["canonical"] = t.Canonical
	}
	if t.TxIndex != nil {
		structMap["tx_index"] = t.TxIndex
	}
	if t.TxStatus != nil {
		structMap["tx_status"] = t.TxStatus
	}
	if t.TxResult != nil {
		structMap["tx_result"] = t.TxResult
	}
	if t.EventCount != nil {
		structMap["event_count"] = t.EventCount
	}
	if t.ParentBlockHash != nil {
		structMap["parent_block_hash"] = t.ParentBlockHash
	}
	if t.IsUnanchored != nil {
		structMap["is_unanchored"] = t.IsUnanchored
	}
	if t.MicroblockHash != nil {
		structMap["microblock_hash"] = t.MicroblockHash
	}
	if t.MicroblockSequence != nil {
		structMap["microblock_sequence"] = t.MicroblockSequence
	}
	if t.MicroblockCanonical != nil {
		structMap["microblock_canonical"] = t.MicroblockCanonical
	}
	if t.ExecutionCostReadCount != nil {
		structMap["execution_cost_read_count"] = t.ExecutionCostReadCount
	}
	if t.ExecutionCostReadLength != nil {
		structMap["execution_cost_read_length"] = t.ExecutionCostReadLength
	}
	if t.ExecutionCostRuntime != nil {
		structMap["execution_cost_runtime"] = t.ExecutionCostRuntime
	}
	if t.ExecutionCostWriteCount != nil {
		structMap["execution_cost_write_count"] = t.ExecutionCostWriteCount
	}
	if t.ExecutionCostWriteLength != nil {
		structMap["execution_cost_write_length"] = t.ExecutionCostWriteLength
	}
	if t.Events != nil {
		structMap["events"] = t.Events
	}
	if t.TxType != nil {
		structMap["tx_type"] = t.TxType
	}
	if t.TokenTransfer != nil {
		structMap["token_transfer"] = t.TokenTransfer
	}
	if t.SmartContract != nil {
		structMap["smart_contract"] = t.SmartContract
	}
	if t.ContractCall != nil {
		structMap["contract_call"] = t.ContractCall
	}
	if t.PoisonMicroblock != nil {
		structMap["poison_microblock"] = t.PoisonMicroblock
	}
	if t.CoinbasePayload != nil {
		structMap["coinbase_payload"] = t.CoinbasePayload
	}
	if t.TenureChangePayload != nil {
		structMap["tenure_change_payload"] = t.TenureChangePayload
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Transaction.
// It customizes the JSON unmarshaling process for Transaction objects.
func (t *Transaction) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		TxId                     *string                        `json:"tx_id,omitempty"`
		Nonce                    *int                           `json:"nonce,omitempty"`
		FeeRate                  *string                        `json:"fee_rate,omitempty"`
		SenderAddress            *string                        `json:"sender_address,omitempty"`
		SponsorNonce             *int                           `json:"sponsor_nonce,omitempty"`
		Sponsored                *bool                          `json:"sponsored,omitempty"`
		SponsorAddress           *string                        `json:"sponsor_address,omitempty"`
		PostConditionMode        *PostConditionModeEnum         `json:"post_condition_mode,omitempty"`
		PostConditions           []PostCondition1               `json:"post_conditions,omitempty"`
		AnchorMode               *TransactionAnchorModeTypeEnum `json:"anchor_mode,omitempty"`
		BlockHash                *string                        `json:"block_hash,omitempty"`
		BlockHeight              *int                           `json:"block_height,omitempty"`
		BurnBlockTime            *int                           `json:"burn_block_time,omitempty"`
		BurnBlockTimeIso         *string                        `json:"burn_block_time_iso,omitempty"`
		ParentBurnBlockTime      *int                           `json:"parent_burn_block_time,omitempty"`
		ParentBurnBlockTimeIso   *string                        `json:"parent_burn_block_time_iso,omitempty"`
		Canonical                *bool                          `json:"canonical,omitempty"`
		TxIndex                  *int                           `json:"tx_index,omitempty"`
		TxStatus                 *TransactionStatusEnum         `json:"tx_status,omitempty"`
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
		TxType                   *TxTypeEnum                    `json:"tx_type,omitempty"`
		TokenTransfer            *TokenTransfer                 `json:"token_transfer,omitempty"`
		SmartContract            *SmartContract                 `json:"smart_contract,omitempty"`
		ContractCall             *ContractCall                  `json:"contract_call,omitempty"`
		PoisonMicroblock         *PoisonMicroblock              `json:"poison_microblock,omitempty"`
		CoinbasePayload          *CoinbasePayload               `json:"coinbase_payload,omitempty"`
		TenureChangePayload      *TenureChangePayload           `json:"tenure_change_payload,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	t.TxId = temp.TxId
	t.Nonce = temp.Nonce
	t.FeeRate = temp.FeeRate
	t.SenderAddress = temp.SenderAddress
	t.SponsorNonce = temp.SponsorNonce
	t.Sponsored = temp.Sponsored
	t.SponsorAddress = temp.SponsorAddress
	t.PostConditionMode = temp.PostConditionMode
	t.PostConditions = temp.PostConditions
	t.AnchorMode = temp.AnchorMode
	t.BlockHash = temp.BlockHash
	t.BlockHeight = temp.BlockHeight
	t.BurnBlockTime = temp.BurnBlockTime
	t.BurnBlockTimeIso = temp.BurnBlockTimeIso
	t.ParentBurnBlockTime = temp.ParentBurnBlockTime
	t.ParentBurnBlockTimeIso = temp.ParentBurnBlockTimeIso
	t.Canonical = temp.Canonical
	t.TxIndex = temp.TxIndex
	t.TxStatus = temp.TxStatus
	t.TxResult = temp.TxResult
	t.EventCount = temp.EventCount
	t.ParentBlockHash = temp.ParentBlockHash
	t.IsUnanchored = temp.IsUnanchored
	t.MicroblockHash = temp.MicroblockHash
	t.MicroblockSequence = temp.MicroblockSequence
	t.MicroblockCanonical = temp.MicroblockCanonical
	t.ExecutionCostReadCount = temp.ExecutionCostReadCount
	t.ExecutionCostReadLength = temp.ExecutionCostReadLength
	t.ExecutionCostRuntime = temp.ExecutionCostRuntime
	t.ExecutionCostWriteCount = temp.ExecutionCostWriteCount
	t.ExecutionCostWriteLength = temp.ExecutionCostWriteLength
	t.Events = temp.Events
	t.TxType = temp.TxType
	t.TokenTransfer = temp.TokenTransfer
	t.SmartContract = temp.SmartContract
	t.ContractCall = temp.ContractCall
	t.PoisonMicroblock = temp.PoisonMicroblock
	t.CoinbasePayload = temp.CoinbasePayload
	t.TenureChangePayload = temp.TenureChangePayload
	return nil
}
