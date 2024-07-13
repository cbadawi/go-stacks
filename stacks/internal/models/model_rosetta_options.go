package models

import (
	"encoding/json"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the RosettaOptions type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RosettaOptions{}

// RosettaOptions The options that will be sent directly to /construction/metadata by the caller.
type RosettaOptions struct {
	// sender's address
	SenderAddress *string `json:"sender_address,omitempty"`
	// Type of operation e.g transfer
	Type *string `json:"type,omitempty"`
	// This value indicates the state of the operations
	Status *string `json:"status,omitempty"`
	// Recipient's address
	TokenTransferRecipientAddress *string `json:"token_transfer_recipient_address,omitempty"`
	// Amount to be transfered.
	Amount *string `json:"amount,omitempty"`
	// Currency symbol e.g STX
	Symbol *string `json:"symbol,omitempty"`
	// Number of decimal places
	Decimals *int32 `json:"decimals,omitempty"`
	// Maximum price a user is willing to pay.
	GasLimit *float32 `json:"gas_limit,omitempty"`
	// Cost necessary to perform a transaction on the network
	GasPrice *float32 `json:"gas_price,omitempty"`
	//  A suggested fee multiplier to indicate that the suggested fee should be scaled. This may be used to set higher fees for urgent transactions or to pay lower fees when there is less urgency.
	SuggestedFeeMultiplier *float32 `json:"suggested_fee_multiplier,omitempty"`
	// Maximum fee user is willing to pay
	MaxFee *string `json:"max_fee,omitempty"`
	// Fee for this transaction
	Fee *string `json:"fee,omitempty"`
	// Transaction approximative size (used to calculate total fee).
	Size *int32 `json:"size,omitempty"`
	// STX token transfer memo.
	Memo *string `json:"memo,omitempty"`
	// Number of cycles when stacking.
	NumberOfCycles *int32 `json:"number_of_cycles,omitempty"`
	// Address of the contract to call.
	ContractAddress *string `json:"contract_address,omitempty"`
	// Name of the contract to call.
	ContractName *string `json:"contract_name,omitempty"`
	// Set the burnchain (BTC) block for stacking lock to start.
	BurnBlockHeight *int32 `json:"burn_block_height,omitempty"`
	// The reward cycle ID for stacking transaction.
	RewardCycleId *int32 `json:"reward_cycle_id,omitempty"`
	// Delegator address for when calling `delegate-stacking`.
	DelegateTo *string `json:"delegate_to,omitempty"`
	// The reward address for stacking transaction. It should be a valid Bitcoin address
	PoxAddr *string `json:"pox_addr,omitempty"`
	// The hex-encoded signer key (buff 33) for PoX.
	SignerKey *string `json:"signer_key,omitempty"`
	// The hex-encoded signer private key for PoX. Specify either this or `signer_signature`, otherwise the PoX transaction requires allow-listing from the signer.
	SignerPrivateKey *string `json:"signer_private_key,omitempty"`
	// The hex-encoded signer signature for PoX. Specify either this or `signer_private_key`, otherwise the PoX transaction requires allow-listing from the signer.
	SignerSignature *string `json:"signer_signature,omitempty"`
	// The maximum amount of STX to stack for PoX. If not specified, the `amount` will be used as the `max-amount` for the PoX transaction.
	PoxMaxAmount *string `json:"pox_max_amount,omitempty"`
	// The auth ID for the PoX transaction. If not specified, a random value will be generated.
	PoxAuthId *string `json:"pox_auth_id,omitempty"`
}

// NewRosettaOptions instantiates a new RosettaOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRosettaOptions() *RosettaOptions {
	this := RosettaOptions{}
	return &this
}

// NewRosettaOptionsWithDefaults instantiates a new RosettaOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRosettaOptionsWithDefaults() *RosettaOptions {
	this := RosettaOptions{}
	return &this
}

// GetSenderAddress returns the SenderAddress field value if set, zero value otherwise.
func (o *RosettaOptions) GetSenderAddress() string {
	if o == nil || utils.IsNil(o.SenderAddress) {
		var ret string
		return ret
	}
	return *o.SenderAddress
}

