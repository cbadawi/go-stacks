package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the PoolDelegation type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PoolDelegation{}

// PoolDelegation struct for PoolDelegation
type PoolDelegation struct {
	// The principal of the pool member that issued the delegation
	Stacker string `json:"stacker"`
	// The pox-addr value specified by the stacker in the delegation operation
	PoxAddr *string `json:"pox_addr,omitempty"`
	// The amount of uSTX delegated by the stacker
	AmountUstx string `json:"amount_ustx"`
	// The optional burnchain block unlock height that the stacker may have specified
	BurnBlockUnlockHeight *int32 `json:"burn_block_unlock_height,omitempty"`
	// The block height at which the stacker delegation transaction was mined at
	BlockHeight int32 `json:"block_height"`
	// The tx_id of the stacker delegation operation
	TxId string `json:"tx_id"`
}

type _PoolDelegation PoolDelegation

// NewPoolDelegation instantiates a new PoolDelegation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoolDelegation(stacker string, amountUstx string, blockHeight int32, txId string) *PoolDelegation {
	this := PoolDelegation{}
	this.Stacker = stacker
	this.AmountUstx = amountUstx
	this.BlockHeight = blockHeight
	this.TxId = txId
	return &this
}

// NewPoolDelegationWithDefaults instantiates a new PoolDelegation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoolDelegationWithDefaults() *PoolDelegation {
	this := PoolDelegation{}
	return &this
}

// GetStacker returns the Stacker field value
func (o *PoolDelegation) GetStacker() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Stacker
}

// GetStackerOk returns a tuple with the Stacker field value
// and a boolean to check if the value has been set.
func (o *PoolDelegation) GetStackerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stacker, true
}

// SetStacker sets field value
func (o *PoolDelegation) SetStacker(v string) {
	o.Stacker = v
}

// GetPoxAddr returns the PoxAddr field value if set, zero value otherwise.
func (o *PoolDelegation) GetPoxAddr() string {
	if o == nil || utils.IsNil(o.PoxAddr) {
		var ret string
		return ret
	}
	return *o.PoxAddr
}

// GetPoxAddrOk returns a tuple with the PoxAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolDelegation) GetPoxAddrOk() (*string, bool) {
	if o == nil || utils.IsNil(o.PoxAddr) {
		return nil, false
	}
	return o.PoxAddr, true
}

// HasPoxAddr returns a boolean if a field has been set.
func (o *PoolDelegation) HasPoxAddr() bool {
	if o != nil && !utils.IsNil(o.PoxAddr) {
		return true
	}

	return false
}

// SetPoxAddr gets a reference to the given string and assigns it to the PoxAddr field.
func (o *PoolDelegation) SetPoxAddr(v string) {
	o.PoxAddr = &v
}

// GetAmountUstx returns the AmountUstx field value
func (o *PoolDelegation) GetAmountUstx() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AmountUstx
}

// GetAmountUstxOk returns a tuple with the AmountUstx field value
// and a boolean to check if the value has been set.
func (o *PoolDelegation) GetAmountUstxOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountUstx, true
}

// SetAmountUstx sets field value
func (o *PoolDelegation) SetAmountUstx(v string) {
	o.AmountUstx = v
}

// GetBurnBlockUnlockHeight returns the BurnBlockUnlockHeight field value if set, zero value otherwise.
func (o *PoolDelegation) GetBurnBlockUnlockHeight() int32 {
	if o == nil || utils.IsNil(o.BurnBlockUnlockHeight) {
		var ret int32
		return ret
	}
	return *o.BurnBlockUnlockHeight
}

// GetBurnBlockUnlockHeightOk returns a tuple with the BurnBlockUnlockHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolDelegation) GetBurnBlockUnlockHeightOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.BurnBlockUnlockHeight) {
		return nil, false
	}
	return o.BurnBlockUnlockHeight, true
}

// HasBurnBlockUnlockHeight returns a boolean if a field has been set.
func (o *PoolDelegation) HasBurnBlockUnlockHeight() bool {
	if o != nil && !utils.IsNil(o.BurnBlockUnlockHeight) {
		return true
	}

	return false
}

// SetBurnBlockUnlockHeight gets a reference to the given int32 and assigns it to the BurnBlockUnlockHeight field.
func (o *PoolDelegation) SetBurnBlockUnlockHeight(v int32) {
	o.BurnBlockUnlockHeight = &v
}

// GetBlockHeight returns the BlockHeight field value
func (o *PoolDelegation) GetBlockHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BlockHeight
}

// GetBlockHeightOk returns a tuple with the BlockHeight field value
// and a boolean to check if the value has been set.
func (o *PoolDelegation) GetBlockHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockHeight, true
}

// SetBlockHeight sets field value
func (o *PoolDelegation) SetBlockHeight(v int32) {
	o.BlockHeight = v
}

// GetTxId returns the TxId field value
func (o *PoolDelegation) GetTxId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TxId
}

// GetTxIdOk returns a tuple with the TxId field value
// and a boolean to check if the value has been set.
func (o *PoolDelegation) GetTxIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxId, true
}

// SetTxId sets field value
func (o *PoolDelegation) SetTxId(v string) {
	o.TxId = v
}

func (o PoolDelegation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PoolDelegation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["stacker"] = o.Stacker
	if !utils.IsNil(o.PoxAddr) {
		toSerialize["pox_addr"] = o.PoxAddr
	}
	toSerialize["amount_ustx"] = o.AmountUstx
	if !utils.IsNil(o.BurnBlockUnlockHeight) {
		toSerialize["burn_block_unlock_height"] = o.BurnBlockUnlockHeight
	}
	toSerialize["block_height"] = o.BlockHeight
	toSerialize["tx_id"] = o.TxId
	return toSerialize, nil
}

func (o *PoolDelegation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"stacker",
		"amount_ustx",
		"block_height",
		"tx_id",
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

	varPoolDelegation := _PoolDelegation{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPoolDelegation)

	if err != nil {
		return err
	}

	*o = PoolDelegation(varPoolDelegation)

	return err
}

type NullablePoolDelegation struct {
	value *PoolDelegation
	isSet bool
}

func (v NullablePoolDelegation) Get() *PoolDelegation {
	return v.value
}

func (v *NullablePoolDelegation) Set(val *PoolDelegation) {
	v.value = val
	v.isSet = true
}

func (v NullablePoolDelegation) IsSet() bool {
	return v.isSet
}

func (v *NullablePoolDelegation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoolDelegation(val *PoolDelegation) *NullablePoolDelegation {
	return &NullablePoolDelegation{value: val, isSet: true}
}

func (v NullablePoolDelegation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoolDelegation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
