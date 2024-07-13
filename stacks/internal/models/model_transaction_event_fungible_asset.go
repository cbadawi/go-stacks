package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the TransactionEventFungibleAsset type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &TransactionEventFungibleAsset{}

// TransactionEventFungibleAsset struct for TransactionEventFungibleAsset
type TransactionEventFungibleAsset struct {
	EventIndex int32                                   `json:"event_index"`
	EventType  string                                  `json:"event_type"`
	TxId       string                                  `json:"tx_id"`
	Asset      TransactionEventFungibleAssetAllOfAsset `json:"asset"`
}

type _TransactionEventFungibleAsset TransactionEventFungibleAsset

// NewTransactionEventFungibleAsset instantiates a new TransactionEventFungibleAsset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionEventFungibleAsset(eventIndex int32, eventType string, txId string, asset TransactionEventFungibleAssetAllOfAsset) *TransactionEventFungibleAsset {
	this := TransactionEventFungibleAsset{}
	this.EventIndex = eventIndex
	this.EventType = eventType
	this.TxId = txId
	this.Asset = asset
	return &this
}

// NewTransactionEventFungibleAssetWithDefaults instantiates a new TransactionEventFungibleAsset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionEventFungibleAssetWithDefaults() *TransactionEventFungibleAsset {
	this := TransactionEventFungibleAsset{}
	return &this
}

// GetEventIndex returns the EventIndex field value
func (o *TransactionEventFungibleAsset) GetEventIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EventIndex
}

// GetEventIndexOk returns a tuple with the EventIndex field value
// and a boolean to check if the value has been set.
func (o *TransactionEventFungibleAsset) GetEventIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventIndex, true
}

// SetEventIndex sets field value
func (o *TransactionEventFungibleAsset) SetEventIndex(v int32) {
	o.EventIndex = v
}

// GetEventType returns the EventType field value
func (o *TransactionEventFungibleAsset) GetEventType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *TransactionEventFungibleAsset) GetEventTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *TransactionEventFungibleAsset) SetEventType(v string) {
	o.EventType = v
}

// GetTxId returns the TxId field value
func (o *TransactionEventFungibleAsset) GetTxId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TxId
}

// GetTxIdOk returns a tuple with the TxId field value
// and a boolean to check if the value has been set.
func (o *TransactionEventFungibleAsset) GetTxIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxId, true
}

// SetTxId sets field value
func (o *TransactionEventFungibleAsset) SetTxId(v string) {
	o.TxId = v
}

// GetAsset returns the Asset field value
func (o *TransactionEventFungibleAsset) GetAsset() TransactionEventFungibleAssetAllOfAsset {
	if o == nil {
		var ret TransactionEventFungibleAssetAllOfAsset
		return ret
	}

	return o.Asset
}

// GetAssetOk returns a tuple with the Asset field value
// and a boolean to check if the value has been set.
func (o *TransactionEventFungibleAsset) GetAssetOk() (*TransactionEventFungibleAssetAllOfAsset, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Asset, true
}

// SetAsset sets field value
func (o *TransactionEventFungibleAsset) SetAsset(v TransactionEventFungibleAssetAllOfAsset) {
	o.Asset = v
}

func (o TransactionEventFungibleAsset) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionEventFungibleAsset) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["event_index"] = o.EventIndex
	toSerialize["event_type"] = o.EventType
	toSerialize["tx_id"] = o.TxId
	toSerialize["asset"] = o.Asset
	return toSerialize, nil
}

func (o *TransactionEventFungibleAsset) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"event_index",
		"event_type",
		"tx_id",
		"asset",
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

	varTransactionEventFungibleAsset := _TransactionEventFungibleAsset{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransactionEventFungibleAsset)

	if err != nil {
		return err
	}

	*o = TransactionEventFungibleAsset(varTransactionEventFungibleAsset)

	return err
}

type NullableTransactionEventFungibleAsset struct {
	value *TransactionEventFungibleAsset
	isSet bool
}

func (v NullableTransactionEventFungibleAsset) Get() *TransactionEventFungibleAsset {
	return v.value
}

func (v *NullableTransactionEventFungibleAsset) Set(val *TransactionEventFungibleAsset) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionEventFungibleAsset) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionEventFungibleAsset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionEventFungibleAsset(val *TransactionEventFungibleAsset) *NullableTransactionEventFungibleAsset {
	return &NullableTransactionEventFungibleAsset{value: val, isSet: true}
}

func (v NullableTransactionEventFungibleAsset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionEventFungibleAsset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
