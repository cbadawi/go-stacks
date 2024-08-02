package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the RosettaConstructionPreprocessRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RosettaConstructionPreprocessRequest{}

// RosettaConstructionPreprocessRequest ConstructionPreprocessRequest is passed to the /construction/preprocess endpoint so that a Rosetta implementation can determine which metadata it needs to request for construction
type RosettaConstructionPreprocessRequest struct {
	NetworkIdentifier NetworkIdentifier      `json:"network_identifier"`
	Operations        []RosettaOperation     `json:"operations"`
	Metadata          map[string]interface{} `json:"metadata,omitempty"`
	MaxFee            []RosettaMaxFeeAmount  `json:"max_fee,omitempty"`
	//  The caller can also provide a suggested fee multiplier to indicate that the suggested fee should be scaled. This may be used to set higher fees for urgent transactions or to pay lower fees when there is less urgency. It is assumed that providing a very low multiplier (like 0.0001) will never lead to a transaction being created with a fee less than the minimum network fee (if applicable). In the case that the caller provides both a max fee and a suggested fee multiplier, the max fee will set an upper bound on the suggested fee (regardless of the multiplier provided).
	SuggestedFeeMultiplier *int32 `json:"suggested_fee_multiplier,omitempty"`
}

type _RosettaConstructionPreprocessRequest RosettaConstructionPreprocessRequest

// NewRosettaConstructionPreprocessRequest instantiates a new RosettaConstructionPreprocessRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRosettaConstructionPreprocessRequest(networkIdentifier NetworkIdentifier, operations []RosettaOperation) *RosettaConstructionPreprocessRequest {
	this := RosettaConstructionPreprocessRequest{}
	this.NetworkIdentifier = networkIdentifier
	this.Operations = operations
	return &this
}

// NewRosettaConstructionPreprocessRequestWithDefaults instantiates a new RosettaConstructionPreprocessRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRosettaConstructionPreprocessRequestWithDefaults() *RosettaConstructionPreprocessRequest {
	this := RosettaConstructionPreprocessRequest{}
	return &this
}

// GetNetworkIdentifier returns the NetworkIdentifier field value
func (o *RosettaConstructionPreprocessRequest) GetNetworkIdentifier() NetworkIdentifier {
	if o == nil {
		var ret NetworkIdentifier
		return ret
	}

	return o.NetworkIdentifier
}

// GetNetworkIdentifierOk returns a tuple with the NetworkIdentifier field value
// and a boolean to check if the value has been set.
func (o *RosettaConstructionPreprocessRequest) GetNetworkIdentifierOk() (*NetworkIdentifier, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkIdentifier, true
}

// SetNetworkIdentifier sets field value
func (o *RosettaConstructionPreprocessRequest) SetNetworkIdentifier(v NetworkIdentifier) {
	o.NetworkIdentifier = v
}

// GetOperations returns the Operations field value
func (o *RosettaConstructionPreprocessRequest) GetOperations() []RosettaOperation {
	if o == nil {
		var ret []RosettaOperation
		return ret
	}

	return o.Operations
}

// GetOperationsOk returns a tuple with the Operations field value
// and a boolean to check if the value has been set.
func (o *RosettaConstructionPreprocessRequest) GetOperationsOk() ([]RosettaOperation, bool) {
	if o == nil {
		return nil, false
	}
	return o.Operations, true
}

// SetOperations sets field value
func (o *RosettaConstructionPreprocessRequest) SetOperations(v []RosettaOperation) {
	o.Operations = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *RosettaConstructionPreprocessRequest) GetMetadata() map[string]interface{} {
	if o == nil || utils.IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaConstructionPreprocessRequest) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || utils.IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *RosettaConstructionPreprocessRequest) HasMetadata() bool {
	if o != nil && !utils.IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *RosettaConstructionPreprocessRequest) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetMaxFee returns the MaxFee field value if set, zero value otherwise.
