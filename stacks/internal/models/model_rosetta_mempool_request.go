package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the RosettaMempoolRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RosettaMempoolRequest{}

// RosettaMempoolRequest Get all Transaction Identifiers in the mempool
type RosettaMempoolRequest struct {
	NetworkIdentifier NetworkIdentifier      `json:"network_identifier"`
	Metadata          map[string]interface{} `json:"metadata,omitempty"`
}

type _RosettaMempoolRequest RosettaMempoolRequest

// NewRosettaMempoolRequest instantiates a new RosettaMempoolRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRosettaMempoolRequest(networkIdentifier NetworkIdentifier) *RosettaMempoolRequest {
	this := RosettaMempoolRequest{}
	this.NetworkIdentifier = networkIdentifier
	return &this
}

// NewRosettaMempoolRequestWithDefaults instantiates a new RosettaMempoolRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRosettaMempoolRequestWithDefaults() *RosettaMempoolRequest {
	this := RosettaMempoolRequest{}
	return &this
}

// GetNetworkIdentifier returns the NetworkIdentifier field value
func (o *RosettaMempoolRequest) GetNetworkIdentifier() NetworkIdentifier {
	if o == nil {
		var ret NetworkIdentifier
		return ret
	}

	return o.NetworkIdentifier
}

// GetNetworkIdentifierOk returns a tuple with the NetworkIdentifier field value
// and a boolean to check if the value has been set.
func (o *RosettaMempoolRequest) GetNetworkIdentifierOk() (*NetworkIdentifier, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkIdentifier, true
}

// SetNetworkIdentifier sets field value
func (o *RosettaMempoolRequest) SetNetworkIdentifier(v NetworkIdentifier) {
	o.NetworkIdentifier = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *RosettaMempoolRequest) GetMetadata() map[string]interface{} {
	if o == nil || utils.IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaMempoolRequest) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || utils.IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *RosettaMempoolRequest) HasMetadata() bool {
	if o != nil && !utils.IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *RosettaMempoolRequest) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o RosettaMempoolRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RosettaMempoolRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["network_identifier"] = o.NetworkIdentifier
	if !utils.IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

func (o *RosettaMempoolRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"network_identifier",
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

	varRosettaMempoolRequest := _RosettaMempoolRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRosettaMempoolRequest)

	if err != nil {
		return err
	}

	*o = RosettaMempoolRequest(varRosettaMempoolRequest)

	return err
}

type NullableRosettaMempoolRequest struct {
	value *RosettaMempoolRequest
	isSet bool
}

func (v NullableRosettaMempoolRequest) Get() *RosettaMempoolRequest {
	return v.value
}

func (v *NullableRosettaMempoolRequest) Set(val *RosettaMempoolRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRosettaMempoolRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRosettaMempoolRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRosettaMempoolRequest(val *RosettaMempoolRequest) *NullableRosettaMempoolRequest {
	return &NullableRosettaMempoolRequest{value: val, isSet: true}
}

func (v NullableRosettaMempoolRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRosettaMempoolRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
