package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the TxSearchResult type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &TxSearchResult{}

// TxSearchResult Transaction search result
type TxSearchResult struct {
	// Indicates if the requested object was found or not
	Found  bool                 `json:"found"`
	Result TxSearchResultResult `json:"result"`
}

type _TxSearchResult TxSearchResult

// NewTxSearchResult instantiates a new TxSearchResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTxSearchResult(found bool, result TxSearchResultResult) *TxSearchResult {
	this := TxSearchResult{}
	this.Found = found
	this.Result = result
	return &this
}

// NewTxSearchResultWithDefaults instantiates a new TxSearchResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTxSearchResultWithDefaults() *TxSearchResult {
	this := TxSearchResult{}
	var found bool = true
	this.Found = found
	return &this
}

// GetFound returns the Found field value
func (o *TxSearchResult) GetFound() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Found
}

// GetFoundOk returns a tuple with the Found field value
// and a boolean to check if the value has been set.
func (o *TxSearchResult) GetFoundOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Found, true
}

// SetFound sets field value
func (o *TxSearchResult) SetFound(v bool) {
	o.Found = v
}

// GetResult returns the Result field value
func (o *TxSearchResult) GetResult() TxSearchResultResult {
	if o == nil {
		var ret TxSearchResultResult
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *TxSearchResult) GetResultOk() (*TxSearchResultResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *TxSearchResult) SetResult(v TxSearchResultResult) {
	o.Result = v
}

func (o TxSearchResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TxSearchResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["found"] = o.Found
	toSerialize["result"] = o.Result
	return toSerialize, nil
}

func (o *TxSearchResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"found",
		"result",
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

	varTxSearchResult := _TxSearchResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTxSearchResult)

	if err != nil {
		return err
	}

	*o = TxSearchResult(varTxSearchResult)

	return err
}

type NullableTxSearchResult struct {
	value *TxSearchResult
	isSet bool
}

func (v NullableTxSearchResult) Get() *TxSearchResult {
	return v.value
}

func (v *NullableTxSearchResult) Set(val *TxSearchResult) {
	v.value = val
	v.isSet = true
}

func (v NullableTxSearchResult) IsSet() bool {
	return v.isSet
}

func (v *NullableTxSearchResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTxSearchResult(val *TxSearchResult) *NullableTxSearchResult {
	return &NullableTxSearchResult{value: val, isSet: true}
}

func (v NullableTxSearchResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTxSearchResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
