package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the TokenTransferTransaction type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &TokenTransferTransaction{}

// TokenTransferTransaction Describes representation of a Type-0 Stacks 2.0 transaction. https://github.com/blockstack/stacks-blockchain/blob/master/sip/sip-005-blocks-and-transactions.md#type-0-transferring-an-asset
type TokenTransferTransaction struct {
	// Transaction ID
	TxId string `json:"tx_id"`
	// Used for ordering the transactions originating from and paying from an account. The nonce ensures that a transaction is processed at most once. The nonce counts the number of times an account's owner(s) have authorized a transaction. The first transaction from an account will have a nonce value equal to 0, the second will have a nonce value equal to 1, and so on.
	Nonce int32 `json:"nonce"`
	// Transaction fee as Integer string (64-bit unsigned integer).
	FeeRate string `json:"fee_rate"`
	// Address of the transaction initiator
	SenderAddress string `json:"sender_address"`
	SponsorNonce  *int32 `json:"sponsor_nonce,omitempty"`
	// Denotes whether the originating account is the same as the paying account
	Sponsored      bool    `json:"sponsored"`
	SponsorAddress *string `json:"sponsor_address,omitempty"`
	//
	PostConditionMode string          `json:"post_condition_mode"`
	PostConditions    []PostCondition `json:"post_conditions"`
	// `on_chain_only`: the transaction MUST be included in an anchored block, `off_chain_only`: the transaction MUST be included in a microblock, `any`: the leader can choose where to include the transaction.
	AnchorMode string `json:"anchor_mode"`
	// Hash of the blocked this transactions was associated with
	BlockHash string `json:"block_hash"`
	// Height of the block this transactions was associated with
	BlockHeight int32 `json:"block_height"`
	// Unix timestamp (in seconds) indicating when this block was mined.
	BlockTime float32 `json:"block_time"`
	// An ISO 8601 (YYYY-MM-DDTHH:mm:ss.sssZ) indicating when this block was mined.
	BlockTimeIso string `json:"block_time_iso"`
	// Height of the anchor burn block.
	BurnBlockHeight int32 `json:"burn_block_height"`
	// Unix timestamp (in seconds) indicating when this block was mined
	BurnBlockTime int32 `json:"burn_block_time"`
	// An ISO 8601 (YYYY-MM-DDTHH:mm:ss.sssZ) timestamp indicating when this block was mined.
	BurnBlockTimeIso string `json:"burn_block_time_iso"`
	// Unix timestamp (in seconds) indicating when this parent block was mined
	ParentBurnBlockTime int32 `json:"parent_burn_block_time"`
	// An ISO 8601 (YYYY-MM-DDTHH:mm:ss.sssZ) timestamp indicating when this parent block was mined.
	ParentBurnBlockTimeIso string `json:"parent_burn_block_time_iso"`
	// Set to `true` if block corresponds to the canonical chain tip
	Canonical bool `json:"canonical"`
	// Index of the transaction, indicating the order. Starts at `0` and increases with each transaction
	TxIndex int32 `json:"tx_index"`
	// Status of the transaction
	TxStatus string                           `json:"tx_status"`
	TxResult AbstractTransactionAllOfTxResult `json:"tx_result"`
	// Number of transaction events
	EventCount int32 `json:"event_count"`
	// Hash of the previous block.
	ParentBlockHash string `json:"parent_block_hash"`
	// True if the transaction is included in a microblock that has not been confirmed by an anchor block.
	IsUnanchored bool `json:"is_unanchored"`
	// The microblock hash that this transaction was streamed in. If the transaction was batched in an anchor block (not included within a microblock) then this value will be an empty string.
	MicroblockHash string `json:"microblock_hash"`
	// The microblock sequence number that this transaction was streamed in. If the transaction was batched in an anchor block (not included within a microblock) then this value will be 2147483647 (0x7fffffff, the max int32 value), this value preserves logical transaction ordering on (block_height, microblock_sequence, tx_index).
	MicroblockSequence int32 `json:"microblock_sequence"`
	// Set to `true` if microblock is anchored in the canonical chain tip, `false` if the transaction was orphaned in a micro-fork.
	MicroblockCanonical bool `json:"microblock_canonical"`
	// Execution cost read count.
	ExecutionCostReadCount int32 `json:"execution_cost_read_count"`
	// Execution cost read length.
	ExecutionCostReadLength int32 `json:"execution_cost_read_length"`
	// Execution cost runtime.
	ExecutionCostRuntime int32 `json:"execution_cost_runtime"`
	// Execution cost write count.
	ExecutionCostWriteCount int32 `json:"execution_cost_write_count"`
	// Execution cost write length.
	ExecutionCostWriteLength int32 `json:"execution_cost_write_length"`
	// List of transaction events
	Events        []TransactionEvent                            `json:"events"`
	TxType        string                                        `json:"tx_type"`
	TokenTransfer TokenTransferTransactionMetadataTokenTransfer `json:"token_transfer"`
}

