package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the StxBalance type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &StxBalance{}

// StxBalance struct for StxBalance
type StxBalance struct {
	Balance                   string `json:"balance"`
	TotalSent                 string `json:"total_sent"`
	TotalReceived             string `json:"total_received"`
	TotalFeesSent             string `json:"total_fees_sent"`
	TotalMinerRewardsReceived string `json:"total_miner_rewards_received"`
	// The transaction where the lock event occurred. Empty if no tokens are locked.
	LockTxId string `json:"lock_tx_id"`
	// The amount of locked STX, as string quoted micro-STX. Zero if no tokens are locked.
	Locked string `json:"locked"`
	// The STX chain block height of when the lock event occurred. Zero if no tokens are locked.
	LockHeight int32 `json:"lock_height"`
	// The burnchain block height of when the lock event occurred. Zero if no tokens are locked.
	BurnchainLockHeight int32 `json:"burnchain_lock_height"`
	// The burnchain block height of when the tokens unlock. Zero if no tokens are locked.
	BurnchainUnlockHeight int32 `json:"burnchain_unlock_height"`
}

type _StxBalance StxBalance

// NewStxBalance instantiates a new StxBalance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStxBalance(balance string, totalSent string, totalReceived string, totalFeesSent string, totalMinerRewardsReceived string, lockTxId string, locked string, lockHeight int32, burnchainLockHeight int32, burnchainUnlockHeight int32) *StxBalance {
	this := StxBalance{}
	this.Balance = balance
	this.TotalSent = totalSent
	this.TotalReceived = totalReceived
	this.TotalFeesSent = totalFeesSent
	this.TotalMinerRewardsReceived = totalMinerRewardsReceived
	this.LockTxId = lockTxId
	this.Locked = locked
	this.LockHeight = lockHeight
	this.BurnchainLockHeight = burnchainLockHeight
	this.BurnchainUnlockHeight = burnchainUnlockHeight
	return &this
}

// NewStxBalanceWithDefaults instantiates a new StxBalance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStxBalanceWithDefaults() *StxBalance {
	this := StxBalance{}
	return &this
}

// GetBalance returns the Balance field value
func (o *StxBalance) GetBalance() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value
// and a boolean to check if the value has been set.
func (o *StxBalance) GetBalanceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Balance, true
}

// SetBalance sets field value
func (o *StxBalance) SetBalance(v string) {
	o.Balance = v
}

// GetTotalSent returns the TotalSent field value
func (o *StxBalance) GetTotalSent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TotalSent
}

// GetTotalSentOk returns a tuple with the TotalSent field value
// and a boolean to check if the value has been set.
func (o *StxBalance) GetTotalSentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalSent, true
}

// SetTotalSent sets field value
func (o *StxBalance) SetTotalSent(v string) {
	o.TotalSent = v
}

// GetTotalReceived returns the TotalReceived field value
func (o *StxBalance) GetTotalReceived() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TotalReceived
}

// GetTotalReceivedOk returns a tuple with the TotalReceived field value
// and a boolean to check if the value has been set.
func (o *StxBalance) GetTotalReceivedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalReceived, true
}

// SetTotalReceived sets field value
func (o *StxBalance) SetTotalReceived(v string) {
	o.TotalReceived = v
}

// GetTotalFeesSent returns the TotalFeesSent field value
func (o *StxBalance) GetTotalFeesSent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TotalFeesSent
}

// GetTotalFeesSentOk returns a tuple with the TotalFeesSent field value
// and a boolean to check if the value has been set.
func (o *StxBalance) GetTotalFeesSentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalFeesSent, true
}

// SetTotalFeesSent sets field value
func (o *StxBalance) SetTotalFeesSent(v string) {
	o.TotalFeesSent = v
}

// GetTotalMinerRewardsReceived returns the TotalMinerRewardsReceived field value
func (o *StxBalance) GetTotalMinerRewardsReceived() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TotalMinerRewardsReceived
}

// GetTotalMinerRewardsReceivedOk returns a tuple with the TotalMinerRewardsReceived field value
// and a boolean to check if the value has been set.
func (o *StxBalance) GetTotalMinerRewardsReceivedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalMinerRewardsReceived, true
}

