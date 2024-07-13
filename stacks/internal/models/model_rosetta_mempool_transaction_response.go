package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the RosettaMempoolTransactionResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RosettaMempoolTransactionResponse{}

// RosettaMempoolTransactionResponse A MempoolTransactionResponse contains an estimate of a mempool transaction. It may not be possible to know the full impact of a transaction in the mempool (ex: fee paid).
type RosettaMempoolTransactionResponse struct {
	Transaction RosettaTransaction     `json:"transaction"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

type _RosettaMempoolTransactionResponse RosettaMempoolTransactionResponse

// NewRosettaMempoolTransactionResponse instantiates a new RosettaMempoolTransactionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRosettaMempoolTransactionResponse(transaction RosettaTransaction) *RosettaMempoolTransactionResponse {
	this := RosettaMempoolTransactionResponse{}
	this.Transaction = transaction
	return &this
}

// NewRosettaMempoolTransactionResponseWithDefaults instantiates a new RosettaMempoolTransactionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRosettaMempoolTransactionResponseWithDefaults() *RosettaMempoolTransactionResponse {
	this := RosettaMempoolTransactionResponse{}
	return &this
}

// GetTransaction returns the Transaction field value
func (o *RosettaMempoolTransactionResponse) GetTransaction() RosettaTransaction {
	if o == nil {
		var ret RosettaTransaction
		return ret
	}

	return o.Transaction
}

// GetTransactionOk returns a tuple with the Transaction field value
// and a boolean to check if the value has been set.
func (o *RosettaMempoolTransactionResponse) GetTransactionOk() (*RosettaTransaction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Transaction, true
}

// SetTransaction sets field value
func (o *RosettaMempoolTransactionResponse) SetTransaction(v RosettaTransaction) {
	o.Transaction = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *RosettaMempoolTransactionResponse) GetMetadata() map[string]interface{} {
	if o == nil || utils.IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaMempoolTransactionResponse) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || utils.IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *RosettaMempoolTransactionResponse) HasMetadata() bool {
	if o != nil && !utils.IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *RosettaMempoolTransactionResponse) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o RosettaMempoolTransactionResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RosettaMempoolTransactionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["transaction"] = o.Transaction
	if !utils.IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

func (o *RosettaMempoolTransactionResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"transaction",
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

	varRosettaMempoolTransactionResponse := _RosettaMempoolTransactionResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRosettaMempoolTransactionResponse)

	if err != nil {
		return err
	}

	*o = RosettaMempoolTransactionResponse(varRosettaMempoolTransactionResponse)

	return err
}

type NullableRosettaMempoolTransactionResponse struct {
	value *RosettaMempoolTransactionResponse
	isSet bool
}

func (v NullableRosettaMempoolTransactionResponse) Get() *RosettaMempoolTransactionResponse {
	return v.value
}

func (v *NullableRosettaMempoolTransactionResponse) Set(val *RosettaMempoolTransactionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRosettaMempoolTransactionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRosettaMempoolTransactionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRosettaMempoolTransactionResponse(val *RosettaMempoolTransactionResponse) *NullableRosettaMempoolTransactionResponse {
	return &NullableRosettaMempoolTransactionResponse{value: val, isSet: true}
}

func (v NullableRosettaMempoolTransactionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRosettaMempoolTransactionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
