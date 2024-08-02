package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the RosettaConstructionSubmitResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RosettaConstructionSubmitResponse{}

// RosettaConstructionSubmitResponse TransactionIdentifier contains the transaction_identifier of a transaction that was submitted to either /construction/submit.
type RosettaConstructionSubmitResponse struct {
	TransactionIdentifier TransactionIdentifier  `json:"transaction_identifier"`
	Metadata              map[string]interface{} `json:"metadata,omitempty"`
}

type _RosettaConstructionSubmitResponse RosettaConstructionSubmitResponse

// NewRosettaConstructionSubmitResponse instantiates a new RosettaConstructionSubmitResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRosettaConstructionSubmitResponse(transactionIdentifier TransactionIdentifier) *RosettaConstructionSubmitResponse {
	this := RosettaConstructionSubmitResponse{}
	this.TransactionIdentifier = transactionIdentifier
	return &this
}

// NewRosettaConstructionSubmitResponseWithDefaults instantiates a new RosettaConstructionSubmitResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRosettaConstructionSubmitResponseWithDefaults() *RosettaConstructionSubmitResponse {
	this := RosettaConstructionSubmitResponse{}
	return &this
}

// GetTransactionIdentifier returns the TransactionIdentifier field value
func (o *RosettaConstructionSubmitResponse) GetTransactionIdentifier() TransactionIdentifier {
	if o == nil {
		var ret TransactionIdentifier
		return ret
	}

	return o.TransactionIdentifier
}

// GetTransactionIdentifierOk returns a tuple with the TransactionIdentifier field value
// and a boolean to check if the value has been set.
func (o *RosettaConstructionSubmitResponse) GetTransactionIdentifierOk() (*TransactionIdentifier, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionIdentifier, true
}

// SetTransactionIdentifier sets field value
func (o *RosettaConstructionSubmitResponse) SetTransactionIdentifier(v TransactionIdentifier) {
	o.TransactionIdentifier = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *RosettaConstructionSubmitResponse) GetMetadata() map[string]interface{} {
	if o == nil || utils.IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaConstructionSubmitResponse) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || utils.IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *RosettaConstructionSubmitResponse) HasMetadata() bool {
	if o != nil && !utils.IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *RosettaConstructionSubmitResponse) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o RosettaConstructionSubmitResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RosettaConstructionSubmitResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["transaction_identifier"] = o.TransactionIdentifier
	if !utils.IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

func (o *RosettaConstructionSubmitResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"transaction_identifier",
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

	varRosettaConstructionSubmitResponse := _RosettaConstructionSubmitResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRosettaConstructionSubmitResponse)

	if err != nil {
		return err
	}

	*o = RosettaConstructionSubmitResponse(varRosettaConstructionSubmitResponse)

	return err
}

type NullableRosettaConstructionSubmitResponse struct {
	value *RosettaConstructionSubmitResponse
	isSet bool
}

func (v NullableRosettaConstructionSubmitResponse) Get() *RosettaConstructionSubmitResponse {
	return v.value
}

func (v *NullableRosettaConstructionSubmitResponse) Set(val *RosettaConstructionSubmitResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRosettaConstructionSubmitResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRosettaConstructionSubmitResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRosettaConstructionSubmitResponse(val *RosettaConstructionSubmitResponse) *NullableRosettaConstructionSubmitResponse {
	return &NullableRosettaConstructionSubmitResponse{value: val, isSet: true}
}

func (v NullableRosettaConstructionSubmitResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRosettaConstructionSubmitResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