// GetSenderAddressOk returns a tuple with the SenderAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaOptions) GetSenderAddressOk() (*string, bool) {
	if o == nil || utils.IsNil(o.SenderAddress) {
		return nil, false
	}
	return o.SenderAddress, true
}

// HasSenderAddress returns a boolean if a field has been set.
func (o *RosettaOptions) HasSenderAddress() bool {
	if o != nil && !utils.IsNil(o.SenderAddress) {
		return true
	}

	return false
}

// SetSenderAddress gets a reference to the given string and assigns it to the SenderAddress field.
func (o *RosettaOptions) SetSenderAddress(v string) {
	o.SenderAddress = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RosettaOptions) GetType() string {
	if o == nil || utils.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaOptions) GetTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RosettaOptions) HasType() bool {
	if o != nil && !utils.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RosettaOptions) SetType(v string) {
	o.Type = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *RosettaOptions) GetStatus() string {
	if o == nil || utils.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaOptions) GetStatusOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *RosettaOptions) HasStatus() bool {
	if o != nil && !utils.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *RosettaOptions) SetStatus(v string) {
	o.Status = &v
}

// GetTokenTransferRecipientAddress returns the TokenTransferRecipientAddress field value if set, zero value otherwise.
func (o *RosettaOptions) GetTokenTransferRecipientAddress() string {
	if o == nil || utils.IsNil(o.TokenTransferRecipientAddress) {
		var ret string
		return ret
	}
	return *o.TokenTransferRecipientAddress
}

// GetTokenTransferRecipientAddressOk returns a tuple with the TokenTransferRecipientAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaOptions) GetTokenTransferRecipientAddressOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TokenTransferRecipientAddress) {
		return nil, false
	}
	return o.TokenTransferRecipientAddress, true
}

// HasTokenTransferRecipientAddress returns a boolean if a field has been set.
func (o *RosettaOptions) HasTokenTransferRecipientAddress() bool {
	if o != nil && !utils.IsNil(o.TokenTransferRecipientAddress) {
		return true
	}

	return false
}

// SetTokenTransferRecipientAddress gets a reference to the given string and assigns it to the TokenTransferRecipientAddress field.
func (o *RosettaOptions) SetTokenTransferRecipientAddress(v string) {
	o.TokenTransferRecipientAddress = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *RosettaOptions) GetAmount() string {
	if o == nil || utils.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaOptions) GetAmountOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *RosettaOptions) HasAmount() bool {
	if o != nil && !utils.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *RosettaOptions) SetAmount(v string) {
	o.Amount = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *RosettaOptions) GetSymbol() string {
	if o == nil || utils.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaOptions) GetSymbolOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *RosettaOptions) HasSymbol() bool {
	if o != nil && !utils.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *RosettaOptions) SetSymbol(v string) {
	o.Symbol = &v
}

// GetDecimals returns the Decimals field value if set, zero value otherwise.
func (o *RosettaOptions) GetDecimals() int32 {
	if o == nil || utils.IsNil(o.Decimals) {
		var ret int32
		return ret
	}
	return *o.Decimals
}

// GetDecimalsOk returns a tuple with the Decimals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaOptions) GetDecimalsOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Decimals) {
		return nil, false
	}
	return o.Decimals, true
}

// HasDecimals returns a boolean if a field has been set.
func (o *RosettaOptions) HasDecimals() bool {
	if o != nil && !utils.IsNil(o.Decimals) {
		return true
	}

	return false
}

// SetDecimals gets a reference to the given int32 and assigns it to the Decimals field.
func (o *RosettaOptions) SetDecimals(v int32) {
	o.Decimals = &v
}

// GetGasLimit returns the GasLimit field value if set, zero value otherwise.
func (o *RosettaOptions) GetGasLimit() float32 {
	if o == nil || utils.IsNil(o.GasLimit) {
		var ret float32
		return ret
	}
	return *o.GasLimit
}

