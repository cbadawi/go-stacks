package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the RosettaPeers type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RosettaPeers{}

// RosettaPeers A Peer is a representation of a node's peer.
type RosettaPeers struct {
	// peer id
	PeerId string `json:"peer_id"`
	// meta data
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

type _RosettaPeers RosettaPeers

// NewRosettaPeers instantiates a new RosettaPeers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRosettaPeers(peerId string) *RosettaPeers {
	this := RosettaPeers{}
	this.PeerId = peerId
	return &this
}

// NewRosettaPeersWithDefaults instantiates a new RosettaPeers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRosettaPeersWithDefaults() *RosettaPeers {
	this := RosettaPeers{}
	return &this
}

// GetPeerId returns the PeerId field value
func (o *RosettaPeers) GetPeerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PeerId
}

// GetPeerIdOk returns a tuple with the PeerId field value
// and a boolean to check if the value has been set.
func (o *RosettaPeers) GetPeerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PeerId, true
}

// SetPeerId sets field value
func (o *RosettaPeers) SetPeerId(v string) {
	o.PeerId = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *RosettaPeers) GetMetadata() map[string]interface{} {
	if o == nil || utils.IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaPeers) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || utils.IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *RosettaPeers) HasMetadata() bool {
	if o != nil && !utils.IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *RosettaPeers) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o RosettaPeers) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RosettaPeers) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["peer_id"] = o.PeerId
	if !utils.IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

func (o *RosettaPeers) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"peer_id",
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

	varRosettaPeers := _RosettaPeers{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRosettaPeers)

	if err != nil {
		return err
	}

	*o = RosettaPeers(varRosettaPeers)

	return err
}

type NullableRosettaPeers struct {
	value *RosettaPeers
	isSet bool
}

func (v NullableRosettaPeers) Get() *RosettaPeers {
	return v.value
}

func (v *NullableRosettaPeers) Set(val *RosettaPeers) {
	v.value = val
	v.isSet = true
}

func (v NullableRosettaPeers) IsSet() bool {
	return v.isSet
}

func (v *NullableRosettaPeers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRosettaPeers(val *RosettaPeers) *NullableRosettaPeers {
	return &NullableRosettaPeers{value: val, isSet: true}
}

func (v NullableRosettaPeers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRosettaPeers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
