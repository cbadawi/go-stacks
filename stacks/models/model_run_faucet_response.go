package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the RunFaucetResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RunFaucetResponse{}

// RunFaucetResponse POST request that initiates a transfer of tokens to a specified testnet address
type RunFaucetResponse struct {
	// Indicates if the faucet call was successful
	Success bool `json:"success"`
	// The transaction ID for the faucet call
	TxId *string `json:"txId,omitempty"`
	// Raw transaction in hex string representation
	TxRaw *string `json:"txRaw,omitempty"`
}

type _RunFaucetResponse RunFaucetResponse

// NewRunFaucetResponse instantiates a new RunFaucetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunFaucetResponse(success bool) *RunFaucetResponse {
	this := RunFaucetResponse{}
	this.Success = success
	return &this
}

// NewRunFaucetResponseWithDefaults instantiates a new RunFaucetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunFaucetResponseWithDefaults() *RunFaucetResponse {
	this := RunFaucetResponse{}
	return &this
}

// GetSuccess returns the Success field value
func (o *RunFaucetResponse) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *RunFaucetResponse) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *RunFaucetResponse) SetSuccess(v bool) {
	o.Success = v
}

// GetTxId returns the TxId field value if set, zero value otherwise.
func (o *RunFaucetResponse) GetTxId() string {
	if o == nil || utils.IsNil(o.TxId) {
		var ret string
		return ret
	}
	return *o.TxId
}

// GetTxIdOk returns a tuple with the TxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunFaucetResponse) GetTxIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TxId) {
		return nil, false
	}
	return o.TxId, true
}

// HasTxId returns a boolean if a field has been set.
func (o *RunFaucetResponse) HasTxId() bool {
	if o != nil && !utils.IsNil(o.TxId) {
		return true
	}

	return false
}

// SetTxId gets a reference to the given string and assigns it to the TxId field.
func (o *RunFaucetResponse) SetTxId(v string) {
	o.TxId = &v
}

// GetTxRaw returns the TxRaw field value if set, zero value otherwise.
func (o *RunFaucetResponse) GetTxRaw() string {
	if o == nil || utils.IsNil(o.TxRaw) {
		var ret string
		return ret
	}
	return *o.TxRaw
}

// GetTxRawOk returns a tuple with the TxRaw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunFaucetResponse) GetTxRawOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TxRaw) {
		return nil, false
	}
	return o.TxRaw, true
}

// HasTxRaw returns a boolean if a field has been set.
func (o *RunFaucetResponse) HasTxRaw() bool {
	if o != nil && !utils.IsNil(o.TxRaw) {
		return true
	}

	return false
}

// SetTxRaw gets a reference to the given string and assigns it to the TxRaw field.
func (o *RunFaucetResponse) SetTxRaw(v string) {
	o.TxRaw = &v
}

func (o RunFaucetResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RunFaucetResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["success"] = o.Success
	if !utils.IsNil(o.TxId) {
		toSerialize["txId"] = o.TxId
	}
	if !utils.IsNil(o.TxRaw) {
		toSerialize["txRaw"] = o.TxRaw
	}
	return toSerialize, nil
}

func (o *RunFaucetResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"success",
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

	varRunFaucetResponse := _RunFaucetResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRunFaucetResponse)

	if err != nil {
		return err
	}

	*o = RunFaucetResponse(varRunFaucetResponse)

	return err
}

type NullableRunFaucetResponse struct {
	value *RunFaucetResponse
	isSet bool
}

func (v NullableRunFaucetResponse) Get() *RunFaucetResponse {
	return v.value
}

func (v *NullableRunFaucetResponse) Set(val *RunFaucetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRunFaucetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRunFaucetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunFaucetResponse(val *RunFaucetResponse) *NullableRunFaucetResponse {
	return &NullableRunFaucetResponse{value: val, isSet: true}
}

func (v NullableRunFaucetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunFaucetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
