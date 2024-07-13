package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the FtBalance type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &FtBalance{}

// FtBalance struct for FtBalance
type FtBalance struct {
	Balance       string `json:"balance"`
	TotalSent     string `json:"total_sent"`
	TotalReceived string `json:"total_received"`
}

type _FtBalance FtBalance

// NewFtBalance instantiates a new FtBalance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFtBalance(balance string, totalSent string, totalReceived string) *FtBalance {
	this := FtBalance{}
	this.Balance = balance
	this.TotalSent = totalSent
	this.TotalReceived = totalReceived
	return &this
}

// NewFtBalanceWithDefaults instantiates a new FtBalance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFtBalanceWithDefaults() *FtBalance {
	this := FtBalance{}
	return &this
}

// GetBalance returns the Balance field value
func (o *FtBalance) GetBalance() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value
// and a boolean to check if the value has been set.
func (o *FtBalance) GetBalanceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Balance, true
}

// SetBalance sets field value
func (o *FtBalance) SetBalance(v string) {
	o.Balance = v
}

// GetTotalSent returns the TotalSent field value
func (o *FtBalance) GetTotalSent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TotalSent
}

// GetTotalSentOk returns a tuple with the TotalSent field value
// and a boolean to check if the value has been set.
func (o *FtBalance) GetTotalSentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalSent, true
}

// SetTotalSent sets field value
func (o *FtBalance) SetTotalSent(v string) {
	o.TotalSent = v
}

// GetTotalReceived returns the TotalReceived field value
func (o *FtBalance) GetTotalReceived() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TotalReceived
}

// GetTotalReceivedOk returns a tuple with the TotalReceived field value
// and a boolean to check if the value has been set.
func (o *FtBalance) GetTotalReceivedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalReceived, true
}

// SetTotalReceived sets field value
func (o *FtBalance) SetTotalReceived(v string) {
	o.TotalReceived = v
}

func (o FtBalance) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FtBalance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["balance"] = o.Balance
	toSerialize["total_sent"] = o.TotalSent
	toSerialize["total_received"] = o.TotalReceived
	return toSerialize, nil
}

func (o *FtBalance) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"balance",
		"total_sent",
		"total_received",
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

	varFtBalance := _FtBalance{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFtBalance)

	if err != nil {
		return err
	}

	*o = FtBalance(varFtBalance)

	return err
}

type NullableFtBalance struct {
	value *FtBalance
	isSet bool
}

func (v NullableFtBalance) Get() *FtBalance {
	return v.value
}

func (v *NullableFtBalance) Set(val *FtBalance) {
	v.value = val
	v.isSet = true
}

func (v NullableFtBalance) IsSet() bool {
	return v.isSet
}

func (v *NullableFtBalance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFtBalance(val *FtBalance) *NullableFtBalance {
	return &NullableFtBalance{value: val, isSet: true}
}

func (v NullableFtBalance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFtBalance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