// GetGasLimitOk returns a tuple with the GasLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaOptions) GetGasLimitOk() (*float32, bool) {
	if o == nil || utils.IsNil(o.GasLimit) {
		return nil, false
	}
	return o.GasLimit, true
}

// HasGasLimit returns a boolean if a field has been set.
func (o *RosettaOptions) HasGasLimit() bool {
	if o != nil && !utils.IsNil(o.GasLimit) {
		return true
	}

	return false
}

// SetGasLimit gets a reference to the given float32 and assigns it to the GasLimit field.
func (o *RosettaOptions) SetGasLimit(v float32) {
	o.GasLimit = &v
}

// GetGasPrice returns the GasPrice field value if set, zero value otherwise.
func (o *RosettaOptions) GetGasPrice() float32 {
	if o == nil || utils.IsNil(o.GasPrice) {
		var ret float32
		return ret
	}
	return *o.GasPrice
}

// GetGasPriceOk returns a tuple with the GasPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaOptions) GetGasPriceOk() (*float32, bool) {
	if o == nil || utils.IsNil(o.GasPrice) {
		return nil, false
	}
	return o.GasPrice, true
}

// HasGasPrice returns a boolean if a field has been set.
func (o *RosettaOptions) HasGasPrice() bool {
	if o != nil && !utils.IsNil(o.GasPrice) {
		return true
	}

	return false
}

// SetGasPrice gets a reference to the given float32 and assigns it to the GasPrice field.
func (o *RosettaOptions) SetGasPrice(v float32) {
	o.GasPrice = &v
}

// GetSuggestedFeeMultiplier returns the SuggestedFeeMultiplier field value if set, zero value otherwise.
func (o *RosettaOptions) GetSuggestedFeeMultiplier() float32 {
	if o == nil || utils.IsNil(o.SuggestedFeeMultiplier) {
		var ret float32
		return ret
	}
	return *o.SuggestedFeeMultiplier
}

// GetSuggestedFeeMultiplierOk returns a tuple with the SuggestedFeeMultiplier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaOptions) GetSuggestedFeeMultiplierOk() (*float32, bool) {
	if o == nil || utils.IsNil(o.SuggestedFeeMultiplier) {
		return nil, false
	}
	return o.SuggestedFeeMultiplier, true
}

// HasSuggestedFeeMultiplier returns a boolean if a field has been set.
func (o *RosettaOptions) HasSuggestedFeeMultiplier() bool {
	if o != nil && !utils.IsNil(o.SuggestedFeeMultiplier) {
		return true
	}

	return false
}

// SetSuggestedFeeMultiplier gets a reference to the given float32 and assigns it to the SuggestedFeeMultiplier field.
func (o *RosettaOptions) SetSuggestedFeeMultiplier(v float32) {
	o.SuggestedFeeMultiplier = &v
}

// GetMaxFee returns the MaxFee field value if set, zero value otherwise.
func (o *RosettaOptions) GetMaxFee() string {
	if o == nil || utils.IsNil(o.MaxFee) {
		var ret string
		return ret
	}
	return *o.MaxFee
}

// GetMaxFeeOk returns a tuple with the MaxFee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaOptions) GetMaxFeeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.MaxFee) {
		return nil, false
	}
	return o.MaxFee, true
}

// HasMaxFee returns a boolean if a field has been set.
func (o *RosettaOptions) HasMaxFee() bool {
	if o != nil && !utils.IsNil(o.MaxFee) {
		return true
	}

	return false
}

// SetMaxFee gets a reference to the given string and assigns it to the MaxFee field.
func (o *RosettaOptions) SetMaxFee(v string) {
	o.MaxFee = &v
}

// GetFee returns the Fee field value if set, zero value otherwise.
func (o *RosettaOptions) GetFee() string {
	if o == nil || utils.IsNil(o.Fee) {
		var ret string
		return ret
	}
	return *o.Fee
}

// GetFeeOk returns a tuple with the Fee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaOptions) GetFeeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Fee) {
		return nil, false
	}
	return o.Fee, true
}

