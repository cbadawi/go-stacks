package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the SearchErrorResultResult type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &SearchErrorResultResult{}

// SearchErrorResultResult struct for SearchErrorResultResult
type SearchErrorResultResult struct {
	// Shows the currenty category of entity it is searched in.
	EntityType string `json:"entity_type"`
}

type _SearchErrorResultResult SearchErrorResultResult

// NewSearchErrorResultResult instantiates a new SearchErrorResultResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchErrorResultResult(entityType string) *SearchErrorResultResult {
	this := SearchErrorResultResult{}
	this.EntityType = entityType
	return &this
}

// NewSearchErrorResultResultWithDefaults instantiates a new SearchErrorResultResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchErrorResultResultWithDefaults() *SearchErrorResultResult {
	this := SearchErrorResultResult{}
	return &this
}

// GetEntityType returns the EntityType field value
func (o *SearchErrorResultResult) GetEntityType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value
// and a boolean to check if the value has been set.
func (o *SearchErrorResultResult) GetEntityTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityType, true
}

// SetEntityType sets field value
func (o *SearchErrorResultResult) SetEntityType(v string) {
	o.EntityType = v
}

func (o SearchErrorResultResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchErrorResultResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["entity_type"] = o.EntityType
	return toSerialize, nil
}

func (o *SearchErrorResultResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"entity_type",
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

	varSearchErrorResultResult := _SearchErrorResultResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSearchErrorResultResult)

	if err != nil {
		return err
	}

	*o = SearchErrorResultResult(varSearchErrorResultResult)

	return err
}

type NullableSearchErrorResultResult struct {
	value *SearchErrorResultResult
	isSet bool
}

func (v NullableSearchErrorResultResult) Get() *SearchErrorResultResult {
	return v.value
}

func (v *NullableSearchErrorResultResult) Set(val *SearchErrorResultResult) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchErrorResultResult) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchErrorResultResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchErrorResultResult(val *SearchErrorResultResult) *NullableSearchErrorResultResult {
	return &NullableSearchErrorResultResult{value: val, isSet: true}
}

func (v NullableSearchErrorResultResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchErrorResultResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
