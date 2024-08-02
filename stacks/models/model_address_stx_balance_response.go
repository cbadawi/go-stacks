package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the AddressStxBalanceResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AddressStxBalanceResponse{}

// AddressStxBalanceResponse GET request that returns address balances
type AddressStxBalanceResponse struct {
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
	BurnchainUnlockHeight int32                       `json:"burnchain_unlock_height"`
	TokenOfferingLocked   *AddressTokenOfferingLocked `json:"token_offering_locked,omitempty"`
}

type _AddressStxBalanceResponse AddressStxBalanceResponse

// NewAddressStxBalanceResponse instantiates a new AddressStxBalanceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressStxBalanceResponse(balance string, totalSent string, totalReceived string, totalFeesSent string, totalMinerRewardsReceived string, lockTxId string, locked string, lockHeight int32, burnchainLockHeight int32, burnchainUnlockHeight int32) *AddressStxBalanceResponse {
	this := AddressStxBalanceResponse{}
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

// NewAddressStxBalanceResponseWithDefaults instantiates a new AddressStxBalanceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressStxBalanceResponseWithDefaults() *AddressStxBalanceResponse {
	this := AddressStxBalanceResponse{}
	return &this
}

// GetBalance returns the Balance field value
func (o *AddressStxBalanceResponse) GetBalance() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value
// and a boolean to check if the value has been set.
func (o *AddressStxBalanceResponse) GetBalanceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Balance, true
}

// SetBalance sets field value
func (o *AddressStxBalanceResponse) SetBalance(v string) {
	o.Balance = v
}

// GetTotalSent returns the TotalSent field value
func (o *AddressStxBalanceResponse) GetTotalSent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TotalSent
}

// GetTotalSentOk returns a tuple with the TotalSent field value
// and a boolean to check if the value has been set.
func (o *AddressStxBalanceResponse) GetTotalSentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalSent, true
}

// SetTotalSent sets field value
func (o *AddressStxBalanceResponse) SetTotalSent(v string) {
	o.TotalSent = v
}

// GetTotalReceived returns the TotalReceived field value
func (o *AddressStxBalanceResponse) GetTotalReceived() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TotalReceived
}

// GetTotalReceivedOk returns a tuple with the TotalReceived field value
// and a boolean to check if the value has been set.
func (o *AddressStxBalanceResponse) GetTotalReceivedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalReceived, true
}

// SetTotalReceived sets field value
func (o *AddressStxBalanceResponse) SetTotalReceived(v string) {
	o.TotalReceived = v
}

// GetTotalFeesSent returns the TotalFeesSent field value
func (o *AddressStxBalanceResponse) GetTotalFeesSent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TotalFeesSent
}

// GetTotalFeesSentOk returns a tuple with the TotalFeesSent field value
// and a boolean to check if the value has been set.
func (o *AddressStxBalanceResponse) GetTotalFeesSentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalFeesSent, true
}

// SetTotalFeesSent sets field value
func (o *AddressStxBalanceResponse) SetTotalFeesSent(v string) {
	o.TotalFeesSent = v
}

// GetTotalMinerRewardsReceived returns the TotalMinerRewardsReceived field value
func (o *AddressStxBalanceResponse) GetTotalMinerRewardsReceived() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TotalMinerRewardsReceived
}

// GetTotalMinerRewardsReceivedOk returns a tuple with the TotalMinerRewardsReceived field value
// and a boolean to check if the value has been set.
func (o *AddressStxBalanceResponse) GetTotalMinerRewardsReceivedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalMinerRewardsReceived, true
}

// SetTotalMinerRewardsReceived sets field value
func (o *AddressStxBalanceResponse) SetTotalMinerRewardsReceived(v string) {
	o.TotalMinerRewardsReceived = v
}

// GetLockTxId returns the LockTxId field value
func (o *AddressStxBalanceResponse) GetLockTxId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LockTxId
}

// GetLockTxIdOk returns a tuple with the LockTxId field value
// and a boolean to check if the value has been set.
func (o *AddressStxBalanceResponse) GetLockTxIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LockTxId, true
}

// SetLockTxId sets field value
func (o *AddressStxBalanceResponse) SetLockTxId(v string) {
	o.LockTxId = v
}

// GetLocked returns the Locked field value
func (o *AddressStxBalanceResponse) GetLocked() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Locked
}

// GetLockedOk returns a tuple with the Locked field value
// and a boolean to check if the value has been set.
func (o *AddressStxBalanceResponse) GetLockedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Locked, true
}

