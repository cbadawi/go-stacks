package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the MempoolTxSearchResult type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &MempoolTxSearchResult{}

// MempoolTxSearchResult Contract search result
type MempoolTxSearchResult struct {
	// Indicates if the requested object was found or not
	Found  bool                        `json:"found"`
	Result MempoolTxSearchResultResult `json:"result"`
}

type _MempoolTxSearchResult MempoolTxSearchResult

// NewMempoolTxSearchResult instantiates a new MempoolTxSearchResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMempoolTxSearchResult(found bool, result MempoolTxSearchResultResult) *MempoolTxSearchResult {
	this := MempoolTxSearchResult{}
	this.Found = found
	this.Result = result
	return &this
}

// NewMempoolTxSearchResultWithDefaults instantiates a new MempoolTxSearchResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMempoolTxSearchResultWithDefaults() *MempoolTxSearchResult {
	this := MempoolTxSearchResult{}
	var found bool = true
	this.Found = found
	return &this
}

// GetFound returns the Found field value
func (o *MempoolTxSearchResult) GetFound() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Found
}

// GetFoundOk returns a tuple with the Found field value
// and a boolean to check if the value has been set.
func (o *MempoolTxSearchResult) GetFoundOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Found, true
}

// SetFound sets field value
func (o *MempoolTxSearchResult) SetFound(v bool) {
	o.Found = v
}

// GetResult returns the Result field value
func (o *MempoolTxSearchResult) GetResult() MempoolTxSearchResultResult {
	if o == nil {
		var ret MempoolTxSearchResultResult
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *MempoolTxSearchResult) GetResultOk() (*MempoolTxSearchResultResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *MempoolTxSearchResult) SetResult(v MempoolTxSearchResultResult) {
	o.Result = v
}

func (o MempoolTxSearchResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MempoolTxSearchResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["found"] = o.Found
	toSerialize["result"] = o.Result
	return toSerialize, nil
}

func (o *MempoolTxSearchResult) UnmarshalJSON(data []byte) (err error) {
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

	varMempoolTxSearchResult := _MempoolTxSearchResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMempoolTxSearchResult)

	if err != nil {
		return err
	}

	*o = MempoolTxSearchResult(varMempoolTxSearchResult)

	return err
}

type NullableMempoolTxSearchResult struct {
	value *MempoolTxSearchResult
	isSet bool
}

func (v NullableMempoolTxSearchResult) Get() *MempoolTxSearchResult {
	return v.value
}

func (v *NullableMempoolTxSearchResult) Set(val *MempoolTxSearchResult) {
	v.value = val
	v.isSet = true
}

func (v NullableMempoolTxSearchResult) IsSet() bool {
	return v.isSet
}

func (v *NullableMempoolTxSearchResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMempoolTxSearchResult(val *MempoolTxSearchResult) *NullableMempoolTxSearchResult {
	return &NullableMempoolTxSearchResult{value: val, isSet: true}
}

func (v NullableMempoolTxSearchResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMempoolTxSearchResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