type _TokenTransferTransaction TokenTransferTransaction

// NewTokenTransferTransaction instantiates a new TokenTransferTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenTransferTransaction(txId string, nonce int32, feeRate string, senderAddress string, sponsored bool, postConditionMode string, postConditions []PostCondition, anchorMode string, blockHash string, blockHeight int32, blockTime float32, blockTimeIso string, burnBlockHeight int32, burnBlockTime int32, burnBlockTimeIso string, parentBurnBlockTime int32, parentBurnBlockTimeIso string, canonical bool, txIndex int32, txStatus string, txResult AbstractTransactionAllOfTxResult, eventCount int32, parentBlockHash string, isUnanchored bool, microblockHash string, microblockSequence int32, microblockCanonical bool, executionCostReadCount int32, executionCostReadLength int32, executionCostRuntime int32, executionCostWriteCount int32, executionCostWriteLength int32, events []TransactionEvent, txType string, tokenTransfer TokenTransferTransactionMetadataTokenTransfer) *TokenTransferTransaction {
	this := TokenTransferTransaction{}
	this.TxId = txId
	this.Nonce = nonce
	this.FeeRate = feeRate
	this.SenderAddress = senderAddress
	this.Sponsored = sponsored
	this.PostConditionMode = postConditionMode
	this.PostConditions = postConditions
	this.AnchorMode = anchorMode
	this.BlockHash = blockHash
	this.BlockHeight = blockHeight
	this.BlockTime = blockTime
	this.BlockTimeIso = blockTimeIso
	this.BurnBlockHeight = burnBlockHeight
	this.BurnBlockTime = burnBlockTime
	this.BurnBlockTimeIso = burnBlockTimeIso
	this.ParentBurnBlockTime = parentBurnBlockTime
	this.ParentBurnBlockTimeIso = parentBurnBlockTimeIso
	this.Canonical = canonical
	this.TxIndex = txIndex
	this.TxStatus = txStatus
	this.TxResult = txResult
	this.EventCount = eventCount
	this.ParentBlockHash = parentBlockHash
	this.IsUnanchored = isUnanchored
	this.MicroblockHash = microblockHash
	this.MicroblockSequence = microblockSequence
	this.MicroblockCanonical = microblockCanonical
	this.ExecutionCostReadCount = executionCostReadCount
	this.ExecutionCostReadLength = executionCostReadLength
	this.ExecutionCostRuntime = executionCostRuntime
	this.ExecutionCostWriteCount = executionCostWriteCount
	this.ExecutionCostWriteLength = executionCostWriteLength
	this.Events = events
	this.TxType = txType
	this.TokenTransfer = tokenTransfer
	return &this
}

// NewTokenTransferTransactionWithDefaults instantiates a new TokenTransferTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenTransferTransactionWithDefaults() *TokenTransferTransaction {
	this := TokenTransferTransaction{}
	return &this
}

// GetTxId returns the TxId field value
func (o *TokenTransferTransaction) GetTxId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TxId
}

// GetTxIdOk returns a tuple with the TxId field value
// and a boolean to check if the value has been set.
func (o *TokenTransferTransaction) GetTxIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxId, true
}