func (o *RosettaConstructionPreprocessRequest) GetMaxFee() []RosettaMaxFeeAmount {
	if o == nil || utils.IsNil(o.MaxFee) {
		var ret []RosettaMaxFeeAmount
		return ret
	}
	return o.MaxFee
}

// GetMaxFeeOk returns a tuple with the MaxFee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaConstructionPreprocessRequest) GetMaxFeeOk() ([]RosettaMaxFeeAmount, bool) {
	if o == nil || utils.IsNil(o.MaxFee) {
		return nil, false
	}
	return o.MaxFee, true
}

// HasMaxFee returns a boolean if a field has been set.
func (o *RosettaConstructionPreprocessRequest) HasMaxFee() bool {
	if o != nil && !utils.IsNil(o.MaxFee) {
		return true
	}

	return false
}

// SetMaxFee gets a reference to the given []RosettaMaxFeeAmount and assigns it to the MaxFee field.
func (o *RosettaConstructionPreprocessRequest) SetMaxFee(v []RosettaMaxFeeAmount) {
	o.MaxFee = v
}

// GetSuggestedFeeMultiplier returns the SuggestedFeeMultiplier field value if set, zero value otherwise.
func (o *RosettaConstructionPreprocessRequest) GetSuggestedFeeMultiplier() int32 {
	if o == nil || utils.IsNil(o.SuggestedFeeMultiplier) {
		var ret int32
		return ret
	}
	return *o.SuggestedFeeMultiplier
}

// GetSuggestedFeeMultiplierOk returns a tuple with the SuggestedFeeMultiplier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaConstructionPreprocessRequest) GetSuggestedFeeMultiplierOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.SuggestedFeeMultiplier) {
		return nil, false
	}
	return o.SuggestedFeeMultiplier, true
}

// HasSuggestedFeeMultiplier returns a boolean if a field has been set.
func (o *RosettaConstructionPreprocessRequest) HasSuggestedFeeMultiplier() bool {
	if o != nil && !utils.IsNil(o.SuggestedFeeMultiplier) {
		return true
	}

	return false
}

// SetSuggestedFeeMultiplier gets a reference to the given int32 and assigns it to the SuggestedFeeMultiplier field.
func (o *RosettaConstructionPreprocessRequest) SetSuggestedFeeMultiplier(v int32) {
	o.SuggestedFeeMultiplier = &v
}

func (o RosettaConstructionPreprocessRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RosettaConstructionPreprocessRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["network_identifier"] = o.NetworkIdentifier
	toSerialize["operations"] = o.Operations
	if !utils.IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !utils.IsNil(o.MaxFee) {
		toSerialize["max_fee"] = o.MaxFee
	}
	if !utils.IsNil(o.SuggestedFeeMultiplier) {
		toSerialize["suggested_fee_multiplier"] = o.SuggestedFeeMultiplier
	}
	return toSerialize, nil
}

func (o *RosettaConstructionPreprocessRequest) UnmarshalJSON(data []byte) (err error) {
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

	varRosettaConstructionPreprocessRequest := _RosettaConstructionPreprocessRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRosettaConstructionPreprocessRequest)

	if err != nil {
		return err
	}

	*o = RosettaConstructionPreprocessRequest(varRosettaConstructionPreprocessRequest)

	return err
}

type NullableRosettaConstructionPreprocessRequest struct {
	value *RosettaConstructionPreprocessRequest
	isSet bool
}

func (v NullableRosettaConstructionPreprocessRequest) Get() *RosettaConstructionPreprocessRequest {
	return v.value
}

func (v *NullableRosettaConstructionPreprocessRequest) Set(val *RosettaConstructionPreprocessRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRosettaConstructionPreprocessRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRosettaConstructionPreprocessRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRosettaConstructionPreprocessRequest(val *RosettaConstructionPreprocessRequest) *NullableRosettaConstructionPreprocessRequest {
	return &NullableRosettaConstructionPreprocessRequest{value: val, isSet: true}
}

func (v NullableRosettaConstructionPreprocessRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRosettaConstructionPreprocessRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
