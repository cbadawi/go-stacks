package models

import (
	"encoding/json"
)

// Block1 represents a Block1 struct.
// A block
type Block1 struct {
	// Set to `true` if block corresponds to the canonical chain tip
	Canonical bool `json:"canonical"`
	// Height of the block
	Height int `json:"height"`
	// Hash representing the block
	Hash string `json:"hash"`
	// The only hash that can uniquely identify an anchored block or an unconfirmed state trie
	IndexBlockHash string `json:"index_block_hash"`
	// Hash of the parent block
	ParentBlockHash string `json:"parent_block_hash"`
	// Unix timestamp (in seconds) indicating when this block was mined.
	BurnBlockTime float64 `json:"burn_block_time"`
	// An ISO 8601 (YYYY-MM-DDTHH:mm:ss.sssZ) indicating when this block was mined.
	BurnBlockTimeIso string `json:"burn_block_time_iso"`
	// Hash of the anchor chain block
	BurnBlockHash string `json:"burn_block_hash"`
	// Height of the anchor chain block
	BurnBlockHeight int `json:"burn_block_height"`
	// Anchor chain transaction ID
	MinerTxid string `json:"miner_txid"`
	// The hash of the last streamed block that precedes this block to which this block is to be appended. Not every anchored block will have a parent microblock stream. An anchored block that does not have a parent microblock stream has the parent microblock hash set to an empty string, and the parent microblock sequence number set to -1.
	ParentMicroblockHash string `json:"parent_microblock_hash"`
	// The hash of the last streamed block that precedes this block to which this block is to be appended. Not every anchored block will have a parent microblock stream. An anchored block that does not have a parent microblock stream has the parent microblock hash set to an empty string, and the parent microblock sequence number set to -1.
	ParentMicroblockSequence int `json:"parent_microblock_sequence"`
	// List of transactions included in the block
	Txs []string `json:"txs"`
	// List of microblocks that were accepted in this anchor block. Not every anchored block will have a accepted all (or any) of the previously streamed microblocks. Microblocks that were orphaned are not included in this list.
	MicroblocksAccepted []string `json:"microblocks_accepted"`
	// List of microblocks that were streamed/produced by this anchor block's miner. This list only includes microblocks that were accepted in the following anchor block. Microblocks that were orphaned are not included in this list.
	MicroblocksStreamed []string `json:"microblocks_streamed"`
	// Execution cost read count.
	ExecutionCostReadCount int `json:"execution_cost_read_count"`
	// Execution cost read length.
	ExecutionCostReadLength int `json:"execution_cost_read_length"`
	// Execution cost runtime.
	ExecutionCostRuntime int `json:"execution_cost_runtime"`
	// Execution cost write count.
	ExecutionCostWriteCount int `json:"execution_cost_write_count"`
	// Execution cost write length.
	ExecutionCostWriteLength int `json:"execution_cost_write_length"`
	// List of txs counts in each accepted microblock
	MicroblockTxCount map[string]float64 `json:"microblock_tx_count"`
	// Transaction ID
	TxId string `json:"tx_id"`
	// Used for ordering the transactions originating from and paying from an account. The nonce ensures that a transaction is processed at most once. The nonce counts the number of times an account's owner(s) have authorized a transaction. The first transaction from an account will have a nonce value equal to 0, the second will have a nonce value equal to 1, and so on.
	Nonce int `json:"nonce"`
	// Transaction fee as Integer string (64-bit unsigned integer).
	FeeRate string `json:"fee_rate"`
	// Address of the transaction initiator
	SenderAddress string `json:"sender_address"`
	SponsorNonce  *int   `json:"sponsor_nonce,omitempty"`
	// Denotes whether the originating account is the same as the paying account
	Sponsored         bool                  `json:"sponsored"`
	SponsorAddress    *string               `json:"sponsor_address,omitempty"`
	PostConditionMode PostConditionModeEnum `json:"post_condition_mode"`
	PostConditions    []PostCondition1      `json:"post_conditions"`
	// `on_chain_only`: the transaction MUST be included in an anchored block, `off_chain_only`: the transaction MUST be included in a microblock, `any`: the leader can choose where to include the transaction.
	AnchorMode TransactionAnchorModeTypeEnum `json:"anchor_mode"`
	// Status of the transaction
	TxStatus MempoolTransactionStatusEnum `json:"tx_status"`
	// A unix timestamp (in seconds) indicating when the transaction broadcast was received by the node.
	ReceiptTime float64 `json:"receipt_time"`
	// An ISO 8601 (YYYY-MM-DDTHH:mm:ss.sssZ) timestamp indicating when the transaction broadcast was received by the node.
	ReceiptTimeIso      string               `json:"receipt_time_iso"`
	TxType              string               `json:"tx_type"`
	TokenTransfer       TokenTransfer        `json:"token_transfer"`
	SmartContract       SmartContract        `json:"smart_contract"`
	ContractCall        ContractCall         `json:"contract_call"`
	PoisonMicroblock    PoisonMicroblock     `json:"poison_microblock"`
	CoinbasePayload     CoinbasePayload      `json:"coinbase_payload"`
	TenureChangePayload *TenureChangePayload `json:"tenure_change_payload,omitempty"`
	// Hash of the blocked this transactions was associated with
	BlockHash string `json:"block_hash"`
	// Height of the block this transactions was associated with
	BlockHeight int `json:"block_height"`
	// Unix timestamp (in seconds) indicating when this parent block was mined
	ParentBurnBlockTime int `json:"parent_burn_block_time"`
	// An ISO 8601 (YYYY-MM-DDTHH:mm:ss.sssZ) timestamp indicating when this parent block was mined.
	ParentBurnBlockTimeIso string `json:"parent_burn_block_time_iso"`
	// Index of the transaction, indicating the order. Starts at `0` and increases with each transaction
	TxIndex int `json:"tx_index"`
	// Result of the transaction. For contract calls, this will show the value returned by the call. For other transaction types, this will return a boolean indicating the success of the transaction.
	TxResult TxResult `json:"tx_result"`
	// Number of transaction events
	EventCount int `json:"event_count"`
	// True if the transaction is included in a microblock that has not been confirmed by an anchor block.
	IsUnanchored bool `json:"is_unanchored"`
	// The microblock hash that this transaction was streamed in. If the transaction was batched in an anchor block (not included within a microblock) then this value will be an empty string.
	MicroblockHash string `json:"microblock_hash"`
	// The microblock sequence number that this transaction was streamed in. If the transaction was batched in an anchor block (not included within a microblock) then this value will be 2147483647 (0x7fffffff, the max int32 value), this value preserves logical transaction ordering on (block_height, microblock_sequence, tx_index).
	MicroblockSequence int `json:"microblock_sequence"`
	// Set to `true` if microblock is anchored in the canonical chain tip, `false` if the transaction was orphaned in a micro-fork.
	MicroblockCanonical bool `json:"microblock_canonical"`
	// List of transaction events
	Events []TransactionEvent1 `json:"events"`
}

