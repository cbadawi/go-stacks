package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the TransactionEventNonFungibleAsset type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &TransactionEventNonFungibleAsset{}

// TransactionEventNonFungibleAsset struct for TransactionEventNonFungibleAsset
type TransactionEventNonFungibleAsset struct {
	EventIndex int32                                      `json:"event_index"`
	EventType  string                                     `json:"event_type"`
	TxId       string                                     `json:"tx_id"`
	Asset      TransactionEventNonFungibleAssetAllOfAsset `json:"asset"`
}

type _TransactionEventNonFungibleAsset TransactionEventNonFungibleAsset

// NewTransactionEventNonFungibleAsset instantiates a new TransactionEventNonFungibleAsset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionEventNonFungibleAsset(eventIndex int32, eventType string, txId string, asset TransactionEventNonFungibleAssetAllOfAsset) *TransactionEventNonFungibleAsset {
	this := TransactionEventNonFungibleAsset{}
	this.EventIndex = eventIndex
	this.EventType = eventType
	this.TxId = txId
	this.Asset = asset
	return &this
}

// NewTransactionEventNonFungibleAssetWithDefaults instantiates a new TransactionEventNonFungibleAsset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionEventNonFungibleAssetWithDefaults() *TransactionEventNonFungibleAsset {
	this := TransactionEventNonFungibleAsset{}
	return &this
}

// GetEventIndex returns the EventIndex field value
func (o *TransactionEventNonFungibleAsset) GetEventIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EventIndex
}

// GetEventIndexOk returns a tuple with the EventIndex field value
// and a boolean to check if the value has been set.
func (o *TransactionEventNonFungibleAsset) GetEventIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventIndex, true
}

// SetEventIndex sets field value
func (o *TransactionEventNonFungibleAsset) SetEventIndex(v int32) {
	o.EventIndex = v
}

// GetEventType returns the EventType field value
func (o *TransactionEventNonFungibleAsset) GetEventType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *TransactionEventNonFungibleAsset) GetEventTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *TransactionEventNonFungibleAsset) SetEventType(v string) {
	o.EventType = v
}

// GetTxId returns the TxId field value
func (o *TransactionEventNonFungibleAsset) GetTxId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TxId
}

// GetTxIdOk returns a tuple with the TxId field value
// and a boolean to check if the value has been set.
func (o *TransactionEventNonFungibleAsset) GetTxIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxId, true
}

// SetTxId sets field value
func (o *TransactionEventNonFungibleAsset) SetTxId(v string) {
	o.TxId = v
}

// GetAsset returns the Asset field value
func (o *TransactionEventNonFungibleAsset) GetAsset() TransactionEventNonFungibleAssetAllOfAsset {
	if o == nil {
		var ret TransactionEventNonFungibleAssetAllOfAsset
		return ret
	}

	return o.Asset
}

// GetAssetOk returns a tuple with the Asset field value
// and a boolean to check if the value has been set.
func (o *TransactionEventNonFungibleAsset) GetAssetOk() (*TransactionEventNonFungibleAssetAllOfAsset, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Asset, true
}

// SetAsset sets field value
func (o *TransactionEventNonFungibleAsset) SetAsset(v TransactionEventNonFungibleAssetAllOfAsset) {
	o.Asset = v
}

func (o TransactionEventNonFungibleAsset) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionEventNonFungibleAsset) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["event_index"] = o.EventIndex
	toSerialize["event_type"] = o.EventType
	toSerialize["tx_id"] = o.TxId
	toSerialize["asset"] = o.Asset
	return toSerialize, nil
}

func (o *TransactionEventNonFungibleAsset) UnmarshalJSON(data []byte) (err error) {
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

	varTransactionEventNonFungibleAsset := _TransactionEventNonFungibleAsset{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransactionEventNonFungibleAsset)

	if err != nil {
		return err
	}

	*o = TransactionEventNonFungibleAsset(varTransactionEventNonFungibleAsset)

	return err
}

type NullableTransactionEventNonFungibleAsset struct {
	value *TransactionEventNonFungibleAsset
	isSet bool
}

func (v NullableTransactionEventNonFungibleAsset) Get() *TransactionEventNonFungibleAsset {
	return v.value
}

func (v *NullableTransactionEventNonFungibleAsset) Set(val *TransactionEventNonFungibleAsset) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionEventNonFungibleAsset) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionEventNonFungibleAsset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionEventNonFungibleAsset(val *TransactionEventNonFungibleAsset) *NullableTransactionEventNonFungibleAsset {
	return &NullableTransactionEventNonFungibleAsset{value: val, isSet: true}
}

func (v NullableTransactionEventNonFungibleAsset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionEventNonFungibleAsset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