// SetTxId sets field value
func (o *TokenTransferTransaction) SetTxId(v string) {
	o.TxId = v
}

// GetNonce returns the Nonce field value
func (o *TokenTransferTransaction) GetNonce() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value
// and a boolean to check if the value has been set.
func (o *TokenTransferTransaction) GetNonceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nonce, true
}

// SetNonce sets field value
func (o *TokenTransferTransaction) SetNonce(v int32) {
	o.Nonce = v
}

// GetFeeRate returns the FeeRate field value
func (o *TokenTransferTransaction) GetFeeRate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FeeRate
}

// GetFeeRateOk returns a tuple with the FeeRate field value
// and a boolean to check if the value has been set.
func (o *TokenTransferTransaction) GetFeeRateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeeRate, true
}

// SetFeeRate sets field value
func (o *TokenTransferTransaction) SetFeeRate(v string) {
	o.FeeRate = v
}

// GetSenderAddress returns the SenderAddress field value
func (o *TokenTransferTransaction) GetSenderAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SenderAddress
}

// GetSenderAddressOk returns a tuple with the SenderAddress field value
// and a boolean to check if the value has been set.
func (o *TokenTransferTransaction) GetSenderAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SenderAddress, true
}

// SetSenderAddress sets field value
func (o *TokenTransferTransaction) SetSenderAddress(v string) {
	o.SenderAddress = v
}

// GetSponsorNonce returns the SponsorNonce field value if set, zero value otherwise.
func (o *TokenTransferTransaction) GetSponsorNonce() int32 {
	if o == nil || utils.IsNil(o.SponsorNonce) {
		var ret int32
		return ret
	}
	return *o.SponsorNonce
}

// GetSponsorNonceOk returns a tuple with the SponsorNonce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenTransferTransaction) GetSponsorNonceOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.SponsorNonce) {
		return nil, false
	}
	return o.SponsorNonce, true
}

// HasSponsorNonce returns a boolean if a field has been set.
func (o *TokenTransferTransaction) HasSponsorNonce() bool {
	if o != nil && !utils.IsNil(o.SponsorNonce) {
		return true
	}

	return false
}

// SetSponsorNonce gets a reference to the given int32 and assigns it to the SponsorNonce field.
func (o *TokenTransferTransaction) SetSponsorNonce(v int32) {
	o.SponsorNonce = &v
}

// GetSponsored returns the Sponsored field value
func (o *TokenTransferTransaction) GetSponsored() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Sponsored
}

// GetSponsoredOk returns a tuple with the Sponsored field value
// and a boolean to check if the value has been set.
func (o *TokenTransferTransaction) GetSponsoredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sponsored, true
}

// SetSponsored sets field value
func (o *TokenTransferTransaction) SetSponsored(v bool) {
	o.Sponsored = v
}

// GetSponsorAddress returns the SponsorAddress field value if set, zero value otherwise.
func (o *TokenTransferTransaction) GetSponsorAddress() string {
	if o == nil || utils.IsNil(o.SponsorAddress) {
		var ret string
		return ret
	}
	return *o.SponsorAddress
}

// GetSponsorAddressOk returns a tuple with the SponsorAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenTransferTransaction) GetSponsorAddressOk() (*string, bool) {
	if o == nil || utils.IsNil(o.SponsorAddress) {
		return nil, false
	}
	return o.SponsorAddress, true
}

// HasSponsorAddress returns a boolean if a field has been set.
func (o *TokenTransferTransaction) HasSponsorAddress() bool {
	if o != nil && !utils.IsNil(o.SponsorAddress) {
		return true
	}

	return false
}

// SetSponsorAddress gets a reference to the given string and assigns it to the SponsorAddress field.
func (o *TokenTransferTransaction) SetSponsorAddress(v string) {
	o.SponsorAddress = &v
}

// GetPostConditionMode returns the PostConditionMode field value
func (o *TokenTransferTransaction) GetPostConditionMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PostConditionMode
}

// GetPostConditionModeOk returns a tuple with the PostConditionMode field value
// and a boolean to check if the value has been set.
func (o *TokenTransferTransaction) GetPostConditionModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PostConditionMode, true
}

