package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the BnsNamesOwnByAddressResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &BnsNamesOwnByAddressResponse{}

// BnsNamesOwnByAddressResponse Retrieves a list of names owned by the address provided.
type BnsNamesOwnByAddressResponse struct {
	Names []string `json:"names"`
}

type _BnsNamesOwnByAddressResponse BnsNamesOwnByAddressResponse

// NewBnsNamesOwnByAddressResponse instantiates a new BnsNamesOwnByAddressResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBnsNamesOwnByAddressResponse(names []string) *BnsNamesOwnByAddressResponse {
	this := BnsNamesOwnByAddressResponse{}
	this.Names = names
	return &this
}

// NewBnsNamesOwnByAddressResponseWithDefaults instantiates a new BnsNamesOwnByAddressResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBnsNamesOwnByAddressResponseWithDefaults() *BnsNamesOwnByAddressResponse {
	this := BnsNamesOwnByAddressResponse{}
	return &this
}

// GetNames returns the Names field value
func (o *BnsNamesOwnByAddressResponse) GetNames() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Names
}

// GetNamesOk returns a tuple with the Names field value
// and a boolean to check if the value has been set.
func (o *BnsNamesOwnByAddressResponse) GetNamesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Names, true
}

// SetNames sets field value
func (o *BnsNamesOwnByAddressResponse) SetNames(v []string) {
	o.Names = v
}

func (o BnsNamesOwnByAddressResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BnsNamesOwnByAddressResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["names"] = o.Names
	return toSerialize, nil
}

func (o *BnsNamesOwnByAddressResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"names",
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

	varBnsNamesOwnByAddressResponse := _BnsNamesOwnByAddressResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBnsNamesOwnByAddressResponse)

	if err != nil {
		return err
	}

	*o = BnsNamesOwnByAddressResponse(varBnsNamesOwnByAddressResponse)

	return err
}

type NullableBnsNamesOwnByAddressResponse struct {
	value *BnsNamesOwnByAddressResponse
	isSet bool
}

func (v NullableBnsNamesOwnByAddressResponse) Get() *BnsNamesOwnByAddressResponse {
	return v.value
}

func (v *NullableBnsNamesOwnByAddressResponse) Set(val *BnsNamesOwnByAddressResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBnsNamesOwnByAddressResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBnsNamesOwnByAddressResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBnsNamesOwnByAddressResponse(val *BnsNamesOwnByAddressResponse) *NullableBnsNamesOwnByAddressResponse {
	return &NullableBnsNamesOwnByAddressResponse{value: val, isSet: true}
}

func (v NullableBnsNamesOwnByAddressResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBnsNamesOwnByAddressResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
