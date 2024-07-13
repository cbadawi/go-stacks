package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the NonFungibleTokenMintWithTxMetadata type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &NonFungibleTokenMintWithTxMetadata{}

// NonFungibleTokenMintWithTxMetadata Non-Fungible Token mint event with transaction metadata
type NonFungibleTokenMintWithTxMetadata struct {
	Recipient  *string                              `json:"recipient,omitempty"`
	EventIndex int32                                `json:"event_index"`
	Value      NonFungibleTokenHoldingWithTxIdValue `json:"value"`
	Tx         Transaction                          `json:"tx"`
}

type _NonFungibleTokenMintWithTxMetadata NonFungibleTokenMintWithTxMetadata

// NewNonFungibleTokenMintWithTxMetadata instantiates a new NonFungibleTokenMintWithTxMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNonFungibleTokenMintWithTxMetadata(eventIndex int32, value NonFungibleTokenHoldingWithTxIdValue, tx Transaction) *NonFungibleTokenMintWithTxMetadata {
	this := NonFungibleTokenMintWithTxMetadata{}
	this.EventIndex = eventIndex
	this.Value = value
	this.Tx = tx
	return &this
}

// NewNonFungibleTokenMintWithTxMetadataWithDefaults instantiates a new NonFungibleTokenMintWithTxMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNonFungibleTokenMintWithTxMetadataWithDefaults() *NonFungibleTokenMintWithTxMetadata {
	this := NonFungibleTokenMintWithTxMetadata{}
	return &this
}

// GetRecipient returns the Recipient field value if set, zero value otherwise.
func (o *NonFungibleTokenMintWithTxMetadata) GetRecipient() string {
	if o == nil || utils.IsNil(o.Recipient) {
		var ret string
		return ret
	}
	return *o.Recipient
}

// GetRecipientOk returns a tuple with the Recipient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonFungibleTokenMintWithTxMetadata) GetRecipientOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Recipient) {
		return nil, false
	}
	return o.Recipient, true
}

// HasRecipient returns a boolean if a field has been set.
func (o *NonFungibleTokenMintWithTxMetadata) HasRecipient() bool {
	if o != nil && !utils.IsNil(o.Recipient) {
		return true
	}

	return false
}

// SetRecipient gets a reference to the given string and assigns it to the Recipient field.
func (o *NonFungibleTokenMintWithTxMetadata) SetRecipient(v string) {
	o.Recipient = &v
}

// GetEventIndex returns the EventIndex field value
func (o *NonFungibleTokenMintWithTxMetadata) GetEventIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EventIndex
}

// GetEventIndexOk returns a tuple with the EventIndex field value
// and a boolean to check if the value has been set.
func (o *NonFungibleTokenMintWithTxMetadata) GetEventIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventIndex, true
}

// SetEventIndex sets field value
func (o *NonFungibleTokenMintWithTxMetadata) SetEventIndex(v int32) {
	o.EventIndex = v
}

// GetValue returns the Value field value
func (o *NonFungibleTokenMintWithTxMetadata) GetValue() NonFungibleTokenHoldingWithTxIdValue {
	if o == nil {
		var ret NonFungibleTokenHoldingWithTxIdValue
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *NonFungibleTokenMintWithTxMetadata) GetValueOk() (*NonFungibleTokenHoldingWithTxIdValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *NonFungibleTokenMintWithTxMetadata) SetValue(v NonFungibleTokenHoldingWithTxIdValue) {
	o.Value = v
}

// GetTx returns the Tx field value
func (o *NonFungibleTokenMintWithTxMetadata) GetTx() Transaction {
	if o == nil {
		var ret Transaction
		return ret
	}

	return o.Tx
}

// GetTxOk returns a tuple with the Tx field value
// and a boolean to check if the value has been set.
func (o *NonFungibleTokenMintWithTxMetadata) GetTxOk() (*Transaction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tx, true
}

// SetTx sets field value
func (o *NonFungibleTokenMintWithTxMetadata) SetTx(v Transaction) {
	o.Tx = v
}

func (o NonFungibleTokenMintWithTxMetadata) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NonFungibleTokenMintWithTxMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Recipient) {
		toSerialize["recipient"] = o.Recipient
	}
	toSerialize["event_index"] = o.EventIndex
	toSerialize["value"] = o.Value
	toSerialize["tx"] = o.Tx
	return toSerialize, nil
}

func (o *NonFungibleTokenMintWithTxMetadata) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"event_index",
		"value",
		"tx",
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

	varNonFungibleTokenMintWithTxMetadata := _NonFungibleTokenMintWithTxMetadata{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNonFungibleTokenMintWithTxMetadata)

	if err != nil {
		return err
	}

	*o = NonFungibleTokenMintWithTxMetadata(varNonFungibleTokenMintWithTxMetadata)

	return err
}

type NullableNonFungibleTokenMintWithTxMetadata struct {
	value *NonFungibleTokenMintWithTxMetadata
	isSet bool
}

func (v NullableNonFungibleTokenMintWithTxMetadata) Get() *NonFungibleTokenMintWithTxMetadata {
	return v.value
}

func (v *NullableNonFungibleTokenMintWithTxMetadata) Set(val *NonFungibleTokenMintWithTxMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableNonFungibleTokenMintWithTxMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableNonFungibleTokenMintWithTxMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNonFungibleTokenMintWithTxMetadata(val *NonFungibleTokenMintWithTxMetadata) *NullableNonFungibleTokenMintWithTxMetadata {
	return &NullableNonFungibleTokenMintWithTxMetadata{value: val, isSet: true}
}

func (v NullableNonFungibleTokenMintWithTxMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNonFungibleTokenMintWithTxMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