// SetPostConditionMode sets field value
func (o *TokenTransferTransaction) SetPostConditionMode(v string) {
	o.PostConditionMode = v
}

// GetPostConditions returns the PostConditions field value
func (o *TokenTransferTransaction) GetPostConditions() []PostCondition {
	if o == nil {
		var ret []PostCondition
		return ret
	}

	return o.PostConditions
}

// GetPostConditionsOk returns a tuple with the PostConditions field value
// and a boolean to check if the value has been set.
func (o *TokenTransferTransaction) GetPostConditionsOk() ([]PostCondition, bool) {
	if o == nil {
		return nil, false
	}
	return o.PostConditions, true
}

// SetPostConditions sets field value
func (o *TokenTransferTransaction) SetPostConditions(v []PostCondition) {
	o.PostConditions = v
}

// GetAnchorMode returns the AnchorMode field value
func (o *TokenTransferTransaction) GetAnchorMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AnchorMode
}

// GetAnchorModeOk returns a tuple with the AnchorMode field value
// and a boolean to check if the value has been set.
func (o *TokenTransferTransaction) GetAnchorModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AnchorMode, true
}

// SetAnchorMode sets field value
func (o *TokenTransferTransaction) SetAnchorMode(v string) {
	o.AnchorMode = v
}

// GetBlockHash returns the BlockHash field value
func (o *TokenTransferTransaction) GetBlockHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BlockHash
}

// GetBlockHashOk returns a tuple with the BlockHash field value
// and a boolean to check if the value has been set.
func (o *TokenTransferTransaction) GetBlockHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockHash, true
}

// SetBlockHash sets field value
func (o *TokenTransferTransaction) SetBlockHash(v string) {
	o.BlockHash = v
}

// GetBlockHeight returns the BlockHeight field value
func (o *TokenTransferTransaction) GetBlockHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BlockHeight
}

// GetBlockHeightOk returns a tuple with the BlockHeight field value
// and a boolean to check if the value has been set.
func (o *TokenTransferTransaction) GetBlockHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockHeight, true
}

// SetBlockHeight sets field value
func (o *TokenTransferTransaction) SetBlockHeight(v int32) {
	o.BlockHeight = v
}

// GetBlockTime returns the BlockTime field value
func (o *TokenTransferTransaction) GetBlockTime() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.BlockTime
}

// GetBlockTimeOk returns a tuple with the BlockTime field value
// and a boolean to check if the value has been set.
func (o *TokenTransferTransaction) GetBlockTimeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockTime, true
}

// SetBlockTime sets field value
func (o *TokenTransferTransaction) SetBlockTime(v float32) {
	o.BlockTime = v
}

// GetBlockTimeIso returns the BlockTimeIso field value
func (o *TokenTransferTransaction) GetBlockTimeIso() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BlockTimeIso
}

// GetBlockTimeIsoOk returns a tuple with the BlockTimeIso field value
// and a boolean to check if the value has been set.
func (o *TokenTransferTransaction) GetBlockTimeIsoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockTimeIso, true
}

// SetBlockTimeIso sets field value
func (o *TokenTransferTransaction) SetBlockTimeIso(v string) {
	o.BlockTimeIso = v
}

// GetBurnBlockHeight returns the BurnBlockHeight field value
func (o *TokenTransferTransaction) GetBurnBlockHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BurnBlockHeight
}

// GetBurnBlockHeightOk returns a tuple with the BurnBlockHeight field value
// and a boolean to check if the value has been set.
func (o *TokenTransferTransaction) GetBurnBlockHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BurnBlockHeight, true
}

// SetBurnBlockHeight sets field value
func (o *TokenTransferTransaction) SetBurnBlockHeight(v int32) {
	o.BurnBlockHeight = v
}

// GetBurnBlockTime returns the BurnBlockTime field value
func (o *TokenTransferTransaction) GetBurnBlockTime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BurnBlockTime
}

