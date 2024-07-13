package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the SmartContractTransaction type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &SmartContractTransaction{}

// SmartContractTransaction Describes representation of a Type-1 Stacks 2.0 transaction. https://github.com/blockstack/stacks-blockchain/blob/master/sip/sip-005-blocks-and-transactions.md#type-1-instantiating-a-smart-contract
type SmartContractTransaction struct {
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
	SmartContract SmartContractTransactionMetadataSmartContract `json:"smart_contract"`
}

type _SmartContractTransaction SmartContractTransaction

// NewSmartContractTransaction instantiates a new SmartContractTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmartContractTransaction(txId string, nonce int32, feeRate string, senderAddress string, sponsored bool, postConditionMode string, postConditions []PostCondition, anchorMode string, blockHash string, blockHeight int32, blockTime float32, blockTimeIso string, burnBlockHeight int32, burnBlockTime int32, burnBlockTimeIso string, parentBurnBlockTime int32, parentBurnBlockTimeIso string, canonical bool, txIndex int32, txStatus string, txResult AbstractTransactionAllOfTxResult, eventCount int32, parentBlockHash string, isUnanchored bool, microblockHash string, microblockSequence int32, microblockCanonical bool, executionCostReadCount int32, executionCostReadLength int32, executionCostRuntime int32, executionCostWriteCount int32, executionCostWriteLength int32, events []TransactionEvent, txType string, smartContract SmartContractTransactionMetadataSmartContract) *SmartContractTransaction {
	this := SmartContractTransaction{}
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
	this.SmartContract = smartContract
	return &this
}

// NewSmartContractTransactionWithDefaults instantiates a new SmartContractTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmartContractTransactionWithDefaults() *SmartContractTransaction {
	this := SmartContractTransaction{}
	return &this
}

// GetTxId returns the TxId field value
func (o *SmartContractTransaction) GetTxId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TxId
}

// GetTxIdOk returns a tuple with the TxId field value
// and a boolean to check if the value has been set.
func (o *SmartContractTransaction) GetTxIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxId, true
}

// SetTxId sets field value
func (o *SmartContractTransaction) SetTxId(v string) {
	o.TxId = v
}

// GetNonce returns the Nonce field value
func (o *SmartContractTransaction) GetNonce() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value
// and a boolean to check if the value has been set.
func (o *SmartContractTransaction) GetNonceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nonce, true
}

// SetNonce sets field value
func (o *SmartContractTransaction) SetNonce(v int32) {
	o.Nonce = v
}

// GetFeeRate returns the FeeRate field value
func (o *SmartContractTransaction) GetFeeRate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FeeRate
}

// GetFeeRateOk returns a tuple with the FeeRate field value
// and a boolean to check if the value has been set.
func (o *SmartContractTransaction) GetFeeRateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeeRate, true
}

// SetFeeRate sets field value
func (o *SmartContractTransaction) SetFeeRate(v string) {
	o.FeeRate = v
}

// GetSenderAddress returns the SenderAddress field value
func (o *SmartContractTransaction) GetSenderAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SenderAddress
}

// GetSenderAddressOk returns a tuple with the SenderAddress field value
// and a boolean to check if the value has been set.
func (o *SmartContractTransaction) GetSenderAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SenderAddress, true
}

// SetSenderAddress sets field value
func (o *SmartContractTransaction) SetSenderAddress(v string) {
	o.SenderAddress = v
}

// GetSponsorNonce returns the SponsorNonce field value if set, zero value otherwise.
func (o *SmartContractTransaction) GetSponsorNonce() int32 {
	if o == nil || utils.IsNil(o.SponsorNonce) {
		var ret int32
		return ret
	}
	return *o.SponsorNonce
}

// GetSponsorNonceOk returns a tuple with the SponsorNonce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartContractTransaction) GetSponsorNonceOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.SponsorNonce) {
		return nil, false
	}
	return o.SponsorNonce, true
}

// HasSponsorNonce returns a boolean if a field has been set.
func (o *SmartContractTransaction) HasSponsorNonce() bool {
	if o != nil && !utils.IsNil(o.SponsorNonce) {
		return true
	}

	return false
}

// SetSponsorNonce gets a reference to the given int32 and assigns it to the SponsorNonce field.
func (o *SmartContractTransaction) SetSponsorNonce(v int32) {
	o.SponsorNonce = &v
}

// GetSponsored returns the Sponsored field value
func (o *SmartContractTransaction) GetSponsored() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Sponsored
}

// GetSponsoredOk returns a tuple with the Sponsored field value
// and a boolean to check if the value has been set.
func (o *SmartContractTransaction) GetSponsoredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sponsored, true
}