// SetTotalMinerRewardsReceived sets field value
func (o *StxBalance) SetTotalMinerRewardsReceived(v string) {
	o.TotalMinerRewardsReceived = v
}

// GetLockTxId returns the LockTxId field value
func (o *StxBalance) GetLockTxId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LockTxId
}

// GetLockTxIdOk returns a tuple with the LockTxId field value
// and a boolean to check if the value has been set.
func (o *StxBalance) GetLockTxIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LockTxId, true
}

// SetLockTxId sets field value
func (o *StxBalance) SetLockTxId(v string) {
	o.LockTxId = v
}

// GetLocked returns the Locked field value
func (o *StxBalance) GetLocked() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Locked
}

// GetLockedOk returns a tuple with the Locked field value
// and a boolean to check if the value has been set.
func (o *StxBalance) GetLockedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Locked, true
}

// SetLocked sets field value
func (o *StxBalance) SetLocked(v string) {
	o.Locked = v
}

// GetLockHeight returns the LockHeight field value
func (o *StxBalance) GetLockHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.LockHeight
}

// GetLockHeightOk returns a tuple with the LockHeight field value
// and a boolean to check if the value has been set.
func (o *StxBalance) GetLockHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LockHeight, true
}

// SetLockHeight sets field value
func (o *StxBalance) SetLockHeight(v int32) {
	o.LockHeight = v
}

// GetBurnchainLockHeight returns the BurnchainLockHeight field value
func (o *StxBalance) GetBurnchainLockHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BurnchainLockHeight
}

// GetBurnchainLockHeightOk returns a tuple with the BurnchainLockHeight field value
// and a boolean to check if the value has been set.
func (o *StxBalance) GetBurnchainLockHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BurnchainLockHeight, true
}

// SetBurnchainLockHeight sets field value
func (o *StxBalance) SetBurnchainLockHeight(v int32) {
	o.BurnchainLockHeight = v
}

// GetBurnchainUnlockHeight returns the BurnchainUnlockHeight field value
func (o *StxBalance) GetBurnchainUnlockHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BurnchainUnlockHeight
}

// GetBurnchainUnlockHeightOk returns a tuple with the BurnchainUnlockHeight field value
// and a boolean to check if the value has been set.
func (o *StxBalance) GetBurnchainUnlockHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BurnchainUnlockHeight, true
}

// SetBurnchainUnlockHeight sets field value
func (o *StxBalance) SetBurnchainUnlockHeight(v int32) {
	o.BurnchainUnlockHeight = v
}

func (o StxBalance) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StxBalance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["balance"] = o.Balance
	toSerialize["total_sent"] = o.TotalSent
	toSerialize["total_received"] = o.TotalReceived
	toSerialize["total_fees_sent"] = o.TotalFeesSent
	toSerialize["total_miner_rewards_received"] = o.TotalMinerRewardsReceived
	toSerialize["lock_tx_id"] = o.LockTxId
	toSerialize["locked"] = o.Locked
	toSerialize["lock_height"] = o.LockHeight
	toSerialize["burnchain_lock_height"] = o.BurnchainLockHeight
	toSerialize["burnchain_unlock_height"] = o.BurnchainUnlockHeight
	return toSerialize, nil
}

func (o *StxBalance) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"balance",
		"total_sent",
		"total_received",
		"total_fees_sent",
		"total_miner_rewards_received",
		"lock_tx_id",
		"locked",
		"lock_height",
		"burnchain_lock_height",
		"burnchain_unlock_height",
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

	varStxBalance := _StxBalance{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStxBalance)

	if err != nil {
		return err
	}

	*o = StxBalance(varStxBalance)

	return err
}

type NullableStxBalance struct {
	value *StxBalance
	isSet bool
}

func (v NullableStxBalance) Get() *StxBalance {
	return v.value
}

func (v *NullableStxBalance) Set(val *StxBalance) {
	v.value = val
	v.isSet = true
}

func (v NullableStxBalance) IsSet() bool {
	return v.isSet
}

func (v *NullableStxBalance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStxBalance(val *StxBalance) *NullableStxBalance {
	return &NullableStxBalance{value: val, isSet: true}
}

func (v NullableStxBalance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStxBalance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
