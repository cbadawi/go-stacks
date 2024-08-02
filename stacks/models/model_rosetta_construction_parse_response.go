package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the RosettaConstructionParseResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RosettaConstructionParseResponse{}

// RosettaConstructionParseResponse RosettaConstructionParseResponse contains an array of operations that occur in a transaction blob. This should match the array of operations provided to /construction/preprocess and /construction/payloads.
type RosettaConstructionParseResponse struct {
	Operations []RosettaOperation `json:"operations"`
	// [DEPRECATED by account_identifier_signers in v1.4.4] All signers (addresses) of a particular transaction. If the transaction is unsigned, it should be empty.
	Signers                  []RosettaNetworkOptionsResponseAllowOperationTypesInner `json:"signers,omitempty"`
	AccountIdentifierSigners []RosettaAccountIdentifier                              `json:"account_identifier_signers,omitempty"`
	Metadata                 map[string]interface{}                                  `json:"metadata,omitempty"`
}

type _RosettaConstructionParseResponse RosettaConstructionParseResponse

// NewRosettaConstructionParseResponse instantiates a new RosettaConstructionParseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRosettaConstructionParseResponse(operations []RosettaOperation) *RosettaConstructionParseResponse {
	this := RosettaConstructionParseResponse{}
	this.Operations = operations
	return &this
}

// NewRosettaConstructionParseResponseWithDefaults instantiates a new RosettaConstructionParseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRosettaConstructionParseResponseWithDefaults() *RosettaConstructionParseResponse {
	this := RosettaConstructionParseResponse{}
	return &this
}

// GetOperations returns the Operations field value
func (o *RosettaConstructionParseResponse) GetOperations() []RosettaOperation {
	if o == nil {
		var ret []RosettaOperation
		return ret
	}

	return o.Operations
}

// GetOperationsOk returns a tuple with the Operations field value
// and a boolean to check if the value has been set.
func (o *RosettaConstructionParseResponse) GetOperationsOk() ([]RosettaOperation, bool) {
	if o == nil {
		return nil, false
	}
	return o.Operations, true
}

// SetOperations sets field value
func (o *RosettaConstructionParseResponse) SetOperations(v []RosettaOperation) {
	o.Operations = v
}

// GetSigners returns the Signers field value if set, zero value otherwise.
func (o *RosettaConstructionParseResponse) GetSigners() []RosettaNetworkOptionsResponseAllowOperationTypesInner {
	if o == nil || utils.IsNil(o.Signers) {
		var ret []RosettaNetworkOptionsResponseAllowOperationTypesInner
		return ret
	}
	return o.Signers
}

// GetSignersOk returns a tuple with the Signers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaConstructionParseResponse) GetSignersOk() ([]RosettaNetworkOptionsResponseAllowOperationTypesInner, bool) {
	if o == nil || utils.IsNil(o.Signers) {
		return nil, false
	}
	return o.Signers, true
}

// HasSigners returns a boolean if a field has been set.
func (o *RosettaConstructionParseResponse) HasSigners() bool {
	if o != nil && !utils.IsNil(o.Signers) {
		return true
	}

	return false
}

// SetSigners gets a reference to the given []RosettaNetworkOptionsResponseAllowOperationTypesInner and assigns it to the Signers field.
func (o *RosettaConstructionParseResponse) SetSigners(v []RosettaNetworkOptionsResponseAllowOperationTypesInner) {
	o.Signers = v
}

// GetAccountIdentifierSigners returns the AccountIdentifierSigners field value if set, zero value otherwise.
func (o *RosettaConstructionParseResponse) GetAccountIdentifierSigners() []RosettaAccountIdentifier {
	if o == nil || utils.IsNil(o.AccountIdentifierSigners) {
		var ret []RosettaAccountIdentifier
		return ret
	}
	return o.AccountIdentifierSigners
}

// GetAccountIdentifierSignersOk returns a tuple with the AccountIdentifierSigners field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaConstructionParseResponse) GetAccountIdentifierSignersOk() ([]RosettaAccountIdentifier, bool) {
	if o == nil || utils.IsNil(o.AccountIdentifierSigners) {
		return nil, false
	}
	return o.AccountIdentifierSigners, true
}

// HasAccountIdentifierSigners returns a boolean if a field has been set.
func (o *RosettaConstructionParseResponse) HasAccountIdentifierSigners() bool {
	if o != nil && !utils.IsNil(o.AccountIdentifierSigners) {
		return true
	}

	return false
}

// SetAccountIdentifierSigners gets a reference to the given []RosettaAccountIdentifier and assigns it to the AccountIdentifierSigners field.
func (o *RosettaConstructionParseResponse) SetAccountIdentifierSigners(v []RosettaAccountIdentifier) {
	o.AccountIdentifierSigners = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *RosettaConstructionParseResponse) GetMetadata() map[string]interface{} {
	if o == nil || utils.IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaConstructionParseResponse) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || utils.IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *RosettaConstructionParseResponse) HasMetadata() bool {
	if o != nil && !utils.IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *RosettaConstructionParseResponse) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o RosettaConstructionParseResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RosettaConstructionParseResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["operations"] = o.Operations
	if !utils.IsNil(o.Signers) {
		toSerialize["signers"] = o.Signers
	}
	if !utils.IsNil(o.AccountIdentifierSigners) {
		toSerialize["account_identifier_signers"] = o.AccountIdentifierSigners
	}
	if !utils.IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

func (o *RosettaConstructionParseResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varRosettaConstructionParseResponse := _RosettaConstructionParseResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRosettaConstructionParseResponse)

	if err != nil {
		return err
	}

	*o = RosettaConstructionParseResponse(varRosettaConstructionParseResponse)

	return err
}

type NullableRosettaConstructionParseResponse struct {
	value *RosettaConstructionParseResponse
	isSet bool
}

func (v NullableRosettaConstructionParseResponse) Get() *RosettaConstructionParseResponse {
	return v.value
}

func (v *NullableRosettaConstructionParseResponse) Set(val *RosettaConstructionParseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRosettaConstructionParseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRosettaConstructionParseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRosettaConstructionParseResponse(val *RosettaConstructionParseResponse) *NullableRosettaConstructionParseResponse {
	return &NullableRosettaConstructionParseResponse{value: val, isSet: true}
}

func (v NullableRosettaConstructionParseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRosettaConstructionParseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