// SetSponsored sets field value
func (o *SmartContractTransaction) SetSponsored(v bool) {
	o.Sponsored = v
}

// GetSponsorAddress returns the SponsorAddress field value if set, zero value otherwise.
func (o *SmartContractTransaction) GetSponsorAddress() string {
	if o == nil || utils.IsNil(o.SponsorAddress) {
		var ret string
		return ret
	}
	return *o.SponsorAddress
}

// GetSponsorAddressOk returns a tuple with the SponsorAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartContractTransaction) GetSponsorAddressOk() (*string, bool) {
	if o == nil || utils.IsNil(o.SponsorAddress) {
		return nil, false
	}
	return o.SponsorAddress, true
}

// HasSponsorAddress returns a boolean if a field has been set.
func (o *SmartContractTransaction) HasSponsorAddress() bool {
	if o != nil && !utils.IsNil(o.SponsorAddress) {
		return true
	}

	return false
}

// SetSponsorAddress gets a reference to the given string and assigns it to the SponsorAddress field.
func (o *SmartContractTransaction) SetSponsorAddress(v string) {
	o.SponsorAddress = &v
}

// GetPostConditionMode returns the PostConditionMode field value
func (o *SmartContractTransaction) GetPostConditionMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PostConditionMode
}

// GetPostConditionModeOk returns a tuple with the PostConditionMode field value
// and a boolean to check if the value has been set.
func (o *SmartContractTransaction) GetPostConditionModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PostConditionMode, true
}

// SetPostConditionMode sets field value
func (o *SmartContractTransaction) SetPostConditionMode(v string) {
	o.PostConditionMode = v
}

// GetPostConditions returns the PostConditions field value
func (o *SmartContractTransaction) GetPostConditions() []PostCondition {
	if o == nil {
		var ret []PostCondition
		return ret
	}

	return o.PostConditions
}

// GetPostConditionsOk returns a tuple with the PostConditions field value
// and a boolean to check if the value has been set.
func (o *SmartContractTransaction) GetPostConditionsOk() ([]PostCondition, bool) {
	if o == nil {
		return nil, false
	}
	return o.PostConditions, true
}

// SetPostConditions sets field value
func (o *SmartContractTransaction) SetPostConditions(v []PostCondition) {
	o.PostConditions = v
}

// GetAnchorMode returns the AnchorMode field value
func (o *SmartContractTransaction) GetAnchorMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AnchorMode
}

// GetAnchorModeOk returns a tuple with the AnchorMode field value
// and a boolean to check if the value has been set.
func (o *SmartContractTransaction) GetAnchorModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AnchorMode, true
}

// SetAnchorMode sets field value
func (o *SmartContractTransaction) SetAnchorMode(v string) {
	o.AnchorMode = v
}

// GetBlockHash returns the BlockHash field value
func (o *SmartContractTransaction) GetBlockHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BlockHash
}

// GetBlockHashOk returns a tuple with the BlockHash field value
// and a boolean to check if the value has been set.
func (o *SmartContractTransaction) GetBlockHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockHash, true
}

// SetBlockHash sets field value
func (o *SmartContractTransaction) SetBlockHash(v string) {
	o.BlockHash = v
}

// GetBlockHeight returns the BlockHeight field value
func (o *SmartContractTransaction) GetBlockHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BlockHeight
}

// GetBlockHeightOk returns a tuple with the BlockHeight field value
// and a boolean to check if the value has been set.
func (o *SmartContractTransaction) GetBlockHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockHeight, true
}

// SetBlockHeight sets field value
func (o *SmartContractTransaction) SetBlockHeight(v int32) {
	o.BlockHeight = v
}

// GetBlockTime returns the BlockTime field value
func (o *SmartContractTransaction) GetBlockTime() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.BlockTime
}

// GetBlockTimeOk returns a tuple with the BlockTime field value
// and a boolean to check if the value has been set.
func (o *SmartContractTransaction) GetBlockTimeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockTime, true
}

// SetBlockTime sets field value
func (o *SmartContractTransaction) SetBlockTime(v float32) {
	o.BlockTime = v
}

// GetBlockTimeIso returns the BlockTimeIso field value
func (o *SmartContractTransaction) GetBlockTimeIso() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BlockTimeIso
}

// GetBlockTimeIsoOk returns a tuple with the BlockTimeIso field value
// and a boolean to check if the value has been set.
func (o *SmartContractTransaction) GetBlockTimeIsoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockTimeIso, true
}