// MarshalJSON implements the json.Marshaler interface for Block1.
// It customizes the JSON marshaling process for Block1 objects.
func (b *Block1) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(b.toMap())
}

// toMap converts the Block1 object to a map representation for JSON marshaling.
func (b *Block1) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["canonical"] = b.Canonical
	structMap["height"] = b.Height
	structMap["hash"] = b.Hash
	structMap["index_block_hash"] = b.IndexBlockHash
	structMap["parent_block_hash"] = b.ParentBlockHash
	structMap["burn_block_time"] = b.BurnBlockTime
	structMap["burn_block_time_iso"] = b.BurnBlockTimeIso
	structMap["burn_block_hash"] = b.BurnBlockHash
	structMap["burn_block_height"] = b.BurnBlockHeight
	structMap["miner_txid"] = b.MinerTxid
	structMap["parent_microblock_hash"] = b.ParentMicroblockHash
	structMap["parent_microblock_sequence"] = b.ParentMicroblockSequence
	structMap["txs"] = b.Txs
	structMap["microblocks_accepted"] = b.MicroblocksAccepted
	structMap["microblocks_streamed"] = b.MicroblocksStreamed
	structMap["execution_cost_read_count"] = b.ExecutionCostReadCount
	structMap["execution_cost_read_length"] = b.ExecutionCostReadLength
	structMap["execution_cost_runtime"] = b.ExecutionCostRuntime
	structMap["execution_cost_write_count"] = b.ExecutionCostWriteCount
	structMap["execution_cost_write_length"] = b.ExecutionCostWriteLength
	structMap["microblock_tx_count"] = b.MicroblockTxCount
	structMap["tx_id"] = b.TxId
	structMap["nonce"] = b.Nonce
	structMap["fee_rate"] = b.FeeRate
	structMap["sender_address"] = b.SenderAddress
	if b.SponsorNonce != nil {
		structMap["sponsor_nonce"] = b.SponsorNonce
	}
	structMap["sponsored"] = b.Sponsored
	if b.SponsorAddress != nil {
		structMap["sponsor_address"] = b.SponsorAddress
	}
	structMap["post_condition_mode"] = b.PostConditionMode
	structMap["post_conditions"] = b.PostConditions
	structMap["anchor_mode"] = b.AnchorMode
	structMap["tx_status"] = b.TxStatus
	structMap["receipt_time"] = b.ReceiptTime
	structMap["receipt_time_iso"] = b.ReceiptTimeIso
	structMap["tx_type"] = b.TxType
	structMap["token_transfer"] = b.TokenTransfer
	structMap["smart_contract"] = b.SmartContract
	structMap["contract_call"] = b.ContractCall
	structMap["poison_microblock"] = b.PoisonMicroblock
	structMap["coinbase_payload"] = b.CoinbasePayload
	if b.TenureChangePayload != nil {
		structMap["tenure_change_payload"] = b.TenureChangePayload
	}
	structMap["block_hash"] = b.BlockHash
	structMap["block_height"] = b.BlockHeight
	structMap["parent_burn_block_time"] = b.ParentBurnBlockTime
	structMap["parent_burn_block_time_iso"] = b.ParentBurnBlockTimeIso
	structMap["tx_index"] = b.TxIndex
	structMap["tx_result"] = b.TxResult
	structMap["event_count"] = b.EventCount
	structMap["is_unanchored"] = b.IsUnanchored
	structMap["microblock_hash"] = b.MicroblockHash
	structMap["microblock_sequence"] = b.MicroblockSequence
	structMap["microblock_canonical"] = b.MicroblockCanonical
	structMap["events"] = b.Events
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Block1.
// It customizes the JSON unmarshaling process for Block1 objects.
func (b *Block1) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Canonical                bool                          `json:"canonical"`
		Height                   int                           `json:"height"`
		Hash                     string                        `json:"hash"`
		IndexBlockHash           string                        `json:"index_block_hash"`
		ParentBlockHash          string                        `json:"parent_block_hash"`
		BurnBlockTime            float64                       `json:"burn_block_time"`
		BurnBlockTimeIso         string                        `json:"burn_block_time_iso"`
		BurnBlockHash            string                        `json:"burn_block_hash"`
		BurnBlockHeight          int                           `json:"burn_block_height"`
		MinerTxid                string                        `json:"miner_txid"`
		ParentMicroblockHash     string                        `json:"parent_microblock_hash"`
		ParentMicroblockSequence int                           `json:"parent_microblock_sequence"`
		Txs                      []string                      `json:"txs"`
		MicroblocksAccepted      []string                      `json:"microblocks_accepted"`
		MicroblocksStreamed      []string                      `json:"microblocks_streamed"`
		ExecutionCostReadCount   int                           `json:"execution_cost_read_count"`
		ExecutionCostReadLength  int                           `json:"execution_cost_read_length"`
		ExecutionCostRuntime     int                           `json:"execution_cost_runtime"`
		ExecutionCostWriteCount  int                           `json:"execution_cost_write_count"`
		ExecutionCostWriteLength int                           `json:"execution_cost_write_length"`
		MicroblockTxCount        map[string]float64            `json:"microblock_tx_count"`
		TxId                     string                        `json:"tx_id"`
		Nonce                    int                           `json:"nonce"`
		FeeRate                  string                        `json:"fee_rate"`
		SenderAddress            string                        `json:"sender_address"`
		SponsorNonce             *int                          `json:"sponsor_nonce,omitempty"`
		Sponsored                bool                          `json:"sponsored"`
		SponsorAddress           *string                       `json:"sponsor_address,omitempty"`
		PostConditionMode        PostConditionModeEnum         `json:"post_condition_mode"`
		PostConditions           []PostCondition1              `json:"post_conditions"`
		AnchorMode               TransactionAnchorModeTypeEnum `json:"anchor_mode"`
		TxStatus                 MempoolTransactionStatusEnum  `json:"tx_status"`
		ReceiptTime              float64                       `json:"receipt_time"`
		ReceiptTimeIso           string                        `json:"receipt_time_iso"`
		TxType                   string                        `json:"tx_type"`
		TokenTransfer            TokenTransfer                 `json:"token_transfer"`
		SmartContract            SmartContract                 `json:"smart_contract"`
		ContractCall             ContractCall                  `json:"contract_call"`
		PoisonMicroblock         PoisonMicroblock              `json:"poison_microblock"`
		CoinbasePayload          CoinbasePayload               `json:"coinbase_payload"`
		TenureChangePayload      *TenureChangePayload          `json:"tenure_change_payload,omitempty"`
		BlockHash                string                        `json:"block_hash"`
		BlockHeight              int                           `json:"block_height"`
		ParentBurnBlockTime      int                           `json:"parent_burn_block_time"`
		ParentBurnBlockTimeIso   string                        `json:"parent_burn_block_time_iso"`
		TxIndex                  int                           `json:"tx_index"`
		TxResult                 TxResult                      `json:"tx_result"`
		EventCount               int                           `json:"event_count"`
		IsUnanchored             bool                          `json:"is_unanchored"`
		MicroblockHash           string                        `json:"microblock_hash"`
		MicroblockSequence       int                           `json:"microblock_sequence"`
		MicroblockCanonical      bool                          `json:"microblock_canonical"`
		Events                   []TransactionEvent1           `json:"events"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	b.Canonical = temp.Canonical
	b.Height = temp.Height
	b.Hash = temp.Hash
	b.IndexBlockHash = temp.IndexBlockHash
	b.ParentBlockHash = temp.ParentBlockHash
	b.BurnBlockTime = temp.BurnBlockTime
	b.BurnBlockTimeIso = temp.BurnBlockTimeIso
	b.BurnBlockHash = temp.BurnBlockHash
	b.BurnBlockHeight = temp.BurnBlockHeight
	b.MinerTxid = temp.MinerTxid
	b.ParentMicroblockHash = temp.ParentMicroblockHash
	b.ParentMicroblockSequence = temp.ParentMicroblockSequence
	b.Txs = temp.Txs
	b.MicroblocksAccepted = temp.MicroblocksAccepted
	b.MicroblocksStreamed = temp.MicroblocksStreamed
	b.ExecutionCostReadCount = temp.ExecutionCostReadCount
	b.ExecutionCostReadLength = temp.ExecutionCostReadLength
	b.ExecutionCostRuntime = temp.ExecutionCostRuntime
	b.ExecutionCostWriteCount = temp.ExecutionCostWriteCount
	b.ExecutionCostWriteLength = temp.ExecutionCostWriteLength
	b.MicroblockTxCount = temp.MicroblockTxCount
	b.TxId = temp.TxId
	b.Nonce = temp.Nonce
	b.FeeRate = temp.FeeRate
	b.SenderAddress = temp.SenderAddress
	b.SponsorNonce = temp.SponsorNonce
	b.Sponsored = temp.Sponsored
	b.SponsorAddress = temp.SponsorAddress
	b.PostConditionMode = temp.PostConditionMode
	b.PostConditions = temp.PostConditions
	b.AnchorMode = temp.AnchorMode
	b.TxStatus = temp.TxStatus
	b.ReceiptTime = temp.ReceiptTime
	b.ReceiptTimeIso = temp.ReceiptTimeIso
	b.TxType = temp.TxType
	b.TokenTransfer = temp.TokenTransfer
	b.SmartContract = temp.SmartContract
	b.ContractCall = temp.ContractCall
	b.PoisonMicroblock = temp.PoisonMicroblock
	b.CoinbasePayload = temp.CoinbasePayload
	b.TenureChangePayload = temp.TenureChangePayload
	b.BlockHash = temp.BlockHash
	b.BlockHeight = temp.BlockHeight
	b.ParentBurnBlockTime = temp.ParentBurnBlockTime
	b.ParentBurnBlockTimeIso = temp.ParentBurnBlockTimeIso
	b.TxIndex = temp.TxIndex
	b.TxResult = temp.TxResult
	b.EventCount = temp.EventCount
	b.IsUnanchored = temp.IsUnanchored
	b.MicroblockHash = temp.MicroblockHash
	b.MicroblockSequence = temp.MicroblockSequence
	b.MicroblockCanonical = temp.MicroblockCanonical
	b.Events = temp.Events
	return nil
}
