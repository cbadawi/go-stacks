package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the MempoolCoinbaseTransaction type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &MempoolCoinbaseTransaction{}

// MempoolCoinbaseTransaction Describes representation of a Type 3 Stacks 2.0 transaction: Poison Microblock
type MempoolCoinbaseTransaction struct {
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
	// Status of the transaction
	TxStatus string `json:"tx_status"`
	// A unix timestamp (in seconds) indicating when the transaction broadcast was received by the node.
	ReceiptTime float32 `json:"receipt_time"`
	// An ISO 8601 (YYYY-MM-DDTHH:mm:ss.sssZ) timestamp indicating when the transaction broadcast was received by the node.
	ReceiptTimeIso  string                                     `json:"receipt_time_iso"`
	TxType          string                                     `json:"tx_type"`
	CoinbasePayload CoinbaseTransactionMetadataCoinbasePayload `json:"coinbase_payload"`
}

type _MempoolCoinbaseTransaction MempoolCoinbaseTransaction

// NewMempoolCoinbaseTransaction instantiates a new MempoolCoinbaseTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMempoolCoinbaseTransaction(txId string, nonce int32, feeRate string, senderAddress string, sponsored bool, postConditionMode string, postConditions []PostCondition, anchorMode string, txStatus string, receiptTime float32, receiptTimeIso string, txType string, coinbasePayload CoinbaseTransactionMetadataCoinbasePayload) *MempoolCoinbaseTransaction {
	this := MempoolCoinbaseTransaction{}
	this.TxId = txId
	this.Nonce = nonce
	this.FeeRate = feeRate
	this.SenderAddress = senderAddress
	this.Sponsored = sponsored
	this.PostConditionMode = postConditionMode
	this.PostConditions = postConditions
	this.AnchorMode = anchorMode
	this.TxStatus = txStatus
	this.ReceiptTime = receiptTime
	this.ReceiptTimeIso = receiptTimeIso
	this.TxType = txType
	this.CoinbasePayload = coinbasePayload
	return &this
}

// NewMempoolCoinbaseTransactionWithDefaults instantiates a new MempoolCoinbaseTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMempoolCoinbaseTransactionWithDefaults() *MempoolCoinbaseTransaction {
	this := MempoolCoinbaseTransaction{}
	return &this
}

// GetTxId returns the TxId field value
func (o *MempoolCoinbaseTransaction) GetTxId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TxId
}

// GetTxIdOk returns a tuple with the TxId field value
// and a boolean to check if the value has been set.
func (o *MempoolCoinbaseTransaction) GetTxIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxId, true
}

// SetTxId sets field value
func (o *MempoolCoinbaseTransaction) SetTxId(v string) {
	o.TxId = v
}

// GetNonce returns the Nonce field value
func (o *MempoolCoinbaseTransaction) GetNonce() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value
// and a boolean to check if the value has been set.
func (o *MempoolCoinbaseTransaction) GetNonceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nonce, true
}

// SetNonce sets field value
func (o *MempoolCoinbaseTransaction) SetNonce(v int32) {
	o.Nonce = v
}

// GetFeeRate returns the FeeRate field value
func (o *MempoolCoinbaseTransaction) GetFeeRate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FeeRate
}

// GetFeeRateOk returns a tuple with the FeeRate field value
// and a boolean to check if the value has been set.
func (o *MempoolCoinbaseTransaction) GetFeeRateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeeRate, true
}

// SetFeeRate sets field value
func (o *MempoolCoinbaseTransaction) SetFeeRate(v string) {
	o.FeeRate = v
}

// GetSenderAddress returns the SenderAddress field value
func (o *MempoolCoinbaseTransaction) GetSenderAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SenderAddress
}

// GetSenderAddressOk returns a tuple with the SenderAddress field value
// and a boolean to check if the value has been set.
func (o *MempoolCoinbaseTransaction) GetSenderAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SenderAddress, true
}

// SetSenderAddress sets field value
func (o *MempoolCoinbaseTransaction) SetSenderAddress(v string) {
	o.SenderAddress = v
}

// GetSponsorNonce returns the SponsorNonce field value if set, zero value otherwise.
func (o *MempoolCoinbaseTransaction) GetSponsorNonce() int32 {
	if o == nil || utils.IsNil(o.SponsorNonce) {
		var ret int32
		return ret
	}
	return *o.SponsorNonce
}

