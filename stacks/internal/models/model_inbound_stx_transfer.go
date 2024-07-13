package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the InboundStxTransfer type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &InboundStxTransfer{}

// InboundStxTransfer A inbound STX transfer with a memo
type InboundStxTransfer struct {
	// Principal that sent this transfer
	Sender string `json:"sender"`
	// Transfer amount in micro-STX as integer string
	Amount string `json:"amount"`
	// Hex encoded memo bytes associated with the transfer
	Memo string `json:"memo"`
	// Block height at which this transfer occurred
	BlockHeight float32 `json:"block_height"`
	// The transaction ID in which this transfer occurred
	TxId string `json:"tx_id"`
	// Indicates if the transfer is from a stx-transfer transaction or a contract-call transaction
	TransferType string `json:"transfer_type"`
	// Index of the transaction within a block
	TxIndex float32 `json:"tx_index"`
}

type _InboundStxTransfer InboundStxTransfer

// NewInboundStxTransfer instantiates a new InboundStxTransfer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInboundStxTransfer(sender string, amount string, memo string, blockHeight float32, txId string, transferType string, txIndex float32) *InboundStxTransfer {
	this := InboundStxTransfer{}
	this.Sender = sender
	this.Amount = amount
	this.Memo = memo
	this.BlockHeight = blockHeight
	this.TxId = txId
	this.TransferType = transferType
	this.TxIndex = txIndex
	return &this
}

// NewInboundStxTransferWithDefaults instantiates a new InboundStxTransfer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInboundStxTransferWithDefaults() *InboundStxTransfer {
	this := InboundStxTransfer{}
	return &this
}

// GetSender returns the Sender field value
func (o *InboundStxTransfer) GetSender() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sender
}

// GetSenderOk returns a tuple with the Sender field value
// and a boolean to check if the value has been set.
func (o *InboundStxTransfer) GetSenderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sender, true
}

// SetSender sets field value
func (o *InboundStxTransfer) SetSender(v string) {
	o.Sender = v
}

// GetAmount returns the Amount field value
func (o *InboundStxTransfer) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *InboundStxTransfer) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *InboundStxTransfer) SetAmount(v string) {
	o.Amount = v
}

// GetMemo returns the Memo field value
func (o *InboundStxTransfer) GetMemo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Memo
}

// GetMemoOk returns a tuple with the Memo field value
// and a boolean to check if the value has been set.
func (o *InboundStxTransfer) GetMemoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Memo, true
}

// SetMemo sets field value
func (o *InboundStxTransfer) SetMemo(v string) {
	o.Memo = v
}

// GetBlockHeight returns the BlockHeight field value
func (o *InboundStxTransfer) GetBlockHeight() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.BlockHeight
}

// GetBlockHeightOk returns a tuple with the BlockHeight field value
// and a boolean to check if the value has been set.
func (o *InboundStxTransfer) GetBlockHeightOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockHeight, true
}

// SetBlockHeight sets field value
func (o *InboundStxTransfer) SetBlockHeight(v float32) {
	o.BlockHeight = v
}

// GetTxId returns the TxId field value
func (o *InboundStxTransfer) GetTxId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TxId
}

// GetTxIdOk returns a tuple with the TxId field value
// and a boolean to check if the value has been set.
func (o *InboundStxTransfer) GetTxIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxId, true
}

// SetTxId sets field value
func (o *InboundStxTransfer) SetTxId(v string) {
	o.TxId = v
}

// GetTransferType returns the TransferType field value
func (o *InboundStxTransfer) GetTransferType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransferType
}

// GetTransferTypeOk returns a tuple with the TransferType field value
// and a boolean to check if the value has been set.
func (o *InboundStxTransfer) GetTransferTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransferType, true
}

// SetTransferType sets field value
func (o *InboundStxTransfer) SetTransferType(v string) {
	o.TransferType = v
}

// GetTxIndex returns the TxIndex field value
func (o *InboundStxTransfer) GetTxIndex() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TxIndex
}

// GetTxIndexOk returns a tuple with the TxIndex field value
// and a boolean to check if the value has been set.
func (o *InboundStxTransfer) GetTxIndexOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxIndex, true
}

// SetTxIndex sets field value
func (o *InboundStxTransfer) SetTxIndex(v float32) {
	o.TxIndex = v
}

func (o InboundStxTransfer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InboundStxTransfer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sender"] = o.Sender
	toSerialize["amount"] = o.Amount
	toSerialize["memo"] = o.Memo
	toSerialize["block_height"] = o.BlockHeight
	toSerialize["tx_id"] = o.TxId
	toSerialize["transfer_type"] = o.TransferType
	toSerialize["tx_index"] = o.TxIndex
	return toSerialize, nil
}

func (o *InboundStxTransfer) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sender",
		"amount",
		"memo",
		"block_height",
		"tx_id",
		"transfer_type",
		"tx_index",
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

	varInboundStxTransfer := _InboundStxTransfer{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInboundStxTransfer)

	if err != nil {
		return err
	}

	*o = InboundStxTransfer(varInboundStxTransfer)

	return err
}

type NullableInboundStxTransfer struct {
	value *InboundStxTransfer
	isSet bool
}

func (v NullableInboundStxTransfer) Get() *InboundStxTransfer {
	return v.value
}

func (v *NullableInboundStxTransfer) Set(val *InboundStxTransfer) {
	v.value = val
	v.isSet = true
}

func (v NullableInboundStxTransfer) IsSet() bool {
	return v.isSet
}

func (v *NullableInboundStxTransfer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInboundStxTransfer(val *InboundStxTransfer) *NullableInboundStxTransfer {
	return &NullableInboundStxTransfer{value: val, isSet: true}
}

func (v NullableInboundStxTransfer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInboundStxTransfer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
