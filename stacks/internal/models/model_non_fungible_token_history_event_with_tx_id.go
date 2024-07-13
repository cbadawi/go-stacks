package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the NonFungibleTokenHistoryEventWithTxId type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &NonFungibleTokenHistoryEventWithTxId{}

// NonFungibleTokenHistoryEventWithTxId Non-Fungible Token history event with transaction id
type NonFungibleTokenHistoryEventWithTxId struct {
	Sender         utils.NullableString `json:"sender,omitempty"`
	Recipient      *string              `json:"recipient,omitempty"`
	EventIndex     int32                `json:"event_index"`
	AssetEventType string               `json:"asset_event_type"`
	TxId           string               `json:"tx_id"`
}

type _NonFungibleTokenHistoryEventWithTxId NonFungibleTokenHistoryEventWithTxId

// NewNonFungibleTokenHistoryEventWithTxId instantiates a new NonFungibleTokenHistoryEventWithTxId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNonFungibleTokenHistoryEventWithTxId(eventIndex int32, assetEventType string, txId string) *NonFungibleTokenHistoryEventWithTxId {
	this := NonFungibleTokenHistoryEventWithTxId{}
	this.EventIndex = eventIndex
	this.AssetEventType = assetEventType
	this.TxId = txId
	return &this
}

// NewNonFungibleTokenHistoryEventWithTxIdWithDefaults instantiates a new NonFungibleTokenHistoryEventWithTxId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNonFungibleTokenHistoryEventWithTxIdWithDefaults() *NonFungibleTokenHistoryEventWithTxId {
	this := NonFungibleTokenHistoryEventWithTxId{}
	return &this
}

// GetSender returns the Sender field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NonFungibleTokenHistoryEventWithTxId) GetSender() string {
	if o == nil || utils.IsNil(o.Sender.Get()) {
		var ret string
		return ret
	}
	return *o.Sender.Get()
}

// GetSenderOk returns a tuple with the Sender field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NonFungibleTokenHistoryEventWithTxId) GetSenderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sender.Get(), o.Sender.IsSet()
}

// HasSender returns a boolean if a field has been set.
func (o *NonFungibleTokenHistoryEventWithTxId) HasSender() bool {
	if o != nil && o.Sender.IsSet() {
		return true
	}

	return false
}

// SetSender gets a reference to the given NullableString and assigns it to the Sender field.
func (o *NonFungibleTokenHistoryEventWithTxId) SetSender(v string) {
	o.Sender.Set(&v)
}

// SetSenderNil sets the value for Sender to be an explicit nil
func (o *NonFungibleTokenHistoryEventWithTxId) SetSenderNil() {
	o.Sender.Set(nil)
}

// UnsetSender ensures that no value is present for Sender, not even an explicit nil
func (o *NonFungibleTokenHistoryEventWithTxId) UnsetSender() {
	o.Sender.Unset()
}

// GetRecipient returns the Recipient field value if set, zero value otherwise.
func (o *NonFungibleTokenHistoryEventWithTxId) GetRecipient() string {
	if o == nil || utils.IsNil(o.Recipient) {
		var ret string
		return ret
	}
	return *o.Recipient
}

// GetRecipientOk returns a tuple with the Recipient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonFungibleTokenHistoryEventWithTxId) GetRecipientOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Recipient) {
		return nil, false
	}
	return o.Recipient, true
}

// HasRecipient returns a boolean if a field has been set.
func (o *NonFungibleTokenHistoryEventWithTxId) HasRecipient() bool {
	if o != nil && !utils.IsNil(o.Recipient) {
		return true
	}

	return false
}

// SetRecipient gets a reference to the given string and assigns it to the Recipient field.
func (o *NonFungibleTokenHistoryEventWithTxId) SetRecipient(v string) {
	o.Recipient = &v
}

// GetEventIndex returns the EventIndex field value
func (o *NonFungibleTokenHistoryEventWithTxId) GetEventIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EventIndex
}

// GetEventIndexOk returns a tuple with the EventIndex field value
// and a boolean to check if the value has been set.
func (o *NonFungibleTokenHistoryEventWithTxId) GetEventIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventIndex, true
}

// SetEventIndex sets field value
func (o *NonFungibleTokenHistoryEventWithTxId) SetEventIndex(v int32) {
	o.EventIndex = v
}

// GetAssetEventType returns the AssetEventType field value
func (o *NonFungibleTokenHistoryEventWithTxId) GetAssetEventType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetEventType
}

// GetAssetEventTypeOk returns a tuple with the AssetEventType field value
// and a boolean to check if the value has been set.
func (o *NonFungibleTokenHistoryEventWithTxId) GetAssetEventTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetEventType, true
}

// SetAssetEventType sets field value
func (o *NonFungibleTokenHistoryEventWithTxId) SetAssetEventType(v string) {
	o.AssetEventType = v
}

// GetTxId returns the TxId field value
func (o *NonFungibleTokenHistoryEventWithTxId) GetTxId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TxId
}

// GetTxIdOk returns a tuple with the TxId field value
// and a boolean to check if the value has been set.
func (o *NonFungibleTokenHistoryEventWithTxId) GetTxIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxId, true
}

// SetTxId sets field value
func (o *NonFungibleTokenHistoryEventWithTxId) SetTxId(v string) {
	o.TxId = v
}

func (o NonFungibleTokenHistoryEventWithTxId) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NonFungibleTokenHistoryEventWithTxId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Sender.IsSet() {
		toSerialize["sender"] = o.Sender.Get()
	}
	if !utils.IsNil(o.Recipient) {
		toSerialize["recipient"] = o.Recipient
	}
	toSerialize["event_index"] = o.EventIndex
	toSerialize["asset_event_type"] = o.AssetEventType
	toSerialize["tx_id"] = o.TxId
	return toSerialize, nil
}

func (o *NonFungibleTokenHistoryEventWithTxId) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"event_index",
		"asset_event_type",
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

	varNonFungibleTokenHistoryEventWithTxId := _NonFungibleTokenHistoryEventWithTxId{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNonFungibleTokenHistoryEventWithTxId)

	if err != nil {
		return err
	}

	*o = NonFungibleTokenHistoryEventWithTxId(varNonFungibleTokenHistoryEventWithTxId)

	return err
}

type NullableNonFungibleTokenHistoryEventWithTxId struct {
	value *NonFungibleTokenHistoryEventWithTxId
	isSet bool
}

func (v NullableNonFungibleTokenHistoryEventWithTxId) Get() *NonFungibleTokenHistoryEventWithTxId {
	return v.value
}

func (v *NullableNonFungibleTokenHistoryEventWithTxId) Set(val *NonFungibleTokenHistoryEventWithTxId) {
	v.value = val
	v.isSet = true
}

func (v NullableNonFungibleTokenHistoryEventWithTxId) IsSet() bool {
	return v.isSet
}

func (v *NullableNonFungibleTokenHistoryEventWithTxId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNonFungibleTokenHistoryEventWithTxId(val *NonFungibleTokenHistoryEventWithTxId) *NullableNonFungibleTokenHistoryEventWithTxId {
	return &NullableNonFungibleTokenHistoryEventWithTxId{value: val, isSet: true}
}

func (v NullableNonFungibleTokenHistoryEventWithTxId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNonFungibleTokenHistoryEventWithTxId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