// HasFee returns a boolean if a field has been set.
func (o *RosettaOptions) HasFee() bool {
	if o != nil && !utils.IsNil(o.Fee) {
		return true
	}

	return false
}

// SetFee gets a reference to the given string and assigns it to the Fee field.
func (o *RosettaOptions) SetFee(v string) {
	o.Fee = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *RosettaOptions) GetSize() int32 {
	if o == nil || utils.IsNil(o.Size) {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaOptions) GetSizeOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *RosettaOptions) HasSize() bool {
	if o != nil && !utils.IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *RosettaOptions) SetSize(v int32) {
	o.Size = &v
}

// GetMemo returns the Memo field value if set, zero value otherwise.
func (o *RosettaOptions) GetMemo() string {
	if o == nil || utils.IsNil(o.Memo) {
		var ret string
		return ret
	}
	return *o.Memo
}

// GetMemoOk returns a tuple with the Memo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaOptions) GetMemoOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Memo) {
		return nil, false
	}
	return o.Memo, true
}

// HasMemo returns a boolean if a field has been set.
func (o *RosettaOptions) HasMemo() bool {
	if o != nil && !utils.IsNil(o.Memo) {
		return true
	}

	return false
}

// SetMemo gets a reference to the given string and assigns it to the Memo field.
func (o *RosettaOptions) SetMemo(v string) {
	o.Memo = &v
}

// GetNumberOfCycles returns the NumberOfCycles field value if set, zero value otherwise.
func (o *RosettaOptions) GetNumberOfCycles() int32 {
	if o == nil || utils.IsNil(o.NumberOfCycles) {
		var ret int32
		return ret
	}
	return *o.NumberOfCycles
}

// GetNumberOfCyclesOk returns a tuple with the NumberOfCycles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaOptions) GetNumberOfCyclesOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.NumberOfCycles) {
		return nil, false
	}
	return o.NumberOfCycles, true
}

// HasNumberOfCycles returns a boolean if a field has been set.
func (o *RosettaOptions) HasNumberOfCycles() bool {
	if o != nil && !utils.IsNil(o.NumberOfCycles) {
		return true
	}

	return false
}

// SetNumberOfCycles gets a reference to the given int32 and assigns it to the NumberOfCycles field.
func (o *RosettaOptions) SetNumberOfCycles(v int32) {
	o.NumberOfCycles = &v
}

// GetContractAddress returns the ContractAddress field value if set, zero value otherwise.
func (o *RosettaOptions) GetContractAddress() string {
	if o == nil || utils.IsNil(o.ContractAddress) {
		var ret string
		return ret
	}
	return *o.ContractAddress
}

// GetContractAddressOk returns a tuple with the ContractAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaOptions) GetContractAddressOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ContractAddress) {
		return nil, false
	}
	return o.ContractAddress, true
}

// HasContractAddress returns a boolean if a field has been set.
func (o *RosettaOptions) HasContractAddress() bool {
	if o != nil && !utils.IsNil(o.ContractAddress) {
		return true
	}

	return false
}

// SetContractAddress gets a reference to the given string and assigns it to the ContractAddress field.
func (o *RosettaOptions) SetContractAddress(v string) {
	o.ContractAddress = &v
}

// GetContractName returns the ContractName field value if set, zero value otherwise.
func (o *RosettaOptions) GetContractName() string {
	if o == nil || utils.IsNil(o.ContractName) {
		var ret string
		return ret
	}
	return *o.ContractName
}

// GetContractNameOk returns a tuple with the ContractName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaOptions) GetContractNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ContractName) {
		return nil, false
	}
	return o.ContractName, true
}

// HasContractName returns a boolean if a field has been set.
func (o *RosettaOptions) HasContractName() bool {
	if o != nil && !utils.IsNil(o.ContractName) {
		return true
	}

	return false
}

// SetContractName gets a reference to the given string and assigns it to the ContractName field.
func (o *RosettaOptions) SetContractName(v string) {
	o.ContractName = &v
}

