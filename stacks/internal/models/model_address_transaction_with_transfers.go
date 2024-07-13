package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the AddressTransactionWithTransfers type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AddressTransactionWithTransfers{}

// AddressTransactionWithTransfers Transaction with STX transfers for a given address
type AddressTransactionWithTransfers struct {
	Tx Transaction `json:"tx"`
	// Total sent from the given address, including the tx fee, in micro-STX as an integer string.
	StxSent string `json:"stx_sent"`
	// Total received by the given address in micro-STX as an integer string.
	StxReceived  string                                             `json:"stx_received"`
	StxTransfers []AddressTransactionWithTransfersStxTransfersInner `json:"stx_transfers"`
	FtTransfers  []AddressTransactionWithTransfersFtTransfersInner  `json:"ft_transfers,omitempty"`
	NftTransfers []AddressTransactionWithTransfersNftTransfersInner `json:"nft_transfers,omitempty"`
}

type _AddressTransactionWithTransfers AddressTransactionWithTransfers

// NewAddressTransactionWithTransfers instantiates a new AddressTransactionWithTransfers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressTransactionWithTransfers(tx Transaction, stxSent string, stxReceived string, stxTransfers []AddressTransactionWithTransfersStxTransfersInner) *AddressTransactionWithTransfers {
	this := AddressTransactionWithTransfers{}
	this.Tx = tx
	this.StxSent = stxSent
	this.StxReceived = stxReceived
	this.StxTransfers = stxTransfers
	return &this
}

// NewAddressTransactionWithTransfersWithDefaults instantiates a new AddressTransactionWithTransfers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressTransactionWithTransfersWithDefaults() *AddressTransactionWithTransfers {
	this := AddressTransactionWithTransfers{}
	return &this
}

// GetTx returns the Tx field value
func (o *AddressTransactionWithTransfers) GetTx() Transaction {
	if o == nil {
		var ret Transaction
		return ret
	}

	return o.Tx
}

// GetTxOk returns a tuple with the Tx field value
// and a boolean to check if the value has been set.
func (o *AddressTransactionWithTransfers) GetTxOk() (*Transaction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tx, true
}

// SetTx sets field value
func (o *AddressTransactionWithTransfers) SetTx(v Transaction) {
	o.Tx = v
}

// GetStxSent returns the StxSent field value
func (o *AddressTransactionWithTransfers) GetStxSent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StxSent
}

// GetStxSentOk returns a tuple with the StxSent field value
// and a boolean to check if the value has been set.
func (o *AddressTransactionWithTransfers) GetStxSentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StxSent, true
}

// SetStxSent sets field value
func (o *AddressTransactionWithTransfers) SetStxSent(v string) {
	o.StxSent = v
}

// GetStxReceived returns the StxReceived field value
func (o *AddressTransactionWithTransfers) GetStxReceived() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StxReceived
}

// GetStxReceivedOk returns a tuple with the StxReceived field value
// and a boolean to check if the value has been set.
func (o *AddressTransactionWithTransfers) GetStxReceivedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StxReceived, true
}

// SetStxReceived sets field value
func (o *AddressTransactionWithTransfers) SetStxReceived(v string) {
	o.StxReceived = v
}

// GetStxTransfers returns the StxTransfers field value
func (o *AddressTransactionWithTransfers) GetStxTransfers() []AddressTransactionWithTransfersStxTransfersInner {
	if o == nil {
		var ret []AddressTransactionWithTransfersStxTransfersInner
		return ret
	}

	return o.StxTransfers
}

// GetStxTransfersOk returns a tuple with the StxTransfers field value
// and a boolean to check if the value has been set.
func (o *AddressTransactionWithTransfers) GetStxTransfersOk() ([]AddressTransactionWithTransfersStxTransfersInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.StxTransfers, true
}

// SetStxTransfers sets field value
func (o *AddressTransactionWithTransfers) SetStxTransfers(v []AddressTransactionWithTransfersStxTransfersInner) {
	o.StxTransfers = v
}