// SetBlockTimeIso sets field value
func (o *SmartContractTransaction) SetBlockTimeIso(v string) {
	o.BlockTimeIso = v
}

// GetBurnBlockHeight returns the BurnBlockHeight field value
func (o *SmartContractTransaction) GetBurnBlockHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BurnBlockHeight
}

// GetBurnBlockHeightOk returns a tuple with the BurnBlockHeight field value
// and a boolean to check if the value has been set.
func (o *SmartContractTransaction) GetBurnBlockHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BurnBlockHeight, true
}

// SetBurnBlockHeight sets field value
func (o *SmartContractTransaction) SetBurnBlockHeight(v int32) {
	o.BurnBlockHeight = v
}

// GetBurnBlockTime returns the BurnBlockTime field value
func (o *SmartContractTransaction) GetBurnBlockTime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BurnBlockTime
}

// GetBurnBlockTimeOk returns a tuple with the BurnBlockTime field value
// and a boolean to check if the value has been set.
func (o *SmartContractTransaction) GetBurnBlockTimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BurnBlockTime, true
}

// SetBurnBlockTime sets field value
func (o *SmartContractTransaction) SetBurnBlockTime(v int32) {
	o.BurnBlockTime = v
}

// GetBurnBlockTimeIso returns the BurnBlockTimeIso field value
func (o *SmartContractTransaction) GetBurnBlockTimeIso() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BurnBlockTimeIso
}

// GetBurnBlockTimeIsoOk returns a tuple with the BurnBlockTimeIso field value
// and a boolean to check if the value has been set.
func (o *SmartContractTransaction) GetBurnBlockTimeIsoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BurnBlockTimeIso, true
}

// SetBurnBlockTimeIso sets field value
func (o *SmartContractTransaction) SetBurnBlockTimeIso(v string) {
	o.BurnBlockTimeIso = v
}

// GetParentBurnBlockTime returns the ParentBurnBlockTime field value
func (o *SmartContractTransaction) GetParentBurnBlockTime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ParentBurnBlockTime
}

// GetParentBurnBlockTimeOk returns a tuple with the ParentBurnBlockTime field value
// and a boolean to check if the value has been set.
func (o *SmartContractTransaction) GetParentBurnBlockTimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentBurnBlockTime, true
}

// SetParentBurnBlockTime sets field value
func (o *SmartContractTransaction) SetParentBurnBlockTime(v int32) {
	o.ParentBurnBlockTime = v
}

// GetParentBurnBlockTimeIso returns the ParentBurnBlockTimeIso field value
func (o *SmartContractTransaction) GetParentBurnBlockTimeIso() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParentBurnBlockTimeIso
}

// GetParentBurnBlockTimeIsoOk returns a tuple with the ParentBurnBlockTimeIso field value
// and a boolean to check if the value has been set.
func (o *SmartContractTransaction) GetParentBurnBlockTimeIsoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentBurnBlockTimeIso, true
}

// SetParentBurnBlockTimeIso sets field value
func (o *SmartContractTransaction) SetParentBurnBlockTimeIso(v string) {
	o.ParentBurnBlockTimeIso = v
}

// GetCanonical returns the Canonical field value
func (o *SmartContractTransaction) GetCanonical() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Canonical
}

// GetCanonicalOk returns a tuple with the Canonical field value
// and a boolean to check if the value has been set.
func (o *SmartContractTransaction) GetCanonicalOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Canonical, true
}

// SetCanonical sets field value
func (o *SmartContractTransaction) SetCanonical(v bool) {
	o.Canonical = v
}

// GetTxIndex returns the TxIndex field value
func (o *SmartContractTransaction) GetTxIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TxIndex
}

// GetTxIndexOk returns a tuple with the TxIndex field value
// and a boolean to check if the value has been set.
func (o *SmartContractTransaction) GetTxIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxIndex, true
}

// SetTxIndex sets field value
func (o *SmartContractTransaction) SetTxIndex(v int32) {
	o.TxIndex = v
}

// GetTxStatus returns the TxStatus field value
func (o *SmartContractTransaction) GetTxStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TxStatus
}

// GetTxStatusOk returns a tuple with the TxStatus field value
// and a boolean to check if the value has been set.
func (o *SmartContractTransaction) GetTxStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxStatus, true
}

// SetTxStatus sets field value
func (o *SmartContractTransaction) SetTxStatus(v string) {
	o.TxStatus = v
}

// GetTxResult returns the TxResult field value
func (o *SmartContractTransaction) GetTxResult() AbstractTransactionAllOfTxResult {
	if o == nil {
		var ret AbstractTransactionAllOfTxResult
		return ret
	}

	return o.TxResult
}