// GetBurnBlockTimeOk returns a tuple with the BurnBlockTime field value
// and a boolean to check if the value has been set.
func (o *TokenTransferTransaction) GetBurnBlockTimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BurnBlockTime, true
}

// SetBurnBlockTime sets field value
func (o *TokenTransferTransaction) SetBurnBlockTime(v int32) {
	o.BurnBlockTime = v
}

// GetBurnBlockTimeIso returns the BurnBlockTimeIso field value
func (o *TokenTransferTransaction) GetBurnBlockTimeIso() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BurnBlockTimeIso
}

// GetBurnBlockTimeIsoOk returns a tuple with the BurnBlockTimeIso field value
// and a boolean to check if the value has been set.
func (o *TokenTransferTransaction) GetBurnBlockTimeIsoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BurnBlockTimeIso, true
}

// SetBurnBlockTimeIso sets field value
func (o *TokenTransferTransaction) SetBurnBlockTimeIso(v string) {
	o.BurnBlockTimeIso = v
}

// GetParentBurnBlockTime returns the ParentBurnBlockTime field value
func (o *TokenTransferTransaction) GetParentBurnBlockTime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ParentBurnBlockTime
}

// GetParentBurnBlockTimeOk returns a tuple with the ParentBurnBlockTime field value
// and a boolean to check if the value has been set.
func (o *TokenTransferTransaction) GetParentBurnBlockTimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentBurnBlockTime, true
}

// SetParentBurnBlockTime sets field value
func (o *TokenTransferTransaction) SetParentBurnBlockTime(v int32) {
	o.ParentBurnBlockTime = v
}

// GetParentBurnBlockTimeIso returns the ParentBurnBlockTimeIso field value
func (o *TokenTransferTransaction) GetParentBurnBlockTimeIso() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParentBurnBlockTimeIso
}

// GetParentBurnBlockTimeIsoOk returns a tuple with the ParentBurnBlockTimeIso field value
// and a boolean to check if the value has been set.
func (o *TokenTransferTransaction) GetParentBurnBlockTimeIsoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentBurnBlockTimeIso, true
}

// SetParentBurnBlockTimeIso sets field value
func (o *TokenTransferTransaction) SetParentBurnBlockTimeIso(v string) {
	o.ParentBurnBlockTimeIso = v
}

// GetCanonical returns the Canonical field value
func (o *TokenTransferTransaction) GetCanonical() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Canonical
}

// GetCanonicalOk returns a tuple with the Canonical field value
// and a boolean to check if the value has been set.
func (o *TokenTransferTransaction) GetCanonicalOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Canonical, true
}

// SetCanonical sets field value
func (o *TokenTransferTransaction) SetCanonical(v bool) {
	o.Canonical = v
}

// GetTxIndex returns the TxIndex field value
func (o *TokenTransferTransaction) GetTxIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TxIndex
}

// GetTxIndexOk returns a tuple with the TxIndex field value
// and a boolean to check if the value has been set.
func (o *TokenTransferTransaction) GetTxIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxIndex, true
}

// SetTxIndex sets field value
func (o *TokenTransferTransaction) SetTxIndex(v int32) {
	o.TxIndex = v
}

// GetTxStatus returns the TxStatus field value
func (o *TokenTransferTransaction) GetTxStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TxStatus
}

// GetTxStatusOk returns a tuple with the TxStatus field value
// and a boolean to check if the value has been set.
func (o *TokenTransferTransaction) GetTxStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxStatus, true
}

// SetTxStatus sets field value
func (o *TokenTransferTransaction) SetTxStatus(v string) {
	o.TxStatus = v
}

// GetTxResult returns the TxResult field value
func (o *TokenTransferTransaction) GetTxResult() AbstractTransactionAllOfTxResult {
	if o == nil {
		var ret AbstractTransactionAllOfTxResult
		return ret
	}

	return o.TxResult
}

// GetTxResultOk returns a tuple with the TxResult field value
// and a boolean to check if the value has been set.
func (o *TokenTransferTransaction) GetTxResultOk() (*AbstractTransactionAllOfTxResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxResult, true
}