// GetFtTransfers returns the FtTransfers field value if set, zero value otherwise.
func (o *AddressTransactionWithTransfers) GetFtTransfers() []AddressTransactionWithTransfersFtTransfersInner {
	if o == nil || utils.IsNil(o.FtTransfers) {
		var ret []AddressTransactionWithTransfersFtTransfersInner
		return ret
	}
	return o.FtTransfers
}

// GetFtTransfersOk returns a tuple with the FtTransfers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressTransactionWithTransfers) GetFtTransfersOk() ([]AddressTransactionWithTransfersFtTransfersInner, bool) {
	if o == nil || utils.IsNil(o.FtTransfers) {
		return nil, false
	}
	return o.FtTransfers, true
}

// HasFtTransfers returns a boolean if a field has been set.
func (o *AddressTransactionWithTransfers) HasFtTransfers() bool {
	if o != nil && !utils.IsNil(o.FtTransfers) {
		return true
	}

	return false
}

// SetFtTransfers gets a reference to the given []AddressTransactionWithTransfersFtTransfersInner and assigns it to the FtTransfers field.
func (o *AddressTransactionWithTransfers) SetFtTransfers(v []AddressTransactionWithTransfersFtTransfersInner) {
	o.FtTransfers = v
}

// GetNftTransfers returns the NftTransfers field value if set, zero value otherwise.
func (o *AddressTransactionWithTransfers) GetNftTransfers() []AddressTransactionWithTransfersNftTransfersInner {
	if o == nil || utils.IsNil(o.NftTransfers) {
		var ret []AddressTransactionWithTransfersNftTransfersInner
		return ret
	}
	return o.NftTransfers
}

// GetNftTransfersOk returns a tuple with the NftTransfers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressTransactionWithTransfers) GetNftTransfersOk() ([]AddressTransactionWithTransfersNftTransfersInner, bool) {
	if o == nil || utils.IsNil(o.NftTransfers) {
		return nil, false
	}
	return o.NftTransfers, true
}

// HasNftTransfers returns a boolean if a field has been set.
func (o *AddressTransactionWithTransfers) HasNftTransfers() bool {
	if o != nil && !utils.IsNil(o.NftTransfers) {
		return true
	}

	return false
}

// SetNftTransfers gets a reference to the given []AddressTransactionWithTransfersNftTransfersInner and assigns it to the NftTransfers field.
func (o *AddressTransactionWithTransfers) SetNftTransfers(v []AddressTransactionWithTransfersNftTransfersInner) {
	o.NftTransfers = v
}

func (o AddressTransactionWithTransfers) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressTransactionWithTransfers) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tx"] = o.Tx
	toSerialize["stx_sent"] = o.StxSent
	toSerialize["stx_received"] = o.StxReceived
	toSerialize["stx_transfers"] = o.StxTransfers
	if !utils.IsNil(o.FtTransfers) {
		toSerialize["ft_transfers"] = o.FtTransfers
	}
	if !utils.IsNil(o.NftTransfers) {
		toSerialize["nft_transfers"] = o.NftTransfers
	}
	return toSerialize, nil
}

func (o *AddressTransactionWithTransfers) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tx",
		"stx_sent",
		"stx_received",
		"stx_transfers",
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

	varAddressTransactionWithTransfers := _AddressTransactionWithTransfers{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddressTransactionWithTransfers)

	if err != nil {
		return err
	}

	*o = AddressTransactionWithTransfers(varAddressTransactionWithTransfers)

	return err
}

type NullableAddressTransactionWithTransfers struct {
	value *AddressTransactionWithTransfers
	isSet bool
}

func (v NullableAddressTransactionWithTransfers) Get() *AddressTransactionWithTransfers {
	return v.value
}

func (v *NullableAddressTransactionWithTransfers) Set(val *AddressTransactionWithTransfers) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressTransactionWithTransfers) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressTransactionWithTransfers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressTransactionWithTransfers(val *AddressTransactionWithTransfers) *NullableAddressTransactionWithTransfers {
	return &NullableAddressTransactionWithTransfers{value: val, isSet: true}
}

func (v NullableAddressTransactionWithTransfers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressTransactionWithTransfers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
