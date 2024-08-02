package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the NonFungibleTokenMintWithTxId type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &NonFungibleTokenMintWithTxId{}

// NonFungibleTokenMintWithTxId Non-Fungible Token mint event with transaction id
type NonFungibleTokenMintWithTxId struct {
	Recipient  *string                              `json:"recipient,omitempty"`
	EventIndex int32                                `json:"event_index"`
	Value      NonFungibleTokenHoldingWithTxIdValue `json:"value"`
	TxId       string                               `json:"tx_id"`
}

type _NonFungibleTokenMintWithTxId NonFungibleTokenMintWithTxId

// NewNonFungibleTokenMintWithTxId instantiates a new NonFungibleTokenMintWithTxId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNonFungibleTokenMintWithTxId(eventIndex int32, value NonFungibleTokenHoldingWithTxIdValue, txId string) *NonFungibleTokenMintWithTxId {
	this := NonFungibleTokenMintWithTxId{}
	this.EventIndex = eventIndex
	this.Value = value
	this.TxId = txId
	return &this
}

// NewNonFungibleTokenMintWithTxIdWithDefaults instantiates a new NonFungibleTokenMintWithTxId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNonFungibleTokenMintWithTxIdWithDefaults() *NonFungibleTokenMintWithTxId {
	this := NonFungibleTokenMintWithTxId{}
	return &this
}

// GetRecipient returns the Recipient field value if set, zero value otherwise.
func (o *NonFungibleTokenMintWithTxId) GetRecipient() string {
	if o == nil || utils.IsNil(o.Recipient) {
		var ret string
		return ret
	}
	return *o.Recipient
}

// GetRecipientOk returns a tuple with the Recipient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonFungibleTokenMintWithTxId) GetRecipientOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Recipient) {
		return nil, false
	}
	return o.Recipient, true
}

// HasRecipient returns a boolean if a field has been set.
func (o *NonFungibleTokenMintWithTxId) HasRecipient() bool {
	if o != nil && !utils.IsNil(o.Recipient) {
		return true
	}

	return false
}

// SetRecipient gets a reference to the given string and assigns it to the Recipient field.
func (o *NonFungibleTokenMintWithTxId) SetRecipient(v string) {
	o.Recipient = &v
}

// GetEventIndex returns the EventIndex field value
func (o *NonFungibleTokenMintWithTxId) GetEventIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EventIndex
}

// GetEventIndexOk returns a tuple with the EventIndex field value
// and a boolean to check if the value has been set.
func (o *NonFungibleTokenMintWithTxId) GetEventIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventIndex, true
}

// SetEventIndex sets field value
func (o *NonFungibleTokenMintWithTxId) SetEventIndex(v int32) {
	o.EventIndex = v
}

// GetValue returns the Value field value
func (o *NonFungibleTokenMintWithTxId) GetValue() NonFungibleTokenHoldingWithTxIdValue {
	if o == nil {
		var ret NonFungibleTokenHoldingWithTxIdValue
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *NonFungibleTokenMintWithTxId) GetValueOk() (*NonFungibleTokenHoldingWithTxIdValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *NonFungibleTokenMintWithTxId) SetValue(v NonFungibleTokenHoldingWithTxIdValue) {
	o.Value = v
}

// GetTxId returns the TxId field value
func (o *NonFungibleTokenMintWithTxId) GetTxId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TxId
}

// GetTxIdOk returns a tuple with the TxId field value
// and a boolean to check if the value has been set.
func (o *NonFungibleTokenMintWithTxId) GetTxIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxId, true
}

// SetTxId sets field value
func (o *NonFungibleTokenMintWithTxId) SetTxId(v string) {
	o.TxId = v
}

func (o NonFungibleTokenMintWithTxId) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NonFungibleTokenMintWithTxId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Recipient) {
		toSerialize["recipient"] = o.Recipient
	}
	toSerialize["event_index"] = o.EventIndex
	toSerialize["value"] = o.Value
	toSerialize["tx_id"] = o.TxId
	return toSerialize, nil
}

func (o *NonFungibleTokenMintWithTxId) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"event_index",
		"value",
		"tx_id",
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

	varNonFungibleTokenMintWithTxId := _NonFungibleTokenMintWithTxId{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNonFungibleTokenMintWithTxId)

	if err != nil {
		return err
	}

	*o = NonFungibleTokenMintWithTxId(varNonFungibleTokenMintWithTxId)

	return err
}

type NullableNonFungibleTokenMintWithTxId struct {
	value *NonFungibleTokenMintWithTxId
	isSet bool
}

func (v NullableNonFungibleTokenMintWithTxId) Get() *NonFungibleTokenMintWithTxId {
	return v.value
}

func (v *NullableNonFungibleTokenMintWithTxId) Set(val *NonFungibleTokenMintWithTxId) {
	v.value = val
	v.isSet = true
}

func (v NullableNonFungibleTokenMintWithTxId) IsSet() bool {
	return v.isSet
}

func (v *NullableNonFungibleTokenMintWithTxId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNonFungibleTokenMintWithTxId(val *NonFungibleTokenMintWithTxId) *NullableNonFungibleTokenMintWithTxId {
	return &NullableNonFungibleTokenMintWithTxId{value: val, isSet: true}
}

func (v NullableNonFungibleTokenMintWithTxId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNonFungibleTokenMintWithTxId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