// GetTxResultOk returns a tuple with the TxResult field value
// and a boolean to check if the value has been set.
func (o *SmartContractTransaction) GetTxResultOk() (*AbstractTransactionAllOfTxResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxResult, true
}

// SetTxResult sets field value
func (o *SmartContractTransaction) SetTxResult(v AbstractTransactionAllOfTxResult) {
	o.TxResult = v
}

// GetEventCount returns the EventCount field value
func (o *SmartContractTransaction) GetEventCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EventCount
}

// GetEventCountOk returns a tuple with the EventCount field value
// and a boolean to check if the value has been set.
func (o *SmartContractTransaction) GetEventCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventCount, true
}

// SetEventCount sets field value
func (o *SmartContractTransaction) SetEventCount(v int32) {
	o.EventCount = v
}

// GetParentBlockHash returns the ParentBlockHash field value
func (o *SmartContractTransaction) GetParentBlockHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParentBlockHash
}

// GetParentBlockHashOk returns a tuple with the ParentBlockHash field value
// and a boolean to check if the value has been set.
func (o *SmartContractTransaction) GetParentBlockHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentBlockHash, true
}

// SetParentBlockHash sets field value
func (o *SmartContractTransaction) SetParentBlockHash(v string) {
	o.ParentBlockHash = v
}

// GetIsUnanchored returns the IsUnanchored field value
func (o *SmartContractTransaction) GetIsUnanchored() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsUnanchored
}

// GetIsUnanchoredOk returns a tuple with the IsUnanchored field value
// and a boolean to check if the value has been set.
func (o *SmartContractTransaction) GetIsUnanchoredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsUnanchored, true
}

// SetIsUnanchored sets field value
func (o *SmartContractTransaction) SetIsUnanchored(v bool) {
	o.IsUnanchored = v
}

// GetMicroblockHash returns the MicroblockHash field value
func (o *SmartContractTransaction) GetMicroblockHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MicroblockHash
}

// GetMicroblockHashOk returns a tuple with the MicroblockHash field value
// and a boolean to check if the value has been set.
func (o *SmartContractTransaction) GetMicroblockHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MicroblockHash, true
}

// SetMicroblockHash sets field value
func (o *SmartContractTransaction) SetMicroblockHash(v string) {
	o.MicroblockHash = v
}

// GetMicroblockSequence returns the MicroblockSequence field value
func (o *SmartContractTransaction) GetMicroblockSequence() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MicroblockSequence
}

// GetMicroblockSequenceOk returns a tuple with the MicroblockSequence field value
// and a boolean to check if the value has been set.
func (o *SmartContractTransaction) GetMicroblockSequenceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MicroblockSequence, true
}

// SetMicroblockSequence sets field value
func (o *SmartContractTransaction) SetMicroblockSequence(v int32) {
	o.MicroblockSequence = v
}

// GetMicroblockCanonical returns the MicroblockCanonical field value
func (o *SmartContractTransaction) GetMicroblockCanonical() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.MicroblockCanonical
}

// GetMicroblockCanonicalOk returns a tuple with the MicroblockCanonical field value
// and a boolean to check if the value has been set.
func (o *SmartContractTransaction) GetMicroblockCanonicalOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MicroblockCanonical, true
}

// SetMicroblockCanonical sets field value
func (o *SmartContractTransaction) SetMicroblockCanonical(v bool) {
	o.MicroblockCanonical = v
}

// GetExecutionCostReadCount returns the ExecutionCostReadCount field value
func (o *SmartContractTransaction) GetExecutionCostReadCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExecutionCostReadCount
}

// GetExecutionCostReadCountOk returns a tuple with the ExecutionCostReadCount field value
// and a boolean to check if the value has been set.
func (o *SmartContractTransaction) GetExecutionCostReadCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExecutionCostReadCount, true
}

// SetExecutionCostReadCount sets field value
func (o *SmartContractTransaction) SetExecutionCostReadCount(v int32) {
	o.ExecutionCostReadCount = v
}

// GetExecutionCostReadLength returns the ExecutionCostReadLength field value
func (o *SmartContractTransaction) GetExecutionCostReadLength() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExecutionCostReadLength
}

// GetExecutionCostReadLengthOk returns a tuple with the ExecutionCostReadLength field value
// and a boolean to check if the value has been set.
func (o *SmartContractTransaction) GetExecutionCostReadLengthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExecutionCostReadLength, true
}

