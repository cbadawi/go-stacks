package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the NetworkIdentifier type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &NetworkIdentifier{}

// NetworkIdentifier The network_identifier specifies which network a particular object is associated with.
type NetworkIdentifier struct {
	// Blockchain name
	Blockchain string `json:"blockchain"`
	// If a blockchain has a specific chain-id or network identifier, it should go in this field. It is up to the client to determine which network-specific identifier is mainnet or testnet.
	Network              string                                 `json:"network"`
	SubNetworkIdentifier *NetworkIdentifierSubNetworkIdentifier `json:"sub_network_identifier,omitempty"`
}

type _NetworkIdentifier NetworkIdentifier

// NewNetworkIdentifier instantiates a new NetworkIdentifier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkIdentifier(blockchain string, network string) *NetworkIdentifier {
	this := NetworkIdentifier{}
	this.Blockchain = blockchain
	this.Network = network
	return &this
}

// NewNetworkIdentifierWithDefaults instantiates a new NetworkIdentifier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkIdentifierWithDefaults() *NetworkIdentifier {
	this := NetworkIdentifier{}
	return &this
}

// GetBlockchain returns the Blockchain field value
func (o *NetworkIdentifier) GetBlockchain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Blockchain
}

// GetBlockchainOk returns a tuple with the Blockchain field value
// and a boolean to check if the value has been set.
func (o *NetworkIdentifier) GetBlockchainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Blockchain, true
}

// SetBlockchain sets field value
func (o *NetworkIdentifier) SetBlockchain(v string) {
	o.Blockchain = v
}

// GetNetwork returns the Network field value
func (o *NetworkIdentifier) GetNetwork() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Network
}

// GetNetworkOk returns a tuple with the Network field value
// and a boolean to check if the value has been set.
func (o *NetworkIdentifier) GetNetworkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Network, true
}

// SetNetwork sets field value
func (o *NetworkIdentifier) SetNetwork(v string) {
	o.Network = v
}

// GetSubNetworkIdentifier returns the SubNetworkIdentifier field value if set, zero value otherwise.
func (o *NetworkIdentifier) GetSubNetworkIdentifier() NetworkIdentifierSubNetworkIdentifier {
	if o == nil || utils.IsNil(o.SubNetworkIdentifier) {
		var ret NetworkIdentifierSubNetworkIdentifier
		return ret
	}
	return *o.SubNetworkIdentifier
}

// GetSubNetworkIdentifierOk returns a tuple with the SubNetworkIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkIdentifier) GetSubNetworkIdentifierOk() (*NetworkIdentifierSubNetworkIdentifier, bool) {
	if o == nil || utils.IsNil(o.SubNetworkIdentifier) {
		return nil, false
	}
	return o.SubNetworkIdentifier, true
}

// HasSubNetworkIdentifier returns a boolean if a field has been set.
func (o *NetworkIdentifier) HasSubNetworkIdentifier() bool {
	if o != nil && !utils.IsNil(o.SubNetworkIdentifier) {
		return true
	}

	return false
}

// SetSubNetworkIdentifier gets a reference to the given NetworkIdentifierSubNetworkIdentifier and assigns it to the SubNetworkIdentifier field.
func (o *NetworkIdentifier) SetSubNetworkIdentifier(v NetworkIdentifierSubNetworkIdentifier) {
	o.SubNetworkIdentifier = &v
}

func (o NetworkIdentifier) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkIdentifier) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["blockchain"] = o.Blockchain
	toSerialize["network"] = o.Network
	if !utils.IsNil(o.SubNetworkIdentifier) {
		toSerialize["sub_network_identifier"] = o.SubNetworkIdentifier
	}
	return toSerialize, nil
}

func (o *NetworkIdentifier) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"blockchain",
		"network",
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

	varNetworkIdentifier := _NetworkIdentifier{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNetworkIdentifier)

	if err != nil {
		return err
	}

	*o = NetworkIdentifier(varNetworkIdentifier)

	return err
}

type NullableNetworkIdentifier struct {
	value *NetworkIdentifier
	isSet bool
}

func (v NullableNetworkIdentifier) Get() *NetworkIdentifier {
	return v.value
}

func (v *NullableNetworkIdentifier) Set(val *NetworkIdentifier) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkIdentifier) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkIdentifier(val *NetworkIdentifier) *NullableNetworkIdentifier {
	return &NullableNetworkIdentifier{value: val, isSet: true}
}

func (v NullableNetworkIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
