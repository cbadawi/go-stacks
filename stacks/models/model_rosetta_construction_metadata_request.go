package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the RosettaConstructionMetadataRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RosettaConstructionMetadataRequest{}

// RosettaConstructionMetadataRequest A ConstructionMetadataRequest is utilized to get information required to construct a transaction. The Options object used to specify which metadata to return is left purposely unstructured to allow flexibility for implementers. Optionally, the request can also include an array of PublicKeys associated with the AccountIdentifiers returned in ConstructionPreprocessResponse.
type RosettaConstructionMetadataRequest struct {
	NetworkIdentifier NetworkIdentifier  `json:"network_identifier"`
	Options           RosettaOptions     `json:"options"`
	PublicKeys        []RosettaPublicKey `json:"public_keys,omitempty"`
}

type _RosettaConstructionMetadataRequest RosettaConstructionMetadataRequest

// NewRosettaConstructionMetadataRequest instantiates a new RosettaConstructionMetadataRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRosettaConstructionMetadataRequest(networkIdentifier NetworkIdentifier, options RosettaOptions) *RosettaConstructionMetadataRequest {
	this := RosettaConstructionMetadataRequest{}
	this.NetworkIdentifier = networkIdentifier
	this.Options = options
	return &this
}

// NewRosettaConstructionMetadataRequestWithDefaults instantiates a new RosettaConstructionMetadataRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRosettaConstructionMetadataRequestWithDefaults() *RosettaConstructionMetadataRequest {
	this := RosettaConstructionMetadataRequest{}
	return &this
}

// GetNetworkIdentifier returns the NetworkIdentifier field value
func (o *RosettaConstructionMetadataRequest) GetNetworkIdentifier() NetworkIdentifier {
	if o == nil {
		var ret NetworkIdentifier
		return ret
	}

	return o.NetworkIdentifier
}

// GetNetworkIdentifierOk returns a tuple with the NetworkIdentifier field value
// and a boolean to check if the value has been set.
func (o *RosettaConstructionMetadataRequest) GetNetworkIdentifierOk() (*NetworkIdentifier, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkIdentifier, true
}

// SetNetworkIdentifier sets field value
func (o *RosettaConstructionMetadataRequest) SetNetworkIdentifier(v NetworkIdentifier) {
	o.NetworkIdentifier = v
}

// GetOptions returns the Options field value
func (o *RosettaConstructionMetadataRequest) GetOptions() RosettaOptions {
	if o == nil {
		var ret RosettaOptions
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *RosettaConstructionMetadataRequest) GetOptionsOk() (*RosettaOptions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Options, true
}

// SetOptions sets field value
func (o *RosettaConstructionMetadataRequest) SetOptions(v RosettaOptions) {
	o.Options = v
}

// GetPublicKeys returns the PublicKeys field value if set, zero value otherwise.
func (o *RosettaConstructionMetadataRequest) GetPublicKeys() []RosettaPublicKey {
	if o == nil || utils.IsNil(o.PublicKeys) {
		var ret []RosettaPublicKey
		return ret
	}
	return o.PublicKeys
}

// GetPublicKeysOk returns a tuple with the PublicKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaConstructionMetadataRequest) GetPublicKeysOk() ([]RosettaPublicKey, bool) {
	if o == nil || utils.IsNil(o.PublicKeys) {
		return nil, false
	}
	return o.PublicKeys, true
}

// HasPublicKeys returns a boolean if a field has been set.
func (o *RosettaConstructionMetadataRequest) HasPublicKeys() bool {
	if o != nil && !utils.IsNil(o.PublicKeys) {
		return true
	}

	return false
}

// SetPublicKeys gets a reference to the given []RosettaPublicKey and assigns it to the PublicKeys field.
func (o *RosettaConstructionMetadataRequest) SetPublicKeys(v []RosettaPublicKey) {
	o.PublicKeys = v
}

func (o RosettaConstructionMetadataRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RosettaConstructionMetadataRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["network_identifier"] = o.NetworkIdentifier
	toSerialize["options"] = o.Options
	if !utils.IsNil(o.PublicKeys) {
		toSerialize["public_keys"] = o.PublicKeys
	}
	return toSerialize, nil
}

func (o *RosettaConstructionMetadataRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"network_identifier",
		"options",
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

	varRosettaConstructionMetadataRequest := _RosettaConstructionMetadataRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRosettaConstructionMetadataRequest)

	if err != nil {
		return err
	}

	*o = RosettaConstructionMetadataRequest(varRosettaConstructionMetadataRequest)

	return err
}

type NullableRosettaConstructionMetadataRequest struct {
	value *RosettaConstructionMetadataRequest
	isSet bool
}

func (v NullableRosettaConstructionMetadataRequest) Get() *RosettaConstructionMetadataRequest {
	return v.value
}

func (v *NullableRosettaConstructionMetadataRequest) Set(val *RosettaConstructionMetadataRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRosettaConstructionMetadataRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRosettaConstructionMetadataRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRosettaConstructionMetadataRequest(val *RosettaConstructionMetadataRequest) *NullableRosettaConstructionMetadataRequest {
	return &NullableRosettaConstructionMetadataRequest{value: val, isSet: true}
}

func (v NullableRosettaConstructionMetadataRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRosettaConstructionMetadataRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