// SetTxResult sets field value
func (o *TokenTransferTransaction) SetTxResult(v AbstractTransactionAllOfTxResult) {
	o.TxResult = v
}

// GetEventCount returns the EventCount field value
func (o *TokenTransferTransaction) GetEventCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EventCount
}

// GetEventCountOk returns a tuple with the EventCount field value
// and a boolean to check if the value has been set.
func (o *TokenTransferTransaction) GetEventCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventCount, true
}

// SetEventCount sets field value
func (o *TokenTransferTransaction) SetEventCount(v int32) {
	o.EventCount = v
}

// GetParentBlockHash returns the ParentBlockHash field value
func (o *TokenTransferTransaction) GetParentBlockHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParentBlockHash
}

// GetParentBlockHashOk returns a tuple with the ParentBlockHash field value
// and a boolean to check if the value has been set.
func (o *TokenTransferTransaction) GetParentBlockHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentBlockHash, true
}

// SetParentBlockHash sets field value
func (o *TokenTransferTransaction) SetParentBlockHash(v string) {
	o.ParentBlockHash = v
}

// GetIsUnanchored returns the IsUnanchored field value
func (o *TokenTransferTransaction) GetIsUnanchored() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsUnanchored
}

// GetIsUnanchoredOk returns a tuple with the IsUnanchored field value
// and a boolean to check if the value has been set.
func (o *TokenTransferTransaction) GetIsUnanchoredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsUnanchored, true
}

// SetIsUnanchored sets field value
func (o *TokenTransferTransaction) SetIsUnanchored(v bool) {
	o.IsUnanchored = v
}

// GetMicroblockHash returns the MicroblockHash field value
func (o *TokenTransferTransaction) GetMicroblockHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MicroblockHash
}

// GetMicroblockHashOk returns a tuple with the MicroblockHash field value
// and a boolean to check if the value has been set.
func (o *TokenTransferTransaction) GetMicroblockHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MicroblockHash, true
}

// SetMicroblockHash sets field value
func (o *TokenTransferTransaction) SetMicroblockHash(v string) {
	o.MicroblockHash = v
}

// GetMicroblockSequence returns the MicroblockSequence field value
func (o *TokenTransferTransaction) GetMicroblockSequence() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MicroblockSequence
}

// GetMicroblockSequenceOk returns a tuple with the MicroblockSequence field value
// and a boolean to check if the value has been set.
func (o *TokenTransferTransaction) GetMicroblockSequenceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MicroblockSequence, true
}

// SetMicroblockSequence sets field value
func (o *TokenTransferTransaction) SetMicroblockSequence(v int32) {
	o.MicroblockSequence = v
}

// GetMicroblockCanonical returns the MicroblockCanonical field value
func (o *TokenTransferTransaction) GetMicroblockCanonical() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.MicroblockCanonical
}

// GetMicroblockCanonicalOk returns a tuple with the MicroblockCanonical field value
// and a boolean to check if the value has been set.
func (o *TokenTransferTransaction) GetMicroblockCanonicalOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MicroblockCanonical, true
}

// SetMicroblockCanonical sets field value
func (o *TokenTransferTransaction) SetMicroblockCanonical(v bool) {
	o.MicroblockCanonical = v
}

// GetExecutionCostReadCount returns the ExecutionCostReadCount field value
func (o *TokenTransferTransaction) GetExecutionCostReadCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExecutionCostReadCount
}

// GetExecutionCostReadCountOk returns a tuple with the ExecutionCostReadCount field value
// and a boolean to check if the value has been set.
func (o *TokenTransferTransaction) GetExecutionCostReadCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExecutionCostReadCount, true
}

// SetExecutionCostReadCount sets field value
func (o *TokenTransferTransaction) SetExecutionCostReadCount(v int32) {
	o.ExecutionCostReadCount = v
}

// GetExecutionCostReadLength returns the ExecutionCostReadLength field value
func (o *TokenTransferTransaction) GetExecutionCostReadLength() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExecutionCostReadLength
}