// GetBurnBlockHeight returns the BurnBlockHeight field value if set, zero value otherwise.
func (o *RosettaOptions) GetBurnBlockHeight() int32 {
	if o == nil || utils.IsNil(o.BurnBlockHeight) {
		var ret int32
		return ret
	}
	return *o.BurnBlockHeight
}

// GetBurnBlockHeightOk returns a tuple with the BurnBlockHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaOptions) GetBurnBlockHeightOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.BurnBlockHeight) {
		return nil, false
	}
	return o.BurnBlockHeight, true
}

// HasBurnBlockHeight returns a boolean if a field has been set.
func (o *RosettaOptions) HasBurnBlockHeight() bool {
	if o != nil && !utils.IsNil(o.BurnBlockHeight) {
		return true
	}

	return false
}

// SetBurnBlockHeight gets a reference to the given int32 and assigns it to the BurnBlockHeight field.
func (o *RosettaOptions) SetBurnBlockHeight(v int32) {
	o.BurnBlockHeight = &v
}

// GetRewardCycleId returns the RewardCycleId field value if set, zero value otherwise.
func (o *RosettaOptions) GetRewardCycleId() int32 {
	if o == nil || utils.IsNil(o.RewardCycleId) {
		var ret int32
		return ret
	}
	return *o.RewardCycleId
}

// GetRewardCycleIdOk returns a tuple with the RewardCycleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaOptions) GetRewardCycleIdOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.RewardCycleId) {
		return nil, false
	}
	return o.RewardCycleId, true
}

// HasRewardCycleId returns a boolean if a field has been set.
func (o *RosettaOptions) HasRewardCycleId() bool {
	if o != nil && !utils.IsNil(o.RewardCycleId) {
		return true
	}

	return false
}

// SetRewardCycleId gets a reference to the given int32 and assigns it to the RewardCycleId field.
func (o *RosettaOptions) SetRewardCycleId(v int32) {
	o.RewardCycleId = &v
}

// GetDelegateTo returns the DelegateTo field value if set, zero value otherwise.
func (o *RosettaOptions) GetDelegateTo() string {
	if o == nil || utils.IsNil(o.DelegateTo) {
		var ret string
		return ret
	}
	return *o.DelegateTo
}

// GetDelegateToOk returns a tuple with the DelegateTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaOptions) GetDelegateToOk() (*string, bool) {
	if o == nil || utils.IsNil(o.DelegateTo) {
		return nil, false
	}
	return o.DelegateTo, true
}

// HasDelegateTo returns a boolean if a field has been set.
func (o *RosettaOptions) HasDelegateTo() bool {
	if o != nil && !utils.IsNil(o.DelegateTo) {
		return true
	}

	return false
}

// SetDelegateTo gets a reference to the given string and assigns it to the DelegateTo field.
func (o *RosettaOptions) SetDelegateTo(v string) {
	o.DelegateTo = &v
}

// GetPoxAddr returns the PoxAddr field value if set, zero value otherwise.
func (o *RosettaOptions) GetPoxAddr() string {
	if o == nil || utils.IsNil(o.PoxAddr) {
		var ret string
		return ret
	}
	return *o.PoxAddr
}

// GetPoxAddrOk returns a tuple with the PoxAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaOptions) GetPoxAddrOk() (*string, bool) {
	if o == nil || utils.IsNil(o.PoxAddr) {
		return nil, false
	}
	return o.PoxAddr, true
}

// HasPoxAddr returns a boolean if a field has been set.
func (o *RosettaOptions) HasPoxAddr() bool {
	if o != nil && !utils.IsNil(o.PoxAddr) {
		return true
	}

	return false
}

// SetPoxAddr gets a reference to the given string and assigns it to the PoxAddr field.
func (o *RosettaOptions) SetPoxAddr(v string) {
	o.PoxAddr = &v
}

// GetSignerKey returns the SignerKey field value if set, zero value otherwise.
func (o *RosettaOptions) GetSignerKey() string {
	if o == nil || utils.IsNil(o.SignerKey) {
		var ret string
		return ret
	}
	return *o.SignerKey
}

