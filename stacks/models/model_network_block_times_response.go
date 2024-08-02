package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the NetworkBlockTimesResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &NetworkBlockTimesResponse{}

// NetworkBlockTimesResponse GET request that returns network target block times
type NetworkBlockTimesResponse struct {
	Mainnet TargetBlockTime `json:"mainnet"`
	Testnet TargetBlockTime `json:"testnet"`
}

type _NetworkBlockTimesResponse NetworkBlockTimesResponse

// NewNetworkBlockTimesResponse instantiates a new NetworkBlockTimesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkBlockTimesResponse(mainnet TargetBlockTime, testnet TargetBlockTime) *NetworkBlockTimesResponse {
	this := NetworkBlockTimesResponse{}
	this.Mainnet = mainnet
	this.Testnet = testnet
	return &this
}

// NewNetworkBlockTimesResponseWithDefaults instantiates a new NetworkBlockTimesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkBlockTimesResponseWithDefaults() *NetworkBlockTimesResponse {
	this := NetworkBlockTimesResponse{}
	return &this
}

// GetMainnet returns the Mainnet field value
func (o *NetworkBlockTimesResponse) GetMainnet() TargetBlockTime {
	if o == nil {
		var ret TargetBlockTime
		return ret
	}

	return o.Mainnet
}

// GetMainnetOk returns a tuple with the Mainnet field value
// and a boolean to check if the value has been set.
func (o *NetworkBlockTimesResponse) GetMainnetOk() (*TargetBlockTime, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mainnet, true
}

// SetMainnet sets field value
func (o *NetworkBlockTimesResponse) SetMainnet(v TargetBlockTime) {
	o.Mainnet = v
}

// GetTestnet returns the Testnet field value
func (o *NetworkBlockTimesResponse) GetTestnet() TargetBlockTime {
	if o == nil {
		var ret TargetBlockTime
		return ret
	}

	return o.Testnet
}

// GetTestnetOk returns a tuple with the Testnet field value
// and a boolean to check if the value has been set.
func (o *NetworkBlockTimesResponse) GetTestnetOk() (*TargetBlockTime, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Testnet, true
}

// SetTestnet sets field value
func (o *NetworkBlockTimesResponse) SetTestnet(v TargetBlockTime) {
	o.Testnet = v
}

func (o NetworkBlockTimesResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkBlockTimesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mainnet"] = o.Mainnet
	toSerialize["testnet"] = o.Testnet
	return toSerialize, nil
}

func (o *NetworkBlockTimesResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"mainnet",
		"testnet",
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

	varNetworkBlockTimesResponse := _NetworkBlockTimesResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNetworkBlockTimesResponse)

	if err != nil {
		return err
	}

	*o = NetworkBlockTimesResponse(varNetworkBlockTimesResponse)

	return err
}

type NullableNetworkBlockTimesResponse struct {
	value *NetworkBlockTimesResponse
	isSet bool
}

func (v NullableNetworkBlockTimesResponse) Get() *NetworkBlockTimesResponse {
	return v.value
}

func (v *NullableNetworkBlockTimesResponse) Set(val *NetworkBlockTimesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkBlockTimesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkBlockTimesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkBlockTimesResponse(val *NetworkBlockTimesResponse) *NullableNetworkBlockTimesResponse {
	return &NullableNetworkBlockTimesResponse{value: val, isSet: true}
}

func (v NullableNetworkBlockTimesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkBlockTimesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