// GetExecutionCostReadLengthOk returns a tuple with the ExecutionCostReadLength field value
// and a boolean to check if the value has been set.
func (o *TokenTransferTransaction) GetExecutionCostReadLengthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExecutionCostReadLength, true
}

// SetExecutionCostReadLength sets field value
func (o *TokenTransferTransaction) SetExecutionCostReadLength(v int32) {
	o.ExecutionCostReadLength = v
}

// GetExecutionCostRuntime returns the ExecutionCostRuntime field value
func (o *TokenTransferTransaction) GetExecutionCostRuntime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExecutionCostRuntime
}

// GetExecutionCostRuntimeOk returns a tuple with the ExecutionCostRuntime field value
// and a boolean to check if the value has been set.
func (o *TokenTransferTransaction) GetExecutionCostRuntimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExecutionCostRuntime, true
}

// SetExecutionCostRuntime sets field value
func (o *TokenTransferTransaction) SetExecutionCostRuntime(v int32) {
	o.ExecutionCostRuntime = v
}

// GetExecutionCostWriteCount returns the ExecutionCostWriteCount field value
func (o *TokenTransferTransaction) GetExecutionCostWriteCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExecutionCostWriteCount
}

// GetExecutionCostWriteCountOk returns a tuple with the ExecutionCostWriteCount field value
// and a boolean to check if the value has been set.
func (o *TokenTransferTransaction) GetExecutionCostWriteCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExecutionCostWriteCount, true
}

// SetExecutionCostWriteCount sets field value
func (o *TokenTransferTransaction) SetExecutionCostWriteCount(v int32) {
	o.ExecutionCostWriteCount = v
}

// GetExecutionCostWriteLength returns the ExecutionCostWriteLength field value
func (o *TokenTransferTransaction) GetExecutionCostWriteLength() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExecutionCostWriteLength
}

// GetExecutionCostWriteLengthOk returns a tuple with the ExecutionCostWriteLength field value
// and a boolean to check if the value has been set.
func (o *TokenTransferTransaction) GetExecutionCostWriteLengthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExecutionCostWriteLength, true
}

// SetExecutionCostWriteLength sets field value
func (o *TokenTransferTransaction) SetExecutionCostWriteLength(v int32) {
	o.ExecutionCostWriteLength = v
}

// GetEvents returns the Events field value
func (o *TokenTransferTransaction) GetEvents() []TransactionEvent {
	if o == nil {
		var ret []TransactionEvent
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *TokenTransferTransaction) GetEventsOk() ([]TransactionEvent, bool) {
	if o == nil {
		return nil, false
	}
	return o.Events, true
}

// SetEvents sets field value
func (o *TokenTransferTransaction) SetEvents(v []TransactionEvent) {
	o.Events = v
}

// GetTxType returns the TxType field value
func (o *TokenTransferTransaction) GetTxType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TxType
}

// GetTxTypeOk returns a tuple with the TxType field value
// and a boolean to check if the value has been set.
func (o *TokenTransferTransaction) GetTxTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxType, true
}

// SetTxType sets field value
func (o *TokenTransferTransaction) SetTxType(v string) {
	o.TxType = v
}

// GetTokenTransfer returns the TokenTransfer field value
func (o *TokenTransferTransaction) GetTokenTransfer() TokenTransferTransactionMetadataTokenTransfer {
	if o == nil {
		var ret TokenTransferTransactionMetadataTokenTransfer
		return ret
	}

	return o.TokenTransfer
}

// GetTokenTransferOk returns a tuple with the TokenTransfer field value
// and a boolean to check if the value has been set.
func (o *TokenTransferTransaction) GetTokenTransferOk() (*TokenTransferTransactionMetadataTokenTransfer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenTransfer, true
}

// SetTokenTransfer sets field value
func (o *TokenTransferTransaction) SetTokenTransfer(v TokenTransferTransactionMetadataTokenTransfer) {
	o.TokenTransfer = v
}

