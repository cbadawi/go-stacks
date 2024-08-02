package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the BurnchainReward type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &BurnchainReward{}

// BurnchainReward Reward payment made on the burnchain
type BurnchainReward struct {
	// Set to `true` if block corresponds to the canonical burchchain tip
	Canonical bool `json:"canonical"`
	// The hash representing the burnchain block
	BurnBlockHash string `json:"burn_block_hash"`
	// Height of the burnchain block
	BurnBlockHeight int32 `json:"burn_block_height"`
	// The total amount of burnchain tokens burned for this burnchain block, in the smallest unit (e.g. satoshis for Bitcoin)
	BurnAmount string `json:"burn_amount"`
	// The recipient address that received the burnchain rewards, in the format native to the burnchain (e.g. B58 encoded for Bitcoin)
	RewardRecipient string `json:"reward_recipient"`
	// The amount of burnchain tokens rewarded to the recipient, in the smallest unit (e.g. satoshis for Bitcoin)
	RewardAmount string `json:"reward_amount"`
	// The index position of the reward entry, useful for ordering when there's more than one recipient per burnchain block
	RewardIndex int32 `json:"reward_index"`
}

type _BurnchainReward BurnchainReward

// NewBurnchainReward instantiates a new BurnchainReward object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBurnchainReward(canonical bool, burnBlockHash string, burnBlockHeight int32, burnAmount string, rewardRecipient string, rewardAmount string, rewardIndex int32) *BurnchainReward {
	this := BurnchainReward{}
	this.Canonical = canonical
	this.BurnBlockHash = burnBlockHash
	this.BurnBlockHeight = burnBlockHeight
	this.BurnAmount = burnAmount
	this.RewardRecipient = rewardRecipient
	this.RewardAmount = rewardAmount
	this.RewardIndex = rewardIndex
	return &this
}

// NewBurnchainRewardWithDefaults instantiates a new BurnchainReward object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBurnchainRewardWithDefaults() *BurnchainReward {
	this := BurnchainReward{}
	return &this
}

// GetCanonical returns the Canonical field value
func (o *BurnchainReward) GetCanonical() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Canonical
}

// GetCanonicalOk returns a tuple with the Canonical field value
// and a boolean to check if the value has been set.
func (o *BurnchainReward) GetCanonicalOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Canonical, true
}

// SetCanonical sets field value
func (o *BurnchainReward) SetCanonical(v bool) {
	o.Canonical = v
}

// GetBurnBlockHash returns the BurnBlockHash field value
func (o *BurnchainReward) GetBurnBlockHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BurnBlockHash
}

// GetBurnBlockHashOk returns a tuple with the BurnBlockHash field value
// and a boolean to check if the value has been set.
func (o *BurnchainReward) GetBurnBlockHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BurnBlockHash, true
}

// SetBurnBlockHash sets field value
func (o *BurnchainReward) SetBurnBlockHash(v string) {
	o.BurnBlockHash = v
}

// GetBurnBlockHeight returns the BurnBlockHeight field value
func (o *BurnchainReward) GetBurnBlockHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BurnBlockHeight
}

// GetBurnBlockHeightOk returns a tuple with the BurnBlockHeight field value
// and a boolean to check if the value has been set.
func (o *BurnchainReward) GetBurnBlockHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BurnBlockHeight, true
}

// SetBurnBlockHeight sets field value
func (o *BurnchainReward) SetBurnBlockHeight(v int32) {
	o.BurnBlockHeight = v
}

// GetBurnAmount returns the BurnAmount field value
func (o *BurnchainReward) GetBurnAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BurnAmount
}

// GetBurnAmountOk returns a tuple with the BurnAmount field value
// and a boolean to check if the value has been set.
func (o *BurnchainReward) GetBurnAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BurnAmount, true
}

// SetBurnAmount sets field value
func (o *BurnchainReward) SetBurnAmount(v string) {
	o.BurnAmount = v
}

// GetRewardRecipient returns the RewardRecipient field value
func (o *BurnchainReward) GetRewardRecipient() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RewardRecipient
}

// GetRewardRecipientOk returns a tuple with the RewardRecipient field value
// and a boolean to check if the value has been set.
func (o *BurnchainReward) GetRewardRecipientOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RewardRecipient, true
}

// SetRewardRecipient sets field value
func (o *BurnchainReward) SetRewardRecipient(v string) {
	o.RewardRecipient = v
}

// GetRewardAmount returns the RewardAmount field value
func (o *BurnchainReward) GetRewardAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RewardAmount
}

// GetRewardAmountOk returns a tuple with the RewardAmount field value
// and a boolean to check if the value has been set.
func (o *BurnchainReward) GetRewardAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RewardAmount, true
}

// SetRewardAmount sets field value
func (o *BurnchainReward) SetRewardAmount(v string) {
	o.RewardAmount = v
}

// GetRewardIndex returns the RewardIndex field value
func (o *BurnchainReward) GetRewardIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RewardIndex
}

// GetRewardIndexOk returns a tuple with the RewardIndex field value
// and a boolean to check if the value has been set.
func (o *BurnchainReward) GetRewardIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RewardIndex, true
}

// SetRewardIndex sets field value
func (o *BurnchainReward) SetRewardIndex(v int32) {
	o.RewardIndex = v
}

func (o BurnchainReward) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BurnchainReward) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["canonical"] = o.Canonical
	toSerialize["burn_block_hash"] = o.BurnBlockHash
	toSerialize["burn_block_height"] = o.BurnBlockHeight
	toSerialize["burn_amount"] = o.BurnAmount
	toSerialize["reward_recipient"] = o.RewardRecipient
	toSerialize["reward_amount"] = o.RewardAmount
	toSerialize["reward_index"] = o.RewardIndex
	return toSerialize, nil
}

func (o *BurnchainReward) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"canonical",
		"burn_block_hash",
		"burn_block_height",
		"burn_amount",
		"reward_recipient",
		"reward_amount",
		"reward_index",
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

	varBurnchainReward := _BurnchainReward{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBurnchainReward)

	if err != nil {
		return err
	}

	*o = BurnchainReward(varBurnchainReward)

	return err
}

type NullableBurnchainReward struct {
	value *BurnchainReward
	isSet bool
}

func (v NullableBurnchainReward) Get() *BurnchainReward {
	return v.value
}

func (v *NullableBurnchainReward) Set(val *BurnchainReward) {
	v.value = val
	v.isSet = true
}

func (v NullableBurnchainReward) IsSet() bool {
	return v.isSet
}

func (v *NullableBurnchainReward) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBurnchainReward(val *BurnchainReward) *NullableBurnchainReward {
	return &NullableBurnchainReward{value: val, isSet: true}
}

func (v NullableBurnchainReward) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBurnchainReward) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
