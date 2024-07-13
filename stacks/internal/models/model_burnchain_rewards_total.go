package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the BurnchainRewardsTotal type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &BurnchainRewardsTotal{}

// BurnchainRewardsTotal Total burnchain rewards made to a recipient
type BurnchainRewardsTotal struct {
	// The recipient address that received the burnchain rewards, in the format native to the burnchain (e.g. B58 encoded for Bitcoin)
	RewardRecipient string `json:"reward_recipient"`
	// The total amount of burnchain tokens rewarded to the recipient, in the smallest unit (e.g. satoshis for Bitcoin)
	RewardAmount string `json:"reward_amount"`
}

type _BurnchainRewardsTotal BurnchainRewardsTotal

// NewBurnchainRewardsTotal instantiates a new BurnchainRewardsTotal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBurnchainRewardsTotal(rewardRecipient string, rewardAmount string) *BurnchainRewardsTotal {
	this := BurnchainRewardsTotal{}
	this.RewardRecipient = rewardRecipient
	this.RewardAmount = rewardAmount
	return &this
}

// NewBurnchainRewardsTotalWithDefaults instantiates a new BurnchainRewardsTotal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBurnchainRewardsTotalWithDefaults() *BurnchainRewardsTotal {
	this := BurnchainRewardsTotal{}
	return &this
}

// GetRewardRecipient returns the RewardRecipient field value
func (o *BurnchainRewardsTotal) GetRewardRecipient() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RewardRecipient
}

// GetRewardRecipientOk returns a tuple with the RewardRecipient field value
// and a boolean to check if the value has been set.
func (o *BurnchainRewardsTotal) GetRewardRecipientOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RewardRecipient, true
}

// SetRewardRecipient sets field value
func (o *BurnchainRewardsTotal) SetRewardRecipient(v string) {
	o.RewardRecipient = v
}

// GetRewardAmount returns the RewardAmount field value
func (o *BurnchainRewardsTotal) GetRewardAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RewardAmount
}

// GetRewardAmountOk returns a tuple with the RewardAmount field value
// and a boolean to check if the value has been set.
func (o *BurnchainRewardsTotal) GetRewardAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RewardAmount, true
}

// SetRewardAmount sets field value
func (o *BurnchainRewardsTotal) SetRewardAmount(v string) {
	o.RewardAmount = v
}

func (o BurnchainRewardsTotal) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BurnchainRewardsTotal) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reward_recipient"] = o.RewardRecipient
	toSerialize["reward_amount"] = o.RewardAmount
	return toSerialize, nil
}

func (o *BurnchainRewardsTotal) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"reward_recipient",
		"reward_amount",
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

	varBurnchainRewardsTotal := _BurnchainRewardsTotal{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBurnchainRewardsTotal)

	if err != nil {
		return err
	}

	*o = BurnchainRewardsTotal(varBurnchainRewardsTotal)

	return err
}

type NullableBurnchainRewardsTotal struct {
	value *BurnchainRewardsTotal
	isSet bool
}

func (v NullableBurnchainRewardsTotal) Get() *BurnchainRewardsTotal {
	return v.value
}

func (v *NullableBurnchainRewardsTotal) Set(val *BurnchainRewardsTotal) {
	v.value = val
	v.isSet = true
}

func (v NullableBurnchainRewardsTotal) IsSet() bool {
	return v.isSet
}

func (v *NullableBurnchainRewardsTotal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBurnchainRewardsTotal(val *BurnchainRewardsTotal) *NullableBurnchainRewardsTotal {
	return &NullableBurnchainRewardsTotal{value: val, isSet: true}
}

func (v NullableBurnchainRewardsTotal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBurnchainRewardsTotal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
