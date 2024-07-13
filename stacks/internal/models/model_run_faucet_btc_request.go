package models

import (
	"encoding/json"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the RunFaucetBtcRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RunFaucetBtcRequest{}

// RunFaucetBtcRequest struct for RunFaucetBtcRequest
type RunFaucetBtcRequest struct {
	// BTC testnet address
	Address *string `json:"address,omitempty"`
}

// NewRunFaucetBtcRequest instantiates a new RunFaucetBtcRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunFaucetBtcRequest() *RunFaucetBtcRequest {
	this := RunFaucetBtcRequest{}
	return &this
}

// NewRunFaucetBtcRequestWithDefaults instantiates a new RunFaucetBtcRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunFaucetBtcRequestWithDefaults() *RunFaucetBtcRequest {
	this := RunFaucetBtcRequest{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *RunFaucetBtcRequest) GetAddress() string {
	if o == nil || utils.IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunFaucetBtcRequest) GetAddressOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *RunFaucetBtcRequest) HasAddress() bool {
	if o != nil && !utils.IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *RunFaucetBtcRequest) SetAddress(v string) {
	o.Address = &v
}

func (o RunFaucetBtcRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RunFaucetBtcRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	return toSerialize, nil
}

type NullableRunFaucetBtcRequest struct {
	value *RunFaucetBtcRequest
	isSet bool
}

func (v NullableRunFaucetBtcRequest) Get() *RunFaucetBtcRequest {
	return v.value
}

func (v *NullableRunFaucetBtcRequest) Set(val *RunFaucetBtcRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRunFaucetBtcRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRunFaucetBtcRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunFaucetBtcRequest(val *RunFaucetBtcRequest) *NullableRunFaucetBtcRequest {
	return &NullableRunFaucetBtcRequest{value: val, isSet: true}
}

func (v NullableRunFaucetBtcRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunFaucetBtcRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