// SetExecutionCostReadLength sets field value
func (o *SmartContractTransaction) SetExecutionCostReadLength(v int32) {
	o.ExecutionCostReadLength = v
}

// GetExecutionCostRuntime returns the ExecutionCostRuntime field value
func (o *SmartContractTransaction) GetExecutionCostRuntime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExecutionCostRuntime
}

// GetExecutionCostRuntimeOk returns a tuple with the ExecutionCostRuntime field value
// and a boolean to check if the value has been set.
func (o *SmartContractTransaction) GetExecutionCostRuntimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExecutionCostRuntime, true
}

// SetExecutionCostRuntime sets field value
func (o *SmartContractTransaction) SetExecutionCostRuntime(v int32) {
	o.ExecutionCostRuntime = v
}

// GetExecutionCostWriteCount returns the ExecutionCostWriteCount field value
func (o *SmartContractTransaction) GetExecutionCostWriteCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExecutionCostWriteCount
}

// GetExecutionCostWriteCountOk returns a tuple with the ExecutionCostWriteCount field value
// and a boolean to check if the value has been set.
func (o *SmartContractTransaction) GetExecutionCostWriteCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExecutionCostWriteCount, true
}

// SetExecutionCostWriteCount sets field value
func (o *SmartContractTransaction) SetExecutionCostWriteCount(v int32) {
	o.ExecutionCostWriteCount = v
}

// GetExecutionCostWriteLength returns the ExecutionCostWriteLength field value
func (o *SmartContractTransaction) GetExecutionCostWriteLength() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExecutionCostWriteLength
}

// GetExecutionCostWriteLengthOk returns a tuple with the ExecutionCostWriteLength field value
// and a boolean to check if the value has been set.
func (o *SmartContractTransaction) GetExecutionCostWriteLengthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExecutionCostWriteLength, true
}

// SetExecutionCostWriteLength sets field value
func (o *SmartContractTransaction) SetExecutionCostWriteLength(v int32) {
	o.ExecutionCostWriteLength = v
}

// GetEvents returns the Events field value
func (o *SmartContractTransaction) GetEvents() []TransactionEvent {
	if o == nil {
		var ret []TransactionEvent
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *SmartContractTransaction) GetEventsOk() ([]TransactionEvent, bool) {
	if o == nil {
		return nil, false
	}
	return o.Events, true
}

// SetEvents sets field value
func (o *SmartContractTransaction) SetEvents(v []TransactionEvent) {
	o.Events = v
}

// GetTxType returns the TxType field value
func (o *SmartContractTransaction) GetTxType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TxType
}

// GetTxTypeOk returns a tuple with the TxType field value
// and a boolean to check if the value has been set.
func (o *SmartContractTransaction) GetTxTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxType, true
}

// SetTxType sets field value
func (o *SmartContractTransaction) SetTxType(v string) {
	o.TxType = v
}

// GetSmartContract returns the SmartContract field value
func (o *SmartContractTransaction) GetSmartContract() SmartContractTransactionMetadataSmartContract {
	if o == nil {
		var ret SmartContractTransactionMetadataSmartContract
		return ret
	}

	return o.SmartContract
}

// GetSmartContractOk returns a tuple with the SmartContract field value
// and a boolean to check if the value has been set.
func (o *SmartContractTransaction) GetSmartContractOk() (*SmartContractTransactionMetadataSmartContract, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SmartContract, true
}

// SetSmartContract sets field value
func (o *SmartContractTransaction) SetSmartContract(v SmartContractTransactionMetadataSmartContract) {
	o.SmartContract = v
}

func (o SmartContractTransaction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SmartContractTransaction) ToMap() (map[string]interface{}, error) {
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
	toSerialize["smart_contract"] = o.SmartContract
	return toSerialize, nil
}

func (o *SmartContractTransaction) UnmarshalJSON(data []byte) (err error) {
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
		"smart_contract",
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

	varSmartContractTransaction := _SmartContractTransaction{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSmartContractTransaction)

	if err != nil {
		return err
	}

	*o = SmartContractTransaction(varSmartContractTransaction)

	return err
}

type NullableSmartContractTransaction struct {
	value *SmartContractTransaction
	isSet bool
}

func (v NullableSmartContractTransaction) Get() *SmartContractTransaction {
	return v.value
}

func (v *NullableSmartContractTransaction) Set(val *SmartContractTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableSmartContractTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableSmartContractTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmartContractTransaction(val *SmartContractTransaction) *NullableSmartContractTransaction {
	return &NullableSmartContractTransaction{value: val, isSet: true}
}

func (v NullableSmartContractTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmartContractTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
