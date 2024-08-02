package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the ContractSearchResult type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ContractSearchResult{}

// ContractSearchResult Contract search result
type ContractSearchResult struct {
	// Indicates if the requested object was found or not
	Found  bool                       `json:"found"`
	Result ContractSearchResultResult `json:"result"`
}

type _ContractSearchResult ContractSearchResult

// NewContractSearchResult instantiates a new ContractSearchResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContractSearchResult(found bool, result ContractSearchResultResult) *ContractSearchResult {
	this := ContractSearchResult{}
	this.Found = found
	this.Result = result
	return &this
}

// NewContractSearchResultWithDefaults instantiates a new ContractSearchResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractSearchResultWithDefaults() *ContractSearchResult {
	this := ContractSearchResult{}
	var found bool = true
	this.Found = found
	return &this
}

// GetFound returns the Found field value
func (o *ContractSearchResult) GetFound() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Found
}

// GetFoundOk returns a tuple with the Found field value
// and a boolean to check if the value has been set.
func (o *ContractSearchResult) GetFoundOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Found, true
}

// SetFound sets field value
func (o *ContractSearchResult) SetFound(v bool) {
	o.Found = v
}

// GetResult returns the Result field value
func (o *ContractSearchResult) GetResult() ContractSearchResultResult {
	if o == nil {
		var ret ContractSearchResultResult
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *ContractSearchResult) GetResultOk() (*ContractSearchResultResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *ContractSearchResult) SetResult(v ContractSearchResultResult) {
	o.Result = v
}

func (o ContractSearchResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContractSearchResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["found"] = o.Found
	toSerialize["result"] = o.Result
	return toSerialize, nil
}

func (o *ContractSearchResult) UnmarshalJSON(data []byte) (err error) {
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

	varContractSearchResult := _ContractSearchResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varContractSearchResult)

	if err != nil {
		return err
	}

	*o = ContractSearchResult(varContractSearchResult)

	return err
}

type NullableContractSearchResult struct {
	value *ContractSearchResult
	isSet bool
}

func (v NullableContractSearchResult) Get() *ContractSearchResult {
	return v.value
}

func (v *NullableContractSearchResult) Set(val *ContractSearchResult) {
	v.value = val
	v.isSet = true
}

func (v NullableContractSearchResult) IsSet() bool {
	return v.isSet
}

func (v *NullableContractSearchResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractSearchResult(val *ContractSearchResult) *NullableContractSearchResult {
	return &NullableContractSearchResult{value: val, isSet: true}
}

func (v NullableContractSearchResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractSearchResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
