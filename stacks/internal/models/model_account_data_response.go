package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the AccountDataResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AccountDataResponse{}

// AccountDataResponse GET request for account data
type AccountDataResponse struct {
	Balance      string `json:"balance"`
	Locked       string `json:"locked"`
	UnlockHeight int32  `json:"unlock_height"`
	Nonce        int32  `json:"nonce"`
	BalanceProof string `json:"balance_proof"`
	NonceProof   string `json:"nonce_proof"`
}

type _AccountDataResponse AccountDataResponse

// NewAccountDataResponse instantiates a new AccountDataResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountDataResponse(balance string, locked string, unlockHeight int32, nonce int32, balanceProof string, nonceProof string) *AccountDataResponse {
	this := AccountDataResponse{}
	this.Balance = balance
	this.Locked = locked
	this.UnlockHeight = unlockHeight
	this.Nonce = nonce
	this.BalanceProof = balanceProof
	this.NonceProof = nonceProof
	return &this
}

// NewAccountDataResponseWithDefaults instantiates a new AccountDataResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountDataResponseWithDefaults() *AccountDataResponse {
	this := AccountDataResponse{}
	return &this
}

// GetBalance returns the Balance field value
func (o *AccountDataResponse) GetBalance() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value
// and a boolean to check if the value has been set.
func (o *AccountDataResponse) GetBalanceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Balance, true
}

// SetBalance sets field value
func (o *AccountDataResponse) SetBalance(v string) {
	o.Balance = v
}

// GetLocked returns the Locked field value
func (o *AccountDataResponse) GetLocked() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Locked
}

// GetLockedOk returns a tuple with the Locked field value
// and a boolean to check if the value has been set.
func (o *AccountDataResponse) GetLockedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Locked, true
}

// SetLocked sets field value
func (o *AccountDataResponse) SetLocked(v string) {
	o.Locked = v
}

// GetUnlockHeight returns the UnlockHeight field value
func (o *AccountDataResponse) GetUnlockHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UnlockHeight
}

// GetUnlockHeightOk returns a tuple with the UnlockHeight field value
// and a boolean to check if the value has been set.
func (o *AccountDataResponse) GetUnlockHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnlockHeight, true
}

// SetUnlockHeight sets field value
func (o *AccountDataResponse) SetUnlockHeight(v int32) {
	o.UnlockHeight = v
}

// GetNonce returns the Nonce field value
func (o *AccountDataResponse) GetNonce() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value
// and a boolean to check if the value has been set.
func (o *AccountDataResponse) GetNonceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nonce, true
}

// SetNonce sets field value
func (o *AccountDataResponse) SetNonce(v int32) {
	o.Nonce = v
}

// GetBalanceProof returns the BalanceProof field value
func (o *AccountDataResponse) GetBalanceProof() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BalanceProof
}

// GetBalanceProofOk returns a tuple with the BalanceProof field value
// and a boolean to check if the value has been set.
func (o *AccountDataResponse) GetBalanceProofOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BalanceProof, true
}

// SetBalanceProof sets field value
func (o *AccountDataResponse) SetBalanceProof(v string) {
	o.BalanceProof = v
}

// GetNonceProof returns the NonceProof field value
func (o *AccountDataResponse) GetNonceProof() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NonceProof
}

// GetNonceProofOk returns a tuple with the NonceProof field value
// and a boolean to check if the value has been set.
func (o *AccountDataResponse) GetNonceProofOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NonceProof, true
}

// SetNonceProof sets field value
func (o *AccountDataResponse) SetNonceProof(v string) {
	o.NonceProof = v
}

func (o AccountDataResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountDataResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["balance"] = o.Balance
	toSerialize["locked"] = o.Locked
	toSerialize["unlock_height"] = o.UnlockHeight
	toSerialize["nonce"] = o.Nonce
	toSerialize["balance_proof"] = o.BalanceProof
	toSerialize["nonce_proof"] = o.NonceProof
	return toSerialize, nil
}

func (o *AccountDataResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"balance",
		"locked",
		"unlock_height",
		"nonce",
		"balance_proof",
		"nonce_proof",
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

	varAccountDataResponse := _AccountDataResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAccountDataResponse)

	if err != nil {
		return err
	}

	*o = AccountDataResponse(varAccountDataResponse)

	return err
}

type NullableAccountDataResponse struct {
	value *AccountDataResponse
	isSet bool
}

func (v NullableAccountDataResponse) Get() *AccountDataResponse {
	return v.value
}

func (v *NullableAccountDataResponse) Set(val *AccountDataResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountDataResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountDataResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountDataResponse(val *AccountDataResponse) *NullableAccountDataResponse {
	return &NullableAccountDataResponse{value: val, isSet: true}
}

func (v NullableAccountDataResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountDataResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
