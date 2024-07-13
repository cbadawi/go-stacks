package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the ServerStatusResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ServerStatusResponse{}

// ServerStatusResponse GET blockchain API status
type ServerStatusResponse struct {
	// the server version that is currently running
	ServerVersion *string `json:"server_version,omitempty"`
	// the current server status
	Status            string              `json:"status"`
	PoxV1UnlockHeight utils.NullableInt32 `json:"pox_v1_unlock_height,omitempty"`
	PoxV2UnlockHeight utils.NullableInt32 `json:"pox_v2_unlock_height,omitempty"`
	PoxV3UnlockHeight utils.NullableInt32 `json:"pox_v3_unlock_height,omitempty"`
	ChainTip          *ChainTip           `json:"chain_tip,omitempty"`
}

type _ServerStatusResponse ServerStatusResponse

// NewServerStatusResponse instantiates a new ServerStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerStatusResponse(status string) *ServerStatusResponse {
	this := ServerStatusResponse{}
	this.Status = status
	return &this
}

// NewServerStatusResponseWithDefaults instantiates a new ServerStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerStatusResponseWithDefaults() *ServerStatusResponse {
	this := ServerStatusResponse{}
	return &this
}

// GetServerVersion returns the ServerVersion field value if set, zero value otherwise.
func (o *ServerStatusResponse) GetServerVersion() string {
	if o == nil || utils.IsNil(o.ServerVersion) {
		var ret string
		return ret
	}
	return *o.ServerVersion
}

// GetServerVersionOk returns a tuple with the ServerVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerStatusResponse) GetServerVersionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ServerVersion) {
		return nil, false
	}
	return o.ServerVersion, true
}

// HasServerVersion returns a boolean if a field has been set.
func (o *ServerStatusResponse) HasServerVersion() bool {
	if o != nil && !utils.IsNil(o.ServerVersion) {
		return true
	}

	return false
}

// SetServerVersion gets a reference to the given string and assigns it to the ServerVersion field.
func (o *ServerStatusResponse) SetServerVersion(v string) {
	o.ServerVersion = &v
}

// GetStatus returns the Status field value
func (o *ServerStatusResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ServerStatusResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ServerStatusResponse) SetStatus(v string) {
	o.Status = v
}

// GetPoxV1UnlockHeight returns the PoxV1UnlockHeight field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServerStatusResponse) GetPoxV1UnlockHeight() int32 {
	if o == nil || utils.IsNil(o.PoxV1UnlockHeight.Get()) {
		var ret int32
		return ret
	}
	return *o.PoxV1UnlockHeight.Get()
}

// GetPoxV1UnlockHeightOk returns a tuple with the PoxV1UnlockHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServerStatusResponse) GetPoxV1UnlockHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PoxV1UnlockHeight.Get(), o.PoxV1UnlockHeight.IsSet()
}

// HasPoxV1UnlockHeight returns a boolean if a field has been set.
func (o *ServerStatusResponse) HasPoxV1UnlockHeight() bool {
	if o != nil && o.PoxV1UnlockHeight.IsSet() {
		return true
	}

	return false
}

// SetPoxV1UnlockHeight gets a reference to the given NullableInt32 and assigns it to the PoxV1UnlockHeight field.
func (o *ServerStatusResponse) SetPoxV1UnlockHeight(v int32) {
	o.PoxV1UnlockHeight.Set(&v)
}

// SetPoxV1UnlockHeightNil sets the value for PoxV1UnlockHeight to be an explicit nil
func (o *ServerStatusResponse) SetPoxV1UnlockHeightNil() {
	o.PoxV1UnlockHeight.Set(nil)
}

// UnsetPoxV1UnlockHeight ensures that no value is present for PoxV1UnlockHeight, not even an explicit nil
func (o *ServerStatusResponse) UnsetPoxV1UnlockHeight() {
	o.PoxV1UnlockHeight.Unset()
}

// GetPoxV2UnlockHeight returns the PoxV2UnlockHeight field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServerStatusResponse) GetPoxV2UnlockHeight() int32 {
	if o == nil || utils.IsNil(o.PoxV2UnlockHeight.Get()) {
		var ret int32
		return ret
	}
	return *o.PoxV2UnlockHeight.Get()
}

// GetPoxV2UnlockHeightOk returns a tuple with the PoxV2UnlockHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServerStatusResponse) GetPoxV2UnlockHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PoxV2UnlockHeight.Get(), o.PoxV2UnlockHeight.IsSet()
}

// HasPoxV2UnlockHeight returns a boolean if a field has been set.
func (o *ServerStatusResponse) HasPoxV2UnlockHeight() bool {
	if o != nil && o.PoxV2UnlockHeight.IsSet() {
		return true
	}

	return false
}

