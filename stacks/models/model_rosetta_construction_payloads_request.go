package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the RosettaConstructionPayloadsRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RosettaConstructionPayloadsRequest{}

// RosettaConstructionPayloadsRequest ConstructionPayloadsRequest is the request to /construction/payloads. It contains the network, a slice of operations, and arbitrary metadata that was returned by the call to /construction/metadata. Optionally, the request can also include an array of PublicKeys associated with the AccountIdentifiers returned in ConstructionPreprocessResponse.
type RosettaConstructionPayloadsRequest struct {
	NetworkIdentifier NetworkIdentifier                            `json:"network_identifier"`
	Operations        []RosettaOperation                           `json:"operations"`
	PublicKeys        []RosettaPublicKey                           `json:"public_keys,omitempty"`
	Metadata          *RosettaConstructionMetadataResponseMetadata `json:"metadata,omitempty"`
}

type _RosettaConstructionPayloadsRequest RosettaConstructionPayloadsRequest

// NewRosettaConstructionPayloadsRequest instantiates a new RosettaConstructionPayloadsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRosettaConstructionPayloadsRequest(networkIdentifier NetworkIdentifier, operations []RosettaOperation) *RosettaConstructionPayloadsRequest {
	this := RosettaConstructionPayloadsRequest{}
	this.NetworkIdentifier = networkIdentifier
	this.Operations = operations
	return &this
}

// NewRosettaConstructionPayloadsRequestWithDefaults instantiates a new RosettaConstructionPayloadsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRosettaConstructionPayloadsRequestWithDefaults() *RosettaConstructionPayloadsRequest {
	this := RosettaConstructionPayloadsRequest{}
	return &this
}

// GetNetworkIdentifier returns the NetworkIdentifier field value
func (o *RosettaConstructionPayloadsRequest) GetNetworkIdentifier() NetworkIdentifier {
	if o == nil {
		var ret NetworkIdentifier
		return ret
	}

	return o.NetworkIdentifier
}

// GetNetworkIdentifierOk returns a tuple with the NetworkIdentifier field value
// and a boolean to check if the value has been set.
func (o *RosettaConstructionPayloadsRequest) GetNetworkIdentifierOk() (*NetworkIdentifier, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkIdentifier, true
}

// SetNetworkIdentifier sets field value
func (o *RosettaConstructionPayloadsRequest) SetNetworkIdentifier(v NetworkIdentifier) {
	o.NetworkIdentifier = v
}

// GetOperations returns the Operations field value
func (o *RosettaConstructionPayloadsRequest) GetOperations() []RosettaOperation {
	if o == nil {
		var ret []RosettaOperation
		return ret
	}

	return o.Operations
}

// GetOperationsOk returns a tuple with the Operations field value
// and a boolean to check if the value has been set.
func (o *RosettaConstructionPayloadsRequest) GetOperationsOk() ([]RosettaOperation, bool) {
	if o == nil {
		return nil, false
	}
	return o.Operations, true
}

// SetOperations sets field value
func (o *RosettaConstructionPayloadsRequest) SetOperations(v []RosettaOperation) {
	o.Operations = v
}

// GetPublicKeys returns the PublicKeys field value if set, zero value otherwise.
func (o *RosettaConstructionPayloadsRequest) GetPublicKeys() []RosettaPublicKey {
	if o == nil || utils.IsNil(o.PublicKeys) {
		var ret []RosettaPublicKey
		return ret
	}
	return o.PublicKeys
}

// GetPublicKeysOk returns a tuple with the PublicKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaConstructionPayloadsRequest) GetPublicKeysOk() ([]RosettaPublicKey, bool) {
	if o == nil || utils.IsNil(o.PublicKeys) {
		return nil, false
	}
	return o.PublicKeys, true
}

// HasPublicKeys returns a boolean if a field has been set.
func (o *RosettaConstructionPayloadsRequest) HasPublicKeys() bool {
	if o != nil && !utils.IsNil(o.PublicKeys) {
		return true
	}

	return false
}

// SetPublicKeys gets a reference to the given []RosettaPublicKey and assigns it to the PublicKeys field.
func (o *RosettaConstructionPayloadsRequest) SetPublicKeys(v []RosettaPublicKey) {
	o.PublicKeys = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *RosettaConstructionPayloadsRequest) GetMetadata() RosettaConstructionMetadataResponseMetadata {
	if o == nil || utils.IsNil(o.Metadata) {
		var ret RosettaConstructionMetadataResponseMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaConstructionPayloadsRequest) GetMetadataOk() (*RosettaConstructionMetadataResponseMetadata, bool) {
	if o == nil || utils.IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *RosettaConstructionPayloadsRequest) HasMetadata() bool {
	if o != nil && !utils.IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given RosettaConstructionMetadataResponseMetadata and assigns it to the Metadata field.
func (o *RosettaConstructionPayloadsRequest) SetMetadata(v RosettaConstructionMetadataResponseMetadata) {
	o.Metadata = &v
}

func (o RosettaConstructionPayloadsRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RosettaConstructionPayloadsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["network_identifier"] = o.NetworkIdentifier
	toSerialize["operations"] = o.Operations
	if !utils.IsNil(o.PublicKeys) {
		toSerialize["public_keys"] = o.PublicKeys
	}
	if !utils.IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

func (o *RosettaConstructionPayloadsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"network_identifier",
		"operations",
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

	varRosettaConstructionPayloadsRequest := _RosettaConstructionPayloadsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRosettaConstructionPayloadsRequest)

	if err != nil {
		return err
	}

	*o = RosettaConstructionPayloadsRequest(varRosettaConstructionPayloadsRequest)

	return err
}

type NullableRosettaConstructionPayloadsRequest struct {
	value *RosettaConstructionPayloadsRequest
	isSet bool
}

func (v NullableRosettaConstructionPayloadsRequest) Get() *RosettaConstructionPayloadsRequest {
	return v.value
}

func (v *NullableRosettaConstructionPayloadsRequest) Set(val *RosettaConstructionPayloadsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRosettaConstructionPayloadsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRosettaConstructionPayloadsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRosettaConstructionPayloadsRequest(val *RosettaConstructionPayloadsRequest) *NullableRosettaConstructionPayloadsRequest {
	return &NullableRosettaConstructionPayloadsRequest{value: val, isSet: true}
}

func (v NullableRosettaConstructionPayloadsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRosettaConstructionPayloadsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