// GetSponsorNonceOk returns a tuple with the SponsorNonce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MempoolCoinbaseTransaction) GetSponsorNonceOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.SponsorNonce) {
		return nil, false
	}
	return o.SponsorNonce, true
}

// HasSponsorNonce returns a boolean if a field has been set.
func (o *MempoolCoinbaseTransaction) HasSponsorNonce() bool {
	if o != nil && !utils.IsNil(o.SponsorNonce) {
		return true
	}

	return false
}

// SetSponsorNonce gets a reference to the given int32 and assigns it to the SponsorNonce field.
func (o *MempoolCoinbaseTransaction) SetSponsorNonce(v int32) {
	o.SponsorNonce = &v
}

// GetSponsored returns the Sponsored field value
func (o *MempoolCoinbaseTransaction) GetSponsored() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Sponsored
}

// GetSponsoredOk returns a tuple with the Sponsored field value
// and a boolean to check if the value has been set.
func (o *MempoolCoinbaseTransaction) GetSponsoredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sponsored, true
}

// SetSponsored sets field value
func (o *MempoolCoinbaseTransaction) SetSponsored(v bool) {
	o.Sponsored = v
}

// GetSponsorAddress returns the SponsorAddress field value if set, zero value otherwise.
func (o *MempoolCoinbaseTransaction) GetSponsorAddress() string {
	if o == nil || utils.IsNil(o.SponsorAddress) {
		var ret string
		return ret
	}
	return *o.SponsorAddress
}

// GetSponsorAddressOk returns a tuple with the SponsorAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MempoolCoinbaseTransaction) GetSponsorAddressOk() (*string, bool) {
	if o == nil || utils.IsNil(o.SponsorAddress) {
		return nil, false
	}
	return o.SponsorAddress, true
}

// HasSponsorAddress returns a boolean if a field has been set.
func (o *MempoolCoinbaseTransaction) HasSponsorAddress() bool {
	if o != nil && !utils.IsNil(o.SponsorAddress) {
		return true
	}

	return false
}

// SetSponsorAddress gets a reference to the given string and assigns it to the SponsorAddress field.
func (o *MempoolCoinbaseTransaction) SetSponsorAddress(v string) {
	o.SponsorAddress = &v
}

// GetPostConditionMode returns the PostConditionMode field value
func (o *MempoolCoinbaseTransaction) GetPostConditionMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PostConditionMode
}

// GetPostConditionModeOk returns a tuple with the PostConditionMode field value
// and a boolean to check if the value has been set.
func (o *MempoolCoinbaseTransaction) GetPostConditionModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PostConditionMode, true
}

// SetPostConditionMode sets field value
func (o *MempoolCoinbaseTransaction) SetPostConditionMode(v string) {
	o.PostConditionMode = v
}

// GetPostConditions returns the PostConditions field value
func (o *MempoolCoinbaseTransaction) GetPostConditions() []PostCondition {
	if o == nil {
		var ret []PostCondition
		return ret
	}

	return o.PostConditions
}

// GetPostConditionsOk returns a tuple with the PostConditions field value
// and a boolean to check if the value has been set.
func (o *MempoolCoinbaseTransaction) GetPostConditionsOk() ([]PostCondition, bool) {
	if o == nil {
		return nil, false
	}
	return o.PostConditions, true
}

// SetPostConditions sets field value
func (o *MempoolCoinbaseTransaction) SetPostConditions(v []PostCondition) {
	o.PostConditions = v
}

// GetAnchorMode returns the AnchorMode field value
func (o *MempoolCoinbaseTransaction) GetAnchorMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AnchorMode
}

// GetAnchorModeOk returns a tuple with the AnchorMode field value
// and a boolean to check if the value has been set.
func (o *MempoolCoinbaseTransaction) GetAnchorModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AnchorMode, true
}

// SetAnchorMode sets field value
func (o *MempoolCoinbaseTransaction) SetAnchorMode(v string) {
	o.AnchorMode = v
}

// GetTxStatus returns the TxStatus field value
func (o *MempoolCoinbaseTransaction) GetTxStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TxStatus
}

// GetTxStatusOk returns a tuple with the TxStatus field value
// and a boolean to check if the value has been set.
func (o *MempoolCoinbaseTransaction) GetTxStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxStatus, true
}

