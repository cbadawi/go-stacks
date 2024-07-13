package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the PostCoreNodeTransactionsError type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PostCoreNodeTransactionsError{}

// PostCoreNodeTransactionsError GET request that returns transactions
type PostCoreNodeTransactionsError struct {
	// The error
	Error string `json:"error"`
	// The reason for the error
	Reason string `json:"reason"`
	// More details about the reason
	ReasonData map[string]interface{} `json:"reason_data"`
	// The relevant transaction id
	Txid string `json:"txid"`
}

type _PostCoreNodeTransactionsError PostCoreNodeTransactionsError

// NewPostCoreNodeTransactionsError instantiates a new PostCoreNodeTransactionsError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostCoreNodeTransactionsError(error_ string, reason string, reasonData map[string]interface{}, txid string) *PostCoreNodeTransactionsError {
	this := PostCoreNodeTransactionsError{}
	this.Error = error_
	this.Reason = reason
	this.ReasonData = reasonData
	this.Txid = txid
	return &this
}

// NewPostCoreNodeTransactionsErrorWithDefaults instantiates a new PostCoreNodeTransactionsError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostCoreNodeTransactionsErrorWithDefaults() *PostCoreNodeTransactionsError {
	this := PostCoreNodeTransactionsError{}
	return &this
}

// GetError returns the Error field value
func (o *PostCoreNodeTransactionsError) GetError() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *PostCoreNodeTransactionsError) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *PostCoreNodeTransactionsError) SetError(v string) {
	o.Error = v
}

// GetReason returns the Reason field value
func (o *PostCoreNodeTransactionsError) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *PostCoreNodeTransactionsError) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *PostCoreNodeTransactionsError) SetReason(v string) {
	o.Reason = v
}

// GetReasonData returns the ReasonData field value
func (o *PostCoreNodeTransactionsError) GetReasonData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.ReasonData
}

// GetReasonDataOk returns a tuple with the ReasonData field value
// and a boolean to check if the value has been set.
func (o *PostCoreNodeTransactionsError) GetReasonDataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.ReasonData, true
}

// SetReasonData sets field value
func (o *PostCoreNodeTransactionsError) SetReasonData(v map[string]interface{}) {
	o.ReasonData = v
}

// GetTxid returns the Txid field value
func (o *PostCoreNodeTransactionsError) GetTxid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Txid
}

// GetTxidOk returns a tuple with the Txid field value
// and a boolean to check if the value has been set.
func (o *PostCoreNodeTransactionsError) GetTxidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Txid, true
}

// SetTxid sets field value
func (o *PostCoreNodeTransactionsError) SetTxid(v string) {
	o.Txid = v
}

func (o PostCoreNodeTransactionsError) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostCoreNodeTransactionsError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["error"] = o.Error
	toSerialize["reason"] = o.Reason
	toSerialize["reason_data"] = o.ReasonData
	toSerialize["txid"] = o.Txid
	return toSerialize, nil
}

func (o *PostCoreNodeTransactionsError) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"error",
		"reason",
		"reason_data",
		"txid",
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

	varPostCoreNodeTransactionsError := _PostCoreNodeTransactionsError{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostCoreNodeTransactionsError)

	if err != nil {
		return err
	}

	*o = PostCoreNodeTransactionsError(varPostCoreNodeTransactionsError)

	return err
}

type NullablePostCoreNodeTransactionsError struct {
	value *PostCoreNodeTransactionsError
	isSet bool
}

func (v NullablePostCoreNodeTransactionsError) Get() *PostCoreNodeTransactionsError {
	return v.value
}

func (v *NullablePostCoreNodeTransactionsError) Set(val *PostCoreNodeTransactionsError) {
	v.value = val
	v.isSet = true
}

func (v NullablePostCoreNodeTransactionsError) IsSet() bool {
	return v.isSet
}

func (v *NullablePostCoreNodeTransactionsError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostCoreNodeTransactionsError(val *PostCoreNodeTransactionsError) *NullablePostCoreNodeTransactionsError {
	return &NullablePostCoreNodeTransactionsError{value: val, isSet: true}
}

func (v NullablePostCoreNodeTransactionsError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostCoreNodeTransactionsError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