// GetSignerKeyOk returns a tuple with the SignerKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaOptions) GetSignerKeyOk() (*string, bool) {
	if o == nil || utils.IsNil(o.SignerKey) {
		return nil, false
	}
	return o.SignerKey, true
}

// HasSignerKey returns a boolean if a field has been set.
func (o *RosettaOptions) HasSignerKey() bool {
	if o != nil && !utils.IsNil(o.SignerKey) {
		return true
	}

	return false
}

// SetSignerKey gets a reference to the given string and assigns it to the SignerKey field.
func (o *RosettaOptions) SetSignerKey(v string) {
	o.SignerKey = &v
}

// GetSignerPrivateKey returns the SignerPrivateKey field value if set, zero value otherwise.
func (o *RosettaOptions) GetSignerPrivateKey() string {
	if o == nil || utils.IsNil(o.SignerPrivateKey) {
		var ret string
		return ret
	}
	return *o.SignerPrivateKey
}

// GetSignerPrivateKeyOk returns a tuple with the SignerPrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaOptions) GetSignerPrivateKeyOk() (*string, bool) {
	if o == nil || utils.IsNil(o.SignerPrivateKey) {
		return nil, false
	}
	return o.SignerPrivateKey, true
}

// HasSignerPrivateKey returns a boolean if a field has been set.
func (o *RosettaOptions) HasSignerPrivateKey() bool {
	if o != nil && !utils.IsNil(o.SignerPrivateKey) {
		return true
	}

	return false
}

// SetSignerPrivateKey gets a reference to the given string and assigns it to the SignerPrivateKey field.
func (o *RosettaOptions) SetSignerPrivateKey(v string) {
	o.SignerPrivateKey = &v
}

// GetSignerSignature returns the SignerSignature field value if set, zero value otherwise.
func (o *RosettaOptions) GetSignerSignature() string {
	if o == nil || utils.IsNil(o.SignerSignature) {
		var ret string
		return ret
	}
	return *o.SignerSignature
}

// GetSignerSignatureOk returns a tuple with the SignerSignature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaOptions) GetSignerSignatureOk() (*string, bool) {
	if o == nil || utils.IsNil(o.SignerSignature) {
		return nil, false
	}
	return o.SignerSignature, true
}

// HasSignerSignature returns a boolean if a field has been set.
func (o *RosettaOptions) HasSignerSignature() bool {
	if o != nil && !utils.IsNil(o.SignerSignature) {
		return true
	}

	return false
}

// SetSignerSignature gets a reference to the given string and assigns it to the SignerSignature field.
func (o *RosettaOptions) SetSignerSignature(v string) {
	o.SignerSignature = &v
}

// GetPoxMaxAmount returns the PoxMaxAmount field value if set, zero value otherwise.
func (o *RosettaOptions) GetPoxMaxAmount() string {
	if o == nil || utils.IsNil(o.PoxMaxAmount) {
		var ret string
		return ret
	}
	return *o.PoxMaxAmount
}

// GetPoxMaxAmountOk returns a tuple with the PoxMaxAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaOptions) GetPoxMaxAmountOk() (*string, bool) {
	if o == nil || utils.IsNil(o.PoxMaxAmount) {
		return nil, false
	}
	return o.PoxMaxAmount, true
}

// HasPoxMaxAmount returns a boolean if a field has been set.
func (o *RosettaOptions) HasPoxMaxAmount() bool {
	if o != nil && !utils.IsNil(o.PoxMaxAmount) {
		return true
	}

	return false
}

// SetPoxMaxAmount gets a reference to the given string and assigns it to the PoxMaxAmount field.
func (o *RosettaOptions) SetPoxMaxAmount(v string) {
	o.PoxMaxAmount = &v
}

// GetPoxAuthId returns the PoxAuthId field value if set, zero value otherwise.
func (o *RosettaOptions) GetPoxAuthId() string {
	if o == nil || utils.IsNil(o.PoxAuthId) {
		var ret string
		return ret
	}
	return *o.PoxAuthId
}

