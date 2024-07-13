package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the RosettaConstructionMetadataResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RosettaConstructionMetadataResponse{}

// RosettaConstructionMetadataResponse The ConstructionMetadataResponse returns network-specific metadata used for transaction construction. Optionally, the implementer can return the suggested fee associated with the transaction being constructed. The caller may use this info to adjust the intent of the transaction or to create a transaction with a different account that can pay the suggested fee. Suggested fee is an array in case fee payment must occur in multiple currencies.
type RosettaConstructionMetadataResponse struct {
	Metadata     RosettaConstructionMetadataResponseMetadata `json:"metadata"`
	SuggestedFee []RosettaAmount                             `json:"suggested_fee,omitempty"`
}

type _RosettaConstructionMetadataResponse RosettaConstructionMetadataResponse

// NewRosettaConstructionMetadataResponse instantiates a new RosettaConstructionMetadataResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRosettaConstructionMetadataResponse(metadata RosettaConstructionMetadataResponseMetadata) *RosettaConstructionMetadataResponse {
	this := RosettaConstructionMetadataResponse{}
	this.Metadata = metadata
	return &this
}

// NewRosettaConstructionMetadataResponseWithDefaults instantiates a new RosettaConstructionMetadataResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRosettaConstructionMetadataResponseWithDefaults() *RosettaConstructionMetadataResponse {
	this := RosettaConstructionMetadataResponse{}
	return &this
}

// GetMetadata returns the Metadata field value
func (o *RosettaConstructionMetadataResponse) GetMetadata() RosettaConstructionMetadataResponseMetadata {
	if o == nil {
		var ret RosettaConstructionMetadataResponseMetadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *RosettaConstructionMetadataResponse) GetMetadataOk() (*RosettaConstructionMetadataResponseMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *RosettaConstructionMetadataResponse) SetMetadata(v RosettaConstructionMetadataResponseMetadata) {
	o.Metadata = v
}

// GetSuggestedFee returns the SuggestedFee field value if set, zero value otherwise.
func (o *RosettaConstructionMetadataResponse) GetSuggestedFee() []RosettaAmount {
	if o == nil || utils.IsNil(o.SuggestedFee) {
		var ret []RosettaAmount
		return ret
	}
	return o.SuggestedFee
}

// GetSuggestedFeeOk returns a tuple with the SuggestedFee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaConstructionMetadataResponse) GetSuggestedFeeOk() ([]RosettaAmount, bool) {
	if o == nil || utils.IsNil(o.SuggestedFee) {
		return nil, false
	}
	return o.SuggestedFee, true
}

// HasSuggestedFee returns a boolean if a field has been set.
func (o *RosettaConstructionMetadataResponse) HasSuggestedFee() bool {
	if o != nil && !utils.IsNil(o.SuggestedFee) {
		return true
	}

	return false
}

// SetSuggestedFee gets a reference to the given []RosettaAmount and assigns it to the SuggestedFee field.
func (o *RosettaConstructionMetadataResponse) SetSuggestedFee(v []RosettaAmount) {
	o.SuggestedFee = v
}

func (o RosettaConstructionMetadataResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RosettaConstructionMetadataResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["metadata"] = o.Metadata
	if !utils.IsNil(o.SuggestedFee) {
		toSerialize["suggested_fee"] = o.SuggestedFee
	}
	return toSerialize, nil
}

func (o *RosettaConstructionMetadataResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"metadata",
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

	varRosettaConstructionMetadataResponse := _RosettaConstructionMetadataResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRosettaConstructionMetadataResponse)

	if err != nil {
		return err
	}

	*o = RosettaConstructionMetadataResponse(varRosettaConstructionMetadataResponse)

	return err
}

type NullableRosettaConstructionMetadataResponse struct {
	value *RosettaConstructionMetadataResponse
	isSet bool
}

func (v NullableRosettaConstructionMetadataResponse) Get() *RosettaConstructionMetadataResponse {
	return v.value
}

func (v *NullableRosettaConstructionMetadataResponse) Set(val *RosettaConstructionMetadataResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRosettaConstructionMetadataResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRosettaConstructionMetadataResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRosettaConstructionMetadataResponse(val *RosettaConstructionMetadataResponse) *NullableRosettaConstructionMetadataResponse {
	return &NullableRosettaConstructionMetadataResponse{value: val, isSet: true}
}

func (v NullableRosettaConstructionMetadataResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRosettaConstructionMetadataResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
