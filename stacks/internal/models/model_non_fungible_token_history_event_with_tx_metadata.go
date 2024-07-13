package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the NonFungibleTokenHistoryEventWithTxMetadata type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &NonFungibleTokenHistoryEventWithTxMetadata{}

// NonFungibleTokenHistoryEventWithTxMetadata Non-Fungible Token history event with transaction metadata
type NonFungibleTokenHistoryEventWithTxMetadata struct {
	Sender         utils.NullableString `json:"sender,omitempty"`
	Recipient      *string              `json:"recipient,omitempty"`
	EventIndex     int32                `json:"event_index"`
	AssetEventType string               `json:"asset_event_type"`
	Tx             Transaction          `json:"tx"`
}

type _NonFungibleTokenHistoryEventWithTxMetadata NonFungibleTokenHistoryEventWithTxMetadata

// NewNonFungibleTokenHistoryEventWithTxMetadata instantiates a new NonFungibleTokenHistoryEventWithTxMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNonFungibleTokenHistoryEventWithTxMetadata(eventIndex int32, assetEventType string, tx Transaction) *NonFungibleTokenHistoryEventWithTxMetadata {
	this := NonFungibleTokenHistoryEventWithTxMetadata{}
	this.EventIndex = eventIndex
	this.AssetEventType = assetEventType
	this.Tx = tx
	return &this
}

// NewNonFungibleTokenHistoryEventWithTxMetadataWithDefaults instantiates a new NonFungibleTokenHistoryEventWithTxMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNonFungibleTokenHistoryEventWithTxMetadataWithDefaults() *NonFungibleTokenHistoryEventWithTxMetadata {
	this := NonFungibleTokenHistoryEventWithTxMetadata{}
	return &this
}

// GetSender returns the Sender field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NonFungibleTokenHistoryEventWithTxMetadata) GetSender() string {
	if o == nil || utils.IsNil(o.Sender.Get()) {
		var ret string
		return ret
	}
	return *o.Sender.Get()
}

// GetSenderOk returns a tuple with the Sender field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NonFungibleTokenHistoryEventWithTxMetadata) GetSenderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sender.Get(), o.Sender.IsSet()
}

// HasSender returns a boolean if a field has been set.
func (o *NonFungibleTokenHistoryEventWithTxMetadata) HasSender() bool {
	if o != nil && o.Sender.IsSet() {
		return true
	}

	return false
}

// SetSender gets a reference to the given NullableString and assigns it to the Sender field.
func (o *NonFungibleTokenHistoryEventWithTxMetadata) SetSender(v string) {
	o.Sender.Set(&v)
}

// SetSenderNil sets the value for Sender to be an explicit nil
func (o *NonFungibleTokenHistoryEventWithTxMetadata) SetSenderNil() {
	o.Sender.Set(nil)
}

// UnsetSender ensures that no value is present for Sender, not even an explicit nil
func (o *NonFungibleTokenHistoryEventWithTxMetadata) UnsetSender() {
	o.Sender.Unset()
}

// GetRecipient returns the Recipient field value if set, zero value otherwise.
func (o *NonFungibleTokenHistoryEventWithTxMetadata) GetRecipient() string {
	if o == nil || utils.IsNil(o.Recipient) {
		var ret string
		return ret
	}
	return *o.Recipient
}

// GetRecipientOk returns a tuple with the Recipient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonFungibleTokenHistoryEventWithTxMetadata) GetRecipientOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Recipient) {
		return nil, false
	}
	return o.Recipient, true
}

// HasRecipient returns a boolean if a field has been set.
func (o *NonFungibleTokenHistoryEventWithTxMetadata) HasRecipient() bool {
	if o != nil && !utils.IsNil(o.Recipient) {
		return true
	}

	return false
}

// SetRecipient gets a reference to the given string and assigns it to the Recipient field.
func (o *NonFungibleTokenHistoryEventWithTxMetadata) SetRecipient(v string) {
	o.Recipient = &v
}

// GetEventIndex returns the EventIndex field value
func (o *NonFungibleTokenHistoryEventWithTxMetadata) GetEventIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EventIndex
}

// GetEventIndexOk returns a tuple with the EventIndex field value
// and a boolean to check if the value has been set.
func (o *NonFungibleTokenHistoryEventWithTxMetadata) GetEventIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventIndex, true
}

// SetEventIndex sets field value
func (o *NonFungibleTokenHistoryEventWithTxMetadata) SetEventIndex(v int32) {
	o.EventIndex = v
}

// GetAssetEventType returns the AssetEventType field value
func (o *NonFungibleTokenHistoryEventWithTxMetadata) GetAssetEventType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetEventType
}

// GetAssetEventTypeOk returns a tuple with the AssetEventType field value
// and a boolean to check if the value has been set.
func (o *NonFungibleTokenHistoryEventWithTxMetadata) GetAssetEventTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetEventType, true
}

// SetAssetEventType sets field value
func (o *NonFungibleTokenHistoryEventWithTxMetadata) SetAssetEventType(v string) {
	o.AssetEventType = v
}

// GetTx returns the Tx field value
func (o *NonFungibleTokenHistoryEventWithTxMetadata) GetTx() Transaction {
	if o == nil {
		var ret Transaction
		return ret
	}

	return o.Tx
}

// GetTxOk returns a tuple with the Tx field value
// and a boolean to check if the value has been set.
func (o *NonFungibleTokenHistoryEventWithTxMetadata) GetTxOk() (*Transaction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tx, true
}

// SetTx sets field value
func (o *NonFungibleTokenHistoryEventWithTxMetadata) SetTx(v Transaction) {
	o.Tx = v
}

func (o NonFungibleTokenHistoryEventWithTxMetadata) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NonFungibleTokenHistoryEventWithTxMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Sender.IsSet() {
		toSerialize["sender"] = o.Sender.Get()
	}
	if !utils.IsNil(o.Recipient) {
		toSerialize["recipient"] = o.Recipient
	}
	toSerialize["event_index"] = o.EventIndex
	toSerialize["asset_event_type"] = o.AssetEventType
	toSerialize["tx"] = o.Tx
	return toSerialize, nil
}

func (o *NonFungibleTokenHistoryEventWithTxMetadata) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"event_index",
		"asset_event_type",
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

	varNonFungibleTokenHistoryEventWithTxMetadata := _NonFungibleTokenHistoryEventWithTxMetadata{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNonFungibleTokenHistoryEventWithTxMetadata)

	if err != nil {
		return err
	}

	*o = NonFungibleTokenHistoryEventWithTxMetadata(varNonFungibleTokenHistoryEventWithTxMetadata)

	return err
}

type NullableNonFungibleTokenHistoryEventWithTxMetadata struct {
	value *NonFungibleTokenHistoryEventWithTxMetadata
	isSet bool
}

func (v NullableNonFungibleTokenHistoryEventWithTxMetadata) Get() *NonFungibleTokenHistoryEventWithTxMetadata {
	return v.value
}

func (v *NullableNonFungibleTokenHistoryEventWithTxMetadata) Set(val *NonFungibleTokenHistoryEventWithTxMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableNonFungibleTokenHistoryEventWithTxMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableNonFungibleTokenHistoryEventWithTxMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNonFungibleTokenHistoryEventWithTxMetadata(val *NonFungibleTokenHistoryEventWithTxMetadata) *NullableNonFungibleTokenHistoryEventWithTxMetadata {
	return &NullableNonFungibleTokenHistoryEventWithTxMetadata{value: val, isSet: true}
}

func (v NullableNonFungibleTokenHistoryEventWithTxMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNonFungibleTokenHistoryEventWithTxMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
