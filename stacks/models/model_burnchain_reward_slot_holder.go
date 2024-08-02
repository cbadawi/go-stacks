package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the BurnchainRewardSlotHolder type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &BurnchainRewardSlotHolder{}

// BurnchainRewardSlotHolder Reward slot holder on the burnchain
type BurnchainRewardSlotHolder struct {
	// Set to `true` if block corresponds to the canonical burchchain tip
	Canonical bool `json:"canonical"`
	// The hash representing the burnchain block
	BurnBlockHash string `json:"burn_block_hash"`
	// Height of the burnchain block
	BurnBlockHeight int32 `json:"burn_block_height"`
	// The recipient address that validly received PoX commitments, in the format native to the burnchain (e.g. B58 encoded for Bitcoin)
	Address string `json:"address"`
	// The index position of the reward entry, useful for ordering when there's more than one slot per burnchain block
	SlotIndex int32 `json:"slot_index"`
}

type _BurnchainRewardSlotHolder BurnchainRewardSlotHolder

// NewBurnchainRewardSlotHolder instantiates a new BurnchainRewardSlotHolder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBurnchainRewardSlotHolder(canonical bool, burnBlockHash string, burnBlockHeight int32, address string, slotIndex int32) *BurnchainRewardSlotHolder {
	this := BurnchainRewardSlotHolder{}
	this.Canonical = canonical
	this.BurnBlockHash = burnBlockHash
	this.BurnBlockHeight = burnBlockHeight
	this.Address = address
	this.SlotIndex = slotIndex
	return &this
}

// NewBurnchainRewardSlotHolderWithDefaults instantiates a new BurnchainRewardSlotHolder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBurnchainRewardSlotHolderWithDefaults() *BurnchainRewardSlotHolder {
	this := BurnchainRewardSlotHolder{}
	return &this
}

// GetCanonical returns the Canonical field value
func (o *BurnchainRewardSlotHolder) GetCanonical() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Canonical
}

// GetCanonicalOk returns a tuple with the Canonical field value
// and a boolean to check if the value has been set.
func (o *BurnchainRewardSlotHolder) GetCanonicalOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Canonical, true
}

// SetCanonical sets field value
func (o *BurnchainRewardSlotHolder) SetCanonical(v bool) {
	o.Canonical = v
}

// GetBurnBlockHash returns the BurnBlockHash field value
func (o *BurnchainRewardSlotHolder) GetBurnBlockHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BurnBlockHash
}

// GetBurnBlockHashOk returns a tuple with the BurnBlockHash field value
// and a boolean to check if the value has been set.
func (o *BurnchainRewardSlotHolder) GetBurnBlockHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BurnBlockHash, true
}

// SetBurnBlockHash sets field value
func (o *BurnchainRewardSlotHolder) SetBurnBlockHash(v string) {
	o.BurnBlockHash = v
}

// GetBurnBlockHeight returns the BurnBlockHeight field value
func (o *BurnchainRewardSlotHolder) GetBurnBlockHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BurnBlockHeight
}

// GetBurnBlockHeightOk returns a tuple with the BurnBlockHeight field value
// and a boolean to check if the value has been set.
func (o *BurnchainRewardSlotHolder) GetBurnBlockHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BurnBlockHeight, true
}

// SetBurnBlockHeight sets field value
func (o *BurnchainRewardSlotHolder) SetBurnBlockHeight(v int32) {
	o.BurnBlockHeight = v
}

// GetAddress returns the Address field value
func (o *BurnchainRewardSlotHolder) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *BurnchainRewardSlotHolder) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *BurnchainRewardSlotHolder) SetAddress(v string) {
	o.Address = v
}

// GetSlotIndex returns the SlotIndex field value
func (o *BurnchainRewardSlotHolder) GetSlotIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SlotIndex
}

// GetSlotIndexOk returns a tuple with the SlotIndex field value
// and a boolean to check if the value has been set.
func (o *BurnchainRewardSlotHolder) GetSlotIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SlotIndex, true
}

// SetSlotIndex sets field value
func (o *BurnchainRewardSlotHolder) SetSlotIndex(v int32) {
	o.SlotIndex = v
}

func (o BurnchainRewardSlotHolder) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BurnchainRewardSlotHolder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["canonical"] = o.Canonical
	toSerialize["burn_block_hash"] = o.BurnBlockHash
	toSerialize["burn_block_height"] = o.BurnBlockHeight
	toSerialize["address"] = o.Address
	toSerialize["slot_index"] = o.SlotIndex
	return toSerialize, nil
}

func (o *BurnchainRewardSlotHolder) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"canonical",
		"burn_block_hash",
		"burn_block_height",
		"address",
		"slot_index",
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

	varBurnchainRewardSlotHolder := _BurnchainRewardSlotHolder{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBurnchainRewardSlotHolder)

	if err != nil {
		return err
	}

	*o = BurnchainRewardSlotHolder(varBurnchainRewardSlotHolder)

	return err
}

type NullableBurnchainRewardSlotHolder struct {
	value *BurnchainRewardSlotHolder
	isSet bool
}

func (v NullableBurnchainRewardSlotHolder) Get() *BurnchainRewardSlotHolder {
	return v.value
}

func (v *NullableBurnchainRewardSlotHolder) Set(val *BurnchainRewardSlotHolder) {
	v.value = val
	v.isSet = true
}

func (v NullableBurnchainRewardSlotHolder) IsSet() bool {
	return v.isSet
}

func (v *NullableBurnchainRewardSlotHolder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBurnchainRewardSlotHolder(val *BurnchainRewardSlotHolder) *NullableBurnchainRewardSlotHolder {
	return &NullableBurnchainRewardSlotHolder{value: val, isSet: true}
}

func (v NullableBurnchainRewardSlotHolder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBurnchainRewardSlotHolder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
