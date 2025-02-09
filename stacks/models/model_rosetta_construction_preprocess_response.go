package models

import (
	"encoding/json"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the RosettaConstructionPreprocessResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RosettaConstructionPreprocessResponse{}

// RosettaConstructionPreprocessResponse RosettaConstructionPreprocessResponse contains options that will be sent unmodified to /construction/metadata. If it is not necessary to make a request to /construction/metadata, options should be omitted. Some blockchains require the PublicKey of particular AccountIdentifiers to construct a valid transaction. To fetch these PublicKeys, populate required_public_keys with the AccountIdentifiers associated with the desired PublicKeys. If it is not necessary to retrieve any PublicKeys for construction, required_public_keys should be omitted.
type RosettaConstructionPreprocessResponse struct {
	Options            *RosettaOptions  `json:"options,omitempty"`
	RequiredPublicKeys []RosettaAccount `json:"required_public_keys,omitempty"`
}

// NewRosettaConstructionPreprocessResponse instantiates a new RosettaConstructionPreprocessResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRosettaConstructionPreprocessResponse() *RosettaConstructionPreprocessResponse {
	this := RosettaConstructionPreprocessResponse{}
	return &this
}

// NewRosettaConstructionPreprocessResponseWithDefaults instantiates a new RosettaConstructionPreprocessResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRosettaConstructionPreprocessResponseWithDefaults() *RosettaConstructionPreprocessResponse {
	this := RosettaConstructionPreprocessResponse{}
	return &this
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *RosettaConstructionPreprocessResponse) GetOptions() RosettaOptions {
	if o == nil || utils.IsNil(o.Options) {
		var ret RosettaOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaConstructionPreprocessResponse) GetOptionsOk() (*RosettaOptions, bool) {
	if o == nil || utils.IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *RosettaConstructionPreprocessResponse) HasOptions() bool {
	if o != nil && !utils.IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given RosettaOptions and assigns it to the Options field.
func (o *RosettaConstructionPreprocessResponse) SetOptions(v RosettaOptions) {
	o.Options = &v
}

// GetRequiredPublicKeys returns the RequiredPublicKeys field value if set, zero value otherwise.
func (o *RosettaConstructionPreprocessResponse) GetRequiredPublicKeys() []RosettaAccount {
	if o == nil || utils.IsNil(o.RequiredPublicKeys) {
		var ret []RosettaAccount
		return ret
	}
	return o.RequiredPublicKeys
}

// GetRequiredPublicKeysOk returns a tuple with the RequiredPublicKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaConstructionPreprocessResponse) GetRequiredPublicKeysOk() ([]RosettaAccount, bool) {
	if o == nil || utils.IsNil(o.RequiredPublicKeys) {
		return nil, false
	}
	return o.RequiredPublicKeys, true
}

// HasRequiredPublicKeys returns a boolean if a field has been set.
func (o *RosettaConstructionPreprocessResponse) HasRequiredPublicKeys() bool {
	if o != nil && !utils.IsNil(o.RequiredPublicKeys) {
		return true
	}

	return false
}

// SetRequiredPublicKeys gets a reference to the given []RosettaAccount and assigns it to the RequiredPublicKeys field.
func (o *RosettaConstructionPreprocessResponse) SetRequiredPublicKeys(v []RosettaAccount) {
	o.RequiredPublicKeys = v
}

func (o RosettaConstructionPreprocessResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RosettaConstructionPreprocessResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	if !utils.IsNil(o.RequiredPublicKeys) {
		toSerialize["required_public_keys"] = o.RequiredPublicKeys
	}
	return toSerialize, nil
}

type NullableRosettaConstructionPreprocessResponse struct {
	value *RosettaConstructionPreprocessResponse
	isSet bool
}

func (v NullableRosettaConstructionPreprocessResponse) Get() *RosettaConstructionPreprocessResponse {
	return v.value
}

func (v *NullableRosettaConstructionPreprocessResponse) Set(val *RosettaConstructionPreprocessResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRosettaConstructionPreprocessResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRosettaConstructionPreprocessResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRosettaConstructionPreprocessResponse(val *RosettaConstructionPreprocessResponse) *NullableRosettaConstructionPreprocessResponse {
	return &NullableRosettaConstructionPreprocessResponse{value: val, isSet: true}
}

func (v NullableRosettaConstructionPreprocessResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRosettaConstructionPreprocessResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