func (o TokenTransferTransaction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenTransferTransaction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tx_id"] = o.TxId
	toSerialize["nonce"] = o.Nonce
	toSerialize["fee_rate"] = o.FeeRate
	toSerialize["sender_address"] = o.SenderAddress
	if !utils.IsNil(o.SponsorNonce) {
		toSerialize["sponsor_nonce"] = o.SponsorNonce
	}
	toSerialize["sponsored"] = o.Sponsored
	if !utils.IsNil(o.SponsorAddress) {
		toSerialize["sponsor_address"] = o.SponsorAddress
	}
	toSerialize["post_condition_mode"] = o.PostConditionMode
	toSerialize["post_conditions"] = o.PostConditions
	toSerialize["anchor_mode"] = o.AnchorMode
	toSerialize["block_hash"] = o.BlockHash
	toSerialize["block_height"] = o.BlockHeight
	toSerialize["block_time"] = o.BlockTime
	toSerialize["block_time_iso"] = o.BlockTimeIso
	toSerialize["burn_block_height"] = o.BurnBlockHeight
	toSerialize["burn_block_time"] = o.BurnBlockTime
	toSerialize["burn_block_time_iso"] = o.BurnBlockTimeIso
	toSerialize["parent_burn_block_time"] = o.ParentBurnBlockTime
	toSerialize["parent_burn_block_time_iso"] = o.ParentBurnBlockTimeIso
	toSerialize["canonical"] = o.Canonical
	toSerialize["tx_index"] = o.TxIndex
	toSerialize["tx_status"] = o.TxStatus
	toSerialize["tx_result"] = o.TxResult
	toSerialize["event_count"] = o.EventCount
	toSerialize["parent_block_hash"] = o.ParentBlockHash
	toSerialize["is_unanchored"] = o.IsUnanchored
	toSerialize["microblock_hash"] = o.MicroblockHash
	toSerialize["microblock_sequence"] = o.MicroblockSequence
	toSerialize["microblock_canonical"] = o.MicroblockCanonical
	toSerialize["execution_cost_read_count"] = o.ExecutionCostReadCount
	toSerialize["execution_cost_read_length"] = o.ExecutionCostReadLength
	toSerialize["execution_cost_runtime"] = o.ExecutionCostRuntime
	toSerialize["execution_cost_write_count"] = o.ExecutionCostWriteCount
	toSerialize["execution_cost_write_length"] = o.ExecutionCostWriteLength
	toSerialize["events"] = o.Events
	toSerialize["tx_type"] = o.TxType
	toSerialize["token_transfer"] = o.TokenTransfer
	return toSerialize, nil
}

func (o *TokenTransferTransaction) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tx_id",
		"nonce",
		"fee_rate",
		"sender_address",
		"sponsored",
		"post_condition_mode",
		"post_conditions",
		"anchor_mode",
		"block_hash",
		"block_height",
		"block_time",
		"block_time_iso",
		"burn_block_height",
		"burn_block_time",
		"burn_block_time_iso",
		"parent_burn_block_time",
		"parent_burn_block_time_iso",
		"canonical",
		"tx_index",
		"tx_status",
		"tx_result",
		"event_count",
		"parent_block_hash",
		"is_unanchored",
		"microblock_hash",
		"microblock_sequence",
		"microblock_canonical",
		"execution_cost_read_count",
		"execution_cost_read_length",
		"execution_cost_runtime",
		"execution_cost_write_count",
		"execution_cost_write_length",
		"events",
		"tx_type",
		"token_transfer",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varTokenTransferTransaction := _TokenTransferTransaction{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTokenTransferTransaction)

	if err != nil {
		return err
	}

	*o = TokenTransferTransaction(varTokenTransferTransaction)

	return err
}

type NullableTokenTransferTransaction struct {
	value *TokenTransferTransaction
	isSet bool
}

func (v NullableTokenTransferTransaction) Get() *TokenTransferTransaction {
	return v.value
}

func (v *NullableTokenTransferTransaction) Set(val *TokenTransferTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenTransferTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenTransferTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenTransferTransaction(val *TokenTransferTransaction) *NullableTokenTransferTransaction {
	return &NullableTokenTransferTransaction{value: val, isSet: true}
}

func (v NullableTokenTransferTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenTransferTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