// SetPoxV2UnlockHeight gets a reference to the given NullableInt32 and assigns it to the PoxV2UnlockHeight field.
func (o *ServerStatusResponse) SetPoxV2UnlockHeight(v int32) {
	o.PoxV2UnlockHeight.Set(&v)
}

// SetPoxV2UnlockHeightNil sets the value for PoxV2UnlockHeight to be an explicit nil
func (o *ServerStatusResponse) SetPoxV2UnlockHeightNil() {
	o.PoxV2UnlockHeight.Set(nil)
}

// UnsetPoxV2UnlockHeight ensures that no value is present for PoxV2UnlockHeight, not even an explicit nil
func (o *ServerStatusResponse) UnsetPoxV2UnlockHeight() {
	o.PoxV2UnlockHeight.Unset()
}

// GetPoxV3UnlockHeight returns the PoxV3UnlockHeight field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServerStatusResponse) GetPoxV3UnlockHeight() int32 {
	if o == nil || utils.IsNil(o.PoxV3UnlockHeight.Get()) {
		var ret int32
		return ret
	}
	return *o.PoxV3UnlockHeight.Get()
}

// GetPoxV3UnlockHeightOk returns a tuple with the PoxV3UnlockHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServerStatusResponse) GetPoxV3UnlockHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PoxV3UnlockHeight.Get(), o.PoxV3UnlockHeight.IsSet()
}

// HasPoxV3UnlockHeight returns a boolean if a field has been set.
func (o *ServerStatusResponse) HasPoxV3UnlockHeight() bool {
	if o != nil && o.PoxV3UnlockHeight.IsSet() {
		return true
	}

	return false
}

// SetPoxV3UnlockHeight gets a reference to the given NullableInt32 and assigns it to the PoxV3UnlockHeight field.
func (o *ServerStatusResponse) SetPoxV3UnlockHeight(v int32) {
	o.PoxV3UnlockHeight.Set(&v)
}

// SetPoxV3UnlockHeightNil sets the value for PoxV3UnlockHeight to be an explicit nil
func (o *ServerStatusResponse) SetPoxV3UnlockHeightNil() {
	o.PoxV3UnlockHeight.Set(nil)
}

// UnsetPoxV3UnlockHeight ensures that no value is present for PoxV3UnlockHeight, not even an explicit nil
func (o *ServerStatusResponse) UnsetPoxV3UnlockHeight() {
	o.PoxV3UnlockHeight.Unset()
}

// GetChainTip returns the ChainTip field value if set, zero value otherwise.
func (o *ServerStatusResponse) GetChainTip() ChainTip {
	if o == nil || utils.IsNil(o.ChainTip) {
		var ret ChainTip
		return ret
	}
	return *o.ChainTip
}

// GetChainTipOk returns a tuple with the ChainTip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerStatusResponse) GetChainTipOk() (*ChainTip, bool) {
	if o == nil || utils.IsNil(o.ChainTip) {
		return nil, false
	}
	return o.ChainTip, true
}

// HasChainTip returns a boolean if a field has been set.
func (o *ServerStatusResponse) HasChainTip() bool {
	if o != nil && !utils.IsNil(o.ChainTip) {
		return true
	}

	return false
}

// SetChainTip gets a reference to the given ChainTip and assigns it to the ChainTip field.
func (o *ServerStatusResponse) SetChainTip(v ChainTip) {
	o.ChainTip = &v
}

func (o ServerStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerStatusResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.ServerVersion) {
		toSerialize["server_version"] = o.ServerVersion
	}
	toSerialize["status"] = o.Status
	if o.PoxV1UnlockHeight.IsSet() {
		toSerialize["pox_v1_unlock_height"] = o.PoxV1UnlockHeight.Get()
	}
	if o.PoxV2UnlockHeight.IsSet() {
		toSerialize["pox_v2_unlock_height"] = o.PoxV2UnlockHeight.Get()
	}
	if o.PoxV3UnlockHeight.IsSet() {
		toSerialize["pox_v3_unlock_height"] = o.PoxV3UnlockHeight.Get()
	}
	if !utils.IsNil(o.ChainTip) {
		toSerialize["chain_tip"] = o.ChainTip
	}
	return toSerialize, nil
}

func (o *ServerStatusResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
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

	varServerStatusResponse := _ServerStatusResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varServerStatusResponse)

	if err != nil {
		return err
	}

	*o = ServerStatusResponse(varServerStatusResponse)

	return err
}

type NullableServerStatusResponse struct {
	value *ServerStatusResponse
	isSet bool
}

func (v NullableServerStatusResponse) Get() *ServerStatusResponse {
	return v.value
}

func (v *NullableServerStatusResponse) Set(val *ServerStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableServerStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableServerStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerStatusResponse(val *ServerStatusResponse) *NullableServerStatusResponse {
	return &NullableServerStatusResponse{value: val, isSet: true}
}

func (v NullableServerStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