// SetLocked sets field value
func (o *AddressStxBalanceResponse) SetLocked(v string) {
	o.Locked = v
}

// GetLockHeight returns the LockHeight field value
func (o *AddressStxBalanceResponse) GetLockHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.LockHeight
}

// GetLockHeightOk returns a tuple with the LockHeight field value
// and a boolean to check if the value has been set.
func (o *AddressStxBalanceResponse) GetLockHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LockHeight, true
}

// SetLockHeight sets field value
func (o *AddressStxBalanceResponse) SetLockHeight(v int32) {
	o.LockHeight = v
}

// GetBurnchainLockHeight returns the BurnchainLockHeight field value
func (o *AddressStxBalanceResponse) GetBurnchainLockHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BurnchainLockHeight
}

// GetBurnchainLockHeightOk returns a tuple with the BurnchainLockHeight field value
// and a boolean to check if the value has been set.
func (o *AddressStxBalanceResponse) GetBurnchainLockHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BurnchainLockHeight, true
}

// SetBurnchainLockHeight sets field value
func (o *AddressStxBalanceResponse) SetBurnchainLockHeight(v int32) {
	o.BurnchainLockHeight = v
}

// GetBurnchainUnlockHeight returns the BurnchainUnlockHeight field value
func (o *AddressStxBalanceResponse) GetBurnchainUnlockHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BurnchainUnlockHeight
}

// GetBurnchainUnlockHeightOk returns a tuple with the BurnchainUnlockHeight field value
// and a boolean to check if the value has been set.
func (o *AddressStxBalanceResponse) GetBurnchainUnlockHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BurnchainUnlockHeight, true
}

// SetBurnchainUnlockHeight sets field value
func (o *AddressStxBalanceResponse) SetBurnchainUnlockHeight(v int32) {
	o.BurnchainUnlockHeight = v
}

// GetTokenOfferingLocked returns the TokenOfferingLocked field value if set, zero value otherwise.
func (o *AddressStxBalanceResponse) GetTokenOfferingLocked() AddressTokenOfferingLocked {
	if o == nil || utils.IsNil(o.TokenOfferingLocked) {
		var ret AddressTokenOfferingLocked
		return ret
	}
	return *o.TokenOfferingLocked
}

// GetTokenOfferingLockedOk returns a tuple with the TokenOfferingLocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressStxBalanceResponse) GetTokenOfferingLockedOk() (*AddressTokenOfferingLocked, bool) {
	if o == nil || utils.IsNil(o.TokenOfferingLocked) {
		return nil, false
	}
	return o.TokenOfferingLocked, true
}

// HasTokenOfferingLocked returns a boolean if a field has been set.
func (o *AddressStxBalanceResponse) HasTokenOfferingLocked() bool {
	if o != nil && !utils.IsNil(o.TokenOfferingLocked) {
		return true
	}

	return false
}

// SetTokenOfferingLocked gets a reference to the given AddressTokenOfferingLocked and assigns it to the TokenOfferingLocked field.
func (o *AddressStxBalanceResponse) SetTokenOfferingLocked(v AddressTokenOfferingLocked) {
	o.TokenOfferingLocked = &v
}

func (o AddressStxBalanceResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressStxBalanceResponse) ToMap() (map[string]interface{}, error) {
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
	if !utils.IsNil(o.TokenOfferingLocked) {
		toSerialize["token_offering_locked"] = o.TokenOfferingLocked
	}
	return toSerialize, nil
}

func (o *AddressStxBalanceResponse) UnmarshalJSON(data []byte) (err error) {
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

	varAddressStxBalanceResponse := _AddressStxBalanceResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddressStxBalanceResponse)

	if err != nil {
		return err
	}

	*o = AddressStxBalanceResponse(varAddressStxBalanceResponse)

	return err
}

type NullableAddressStxBalanceResponse struct {
	value *AddressStxBalanceResponse
	isSet bool
}

func (v NullableAddressStxBalanceResponse) Get() *AddressStxBalanceResponse {
	return v.value
}

func (v *NullableAddressStxBalanceResponse) Set(val *AddressStxBalanceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressStxBalanceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressStxBalanceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressStxBalanceResponse(val *AddressStxBalanceResponse) *NullableAddressStxBalanceResponse {
	return &NullableAddressStxBalanceResponse{value: val, isSet: true}
}

func (v NullableAddressStxBalanceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressStxBalanceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