// SetTxStatus sets field value
func (o *MempoolCoinbaseTransaction) SetTxStatus(v string) {
	o.TxStatus = v
}

// GetReceiptTime returns the ReceiptTime field value
func (o *MempoolCoinbaseTransaction) GetReceiptTime() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ReceiptTime
}

// GetReceiptTimeOk returns a tuple with the ReceiptTime field value
// and a boolean to check if the value has been set.
func (o *MempoolCoinbaseTransaction) GetReceiptTimeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReceiptTime, true
}

// SetReceiptTime sets field value
func (o *MempoolCoinbaseTransaction) SetReceiptTime(v float32) {
	o.ReceiptTime = v
}

// GetReceiptTimeIso returns the ReceiptTimeIso field value
func (o *MempoolCoinbaseTransaction) GetReceiptTimeIso() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReceiptTimeIso
}

// GetReceiptTimeIsoOk returns a tuple with the ReceiptTimeIso field value
// and a boolean to check if the value has been set.
func (o *MempoolCoinbaseTransaction) GetReceiptTimeIsoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReceiptTimeIso, true
}

// SetReceiptTimeIso sets field value
func (o *MempoolCoinbaseTransaction) SetReceiptTimeIso(v string) {
	o.ReceiptTimeIso = v
}

// GetTxType returns the TxType field value
func (o *MempoolCoinbaseTransaction) GetTxType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TxType
}

// GetTxTypeOk returns a tuple with the TxType field value
// and a boolean to check if the value has been set.
func (o *MempoolCoinbaseTransaction) GetTxTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxType, true
}

// SetTxType sets field value
func (o *MempoolCoinbaseTransaction) SetTxType(v string) {
	o.TxType = v
}

// GetCoinbasePayload returns the CoinbasePayload field value
func (o *MempoolCoinbaseTransaction) GetCoinbasePayload() CoinbaseTransactionMetadataCoinbasePayload {
	if o == nil {
		var ret CoinbaseTransactionMetadataCoinbasePayload
		return ret
	}

	return o.CoinbasePayload
}

// GetCoinbasePayloadOk returns a tuple with the CoinbasePayload field value
// and a boolean to check if the value has been set.
func (o *MempoolCoinbaseTransaction) GetCoinbasePayloadOk() (*CoinbaseTransactionMetadataCoinbasePayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CoinbasePayload, true
}

// SetCoinbasePayload sets field value
func (o *MempoolCoinbaseTransaction) SetCoinbasePayload(v CoinbaseTransactionMetadataCoinbasePayload) {
	o.CoinbasePayload = v
}

func (o MempoolCoinbaseTransaction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MempoolCoinbaseTransaction) ToMap() (map[string]interface{}, error) {
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
	toSerialize["tx_status"] = o.TxStatus
	toSerialize["receipt_time"] = o.ReceiptTime
	toSerialize["receipt_time_iso"] = o.ReceiptTimeIso
	toSerialize["tx_type"] = o.TxType
	toSerialize["coinbase_payload"] = o.CoinbasePayload
	return toSerialize, nil
}

func (o *MempoolCoinbaseTransaction) UnmarshalJSON(data []byte) (err error) {
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
		"tx_status",
		"receipt_time",
		"receipt_time_iso",
		"tx_type",
		"coinbase_payload",
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

	varMempoolCoinbaseTransaction := _MempoolCoinbaseTransaction{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMempoolCoinbaseTransaction)

	if err != nil {
		return err
	}

	*o = MempoolCoinbaseTransaction(varMempoolCoinbaseTransaction)

	return err
}

type NullableMempoolCoinbaseTransaction struct {
	value *MempoolCoinbaseTransaction
	isSet bool
}

func (v NullableMempoolCoinbaseTransaction) Get() *MempoolCoinbaseTransaction {
	return v.value
}

func (v *NullableMempoolCoinbaseTransaction) Set(val *MempoolCoinbaseTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableMempoolCoinbaseTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableMempoolCoinbaseTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMempoolCoinbaseTransaction(val *MempoolCoinbaseTransaction) *NullableMempoolCoinbaseTransaction {
	return &NullableMempoolCoinbaseTransaction{value: val, isSet: true}
}

func (v NullableMempoolCoinbaseTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMempoolCoinbaseTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