// GetPoxAuthIdOk returns a tuple with the PoxAuthId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaOptions) GetPoxAuthIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.PoxAuthId) {
		return nil, false
	}
	return o.PoxAuthId, true
}

// HasPoxAuthId returns a boolean if a field has been set.
func (o *RosettaOptions) HasPoxAuthId() bool {
	if o != nil && !utils.IsNil(o.PoxAuthId) {
		return true
	}

	return false
}

// SetPoxAuthId gets a reference to the given string and assigns it to the PoxAuthId field.
func (o *RosettaOptions) SetPoxAuthId(v string) {
	o.PoxAuthId = &v
}

func (o RosettaOptions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RosettaOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.SenderAddress) {
		toSerialize["sender_address"] = o.SenderAddress
	}
	if !utils.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !utils.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !utils.IsNil(o.TokenTransferRecipientAddress) {
		toSerialize["token_transfer_recipient_address"] = o.TokenTransferRecipientAddress
	}
	if !utils.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !utils.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !utils.IsNil(o.Decimals) {
		toSerialize["decimals"] = o.Decimals
	}
	if !utils.IsNil(o.GasLimit) {
		toSerialize["gas_limit"] = o.GasLimit
	}
	if !utils.IsNil(o.GasPrice) {
		toSerialize["gas_price"] = o.GasPrice
	}
	if !utils.IsNil(o.SuggestedFeeMultiplier) {
		toSerialize["suggested_fee_multiplier"] = o.SuggestedFeeMultiplier
	}
	if !utils.IsNil(o.MaxFee) {
		toSerialize["max_fee"] = o.MaxFee
	}
	if !utils.IsNil(o.Fee) {
		toSerialize["fee"] = o.Fee
	}
	if !utils.IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !utils.IsNil(o.Memo) {
		toSerialize["memo"] = o.Memo
	}
	if !utils.IsNil(o.NumberOfCycles) {
		toSerialize["number_of_cycles"] = o.NumberOfCycles
	}
	if !utils.IsNil(o.ContractAddress) {
		toSerialize["contract_address"] = o.ContractAddress
	}
	if !utils.IsNil(o.ContractName) {
		toSerialize["contract_name"] = o.ContractName
	}
	if !utils.IsNil(o.BurnBlockHeight) {
		toSerialize["burn_block_height"] = o.BurnBlockHeight
	}
	if !utils.IsNil(o.RewardCycleId) {
		toSerialize["reward_cycle_id"] = o.RewardCycleId
	}
	if !utils.IsNil(o.DelegateTo) {
		toSerialize["delegate_to"] = o.DelegateTo
	}
	if !utils.IsNil(o.PoxAddr) {
		toSerialize["pox_addr"] = o.PoxAddr
	}
	if !utils.IsNil(o.SignerKey) {
		toSerialize["signer_key"] = o.SignerKey
	}
	if !utils.IsNil(o.SignerPrivateKey) {
		toSerialize["signer_private_key"] = o.SignerPrivateKey
	}
	if !utils.IsNil(o.SignerSignature) {
		toSerialize["signer_signature"] = o.SignerSignature
	}
	if !utils.IsNil(o.PoxMaxAmount) {
		toSerialize["pox_max_amount"] = o.PoxMaxAmount
	}
	if !utils.IsNil(o.PoxAuthId) {
		toSerialize["pox_auth_id"] = o.PoxAuthId
	}
	return toSerialize, nil
}

type NullableRosettaOptions struct {
	value *RosettaOptions
	isSet bool
}

func (v NullableRosettaOptions) Get() *RosettaOptions {
	return v.value
}

func (v *NullableRosettaOptions) Set(val *RosettaOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableRosettaOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableRosettaOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRosettaOptions(val *RosettaOptions) *NullableRosettaOptions {
	return &NullableRosettaOptions{value: val, isSet: true}
}

func (v NullableRosettaOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRosettaOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
