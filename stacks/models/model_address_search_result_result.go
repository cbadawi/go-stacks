package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the AddressSearchResultResult type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AddressSearchResultResult{}

// AddressSearchResultResult This object carries the search result
type AddressSearchResultResult struct {
	// The id used to search this query.
	EntityId   string                     `json:"entity_id"`
	EntityType string                     `json:"entity_type"`
	Metadata   *AddressStxBalanceResponse `json:"metadata,omitempty"`
}

type _AddressSearchResultResult AddressSearchResultResult

// NewAddressSearchResultResult instantiates a new AddressSearchResultResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressSearchResultResult(entityId string, entityType string) *AddressSearchResultResult {
	this := AddressSearchResultResult{}
	this.EntityId = entityId
	this.EntityType = entityType
	return &this
}

// NewAddressSearchResultResultWithDefaults instantiates a new AddressSearchResultResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressSearchResultResultWithDefaults() *AddressSearchResultResult {
	this := AddressSearchResultResult{}
	return &this
}

// GetEntityId returns the EntityId field value
func (o *AddressSearchResultResult) GetEntityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value
// and a boolean to check if the value has been set.
func (o *AddressSearchResultResult) GetEntityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityId, true
}

// SetEntityId sets field value
func (o *AddressSearchResultResult) SetEntityId(v string) {
	o.EntityId = v
}

// GetEntityType returns the EntityType field value
func (o *AddressSearchResultResult) GetEntityType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value
// and a boolean to check if the value has been set.
func (o *AddressSearchResultResult) GetEntityTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityType, true
}

// SetEntityType sets field value
func (o *AddressSearchResultResult) SetEntityType(v string) {
	o.EntityType = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *AddressSearchResultResult) GetMetadata() AddressStxBalanceResponse {
	if o == nil || utils.IsNil(o.Metadata) {
		var ret AddressStxBalanceResponse
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressSearchResultResult) GetMetadataOk() (*AddressStxBalanceResponse, bool) {
	if o == nil || utils.IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *AddressSearchResultResult) HasMetadata() bool {
	if o != nil && !utils.IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given AddressStxBalanceResponse and assigns it to the Metadata field.
func (o *AddressSearchResultResult) SetMetadata(v AddressStxBalanceResponse) {
	o.Metadata = &v
}

func (o AddressSearchResultResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressSearchResultResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["entity_id"] = o.EntityId
	toSerialize["entity_type"] = o.EntityType
	if !utils.IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

func (o *AddressSearchResultResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"entity_id",
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

	varAddressSearchResultResult := _AddressSearchResultResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddressSearchResultResult)

	if err != nil {
		return err
	}

	*o = AddressSearchResultResult(varAddressSearchResultResult)

	return err
}

type NullableAddressSearchResultResult struct {
	value *AddressSearchResultResult
	isSet bool
}

func (v NullableAddressSearchResultResult) Get() *AddressSearchResultResult {
	return v.value
}

func (v *NullableAddressSearchResultResult) Set(val *AddressSearchResultResult) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressSearchResultResult) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressSearchResultResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressSearchResultResult(val *AddressSearchResultResult) *NullableAddressSearchResultResult {
	return &NullableAddressSearchResultResult{value: val, isSet: true}
}

func (v NullableAddressSearchResultResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressSearchResultResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
