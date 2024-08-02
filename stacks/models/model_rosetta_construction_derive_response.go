package models

import (
	"encoding/json"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the RosettaConstructionDeriveResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RosettaConstructionDeriveResponse{}

// RosettaConstructionDeriveResponse ConstructionDeriveResponse is returned by the /construction/derive endpoint.
type RosettaConstructionDeriveResponse struct {
	// [DEPRECATED by account_identifier in v1.4.4] Address in network-specific format.
	Address           *string                   `json:"address,omitempty"`
	AccountIdentifier *RosettaAccountIdentifier `json:"account_identifier,omitempty"`
	Metadata          map[string]interface{}    `json:"metadata,omitempty"`
}

// NewRosettaConstructionDeriveResponse instantiates a new RosettaConstructionDeriveResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRosettaConstructionDeriveResponse() *RosettaConstructionDeriveResponse {
	this := RosettaConstructionDeriveResponse{}
	return &this
}

// NewRosettaConstructionDeriveResponseWithDefaults instantiates a new RosettaConstructionDeriveResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRosettaConstructionDeriveResponseWithDefaults() *RosettaConstructionDeriveResponse {
	this := RosettaConstructionDeriveResponse{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *RosettaConstructionDeriveResponse) GetAddress() string {
	if o == nil || utils.IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaConstructionDeriveResponse) GetAddressOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *RosettaConstructionDeriveResponse) HasAddress() bool {
	if o != nil && !utils.IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *RosettaConstructionDeriveResponse) SetAddress(v string) {
	o.Address = &v
}

// GetAccountIdentifier returns the AccountIdentifier field value if set, zero value otherwise.
func (o *RosettaConstructionDeriveResponse) GetAccountIdentifier() RosettaAccountIdentifier {
	if o == nil || utils.IsNil(o.AccountIdentifier) {
		var ret RosettaAccountIdentifier
		return ret
	}
	return *o.AccountIdentifier
}

// GetAccountIdentifierOk returns a tuple with the AccountIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaConstructionDeriveResponse) GetAccountIdentifierOk() (*RosettaAccountIdentifier, bool) {
	if o == nil || utils.IsNil(o.AccountIdentifier) {
		return nil, false
	}
	return o.AccountIdentifier, true
}

// HasAccountIdentifier returns a boolean if a field has been set.
func (o *RosettaConstructionDeriveResponse) HasAccountIdentifier() bool {
	if o != nil && !utils.IsNil(o.AccountIdentifier) {
		return true
	}

	return false
}

// SetAccountIdentifier gets a reference to the given RosettaAccountIdentifier and assigns it to the AccountIdentifier field.
func (o *RosettaConstructionDeriveResponse) SetAccountIdentifier(v RosettaAccountIdentifier) {
	o.AccountIdentifier = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *RosettaConstructionDeriveResponse) GetMetadata() map[string]interface{} {
	if o == nil || utils.IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaConstructionDeriveResponse) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || utils.IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *RosettaConstructionDeriveResponse) HasMetadata() bool {
	if o != nil && !utils.IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *RosettaConstructionDeriveResponse) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o RosettaConstructionDeriveResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RosettaConstructionDeriveResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !utils.IsNil(o.AccountIdentifier) {
		toSerialize["account_identifier"] = o.AccountIdentifier
	}
	if !utils.IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullableRosettaConstructionDeriveResponse struct {
	value *RosettaConstructionDeriveResponse
	isSet bool
}

func (v NullableRosettaConstructionDeriveResponse) Get() *RosettaConstructionDeriveResponse {
	return v.value
}

func (v *NullableRosettaConstructionDeriveResponse) Set(val *RosettaConstructionDeriveResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRosettaConstructionDeriveResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRosettaConstructionDeriveResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRosettaConstructionDeriveResponse(val *RosettaConstructionDeriveResponse) *NullableRosettaConstructionDeriveResponse {
	return &NullableRosettaConstructionDeriveResponse{value: val, isSet: true}
}

func (v NullableRosettaConstructionDeriveResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRosettaConstructionDeriveResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
