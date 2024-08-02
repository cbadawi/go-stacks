package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the RosettaMempoolResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RosettaMempoolResponse{}

// RosettaMempoolResponse A MempoolResponse contains all transaction identifiers in the mempool for a particular network_identifier.
type RosettaMempoolResponse struct {
	//
	TransactionIdentifiers []TransactionIdentifier `json:"transaction_identifiers"`
	Metadata               map[string]interface{}  `json:"metadata,omitempty"`
}

type _RosettaMempoolResponse RosettaMempoolResponse

// NewRosettaMempoolResponse instantiates a new RosettaMempoolResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRosettaMempoolResponse(transactionIdentifiers []TransactionIdentifier) *RosettaMempoolResponse {
	this := RosettaMempoolResponse{}
	this.TransactionIdentifiers = transactionIdentifiers
	return &this
}

// NewRosettaMempoolResponseWithDefaults instantiates a new RosettaMempoolResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRosettaMempoolResponseWithDefaults() *RosettaMempoolResponse {
	this := RosettaMempoolResponse{}
	return &this
}

// GetTransactionIdentifiers returns the TransactionIdentifiers field value
func (o *RosettaMempoolResponse) GetTransactionIdentifiers() []TransactionIdentifier {
	if o == nil {
		var ret []TransactionIdentifier
		return ret
	}

	return o.TransactionIdentifiers
}

// GetTransactionIdentifiersOk returns a tuple with the TransactionIdentifiers field value
// and a boolean to check if the value has been set.
func (o *RosettaMempoolResponse) GetTransactionIdentifiersOk() ([]TransactionIdentifier, bool) {
	if o == nil {
		return nil, false
	}
	return o.TransactionIdentifiers, true
}

// SetTransactionIdentifiers sets field value
func (o *RosettaMempoolResponse) SetTransactionIdentifiers(v []TransactionIdentifier) {
	o.TransactionIdentifiers = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *RosettaMempoolResponse) GetMetadata() map[string]interface{} {
	if o == nil || utils.IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaMempoolResponse) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || utils.IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *RosettaMempoolResponse) HasMetadata() bool {
	if o != nil && !utils.IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *RosettaMempoolResponse) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o RosettaMempoolResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RosettaMempoolResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["transaction_identifiers"] = o.TransactionIdentifiers
	if !utils.IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

func (o *RosettaMempoolResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"transaction_identifiers",
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

	varRosettaMempoolResponse := _RosettaMempoolResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRosettaMempoolResponse)

	if err != nil {
		return err
	}

	*o = RosettaMempoolResponse(varRosettaMempoolResponse)

	return err
}

type NullableRosettaMempoolResponse struct {
	value *RosettaMempoolResponse
	isSet bool
}

func (v NullableRosettaMempoolResponse) Get() *RosettaMempoolResponse {
	return v.value
}

func (v *NullableRosettaMempoolResponse) Set(val *RosettaMempoolResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRosettaMempoolResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRosettaMempoolResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRosettaMempoolResponse(val *RosettaMempoolResponse) *NullableRosettaMempoolResponse {
	return &NullableRosettaMempoolResponse{value: val, isSet: true}
}

func (v NullableRosettaMempoolResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRosettaMempoolResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
