package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the AddressTransaction type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AddressTransaction{}

// AddressTransaction Address transaction with STX, FT and NFT transfer summaries
type AddressTransaction struct {
	Tx Transaction `json:"tx"`
	// Total sent from the given address, including the tx fee, in micro-STX as an integer string.
	StxSent string `json:"stx_sent"`
	// Total received by the given address in micro-STX as an integer string.
	StxReceived string                    `json:"stx_received"`
	Events      *AddressTransactionEvents `json:"events,omitempty"`
}

type _AddressTransaction AddressTransaction

// NewAddressTransaction instantiates a new AddressTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressTransaction(tx Transaction, stxSent string, stxReceived string) *AddressTransaction {
	this := AddressTransaction{}
	this.Tx = tx
	this.StxSent = stxSent
	this.StxReceived = stxReceived
	return &this
}

// NewAddressTransactionWithDefaults instantiates a new AddressTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressTransactionWithDefaults() *AddressTransaction {
	this := AddressTransaction{}
	return &this
}

// GetTx returns the Tx field value
func (o *AddressTransaction) GetTx() Transaction {
	if o == nil {
		var ret Transaction
		return ret
	}

	return o.Tx
}

// GetTxOk returns a tuple with the Tx field value
// and a boolean to check if the value has been set.
func (o *AddressTransaction) GetTxOk() (*Transaction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tx, true
}

// SetTx sets field value
func (o *AddressTransaction) SetTx(v Transaction) {
	o.Tx = v
}

// GetStxSent returns the StxSent field value
func (o *AddressTransaction) GetStxSent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StxSent
}

// GetStxSentOk returns a tuple with the StxSent field value
// and a boolean to check if the value has been set.
func (o *AddressTransaction) GetStxSentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StxSent, true
}

// SetStxSent sets field value
func (o *AddressTransaction) SetStxSent(v string) {
	o.StxSent = v
}

// GetStxReceived returns the StxReceived field value
func (o *AddressTransaction) GetStxReceived() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StxReceived
}

// GetStxReceivedOk returns a tuple with the StxReceived field value
// and a boolean to check if the value has been set.
func (o *AddressTransaction) GetStxReceivedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StxReceived, true
}

// SetStxReceived sets field value
func (o *AddressTransaction) SetStxReceived(v string) {
	o.StxReceived = v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *AddressTransaction) GetEvents() AddressTransactionEvents {
	if o == nil || utils.IsNil(o.Events) {
		var ret AddressTransactionEvents
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressTransaction) GetEventsOk() (*AddressTransactionEvents, bool) {
	if o == nil || utils.IsNil(o.Events) {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *AddressTransaction) HasEvents() bool {
	if o != nil && !utils.IsNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given AddressTransactionEvents and assigns it to the Events field.
func (o *AddressTransaction) SetEvents(v AddressTransactionEvents) {
	o.Events = &v
}

func (o AddressTransaction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressTransaction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tx"] = o.Tx
	toSerialize["stx_sent"] = o.StxSent
	toSerialize["stx_received"] = o.StxReceived
	if !utils.IsNil(o.Events) {
		toSerialize["events"] = o.Events
	}
	return toSerialize, nil
}

func (o *AddressTransaction) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tx",
		"stx_sent",
		"stx_received",
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

	varAddressTransaction := _AddressTransaction{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddressTransaction)

	if err != nil {
		return err
	}

	*o = AddressTransaction(varAddressTransaction)

	return err
}

type NullableAddressTransaction struct {
	value *AddressTransaction
	isSet bool
}

func (v NullableAddressTransaction) Get() *AddressTransaction {
	return v.value
}

func (v *NullableAddressTransaction) Set(val *AddressTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressTransaction(val *AddressTransaction) *NullableAddressTransaction {
	return &NullableAddressTransaction{value: val, isSet: true}
}

func (v NullableAddressTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
